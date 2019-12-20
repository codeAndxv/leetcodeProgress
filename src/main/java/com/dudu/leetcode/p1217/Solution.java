package com.dudu.leetcode.p1217;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 10/30/2019 4:26 PM
 * @email zhao.lu@parcelx.io
 */
class Solution {
    public int minCostToMoveChips(int[] chips) {
        int oddNum = 0, evenNum = 0;
        for(int location : chips){
            if(location%2==0){
                evenNum++;
            }else {
                oddNum++;
            }
        }
        return Math.min(oddNum, evenNum);
    }
}