package Leetcode;

import java.util.ArrayList;
import java.util.List;

public class Combination_Sum_39 {
    public static void main(String[] args) {
        Combination_Sum_39 app = new Combination_Sum_39();
        int[] candidates = {2,3,6,7};
        int target = 7;
        List<List<Integer>> res = app.combinationSum(candidates, target);
        for (List<Integer> list : res){
            for (Integer num : list){
                System.out.print(num + " ");
            }
            System.out.println();
        }
    }

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        backtrack(candidates, target, 0, 0, new ArrayList<>(), res);
        return res;
    }

    public void backtrack(int[] candidates, int target, int start, int sum, List<Integer> list, List<List<Integer>> res){
        if (sum == target){
            res.add(new ArrayList<>(list));
        } else if (sum > target) {
            return;
        }

        for (int i = start; i< candidates.length; i++){
            sum += candidates[i];
            if (sum > target) {
                sum -= candidates[i];
                continue;
            }
            list.add(candidates[i]);
            backtrack(candidates, target, i, sum, list, res);
            list.removeLast();
            sum -= candidates[i];
        }
    }
}
