package com.dudu.leetcode.p647;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/28/2019 2:50 PM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int countSubstrings(String s) {
        boolean[][] paStatus = new boolean[3][s.length()];
        int result = 0;
        for (int len = 1; len <= s.length(); len++) {
            for (int index = 0; index < s.length() + 1 - len; index++) {
                if (len == 1) {
                    paStatus[(len - 1) % 3][index] = true;
                } else {
                    paStatus[(len - 1) % 3][index] = (len == 2 || paStatus[(len - 1 - 2 + 3) % 3][index + 1]) && s.charAt(index) == s.charAt(index + len - 1);
                }
                result = paStatus[(len - 1) % 3][index] ? result + 1 : result;
            }
        }
        return result;
    }
}
