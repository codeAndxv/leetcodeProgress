package com.dudu.leetcode.p494;

import java.util.Arrays;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/10/2019 6:14 PM
 * @email zhao.lu@parcelx.io
 */
public class Solution1 {
    /**
     * It can be easily observed that in the last approach, a lot of redundant function calls could be made with the same value of ii as the current index and the same value of sum as the current sum, since the same values could be obtained through multiple paths in the recursion tree. In order to remove this redundancy, we make use of memoization as well to store the results which have been calculated earlier.
     *
     * Thus, for every call to calculate(nums, i, sum, S), we store the result obtained in memo[i][sum + 1000]. The factor of 1000 has been added as an offset to the sum value to map all the sums possible to positive integer range. By making use of memoization, we can prune the search space to a good extent.
     */
    int count = 0;
    public int findTargetSumWays(int[] nums, int S) {
        int[][] memo = new int[nums.length][2001];
        for (int[] row: memo)
            Arrays.fill(row, Integer.MIN_VALUE);
        return calculate(nums, 0, 0, S, memo);
    }
    public int calculate(int[] nums, int i, int sum, int S, int[][] memo) {
        if (i == nums.length) {
            if (sum == S)
                return 1;
            else
                return 0;
        } else {
            if (memo[i][sum + 1000] != Integer.MIN_VALUE) {
                return memo[i][sum + 1000];
            }
            int add = calculate(nums, i + 1, sum + nums[i], S, memo);
            int subtract = calculate(nums, i + 1, sum - nums[i], S, memo);
            memo[i][sum + 1000] = add + subtract;
            return memo[i][sum + 1000];
        }
    }
}
