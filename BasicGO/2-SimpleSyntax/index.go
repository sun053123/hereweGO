package main

import "fmt"

func main() {

	// Varible ทั้งหมดจะสามารถ ใส่ค่าไปได้เลย หรือ ประกาศ เฉย ๆ ก็ได้

	var i1 int    // ค่าเริ่มต้น 0
	var s1 string // ค่าเริ่มต้น ""
	var b1 bool   // ค่าเริ่มต้น false

	var in1 int = 20
	var st1 string = "hello_world"
	var bo1 bool = true

	fmt.Println("Varible : ", i1, s1, b1, in1, st1, bo1)

	// Type inference สามารถใส่ค่าประกาศได้เลยโดยไม่ต้องมีตัวบอก type  มันจะ detect เอง
	// หมายเหตุ สามารถทำได้แค่ใน function เท่านั้น

	var i2 = 20
	var s2 = "hello_world"
	var b2 = true

	in2 := 30
	st2 := "hi world"
	bo2 := true

	fmt.Println("Type Inference : ", i2, s2, b2, in2, st2, bo2)

	// Constant เป็น Imutable *ไม่สามารถเปลี่ยนค่าได้* ต้องทำการ assign ค่าเริ่มต้นก่อน
	// จะประกาศ Type หรือไม่ก็ได้ และ ไม่ต้องเรียกใช้งานก็ไม่ error
	const defaultValue int = 1
	const defaultTitle = "Go"
	const defaultStatus = true


	const (
		sunday = iota // iota คือ enum ประกาศไล่ตั้งแต่ 0 แล้ว + 1 ไปเรื่อย ๆ // สามารถ iota + 1 เพื่อเริ่มตั้งแต่ 1 ก็ได้
		monday
		tuesday
		wendsday
		thursday
		friday
		saturday
	)
	fmt.Println(sunday, monday, tuesday)

	// ลองใช้ function calc
	calc()
	// ลองใช้ การเรียก function ใน print เพราะตัวมัน return กับมาเป็น string
	fmt.Println("Return a + b = ", calcReturn())

	// ลองส่ง varible ที่ประกาศจาก Main 
	a := 1
	b := 2
	calcRecive(a, b)

	fmt.Println("Recive and Return a + b = ", calcReciveReturn(a, b))

	c, d := swap(a, b)
	fmt.Printf("func swap a = %d , b = %d \n", c, d)

	// Example เรียก function เพื่อต่อ string ไปเรื่อย ๆ

	const (
		r = "Asun "
		t = "Hi my name is  "
		y = "Apple"
		u = "i like "
	)

	fmt.Println(stringAdd(r, t, y, u))

	// print I am Asun and i like megumin more than Apple

	strr := stringAdd2(r,y,u)
	fmt.Println(strr)

}

// function เรียกทำงาน ไม่มีการส่งค่า และ รับค่าใด ๆ
func calc() {

	a := 1
	b := 2
	c := a + b

	fmt.Println("on calc func = ", c)

}

func calcReturn() int {
	a := 1
	b := 2

	return a + b
}

func calcRecive(a, b int) {
	fmt.Println(a + b)
}

func calcReciveReturn(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func stringAdd(r, t, y, u string) string {
	return fmt.Sprint(t, r, u, y)
}

func stringAdd2(r,y,u string) string {
	return fmt.Sprintf("I am %sand %smegumin more than %s", r , u, y)
}