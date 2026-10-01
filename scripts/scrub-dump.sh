#!/usr/bin/env bash
# Redact cloud credentials from a controlplane dump before it is uploaded.
#
# The Alibaba Cloud SDKs sign requests with query parameters, so a failed call
# puts the access key id and the request signature into the error string. That
# error reaches the managed resource's status conditions and the provider log,
# both of which the dump collects -- and workflow artifacts on a public
# repository are downloadable by anyone.
#
# Usage: scrub-dump.sh <dump-directory>
set -euo pipefail

DUMP_DIR="${1:-}"
if [[ -z "${DUMP_DIR}" || ! -d "${DUMP_DIR}" ]]; then
  echo "usage: $(basename "$0") <dump-directory>" >&2
  exit 1
fi

# Query parameters whose values identify or authenticate the caller. The secret
# key is never sent, so it cannot appear here, but the key id names the
# credential in use and the signature is derived from the secret.
PARAMS=(AccessKeyId Signature SignatureNonce SecurityToken)

# Access key ids are stable and greppable, so redact them wherever they appear,
# not only inside a query string.
KEY_ID_PATTERN='LTAI[0-9A-Za-z]{12,24}'

# grep exits 1 when it finds nothing, which is not an error here.
count_matches() {
  local pattern="$1"
  grep -rEoI "${pattern}" "${DUMP_DIR}" 2>/dev/null | grep -cve '^$' || true
}

before=0
for param in "${PARAMS[@]}"; do
  before=$((before + $(count_matches "${param}=[^&\"'[:space:]]+")))
done
before=$((before + $(count_matches "${KEY_ID_PATTERN}")))

while IFS= read -r -d '' file; do
  for param in "${PARAMS[@]}"; do
    # Stop at the value's delimiter so the rest of the URL stays readable.
    perl -pi -e "s/${param}=[^&\"'\\s]+/${param}=REDACTED/g" "${file}"
  done
  perl -pi -e "s/${KEY_ID_PATTERN}/REDACTED/g" "${file}"
done < <(find "${DUMP_DIR}" -type f -print0)

# Anything still carrying a value other than REDACTED means a pattern missed.
leaked=0
for param in "${PARAMS[@]}"; do
  remaining=$(grep -rEoI "${param}=[^&\"'[:space:]]+" "${DUMP_DIR}" 2>/dev/null | grep -cv "=REDACTED$" || true)
  leaked=$((leaked + remaining))
done
leaked=$((leaked + $(count_matches "${KEY_ID_PATTERN}")))

echo "scrub-dump: redacted ${before} credential reference(s) in ${DUMP_DIR}"
if [[ "${leaked}" -ne 0 ]]; then
  echo "scrub-dump: ${leaked} reference(s) survived redaction" >&2
  exit 1
fi
