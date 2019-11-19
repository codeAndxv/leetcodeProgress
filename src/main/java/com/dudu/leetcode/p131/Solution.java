package com.dudu.leetcode.p131;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 7/26/2019 4:52 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    private List<List<List>> store;

    public List<List<String>> partition(String s) {
        if (s == null || s.length() == 0) {
            return null;
        }
        int len = s.length();
        if (store == null) {
            store = new ArrayList<>();
            for (int index = 0; index <= len; index++) {
                store.add(new ArrayList<>());
                for (int innerindex = 0; innerindex <= len; innerindex++) {
                    store.get(index).add(new ArrayList<>());
                }
            }
        }
        List<List<String>> result = new LinkedList<>();
        if (s.length() == 1) {
            result.add(Stream.of(s).collect(Collectors.toList()));
            return result;
        } else {
            for (int index = 1; index <= s.length(); index++) {
                String leftStr = s.substring(0, index);
                String rightStr = s.substring(index);
                if (isPalindrome(leftStr)) {
                    if(store.get(index).get(s.length()) == null || store.get(index).get(s.length()).size() == 0){
                        store.get(index).set(s.length(), partition(rightStr));
                    }
                    List<List<String>> rightResult = store.get(index).get(s.length());
                    List<String> allTemResult;
                    List<String> perLeftResult = Stream.of(leftStr).collect(Collectors.toList());
                    if (rightResult != null) {
                        for (List<String> perRightResult : rightResult) {
                            allTemResult = new LinkedList<>();
                            allTemResult.addAll(perLeftResult);
                            allTemResult.addAll(perRightResult);
                            result.add(allTemResult);
                        }
                    } else {
                        allTemResult = new LinkedList<>(perLeftResult);
                        result.add(allTemResult);
                    }
                }
            }
            return result;
        }
    }

    private boolean isPalindrome(String s) {
        int len = s.length() - 1;
        for (int index = 0; index <= len / 2; index++) {
            if (s.charAt(index) != s.charAt(len - index)) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        solution.partition("aab");
    }


}
