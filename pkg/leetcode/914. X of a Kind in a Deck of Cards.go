package leetcode

func hasGroupsSizeX(deck []int) bool {
	numMap := make(map[int]int)
	for _, value := range deck {
		if _, exist := numMap[value]; !exist {
			numMap[value] = 1
		} else {
			numMap[value] += 1
		}
	}
	var lastFactors []int
	for _, value := range numMap {
		currentFactores := primeFactors(value)
		if len(currentFactores) < 1 {
			return false
		}
		if lastFactors == nil {
			lastFactors = currentFactores
			continue
		} else {
			lastFactors = findCommonSubset(lastFactors, currentFactores)
		}
		if len(lastFactors) < 1 {
			return false
		}
	}
	return true
}

// 判断一个数是否为质数
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// 找出给定数的所有不重复的质数因子
func primeFactors(num int) []int {
	factors := []int{}
	seen := make(map[int]bool)

	for i := 2; i <= num; i++ {
		if num%i == 0 && isPrime(i) && !seen[i] {
			factors = append(factors, i)
			seen[i] = true
			num /= i
			i--
		}
	}
	return factors
}

// 找出两个[]int的共同子集
func findCommonSubset(arr1, arr2 []int) []int {
	// 将第一个[]int中的元素放入哈希集合中
	set := make(map[int]bool)
	for _, num := range arr1 {
		set[num] = true
	}

	// 遍历第二个[]int，检查每个元素是否在哈希集合中
	commonSubset := []int{}
	for _, num := range arr2 {
		if set[num] {
			commonSubset = append(commonSubset, num)
		}
	}

	return commonSubset
}
