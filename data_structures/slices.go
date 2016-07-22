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
	if len(list1) != len(list2) {
		return false
	}
	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func partialReverse(src []int, from, to int) []int {
	partialSlice := []int {}
	for i := from; i <= to; i ++ {
		partialSlice = append(partialSlice, src[i])
	}
	for i, j := 0, len(partialSlice) - 1; i < j; i, j = i + 1, j - 1 {
		partialSlice[i], partialSlice[j] = partialSlice[j], partialSlice[i]
	}
	return partialSlice
}
