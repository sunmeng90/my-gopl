package main

import (
	"bytes"
)
import "strings"

func comma(s string) string {
	var buf bytes.Buffer
	intSign(&buf, &s)
	firstNumIdx := 0
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		firstNumIdx = 1
	}
	dotIdx := strings.LastIndex(s, ".")
	if dotIdx == -1 {
		intComma(&buf, &s, firstNumIdx, len(s)-1)
	} else {
		intComma(&buf, &s, firstNumIdx, dotIdx-1)
		intAfterDot(&buf, &s, dotIdx+1)
	}
	return buf.String()
}

func intSign(buf *bytes.Buffer, s *string) {
	if strings.HasPrefix(*s, "-") {
		buf.WriteByte('-')
	}
}

// add comma to integers
func intComma(buf *bytes.Buffer, s *string, start int, end int) {
	re := (end - start + 1) % 3
	intPart1EndIdx := start + re - 1
	for i := start; i <= end; i++ {
		buf.WriteByte((*s)[i])
		if (i != end) && (i == intPart1EndIdx || (i-intPart1EndIdx)%3 == 0) {
			buf.WriteString(",")
		}
	}
}

func intAfterDot(buf *bytes.Buffer, s *string, idxAfterDot int) {
	n := len(*s)
	buf.WriteByte('.')
	for i := idxAfterDot; i < n; i++ {
		buf.WriteByte((*s)[i])
	}
}
