package output

import (
	"fmt"
	"io"
)

type Printer struct {
	format  string
	writer  io.Writer
	rawBody func() []byte
}

func New(format string, w io.Writer) *Printer {
	return &Printer{format: format, writer: w}
}

// SetRawBody registers a provider for the raw API response body. When the
// format is "raw", Print writes the value returned by this provider instead of
// encoding the passed data.
func (p *Printer) SetRawBody(fn func() []byte) {
	p.rawBody = fn
}

func (p *Printer) Print(data any) error {
	switch p.format {
	case "table":
		return printTable(p.writer, data)
	case "json":
		return printJSON(p.writer, data)
	case "yaml":
		return printYAML(p.writer, data)
	case "raw":
		return p.printRawBody()
	default:
		return fmt.Errorf("unknown output format: %s", p.format)
	}
}

func (p *Printer) printRawBody() error {
	if p.rawBody == nil {
		return fmt.Errorf("raw response body unavailable")
	}
	_, err := fmt.Fprintln(p.writer, string(p.rawBody()))
	return err
}
