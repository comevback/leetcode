package Leetcode;

public class Three_Sum_Closest_16 {

    public static void main(String[] args) {
        int[] nums = {-1, 2, 1, -4};
        Three_Sum_Closest_16 app = new Three_Sum_Closest_16();
        app.mergeSort(nums, 0, nums.length);
        for (int num : nums) {
            System.out.println(num);
        }
    }

    public int threeSumClosest(int[] nums, int target) {
        Three_Sum_Closest_16 app = new Three_Sum_Closest_16();
        app.mergeSort(nums, 0, nums.length);
        int minDiff = Integer.MAX_VALUE;
        int res = Integer.MAX_VALUE;

        for (int i = 0; i < nums.length; i++) {
            int head = 0;
            int tail = nums.length - 1;

            while (head < tail) {
                if (head == i) {
                    head += 1;
                } else if (tail == i) {
                    tail -= 1;
                } else {
                    int sum = nums[i] + nums[head] + nums[tail];
                    int diff = Math.abs(sum - target);
                    if (diff < minDiff) {
                        minDiff = diff;
                        res = sum;
                    }
                    if (sum == target) {
                        return sum;
                    } else if (sum < target) {
                        head += 1;
                    } else {
                        tail -= 1;
                    }
                }
            }
        }

        return res;
    }


    public void mergeSort(int[] nums, int left, int right) {
        if (left >= right - 1) {
            return;
        }

        int mid = left + (right - left) / 2;
        mergeSort(nums, left, mid);
        mergeSort(nums, mid, right);
        merge(nums, left, mid, right);
    }

    public void merge(int[] nums, int left, int mid, int right) {
        int[] sorted = new int[right - left];
        int index = 0;
        int p1 = left;
        int p2 = mid;

        while (p1 < mid && p2 < right){
            if (nums[p1] <= nums[p2]){
                sorted[index] = nums[p1];
                p1 += 1;
            }else{
                sorted[index] = nums[p2];
                p2 += 1;
            }

            index += 1;
        }

        while (p1 < mid){
            sorted[index] = nums[p1];
            p1 += 1;
            index += 1;
        }

        while (p2 < right){
            sorted[index] = nums[p2];
            p2 += 1;
            index += 1;
        }

        System.arraycopy(sorted, 0, nums, left, sorted.length);
    }
}
