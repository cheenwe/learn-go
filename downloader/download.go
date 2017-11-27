package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

/*
 *通过传入下载URL和名称来下载命名
 *获取配置文件里面的存放路径
 *然后用http.Get访问下载地址
 *最后用io.Copy拷贝文件到本地
 */

func downloadFile(downUrl, downPath string, filename string) {
	if _, err := os.Stat(downPath); err != nil {
		os.MkdirAll(downPath, 0744)
	}
	res, _ := http.Get(downUrl)
	file, _ := os.Create(downPath + "/" + filename)
	if res.Body != nil {
		defer res.Body.Close()
	}
	io.Copy(file, res.Body)
	fmt.Println("下载完成")
}

func downloadFast(downUrl, downPath string, filename string) {
	err := os.MkdirAll(downPath, 0744)
	if err != nil {
		fmt.Println("文件下载路径创建出错")
	}
	resp, err := http.Get(downUrl)
	if err != nil {
		fmt.Println("下载文件错误")
	}
	defer resp.Body.Close()
	fd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("下载文件出错")
	}
	err = ioutil.WriteFile(filename, fd, 0666)
	if err != nil {
		fmt.Println("写入文件出错")
	}
	fmt.Println("文件下载成功")
}

func main() {
	downUrl := "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2788538611,1610427494&fm=27&gp=0.jpg"
	savePath := "/Users/chenwei/tmp/000001123"
	filename := "123.jpg"

	startAt := time.Now()
	downloadFile(downUrl, savePath, filename)
	down_used := time.Since(startAt)
	fmt.Println("下载用时: ", down_used)

	//	第二次下载同一文件,速度会变快,与方法无太大关系
	startAt1 := time.Now()
	savePath1 := "/Users/chenwei/tmp/000001122"
	downloadFast(downUrl, savePath1, filename)
	//	downloadFile(downUrl, savePath1, filename)
	down_used1 := time.Since(startAt1)
	fmt.Println("下载用时1: ", down_used1)
}
