package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// 实现运行脚本脚本，若是失败，将下载
func main() {

	// url := "C:/Users/Administrator/Desktop/main.exe"
	// 门店的桌面路径一般是上面的
	url := "C:/Users/ZHY/Desktop/main.exe"
	urlpath := "http://localhost:8200/pub_upload/2021-08-17/cdlea9vrq41ct4lpvr.exe"

	cmd := exec.Command(url)
	err := cmd.Run()

	if err != nil {
		resp, err := http.Get(urlpath)
		if err != nil {

			fmt.Println("请求失败")
		}
		if resp != nil {

			defer resp.Body.Close()
			fmt.Println("获得body")
			fmt.Println(resp.Body)

			if err != nil {

				fmt.Println("解析失败")
			}
			// mainrul, err := os.MkdirAll(url, 0777)
			mainrul, err := os.Create(url)

			if err != nil {
				fmt.Println(err)
				fmt.Println("创建失败")
			}

			io.Copy(mainrul, resp.Body)
			mainrul.Close()

		} else {
			fmt.Println("body 是空的")
		}

		cmd1 := exec.Command(url)
		err1 := cmd1.Run()

		if err1 != nil {

			fmt.Println("第二次运行失败")
		}

	} else {
		fmt.Println("启动已经完成")
	}

}
