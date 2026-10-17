# VaultMax Status — 2026-10-17

**Repo (private):** https://github.com/danpennic-sys/vaultmax  
**Export:** ECCN 5E002 — do not public-release crypto source until BIS notification.

## Tree

```
core/        vault_core, max_headroom, max_persistent, max_on_rails   (AGPL core)
hardening/   node_hardening, hmac_hardening, credential_isolation
network/     network_healing, secure_wire
ops/         rotation_scheduler, self_designing, reset_15
guards/      Domains 1–5 control-plane enforcement
docs/        RFC 8785/9557 suite, validation_manifest
web/         vaultmax-carrd.html
```

## Guards (issues register → code)

| Domain | Module | Role |
|--------|--------|------|
| 1 Legacy vulns | `runtime_integrity.py` | Fail-closed admission: hostile vectors, HMAC, policy digest, no nondet fields |
| 2 Engine drift | `sovereign_drift.py` | Ritual/feed/pulse/latency thresholds; succession halt |
| 3 Sentinel Ops | `sentinel_ops.py` | State-engine only; worker≠API; AI router+log; AI advisory-only; anomaly→opsAdvisor |
| 4 QuantumGuard | `quantum_guard.py` | Constitutional / HARD / RATE / HIGH_IMPACT; RF advisory queue |
| 5 Audit chain | `audit_chain.py` | FSM transitions; hash chain; seal↔audit_event; recorded revocation |

All five: `python3 guards/<module>.py` smoke green.

## Dual license

- AGPL: `core/vault_core.py`, `core/max_headroom.py`
- Commercial: hardening, network, ops, guards (stack)

## Next

1. Full `git push` from workstation (API tree still partial)
2. `pyproject.toml` package layout
3. Wire guards into `boot_full_stack` / secure_wire path
4. BIS notification before any public crypto release
