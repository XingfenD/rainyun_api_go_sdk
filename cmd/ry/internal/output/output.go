package output

import (
	"encoding/json"
	"fmt"
	"io"
)

type Printer struct {
	format string
	writer io.Writer
}

func New(format string, w io.Writer) *Printer {
	return &Printer{format: format, writer: w}
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
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(p.writer, string(b))
		return err
	default:
		return fmt.Errorf("unknown output format: %s", p.format)
	}
}

func (p *Printer) PrintRaw(raw []byte) error {
	_, err := fmt.Fprintln(p.writer, string(raw))
	return err
}
