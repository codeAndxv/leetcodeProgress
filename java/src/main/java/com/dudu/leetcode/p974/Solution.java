package com.dudu.leetcode.p974;

/**
 * @Author zhao.lu
 * @Date 12/21/2019 2:14 PM
 * @Version 1.0
 * @Email yx163luzhao@163.com
 * @Description
 **/
public class Solution {
    public int subarraysDivByK(int[] A, int K) {
        int[] recordNum = new int[K];
        int pointer = 0;
        // handle A
        for (int i = 0; i < A.length; i++) {
            A[i] = (A[i] % K + K) % K;
        }
        int result = 0;
        // move the pointer, and add num according to pointer and A[i], add result in the end
        for(int i = 0;i <A.length; i++){
            pointer = ((-A[i] + K)%K + pointer) % K;
            recordNum[(pointer + A[i]) % K] += 1;
            result += recordNum[pointer];
        }
        return result;
    }
}
