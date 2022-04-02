package query

// ForEach 遍历每个值
// 如果结果表示不存在的值，则不会迭代任何值。
// 如果结果是Object，迭代器将传递每个项的键和值。如果结果是Array，迭代器将只传递每一项的值。
// 如果结果不是JSON数组或对象，迭代器将返回一个等于结果的值。
func (t Result) ForEach(iterator func(key, value Result) bool) {
	if !t.Exists() {
		return
	}
	if t.Type != JSON {
		iterator(Result{}, t)
		return
	}
	json := t.Raw
	var obj bool
	var i int
	var key, value Result
	for ; i < len(json); i++ {
		if json[i] == '{' {
			i++
			key.Type = String
			obj = true
			break
		} else if json[i] == '[' {
			i++
			key.Type = Number
			key.Num = -1
			break
		}
		if json[i] > ' ' {
			return
		}
	}
	var str string
	var vesc bool
	var ok bool
	var idx int
	for ; i < len(json); i++ {
		if obj {
			if json[i] != '"' {
				continue
			}
			s := i
			i, str, vesc, ok = parseString(json, i+1)
			if !ok {
				return
			}
			if vesc {
				key.Str = unescape(str[1 : len(str)-1])
			} else {
				key.Str = str[1 : len(str)-1]
			}
			key.Raw = str
			key.Index = s + t.Index
		} else {
			key.Num += 1
		}
		for ; i < len(json); i++ {
			if json[i] <= ' ' || json[i] == ',' || json[i] == ':' {
				continue
			}
			break
		}
		s := i
		i, value, ok = parseAny(json, i, true)
		if !ok {
			return
		}
		if t.Indexes != nil {
			if idx < len(t.Indexes) {
				value.Index = t.Indexes[idx]
			}
		} else {
			value.Index = s + t.Index
		}
		if !iterator(key, value) {
			return
		}
		idx++
	}
}
