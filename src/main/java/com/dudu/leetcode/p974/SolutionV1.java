package com.dudu.leetcode.p974;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/20/2019 11:39 AM
 * @email zhao.lu@parcelx.io
 */
public class SolutionV1 {
    // v1, Memory Limit Exceeded
    public int subarraysDivByK(int[] A, int K) {
        int result = 0;
        int[][]  recordTable = new int[A.length][A.length+1];
        for(int i = 0; i<A.length; i++){
            for(int j = i+1; j<A.length+1; j++){
                recordTable[i][j] = j==i+1? A[i]:recordTable[i][j-1]+A[j-1];
                if(recordTable[i][j] % K ==0){
                    result ++;
                }
            }
        }
        return result;
    }
}
