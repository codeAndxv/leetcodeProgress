package com.dudu.leetcode.p5291;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/22/2019 11:09 AM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int findNumbers(int[] nums) {
        int result = 0;
        for(int num:nums){
            int digit = 1;
            while (num > 9){
                num/=10;
                digit++;
            }
            if(digit%2==0){
                result++;
            }
        }
        return result;
    }
}
