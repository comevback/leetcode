package Leetcode;

public class Next_Permutation_31 {
    public static void main(String[] args) {
        int[] arr = {1,3,2};
        Next_Permutation_31 app = new Next_Permutation_31();
        app.nextPermutation(arr);
        for (int num : arr){
            System.out.println(num);
        }
    }

    public void nextPermutation(int[] nums) {
        if (nums.length < 2){
            return;
        }

        int left = nums.length - 2;
        int right = nums.length - 1;

        if (nums[left] < nums[right]){
            swap(nums, left, right);
            return;
        }

        while (left > 0 && nums[left-1] >= nums[left]){
            left -= 1;
        }

        if (left == 0){
            reverse(nums, 0, nums.length-1);
            return;
        }

        while (nums[left - 1] >= nums[right]){
            right -= 1;
        }

        swap(nums, left-1, right);
        reverse(nums, left, nums.length-1);
    }

    public void reverse(int[] nums, int head, int tail){
        if (head<0 || tail<0 || head>=nums.length || tail>=nums.length){
            throw new RuntimeException("out of range");
        }
        while (head < tail){
            swap(nums,head,tail);
            head += 1;
            tail -= 1;
        }
    }

    public void swap(int[] nums, int left, int right){
        if (left<0 || right<0 || left>=nums.length || right>nums.length){
            throw new RuntimeException("out of range");
        }

        int temp = nums[left];
        nums[left] = nums[right];
        nums[right] = temp;
    }
}
