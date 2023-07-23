package com.dudu.leetcode.p98;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/12/2019 10:32 AM
 * @email zhao.lu@parcelx.io
 */
// Definition for a binary tree node.
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }

    public static TreeNode constructTree(Integer[] input){
        TreeNode treeNode = null;
        if(input.length != 0){
            treeNode = new TreeNode(input[0]);
            if(1 < input.length && input[1] != null){
                treeNode.left = new TreeNode(input[1]);
                assignValue(treeNode.left, input, 1);
            }
            if(2 < input.length && input[2] != null){
                treeNode.right = new TreeNode(input[2]);
                assignValue(treeNode.right, input, 2);
            }
        }
        return treeNode;
    }

    private static void assignValue(TreeNode treeNode, Integer[] input, int index){
        if(index *2 + 1 < input.length && input[index *2 +1] != null){
            treeNode.left = new TreeNode(input[index *2 + 1]);
            assignValue(treeNode.left, input, index *2 + 1);
        }
        if(index *2 + 2 < input.length && input[index *2 + 2] != null){
            treeNode.right = new TreeNode(input[index *2 + 2]);
            assignValue(treeNode.right, input, index *2 + 2);
        }
    }
}
