package main

import "fmt"

type Student struct {
	Name   string
	Number int
	Grade  int
}

type Teacher struct {
	Name string
}

type Person interface {
	getEmail() string
}

func main() {
	var s, t Person
	s = Student{
		Name:   "Yamada",
		Number: 999,
		Grade:  5,
	}
	t = Teacher{Name: "Tsubomi"}

	cxtStu := SendEmail(s)
	fmt.Println(cxtStu)

	cxtTea := SendEmail(t)
	fmt.Println(cxtTea)
}

func (s Student) getEmail() string {
	return fmt.Sprintf("%s@student.ed.jp", s.Name)
}

func (t Teacher) getEmail() string {
	return fmt.Sprintf("%s@teacher.ed.jp", t.Name)
}

func SendEmail(p Person) string {
	from := p.getEmail()
	context := `
  送信元 : ` + from + `
  これはテスト用のメールです。
  よろしくお願いします。
  `

	return context
}
