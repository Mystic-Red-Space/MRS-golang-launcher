package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(640, 320)
	window.SetWindowTitle("MRS Launcher")

	widget := widgets.NewQWidget(window, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	idInput := widgets.NewQLineEdit(widget)
	idInput.SetPlaceholderText("E-mail")
	idInput.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit)
	widget.Layout().AddWidget(idInput)
	window.Show()
	app.Exec()

}
