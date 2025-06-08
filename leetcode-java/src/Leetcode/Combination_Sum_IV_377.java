package Leetcode;

public class Combination_Sum_IV_377 {
    public static void main(String[] args) {
        int[] arr = {4,2,1};
        int target = 32;
        Combination_Sum_IV_377 app = new Combination_Sum_IV_377();
        int res = app.combinationSum4(arr, target);
        System.out.println(res);
    }

    public int combinationSum4(int[] nums, int target) {
        int[] dp = new int[target + 1];
        dp[0] = 1;
        for (int i = 0; i <= target ; i++){
            for (int j = 0; j < nums.length; j++){
                if (nums[j] <= i){
                    dp[i] += dp[i - nums[j]];
                }
            }
        }
        return dp[target];
    }
}
