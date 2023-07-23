package com.dudu.leetcode.p5297;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/29/2019 11:24 AM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    private boolean reach = false;
    public boolean canReach(int[] arr, int start) {
        boolean[] flags = new boolean[arr.length];
        traversal(arr, start, flags);
        return reach;
    }
    private void traversal(int[] arr, int index, boolean[] flags) {
        if (!reach) {
            if (!flags[index]) {
                flags[index] = true;
                if (arr[index] == 0) {
                    reach = true;
                    return;
                }
                if (index - arr[index] >= 0 && index - arr[index] < arr.length) {
                    traversal(arr, index - arr[index], flags);
                }
                if (index + arr[index] >= 0 && index + arr[index] < arr.length) {
                    traversal(arr, index + arr[index], flags);
                }
            }
        }
    }
}
