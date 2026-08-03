package common

import (
	"reflect"
	"testing"
)

type queryParamsSample struct {
	IsRGpu  *bool                   `json:"is_rgpu"`
	Limit   int                     `json:"limit"`
	Name    string                  `json:"name"`
	Options StandQueryParameters    `json:"options"`
	Tags    []string                `json:"tags"`
	Skip    string                  `json:"-"`
	NoTag   string
}

func TestMarshalQueryParams(t *testing.T) {
	isGPU := true
	s := &queryParamsSample{
		IsRGpu: &isGPU,
		Limit:  10,
		Name:   "hello",
		Options: StandQueryParameters{
			Page:    2,
			PerPage: 20,
		},
		Tags: []string{"a", "b"},
	}

	m, err := MarshalQueryParams(s)
	if err != nil {
		t.Fatalf("MarshalQueryParams() error = %v", err)
	}

	want := map[string]string{
		"is_rgpu": "true",
		"limit":   "10",
		"name":    "hello",
		"options": `{"columnFilters":{},"sort":null,"page":2,"perPage":20}`,
		"tags":    `["a","b"]`,
	}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("MarshalQueryParams() = %v, want %v", m, want)
	}
}

func TestMarshalQueryParamsNilPtr(t *testing.T) {
	s := &queryParamsSample{} // IsRGpu 为 nil
	m, err := MarshalQueryParams(s)
	if err != nil {
		t.Fatalf("MarshalQueryParams() error = %v", err)
	}

	// nil 指针跳过，无 json tag 和 "-" 的字段跳过
	if _, ok := m["is_rgpu"]; ok {
		t.Errorf("nil 指针字段 is_rgpu 不应输出")
	}
	if _, ok := m["Skip"]; ok {
		t.Errorf("tag 为 '-' 的字段不应输出")
	}
	if _, ok := m["NoTag"]; ok {
		t.Errorf("无 json tag 的字段不应输出")
	}
}

func TestMarshalQueryParamsNonStruct(t *testing.T) {
	_, err := MarshalQueryParams(42)
	if err == nil {
		t.Errorf("非 struct 入参应返回错误")
	}
}
