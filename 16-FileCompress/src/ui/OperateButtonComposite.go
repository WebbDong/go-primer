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

func CreateOperateButtonComposite(onStartDecompressButtonClick func(), onStartCompressButtonClick func()) (*OperateButtonComposite, Composite) {
	obComposite := OperateButtonComposite{}
	composite := Composite{
		AssignTo: &obComposite.Composite,
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				AssignTo: &obComposite.startDecompressBtn,
				Text:     "开始解压",
				OnClicked: func() {
					onStartDecompressButtonClick()
				},
			},
			PushButton{
				AssignTo: &obComposite.startCompressBtn,
				Text:     "开始压缩",
				OnClicked: func() {
					onStartCompressButtonClick()
				},
			},
		},
	}
	return &obComposite, composite
}
