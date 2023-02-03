func twoSum(nums []int, target int) []int {
    var result []int
    for i, value := range nums {
        for j := i + 1; j < len(nums); j++ {
            if int64(value) + int64(nums[j]) == int64(target) {
                result = []int{i, j}
                return result
            }
        }
    }
    return result
}
