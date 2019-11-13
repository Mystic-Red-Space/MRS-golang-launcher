package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"strconv"
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
	idInput.SetEchoMode(widgets.QLineEdit__Normal)
	widget.Layout().AddWidget(idInput)

	pwInput := widgets.NewQLineEdit(widget)
	pwInput.SetPlaceholderText("Password")
	pwInput.SetEchoMode(widgets.QLineEdit__Password)
	widget.Layout().AddWidget(pwInput)

	loginButton := widgets.NewQPushButton(widget)
	loginButton.ConnectPressed(func() {
		authRes, err := mclogin(idInput.Text(), pwInput.Text())
		dialog := widgets.NewQDialog(widget, 0)
		message := widgets.NewQLabel(dialog, 0)
		if err == nil {
			message.SetText(authRes.SelectedProfile.Name + "\n" + authRes.User.ID)
		} else {
			message.SetText(strconv.Itoa(err.StatusCode) + " " + err.Error + "\n" + err.ErrorMessage)
		}
		dialog.Layout().AddWidget(message)
		dialog.Show()
	})

	window.Show()
	app.Exec()

}
