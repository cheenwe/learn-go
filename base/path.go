package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkOrCreatePath(name string) {
	//	filepath.Abs(filepath.Dir(os.Args[0])) //执行文件所在目录
	//	filepath.Walk(root, visit) //显示所有文件夹、子文件夹、文件、子文件
	path := filepath.Dir(name)
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("路径已存在...")
	} else {
		fmt.Println("创建文件夹...")
		err := os.MkdirAll(path, 0744)
		if err != nil {
			fmt.Println("文件路径: %s", err)
		}
	}
}

func checkOrCreateFile(name string) {
	//	dir := filepath.Dir(name)
	//	filename := filepath.Base(name)
	//	ext := filepath.Ext(filename)
	//	prefix := filename[:len(filename)-len(ext)]
	//	files, _ := ioutil.ReadDir(myfolder) //遍历文件夹下目录
	//	checkOrCreatePath(name)

	//	os.Rename(originalPath, newPath) //重命名

	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		_, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

		fmt.Println("创建文件....")
		if err != nil {
			fmt.Println("创建文件失败", err)
		}
	} else {
		fmt.Println("文件存在啦....")
	}
}

func isExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func deleteAllFiles(dir string, files []string) {
	for _, f := range files {
		_ = os.Remove(filepath.Join(dir, filepath.Base(f)))
	}
}

func deleteDirectory(dir string) {
	_ = os.RemoveAll(dir)
}

// dir returns the directory for the current filename.
func dir(name string) string {
	return filepath.Dir(filepath.Base(name))
}

func main() {

	file := "/Users/chenwei/workspace/go/learn/base/1/t12rr1qq112est111/a1.log"
	file1 := "/Users/chenwei/workspace/go/learn/base/1/t12rr1qq112est111/a.log"
	ff := dir(file)
	fmt.Println(ff)

	checkOrCreatePath(file)
	checkOrCreateFile(file)
	checkOrCreateFile(file1)
	result := isExists(file)
	fmt.Println(result)

	//	deleteAllFiles("/Users/chenwei/workspace/go/learn/base/1/t12rr1qq112est111/", []string{"d.log", "a1.log"})

	//	deleteDirectory("/Users/chenwei/workspace/go/learn/base/1")
}
