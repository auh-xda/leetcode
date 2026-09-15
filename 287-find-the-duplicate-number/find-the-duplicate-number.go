func findDuplicate(nums []int) int {

    seen := make(map[int]int)

    for _, num := range nums {
        if seen[num] == 1 {
            return num
        }
        seen[num]++
    }

    return -1
}