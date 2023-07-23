package com.dudu.leetcode.p518;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/26/2019 11:25 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int change(int amount, int[] coins) {
        Integer[] nums = new Integer[amount + 1];
        nums[0] = 1;
        for (int index = 1; index <= amount; index++) {
            int result = 0;
            for (int coin : coins) {
                if (index >= coin) {
                    if(nums[index-coin] != 0){
                        result+=nums[index-coin];
                    }
                }
            }
            nums[index] = result;
        }
        return nums[amount];
    }
}
