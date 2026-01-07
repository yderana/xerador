package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

var splitter = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func words(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := splitter.Split(s, -1)

	var out []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		// pecah camelCase / PascalCase jadi kata
		var buff []rune
		for i, r := range []rune(p) {
			if i > 0 && unicode.IsUpper(r) && (i+1 < len([]rune(p)) && unicode.IsLower([]rune(p)[i+1])) {
				out = append(out, strings.ToLower(string(buff)))
				buff = buff[:0]
			}
			buff = append(buff, r)
		}
		if len(buff) > 0 {
			out = append(out, strings.ToLower(string(buff)))
		}
	}
	return out
}

func UpperFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func CamelCase(s string) string {
	w := words(s)
	if len(w) == 0 {
		return ""
	}
	for i := range w {
		if i == 0 {
			w[i] = strings.ToLower(w[i])
		} else {
			w[i] = UpperFirst(strings.ToLower(w[i]))
		}
	}
	return strings.Join(w, "")
}

func KebabCase(s string) string {
	w := words(s)
	if len(w) == 0 {
		return ""
	}
	return strings.Join(w, "-")
}

func StartCase(s string) string {
	ws := words(s) // gunakan fungsi words internal yang sudah kamu punya untuk CamelCase/KebabCase
	if len(ws) == 0 {
		return ""
	}
	for i := range ws {
		ws[i] = UpperFirst(strings.ToLower(ws[i]))
	}
	return strings.Join(ws, " ")
}
