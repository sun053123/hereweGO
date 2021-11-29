package main

import "fmt"

func main() {
	// ทำได้เฉพาะ Sort Array เท่านั้น การ Random ค่าจึงจะต้องเป็นแบบ min => max
	var arr []int
	maxsize := 100
	//ทำการ assign ค่าให้ arr โดยเริ่มจาก 1 - 100
	for i := 0; i <= maxsize; i++ {
		arr = append(arr, i)
	}
	fmt.Println("generate arr finished ...")

	target := 55 //เปลี่ยนแปลงค่านี้ เพื่อจะ set target
	var low int = 0
	var high int = len(arr) - 1

	a := findBinarySearch(target, low, high, arr)
	b := findLinearSearch(target, low, high, arr)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Program exit . . .")

}

func findBinarySearch(target, low, high int, arr []int) string {

	index := -1 //set index -1 เมื่อไม่สามารถหาได้ จะ return -1 กลับไป
	count := 0

	for low <= high {
		mid := (low + high) / 2
		count++
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] == target {
			index = mid
			return fmt.Sprintf("BS finished : index = %d, target = %d, count = %d", index, target, count)

			// BigO = O(log N)
			// best case = 1 เมื่อตรงกลาง arr พอดี
			// worst case = 7 *เมื่อเราใช้ความยาวของ arr = 100 เพราะ 2^7 = 128 , 2^6 = 64 จะอยู่ตรงกลาง
		}
	}
	return fmt.Sprintf("BS finished : index = %d, target = %d, count = %d", index, target, count)
}

func findLinearSearch(target, low, high int, arr []int) string {

	index := -2
	count := 0

	for i := 0; i < len(arr); i++ {
		count++
		if target == arr[i] {
			index = i
			return fmt.Sprintf("LS finished : index = %d, target = %d, count = %d", index, target, count)
		}
	}

	return fmt.Sprintf("LS finished : index = %d, target = %d, count = %d", index, target, count)
}
