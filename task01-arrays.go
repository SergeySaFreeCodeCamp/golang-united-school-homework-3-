package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	result = float32(0)

	for indx1, indx2 := 0, len(input)-1; indx1 < indx2; {
		result += input[indx1] + input[indx2]
		indx1++
		indx2--
		if indx1 == indx2 {
			result += input[indx1]
		}
	}
	result = result / float32(len(input))
	return
}
