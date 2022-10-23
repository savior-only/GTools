package gtools

import (
	"changeme/configs"
	"changeme/util"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) OpenMdSaveFileWindow() *util.Resp {
	option := runtime.SaveDialogOptions{
		DefaultFilename: "new",
		Title:           "保存文件",
		Filters:         [](runtime.FileFilter){configs.MdFilter},
	}
	fpath, err := runtime.SaveFileDialog(a.ctx, option)
	if err != nil {
		a.Log.Error("新markdown路径获取失败")
		return util.Error("新markdown路径获取失败")
	}
	return util.Success(fpath)
}
