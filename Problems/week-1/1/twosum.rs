impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut result: Vec<i32> = vec![0, 0];
        let mut i: usize = 0;
        let mut j: usize;
        let totalNum = nums.len() as usize;
        while i < totalNum {
            j = i + 1;
            while j < totalNum {
                let sum = nums[i] + nums[j];
                if sum == target {
                    result[0] = i as i32;
                    result[1] = j as i32;
                }
                j = j + 1;
            }
            i = i + 1;
        }
        result
    }
}
