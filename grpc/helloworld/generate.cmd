@echo off

goto start
这块是注释内容
@rem     目标文件夹(如果不指定默认是当前文件夹)           指定生成在那个文件夹下          目标文件
@rem --proto_path=./  目标文件夹(如果不指定默认是当前文件夹)
@rem --go_out=plugins=grpc:./  指定生成pb.go文件在.proto同目录
@rem ./proto/person.proto  要生成的目标文件
:start
protoc  --go_out=plugins=grpc:./  ./proto/person.proto