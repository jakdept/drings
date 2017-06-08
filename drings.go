package drings

import "strings"

// SplitAndTrimSpace splits a given string based on a delimiter, then trims
// spaces from every resulting string.
func SplitAndTrimSpace(in, sep string) []string {
	valueArr := strings.Split(in, sep)
	for i := 0; i < len(valueArr); {
		valueArr[i] = strings.TrimSpace(valueArr[i])
		if valueArr[i] != "" {
			// either it's not empty and increment
			i++
		} else {
			// or eliminate this cell because it's empty
			if i < len(valueArr)-1 {
				valueArr = append(valueArr[:i], valueArr[i+1])
			} else {
				valueArr = valueArr[:i]
			}
		}
	}
	return valueArr
}

// MaxLen returns the length of the longest string in an array
func MaxLen(arr []string) int {
	var i, max int
	for i < len(arr) {
		if len(arr[i]) > max {
			max = len(arr[i])
		}
	}
	return max
}

// PadRight pads a string with the given padding string, to make a string of the
// given length. The padding string may be cut off in the middle if it is
// multiple characters, but the returned string will be of the given length.
func PadRight(s, pad string, length int) string {
	if pad == "" {
		return s
	}
	for len(s) < length {
		s = s + pad
	}
	return s[:length]
}

// PadLeft pads a string with the given padding string, to make a string of the
// given length. The padding string may be cut off in the middle if it is
// multiple characters, but the returned string will be of the given length.
func PadLeft(s, pad string, length int) string {
	if pad == "" {
		return s
	}
	for len(s) < length {
		s = pad + s
	}
	return s[len(s)-length:]
}

// PadAllRight pads all strings in a slice with the given padding string, to
// make strings of the given length. It calls PadRight on each string.
func PadAllRight(s []string, pad string, length int) []string {
	for i := 0; i < len(s); i++ {
		s[i] = PadRight(s[i], pad, length)
	}
	return s
}

// PadAllLeft pads all strings in a slice with the given padding string, to
// make strings of the given length. It calls PadLeft on each string.
func PadAllLeft(s []string, pad string, length int) []string {
	for i := 0; i < len(s); i++ {
		s[i] = PadLeft(s[i], pad, length)
	}
	return s
}
