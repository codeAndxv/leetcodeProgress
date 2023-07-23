package com.dudu.leetcode.p5296;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/29/2019 10:41 AM
 * @Version 1.0
 * @Description
 **/
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }

    static void construct(TreeNode root, int[] nums, int index){
        if(index*2+1<nums.length){
            root.left = new TreeNode(nums[index*2+1]);
            construct(root.left, nums, index*2+1);
        }
        if(index*2+2<nums.length){
            root.right = new TreeNode(nums[index*2+2]);
            construct(root.right, nums, index*2+2);
        }
    }

}
