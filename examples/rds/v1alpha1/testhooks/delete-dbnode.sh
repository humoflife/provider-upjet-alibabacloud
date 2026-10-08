#!/usr/bin/env bash
# Deletes the DBNode before the Instance it belongs to.
#
# uptest issues every delete with --wait=false and only waits for completion in
# a later step, so listing DBNode ahead of Instance in the manifest does not
# make it finish first -- the two race, and the Instance usually wins. Once the
# Instance is gone, DescribeDBInstanceAttribute can no longer resolve the node,
# upstream Read returns an error rather than clearing the id, and upjet has no
# hook to treat a provider error as "resource gone". The DBNode then keeps its
# finalizer and teardown hangs until someone removes it by hand.
#
# Running as a pre-delete hook on the Instance keeps the node on the path that
# does work: deleted while its parent is still around.
set -euo pipefail

KUBECTL="${KUBECTL:-kubectl}"
DBNODE="dbnode.rds.alibabacloud.crossplane.io/default"
# Well inside the 40m chainsaw exec budget for the whole delete step, which
# this hook shares with every other resource in the example.
TIMEOUT="${UPTEST_DBNODE_DELETE_TIMEOUT:-20m}"

if ! "${KUBECTL}" get "${DBNODE}" >/dev/null 2>&1; then
  echo "pre-delete-hook: ${DBNODE} is not present, nothing to do"
  exit 0
fi

echo "pre-delete-hook: deleting ${DBNODE} while its parent Instance still exists"
"${KUBECTL}" delete "${DBNODE}" --wait=false --ignore-not-found
"${KUBECTL}" wait "${DBNODE}" --for=delete --timeout="${TIMEOUT}"
echo "pre-delete-hook: ${DBNODE} is gone, the Instance can now be deleted"
