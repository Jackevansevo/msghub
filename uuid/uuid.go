package uuid

import (
	"bytes"
	"os/exec"
)

// UUID (Universally unique identifier)
type UUID string

// New returns a new UUID
func New() (UUID, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", nil
	}
	return UUID(string(bytes.TrimSuffix(out, []byte("\n")))), nil
}
