## 运行方式

```
service端：
    go run service.go 
```
## 访问方式
1. echo -e '{"method":"chaoyue.Hello","params":["hello2222"],"id":3}' | nc localhost 1234

2. client 访问
   go run client.go