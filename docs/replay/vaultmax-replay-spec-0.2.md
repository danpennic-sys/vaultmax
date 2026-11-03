# VaultMax Replay Acknowledgment — Minimal Checkable Spec
**Version:** 0.2 (expanded contracts)  
**Status:** Proposed  
**Packet:** CP‑20260921‑1920‑002  
**Previous:** 0.1  

## Purpose
Operator-initiated acknowledgment of a replay event that produces a signed, hash-chained receipt.  
Nothing is claimed as implemented.

## Required Artifacts (still missing)
- Operator entry point (`vaultmax_replay_b1.sh` or equivalent)
- Receipt schema implementation
- Signing key source and algorithm
- State store
- Policy decision function

## State Machine

| State        | Entry                          | Exit                           | Side Effect                  |
|--------------|--------------------------------|--------------------------------|------------------------------|
| Detected     | Instrument + gate resolved     | Policy + timeout timer start   | Create Pending record        |
| Pending      | Record created                 | Ack received or timeout        | None                         |
| Acknowledged | Valid signature + policy allow | Terminal                       | Append signed receipt        |
| Failed       | Any rejection or timeout       | Terminal                       | Append failure receipt       |

### Transition Rules
- Only Pending may move to Acknowledged or Failed.
- Acknowledged and Failed are terminal.
- Every transition produces exactly one receipt.

## Receipt Schema (canonical JSON)

```json
{
  "id": "string (uuid)",
  "prev_hash": "hex (sha256 of previous receipt or genesis)",
  "operator_id": "string",
  "gate_id": "string",
  "instrument_hash": "hex",
  "action": "ack | fail",
  "reason": "string | null",
  "timestamp": "ISO-8601",
  "signature": "hex"
}
```

Canonicalization: UTF-8 JSON with sorted keys, no whitespace.  
Signature is over the canonical bytes of the object excluding the `signature` field.

## Failure Codes

| Code          | Condition                              |
|---------------|----------------------------------------|
| E_CHECKSUM    | instrument_hash mismatch               |
| E_SIGNATURE   | signature verification failure         |
| E_POLICY      | policy function returned deny          |
| E_TIMEOUT     | no valid ack within window             |
| E_MALFORMED   | receipt or instrument fails schema     |

## Timeout
- Default window: 300 seconds from Pending entry.
- Clock source: monotonic system clock at operator host (must be recorded in receipt notes if used).
- On expiry the system writes a Failed receipt with `E_TIMEOUT`.

## Operator Interface Contract

The operator entry point must accept:

```
vaultmax_replay_b1.sh \
  --instrument <path-or-hash> \
  --gate <gate_id> \
  --action ack|fail \
  [--reason <code>]
```

Exit codes:
- 0 = receipt written
- 1 = policy or signature rejection
- 2 = timeout or store error
- 3 = malformed input

## Policy Interface (minimal)

```
decide(gate_id, instrument_hash, operator_id) → allow | deny
```

Implementation is external. The script must call it before signing an ack.

## Loader Metadata (per segment)

| Field         | Type   | Required |
|---------------|--------|----------|
| segment_id    | string | yes      |
| status        | staged \| validated | yes |
| receipt_id    | string | on validated |
| operator_id   | string | yes      |
| created_at    | string | yes      |
| validated_at  | string | on validated |

A segment may move from staged to validated only by presence of an Acknowledged receipt.

## Open Items Remaining
1. Concrete signing algorithm and key management
2. Actual state store (file, sqlite, or chain)
3. Policy function implementation
4. Genesis hash definition
5. Test vectors

## Disposition
This document is a contract template.  
It becomes binding only when the listed artifacts exist and produce the receipts defined here under controlled execution.

**End of version 0.2.**
