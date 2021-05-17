package main

import (
	"fmt"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
	"os/exec"
)

func main() {
	a := app.NewWithID("com.github.something-something-something.go-ffmpeg-ui")
	
	
	ffmpegLocation := binding.NewString()
	ffmpegLocation.Set(a.Preferences().StringWithFallback("ffmpegBin",""))
	ffmpegLocation.AddListener(binding.NewDataListener(func() {
		ffmpegLoc,err:=ffmpegLocation.Get()
		if err==nil{
			a.Preferences().SetString("ffmpegBin",ffmpegLoc)
		}
		
	}))
	importFileName := binding.NewString()
	importFileName.Set("")
	exportFileName := binding.NewString()
	exportFileName.Set("")
	ffmpegStatus := binding.NewString()
	ffmpegStatus.Set("Not Run")
	fmt.Println("Hi")
	
	//a.SendNotification(fyne.NewNotification("ffmpeg ui wraper","Starting"))
	w := a.NewWindow("ffmpeg ui")
	c := container.New(layout.NewGridLayout(2))

	c.Add(widget.NewButton("Click to select ffmpeg Executable", func() {
		loc, err := zenity.SelectFile()
		if err == nil {
			ffmpegLocation.Set(loc)
		}
	}))
	c.Add(widget.NewLabelWithData(ffmpegLocation))
	c.Add(
		widget.NewButton("import", func() {
			imFile, err := zenity.SelectFile()
			if err == nil {
				importFileName.Set(imFile)
			}
		}))
	c.Add(widget.NewLabelWithData(importFileName))

	c.Add(widget.NewButton("export", func() {

		outFile, err := zenity.SelectFileSave(zenity.Filename("ffmpeg.mp4"))
		if err == nil {
			exportFileName.Set(outFile)
		}

	}))
	c.Add(widget.NewLabelWithData(exportFileName))

	c.Add(widget.NewButton("Run conversion", func() {
		fmt.Println("A")
		impath, imerr := importFileName.Get()
		outpath, outerr := exportFileName.Get()
		ffmpegLoc, ffmpegLocErr := ffmpegLocation.Get()
		
		if imerr == nil && outerr == nil && ffmpegLocErr == nil && ffmpegLoc != "" {
			fmt.Println("B")
			cmd := exec.Command(strings.TrimSpace(ffmpegLoc), "-y", "-i", strings.TrimSpace(impath), strings.TrimSpace(outpath))
			ffmpegStatus.Set("Running")
			fmt.Println("C")
			ffmpegText, err := cmd.CombinedOutput()
			fmt.Println("D")
			fmt.Printf("%s\n", ffmpegText)
			//cmd.Run()
			if err != nil {
				fmt.Println("ERROR")
				fmt.Println(err)
			}
			ffmpegStatus.Set("Finnished")
			a.SendNotification(fyne.NewNotification("ffmpeg ui wraper", "Finished converting"))
		} else {
			fmt.Println("X")
		}
		fmt.Println("Z")
	}))
	c.Add(widget.NewLabelWithData(ffmpegStatus))
	w.SetContent(c)

	w.Resize(fyne.Size{
		Width:  800,
		Height: 800,
	})
	//w.SetFixedSize(true)
	//w.CenterOnScreen()
	w.ShowAndRun()

}

// func setDialogSizeAndShow(d dialog.Dialog){
// 	d.Resize(fyne.Size{
// 		Width: 800,
// 		Height: 800,
// 	})
// 	d.Show()
// }

// func importFile(f func(fyne.URIReadCloser, error), win fyne.Window) {
// 	d:=dialog.NewFileOpen(f,win)
// 	setDialogSizeAndShow(d)

// }
// func exportFile(f func(fyne.URIWriteCloser, error), win fyne.Window) {
// 	d:=dialog.NewFileSave(f,win)
// 	d.SetFileName("ffmpeg.mp4")
// 	setDialogSizeAndShow(d)

// }
