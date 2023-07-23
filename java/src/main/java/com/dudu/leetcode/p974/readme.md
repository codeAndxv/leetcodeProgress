#### 974. Subarray Sums Divisible by K
Given an array `A` of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by `K`.
Example 1:
```
Input: A = [4,5,0,-2,-3,1], K = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by K = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
```

Note:
```$xslt
1. 1 <= A.length <= 30000
2. -10000 <= A[i] <= 10000
3. 2 <= K <= 10000
```

这道题可以好好聊聊。非常有意思的一道题。

我的解题阶段，三个阶段。

#### 第一个阶段
没有什么头绪，最直接的办法就是把所有的子数组的和都求出来。然后看是不是能够被整除。

但这样会出现很多的重复操作。
#### 第二个阶段， 动态规划
出现很多重复的操作，很直接的想法就是使用动态规划来记录已经计算过的结果。

那么第一版就出炉了， 
```$xslt
       int result = 0;
        int[][]  recordTable = new int[A.length][A.length+1];
        for(int i = 0; i<A.length; i++){
            for(int j = i+1; j<A.length+1; j++){
                recordTable[i][j] = j==i+1? A[i]:recordTable[i][j-1]+A[j-1];
                if(recordTable[i][j] % K ==0){
                    result ++;
                }
            }
        }
        return result;
```
然后报Memory Limit Exceeded。

那么思考怎么优化，其实很明显，我们并不需要把之前计算过的结果都记录下来，只需要在计算出结果后判断这个结果是不是能够被整除就好了。

上面的代码，S(i,j)代表从i到j的子数组的和，不包括A[j]。

优化后的代码
```$xslt
 int result = 0;
        int recordTable = 0;
        // handle A
        for(int i = 0; i<A.length;i++){
            A[i] = A[i] % K ;
        }
        for (int i = 0; i < A.length; i++) {
            for (int j = i + 1; j < A.length + 1; j++) {
                recordTable = (j == i + 1 ? A[i] : recordTable + A[j - 1]) % K;
                if (recordTable % K == 0) {
                    result++;
                }
            }
        }
        return result;
```
很好内存超出的错误已经解决了，但爆出了新的错误，Time Limit Exceeded

时间超出的错误。优化后的代码的时间复杂度为O(A.len^2)，如果A.len很大时，就出现了上面的时间超出的错误。
#### 第三个阶段
其实我们并不关心实际的子数组的和是多少，以上面的思路继续下去，当循环到下一个数值时，我关心之前能和下一个数值形成K的倍数的数量。

很接近关键了， 我们建立一个recordNum int[K]，把K的余数等于index的数量记录到recordNum[index]上。

以A = [4,5,0,-2,-3,1], K=5为例。

recordNum的初始状态为[0,0,0,0,0],碰到4后，我们首先把recordNum向前推4位。然后把recordNum[4] += 1;

推m位的操作是把recordNum当成一个环，recordNum[(i+m)%K]=recordNum[i]。做完这个操作后，reocrdNum[0]的值就是加上A[i]后的所有的子数组能够被K整除的数量。

但还有一个很重要的优化，如果按照上面的操作，时间复杂度会是O(K*A.len)。其实推m位的操作可以理解成是一个基准指针后退m位的操作。

以上面的[0,0,0,0,0]为例， pointer一开始指向0， 当碰到4后，pointer后退4位，相当于前进1位，指向了1，同时因为基准变了，应该把recodeNum[(pointer+4)%K] += 1.

就变成了recodeNum[0]+=1。当前的recodeNum[pointer]就是和上面一样需要的数量。

这样的结果就是时间复杂度变成了O(A.len)。

代码：
```$xslt
public int subarraysDivByK(int[] A, int K) {
        int[] recordNum = new int[K];
        int pointer = 0;
        // handle A
        for (int i = 0; i < A.length; i++) {
            A[i] = (A[i] % K + K) % K;
        }
        int result = 0;
        // move the pointer, and add num according to pointer and A[i], add result in the end
        for(int i = 0;i <A.length; i++){
            pointer = ((-A[i] + K)%K + pointer) % K;
            recordNum[(pointer + A[i]) % K] += 1;
            result += recordNum[pointer];
        }
        return result;
    }
```

满满的细节啊。