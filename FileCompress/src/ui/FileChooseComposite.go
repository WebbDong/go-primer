package ui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type FileChooseComposite struct {
	*walk.Composite
	mainWindow *walk.MainWindow
	// 解压文件路径文本框
	decompressPathLineEdit *walk.LineEdit
	// 解压后文件保存路径文本框
	decompressSavePathLineEdit *walk.LineEdit
	// 压缩文件路径文本框
	compressPathLineEdit *walk.LineEdit
	// 压缩后文件保存路径文本框
	compressSavePathLineEdit *walk.LineEdit
	// 选择解压文件路径按钮
	selectDecompressPathBtn *walk.PushButton
	// 选择解压后文件保存路径按钮
	selectDecompressSavePathBtn *walk.PushButton
	// 选择压缩文件路径按钮
	selectCompressPathBtn *walk.PushButton
	// 选择压缩后文件保存路径按钮
	selectCompressSavePathBtn *walk.PushButton
	// 状态标签
	statusLabel *walk.Label

	decompressPath     string
	decompressSavePath string
	compressPath       string
	compressSavePath   string
}

// 文件选择区域
func CreateFileChooseComposite() (*FileChooseComposite, Composite) {
	fcComposite := FileChooseComposite{}
	// LineEdit是文本框
	composite := Composite{
		AssignTo: &fcComposite.Composite,
		Layout:   Grid{Columns: 2, Spacing: 10},
		Children: []Widget{
			LineEdit{
				AssignTo: &fcComposite.decompressPathLineEdit,
				Enabled:  false,
			},
			PushButton{
				AssignTo: &fcComposite.selectDecompressPathBtn,
				Text:     "选择解压文件",
				OnClicked: func() {
					accepted, path := openFileDialog(fcComposite.mainWindow, "选择解压文件")
					if accepted {
						fcComposite.decompressPathLineEdit.SetText(path)
						fcComposite.decompressPath = path
					}
				},
			},
			LineEdit{
				AssignTo: &fcComposite.decompressSavePathLineEdit,
				OnEditingFinished: func() {
					fcComposite.decompressSavePath = fcComposite.decompressSavePathLineEdit.Text()
				},
			},
			PushButton{
				AssignTo: &fcComposite.selectDecompressSavePathBtn,
				Text:     "选择解压保存文件夹",
				OnClicked: func() {
					accepted, path := openDirDialog(fcComposite.mainWindow, "选择解压保存文件夹")
					if accepted {
						fcComposite.decompressSavePathLineEdit.SetText(path)
						fcComposite.decompressSavePath = path
					}
				},
			},
			LineEdit{
				AssignTo: &fcComposite.compressPathLineEdit,
				Enabled:  false,
			},
			PushButton{
				AssignTo: &fcComposite.selectCompressPathBtn,
				Text:     "选择压缩文件",
				OnClicked: func() {
					accepted, path := openFileDialog(fcComposite.mainWindow, "选择压缩文件")
					if accepted {
						fcComposite.compressPathLineEdit.SetText(path)
						fcComposite.compressPath = path
					}
				},
			},
			LineEdit{
				AssignTo: &fcComposite.compressSavePathLineEdit,
				OnEditingFinished: func() {
					fcComposite.compressSavePath = fcComposite.compressSavePathLineEdit.Text()
				},
			},
			PushButton{
				AssignTo: &fcComposite.selectCompressSavePathBtn,
				Text:     "选择压缩保存文件夹",
				OnClicked: func() {
					accepted, path := openDirDialog(fcComposite.mainWindow, "选择压缩保存文件夹")
					if accepted {
						fcComposite.compressSavePathLineEdit.SetText(path)
						fcComposite.compressSavePath = path
					}
				},
			},
			Label{
				AssignTo: &fcComposite.statusLabel,
				Text:     "压缩成功",
			},
		},
	}
	return &fcComposite, composite
}
