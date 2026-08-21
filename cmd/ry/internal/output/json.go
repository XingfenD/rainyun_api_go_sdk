package output

import (
	"io"

	"github.com/bytedance/sonic"
)

func printJSON(w io.Writer, data any) error {
	b, err := sonic.ConfigStd.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))
	return err
}
