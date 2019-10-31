在讨论区发现了一个好的答案
```
class Solution {
    public int maxArea(int[] height) {
        
        int low = 0;
        int high = height.length-1;
        int maxArea = 0;
        int distance = height.length-1;
        while(low < high) {
            int area = Math.min(height[low], height[high]) * distance;
            maxArea = area > maxArea ? area : maxArea;
            if(height[low] < height[high]) {
                low++;
            } else {
                high--;
            }
            distance--;
        }
        return maxArea;
    }
}

```

