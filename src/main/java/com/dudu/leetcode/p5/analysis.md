### 解题思想
像这种题目，都是要利用之前的计算结果。如果找不到合理的关系，会导致时间复杂度变成指数级别。

而找到前一次计算与后一次计算之间的关系，是将时间复杂度从指数级别到多项式级别的关键。

### 本题思想
回文字符串有两种，一种是偶数长度的回文，一种是奇数长度的回文。

而不管哪种回文，去掉首尾两个字符，所剩下的字符串也一定是回文字符串。在这里认为空字符串也是一种回文字符串。

那么对于一个最长回文字符串一定是有一个回文字符串生成的。这就意味着我们只需要关注那些可能成为最长回文字符串的回文字符串。

### 实例分析 
以"babad"为例，我们遍历整个字符串。记最长回文字符串的长度为l, 有可能成为最长回文字符串的字符串放在一个list中。

首先是'b'，记l为1， 把'b'放在list中。

然后是'a'，因为'b'!='a'，并且b的左边没有字符与a相等。所以记l为1，把b从list中删除，并且把a放在list中。

再是'b'，因为'b'!='a'，a的左边有字符与b相等。所以记l为3， 把a从list中删除，并bab放入，同时再把b放入（因为该位置的b有可能成为最长回文字符串的一部分。）

再是a， 因为'b'!='a'，b的左边有字符与a相等，所以记l为3，把b从list中删除，并把aba放入。
