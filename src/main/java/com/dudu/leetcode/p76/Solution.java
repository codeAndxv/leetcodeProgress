package com.dudu.leetcode.p76;

import java.util.*;
import java.util.function.BiFunction;
import java.util.function.Consumer;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/23/2019 5:39 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public String minWindow(String s, String t) {
        Map<Character, Integer> bitMap = new HashMap<>();
        Map<Character, Integer> targetMap = new HashMap<>();
        t.codePoints().mapToObj(value -> (char) value).forEach(character -> bitMap.put(character, 0));
        t.codePoints().mapToObj(value -> (char) value).forEach(character -> {
            targetMap.putIfAbsent(character, 0);
            targetMap.computeIfPresent(character, (character1, integer) -> integer+1);
        });
        Set<Character> recordStatus = new HashSet<>();
        int i = 0, j = 0;
        int minLen = Integer.MAX_VALUE, recordI = 0, recordJ = 0;
        for (; j < s.length(); j++) {
            if (bitMap.containsKey(s.charAt(j))) {
                bitMap.computeIfPresent(s.charAt(j), (character, integer) -> integer + 1);
                if(bitMap.get(s.charAt(j)).equals(targetMap.get(s.charAt(j)))){
                    recordStatus.add(s.charAt(j));
                }
            }
            if (recordStatus.size() == bitMap.keySet().size()) {
                for (; i < j; i++) {
                    if (!bitMap.containsKey(s.charAt(i)) || bitMap.get(s.charAt(i)) > targetMap.get(s.charAt(i))) {
                        bitMap.computeIfPresent(s.charAt(i), (character, integer) -> integer - 1);
                    }else {
                        break;
                    }
                }
                if (minLen > j - i) {
                    minLen = j - i;
                    recordI = i;
                    recordJ = j;
                }
            }
        }
        if(recordStatus.size() == bitMap.keySet().size()){
            return s.substring(recordI, recordJ + 1);
        }else {
            return "";
        }
    }
}
