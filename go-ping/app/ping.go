package app

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

type Chann struct {
	searchRequest chan int
	workDone      chan int
}

var all_c = Chann{
	searchRequest: make(chan int),
	workDone:      make(chan int),
}

var (
	icmp        ICMP
	laddr       = net.IPAddr{IP: net.ParseIP("ip")}
	num         int
	timeout     int64
	size        int
	stop        bool
	success_row int16
	failed_row  int16
	file        *excelize.File
	workerDone  int = 2
	lock        sync.Mutex
	wg          sync.WaitGroup
	limitwork   int
)

var rows [][]string

func PingInit(limitwork int) {
	ParseArgs()
	rows = read_excel()
	success_row = 1
	failed_row = 1
	var rows_num = len(rows)
	go if_ping_ok(1, 1)
	wait(rows_num)
}

func wait(rows_num int) {
	for {
		select {
		case row_num, ok := <-all_c.searchRequest:
			if !ok {
				fmt.Println(ok)
			}
			row_num++
			workerDone++
			if row_num < rows_num {
				go if_ping_ok(row_num, workerDone)
			}
		case workerDone, ok := <-all_c.workDone:
			for workerDone == limitwork {
				time.Sleep(time.Second)
			}
			if !ok {
				fmt.Println(ok)
			}
			fmt.Println(workerDone)
			if workerDone == rows_num {
				file.SetActiveSheet(1)
				if err := file.SaveAs("./host.xlsx"); err != nil {
					fmt.Println(err)
				}
				fmt.Println("已全部完成！！！")
				os.Exit(1)
			}
		default:
			time.Sleep(time.Millisecond * 1)
		}
	}
}

//逐行判断IP地址是否可达。
func if_ping_ok(row_num, workerDone int) {
	desIp := rows[row_num]
	all_c.searchRequest <- row_num
	conn, err := net.DialTimeout("ip4:icmp", desIp[0], time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	//icmp头部填充
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 1
	icmp.SequenceNum = 1

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp) // 以大端模式写入
	data := make([]byte, size)
	buffer.Write(data)
	data = buffer.Bytes()

	var SuccessTimes int // 成功次数
	var FailTimes int    // 失败次数
	var minTime int = int(math.MaxInt32)
	var maxTime int
	var totalTime int

	for i := 0; i < num; i++ {
		icmp.SequenceNum = uint16(1)
		// 检验和设为0
		data[2] = byte(0)
		data[3] = byte(0)

		data[6] = byte(icmp.SequenceNum >> 8)
		data[7] = byte(icmp.SequenceNum)
		icmp.Checksum = CheckSum(data)
		data[2] = byte(icmp.Checksum >> 8)
		data[3] = byte(icmp.Checksum)

		// 开始时间
		t1 := time.Now()
		conn.SetDeadline(t1.Add(time.Duration(time.Duration(timeout) * time.Millisecond)))
		//n, err := conn.Write(data)
		_, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 65535)
		//n, err = conn.Read(buf)
		_, err = conn.Read(buf)
		if err != nil {
			FailTimes++
			continue
		}
		et := int(time.Since(t1) / 1000000)
		if minTime > et {
			minTime = et
		}
		if maxTime < et {
			maxTime = et
		}
		totalTime += et
		SuccessTimes++
		time.Sleep(1 * time.Second)
	}
	lock.Lock()
	fmt.Printf("[%s] 的 Ping 统计信息: 数据包: 已发送 = %d，已接收 = %d，丢包率 = %.2f%%,\n", desIp[0], SuccessTimes+FailTimes, SuccessTimes, float64(FailTimes*100)/float64(SuccessTimes+FailTimes))
	success_num := int16(FailTimes*100) / int16(SuccessTimes+FailTimes)
	if success_num == 0 {
		success_row++
		r := fmt.Sprintf("A%d", success_row)
		file.SetCellValue("Ping_Success", r, desIp[0])
	} else {
		failed_row++
		r := fmt.Sprintf("A%d", failed_row)
		file.SetCellValue("Ping_Failed", r, desIp[0])
	}
	all_c.workDone <- workerDone
	lock.Unlock()
}

func CheckSum(data []byte) uint16 {
	var sum uint32
	var length = len(data)
	var index int
	for length > 1 { // 溢出部分直接去除
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length == 1 {
		sum += uint32(data[index])
	}
	sum = uint32(sum>>16) + uint32(sum)
	sum = uint32(sum>>16) + uint32(sum)
	return uint16(^sum)
}

func ParseArgs() {
	flag.Int64Var(&timeout, "w", 1500, "等待每次回复的超时时间(毫秒)")
	flag.IntVar(&num, "n", 4, "要发送的请求数")
	flag.IntVar(&size, "l", 32, "要发送缓冲区大小")
	flag.BoolVar(&stop, "t", false, "Ping 指定的主机，直到停止")
	flag.Parse()
}

//读取excel文件，获取文件中的所有信息
func read_excel() (rows [][]string) {
	var err error
	file, err = excelize.OpenFile("host.xlsx")
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		return
	}
	fmt.Printf("    %v\n", file.GetSheetMap())
	var num int
	fmt.Print("请输入需要使用的sheet编号(1/2/3)：")
	fmt.Scanln(&num)
	rows = file.GetRows(file.GetSheetName(num))
	file.NewSheet("Ping_Success")
	file.NewSheet("Ping_Failed")
	file.SetCellValue("Ping_Failed", "A1", "failed_ip")
	file.SetCellValue("Ping_Success", "A1", "success_ip")
	return rows
}
