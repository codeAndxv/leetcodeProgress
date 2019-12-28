package com.dudu.leetcode.p59;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/28/2019 12:22 PM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int[][] generateMatrix(int n) {
        int[][] matrix = new int[n][n];
        if (n >= 1) {
            int value = 1, iTop = 0, iBottom = matrix.length - 1, iLeft = 0, iRight = matrix[0].length - 1;
            while (iTop <= iBottom && iLeft <= iRight) {
                for (int i = iLeft; i <= iRight; i++) {
                    matrix[iTop][i] = value;
                    value++;
                }
                iTop++;
                if (iTop > iBottom) {
                    break;
                }
                for (int i = iTop; i <= iBottom; i++) {
                    matrix[i][iRight] = value;
                    value++;
                }
                iRight--;
                if (iRight < iLeft) {
                    break;
                }
                for (int i = iRight; i >= iLeft; i--) {
                    matrix[iBottom][i] = value;
                    value++;
                }
                iBottom--;
                if (iBottom < iTop) {
                    break;
                }
                for (int i = iBottom; i >= iTop; i--) {
                    matrix[i][iLeft] = value;
                    value++;
                }
                iLeft++;
            }
        }
        return matrix;
    }
}
