package com.dudu.leetcode.p438;


import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/10/2019 11:02 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        int[] recordNum = new int[26];
        for(char c : p.toCharArray()){
            recordNum[c-'a']++;
        }
        List<Integer> result = new LinkedList<>();
        int l = 0, r = 0, len = p.length();
        while(r < s.length()){
            if (recordNum[s.charAt(r++)-'a']-- >0){
                len--;
            }
            while (len == 0){
                if(r-l == p.length()){
                    result.add(l);
                }
                if(recordNum[s.charAt(l++)-'a'] ++ == 0){
                    len ++;
                }
            }

        }
        return result;
    }
}
