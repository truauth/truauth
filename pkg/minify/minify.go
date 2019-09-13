package minify

func FromBytes(bytes *[]byte) string {
	byteStr := string(*bytes)

	mincss := minifyCSS(byteStr)
	minjs := minifyJs(mincss)
	minhtml := minifyHTML(minjs)

	return minhtml
}
