package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

var id int

type Result struct {
	job *Job
	// 求和
	sum int
}

func main() {
	// 需要两个通道
	// 1、job通道
	jobChan := make(chan *Job, 128)
	// 2、结果通道
	resultChan := make(chan *Result, 128)
	// 3、创建工作池
	createPool(64, jobChan, resultChan)
	// 4、开个打印的协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	// 5、循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行计算，遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 获取需要计算的随机数
				r_num := job.RandNum
				// 随机数位数相加减
				var sum int
				for r_num != 0 {
					// 求余
					tmp := r_num % 10
					// 求和
					sum += tmp
					// 相除
					r_num /= 10
				}
				// 获取想要的结果
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
