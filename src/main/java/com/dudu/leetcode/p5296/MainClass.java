package com.dudu.leetcode.p5296;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/29/2019 10:58 AM
 * @Version 1.0
 * @Description
 **/
public class MainClass {
    public static void main(String[] args) {
        TreeNode root1 = new TreeNode(0);
        TreeNode root2 = new TreeNode(5);
        TreeNode.construct(root1, new int[]{0,-10, 10}, 0);
        TreeNode.construct(root2, new int[]{5,1,7,0,2}, 0);
        Solution solution = new Solution();
        solution.getAllElements(root1, root2);
    }
}
