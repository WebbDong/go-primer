package ui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"service"
)

// 主窗口
type CompressMainWindow struct {
	*walk.MainWindow
	// 左边文件选择区域
	fileChooseComposite *FileChooseComposite
	// 右边操作按钮区域
	operateBtnComposite *OperateButtonComposite

	compressService service.CompressServicer
}

func (c *CompressMainWindow) Show() {
	fileChooseComposite, fcComposite := CreateFileChooseComposite()
	c.fileChooseComposite = fileChooseComposite
	c.compressService = &service.CompressService{}
	operateBtnComposite, obComposite := CreateOperateButtonComposite(
		// 开始解压按钮点击
		func() {
			isSuccess := c.compressService.Decompress(fileChooseComposite.decompressPath, fileChooseComposite.decompressSavePath)
			if isSuccess {
				fileChooseComposite.statusLabel.SetText("解压成功")
			} else {
				fileChooseComposite.statusLabel.SetText("解压失败")
			}
		},
		// 开始压缩按钮点击
		func() {
			isSuccess := c.compressService.Compress(fileChooseComposite.compressPath, fileChooseComposite.compressSavePath)
			if isSuccess {
				fileChooseComposite.statusLabel.SetText("压缩成功")
			} else {
				fileChooseComposite.statusLabel.SetText("压缩成功")
			}
		},
	)
	c.operateBtnComposite = operateBtnComposite

	windowSize := Size{600, 400}
	// Layout: VBox是垂直布局，HBox是水平布局
	err := MainWindow{
		AssignTo: &c.MainWindow,
		Title:    "文件压缩工具",
		Size:     windowSize,
		Layout:   HBox{},
		Children: []Widget{
			fcComposite,
			obComposite,
		},
	}.Create()
	fileChooseComposite.mainWindow = c.MainWindow
	/*
		_, err := MainWindow{
			AssignTo: &c.MainWindow,
			Title:    "文件压缩工具",
			Size:  Size{600, 400},
			Layout:   HBox{},
			Children: []Widget{
				fcComposite,
			},
		}.Run()
	*/
	c.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// 打开文件选择框
func openFileDialog(owner *walk.MainWindow, title string) (accepted bool, path string) {
	dialog := walk.FileDialog{
		Title:  title,
		Filter: "所有文档(*.*)|*.*|文本文档(*.txt)|*.txt|压缩文件(*.zip)|*.zip",
	}
	// accepted返回值：当用户再文件选择框中点击取消打开时，为true，否则为false
	accepted, err := dialog.ShowOpen(owner)
	if err != nil {
		fmt.Println(err)
	}
	path = dialog.FilePath
	return
}

// 打开文夹件选择框
func openDirDialog(owner *walk.MainWindow, title string) (accepted bool, path string) {
	dialog := walk.FileDialog{
		Title: title,
	}
	// accepted返回值：当用户再文件选择框中点击取消打开时，为true，否则为false
	accepted, err := dialog.ShowBrowseFolder(owner)
	if err != nil {
		fmt.Println(err)
	}
	path = dialog.FilePath
	return
}
