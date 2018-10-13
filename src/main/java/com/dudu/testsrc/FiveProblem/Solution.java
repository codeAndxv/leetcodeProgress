package com.codebest.FiveProblem;

/**
 * @author code
 * @create 2018-10-06 22:11
 */

class Solution {
    public String longestPalindrome(String s) {
        String resultStr = "";
        int resultInt = 0;
        //有两种可能性，一种是偶数长度的回文字符串，另一种是奇数长度的回文字符串，
        //判断s.chatAt(i)==s.chatAt(i-(F[s.length-1]+2)),i from (s.length-F[s.length-1]-2+1)/2 to s.length-F[s.length-1]-2
        //F[s.length] = max{F[s.length-1],F[s.length]}

        int resultInts[] = new int[s.length()+1];
        resultInts[0] = 0;

        for(int i = 1 ;i <s.length()+1;i++){
            int tem = 0;
            String temStr = "";
            if((i-resultInts[i-1]-2>=0) && isPalindromic(s.substring(i-resultInts[i-1]-2,i))){
                tem = resultInts[i-1]+2;
                temStr = s.substring(i-resultInts[i-1]-2,i);
            }
            if((tem==0)&&(i-resultInts[i-1]-1>=0)&&isPalindromic(s.substring(i-resultInts[i-1]-1, i))){
                tem = resultInts[i-1]+1;
                temStr = s.substring(i-resultInts[i-1]-1, i);
            }
            if(resultInts[i-1] < tem){
                resultInts[i] = tem;
                resultInt = tem;
                resultStr = temStr;
            }else {
                resultInts[i] = resultInts[i-1];
                resultInt = resultInts[i-1];
            }
        }

        return resultStr;
    }

    public boolean isPalindromic(String str){


        for(int i=0;i<(str.length()+1)/2;i++){
            if(str.charAt(i)!=str.charAt(str.length()-1-i)){
                return false;
            }
        }
        return true;
    }



    public static void main(String args[]){
        Solution solution = new Solution();
        System.out.println(solution.isPalindromic("aba"));
    }


}