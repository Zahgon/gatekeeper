package gator

func ParseYaml(yamlBytes []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

// Pass through JSON since k8s parsing logic doesn't fully handle objects
// parsed directly from YAML. Without passing through JSON, the OPA client
// panics when handed scalar types it doesn't recognize.
func FixYAML(obj map[string]interface{}, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func parseJSON(jsonBytes []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }
