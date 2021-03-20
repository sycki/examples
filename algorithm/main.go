package main

import "fmt"

func main() {
	var arr = []int{5, 2, 3, 5, 3, 4, 9}
	v := maxProfitD(arr)
	println(count)
	println(v)
}

// 在两个数组中查找中位数，使用归并方式实现
// v := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, a, b, al, bl int = 0, 0, 0, len(nums1), len(nums2)
	r := make([]int, al+bl)
	for {
		if al <= a {
			for _, v := range nums2[b:] {
				r[i] = v
				i++
			}
			break
		}
		if bl <= b {
			for _, v := range nums1[a:] {
				r[i] = v
				i++
			}
			break
		}
		if nums1[a] < nums2[b] {
			r[i] = nums1[a]
			a++
			i++
		} else {
			r[i] = nums2[b]
			b++
			i++
		}
	}
	if i == 0 {
		return 0
	}

	fmt.Println(r)
	if i%2 == 1 {
		return float64(r[i/2])
	} else {
		center := i / 2
		return (float64(r[center-1]) + float64(r[center])) / 2
	}
}

// 数组中的值表示每一天的价格，只能进行一次交易，计算最大收益
// var arr = []int{5, 2, 3, 5, 3, 4, 2, 3}
// return 3
func maxProfit(prices []int) int {
	var min, max int = prices[0], 0
	for _, p := range prices {
		if p < min {
			min = p
		}
		v := p - min
		if v > max {
			max = v
		}
	}
	return max
}

// 同上，但使用动态规划实现
// var arr = []int{5, 2, 3, 5, 3, 4, 2, 3}
// return 3
func maxProfitD(prices []int) int {
	var dp = make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(dp)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 数组中的值表示每一天的价格，只能进行k次交易，计算最大收益
// var arr = []int{5, 2, 3, 5, 2, 4, 3, 6}
// return 4
// func maxProfitK(prices []int) int {
// 	var r = make([][2][2]int, len(prices))
// 	for i := 1; i < len(prices); i++ {
// 		for k := 0; k < times; k ++ {
// 			r[i][k][0]
// 		}
// 	}
// 	return min, max
// }

// 数组中的值表示每一天的价格，可以进行任意次交易，计算最大收益
// var arr = []int{5, 2, 3, 5, 3, 4, 2, 3}
// return 5
func maxProfitM(prices []int) int {
	var max int = 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			max += profit
		}
	}
	return max
}

// 将链表反转
// a>b>c => a<b<c
// return c
type Node struct {
	Id   string
	Next *Node
}

func reverse(current, parent *Node) *Node {
	if current == nil {
		return parent
	}
	next := current.Next
	current.Next = parent
	parent = current
	return reverse(next, parent)
}

// 计算第N个斐波那契数，普通递归，O(2^N)
// var arr = []int{5, 2, 3, 5, 3, 4, 2, 3}
// return 3
var count int

func fib(n int) int {
	count++
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 计算第N个斐波那契数，备忘录方式，避免重复计算，O(N)
// var arr = []int{5, 2, 3, 5, 3, 4, 2, 3}
// return 3
func fibEntry(n int, entry []int) int {
	count++
	if n == 1 || n == 2 {
		return 1
	}
	left := n - 1
	right := n - 2
	if entry[left] == 0 {
		entry[left] = fibEntry(left, entry)
	}
	if entry[right] == 0 {
		entry[right] = fibEntry(right, entry)
	}
	return entry[left] + entry[right]
}

// 计算第N个斐波那契数，动规模式，避免重复计算，O(N)
// var arr = []int{5, 2, 3, 5, 3, 4, 2, 3}
// return 3
func fibD(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n)
	dp[0], dp[1] = 1, 1
	for i := 3; i <= n; i++ {
		count++
		dp[i-1] = dp[i-1-1] + dp[i-1-2]
	}
	return dp[n-1]
}

func coinChange(n int, coins []int, entry []int) int {
	count++
	if n < 1 {
		return 0
	}
	var r int = n
	for _, coin := range coins {
		var problem = n - coin
		if problem < 1 {
			continue
		}
		if entry[problem] == 0 {
			entry[problem] = coinChange(n-coin, coins, entry)
		}
		child := entry[problem]
		v := 1 + child
		if v > 0 && v < r {
			r = v
		}
	}
	return r
}
