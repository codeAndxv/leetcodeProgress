package com.dudu.leetcode.p179;

import java.util.Arrays;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/31/2019 2:25 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public String largestNumber(int[] nums) {
        int zeroCount = nums.length;
        nums = Arrays.stream(nums).filter(value -> value != 0).toArray();
        zeroCount -= nums.length;
        sort(nums, 0, nums.length - 1);
        StringBuilder stringBuilder = new StringBuilder();
        for (int num : nums) {
            stringBuilder.append(num);
        }
        if (stringBuilder.length() != 0) {
            for (int i = 0; i < zeroCount; i++) {
                stringBuilder.append(0);
            }
        } else {
            if (zeroCount != 0) {
                stringBuilder.append(0);
            }
        }
        return stringBuilder.toString();
    }

    private void sort(int[] nums, int left, int right) {
        if (left < right) {
            int center = (left + right) / 2;
            sort(nums, left, center);
            sort(nums, center + 1, right);
            int[] tem = Arrays.copyOfRange(nums, left, right + 1);
            int pointer1 = 0, pointer2 = center + 1 - left, pointer = left;
            while (pointer1 <= center - left || pointer2 <= right - left) {
                if (pointer1 <= center - left && pointer2 <= right - left) {
                    if (compare(tem[pointer1], tem[pointer2])) {
                        nums[pointer] = tem[pointer1];
                        pointer++;
                        pointer1++;
                    } else {
                        nums[pointer] = tem[pointer2];
                        pointer++;
                        pointer2++;
                    }
                } else if (pointer1 <= center - left) {
                    nums[pointer] = tem[pointer1];
                    pointer++;
                    pointer1++;
                } else {
                    nums[pointer] = tem[pointer2];
                    pointer++;
                    pointer2++;
                }
            }
        }
    }

    /**
     * if num1 >= num2, return true
     *
     * @param num1
     * @param num2
     * @return
     */
    public boolean compare(int num1, int num2) {
        return num1 * (Math.pow(10, String.valueOf(num2).length()) - 1) >= num2 * (Math.pow(10, String.valueOf(num1).length()) - 1);
    }
}
