package Leetcode;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

public class Valid_Sudoku_36 {
    public static void main(String[] args) {
        Valid_Sudoku_36 solution = new Valid_Sudoku_36();
        char[][] board = {
            {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
            {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
            {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
            {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
            {'4', '.', '6', '8', '.', '3', '.', '.', '1'},
            {'7', '.', '.', '2', '.', '5', '.', '.', '.'},
            {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
            {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
            {'.', '.', '.', '.', '8', '.', '.', '7', '9'}
        };
        boolean res = solution.isValidSudoku(board);
        System.out.println(res); // Output: true
    }

    public boolean isValidSudoku(char[][] board) {
        boolean res = true;
        List<HashSet> matMap = new ArrayList<>();
        List<HashSet> rowMap = new ArrayList<>();
        List<HashSet> columnMap = new ArrayList<>();

        for (int i = 0; i < 9; i++){
            matMap.add(new HashSet<>());
            rowMap.add(new HashSet<>());
            columnMap.add(new HashSet<>());
        }

        for (int i = 0; i < board.length; i++){
            for (int j = 0; j < board[0].length; j++){
                if (board[i][j] == '.'){
                    continue;
                }
                int num = Integer.parseInt(Character.toString(board[i][j]));
                int matNum = 0;
                if (i < 3){
                    if (j < 3){
                        matNum = 0;
                    } else if (j > 2 && j < 6) {
                        matNum = 1;
                    } else if (j > 5) {
                        matNum = 2;
                    }
                } else if (i > 2 && i < 6) {
                    if (j < 3){
                        matNum = 3;
                    } else if (j > 2 && j < 6) {
                        matNum = 4;
                    } else if (j > 5) {
                        matNum = 5;
                    }
                } else if (i > 5) {
                    if (j < 3){
                        matNum = 6;
                    } else if (j > 2 && j < 6) {
                        matNum = 7;
                    } else if (j > 5) {
                        matNum = 8;
                    }
                }

                if (rowMap.get(i).contains(num)){
                    return false;
                }

                if (columnMap.get(j).contains(num)){
                    return false;
                }

                if (matMap.get(matNum).contains(num)){
                    return false;
                }

                rowMap.get(i).add(num);
                columnMap.get(j).add(num);
                matMap.get(matNum).add(num);
            }
        }

        return res;
    }
}
