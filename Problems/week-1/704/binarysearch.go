func search(nums []int, target int) int {
    result, left, right := -1, 0, len(nums)-1
    for left <= right {
        middle := int((right + left)/2)
        if num := nums[middle]; num == target {
            result = middle
            break
        } else if num > target {
            right = middle - 1
        } else {
            left = middle + 1
        }
    }
    return result
}
