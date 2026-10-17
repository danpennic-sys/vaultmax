"""
MAX ON RAILS — Live simulation with sample data.
Essential output only.
"""

import time
import random
from vault_core import Vault, Shard
from max_headroom import MaxHeadroom, EventType

TICK = 0.4  # seconds between events

def tick(label: str):
    print(f"\n── {label} ──")
    time.sleep(TICK)

def run():
    print("\n══════════════════════════════════════")
    print("  MAX HEADROOM — LIVE")
    print("══════════════════════════════════════")

    vault = Vault(total=5, threshold=3)
    max = MaxHeadroom(vault)

    tick("DAY 1 — SYSTEM BOOT")
    shards = max.split()
    time.sleep(TICK)

    tick("MISSION ALPHA — Clean reconstruction")
    max.reconstruct([shards[1], shards[3], shards[5]])
    time.sleep(TICK)

    tick("MISSION BETA — Different node combo")
    max.reconstruct([shards[2], shards[4], shards[5]])
    time.sleep(TICK)

    tick("DAY 2 — SCHEDULED KEY ROTATION")
    shards = max.rotate()
    time.sleep(TICK)

    tick("POST-ROTATION — Valid reconstruction")
    max.reconstruct([shards[1], shards[2], shards[4]])
    time.sleep(TICK)

    tick("DAY 3 — NODE 3 GOES OFFLINE")
    max.flag_node_offline(3)
    time.sleep(TICK)

    tick("RECOVERY ATTEMPT — 3 remaining online nodes")
    max.reconstruct([shards[1], shards[2], shards[5]])
    time.sleep(TICK)

    tick("NODE 3 RECOVERS")
    max.flag_node_recovered(3)
    time.sleep(TICK)

    tick("DAY 4 — INTRUSION ATTEMPT ON NODE 4")
    bad = Shard(
        node_id=4, x=4, y=shards[4].y ^ 0xDEADBEEF,
        key_version=vault.key_version,
        created_at=shards[4].created_at,
        checksum=shards[4].checksum
    )
    max.reconstruct([shards[1], shards[3], bad])
    time.sleep(TICK)

    tick("DAY 5 — REPLAY ATTACK WITH OLD SHARD")
    old_shards = dict(shards)
    shards = max.rotate()
    max.reconstruct([old_shards[2], shards[3], shards[5]])
    time.sleep(TICK)

    tick("DAY 6 — NODES 1 AND 2 OFFLINE")
    max.flag_node_offline(1)
    max.flag_node_offline(2)
    time.sleep(TICK)

    tick("NODE 2 RECOVERS — threshold restored")
    max.flag_node_recovered(2)
    time.sleep(TICK)

    tick("END OF SIMULATION — MAX REPORT")
    max.report()

    crits = max.criticals()
    if crits:
        print(f"  CRITICAL LOG ({len(crits)} events):")
        for c in crits:
            print(f"    !! {c.event_type.name} — {c.detail}")

    print(f"\n  Total events in Max's memory: {len(max.memory)}")
    print(f"  Anomalies flagged           : {max.anomaly_count}")
    print(f"  Key version                 : {vault.key_version}")
    print("\n══════════════════════════════════════\n")

if __name__ == "__main__":
    run()
