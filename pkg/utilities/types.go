package utilities

// ArrContainsString check if str array contains string
func ArrContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
