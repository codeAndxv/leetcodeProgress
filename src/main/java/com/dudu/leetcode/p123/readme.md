#### 123. Best Time to Buy and Sell Stock III
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

**Note**: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

##### Example 1:
```text
Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
```
##### Example 2:
```text
Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
```
##### Example 3:
```text
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

这道题没有想出来答案，最后是看leetcode上一位网友的解法，https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/457495/Java-DP-Solution-1ms-beats-99.74。

感觉在解这种题的时候，总是差了一点。