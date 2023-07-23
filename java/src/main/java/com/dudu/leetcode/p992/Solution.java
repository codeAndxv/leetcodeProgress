package com.dudu.leetcode.p992;

import java.util.HashMap;
import java.util.Map;
import java.util.function.BiFunction;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/25/2019 9:39 PM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int subarraysWithKDistinct(int[] A, int K) {
        Map<Integer, Integer> recordStatus = new HashMap<>();
        int left = 0, right = 0;
        recordStatus.putIfAbsent(A[0], 1);
        int result = 0;
        while (right < A.length - 1) {
            right++;
            if (!recordStatus.containsKey(A[right])) {
                while (left < right && recordStatus.keySet().size() == K) {
                    recordStatus.computeIfPresent(A[left], (integer, integer2) -> integer2 - 1);
                    left++;
                    if (recordStatus.get(A[left - 1]) == 0) {
                        recordStatus.remove(A[left - 1]);
                    }
                    if (recordStatus.keySet().size() == K) {
                        System.out.println("" + left + "," + (right-1));
                        result++;
                    } else {
                        break;
                    }
                }
                recordStatus.putIfAbsent(A[right], 1);
            } else {
                recordStatus.computeIfPresent(A[right], (integer, integer2) -> integer2 + 1);
                Map<Integer, Integer> temStatus = new HashMap<>();
                int temLeft = left;
                while (temLeft < right && recordStatus.keySet().size() == K) {
                    temLeft++;
                    if (recordStatus.get(A[temLeft]) > 1) {
                        recordStatus.compute(A[temLeft], (integer, integer2) -> integer2 - 1);
                        if (recordStatus.get(A[temLeft]) == 0) {
                            recordStatus.remove(A[temLeft]);
                        }
                        temStatus.putIfAbsent(A[temLeft], 0);
                        temStatus.compute(A[temLeft], (integer, integer2) -> integer2 + 1);
                        if (recordStatus.keySet().size() == K) {
                            System.out.println("" + temLeft + "," + (right-1));
                            result++;
                        }
                    } else {
                        break;
                    }
                }
                for (Integer key : temStatus.keySet()) {
                    recordStatus.computeIfPresent(key, (integer, integer2) -> integer2 + temStatus.get(key));
                    recordStatus.putIfAbsent(key, temStatus.get(key));
                }
            }
            if (recordStatus.keySet().size() == K) {
                System.out.println("" + left + "," + (right));
                result++;
            }

        }
        return result;
    }
}
