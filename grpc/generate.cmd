@echo off
goto start
:: 就是如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录。
:: 如果没有指定 –I 参数，则在当前目录进行搜索
:start
protoc -I./helloworld/proto --go_out=plugins=grpc:./helloworld/proto/ ./helloworld/proto/hello_world.proto

goto start
这块是注释内容
@rem     目标文件夹(如果不指定默认是当前文件夹)           指定生成在那个文件夹下          目标文件
@rem --proto_path=./  目标文件夹(如果不指定默认是当前文件夹)
@rem --go_out=plugins=grpc:./  指定生成pb.go文件在.proto同目录
@rem ./proto/person.proto  要生成的目标文件
:start
protoc  --go_out=plugins=grpc:./  ./helloworld/proto/person.proto