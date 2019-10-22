# snoar-test

##在自己服务器上安装go环境
##在服务器上下载sonarqube 和sonar-scanner

go test -coverprofile=coverage.out 命令行查看单元测试覆盖率

#嵌套单元测试命令
go test -json -covermode=atomic -v  -coverprofile convert-sonar-test.out   ./...  
#生成报告以后再sonar 配置中指定这个报告路径 然后
sonar-scanner


