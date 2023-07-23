package com.dudu.leetcode.p885;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/28/2019 12:35 PM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public int[][] spiralMatrixIII(int R, int C, int r0, int c0) {
        List<int[]> list = new LinkedList<>();
        int iTop = r0, iBottom = r0, iLeft = c0, iRight = c0;
        list.add(new int[]{r0,c0});
        while (!(iTop < 0 && iBottom > R && iLeft < 0 && iRight > C)) {
            iRight++;
            for (int i = iLeft+1; i <= iRight; i++) {
                if (isExist(iTop, i, R, C)) {
                    list.add(new int[]{iTop, i});
                }
            }
            iBottom++;
            for (int i = iTop + 1; i <= iBottom; i++) {
                if (isExist(i, iRight, R, C)) {
                    list.add(new int[]{i, iRight});
                }
            }
            iLeft--;
            for (int i = iRight - 1; i >= iLeft; i--) {
                if (isExist(iBottom, i, R, C)) {
                    list.add(new int[]{iBottom, i});
                }
            }
            iTop--;
            for (int i = iBottom - 1; i >= iTop; i--) {
                if (isExist(i, iLeft, R, C)) {
                    list.add(new int[]{i, iLeft});
                }
            }
        }
        int[][] result = new int[list.size()][2];
        for(int index = 0; index<list.size();index++){
            result[index] = list.get(index);
        }
        return result;
    }

    private boolean isExist(int dimen1, int dimen2, int R, int C) {
        return dimen1 >= 0 && dimen1 < R && dimen2 >= 0 && dimen2 < C;
    }
}
