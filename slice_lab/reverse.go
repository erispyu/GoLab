package main

func reverse(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	reverse(nums)
// 	fmt.Println(nums)
// 	reverse(nums[3:])
// 	fmt.Println(nums)
// }
