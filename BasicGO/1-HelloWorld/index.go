package main

//basic package จะชื่ออะไรก็ได้ และไม่ได้สนการตั้งชื่อไฟล์

//เมื่อจะเขียน go ต้อง init ด้วยคำสั่ง go mod init github.com/<username>/<repo> หรือ go mod init <filename> เมื่อไม่อยาก public
//เป็นการ set environment ของ golang
//ทั้งนี้ Golang Compile เป็น Binary จึง Lightweight และ ทำงานรวดเร็ว สามารถสั่ง build ได้ด้วยคำสั่ง go build <filename.go> จะได้ bin file
//Test ของ Goland สามารถทืำได้เลยเมื่อตั้งชื่อไฟล์เหมือนเดิมที่จะ Test โดยใส่ test ด้านหลัง ex. main_test.go

//Golang สามารถเขียนได้บน IDE อะไรก็ได้เช่น Visual Studio Code, Goland, Vim, Sublimetext, etc.
//ใน Vscode มี extendstion แนะนำคือ 1.Go, 2.Error Lens, 3.Test Explorer UI

//ศึกษาเพิ่มเติมได้ที่ https://go.dev/tour/welcome/1

import "fmt"

// เมื่อจะทำการ Print อะไรออกมา สำคัญจะต้องประกาศ import "fmt"

// เป็น Function แรกที่จะทำงานเมื่อกด Run จะไป Build เป็น Cache ข้างหลังบ้่านมาแสดง
// สามารถสั่งรันด้วยคำสั่ง go

func main() {

	//print แบบที่ 1
	fmt.Println("Hello World.")

	//print แบบที่ 2
	text1 := "Hello"
	var text2 = "World."
	fmt.Println(text1, text2)

	//print แบบที่ 3
	fmt.Printf("%s %s \n", text1, text2)

	//print แบบที่ 4
	fmt.Println(Echo1(text1, text2))

	//print แบบที่ 5
	a := Echo2(text1, text2)
	fmt.Println(a)

}

func Echo1(text1, text2 string) string {
	return text1 + " " + text2
}

func Echo2(text1, text2 string) string {
	return fmt.Sprintf(text1 + " " + text2)
}
