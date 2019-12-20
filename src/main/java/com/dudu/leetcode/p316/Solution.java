package com.dudu.leetcode.p316;

import java.util.LinkedList;
import java.util.List;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/18/2019 2:39 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    private final int SCALE = 26;
    // slide window
    // failure try, cannot resolve the problem, for "bcabc", expect is "abc"
    public String removeDuplicateLetters(String s) {
        List<Character> slideWindow = new LinkedList<>();
        for (Character c : s.toCharArray()) {
            if (!slideWindow.contains(c)) {
                slideWindow.add(c);
            } else if (needReplace(slideWindow, c)) {
                slideWindow.remove(c);
                slideWindow.add(c);
            }
        }
        StringBuilder stringBuilder = new StringBuilder();
        for(Character c: slideWindow){
            stringBuilder.append(c);
        }
        return stringBuilder.toString();
    }

    private boolean needReplace(List<Character> slideWindow, Character c) {
        List<Character> tem =copy(slideWindow);
        tem.remove(c);
        tem.add(c);
        return calculate(slideWindow) > calculate(tem);
    }

    private List<Character> copy(List<Character> sliderWindow){
        List<Character> result= new LinkedList<>();
        for(Character character:sliderWindow){
            result.add(character);
        }
        return result;
    }

    private long calculate(List<Character> slideWindow) {
        long sum = 0;
        for (int i = 0; i < slideWindow.size(); i++) {
            sum = sum + (slideWindow.get(i) - 'a') * pow(SCALE, slideWindow.size() - 1 - i);
        }
        return sum;
    }

    private long pow(int i1, int i2) {
        long result = 1;
        for (int i = 0; i < i2; i++) {
            result = result * i1;
        }
        return result;
    }
}
