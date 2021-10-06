# Goroutine 池

## 一、worker pool（goroutine池）

### 1、本质上是生产者消费者模型可以有效控制goroutine数量，防止暴涨
<br>需求：计算一个数字的各个位数之和，例如数字123，
<br>结果：1+2+3=6

<p>随机生成数字进行计算控制台输出结果如下：<p>

```
job id:16006 randnum:1988095020245153943 result:78
job id:16007 randnum:1580759098055997200 result:89
job id:16008 randnum:231869647607668553 result:92
job id:16009 randnum:1798521333646810072 result:76
job id:16010 randnum:2073844353518139720 result:75
job id:16011 randnum:1622183438063966468 result:86
job id:16012 randnum:5840424132984009251 result:71
job id:16013 randnum:4086304361953169998 result:94
job id:16014 randnum:9223148630620769039 result:80
```

### 2、使用方法
```
go run main.go
```