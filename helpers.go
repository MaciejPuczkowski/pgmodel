package pgmodel

import (
	"bytes"
	"text/template"
	"unicode"

	"github.com/MaciejPuczkowski/errs"
)

func shouldQuoteIdentifier(s string) bool {
	if isKeyword(s) {
		return true
	}
	for _, c := range s {
		if unicode.IsUpper(c) || unicode.IsDigit(c) {
			return true
		}
	}
	return false
}
func isKeyword(s string) bool {
	var keywords map[string]struct{}
	if _, ok := keywords[s]; ok {
		return true
	}
	return false
}

func fromTemplate(tmpl *template.Template, data any) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", errs.Wrap(err)
	}
	return buf.String(), nil
}

func quoteIdentifier(s string) string {
	return `"` + s + `"`
}
