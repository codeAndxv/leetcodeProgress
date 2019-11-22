package com.dudu.leetcode.p42;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 11/22/2019 2:12 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    /**
     * simulate the actual rain process
     *
     * @param height
     * @return
     */
    public int trap(int[] height) {
        int rainedArea = 0;
        Wall[] walls = new Wall[height.length];
        // init wall array
        for(int i = 0;i<height.length;i++){
            walls[i] = new Wall();
        }
        for(int i = 0; i<height.length;i++){

        }
        for (int i = 0; i < height.length; i++) {
            rainedArea += Math.max(findMinOfMaxHeight(i, height) - height[i], 0);
        }
        return rainedArea;
    }

    private class Wall{
        int leftHeight;
        int rightHeight;

        public Wall() {
        }

        public Wall(int leftHeight, int rightHeight) {
            this.leftHeight = leftHeight;
            this.rightHeight = rightHeight;
        }

        public int getLeftHeight() {
            return leftHeight;
        }

        public Wall setLeftHeight(int leftHeight) {
            this.leftHeight = leftHeight;
            return this;
        }

        public int getRightHeight() {
            return rightHeight;
        }

        public Wall setRightHeight(int rightHeight) {
            this.rightHeight = rightHeight;
            return this;
        }
    }

    private int findMinOfMaxHeight(int currentIndex, int[] height) {
        int leftHeight = 0, rightHeight = 0;
        // find the first height greater than currentHeight
        for (int i = currentIndex - 1; i >= 0; i--) {
            leftHeight = Math.max(leftHeight, height[i]);
        }
        for (int i = currentIndex + 1; i < height.length; i++) {
            rightHeight = Math.max(rightHeight, height[i]);
        }
        return Math.min(leftHeight, rightHeight);
    }
}