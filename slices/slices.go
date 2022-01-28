package slices

func Sum(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumAll(slices ...[]int) (output []int) {

	for _, s := range slices {
		// current := 0
		// for _, num := range s {
		//     current += num
		// }
		// output = append(output, current)
		output = append(output, Sum(s))
	}
	return output
}
