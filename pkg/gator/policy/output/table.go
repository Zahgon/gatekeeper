package output

import (
	"io"
)

// TablePrinter outputs results in human-readable table format.
type TablePrinter struct{}

// PrintPolicies outputs a table of installed policies.
func (p *TablePrinter) PrintPolicies(w io.Writer, policies []PolicyInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Header

// Rows

// Just the date part

// PrintSearchResults outputs a table of search results.
func (p *TablePrinter) PrintSearchResults(w io.Writer, results []SearchResult) error {
	_ = "STUB: not implemented"
	return nil
}

// Header

// Rows

// PrintUpdateResult outputs catalog update results as a human-readable table.
func (p *TablePrinter) PrintUpdateResult(w io.Writer, result *UpdateResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintMessage outputs a simple message.
func (p *TablePrinter) PrintMessage(w io.Writer, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintInstallResult outputs install results as a human-readable table.
func (p *TablePrinter) PrintInstallResult(w io.Writer, result *InstallResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintUninstallResult outputs uninstall results as a human-readable table.
func (p *TablePrinter) PrintUninstallResult(w io.Writer, result *UninstallResult) error {
	_ = "STUB: not implemented"
	return nil
}

// PrintUpgradeResult outputs upgrade results as a human-readable table.
func (p *TablePrinter) PrintUpgradeResult(w io.Writer, result *UpgradeResult) error {
	_ = "STUB: not implemented"
	return nil
}
