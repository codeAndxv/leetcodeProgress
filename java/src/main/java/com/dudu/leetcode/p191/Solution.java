package com.dudu.leetcode.p191;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/11/2019 3:33 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    // you need to treat n as an unsigned value
    public int hammingWeight(int n) {
        int num = 0;
        while (n != 0){
            num = (n & 1) == 1? num + 1 : num;
            n = n >>> 1;
        }
        return num;
    }
}
