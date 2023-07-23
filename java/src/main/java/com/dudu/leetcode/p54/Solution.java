package com.dudu.leetcode.p54;

import java.util.LinkedList;
import java.util.List;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/28/2019 11:50 AM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new LinkedList<>();
        if (matrix.length >= 1) {
            int iTop = 0, iBottom = matrix.length - 1, iLeft = 0, iRight = matrix[0].length - 1;
            while (iTop <= iBottom && iLeft <= iRight) {
                for (int i = iLeft; i <= iRight; i++) {
                    result.add(matrix[iTop][i]);
                }
                iTop++;
                if (iTop > iBottom) {
                    break;
                }
                for (int i = iTop; i <= iBottom; i++) {
                    result.add(matrix[i][iRight]);
                }
                iRight--;
                if (iRight < iLeft) {
                    break;
                }
                for (int i = iRight; i >= iLeft; i--) {
                    result.add(matrix[iBottom][i]);
                }
                iBottom--;
                if (iBottom < iTop) {
                    break;
                }
                for (int i = iBottom; i >= iTop; i--) {
                    result.add(matrix[i][iLeft]);
                }
                iLeft++;
            }
        }
        return result;
    }
}