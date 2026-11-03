#!/usr/bin/env bash
# vaultmax_replay_b1.sh
# Operator entry point skeleton for VaultMax replay acknowledgment
# Spec: vaultmax-replay-spec-0.2.md
# Status: skeleton only

set -euo pipefail

# Defaults
TIMEOUT_SEC=300
STATE_DIR="${STATE_DIR:-/var/lib/vaultmax/replay}"
RECEIPT_BIN="${RECEIPT_BIN:-./vaultmax-receipt}"

usage() {
  echo "Usage: $0 --instrument <path-or-hash> --gate <gate_id> --action ack|fail [--reason <code>]"
  exit 3
}

# Parse args
INSTRUMENT=""
GATE=""
ACTION=""
REASON=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --instrument) INSTRUMENT="$2"; shift 2 ;;
    --gate) GATE="$2"; shift 2 ;;
    --action) ACTION="$2"; shift 2 ;;
    --reason) REASON="$2"; shift 2 ;;
    *) usage ;;
  esac
done

[[ -z "$INSTRUMENT" || -z "$GATE" || -z "$ACTION" ]] && usage
[[ "$ACTION" != "ack" && "$ACTION" != "fail" ]] && usage

# Compute instrument hash if path given
if [[ -f "$INSTRUMENT" ]]; then
  INSTRUMENT_HASH=$(sha256sum "$INSTRUMENT" | awk '{print $1}')
else
  INSTRUMENT_HASH="$INSTRUMENT"
fi

# Call policy (stub)
if [[ "$ACTION" == "ack" ]]; then
  # TODO: replace with real policy call
  # decide "$GATE" "$INSTRUMENT_HASH" "$OPERATOR_ID" || exit 1
  echo "policy: stub allow" >&2
fi

# Write receipt via helper
"$RECEIPT_BIN" \
  --action "$ACTION" \
  --gate "$GATE" \
  --instrument-hash "$INSTRUMENT_HASH" \
  --reason "$REASON" \
  --state-dir "$STATE_DIR"

exit $?
