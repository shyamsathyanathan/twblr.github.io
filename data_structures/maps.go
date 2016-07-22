package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	result := []string {}
	for key, value := range data {
		for _, itst := range value {
			if itst == interest {
				result = append(result, key)
			}
		}
	}
	return result
}

func contains(src []string, elem string) bool {
	for _, value := range src {
		if value == elem {
			return true
		}
	}
	return false
}
