function twoSum(nums: number[], target: number): number[] {
    var result: number[];
    for (var i = 0, len = nums.length, find = false; i < len && !find; i++) {
        for (var j = i + 1; j < len && !find; j++) {
            if (nums[i] + nums[j] == target) {
                result = [i, j];
                find = true;
            };
        };
    };
    return result
};
