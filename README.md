# VaultMax

**Self-Designing · Self-Healing · Mil-Spec Rotating Split-Key Node Vault**

> *"I know everything that happens in this vault. I forget nothing."*  
> — Max Headroom, Vault Keeper

**This repository is private.** Encryption software is EAR-controlled (ECCN **5E002**). Do not make public or export until BIS notification / License Exception TSU is filed. See export checklist.

---

## Layout

```
core/           AGPL core — vault_core, max_headroom, max_persistent, max_on_rails
hardening/      Commercial — node_hardening, hmac_hardening, credential_isolation
network/        Commercial — network_healing (TCP), secure_wire (mTLS + AES-GCM)
ops/            Commercial — rotation_scheduler, self_designing, reset_15
docs/           RFC 8785/9557 suite, validation manifest
web/            Marketing landing (Carrd-style)
LICENSE_AGPL.txt
LICENSE_COMMERCIAL.txt
```

## Dual license

| Scope | License |
|-------|---------|
| `core/vault_core.py`, `core/max_headroom.py` | **AGPL v3** |
| Hardening, network, ops, secure wire | **Commercial** — see `LICENSE_COMMERCIAL.txt` |

Contact: legal@vaultmax.io · compliance@vaultmax.io

## Quick start (dev)

```bash
pip install cryptography
cd core && python vault_core.py          # Phase 1 smoke
python max_headroom.py                  # Keeper smoke
```

## Security model (summary)

- Shamir 3-of-5 over 2^127-1
- AES-256-GCM at rest per node (PBKDF2 600k)
- Optional mTLS + application AES-GCM + replay protection (`secure_wire`)
- HMAC key isolation (`hmac_hardening`) and per-node credential isolation
- Append-only signed event log (`max_persistent`)
- Adaptive rotation / self-heal / self-design

## Export control

ECCN 5E002. Not for embargoed destinations. File BIS encryption notification before any public source release.

---

*Built with Max Headroom watching.*
