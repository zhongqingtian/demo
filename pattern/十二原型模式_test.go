package pattern

import "testing"

func TestPrototype(t *testing.T) {
	message := Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.2").
		WithDestPort(8080).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()
	// 复制一份消息
	newMessage := message.clone().(*Message) // 不用知道细节，clone 出一个一模一样的对象
	if newMessage.Header.SrcAddr != message.Header.SrcAddr {
		t.Errorf("Clone Message failed.")
	}
	if newMessage.Body.Items[0] != message.Body.Items[0] {
		t.Errorf("Clone Message failed.")
	}
}
