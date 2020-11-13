package service

func isStrInStrSlice(base []string, target string) bool {
	for _, b := range base {
		if b == target {
			return true
		}
	}
	return false
}
