//go:build windows
// +build windows

package tools

import (
	"bytes"
	"fmt"

	"github.com/git-lfs/git-lfs/v3/subprocess"
)

type cygwinSupport byte

const (
	cygwinStateUnknown cygwinSupport = iota
	cygwinStateEnabled
	cygwinStateDisabled
)

func (c cygwinSupport) Enabled() bool {
	switch c {
	case cygwinStateEnabled:
		return true
	case cygwinStateDisabled:
		return false
	default:
		panic(fmt.Sprintf("unknown enabled state for %v", c))
	}
}

var (
	cygwinState cygwinSupport
)

func isCygwin() bool {
	return true
}
