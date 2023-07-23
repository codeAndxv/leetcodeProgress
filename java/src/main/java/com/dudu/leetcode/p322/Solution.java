package com.dudu.leetcode.p322;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/26/2019 10:44 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int coinChange(int[] coins, int amount) {
        Integer[] nums = new Integer[amount + 1];
        nums[0] = 0;
        for (int index = 1; index <= amount; index++) {
            int result = -1;
            for (int coin : coins) {
                if (index >= coin) {
                    int tem = nums[index - coin] == -1 ? -1 : nums[index - coin] + 1;
                    if(tem != -1){
                        if(result == -1){
                            result = tem;
                        }else {
                            result = Math.min(result, tem);
                        }
                    }
                }
            }
            nums[index] = result;
        }
        return nums[amount];
    }
}
