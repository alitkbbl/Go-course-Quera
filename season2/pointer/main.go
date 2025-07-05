package main

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	minNum := (*numbers)[0]
	for _, num := range (*numbers)[1:] {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}
func ReverseSlice(numbers *[]int) {
	nums := *numbers
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func SwapElements(numbers *[]int, i, j int) {
	n := len(*numbers)
	if i < 0 || i >= n || j < 0 || j >= n {
		return
	}
	(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}
