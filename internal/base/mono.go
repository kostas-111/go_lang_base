package base

func Mono(nums []int) bool {

    arraySize := len(nums)

    if arraySize <= 2 {
        return true
    }

    isIncreasing := true
    isDecreasing := true

    for i := 0; i < arraySize - 1; i++ {
        if nums[i] > nums[i + 1] {
            isIncreasing = false
        }
        if nums[i] < nums[i + 1] {
            isDecreasing = false
        }
    }
	return isIncreasing || isDecreasing
}