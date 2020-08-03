package interview

type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main002() {
	//var peo People = Stduent{} // 类型必须一一匹配，否则报错，指针类型可以同时调用结构体类型以及指针类型。结构体类型不能调用指针类型
	//think := "bitch"
	//log.Print(peo.Speak(think))
}
