package Leetcode;

import java.util.ArrayList;
import java.util.List;

public class Combination_Sum_III_216 {
    public static void main(String[] args) {

    }

    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> res = new ArrayList<>();
        backtrack(k, n, 1, 0, res, new ArrayList<>());
        return res;
    }

    public void backtrack(int k, int n, int start, int sum, List<List<Integer>> res, List<Integer> list){
        if (list.size() > k){
            return;
        }

        if (sum == n && list.size() == k){
            res.add(new ArrayList<>(list));
        }

        int minNum = Math.min(9 - start + 1, k - list.size());
        int leftSum = 0;
        for (int i = 9; i > 9 - minNum; i--){
            leftSum += i;
        }

        if (leftSum < n - sum){
            return;
        }

        for (int i = start; i <= 9; i++){
            sum += i;
            if (sum > n){
                sum -= i;
                break;
            }

            list.add(i);
            backtrack(k, n, i + 1, sum, res, list);
            list.removeLast();
            sum -= i;
        }
    }

    /**
     * 优化建议，最大可能和最小可能都要考虑，如果以这个点为起点得到的最小和都大于所需的，也剪掉。
     * int remainCount = k - list.size();
     * int maxPossibleSum = 0, minPossibleSum = 0;
     *
     * // 最大可能和（从9往下取）
     * for (int i = 9; i > 9 - remainCount; i--) {
     *     maxPossibleSum += i;
     * }
     *
     * // 最小可能和（从start往上取）
     * for (int i = start; i < start + remainCount; i++) {
     *     minPossibleSum += i;
     * }
     *
     * // 若总和已超范围，提前剪枝
     * if (sum + minPossibleSum > n || sum + maxPossibleSum < n) {
     *     return;
     * }
     */
}
