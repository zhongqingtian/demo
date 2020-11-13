package pattern

import "sync"

type Message struct {
	Header *Header
	Body   *Body
}
type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}
type Body struct {
	Items []string
}

// 如果按照直接的对象创建方式，创建逻辑应该是这样的：
// 多层的嵌套实例化
func newDuiXiang() {
	message := Message{
		Header: &Header{
			SrcAddr:  "192.168.0.1",
			SrcPort:  1234,
			DestAddr: "192.168.0.2",
			DestPort: 8080,
			Items:    make(map[string]string),
		},
		Body: &Body{
			Items: make([]string, 0),
		},
	}
	// 需要知道对象的实现细节
	message.Header.Items["contents"] = "application/json"
	message.Body.Items = append(message.Body.Items, "record1")
	message.Body.Items = append(message.Body.Items, "record2")
}

/*-----------------------------下面使用建造者模式-----------------------------*/

// 虽然Message结构体嵌套的层次不多，但是从其创建的代码来看，确实存在对对象使用者不友好和代码可读性差的缺点。下面我们引入建造者模式对代码进行重构：
// Message对象的Builder对象
type builder struct {
	once *sync.Once
	msg  *Message
}

// 返回Builder对象
func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &Message{Header: &Header{}, Body: &Body{}},
	}
}

// 以下是对Message成员对构建方法
func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}
func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}
func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}
func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}
func (b *builder) WithHeaderItem(key, value string) *builder {
	// 保证map只初始化一次
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}
func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

// 创建Message对象，在最后一步调用
func (b *builder) Build() *Message {
	return b.msg
}
