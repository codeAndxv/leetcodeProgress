package com.dudu.leetcode.p10;

import java.util.LinkedList;
import java.util.List;
import java.util.function.IntFunction;

/**
 * @Author zhao.lu
 * @Email yx163luzhao@163.com
 * @Date 12/28/2019 7:37 PM
 * @Version 1.0
 * @Description
 **/
public class Solution {
    public boolean isMatch(String s, String p) {
        StatusNode node = null, startNode = null;
        for (int i = 0; i < p.length(); i++) {
            if (startNode == null) {
                startNode = new StatusNode(String.valueOf(p.charAt(i)));
                node = startNode;
            } else {
                if (p.charAt(i) == '*') {
                    if (node.value.length() > 1) {
                        int index = node.value.length() - 2;
                        for (; index >= 0; index--) {
                            if (node.value.charAt(index) != node.value.charAt(node.value.length() - 1)) {
                                index++;
                                break;
                            }
                        }
                        if (index == 0) {
                            node.value = node.value.substring(0, 1);
                        } else {
                            node.nextNode = new StatusNode(node.value.substring(index));
                            node.value = node.value.substring(0, index);
                            node = node.nextNode;
                        }
                    }
                    node.loop = true;
                } else if (p.charAt(i) == '.') {
                    node.nextNode = new StatusNode(".");
                    node = node.nextNode;
                } else {
                    if (!node.value.equals(".") && !node.loop) {
                        node.value = node.value + p.charAt(i);
                    } else if (node.loop && node.value.equals(String.valueOf(p.charAt(i)))) {

                    } else {
                        node.nextNode = new StatusNode(String.valueOf(p.charAt(i)));
                        node = node.nextNode;
                    }
                }
            }
        }
        return false;
    }

    private boolean match(StatusNode node, String s, int matchIndex) {
        if (node != null) {
            if (node.loop) {
                String temStr = "";
                int len = 0;
                for(;len<s.length(); len++){

                }
                while (s.substring(matchIndex, matchIndex + len).equals(temStr)) {
                    match(node.nextNode, s, matchIndex + len);
                    len++;

                }
            } else {
                return (s.substring(matchIndex, matchIndex + node.value.length()).equals(node.value)
                        || (node.value.equals("."))) && (match(node.nextNode, s, matchIndex + node.value.length()));
            }
        }
        return false;
    }

    class StatusNode {
        boolean loop = false;
        String value;
        StatusNode nextNode;

        StatusNode(String value) {
            this.value = value;
        }
    }
}
