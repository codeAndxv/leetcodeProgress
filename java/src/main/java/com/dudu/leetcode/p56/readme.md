### 56. Merge Intervals  Medium
Given a collection of intervals, merge all overlapping intervals.

#### Example 1:
```$xslt
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```
#### Example 2:
```$xslt
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

```
**NOTE**: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

#### Result
Runtime: 4 ms, faster than 94.78% of Java online submissions for Merge Intervals.

Memory Usage: 44.4 MB, less than 34.79% of Java online submissions for Merge Intervals.

#### solution
本来没想写思路，但没想得到结果还不错，就顺手写一下。

这道题乍看起来很简单，但仔细想，又感觉有点难。

在现在的我看来，这其实是排序算法的变种。

对于一个新范围`r`的插入，`r`会影响所有跟它有关的范围。**note**: `r`是闭区间。
所以在新范围`r`插入时，所要做的事情就是把合并`r`所有影响到的范围以及`r`，组成一个新的最大范围。

我最终提交的代码是插入排序的变种。 可以进一步改进使用合并排序，这样时间复杂度会变成O(nlogn)。

