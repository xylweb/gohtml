package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/del-xiong/miniblink"
)

// 获取当前执行程序所在的绝对路径
func Asset(f string) string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return fmt.Sprintf("file:///%s", filepath.Join(res, f))
}
func main() {
	//设置调试模式
	miniblink.SetDebugMode(true)
	//初始化miniblink模块
	err := miniblink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}
	// 启动1366x920普通浏览器
	view := miniblink.NewWebView(false, 1366, 920)
	// 启动1366x920透明浏览器(只有web界面会显示)
	//view := miniblink.NewWebView(true, 1366, 920)
	// 加载github
	//view.LoadURL("https://www.baidu.com")

	view.LoadURL(Asset("main.html"))
	// 设置窗体标题(会被web页面标题覆盖)
	view.SetWindowTitle("miniblink window")
	// 移动到屏幕中心位置
	view.MoveToCenter()
	view.Inject("MoveWindow", func(x, y int32, relative bool) (int, error) {
		rectx, _ := view.GetWindowRect()
		if (relative && x+rectx > 1000) || (!relative && x > 1000) {
			return 0, errors.New("x位置不能超过1000")
		}
		view.Move(x, y, relative)
		time.Sleep(time.Second)
		return int(time.Now().Unix()), nil
	})
	// 显示窗口
	view.ShowWindow()
	// 开启调试模式(会调起chrome调试页面)
	//view.ShowDevTools()
	<-make(chan bool)
}
