package gator

import (
	"io"

	constraintclient "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers/rego"
)

type Opt func() ([]constraintclient.Opt, []rego.Arg, error)

func NewOPAClient(includeTrace bool, opts ...Opt) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func WithK8sCEL() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithPrintHook(w io.Writer) Opt { _ = "STUB: not implemented"; return *new(Opt) }
