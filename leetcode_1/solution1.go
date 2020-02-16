package solution1

//1. target - num1 = res
//2. check if res = num...
//3. return two index
func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if target-nums[i] == nums[j] && i != j {
				res = append(res, i, j)
				goto loop
			}
		}
	}
loop:
	return res
}
