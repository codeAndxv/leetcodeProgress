package com.dudu.leetcode.p121;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/23/2019 10:36 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int maxProfit(int[] prices) {
        int result = 0;
        if(prices.length >1){
            int minPrice = prices[0];
            for(int i = 1; i<prices.length;i++){
                result = Math.max(result, prices[i]- minPrice);
                minPrice = Math.min(minPrice, prices[i]);
            }
        }
        return result;
    }
}
