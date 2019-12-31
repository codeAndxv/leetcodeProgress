package com.dudu.leetcode.p20;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/30/2019 10:49 AM
 * @email zhao.lu@parcelx.io
 */
public class Solution {
    public boolean isValid(String s) {
        Stack stack = new Stack();
        Map<Character,Character> elementMap = new HashMap<>();
        elementMap.put('(', ')');
        elementMap.put('[', ']');
        elementMap.put('{', '}');
        for(Character c:s.toCharArray()){
            if(elementMap.containsKey(c)){
                stack.push(c);
            }else {
                Character topC = stack.peek();
                if(topC==null){
                    return false;
                }
                if(elementMap.get(topC) != c){
                    return false;
                }
                stack.pop();
            }
        }
        return stack.list.size() == 0;

    }
    class Stack{
        List<Character> list = new LinkedList<>();
        boolean push(Character c){
            return list.add(c);
        }

        Character pop(){
            if(list.size() >0){
                return list.remove(list.size()-1);
            }
            return null;
        }

        Character peek(){
            if(list.size() >0){
                return list.get(list.size()-1);
            }
            return null;
        }
    }
}
