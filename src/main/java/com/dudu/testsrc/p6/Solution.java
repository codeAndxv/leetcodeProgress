package com.codebest.SixProblem;

/**
 * @author code
 * @create 2018-10-07 11:55
 */

class Solution {
    public String convert(String s, int numRows) {
        //整体思路是   通过计算余数，0~(numRows-1)*2-1,第一行是0   最后一行是numRows-1,中间行是i和(numRows-1)*2-i
        int length = s.length();
        StringBuilder builder = new StringBuilder("");

        //约数
        int num = (numRows-1)*2;
        //从第一行开始算
        if(num>0){
            for(int i = 0;i<numRows;i++){
                for(int j = 0;j<=length/num;j++){
                    if(i==0||i==num/2){
                        if(j*num+i<length){
                            builder.append(s.charAt(j*num+i));
                        }
                    }else{
                        if(j*num+i<length){
                            builder.append(s.charAt(j*num+i));
                        }
                        if(j*num+num-i<length){
                            builder.append(s.charAt(j*num+num-i));
                        }
                    }
                }
            }
            return builder.toString();
        }else{
            return s;
        }


    }
}