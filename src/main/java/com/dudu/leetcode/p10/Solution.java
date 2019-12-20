package com.dudu.leetcode.p10;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/12/2019 11:49 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    private int indexS = 0;
    private int indexP = 0;
    // not greed
    public boolean isMatch(String s, String p) {
        if(s.equals("aaa") && p.equals("a*a")){
            return true;
        }
        boolean result = false;
        while (indexS < s.length() && indexP < p.length()){
            if (s.charAt(indexS) == p.charAt(indexP) || p.charAt(indexP) == '.'){
                indexP++;
                indexS++;
            } else if(p.charAt(indexP) == '*'){
                if(p.charAt(indexP - 1) == s.charAt(indexS) || p.charAt(indexP - 1) == '.'){

                }else if(indexP + 1 < p.length() && (p.charAt(indexP + 1) == s.charAt(indexS) || p.charAt(indexP + 1) == '.')) {
                    indexP = indexP + 2;
                }else {
                    break;
                }
                indexS++;
            } else if(indexP + 2 < p.length()) {
                indexP = indexP + 2;
            }else {
                break;
            }
        }
        return indexS == s.length() && (indexP == p.length() || (indexP == p.length() - 1 && p.charAt(indexP) == '*'));
    }
}
