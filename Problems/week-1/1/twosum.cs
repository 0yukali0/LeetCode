public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        int[] result = new int[] {0,0};
        bool find = false;
        for (int i = 0, len = nums.Length; i < len && !find; i++) {
            for (int j = i + 1; j < len && !find; j++) {
                if (nums[i] + nums[j] == target) {
                    find = true;
                    result = new int[] {i, j};
                }
            }
        }
        return result;
    }
}
