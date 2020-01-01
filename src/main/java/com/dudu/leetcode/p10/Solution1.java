package com.dudu.leetcode.p10;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/12/2019 11:49 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution1 {
    // not greed
    public boolean isMatch(String s, String p) {
        p = simplyPattern(p);
        int indexS = 0, indexP = 0;
        boolean result = false;
        while (indexS < s.length() && indexP < p.length()) {
            if (s.charAt(indexS) == p.charAt(indexP) || p.charAt(indexP) == '.') {
                indexP++;
                indexS++;
            } else if (p.charAt(indexP) == '*') {
                if (p.charAt(indexP - 1) == s.charAt(indexS) || p.charAt(indexP - 1) == '.') {

                } else if (indexP + 1 < p.length() && (p.charAt(indexP + 1) == s.charAt(indexS) || p.charAt(indexP + 1) == '.')) {
                    indexP = indexP + 2;
                } else {
                    break;
                }
                indexS++;
            } else if (indexP + 2 < p.length()) {
                indexP = indexP + 2;
            } else {
                break;
            }
        }
        return indexS == s.length() && (indexP == p.length() || (indexP == p.length() - 1 && p.charAt(indexP) == '*'));
    }

    private String simplyPattern(String p) {
        for (int index = 0; index < p.length(); index++) {
            if (p.charAt(index) == '*') {
                int left = index - 2;
                for (; left >= 0;) {
                    if (p.charAt(left) != p.charAt(index - 1)) {
                        left++;
                        break;
                    }
                    left--;
                }
                left = Math.max(0, left);
                int right = index + 1;
                for (; right < p.length();) {
                    if (p.charAt(right) != p.charAt(index - 1)) {
                        right--;
                        break;
                    }
                    right++;
                }
                p = p.substring(0, left + 1) + p.substring(right);
            }
        }
        return p;
    }
}
