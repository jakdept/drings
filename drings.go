package drings

import (
	"sort"
	"strings"
)

// MaxLen returns the length of the longest string in an array
func MaxLen(arr []string) int {
	var i, max int
	for ; i < len(arr); i++ {
		if max < len(arr[i]) {
			max = len(arr[i])
		}
	}
	return max
}

// Dedup takes a string slice and returns the string slice sorted and deduplicated
func Dedup(s []string) []string {
	sort.Strings(s)
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
		} else {
			i++
		}
	}
	return s
}

// Copy returns a deepcopy of a stringslice
func DeepCopy(src []string) []string {
	var dst []string
	for _, each := range src {
		dst = append(dst, each)
	}
	return dst
}

// SplitAndTrimSpace splits a given string based on a delimiter, then trims
// spaces from every resulting string.
func SplitAndTrimSpace(in, sep string) []string {
	valueArr := strings.Split(in, sep)
	for i := 0; i < len(valueArr); {
		valueArr[i] = strings.TrimSpace(valueArr[i])
		if len(valueArr[i]) > 0 {
			// either it's not empty and go onto the next
			i++
		} else {
			// or eliminate this cell because it's empty
			valueArr = append(valueArr[:i], valueArr[i+1:]...)
		}
	}
	return valueArr
}

// PadRight pads a string with the given padding string, to make a string of the
// given length. The padding string may be cut off in the middle if it is
// multiple characters, but the returned string will be of the given length.
func PadRight(s, pad string, length int) string {
	if pad == "" || length < len(s) {
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
	if pad == "" || length < len(s) {
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
