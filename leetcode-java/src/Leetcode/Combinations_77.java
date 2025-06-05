package Leetcode;

import java.util.ArrayList;
import java.util.List;

public class Combinations_77 {
    public static void main(String[] args) {
        Combinations_77 app = new Combinations_77();
        List<List<Integer>> res = app.combine(13,13);
        for (List<Integer> list : res){
            for (Integer num : list){
                System.out.println(num);
            }
        }
    }

    public List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> res = new ArrayList<>();
        backtrack(n, k,1, new ArrayList<>(), res);
        return res;
    }

    public void backtrack(int n, int k, int start, List<Integer> path, List<List<Integer>> res){
        if (path.size() >= k){
            res.add(new ArrayList<>(path));
        }else{
            for (int i = start; i <= n - (k - path.size()) + 1; i++){
                path.add(i);
                backtrack(n, k, i+1, path, res);
                path.removeLast();
            }
        }
    }
}