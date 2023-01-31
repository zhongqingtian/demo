package clickhouse

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestAsyncInsert(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "临时测试1",
			args: args{users: []User{
				{Id: 3, Name: "lixxx", Age: 107},
				{Id: 4, Name: "lixxx2", Age: 107}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AsyncInsert(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("AsyncInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAsyncInsertExample(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		err := AsyncInsertExample(Example{Id: uint64(i), Name: "lixxx", Age: 107, Ids: []uint8{1, 2, uint8(i)}, Ctime: time.Now()})
		if err != nil {
			t.Log(err.Error())
		}
	}
	fmt.Println("插入完成 ", time.Since(start).Seconds())
	time.Sleep(time.Minute)
}

func TestBatchInsertExample(t *testing.T) {
	start := time.Now()
	ch := make(chan int, 10)
	var examples []Example
	for i := 99999900; i < 200000200; i++ {
		examples = append(examples, Example{Id: uint64(i), Name: "lixxx争议活得好好的海岛大亨好的好的哈哈哈v", Age: 107, Ids: []uint8{1, 2, uint8(i)}, Ctime: time.Now()})
		if i%100 == 0 {
			ch <- 1
			go func(examples []Example) {
				defer func() {
					<-ch
				}()
				err := BatchInsertExample(examples)
				if err != nil {
					t.Log(err.Error())
				}
			}(examples)
			examples = make([]Example, 0)
		}
	}
	fmt.Println("插入完成 ", time.Since(start).Seconds())
	time.Sleep(time.Minute)
}

func TestBatchInsert(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "临时测试1",
			args: args{users: []User{
				{Id: 5, Name: "luuuu", Age: 105, Ids: []uint8{55, 127, 255}, Ctime: time.Now()},
				{Id: 6, Name: "luuuu2", Age: 109, Ids: []uint8{59, 128, 255}, Ctime: time.Now()}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BatchInsert(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("BatchInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestColumnInsert(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "临时测试2",
			args: args{users: []User{
				{Id: 9, Name: "text1", Age: 109, Ids: []uint8{59, 128, 255}, Ctime: time.Now()},
				{Id: 10, Name: "text2", Age: 119, Ids: []uint8{59, 128, 255}, Ctime: time.Now()}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ColumnInsert(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("ColumnInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTable(); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueryParameter(t *testing.T) {
	tests := []struct {
		name    string
		want    uint64
		wantErr bool
	}{
		{
			name:    "test1",
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryParameter()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryParameter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryRow(t *testing.T) {
	tests := []struct {
		name    string
		want    []User
		wantErr bool
	}{
		{
			name:    "临时测试2",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryRow()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryRow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryWithParameters(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "临时测试2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QueryWithParameters(); (err != nil) != tt.wantErr {
				t.Errorf("QueryWithParameters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	type args struct {
		read User
	}
	tests := []struct {
		name    string
		args    args
		want    []User
		wantErr bool
	}{
		{name: "test",
			args: args{read: User{
				Id: 9, Name: "text1", Age: 109},
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Select(tt.args.read)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
		})
	}
}
