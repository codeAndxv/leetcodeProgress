package leetcode

func ComputeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	nums1 := []int{ax1, ax2, bx1, bx2}
	nums2 := []int{ay1, ay2, by1, by2}
	nums1 = sortN2(nums1)
	nums2 = sortN2(nums2)
	overlap := (nums1[2] - nums1[1]) * (nums2[2] - nums2[1])
	all := abs(ax1-ax2)*abs(ay1-ay2) + abs(bx1-bx2)*abs(by1-by2)

	if (inRange(bx1, ax1, ax2) || inRange(bx2, ax1, ax2) || inRange(ax1, bx1, bx2) || inRange(ax2, bx1, bx2)) && (inRange(by1, ay1, ay2) || inRange(by2, ay1, ay2) || inRange(ay1, by1, by2) || inRange(ay2, by1, by2)) {
		return all - overlap
	} else {
		return all
	}
}

func inRange(target int, num1 int, num2 int) bool {
	return target >= num1 && target <= num2
}

/**
class Solution {
    public int computeArea(int ax1, int ay1, int ax2, int ay2, int bx1, int by1, int bx2, int by2) {
        int x = Math.max(0, Math.min(ax2, bx2) - Math.max(ax1, bx1));
        int y = Math.max(0, Math.min(ay2, by2) - Math.max(ay1, by1));
        return (ax2 - ax1) * (ay2 - ay1) + (bx2 - bx1) * (by2 - by1) - (x * y);
    }
}

作者：宫水三叶
链接：https://leetcode.cn/problems/rectangle-area/solutions/1024841/gong-shui-san-xie-yun-yong-rong-chi-yuan-hzit/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
