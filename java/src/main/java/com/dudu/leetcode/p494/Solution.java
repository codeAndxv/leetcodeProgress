package com.dudu.leetcode.p494;

import java.util.concurrent.atomic.AtomicInteger;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/10/2019 3:31 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int findTargetSumWays(int[] nums, int S) {
        BinaryTree tree = new BinaryTree();
        AtomicInteger result = new AtomicInteger(0);
        buildTree(tree, nums, S, 0, result);
        return result.get();
    }

    private void buildTree(BinaryTree tree,int[] nums, int S, int index, AtomicInteger result){
        if(index < nums.length){
            tree.addLeft(-nums[index]);
            buildTree(tree.left, nums, S, index + 1, result);
            tree.addRight(nums[index]);
            buildTree(tree.right, nums, S, index + 1, result);
            if(index == nums.length-1){
                if(tree.left.totalValue == S){
                    result.incrementAndGet();
                }
                if(tree.right.totalValue == S){
                    result.incrementAndGet();
                }
            }
        }
    }
    class BinaryTree{
        int value;
        BinaryTree left;
        BinaryTree right;
        int totalValue;
        public void addLeft(int value){
            BinaryTree left = new BinaryTree();
            left.value = value;
            left.totalValue = this.totalValue + value;
            this.left = left;
        }

        public void addRight(int value){
            BinaryTree right = new BinaryTree();
            right.value = value;
            right.totalValue = this.totalValue + value;
            this.right = right;
        }
    }
}
