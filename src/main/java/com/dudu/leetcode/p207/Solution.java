package com.dudu.leetcode.p207;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/17/2019 10:28 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    private boolean result = true;
    private List<Integer> allNode = new LinkedList<>();
    // judge whether exist ring
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        Map<Integer, List<Integer>> courses = new HashMap<>();
        for (int[] pair : prerequisites) {
            courses.putIfAbsent(pair[0], new LinkedList<>());
            courses.get(pair[0]).add(pair[1]);
        }
        for(int i = 0; i<numCourses; i++){
            allNode.add(i);
        }
        // if only have in-degree, can be deleted
        for(int[] pair : prerequisites){
            if(allNode.contains(pair[0])){
                List<Integer> recordList = new LinkedList<>();
                traverse(courses, pair[0], recordList);
            }
        }
        return result;
    }

    private void traverse(Map<Integer, List<Integer>> courses, Integer node, List<Integer> recordList){
        if(!result){
            return;
        }
        if(recordList.contains(node)){
            result = false;
            return;
        }
        recordList.add(node);
        allNode.remove(node);
        if(courses.get(node) != null){
            for(Integer nextNode : courses.get(node)){
                if(result){
                    traverse(courses, nextNode, recordList);
                }
            }
        }
        recordList.remove(node);
    }
}
