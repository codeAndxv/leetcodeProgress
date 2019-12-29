package com.dudu.leetcode.p5296;

import java.util.LinkedList;
import java.util.List;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/29/2019 10:40 AM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public List<Integer> getAllElements(TreeNode root1, TreeNode root2) {
        List<Integer> result = new LinkedList<>();
        // just merge sort
        List<Integer> result1 = new LinkedList<>(), result2 = new LinkedList<>();
        inorderTraversal(result1,root1);
        inorderTraversal(result2, root2);
        for(int index1=0, index2=0; index1<result1.size()||index2<result2.size();){
            if(index2<result2.size()&&index1<result1.size()){
                if(result1.get(index1)<=result2.get(index2)){
                    result.add(result1.get(index1));
                    index1++;
                }else {
                    result.add(result2.get(index2));
                    index2++;
                }
            }else if(index2>= result2.size()){
                result.add(result1.get(index1));
                index1++;
            }else {
                result.add(result2.get(index2));
                index2++;
            }

        }
        return result;
    }

    private void inorderTraversal(List<Integer> result, TreeNode node){
        if(node!=null){
            if(node.left!=null){
                inorderTraversal(result, node.left);
            }
            result.add(node.val);
            if(node.right!=null){
                inorderTraversal(result, node.right);
            }
        }
    }
}
