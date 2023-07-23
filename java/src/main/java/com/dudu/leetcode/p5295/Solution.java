package com.dudu.leetcode.p5295;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/29/2019 10:32 AM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int[] sumZero(int n) {
        int[] result = new int[n];
        int index = 0;
        for(int i = n/2; i>0; i--){
            result[index++] = i;
            result[index++] = -i;
        }
        if(n%2!=0){
            result[index]=0;
        }
        return result;
    }
}
