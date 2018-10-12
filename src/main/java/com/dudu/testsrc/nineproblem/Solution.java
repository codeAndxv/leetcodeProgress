package com.dudu.testsrc.nineproblem;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public boolean isPalindrome(int x) {
        if(x<0){
            return false;
        }
        List<Integer> integerList = new ArrayList<>();
        int shang = x;
        while(shang/10 > 0){
            integerList.add(shang%10);
            shang = shang/10;
        }
        integerList.add(shang%10);

        int[] nums = new int[integerList.size()];
        int i = 0;
        for(int num:integerList){
            nums[i] = num;
            i++;
        }
        for(i = 0; i<(nums.length+1/2); i++){
            if(nums[i] != nums[nums.length-1-i]){
                return false;
            }
        }
        return true;
    }
    @Test
    public void test(){
        isPalindrome(121);
    }
}
