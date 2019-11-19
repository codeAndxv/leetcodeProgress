*最近发现，在睡觉之前看一道leetcode上级别为hard的题目能够明显的提高入睡速度。*

今天看了第76题，题目如下：

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

**Example**:
```.java
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"

```
**Note**:

- If there is no such window in S that covers all characters in T, return the empty string "".
- If there is such window, you are guaranteed that there will always be only one unique minimum window in S.


### solution:
直接讲一下解题思路，不码代码了。

从要求的O(n)可以知道，这道题一定是需要充分所得信息，从一次遍历中获取到信息。

而对于一个字符串来说，最重要的信息就是字符以及字符对应的位置。

那么直接遍历整个字符串，并且把信息存起来，如下：

```.env
A   B   C   D   E   N   O
0   3   5   1   4   11  2
10  9   12  7   8       6
```
对于目标字符串`ABC`，我们关注的是前三列。
### 但问题来了，我发现这么抽信息并没有什么作用，还是需要计算最小的窗口。

那么就直接想怎么计算最小窗口吧

最小窗口，嗯，队列这个数据结构好像很不错。

那么方法有了：

```
队列：
---------
A0, B3, C5
---------
||
||
---------
A0, B3, C5, B9, A10
---------
||
||
---------
C5, B9, A10
---------
||
||
---------
B9, A10, C12
---------
```
过程很明显吧， 每次把目标字符串的字符信息放进入，如果和队列头字符相同，就把队列头剔除。