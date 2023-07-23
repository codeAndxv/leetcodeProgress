package com.dudu.leetcode.p887;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/25/2019 10:37 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int superEggDrop(int K, int N) {
        int[][] recordTable = new int[K + 1][N + 1];
        for (int index = 1; index <= N; index++) {
            recordTable[1][index] = index;
        }
        for (int k = 2; k <= K; k++) {
            for (int index = 1; index <= N; index++) {
                int leftLen = 0;
                int rightLen = 0;
                if(index % 2 ==0){
                    leftLen = index / 2;
                    rightLen = index + 1 - leftLen;
                }else {
                    leftLen = (index + 1) / 2;
                    rightLen = index + 1 - leftLen;
                }
                recordTable[k][index] = Math.min(recordTable[k][index - 1] + 1, Math.max(leftLen - 1 < 0 ? 0 : recordTable[k - 1][leftLen - 1], rightLen - 1 < 0 ? 0 : recordTable[k][rightLen - 1]) + 1);
            }
        }
        return recordTable[K][N];
    }
}
