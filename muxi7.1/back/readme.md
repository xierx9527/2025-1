后端
运行go run demo2.go jwt.go
在apifox /login传入json{
  "username": "admin",
  "password": "admin123"
}得到token
在专有接口header添加参数authrization Bearer {token}
根据接口类型传入不同参数，按id查询，删除
上传接口传入json
所有更改操作同步至本地数据库root
