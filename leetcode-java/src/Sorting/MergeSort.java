package Sorting;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class MergeSort {
    public static void main(String[] args) {
        int[][] arr = new int[][]{{8, 10}, {1, 3}, {15, 18}, {2, 6}};
        mergeSort(arr, 0, arr.length);
        for (int[] num : arr) {
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

    public static void mergeSort(int[][] nums, int left, int right) {
        if (left >= right - 1) {
            return;
        }

        int mid = left + (right - left) / 2;
        mergeSort(nums, left, mid);
        mergeSort(nums, mid, right);
        merge(nums, left, mid, right);
    }

    public static void merge(int[][] nums, int left, int mid, int right) {
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
