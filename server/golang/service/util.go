package service

func isStrInStrSlice(base []string, target string) bool {
	for _, b := range base {
		if b == target {
			return true
		}
	}
	return false
}

func mapToSlice(targetMap map[string]int64) []string {
	var slice []string
	for key := range targetMap {
		slice = append(slice, key)
	}
	return slice
}

func sliceToMap(targetSlice []string) map[string]int64 {
	var newMap = map[string]int64{}
	for idx, val := range targetSlice {
		newMap[val] = int64(idx)
	}
	return newMap
}

func isStrInMap(baseMap map[string]int64, target string) bool {
	_, ok := baseMap[target]
	return ok
}
