package com.dudu.leetcode.p98;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/12/2019 11:03 AM
 * @email zhao.lu@parcelx.io
 */
public class MainClass {
    public static void main(String[] args) {
        Integer[] input = new Integer[]{10,5,15,null,null,6,20};
        TreeNode treeNode = TreeNode.constructTree(input);
        Solution solution = new Solution();
        System.out.println(solution.isValidBST(treeNode));
    }
}
