package com.dudu.leetcode.p10;

import java.util.HashSet;
import java.util.Set;
import java.util.Stack;

public class Solution {
    public boolean isMatch(String s, String p) {
        //思路是分别为两个字符串创建两个栈
        Stack<Character> sStack = new Stack<>();
        Stack<Character> pStack = new Stack<>();

        Set<Character> sSet = new HashSet<Character>() {{
            for (char c = 'a'; c <= 'z'; c++) {
                add(c);
            }
        }};
        Set<Character> pSet = new HashSet<Character>(sSet);
        pSet.add('.');
        pSet.add('*');

        //然后把字符串每个字符都放到栈里
        for (int i = 0; i < s.length(); i++) {
            if (sSet.contains(s.charAt(i))) {
                sStack.add(s.charAt(i));
            } else {
                return false;
            }
        }
        for (int j = 0; j < p.length(); j++) {
            if (pSet.contains(p.charAt(j))) {
                pStack.add(p.charAt(j));
            } else {
                return false;
            }
        }
        //是否是循环状态
        boolean isCycle = false;
        //ptChar 用来存放cycle的char
        Character sChar, pChar, ptChar = null;
        while (!pStack.isEmpty() && !sStack.isEmpty()) {
            sChar = sStack.pop();
            if (isCycle) {
                if (ptChar != null) {

                    while (!sStack.isEmpty() && (sStack.peek() == ptChar || ptChar == '.')) {
                        sStack.pop();
                    }

                } else {
                    pChar = pStack.pop();

                    if (pChar == '*') {
                        isCycle = true;
                    }
                    ptChar = pChar;
                }
            }
        }
        return true;
    }
}
