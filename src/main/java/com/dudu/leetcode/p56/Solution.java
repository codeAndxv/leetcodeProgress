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
            boolean merge = false;
            for (int index = 0; index < list.size(); index++) {
                if (list.get(index)[0] >= interval[0]) {
                    list.add(index, interval);
                    merge = true;
                }
            }
            if (!merge) {
                list.add(interval);
            }
        }
        ListIterator<int[]> iterator = list.listIterator();
        Integer right = null;
        while (iterator.hasNext()){
            int[] interval = iterator.next();
            if(right == null){
                right = interval[1];
            }else {
                if(interval[0]<=iterator.previous()[1]){
                    iterator.next();

                }
            }
        }
    }
}
