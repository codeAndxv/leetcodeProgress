package com.dudu.leetcode.p84;

import java.util.Arrays;
import java.util.OptionalInt;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 11/20/2019 2:23 PM
 * @email zhao.lu@parcelx.io
 */
class Solution {
    public int largestRectangleArea(int[] heights) {
        OptionalInt largestHeightOptional = Arrays.stream(heights).max();
        if(!largestHeightOptional.isPresent()){
            return 0;
        }
        int largestHeight = largestHeightOptional.getAsInt();
        int largestArea = 0;
        int[] heightCopy = Arrays.stream(heights).distinct().toArray();
        Arrays.sort(heightCopy);
        for (int temHeight : heightCopy) {
            largestArea = Math.max(largestArea, largestWidth(heights, temHeight) * temHeight);
        }
        return largestArea;
    }



    private int largestWidth(int[] heights, int targetHeight) {
        int largestWidth = 0, temWidth = 0;
        for (int height : heights) {
            if (height >= targetHeight) {
                temWidth++;
                largestWidth = Math.max(largestWidth, temWidth);
            } else {
                temWidth = 0;
            }
        }
        return largestWidth;
    }
}