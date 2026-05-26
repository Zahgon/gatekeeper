package v1beta1

// dashExtractor unpacks the status resource name, unescaping `-`.
func dashExtractor(val string) []string { _ = "STUB: not implemented"; return nil }

// DashPacker puts a list of strings into a dash-separated format. Note that
// it cannot handle empty strings, as that makes the dash separator for the empty
// string reduce to an escaped dash. This is fine because none of the packed strings
// are allowed to be empty. If this changes in the future, we could create a placeholder
// for the empty string, say `b`, and replace all instances of `b` in the input
// stream with `bb`, which could then be unfolded. If we need that, we are already
// changing the schema of the status resource, and therefore don't need to deal with
// it now. It also doesn't handle the case where a value begins or ends with a dash,
// which is also disallowed by the schema (and would require an additional placeholder
// character to fix). Finally, note that it is impossible to distinguish between
// a nil list of strings and a list of one empty string.
func DashPacker(vals ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }
