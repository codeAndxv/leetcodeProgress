package com.codebest.FourProblem;

/**
 * @author code
 * @create 2018-09-28 10:11
 */

/**
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 * Example 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 */

class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length, n = nums2.length;

        int[] nums3 = new int[m+n];
        int m1 = 0, n1 = 0;

        while(m1<m || n1<n){

            if(n1<n && m1<m){
                if(nums1[m1]>nums2[n1]){
                    nums3[m1+n1] = nums2[n1];
                    n1++;
                }else {
                    nums3[m1+n1] = nums1[m1];
                    m1++;
                }
            }else if(n1<n){
                nums3[m1+n1] = nums2[n1];
                n1++;
            }else if(m1<m){
                nums3[m1+n1] = nums1[m1];
                m1++;
            }
        }
        if ((m+n)%2==0){
            return (nums3[(m+n)/2-1]+nums3[(m+n)/2])/2.0;
        }else {
            return nums3[(m+n+1)/2-1];
        }

    }
}