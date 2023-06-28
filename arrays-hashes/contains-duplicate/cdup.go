package containsdup

func ContainsDuplicate(nums []int) bool {
	mem := map[int]struct{}{}

	for _, num := range nums {
		if _, found := mem[num]; found {
			return true
		}
		mem[num] = struct{}{}
	}

	return false
}
