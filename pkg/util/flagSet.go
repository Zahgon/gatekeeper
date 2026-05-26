package util

import (
	"flag"
)

type FlagSet map[string]bool

var _ flag.Value = FlagSet{}

func NewFlagSet() FlagSet { _ = "STUB: not implemented"; return *new(FlagSet) }

func (l FlagSet) ToSlice() []string { _ = "STUB: not implemented"; return nil }

func (l FlagSet) String() string { _ = "STUB: not implemented"; return "" }

func (l FlagSet) Set(s string) error { _ = "STUB: not implemented"; return nil }
