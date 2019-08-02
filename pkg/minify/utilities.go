package minify

import (
	"math"
	"strings"
)

func genSeedString(seed int) string {
	alp := []rune("abcdefghijklmnopqrstuvwxyz")

	if seed > len(alp) {
		genSeed := ""
		seedCount := math.Abs(float64(len(alp)) / float64(seed))

		for idx := 0; idx < int(seedCount); idx++ {
			genSeed = "_" + string(alp[idx])
		}

		return genSeed
	}

	return string(alp[seed])
}

func getStringInBetween(str string, start string, end string) (result string) {
    s := strings.Index(str, start)
    if s == -1 {
        return
    }
    s += len(start)
    e := strings.Index(str, end)
    return str[s:e]
}

func formatNormal(str string) string {
	nowhitespace := strings.ReplaceAll(str, "  ", "")

	return strings.ReplaceAll(nowhitespace, "\n", "")
}