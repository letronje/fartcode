package main

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	mem := map[int]struct{}{}

	for _, num := range nums {
		if _, found := mem[num]; found {
			return true
		}
		mem[num] = struct{}{}
	}

	return false
}

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sc := map[rune]int{}
	tc := map[rune]int{}

	for _, c := range []rune(s) {
		sc[c] += 1
	}

	for _, c := range []rune(t) {
		tc[c] += 1
	}

	for c, cnt := range sc {
		if tc[c] != cnt {
			return false
		}
	}

	return true
}

func main() {

}
