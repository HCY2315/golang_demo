package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// File is an interface to access the file part of a multipart message.
//文件是访问多部分消息的文件部分的接口。
// Its contents may be either stored in memory or on disk.
//它的内容可以存储在内存中，也可以存储在磁盘上。
// If stored on disk, the File's underlying concrete type will be an *os.File.
//如果存储在磁盘上，文件的底层具体类型将是*操作系统文件.

//获取文件大小
func GetSize(f multipart.File) (int, error) {
	// ReadAll reads from r until an error or EOF and returns the data it read.
	//ReadAll从r读取，直到出现错误或EOF，并返回读取的数据。
	// A successful call returns err == nil, not err == EOF. Because ReadAll is
	//成功的调用返回err==nil，而不是err==EOF。因为ReadAll是
	// defined to read from src until EOF, it does not treat an EOF from Read
	//定义为从src读取直到EOF，它不处理从read读取的EOF
	// as an error to be reported.
	//作为要报告的错误。
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

//获取文件后缀名
func GetExt(fileName string) string {
	//Ext返回path使用的文件扩展名。
	//扩展名是从最后一个点开始的后缀在路径的最后斜杠分隔元素；
	//如果没有点，则为空。
	return path.Ext(fileName)
}

//检查文件或者目录是否存在
func CheckNotExist(src string) bool {
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	//Stat返回描述命名文件的FileInfo。
	//如果有错误，它将是*PathError类型。
	_, err := os.Stat(src)
	//IsNotExist返回一个布尔值，指示用户是否知道错误。
	//报告文件或目录不存在。满足于ErrNotExist以及一些syscall错误。
	return os.IsNotExist(err)
}

//检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

//判断是否能够创建文件
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

//创建文件夹
func MkDir(src string) error {
	// ModePerm FileMode = 0777 // Unix permission bits
	//MkdirAll创建一个名为path的目录，以及任何必要的父母，并返回零，否则返回错误。
	//perm（umask之前）的权限位用于所有MkdirAll创建的目录。
	//如果path已经是一个目录，MkdirAll将不执行任何操作返回nil。
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	//OpenFile是一种通用的开放调用；大多数用户将使用open或者改为创建。
	//它打开具有指定标志的命名文件（仅限等）。
	//如果文件不存在，则O\u CREATE标志通过后，它将使用模式perm创建（在umask之前）。
	//如果成功，返回文件上的方法可以用于I/O。如果有错误，它将是*PathError类型。
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	//Getwd返回与当前目录。如果当前目录可以通过多条路径到达（由于符号链接），Getwd可以返回其中任何一个。
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	//os.O_APPEND 可追加内容
	//os.O_CREATE 创建文件，如果文件不存在
	//os.O_RDWR 可读可写
	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
