```
直接运行
go run first_go/src/main/hello_word.go

编译为二进制包
go build first_go/src/main/hello_word.go

直接运行
[root@virtualworkstation study_go]# ll
total 2020
drwxr-xr-x 4 root root      28 Apr  8 19:09 first_go
-rwxr-xr-x 1 root root 2068300 Apr  8 19:13 hello_word
[root@virtualworkstation study_go]# ./hello_word 
Hello World
```
返回程序状态值
```
[root@virtualworkstation study_go]# go run first_go/src/main/hello_word.go
Hello World
[root@virtualworkstation study_go]# echo $?
0
[root@virtualworkstation study_go]# go run first_go/src/main/hello_word.go
Hello World
exit status 1
[root@virtualworkstation study_go]# echo $?
1
```
main函数不支持参数传入，通过os.Args获取命令行参数
```

```

编写测试程序文件：XXX_test.go
定义测试方法：func TestXXX(){}