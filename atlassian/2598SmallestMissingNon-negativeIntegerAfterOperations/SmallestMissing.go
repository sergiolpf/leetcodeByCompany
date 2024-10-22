package smallestmissing

func findSmallestInteger(nums []int, value int) int {

	frequency := make([]int, value)

	for _, num := range nums {

		//module := (value + num%value) % value
		module := num % value
		if module < 0 {
			module += value
		}

		frequency[module]++
	}

	minn := 0
	for i := 0; i < value; i++ {
		if frequency[i] < frequency[minn] {
			minn = i
		}
	}

	return value*frequency[minn] + minn

}
