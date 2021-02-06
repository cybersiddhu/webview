package main

import (
	"github.com/webview/webview"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Dicty stock center")
	w.SetSize(2560, 1440, webview.HintNone)
	w.Navigate("https://dictycr.org/stockcenter")
	w.Run()

}
