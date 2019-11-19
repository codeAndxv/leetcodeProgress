package com.dudu.leetcode.luogup1462;

import java.util.*;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 9/24/2019 11:42 AM
 * @email zhao.lu@parcelx.io
 */
public class Main {
    static class Node{
        public Integer fee;
        public SortedMap<Integer, Node> childNode = new TreeMap<>();
        public SortedMap<Integer, Road> roadList = new TreeMap<>();
    }

    static class Road{
        public Map<Integer, Node> nodeMap;
        public Integer bleed;
    }

    static SortedMap<Integer, Node> nodeMap;
    static SortedMap<Integer, Road> roadMap;

    static Integer minFee;
    static Integer allBleed;
    static Stack<Integer> integerStack;

    public static void tt(Node startNode){

    }

    public static void main(String[] args) {
        allBleed = 8;
    }

}
