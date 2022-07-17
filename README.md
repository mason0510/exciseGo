#Requirements
go1.17.8
redis 7.0.0
mac 11.5.2
grom 1.23.8


#Userage
- go run main.go
- http://localhost:8080/swagger/index.html



#Main function
## main tables
[]: # Language: go
[]: # Path: main.go
1.问题表
2.用户表
3.分类表
4.提交表
5.问题分类表

##Basic function
*用户模块
** 密码登录
** 注册
** 修改密码
** 查看个人信息


*问题管理模块
** 新建问题  关联分类 测试用例
** 查看问题
** 查看问题详情
** 提交问题
** 查看提交列表
** 查看提交详情


*判题模块
**提交记录
**代码提交与判断


*排名模块
**排名的列表情况

#Directory structure
- define
- docs
- models
- routers
- services
- helpers  
- test
- tools
- main.go

#Problem instructions
不包含前端部分

# 许可
License：MIT




