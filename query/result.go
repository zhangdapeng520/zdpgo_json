package query

import (
	"strconv"
	"strings"
	"time"
)

// Result 表示从Get()返回的json值。
type Result struct {
	Type    Type    // json类型
	Raw     string  // 格式化的json
	Str     string  // 字符串的json
	Num     float64 // json数字
	Index   int     // 对于原始json中的原始值，0表示索引未知
	Indexes []int   // 在包含“#”查询字符的路径上匹配的所有元素。
}

// String 返回值的字符串表示形式。
func (t Result) String() string {
	switch t.Type {
	default:
		return ""
	case False:
		return "false"
	case Number:
		if len(t.Raw) == 0 {
			// calculated result
			return strconv.FormatFloat(t.Num, 'f', -1, 64)
		}
		var i int
		if t.Raw[0] == '-' {
			i++
		}
		for ; i < len(t.Raw); i++ {
			if t.Raw[i] < '0' || t.Raw[i] > '9' {
				return strconv.FormatFloat(t.Num, 'f', -1, 64)
			}
		}
		return t.Raw
	case String:
		return t.Str
	case JSON:
		return t.Raw
	case True:
		return "true"
	}
}

// Bool 返回一个布尔表示。
func (t Result) Bool() bool {
	switch t.Type {
	default:
		return false
	case True:
		return true
	case String:
		b, _ := strconv.ParseBool(strings.ToLower(t.Str))
		return b
	case Number:
		return t.Num != 0
	}
}

// Int 返回一个int表示
func (t Result) Int() int64 {
	switch t.Type {
	default:
		return 0
	case True:
		return 1
	case String:
		n, _ := parseInt(t.Str)
		return n
	case Number:
		// 尝试直接将float64转换为int64
		if i, ok := safeInt(t.Num); ok {
			return i
		}

		// 现在试着解析原始字符串
		if i, ok := parseInt(t.Raw); ok {
			return i
		}

		// 回退到标准转换
		return int64(t.Num)
	}
}

// Uint 返回无符号整数表示形式。
func (t Result) Uint() uint64 {
	switch t.Type {
	default:
		return 0
	case True:
		return 1
	case String:
		n, _ := parseUint(t.Str)
		return n
	case Number:
		// try to directly convert the float64 to uint64
		i, ok := safeInt(t.Num)
		if ok && i >= 0 {
			return uint64(i)
		}
		// now try to parse the raw string
		u, ok := parseUint(t.Raw)
		if ok {
			return u
		}
		// fallback to a standard conversion
		return uint64(t.Num)
	}
}

// Float 返回一个float64的表示形式
func (t Result) Float() float64 {
	switch t.Type {
	default:
		return 0
	case True:
		return 1
	case String:
		n, _ := strconv.ParseFloat(t.Str, 64)
		return n
	case Number:
		return t.Num
	}
}

// Time 返回一个 time.Time 的表示形式
func (t Result) Time() time.Time {
	res, _ := time.Parse(time.RFC3339, t.String())
	return res
}

// Array 返回一个array数组表示形式
// 如果结果表示空值或不存在，则返回一个空数组。
// 如果结果不是一个JSON数组，返回值将是一个包含一个结果的数组。
func (t Result) Array() []Result {
	if t.Type == Null {
		return []Result{}
	}
	if !t.IsArray() {
		return []Result{t}
	}
	r := t.arrayOrMap('[', false)
	return r.a
}

// Map 返回值的映射。结果应该是一个JSON对象。如果结果不是一个JSON对象，返回值将是一个空映射。
func (t Result) Map() map[string]Result {
	if t.Type != JSON {
		return map[string]Result{}
	}
	r := t.arrayOrMap('{', false)
	return r.o
}

// Value 返回以下类型之一: bool, float64, Number, string, nil, map[string]interface{}, []interface{}
func (t Result) Value() interface{} {
	if t.Type == String {
		return t.Str
	}
	switch t.Type {
	default:
		return nil
	case False:
		return false
	case Number:
		return t.Num
	case JSON:
		r := t.arrayOrMap(0, true)
		if r.vc == '{' {
			return r.oi
		} else if r.vc == '[' {
			return r.ai
		}
		return nil
	case True:
		return true
	}
}

// IsObject 判断结果是不是JSON对象
func (t Result) IsObject() bool {
	return t.Type == JSON && len(t.Raw) > 0 && t.Raw[0] == '{'
}

// IsArray 判断结果是不是一个JSON数组
func (t Result) IsArray() bool {
	return t.Type == JSON && len(t.Raw) > 0 && t.Raw[0] == '['
}

// Exists 如果值存在，返回true。
//
//  if query.Get(json, "name.last").Exists(){
//		println("value exists")
//  }
func (t Result) Exists() bool {
	return t.Type != Null || len(t.Raw) != 0
}
