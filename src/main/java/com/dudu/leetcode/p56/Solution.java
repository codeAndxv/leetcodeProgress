package com.dudu.leetcode.p56;

import java.util.LinkedList;
import java.util.List;
import java.util.ListIterator;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/31/2019 3:37 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int[][] merge(int[][] intervals) {
        List<int[]> list = new LinkedList<>();
        for (int[] interval : intervals) {
            List<Integer> involve = new LinkedList<>();
            int index = 0;
            while (index < list.size() && interval[1] >= list.get(index)[0]) {
                if (list.get(index)[1] >= interval[0]) {
                    involve.add(index);
                    interval[0] = Math.min(list.get(index)[0], interval[0]);
                    interval[1] = Math.max(list.get(index)[1], interval[1]);
                }
                index++;
            }
            for (int i = involve.size() - 1; i >= 0; i--) {
                list.remove((int)involve.get(i));
            }
            list.add(involve.size() == 0 ? index : involve.get(0), interval);
        }
        int[][] result = new int[list.size()][2];
        for (int i = 0; i < list.size(); i++) {
            result[i] = list.get(i);
        }
        return result;
    }
}
