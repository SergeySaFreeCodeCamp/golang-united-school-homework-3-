package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, v := range keys {
		result = append(result, input[v])
	}

	return
}
