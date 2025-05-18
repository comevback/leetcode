package Sorting;

public class QuickSort {
    public static void main(String[] args) {
        int[] arr = new int[]{8, 10, 1, 8, 3, 15, 18, 2, 8, 6};
        QuickSort qs = new QuickSort();
        qs.quickSort(arr, 0, arr.length);
        for (int num : arr) {
            System.out.println(num);
        }
    }

    public void quickSort(int[] nums, int left, int right) {
        if (left >= right - 1) {
            return;
        }
        int pivot = nums[left];
        int head = left;
        int tail = right - 1;
        while (head < tail) {
            while (tail > head && nums[tail] >= pivot) {
                tail -= 1;
            }
            while (head < tail && nums[head] <= pivot) {
                head += 1;
            }

            swap(nums, head, tail);
        }

        swap(nums, head, left);

        quickSort(nums, left, head);
        quickSort(nums, head + 1, right);
    }

    public void swap(int[] nums, int pos1, int pos2) {
        int temp = nums[pos1];
        nums[pos1] = nums[pos2];
        nums[pos2] = temp;
    }
}
