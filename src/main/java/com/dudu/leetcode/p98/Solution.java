package com.dudu.leetcode.p98;

import java.util.LinkedList;
import java.util.List;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/12/2019 10:31 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    private List<Integer> list = new LinkedList<>();
    public boolean isValidBST(TreeNode root) {
        leftTraversalTree(root);
        Integer maxNum = null;
        for(Integer num : list){
            if(maxNum == null){
                maxNum = num;
            }else if(num <= maxNum) {
                return false;
            }else {
                maxNum = num;
            }
        }
        return true;
    }
    private void leftTraversalTree(TreeNode root){
        if(root != null){
            if(root.left != null){
                leftTraversalTree(root.left);
            }
            list.add(root.val);
            if(root.right != null){
                leftTraversalTree(root.right);
            }
        }
    }
}