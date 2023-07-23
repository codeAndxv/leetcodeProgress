package com.dudu.leetcode.p123;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/23/2019 11:55 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int maxProfit(int[] prices) {
        if(prices==null || prices.length<=1)
            return 0;
        if(prices.length==2)
            return Math.max(0, prices[1]-prices[0]);
        return maxProfitOptimized(prices, 3);
    }

    private int maxProfitOptimized(int[] prices, int k) {
        int[][] T = new int[k+1][prices.length];
        for(int i=1;i<T.length;i++) {
            int maxDiff = T[i-1][0] - prices[0];
            for(int j=1;j<T[0].length;j++) {
                T[i][j] = Math.max(T[i][j-1], prices[j]+maxDiff);
                maxDiff = Math.max(maxDiff, T[i-1][j] - prices[j]);
            }
        }
        return T[k][prices.length-1];
    }
}
