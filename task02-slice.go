package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	result = append(result, input...)

	for indx1, indx2 := 0, len(input)-1; indx1 < indx2; {
		result[indx1], result[indx2] = input[indx2], input[indx1]
		indx1++
		indx2--
		if indx1 == indx2 {
			result[indx1] = input[indx1]
		}
	}

	return
}
