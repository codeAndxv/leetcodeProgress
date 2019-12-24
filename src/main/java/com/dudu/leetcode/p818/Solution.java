package com.dudu.leetcode.p818;

import java.util.LinkedList;
import java.util.List;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/23/2019 4:08 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public int racecar(int target) {
        List<Integer> list = new LinkedList<>();
        int tem = 0, i = 1;
        while (tem < target) {
            i *= 2;
            tem = i - 1;
            list.add(tem);
        }
        tem = 0;
        int result = 0;
        boolean direction = true;
        while (tem != target){
            if(tem < target){
                if(direction){

                }
            }
        }
        return 0;
    }

    private int getMinTime(List<Integer> list, int distance){
        int result = 0;
        for(int index = list.size()-1; index>=0; index--){
            if(distance > list.get(index)){

            }
        }
        return 0;
    }



}
