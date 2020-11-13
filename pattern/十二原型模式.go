package pattern

/*
原型模式主要解决对象复制的问题，它的核心就是clone()方法，返回Prototype对象的复制品。在程序设计过程中，往往会遇到有一些场景需要大量相同的对象，如果不使用原型模式，那么我们可能会这样进行对象的创建：新创建一个相同对象的实例，然后遍历原始对象的所有成员变量， 并将成员变量值复制到新对象中。这种方法的缺点很明显，那就是使用者必须知道对象的实现细节，导致代码之间的耦合。另外，对象很有可能存在除了对象本身以外不可见的变量，这种情况下该方法就行不通了。
对于这种情况，更好的方法就是使用原型模式，将复制逻辑委托给对象本身，这样，上述两个问题也都迎刃而解了。
*/

// 原型复制抽象接口
type Prototype interface {
	clone() Prototype
}

// 和建造者模式公用一个struct
/*type Message2 struct {
	Header *Header
	Body   *Body
}*/

func (m *Message) clone() Prototype {
	msg := *m
	return &msg
}
