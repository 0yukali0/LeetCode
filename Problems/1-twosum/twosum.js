/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var result = [0, 0];
    for (var i = 0, len = nums.length, find = false; i < len && !find; i++) {
        for (var j = i + 1; j < len && !find; j++) {
            if (nums[i] + nums[j] == target) {
                result[0] = i;
                result[1] = j;
                find = true;
            }
        }
    }
    return result;
};
