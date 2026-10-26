#!/usr/bin/env bash
# Tests hack/cloud-cleanup.sh against a stub `aliyun` on PATH, so no call ever
# reaches Alibaba Cloud. Run it from anywhere: hack/cloud-cleanup-test.sh
set -euo pipefail

SCRIPT="$(cd "$(dirname "$0")" && pwd)/cloud-cleanup.sh"
WORK="$(mktemp -d)"
trap 'rm -rf "${WORK}"' EXIT
mkdir -p "${WORK}/bin"

# The stub logs every call and answers from files in $STUB_DIR:
#   creds        "bad" makes every call fail as with a rotated key
#   delete/<id>  "fail:<code>" fails every delete of <id>; "fail-once:<code>"
#                fails only the first; anything else, or no file, succeeds
#   exists/<id>  how many resources with <id> the describe call finds
#                (default 1), or "error" to make that call fail
#   fail/<action> makes that API action fail with the code in the file
#   response/<action> is printed as the response to that action
cat >"${WORK}/bin/aliyun" <<'STUB'
#!/usr/bin/env bash
set -euo pipefail
echo "$*" >>"${STUB_DIR}/calls"
fail() {
  # Like the real CLI, errors go to stderr.
  echo "{\"schemaVersion\":\"v1\",\"message\":\"stub error\",\"error_code\":\"$1\",\"status_code\":400}" >&2
  exit 2
}
if [ "$(cat "${STUB_DIR}/creds" 2>/dev/null)" = bad ]; then
  fail InvalidAccessKeyId.NotFound
fi
action="$2"
[ -f "${STUB_DIR}/fail/${action}" ] && fail "$(cat "${STUB_DIR}/fail/${action}")"
if [ -f "${STUB_DIR}/response/${action}" ]; then
  cat "${STUB_DIR}/response/${action}"
  exit 0
fi
# The id is the value of the first argument that names one.
id=""
args=("$@")
for ((i = 0; i < ${#args[@]} - 1; i++)); do
  case "${args[i]}" in
    --VpcId | --SecurityGroupId | --LoadBalancerId | --LoadBalancerIds.1 | --ListenerId | --ListenerIds.1 | \
      --InstanceId | --InstanceId.1 | --ClusterId | --RoleName | --ServiceId | --EndpointId | --DBInstanceId | \
      --VSwitchId | --RuleId | --InstanceIds)
      id="${args[i + 1]}"
      break
      ;;
  esac
done
id="${id//[\[\]\"]/}"
case "${action}" in
  Delete* | Release* | UpdateHubClusterFeature | DisableDeletionProtection)
    behavior="$(cat "${STUB_DIR}/delete/${id}" 2>/dev/null || true)"
    case "${behavior}" in
      fail-once:*)
        if [ ! -f "${STUB_DIR}/failed-once-${id}" ]; then
          touch "${STUB_DIR}/failed-once-${id}"
          fail "${behavior#fail-once:}"
        fi
        ;;
      fail:*) [[ "${action}" == Delete* ]] && fail "${behavior#fail:}" ;;
    esac
    echo '{}'
    ;;
  GetCallerIdentity) echo '{"IdentityType":"RAMUser"}' ;;
  *)
    n="$(cat "${STUB_DIR}/exists/${id}" 2>/dev/null || echo 1)"
    [ "${n}" = error ] && fail Throttling
    items="[]"
    [ "${n}" = 0 ] || items="[{\"id\":\"${id}\"}]"
    case "${action}" in
      DescribeVpcs) echo "{\"Vpcs\":{\"Vpc\":${items}}}" ;;
      DescribeSecurityGroups) echo "{\"SecurityGroups\":{\"SecurityGroup\":${items}}}" ;;
      ListLoadBalancers) echo "{\"LoadBalancers\":${items}}" ;;
      ListListeners) echo "{\"Listeners\":${items}}" ;;
      GetInstanceList) echo "{\"InstanceList\":{\"InstanceVO\":${items}}}" ;;
      *) echo '{}' ;;
    esac
    ;;
esac
STUB
chmod +x "${WORK}/bin/aliyun"

failures=0
case_dir=""

# start <name> resets the stub's state for a new case.
start() {
  case_dir="${WORK}/$1"
  mkdir -p "${case_dir}/delete" "${case_dir}/exists" "${case_dir}/fail" "${case_dir}/response"
  : >"${case_dir}/calls"
  echo "--- $1"
}

# run <args>... runs the script against the stub, saving output and status.
run() {
  set +e
  STUB_DIR="${case_dir}" PATH="${WORK}/bin:${PATH}" PASSES="${PASSES:-3}" PASS_WAIT=0 \
    bash "${SCRIPT}" "$@" >"${case_dir}/out" 2>&1
  echo $? >"${case_dir}/status"
  set -e
}

targets() {
  printf '%s\n' "$@" >"${case_dir}/targets"
  run delete "${case_dir}/targets"
}

check() {
  local description="$1"
  shift
  if "$@"; then
    echo "ok: ${description}"
  else
    echo "FAIL: ${description}"
    sed 's/^/    /' "${case_dir}/out"
    failures=$((failures + 1))
  fi
}
status_is() { [ "$(cat "${case_dir}/status")" = "$1" ]; }
output_has() { grep -qF -- "$1" "${case_dir}/out"; }
called() { grep -qF -- "$1" "${case_dir}/calls"; }
not_called() { ! grep -qF -- "$1" "${case_dir}/calls"; }
first_line_of() { grep -nF -- "$1" "${case_dir}/calls" | head -1 | cut -d: -f1; }

