package output

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"unicode"
)

func displayWidth(s string) int {
	w := 0
	for _, r := range s {
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) || unicode.Is(unicode.Katakana, r) || unicode.Is(unicode.Hiragana, r) {
			w += 2
		} else {
			w++
		}
	}
	return w
}

func padRight(s string, width int) string {
	dw := displayWidth(s)
	if dw >= width {
		return s
	}
	return s + strings.Repeat(" ", width-dw)
}

func printTable(w io.Writer, data any) error {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		return printTableSlice(w, v)
	}
	return printTableSingle(w, v)
}

func printTableSlice(w io.Writer, v reflect.Value) error {
	if v.Len() == 0 {
		fmt.Fprintln(w, "(empty)")
		return nil
	}
	elemType := v.Index(0).Type()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	headers, cols := extractTableTags(elemType)
	if len(headers) == 0 {
		return fmt.Errorf("no table tags found on type %s", elemType.Name())
	}

	rows := make([][]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		rows[i] = extractRow(v.Index(i), cols)
	}

	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = displayWidth(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if displayWidth(cell) > widths[i] {
				widths[i] = displayWidth(cell)
			}
		}
	}

	sep := "  "
	renderRow(w, headers, widths, sep)

	separator := make([]string, len(headers))
	for i, ww := range widths {
		separator[i] = strings.Repeat("-", ww)
	}
	renderRow(w, separator, widths, sep)

	for _, row := range rows {
		renderRow(w, row, widths, sep)
	}

	return nil
}

func renderRow(w io.Writer, cells []string, widths []int, sep string) {
	parts := make([]string, len(cells))
	for i, cell := range cells {
		parts[i] = padRight(cell, widths[i])
	}
	fmt.Fprintln(w, strings.Join(parts, sep))
}

func printTableSingle(w io.Writer, v reflect.Value) error {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("table")
		if tag == "" {
			continue
		}
		fmt.Fprintf(w, "%-10s %s\n", tag+":", formatCell(v.Field(i)))
	}
	return nil
}

func formatCell(fv reflect.Value) string {
	if fv.Kind() == reflect.Ptr {
		if fv.IsNil() {
			return "-"
		}
		return fmt.Sprintf("%v", fv.Interface())
	}
	return fmt.Sprintf("%v", fv.Interface())
}

func extractTableTags(t reflect.Type) ([]string, []int) {
	var headers []string
	var cols []int
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("table")
		if tag == "" {
			continue
		}
		headers = append(headers, strings.ToUpper(tag))
		cols = append(cols, i)
	}
	return headers, cols
}

func extractRow(v reflect.Value, cols []int) []string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	row := make([]string, len(cols))
	for i, idx := range cols {
		row[i] = fmt.Sprintf("%v", v.Field(idx).Interface())
	}
	return row
}
