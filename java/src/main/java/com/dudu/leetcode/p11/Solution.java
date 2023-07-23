package com.dudu.leetcode.p11;

public class Solution {
    public int maxArea(int[] height) {
        int size = height.length;
        //设置最大面积数组，为什么是size-1，是因为有两端才能构成一个范围，所以直接从第二个高度开始的
        //算法的思路是判断area[i] = max{area[i-1],以及由新的一端也就是i端与每一个前面的端形成一个面积，一共是i个面积}
        int[] area = new int[size-1];

        for(int i = 0; i<area.length; i++){
            int temArea = getArea(height, i);
            if(i == 0){
                area[i] = temArea;
            }else {
                area[i] = Math.max(area[i-1],temArea);
            }
        }

        return area[area.length-1];
    }
    //返回一个最大的面积，一端是i,因为前面说到的，最大面积数组的size是i-1，所以i对应到的height[]是i+1
    public int getArea(int[] height, int i){
        int temArea = 0;
        for(int j = 0; j< i+1; j++){
            int nowArea = Math.min(height[j], height[i+1])*(i+1-j);
            if(temArea < nowArea){
                temArea = nowArea;
            }
        }
        return temArea;
    }

}
