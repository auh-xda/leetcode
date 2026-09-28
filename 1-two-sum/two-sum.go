func twoSum(nums []int, target int) []int {
   	seen, i := make(map[int]int), 0

	for i < len(nums) {
		diff := target - nums[i]
		seenIdx, exists := seen[diff]
		if exists {
			return []int{i, seenIdx}
		} else {
			seen[nums[i]] = i
		}
		i++
	}

	return []int{-1, -1}
}