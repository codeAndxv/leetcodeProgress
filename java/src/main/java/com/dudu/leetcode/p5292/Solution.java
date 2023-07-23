package com.dudu.leetcode.p5292;

import java.util.*;
import java.util.function.BiFunction;
import java.util.stream.Collectors;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/22/2019 11:29 AM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public boolean isPossibleDivide(int[] nums, int k) {
        TreeMap<Integer, Integer> sortedMap = new TreeMap<>();
        for(int num : nums){
            sortedMap.putIfAbsent(num, 0);
            sortedMap.computeIfPresent(num, (integer, integer2) -> integer2 + 1);
        }
        List<Integer> keys = new ArrayList<>(sortedMap.keySet());
        for(int index = 0; index <keys.size();){
            if(sortedMap.get(keys.get(index)) == 0){
                index++;
            }else {
                for(int tem = 0; tem<k; tem++){
                    if(index+tem >= keys.size()){
                        return false;
                    }
                    if(sortedMap.get(keys.get(index+tem)) == null || sortedMap.get(keys.get(index+tem))<1){
                        return false;
                    }else {
                        sortedMap.replace(keys.get(index+tem), sortedMap.get(keys.get(index+tem)) - 1);
                    }
                }
            }
        }
        return true;
    }
}
