package schema

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
)

type IDSet map[types.ID]bool

func (c IDSet) String() string { _ = "STUB: not implemented"; return "" }

func (c IDSet) ToList() []types.ID { _ = "STUB: not implemented"; return nil }
