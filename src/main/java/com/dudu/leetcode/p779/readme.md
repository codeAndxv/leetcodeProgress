这个题目也是比较有意思， 大概思考了半个多小时吧。 根据N的次数，逐轮进行变换，肯定是最愚蠢的选择。

思考一下，在逐轮计算中，出现了大量的不需要的位。而我们只关心最后K-th的变换。

那么思考

```.env
       0                        1
     /  \                     /  \
    0   1                    1   0 
```
我们只需要判断在得到K-th中发生的变换即可。

以N=4,K=5为例，5的路径为 5->3->2->1, 变位顺位为0->1->1->1