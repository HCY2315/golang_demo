package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"golang.org/x/crypto/ssh"
)

//定义能够成功登录的主机所对应的行号的列表
var (
	auth_success []int
	success_row  int16
	failed_row   int16
)

//获取ssh主机连接输出文件
func link_log() {
	log_file, err := os.OpenFile("./link_host.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed：", err)
	}
	log.SetOutput(log_file)
}

//获取命令的返回信息
func cmd_log() {
	log_file, err := os.OpenFile("./cmd.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed：", err)
	}
	log.SetOutput(log_file)
}

//判断列表中的主机是否能够正常登录
func ssh_auth() (auth_list []int, rows [][]string) {
	success_row = 0
	failed_row = 0
	file, err := excelize.OpenFile("host.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("    %v\n", file.GetSheetMap())
	var num int
	fmt.Print("请输入需要使用的sheet编号(1/2/3)：")
	fmt.Scanln(&num)
	rows = file.GetRows(file.GetSheetName(num))
	// link_log()
	auth_success = make([]int, 0)
	row_nums := len(rows)
	var row_num int
	fmt.Printf("开始检测：-------------------------------------\n")
	log.Printf("开始检测：-------------------------------------\n")

	//创建新的workbooks
	success_row = 1
	failed_row = 1
	file.NewSheet("Ssh_Success")
	file.NewSheet("Ssh_Failed")
	file.SetCellValue("Ssh_Success", "A1", "ip")
	file.SetCellValue("Ssh_Success", "B1", "port")
	file.SetCellValue("Ssh_Success", "C1", "user")
	file.SetCellValue("Ssh_Success", "D1", "password")
	file.SetCellValue("Ssh_Failed", "A1", "ip")
	file.SetCellValue("Ssh_Failed", "B1", "port")
	file.SetCellValue("Ssh_Failed", "C1", "user")
	file.SetCellValue("Ssh_Failed", "D1", "password")
	for row_num = 1; row_num < row_nums; row_num++ {
		var ip string
		var port string
		var user string
		var auth string
		ip = rows[row_num][0]
		port = rows[row_num][1]
		user = rows[row_num][2]
		auth = rows[row_num][3]
		host := ip + ":" + port
		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(auth),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		_, err := ssh.Dial("tcp", host, config)
		if err != nil {
			fmt.Printf("[%s] 主机能够登录失败：%s\n", ip, err)
			log.Printf("[%s] 主机能够登录失败：%s\n", ip, err)
			failed_row++
			A := fmt.Sprintf("A%d", failed_row)
			B := fmt.Sprintf("B%d", failed_row)
			C := fmt.Sprintf("C%d", failed_row)
			D := fmt.Sprintf("D%d", failed_row)
			file.SetCellValue("Ssh_Failed", A, ip)
			file.SetCellValue("Ssh_Failed", B, port)
			file.SetCellValue("Ssh_Failed", C, user)
			file.SetCellValue("Ssh_Failed", D, auth)
		} else {
			fmt.Printf("[%s] 主机能够成功登录！！！！\n", ip)
			log.Printf("[%s] 主机能够成功登录！！！！\n", ip)
			success_row++
			A := fmt.Sprintf("A%d", success_row)
			B := fmt.Sprintf("B%d", success_row)
			C := fmt.Sprintf("C%d", success_row)
			D := fmt.Sprintf("D%d", success_row)
			file.SetCellValue("Ssh_Success", A, ip)
			file.SetCellValue("Ssh_Success", B, port)
			file.SetCellValue("Ssh_Success", C, user)
			file.SetCellValue("Ssh_Success", D, auth)
			auth_success = append(auth_success, row_num)
		}
	}
	file.SetActiveSheet(1)
	if err := file.SaveAs("./host.xlsx"); err != nil {
		fmt.Println(err)
	}
	return auth_success, rows
}

//向各个主机中发送linux命令
func ssh_shell(rows [][]string, auth_list []int, str1 string, str2 string) {
	cmd_log()
	var auth_num int
	fmt.Printf("命令：[%s]===================================================\n", str1)
	log.Printf("命令：[%s]===================================================\n", str2)
	for auth_num = 0; auth_num < len(auth_list); auth_num++ {
		var ip string
		var port string
		var user string
		var auth string
		ip = rows[auth_list[auth_num]][0]
		port = rows[auth_list[auth_num]][1]
		user = rows[auth_list[auth_num]][2]
		auth = rows[auth_list[auth_num]][3]
		host := ip + ":" + port
		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(auth),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		client, _ := ssh.Dial("tcp", host, config)
		session, _ := client.NewSession()
		defer session.Close()

		//处理传入的命令
		cmd, err := session.CombinedOutput(str1)
		if err != nil {
			fmt.Println("远程执行cmd失败", err)
		}
		fmt.Println()
		fmt.Printf("[%s] %s：\n", ip, str2)
		fmt.Printf("%s\n%s\n", str2, string(cmd))
		log.Printf("：[%s] %s\n%s\n", ip, str2, string(cmd))
	}
}

func main() {
	fmt.Println("-------------------正在测试远程主机的连通性-------------------------")
	auth_list, rows := ssh_auth()
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd1, _ := reader.ReadString('\n')
		cmd2, _ := reader.ReadString('\n')
		fmt.Print("请输入需要的命令：")
		cmd1 = strings.TrimSpace(cmd1)
		fmt.Print("请输入此命令查询的主题：")
		cmd2 = strings.TrimSpace(cmd2)
		ssh_shell(rows, auth_list, cmd1, cmd2)
	}
}
