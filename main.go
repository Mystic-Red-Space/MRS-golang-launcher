package main

import "github.com/asticode/go-astilectron"

func main() {
	// Initialize astilectron
	var a, _ = astilectron.New(astilectron.Options{
		AppName:            "MRS Launcher",
		AppIconDefaultPath: "favicon.png",  // If path is relative, it must be relative to the data directory
		AppIconDarwinPath:  "favicon.icns", // Same here
		BaseDirectoryPath:  "<where you want the provisioner to install the dependencies>",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	// Blocking pattern
	a.Wait()
	var w, _ = a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	})
	w.Create()
}
