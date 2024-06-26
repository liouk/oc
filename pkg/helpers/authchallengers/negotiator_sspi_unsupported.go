//go:build !windows
// +build !windows

package authchallengers

import (
	"io"
)

func SSPIEnabled() bool {
	return false
}

func NewSSPINegotiator(string, string, string, string, io.Reader) Negotiator {
	return newUnsupportedNegotiator("SSPI")
}
