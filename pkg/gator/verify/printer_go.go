package verify

import (
	"errors"
)

type PrinterGo struct{}

var _ Printer = PrinterGo{}

// ErrWritingString means there was a problem writing output to the writer
// passed to Print.
var ErrWritingString = errors.New("writing output")

func (p PrinterGo) Print(w StringWriter, r []SuiteResult, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PrinterGo) PrintSuite(w StringWriter, r *SuiteResult, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PrinterGo) PrintTest(w StringWriter, r *TestResult, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PrinterGo) PrintCase(w StringWriter, r *CaseResult, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

// if using verbose to print, let's keep the trace at the same level
