package output

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

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
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > widths[i] {
				widths[i] = len(cell)
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
		parts[i] = fmt.Sprintf("%-*s", widths[i], cell)
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
		val := fmt.Sprintf("%v", v.Field(i).Interface())
		fmt.Fprintf(w, "%-10s %s\n", tag+":", val)
	}
	return nil
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
