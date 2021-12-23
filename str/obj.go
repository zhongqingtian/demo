package str

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

type Field struct {
	Type  string      `json:"type"`
	Trans string      `json:"trans"`
	Value interface{} `json:"value"`
}
type Tr struct {
	Name  Field `json:"name"`
	Age   Field `json:"age"`
	Is    Field `json:"is"`
	Ctime Field `json:"ctime"`
}

func Compare() []string {
	s := `{
    "name":"ls",
    "age":100,
    "is":false,
    "ctime": 12121
}`
	s2 := `{
    "name":"ls",
    "age":0,
    "is":false,
    "ctime": 1
}`
	temp := Tr{}
	temp22 := Tr{}

	err := json.Unmarshal([]byte(s), &temp)
	logrus.Info(err)
	err = json.Unmarshal([]byte(s2), &temp22)
	logrus.Info(err)
	var diff DiffReporter
	ok := cmp.Equal(&temp, &temp22, cmp.Reporter(&diff))
	fmt.Println(ok)
	var msg []string
	for _, item := range diff.String() {
		msg = append(msg, item)
	}
	return msg
}

type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *DiffReporter) String() []string {
	var msg []string
	for _, item := range r.diffs {
		msg = append(msg, item)
	}
	return msg
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		//fmt.Println("======================")
		//fmt.Printf("pop之前:%v\n", r.path)
		//fmt.Printf("结构:%v, %v\n", vx, vy)
		//fmt.Printf("path的长度:%v\n", len(r.path))
		r.PopStep()
		//vx2, vy2 := r.path.Last().Values()
		//fmt.Printf("一次pop之后:%v\n", r.path)
		//fmt.Printf("结构:%v, %v\n", vx2, vy2)
		//fmt.Printf("path的长度:%v\n", len(r.path))
		r.PopStep()
		//vx3, vy3 := r.path.Last().Values()
		//fmt.Printf("两次pop之后:%v\n", r.path)
		//fmt.Printf("结构:%v, %v\n", vx3, vy3)
		//fmt.Printf("path的长度:%v\n", len(r.path))
		//fmt.Println("======================")
		x, _ := r.path.Last().Values()
		//fmt.Println(x)
		v := x.FieldByName("Trans").String()
		//fmt.Println(v)
		r.path = r.path[:len(r.path)+1]
		r.PushStep(r.path.Index(0))
		r.diffs = append(r.diffs, fmt.Sprintf("【工单模板填充内容中】%v 从 %v 变更为 %v", v, vx, vy))
	}
	return
}
