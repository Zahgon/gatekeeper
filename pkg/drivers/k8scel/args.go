package k8scel

type Arg func(*Driver) error

// GatherStats starts collecting various stats around the
// underlying engine's calls.
func GatherStats() Arg { _ = "STUB: not implemented"; return *new(Arg) }
