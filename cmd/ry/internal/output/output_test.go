package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
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
	raw := []byte(`{"raw":"data"}`)
	var buf bytes.Buffer
	printer := New("raw", &buf)
	if err := printer.PrintRaw(raw); err != nil {
		t.Fatalf("PrintRaw error: %v", err)
	}
	if buf.String() != `{"raw":"data"}`+"\n" {
		t.Errorf("raw output = %q", buf.String())
	}
}

// 新增:Print 在 raw 格式下应输出 JSON
func TestRawPrintFormat(t *testing.T) {
	items := []testItem{{ID: "1", Name: "Alice", Age: 30}}
	var buf bytes.Buffer
	printer := New("raw", &buf)
	if err := printer.Print(items); err != nil {
		t.Fatalf("Print raw error: %v", err)
	}
	var parsed []testItem
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("raw Print produced invalid JSON: %v\nbuf=%s", err, buf.String())
	}
	if len(parsed) != 1 || parsed[0].ID != "1" {
		t.Errorf("raw Print content mismatch: %+v", parsed)
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
