package model

// RawData 包裹无公开结构的响应数据,保证 table 模式有输出
// (内容为 map/slice 时打印原始值,详细结构请用 -o json)。
type RawData struct {
	Data any `json:"data" table:"DATA"`
}
