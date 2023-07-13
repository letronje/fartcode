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

func twoSumOld(nums []int, target int) []int {
	m := map[int][]int{}

	for idx, n := range nums {
		m[n] = append(m[n], idx)
	}

	for idx, n := range nums {
		d := target - n

		indices, found := m[d]
		if !found {
			continue
		}

		if d == n {
			if len(indices) >= 2 {
				otherIndex := -1
				for _, idx2 := range indices {
					if idx2 != idx {
						otherIndex = idx2
						break
					}
				}
				return []int{idx, otherIndex}
			} else {
				continue
			}

		} else {
			if len(indices) == 1 {
				return []int{idx, indices[0]}
			}
		}
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for idx, n := range nums {

		d := target - n
		if otherIdx, found := m[d]; found {
			return []int{otherIdx, idx}
		}

		m[n] = idx
	}

	return nil
}

func main() {

}
