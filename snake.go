package strcase

import "strings"

func SnakeCase(str string) string {
	return CaseDelimitedScreaming(str, '_', false)
}

func SnakeCaseScreaming(str string) string {
	return CaseDelimitedScreaming(str, '_', true)
}

func KebabCase(str string) string {
	return CaseDelimitedScreaming(str, '-', false)
}

func KebabCaseScreaming(str string) string {
	return CaseDelimitedScreaming(str, '-', true)
}

// 将字符串中的符号分割转换
// @param str		需要处理的字符
// @param delimiter 分割符
// @param screaming 是否大写转换
func CaseDelimitedScreaming(str string, delimiter uint8, screaming bool) string {
	str = strings.TrimSpace(str)
	n := strings.Builder{}
	n.Grow(len(str) + 2) // nominal 2 bytes of extra space for inserted delimiters
	for i, v := range []byte(str) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow && screaming {
			v += 'A'
			v -= 'a'
		} else if vIsCap && !screaming {
			v += 'a'
			v -= 'A'
		}

		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if i+1 < len(str) {
			next := str[i+1]
			vIsNum := v >= '0' && v <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			// add underscore if next letter case type is changed
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
				if vIsCap && nextIsLow {
					if prevIsCap := i > 0 && str[i-1] >= 'A' && str[i-1] <= 'Z'; prevIsCap {
						n.WriteByte(delimiter)
					}
				}
				n.WriteByte(v)
				if vIsLow || vIsNum || nextIsNum {
					n.WriteByte(delimiter)
				}
				continue
			}
		}

		if v == ' ' || v == '_' || v == '-' || v == '.' {
			n.WriteByte(delimiter)
		} else {
			n.WriteByte(v)
		}
	}

	return n.String()
}
