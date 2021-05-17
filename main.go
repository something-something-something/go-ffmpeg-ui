package main

import (
	"fmt"
	"os/exec"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

)

func main() {
	ffmpegLocation:=binding.NewString()
	ffmpegLocation.Set("")

	importFileName := binding.NewString()
	importFileName.Set("")
	exportFileName := binding.NewString()
	exportFileName.Set("")
	ffmpegStatus := binding.NewString()
	ffmpegStatus.Set("Not Run")
	fmt.Println("Hi")
	a := app.NewWithID("com.github.something-something-something.go-ffmpeg-ui")
	//a.SendNotification(fyne.NewNotification("ffmpeg ui wraper","Starting"))
	w := a.NewWindow("ffmpeg")
	c := container.New(layout.NewGridLayout(2))

	c.Add(	widget.NewButton("Click to select ffmpeg Executable", func(){
		dialog.ShowFileOpen(func(u fyne.URIReadCloser,err error){
			if err == nil && u != nil {
				ffmpegLocation.Set(u.URI().Path())	
			}
		},w)
	}))
	c.Add(widget.NewLabelWithData(ffmpegLocation))
	c.Add(
		widget.NewButton("import", func() {
		
			importFile(func(u fyne.URIReadCloser, err error) {

				if err == nil && u != nil {
					importFileName.Set(u.URI().Path())
				}

			}, w)

		}))
		c.Add(widget.NewLabelWithData(importFileName))
		
		c.Add(widget.NewButton("export", func() {
			exportFile(func(u fyne.URIWriteCloser, err error) {

				if err == nil && u != nil {
					exportFileName.Set(u.URI().Path())
				}

			}, w)

		}))
		c.Add(widget.NewLabelWithData(exportFileName))
		
		c.Add(widget.NewButton("Run conversion", func() {
			fmt.Println("A")
			impath,imerr:=importFileName.Get()
			outpath,outerr:=exportFileName.Get()
			ffmpegLoc,ffmpegLocErr:=ffmpegLocation.Get()
			
			if imerr==nil && outerr==nil && ffmpegLocErr==nil && ffmpegLoc!="" {
				fmt.Println("B")
				cmd:=exec.Command(ffmpegLoc,"-y","-i",impath,outpath)
				ffmpegStatus.Set("Running")
				fmt.Println("C")
				ffmpegText,err:=cmd.CombinedOutput();
				fmt.Println("D")
				fmt.Printf("%s\n", ffmpegText)
				//cmd.Run()
				if err!=nil{
					fmt.Println("ERROR")
					fmt.Println(err)
				}
				ffmpegStatus.Set("Finnished")
				a.SendNotification(fyne.NewNotification("ffmpeg ui wraper","Finished converting"))
			} else {
				fmt.Println("X")
			}
			fmt.Println("Z")
		}))
		c.Add(widget.NewLabelWithData(ffmpegStatus))
	w.SetContent(c)

		
	w.Resize(fyne.Size{
		Width: 800,
		Height: 800,
	})
	//w.SetFixedSize(true)
	//w.CenterOnScreen()
	w.ShowAndRun()
	
	
}
func importFile(f func(fyne.URIReadCloser, error), win fyne.Window) {

	dialog.ShowFileOpen(f, win)

}
func exportFile(f func(fyne.URIWriteCloser, error), win fyne.Window) {

	dialog.ShowFileSave(f, win)

}
