package cat

import "bytes"

func cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

func buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
