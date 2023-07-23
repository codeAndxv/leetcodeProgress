package com.dudu.leetcode.p91;

import java.util.HashMap;
import java.util.Map;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/12/2019 4:27 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    private Map<String, Integer> store = new HashMap<>();
    public int numDecodings(String s) {
        // init
        store.putIfAbsent("0", 0);
        for(int i=1; i<10; i++){
            store.putIfAbsent(String.valueOf(i), 1);
        }
        return numDecoding(s);
    }

    private int numDecoding(String s) {
        if(store.get(s) == null){
            // 12
            int num = 0;
            if (s.charAt(0) == '0') {
                return 0;
            }
            if(s.length() == 2){
                if(Integer.parseInt(s) <= 26 && Integer.parseInt(s) >=1){
                    num++;
                }
                String left = s.substring(0, 1);
                String right = s.substring(1);
                num = num + numDecoding(left) *  numDecoding(right);
            }else {
                String left = s.substring(0, 1);
                String right = s.substring(1);
                num = num + numDecoding(left) *  numDecoding(right);
                left = s.substring(0, 2);
                int leftInt = Integer.parseInt(left);
                if(leftInt <= 26){
                    right = s.substring(2);
                    num = num + numDecoding(right);
                }
            }
            store.putIfAbsent(s, num);
        }
        return store.get(s);
    }

}
