package main

import (
	"changeme/utils"
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Printf("name: %v\n", name)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 获取md文档的全部内容
func (a *App) GetMdContent(path string) string {
	return utils.GetFileContent(path)
}

func (a *App) UploadImgByPicgo(name string) string {
	os.Chdir("/Users/pixiao/GTools")
	b, err := exec.Command("sh", "-c", "./picgo.sh").Output()
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
	// resultStr := string(b)
	// fmt.Printf("resultStr: %v\n", resultStr)
	// s := strings.Trim(resultStr, " ")
	// isContain := strings.ContainsAny(s, "https://")
	// if isContain {
	// 	s2 := strings.Split(s, "[PicGo SUCCESS]:")
	// 	fmt.Printf("s: %v\n", s2[1])
	// 	return s2[1]
	// } else {
	// 	return resultStr
	// }
}

func (a *App) OpenUrl(url string) {
	fmt.Printf("url: %v\n", url)
	runtime.BrowserOpenURL(a.ctx, "http://www.baidu.com")
}

func (a *App) GetWd() string {
	os.Chdir("/Users/pixiao/GTools")
	dir, _ := os.Getwd()
	return dir
}

// 添加md文件夹列表
func (a *App) AddDirPath(path string) string {
	return utils.AddDirPathToJdb(path)
}

// 获取md文件夹列表数据
func (a *App) GetDirList() string {
	return utils.GetPathData()
}

// 保存md文件内容
func (a *App) SaveMdContent(path string, content string) string {
	return utils.SaveFileContent(path, content)
}

// 删除md文件夹列表
func (a *App) DelMdDir(path string) string {
	return utils.DelDirPathFromJdb(path)
}
