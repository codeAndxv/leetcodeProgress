package com.dudu.testsrc.quicksort;

import com.alibaba.fastjson.JSONObject;

import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 9/12/2019 4:43 PM
 * @email zhao.lu@parcelx.io
 */
public class QuickSort {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Integer> source = Arrays.asList(6, 4, 4, 2, 6, 3, 9, 7, 8, 3, 2, 8, 5);
        System.out.println("sources list" + JSONObject.toJSON(source));
        dichotomy(0, source.size(), source);
        System.out.println("result list" + JSONObject.toJSON(source));
    }

    private static void dichotomy(int start, int end, List<Integer> source) {
        int hole = start, flagValue = source.get(hole), left = start, right = end, length = end;
        while (left < right) {
            while (right > 0 && right > left) {
                right--;
                if (source.get(right) < flagValue) {
                    source.set(hole, source.get(right));
                    hole = right;
                    break;
                }
            }
            while (left < length && left < right) {
                left++;
                if (source.get(left) >= flagValue) {
                    source.set(hole, source.get(left));
                    hole = left;
                    break;
                }
            }
            if (left == right) {
                source.set(hole, flagValue);
            }
        }
        if (start < left - 1) {
            dichotomy(start, left, source);
        }
        if (left + 1 < end - 1) {
            dichotomy(left + 1, end, source);
        }
    }

}
