package sumtozero

type SumToZero struct {
	NumberList []int
}

func (stz *SumToZero) RemoveZeroSumSublists() []int {
	if stz.NumberList == nil {
		return []int{}
	}

	var stack []int
	sumMap := make(map[int]int)
	prefixSum := 0

	for _, num := range stz.NumberList {
		prefixSum += num

		if startIndex, found := sumMap[prefixSum]; found {
			// Remove the subarray from the stack
			stack = stack[:startIndex]
		} else if prefixSum == 0 {
			// Clear the stack if the prefix sum is zero
			stack = nil
		} else {
			// Append the current number to the stack and update the sumMap
			stack = append(stack, num)
			sumMap[prefixSum] = len(stack)
		}
	}

	return stack
}
