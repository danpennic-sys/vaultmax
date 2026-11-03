// vaultmax-receipt.go
// Minimal receipt writer skeleton
// Spec: vaultmax-replay-spec-0.2.md
// Build: go build -o vaultmax-receipt vaultmax-receipt.go

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Receipt struct {
	ID             string  `json:"id"`
	PrevHash       string  `json:"prev_hash"`
	OperatorID     string  `json:"operator_id"`
	GateID         string  `json:"gate_id"`
	InstrumentHash string  `json:"instrument_hash"`
	Action         string  `json:"action"`
	Reason         *string `json:"reason"`
	Timestamp      string  `json:"timestamp"`
	Signature      string  `json:"signature"`
}

func main() {
	action := flag.String("action", "", "ack or fail")
	gate := flag.String("gate", "", "gate id")
	instrumentHash := flag.String("instrument-hash", "", "hex hash")
	reason := flag.String("reason", "", "failure code")
	stateDir := flag.String("state-dir", "/var/lib/vaultmax/replay", "state directory")
	operatorID := flag.String("operator-id", "operator-stub", "operator identifier")
	flag.Parse()

	if *action == "" || *gate == "" || *instrumentHash == "" {
		fmt.Fprintln(os.Stderr, "missing required flags")
		os.Exit(3)
	}

	// Load previous hash (stub: genesis if none)
	prevHash := loadPrevHash(*stateDir)

	var reasonPtr *string
	if *reason != "" {
		reasonPtr = reason
	}

	r := Receipt{
		ID:             fmt.Sprintf("rcp-%d", time.Now().UnixNano()),
		PrevHash:       prevHash,
		OperatorID:     *operatorID,
		GateID:         *gate,
		InstrumentHash: *instrumentHash,
		Action:         *action,
		Reason:         reasonPtr,
		Timestamp:      time.Now().UTC().Format(time.RFC3339),
		Signature:      "unsigned-stub",
	}

	// Canonicalize and hash (signature stub)
	canonical, _ := json.Marshal(r)
	h := sha256.Sum256(canonical)
	r.Signature = hex.EncodeToString(h[:]) // replace with real signing

	// Persist
	if err := persist(*stateDir, r); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Println(r.ID)
}

func loadPrevHash(dir string) string {
	path := filepath.Join(dir, "latest.hash")
	data, err := os.ReadFile(path)
	if err != nil {
		return "genesis"
	}
	return string(data)
}

func persist(dir string, r Receipt) error {
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(r, "", "  ")
	path := filepath.Join(dir, r.ID+".json")
	if err := os.WriteFile(path, data, 0640); err != nil {
		return err
	}
	// Update latest hash
	h := sha256.Sum256(data)
	return os.WriteFile(filepath.Join(dir, "latest.hash"), []byte(hex.EncodeToString(h[:])), 0640)
}
