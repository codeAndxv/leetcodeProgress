package com.dudu.leetcode.p779;

import java.util.LinkedList;
import java.util.List;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/19/2019 6:16 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int kthGrammar(int N, int K) {
        List<Integer> path = new LinkedList<>();
        while (K != 1){
            path.add(K);
            K = K / 2 + K % 2;
        }
        int result = 0, start = 1;
        for (int index = path.size() - 1; index >= 0; index--){
            if(path.get(index) == 2 * start){
                if(result == 0){
                    result = 1;
                }else {
                    result = 0;
                }
                start = 2 * start;
            }else {
                start = 2 * start - 1;
            }
        }
        return result;
    }
}
