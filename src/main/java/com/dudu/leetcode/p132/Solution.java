package com.dudu.leetcode.p132;

import java.util.LinkedList;
import java.util.List;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/31/2019 1:08 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public boolean find132pattern(int[] nums) {
        List<int[]> collection = new LinkedList<>();
        Integer tem = null, leftNum = null, rightNum = null;
        for (int num : nums) {
            if (tem == null) {
                tem = num;
                leftNum = tem;
            } else {
                if (tem <= num) {
                    tem = num;
                    rightNum = rightNum == null ? num : Math.max(rightNum, num);
                } else {
                    if (rightNum != null) {
                        collection.add(new int[]{leftNum, rightNum});
                    }
                    tem = num;
                    leftNum = tem;
                }
            }
            for(int[] range:collection){
                if(range[0]<num && range[1]> num){
                    return true;
                }
            }
        }
        return false;
    }
}
