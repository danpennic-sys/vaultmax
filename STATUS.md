# Status — 2026-10-17

- Private repo created: https://github.com/danpennic-sys/vaultmax
- Core + hardening + network + ops modules staged
- Dual license files present (AGPL core vs commercial stack)
- **Export:** do not flip public until BIS notification complete

Next:
1. Package as installable module (`pyproject.toml`)
2. Wire HardenedVault + CredentialVault into boot_full_stack
3. Prefer secure_wire over plain TCP in production path
