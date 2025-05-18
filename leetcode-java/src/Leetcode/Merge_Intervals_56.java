package Leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * 56. Merge Intervals
 * 思路：首先把二维数组按照第一个元素进行排序，然后遍历这个数组，判断当前的区间是否和上一个区间重叠，
 * 如果重叠就更新上一个区间的结束位置，否则就把当前区间加入结果中，另开始一个新的区间。
 */

public class Merge_Intervals_56 {
    public static void main(String[] args) {
        int[][] arr = new int[][]{{8, 10}, {1, 3}, {15, 18}, {2, 6}};
        Merge_Intervals_56 ms = new Merge_Intervals_56();
        int[][] res = ms.merge(arr);
        for (int[] num : res) {
            System.out.println(Arrays.toString(num));
        }
    }

    public int[][] merge(int[][] intervals) {
        List<int[]> res = new ArrayList<>();
        mergeSort(intervals, 0, intervals.length);
        int head = intervals[0][0];
        int tail = intervals[0][1];

        for (int i = 0; i < intervals.length; i++) {
            if (intervals[i][0] > tail) {
                res.add(new int[]{head, tail});
                head = intervals[i][0];
                tail = intervals[i][1];
            } else if (intervals[i][1] > tail) {
                tail = intervals[i][1];
            }
        }

        res.add(new int[]{head, tail});
        int[][] arrRes = res.toArray(new int[res.size()][]);

        return arrRes;
    }

    public void mergeSort(int[][] nums, int left, int right) {
        if (left >= right - 1) {
            return;
        }

        int mid = left + (right - left) / 2;
        mergeSort(nums, left, mid);
        mergeSort(nums, mid, right);
        mergeInSort(nums, left, mid, right);
    }

    public void mergeInSort(int[][] nums, int left, int mid, int right) {
        int[][] sorted = new int[right - left][];
        int index = 0;
        int pl = left;
        int pr = mid;

        while (pl < mid && pr < right) {
            if (nums[pl][0] <= nums[pr][0]) {
                sorted[index] = nums[pl];
                pl += 1;
            } else {
                sorted[index] = nums[pr];
                pr += 1;
            }
            index += 1;
        }

        while (pl < mid) {
            sorted[index] = nums[pl];
            pl += 1;
            index += 1;
        }

        while (pr < right) {
            sorted[index] = nums[pr];
            pr += 1;
            index += 1;
        }

        System.arraycopy(sorted, 0, nums, left, sorted.length);
    }
}
