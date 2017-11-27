package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func downloadFile(downUrl, downPath string, filename string) {
	/*
	 *通过传入下载URL和名称来下载命名
	 *获取配置文件里面的存放路径
	 *然后用http.Get访问下载地址
	 *最后用io.Copy拷贝文件到本地
	 */

	if _, err := os.Stat(downPath); err != nil {
		os.MkdirAll(downPath, 0744)
	} else {
		os.Mkdir(downPath, 0744)
	}

	res, _ := http.Get(downUrl)
	file, _ := os.Create(downPath + "/" + filename)
	if res.Body != nil {
		defer res.Body.Close()
	}
	io.Copy(file, res.Body)
	fmt.Println("下载完成")
}

func main() {
	startAt := time.Now()
	downUrl := "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2788538611,1610427494&fm=27&gp=0.jpg"
	savePath := "/tmp/000001/123"
	filename := "123.jpg"

	downloadFile(downUrl, savePath, filename)

	down_used := time.Since(startAt)
	fmt.Println("下载用时: ", downused)
}
