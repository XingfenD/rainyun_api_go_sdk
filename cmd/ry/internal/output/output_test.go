package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

type testItem struct {
	ID   string `json:"id"   table:"ID"`
	Name string `json:"name" table:"NAME"`
	Age  int    `json:"age"  table:"AGE"`
}

func TestTableFormat(t *testing.T) {
	items := []testItem{
		{ID: "1", Name: "Alice", Age: 30},
		{ID: "2", Name: "Bob", Age: 25},
	}
	var buf bytes.Buffer
	printer := New("table", &buf)
	if err := printer.Print(items); err != nil {
		t.Fatalf("Print table error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "ID") || !strings.Contains(out, "Alice") {
		t.Errorf("table output missing headers or data:\n%s", out)
	}
}

func TestJSONFormat(t *testing.T) {
	items := []testItem{{ID: "1", Name: "Alice", Age: 30}}
	var buf bytes.Buffer
	printer := New("json", &buf)
	if err := printer.Print(items); err != nil {
		t.Fatalf("Print json error: %v", err)
	}
	var parsed []testItem
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if len(parsed) != 1 || parsed[0].ID != "1" {
		t.Errorf("JSON content mismatch: %+v", parsed)
	}
}

func TestYAMLFormat(t *testing.T) {
	items := []testItem{{ID: "1", Name: "Alice", Age: 30}}
	var buf bytes.Buffer
	printer := New("yaml", &buf)
	if err := printer.Print(items); err != nil {
		t.Fatalf("Print yaml error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "id:") || !strings.Contains(out, "Alice") {
		t.Errorf("yaml output missing data:\n%s", out)
	}
}

func TestRawFormat(t *testing.T) {
	var buf bytes.Buffer
	printer := New("raw", &buf)
	printer.SetRawBody(func() []byte { return []byte(`{"raw":"data"}`) })
	if err := printer.Print("ignored"); err != nil {
		t.Fatalf("Print raw error: %v", err)
	}
	if buf.String() != `{"raw":"data"}`+"\n" {
		t.Errorf("raw output = %q", buf.String())
	}
}

func TestRawFormatWithoutProvider(t *testing.T) {
	var buf bytes.Buffer
	printer := New("raw", &buf)
	if err := printer.Print("ignored"); err == nil {
		t.Error("expected error when raw provider is unset")
	}
}

func TestSingleItem(t *testing.T) {
	item := testItem{ID: "1", Name: "Solo", Age: 99}
	var buf bytes.Buffer
	printer := New("json", &buf)
	if err := printer.Print(item); err != nil {
		t.Fatalf("Print single item error: %v", err)
	}
	if !strings.Contains(buf.String(), "Solo") {
		t.Errorf("single item output missing data:\n%s", buf.String())
	}
}

func TestUnknownFormat(t *testing.T) {
	var buf bytes.Buffer
	printer := New("unknown", &buf)
	err := printer.Print([]testItem{})
	if err == nil {
		t.Error("expected error for unknown format")
	}
}

func TestTableSingleNilPointer(t *testing.T) {
	type withPtr struct {
		Name string     `table:"NAME"`
		When *time.Time `table:"WHEN"`
	}
	var buf bytes.Buffer
	printer := New("table", &buf)
	if err := printer.Print(withPtr{Name: "x"}); err != nil {
		t.Fatalf("Print error: %v", err)
	}
	if !strings.Contains(buf.String(), "WHEN:      -") {
		t.Errorf("nil pointer not rendered as '-':\n%s", buf.String())
	}
}

func TestTableCJKAlignment(t *testing.T) {
	type product struct {
		ID   string `table:"ID"`
		Name string `table:"NAME"`
		Type string `table:"TYPE"`
	}
	items := []product{
		{ID: "1", Name: "TrustAsia域名型SSL证书", Type: "dv"},
		{ID: "2", Name: "GeoSSL域名型通配符SSL证书", Type: "dv"},
		{ID: "3", Name: "PlainASCII", Type: "ov"},
	}
	var buf bytes.Buffer
	printer := New("table", &buf)
	if err := printer.Print(items); err != nil {
		t.Fatalf("Print error: %v", err)
	}
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	if len(lines) < 5 {
		t.Fatalf("expected 5 lines (header+sep+3rows), got %d", len(lines))
	}
	idxType := strings.Index(lines[0], "TYPE")
	for i, line := range lines {
		if i == 1 {
			continue
		}
		pos := strings.Index(line, lines[i][idxType:idxType+4])
		if pos != idxType {
			t.Errorf("line %d: TYPE column at pos %d, want %d\n%s", i, pos, idxType, buf.String())
		}
	}
}
