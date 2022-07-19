package try

import (
	"fmt"
	"time"
)

var nums []int
var mid int

func GuessIt(guess int, numberPool int) (tries int, result bool) {

	//for i := len(nums); i > 0; i-- {
	for i := 1; i <= numberPool; i++ {
		nums = append(nums, i)
		// nums[i-1] = i * 2
	}

	fmt.Println(nums)

	low := 0
	high := len(nums) - 1

	for low <= high {
		tries++
		mid = (low + high) / 2

		if nums[mid] == guess {
			fmt.Println(low, high, mid, nums[mid], guess)

			fmt.Println(nums[mid], " - верно!")

			result = true
			return
		}

		if nums[mid] > guess {
			fmt.Println(low, high, mid, nums[mid], guess)

			fmt.Println(nums[mid], " - слишком много... Пробуй снова.")
			time.Sleep(time.Second * time.Duration(1))
			high = mid - 1
		} else {
			fmt.Println(low, high, mid, nums[mid], guess)

			fmt.Println(nums[mid], " - слишком мало... Пробуй снова.")
			time.Sleep(time.Second * time.Duration(1))
			low = mid + 1
		}

	}
	fmt.Println(low, high, mid, nums[mid], guess)

	fmt.Println(guess, " - такого числа нет...")
	result = false

	return
}
