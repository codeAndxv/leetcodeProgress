package com.dudu.leetcode.p3;

import java.util.HashSet;

/**
 * @author code
 * @create 2018-09-27 14:43
 */


/**
 * Given a string, find the length of the longest substring without repeating characters.
 *
 * Example 1:
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 * Example 2:
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 * Example 3:
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 *              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int len = s.length();
        int[] nums = new int[len];

        nums[0] = 1;

        if(len <= 0){
            return 0;
        }

        for(int i=1; i<len; i++){
            if(s.charAt(i) == s.charAt(i-1)){
                nums[i] = nums[i-1];
            }else{
                if(isAddOne(s, i-nums[i-1],i)){
                    nums[i] = nums[i-1]+1;
                }else {
                    nums[i] = nums[i-1];
                }
            }
        }

        return nums[len-1];
    }

    //判断是否有重复的字符，没有重复字符可以添加一。如果没有返回true
    public boolean isAddOne(String str, int from, int to){

        HashSet set = new HashSet();

        for(int i=from; i<=to; i++){
            set.add(str.charAt(i));
        }
        if(set.size()==(to - from + 1)){
            return true;
        }
        return false;
    }
}