package demo

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"math/rand"
	"sync"

	"time"
)

func GetPerm() []int {
	rand.Seed(time.Now().UnixNano())
	address := []string{"1.0.0.1:50", "12.12.3.5:211", "191.23.169.58:55", "191.23.169.58:55", "191.23.169.58:55", "191.23.169.58:55"}
	perm := rand.Perm(len(address)) //  [1 4 5 2 0 3] 获取到伪随机数队列切片
	return perm
}

func ForFiles(indexFileDir string) error {
	// 遍历目录
	files, err := ioutil.ReadDir(indexFileDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println("IsDir=", file.IsDir())
		fmt.Println("Size=", file.Size())
		fmt.Println("Name=", file.Name())
		fmt.Println("Mode=", file.Mode())
		fmt.Println("ModTime=", file.ModTime())
		fmt.Println("Sys=", file.Sys())
		fmt.Println("-----------------------------")
	}
	return nil
}

func ForSelect(ch chan int8) bool {
	for {
		select {
		case ch <- int8(0):
			return true
		default:
			return false
		}
	}

}

func FixBug() {
	rates := []int32{1, 2, 3, 4, 5, 6}
	for star, rate := range rates {
		if star+1 < 1 {
			panic("")
		}

		fmt.Println(star, rate)
	}
}

func IoWriter()  {
	// os.Stderr =
}


type ErrGroup struct {
	group errgroup.Group
}

func NewErrGroup() ErrGroup {
	return ErrGroup{group: errgroup.Group{}}
}

func (e *ErrGroup) Go(f func() error) {
	e.group.Go(func() (err error) {
		defer func() {
			if x := recover(); x != nil {
				err = fmt.Errorf("recover:%+v", x)
			}
		}()
		return f()
	})
}

func (e *ErrGroup) Wait() error {
	return e.group.Wait()
}

func EGroup()  {
	go func() {
		g := NewErrGroup()
		g.Go(func() error {
			panic("xxxxxxxxxxxxxxxxxxxxxx")
			return nil
		})

		if err := g.Wait(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("主程正常...")
	time.Sleep(time.Second * 5)
	fmt.Println("安全退出...")
}

type User struct {
	User string
}
type SS struct {
	List []*User
}
func SyncWait()  {
	wg := sync.WaitGroup{}
	fmt.Println("start")
	wg.Wait()
	fmt.Println("end")

	ss := SS{}
	users := make([]*User,0)
	users = append(users, &User{
		User: "张三",
	})

	ss.List = append(users,ss.List...)
	fmt.Println(ss)
}