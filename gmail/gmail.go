package main

import (
	"github.com/webview/webview"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Google web mail")
	w.SetSize(2560, 1440, webview.HintNone)
	w.Navigate("https://gmail.com")
	w.Run()

}
