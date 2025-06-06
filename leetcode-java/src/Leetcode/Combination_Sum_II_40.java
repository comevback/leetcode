package Leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Combination_Sum_II_40 {
    public static void main(String[] args) {
        Combination_Sum_II_40 app = new Combination_Sum_II_40();
        int[] candidates = {10,1,2,7,6,1,5};
        int target = 8;
        List<List<Integer>> res = app.combinationSum2(candidates, target);
        for (List<Integer> list : res){
            for (Integer num : list){
                System.out.print(num + " ");
            }
            System.out.println();
        }
    }

    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> res = new ArrayList<>();
        backtrack(candidates, target, 0, 0, res, new ArrayList<>());

        return res;
    }

    public void backtrack(int[] candidates, int target, int start, int sum, List<List<Integer>> resSet, List<Integer> list){
        if (sum == target){
            resSet.add(new ArrayList<>(list));
            return;
        }

        for (int i = start; i < candidates.length; i++){
            if (i > start && candidates[i] == candidates[i-1]) continue;
            sum += candidates[i];
            if (sum > target){
                sum -= candidates[i];
                continue;
            }
            list.add(candidates[i]);
            backtrack(candidates, target, i + 1, sum, resSet, list);
            list.removeLast();
            sum -= candidates[i];
        }
    }
}
