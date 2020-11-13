package pattern

/*桥接模式本质上就是面向接口编程，可以给系统带来很好的灵活性和可扩展性。*/
/*桥接模式将模块的抽象部分和实现部分进行分离，让它们能够往各自的方向扩展，从而达到解耦的目的。*/
type Input1 interface {
	Plugin
	Receive() *Message
}

type Filter1 interface {
	Plugin
	Process(msg *Message) *Message
}

type Output1 interface {
	Plugin
	Send(msg *Message)
}

// 一个Pipeline由input、filter、output三个Plugin组成
type Pipeline2 struct {
	//status Status
	input  Input1
	filter Filter1
	output Output1
}

// 通过抽象接口来使用，看不到底层的实现细节
func (p *Pipeline2) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}