start bad-credentials
echo bad >"${case_dir}/creds"
targets "us-west-1 vpc-a"
check "a bad key fails the run" status_is 1
check "a bad key is reported" output_has "cannot authenticate to Alibaba Cloud: InvalidAccessKeyId.NotFound"
check "a bad key deletes nothing" not_called "DeleteVpc"

start bad-credentials-inventory
echo bad >"${case_dir}/creds"
run inventory us-west-1
check "inventory with a bad key fails" status_is 1

start unknown-named-type
targets "us-west-1 kafak:x"
check "a mistyped type is rejected" status_is 1
check "the rejection names the target" output_has 'unknown resource type in "kafak:x"'
check "nothing is called for a rejected target" not_called "Delete"

start unknown-prefix
targets "us-west-1 foo-1"
check "an id without a known prefix is rejected" status_is 1

start prefixes
targets "us-west-1 epsrv-1" "us-west-1 ep-1" "us-west-1 rule-1" "us-west-1 r-1" "us-west-1 rr-1" "us-west-1 rm-1"
check "all targets deleted" status_is 0
check "epsrv- is an endpoint service" called "privatelink DeleteVpcEndpointService --RegionId us-west-1 --ServiceId epsrv-1"
check "ep- is an endpoint" called "privatelink DeleteVpcEndpoint --RegionId us-west-1 --EndpointId ep-1"
check "rule- is an ALB rule, not Tair" called "alb DeleteRule --region us-west-1 --RuleId rule-1"
check "r- is Tair" called "r-kvstore DeleteInstance --region us-west-1 --InstanceId r-1"
check "rr- is an RDS read-only instance" called "rds DeleteDBInstance --region us-west-1 --DBInstanceId rr-1"

start order-and-retry
echo "fail-once:DependencyViolation.SecurityGroup" >"${case_dir}/delete/vpc-a"
targets "us-west-1 vpc-a" "us-west-1 sg-a" "us-west-1 alb-a" "us-west-1 lsn-a"
check "the run succeeds" status_is 0
check "the summary counts every delete" output_has "Deleted 4, not found 0, failed 0."
check "the blocked VPC was retried" output_has "pass 1: \`vpc-a\` (vpc, us-west-1) not deleted yet: \`DependencyViolation.SecurityGroup"
check "listener before load balancer" [ "$(first_line_of DeleteListener)" -lt "$(first_line_of "DeleteLoadBalancer")" ]
check "load balancer before security group" [ "$(first_line_of "DeleteLoadBalancer")" -lt "$(first_line_of DeleteSecurityGroup)" ]
check "security group before VPC" [ "$(first_line_of DeleteSecurityGroup)" -lt "$(first_line_of DeleteVpc)" ]

start already-gone
echo "fail:InvalidVpcId.NotFound" >"${case_dir}/delete/vpc-gone"
echo 0 >"${case_dir}/exists/vpc-gone"
targets "us-west-1 vpc-gone"
check "a missing resource is not a failure" status_is 0
check "a missing resource is reported as such" output_has "\`vpc-gone\` (vpc) does not exist in us-west-1; nothing to delete"
check "a run that deleted nothing warns" output_has "::warning::nothing was deleted"

start check-fails
echo "fail:InternalError" >"${case_dir}/delete/vpc-a"
echo error >"${case_dir}/exists/vpc-a"
targets "us-west-1 vpc-a"
check "a failed delete whose check also fails is a failure" status_is 1
check "it is never reported as gone" output_has "Deleted 0, not found 0, failed 1."

start still-there
echo "fail:IncorrectStatus" >"${case_dir}/delete/vpc-a"
targets "us-west-1 vpc-a"
check "a failed delete of an existing resource is a failure" status_is 1

start missing-key-is-not-gone
echo "fail:InvalidVpcId.NotFound" >"${case_dir}/delete/vpc-a"
echo '{"VPCs":{"Vpc":[]}}' >"${case_dir}/response/DescribeVpcs"
targets "us-west-1 vpc-a"
check "an existence check with a renamed key fails the target" status_is 1

start preparation-failure
echo "Forbidden.RAM" >"${case_dir}/fail/UpdateHubClusterFeature"
targets "cn-beijing ackone:c123"
check "the delete still runs" called "adcp DeleteHubCluster --region cn-beijing --ClusterId c123"
check "the failed preparation is reported" output_has "could not turn off ArgoCD and deletion protection: Forbidden.RAM"

start inventory
cat >"${case_dir}/response/DescribeVpcs" <<'JSON'
{"TotalCount": 5, "Vpcs": {"Vpc": [{"VpcId": "vpc-1", "VpcName": "a|b"}, {"VpcId": "vpc-2", "VpcName": null}]}}
JSON
echo '{"VSwitchSet": {}}' >"${case_dir}/response/DescribeVSwitches"
echo "Forbidden.RAM" >"${case_dir}/fail/DescribeSecurityGroups"
run inventory us-west-1 "" cn-beijing
check "inventory succeeds" status_is 0
check "truncation is shown" output_has "Showing 2 of 5."
check "a pipe in a cell is escaped" output_has '| vpc-1 | a\|b |'
check "a null cell is empty" output_has '| vpc-2 |  |'
check "a renamed key is reported" output_has "Unexpected response: no \`.VSwitches.VSwitch\` in it. Top-level keys: \`VSwitchSet\`."
check "a failed list is reported" output_has "Could not list: \`Forbidden.RAM: stub error\`"
check "an empty region is skipped" [ "$(grep -c '^### ' "${case_dir}/out")" -eq 3 ]

echo
if [ "${failures}" -gt 0 ]; then
  echo "${failures} check(s) failed"
  exit 1
fi
echo "all checks passed"
