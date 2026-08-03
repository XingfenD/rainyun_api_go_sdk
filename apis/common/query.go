package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
)

// MarshalQueryParams 基于 struct 字段的 json tag 自动构建 URL 查询参数 map。
//
// 规则：
//   - 字段必须有 json tag，tag 名（逗号前的部分）作为参数 key
//   - 指针字段为 nil 时跳过（对应 optional 语义）；非 nil 时解引用处理
//   - 基础类型（string / int / uint / float / bool）转为字符串
//   - 复杂类型（struct / map / slice 等）自动 marshal 为 JSON 字符串
//   - 无 json tag 或 tag 为 "-" 的字段被忽略
func MarshalQueryParams(v any) (map[string]string, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil, nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("MarshalQueryParams: 期望 struct，实际 %s", rv.Kind())
	}

	rst := make(map[string]string)
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")

		fv := rv.Field(i)
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}

		val, err := queryValue(fv)
		if err != nil {
			return nil, fmt.Errorf("MarshalQueryParams: 字段 %s: %w", field.Name, err)
		}
		rst[name] = val
	}
	return rst, nil
}

func queryValue(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	default:
		if !v.CanInterface() {
			return "", fmt.Errorf("无法读取未导出字段")
		}
		b, err := sonic.Marshal(v.Interface())
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
}
