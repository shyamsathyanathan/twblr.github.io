package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	result := []int32{}
	for _, value := range vals {
		result = append(result, op(value))
	}
	return result
}

func filterInts(op filterOperation, vals []int32) []int32 {
	result := []int32{}
	for _, value := range vals {
		if op(value) {
			result = append(result, value)
		}
	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {
	for _, value := range newValues {
		dest = append(dest, value)
	}
	return dest
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
