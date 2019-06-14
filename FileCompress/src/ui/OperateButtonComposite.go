package ui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type OperateButtonComposite struct {
	*walk.Composite
	startDecompressBtn *walk.PushButton
	startCompressBtn   *walk.PushButton
}

func CreateOperateButtonComposite() (*OperateButtonComposite, Composite) {
	obComposite := OperateButtonComposite{}
	composite := Composite{
		AssignTo: &obComposite.Composite,
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				AssignTo: &obComposite.startDecompressBtn,
				Text:     "开始解压",
			},
			PushButton{
				AssignTo: &obComposite.startCompressBtn,
				Text:     "开始压缩",
			},
		},
	}
	return &obComposite, composite
}
