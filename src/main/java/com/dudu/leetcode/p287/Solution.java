package com.dudu.leetcode.p287;

import java.util.HashSet;
import java.util.Set;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/21/2019 7:08 PM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int findDuplicate(int[] nums) {
        Set<Integer> seen = new HashSet<Integer>();
        for (int num : nums) {
            if (seen.contains(num)) {
                return num;
            }
            seen.add(num);
        }

        return -1;
    }
}
