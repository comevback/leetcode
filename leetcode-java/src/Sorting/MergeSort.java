package Sorting;

public class MergeSort {
    public static void main(String[] args) {
        int[] arr = new int[]{8, 10, 1, 3, 15, 18, 2, 6};
        mergeSort(arr, 0, arr.length);
        for (int num : arr) {
            System.out.println(num);
        }
    }

    public static void mergeSort(int[] nums, int left, int right) {
        if (left >= right - 1) {
            return;
        }

        int mid = left + (right - left) / 2;
        mergeSort(nums, left, mid);
        mergeSort(nums, mid, right);
        merge(nums, left, mid, right);
    }

    public static void merge(int[] nums, int left, int mid, int right) {
        int[] sorted = new int[right - left];
        int index = 0;
        int pl = left;
        int pr = mid;

        while (pl < mid && pr < right) {
            if (nums[pl] <= nums[pr]) {
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
