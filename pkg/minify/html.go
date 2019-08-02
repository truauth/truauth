package minify

import (
	"strings"
)

func minifyHTML(doc string) string {
	scripts := getStringInBetween(doc, "<script>", "</script>")

	formattableHTML := strings.Replace(doc, scripts, "__FormatID__", -1)
	formattedHTML := formatNormal(formattableHTML)

	return strings.Replace(formattedHTML, "__FormatID__", scripts, -1)
}
