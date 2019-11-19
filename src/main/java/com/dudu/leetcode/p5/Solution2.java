package com.dudu.leetcode.p5;

import java.util.*;


/**
 * @author zhaolu
 * @version 1.0
 * @datetime 10/31/2019 11:28 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution2 {

    private Set<PalindromeString> palindromeStringList = new HashSet<>();
    LongLength longLength = new LongLength();

    public String longestPalindrome(String s) {
        for (int i = 0; i < s.length(); i++) {
            reSize(i, s.charAt(i), s);
        }
        for (PalindromeString palindromeString : palindromeStringList) {
            if ((palindromeString.right - palindromeString.left) == longLength.length) {
                return s.substring(palindromeString.left, palindromeString.right);
            }
        }
        return "";
    }

    private void reSize(int index, char c, String s) {
        Iterator<PalindromeString> iterator = palindromeStringList.iterator();
        Set<PalindromeString> palindromeStringListTem = new HashSet<>();
        while (iterator.hasNext()) {
            PalindromeString tem = iterator.next();
            // 如果是这个情况，说明这个对象是刚放进入的，直接跳过就行
            if (tem.right == index) {
                if(isValid(tem.left-1, s)){
                    if(isPalindromic(s.substring(tem.left-1, tem.right+1))){
                        palindromeStringListTem.add(new PalindromeString(tem.left - 1, tem.right + 1));
                        longLength.length = Math.max(longLength.length, tem.right - tem.left + 2);
                    }else if(isPalindromic(s.substring(tem.left, tem.right + 1))){
                        palindromeStringListTem.add(new PalindromeString(tem.left, tem.right + 1));
                        longLength.length = Math.max(longLength.length, tem.right - tem.left + 1);
                    }
                }else {
                    if(isPalindromic(s.substring(tem.left, tem.right + 1))){
                        palindromeStringListTem.add(new PalindromeString(tem.left, tem.right + 1));
                        longLength.length = Math.max(longLength.length, tem.right - tem.left + 1);
                    }
                }
            }
            if ((tem.right - tem.left) != 1 && (tem.right - tem.left) >= longLength.length) {
                palindromeStringListTem.add(tem);
            }
        }
        palindromeStringListTem.add(new PalindromeString(index, index + 1));
        longLength.length = Math.max(longLength.length, 1);
        palindromeStringList = palindromeStringListTem;
    }

    private boolean isValid(int index, String s) {
        return index >= 0 && index < s.length();
    }

    private class LongLength {
        private int length;

        public int getLength() {
            return length;
        }

        public LongLength setLength(int length) {
            this.length = length;
            return this;
        }
    }

    private class PalindromeString {
        private int left;
        private int right;

        public int getLeft() {
            return left;
        }

        public PalindromeString setLeft(int left) {
            this.left = left;
            return this;
        }

        public int getRight() {
            return right;
        }

        public PalindromeString setRight(int right) {
            this.right = right;
            return this;
        }

        public PalindromeString(int left, int right) {
            this.left = left;
            this.right = right;
        }

        @Override
        public boolean equals(Object obj) {
            return (obj instanceof  PalindromeString) && (this.getLeft() == ((PalindromeString) obj).getLeft())
                    && (this.getRight() == ((PalindromeString) obj).getRight());
        }

        @Override
        public int hashCode() {
            int result = 17;
            result = 31 * result + left;
            result = 15 * result + right;
            return result;
        }
    }

    private boolean isPalindromic(String str) {
        for (int i = 0; i < (str.length() + 1) / 2; i++) {
            if (str.charAt(i) != str.charAt(str.length() - 1 - i)) {
                return false;
            }
        }
        return true;
    }

    public static void main(String args[]) {
        Solution2 solution = new Solution2();
        System.out.println(solution.longestPalindrome("aaaa"));
    }
}
