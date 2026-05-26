package output

import (
	"io"
)

// JSONPrinter outputs results in JSON format.
type JSONPrinter struct{}

// JSONOutputVersion is the apiVersion for JSON output schema.
const JSONOutputVersion = "gator.gatekeeper.sh/v1alpha1"

// PrintPolicies outputs installed policies as JSON.
func (p *JSONPrinter) PrintPolicies(w io.Writer, policies []PolicyInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintSearchResults outputs search results as JSON.
func (p *JSONPrinter) PrintSearchResults(w io.Writer, results []SearchResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintUpdateResult outputs catalog update results as JSON.
func (p *JSONPrinter) PrintUpdateResult(w io.Writer, result *UpdateResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintMessage outputs a message as JSON.
func (p *JSONPrinter) PrintMessage(w io.Writer, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintInstallResult outputs install results as JSON.
func (p *JSONPrinter) PrintInstallResult(w io.Writer, result *InstallResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintUninstallResult outputs uninstall results as JSON.
func (p *JSONPrinter) PrintUninstallResult(w io.Writer, result *UninstallResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintUpgradeResult outputs upgrade results as JSON.
func (p *JSONPrinter) PrintUpgradeResult(w io.Writer, result *UpgradeResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *JSONPrinter) writeJSON(w io.Writer, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
