package minify

import (
	"strings"
	"strconv"
)

func minifyCSS(doc string) string {
	if !strings.Contains(doc, "<style>") {
		return doc
	}
	originalStyles := getStringInBetween(doc, "<style>", "</style>")
	styleDoc := getStringInBetween(doc, "<style>", "</style>")
	updatedStyles := getStringInBetween(doc, "<style>", "</style>")

	resolved := map[string]string{}
	
	completed := false
	for !completed {
		indexOfPossibleClass := strings.Index(styleDoc, ".")

		if indexOfPossibleClass < 0 {
			completed = true
			break
		}
		
		possibleClass := strings.Split(styleDoc[indexOfPossibleClass+1:], " ")[0]
		styleDoc = styleDoc[indexOfPossibleClass+1:]

		firstChar := string([]rune(possibleClass)[0])

		_, err := strconv.Atoi(firstChar)
		ok := resolved[possibleClass]; if ok != "" || err == nil {
			continue
		}

		commaExist := false
		if strings.Contains(possibleClass, ",") {
			commaExist = true
			possibleClass = strings.Replace(possibleClass, ",", "", -1)
		}

		if strings.Contains(possibleClass, "hover") {
			possibleClass = strings.Replace(possibleClass, ":hover", "", -1)
			possibleClass = strings.Replace(possibleClass, ",", "", -1)

			val := resolved[possibleClass]; if val == "" {
				val = genSeedString(len(resolved))
			}
			
			resolved[possibleClass] = val
		
			eod := ":hover"
			if commaExist {
				eod = eod + ","
			}

			updatedStyles = strings.ReplaceAll(updatedStyles, possibleClass + eod, val + eod)

			doc = strings.ReplaceAll(doc, possibleClass + `"`, val + `"`) // does the html
			continue;
		}

		generated := genSeedString(len(resolved))
		resolved[possibleClass] = generated

		eod := " "
		if commaExist {
			eod = eod + ","
		}

		updatedStyles = strings.ReplaceAll(updatedStyles, possibleClass + eod, generated)

		doc = strings.ReplaceAll(doc, possibleClass + `"`, generated + `"`) // does the html
	}

	return strings.Replace(doc, originalStyles, formatNormal(updatedStyles), -1)
}
