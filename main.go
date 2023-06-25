package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	root := tview.NewPages()
	root.SetBorder(true).
		SetTitle("ROO Packet Decoder").
		SetTitleAlign(tview.AlignCenter)
	packetPage := NewFlex(false)
	packetPage.SetDirection(tview.FlexRow)

	view := NewTextView(app)
	view.SetText("viewer")
	data := NewFlex(false)
	sentPackets := NewFlex(true)
	receivedPackets := NewFlex(true)

	data.AddItem(sentPackets, 0, 1, false)
	data.AddItem(receivedPackets, 0, 1, false)
	packetPage.AddItem(data, 0, 1, false)
	packetPage.AddItem(view, 5, 1, false)

	receivedPacketsText := NewTextView(app)
	receivedPackets.AddItem(receivedPacketsText, 0, 2, false)

	sentPacketsText := NewTextView(app)
	sentPackets.AddItem(sentPacketsText, 0, 1, false)

	go func() {
		for i := 1; i <= 1024; i++ {
			<-time.Tick(500 * time.Millisecond)
			fmt.Fprintln(receivedPacketsText, "ping!", i)

			if i%10 == 0 {
				fmt.Fprintln(sentPacketsText, "pong!", i)
			}
		}
	}()

	root.AddPage("packet", packetPage, true, true)
	_ = app.SetRoot(root, true).EnableMouse(true).Run()
}

func NewTextView(app *tview.Application) *tview.TextView {
	tv := tview.NewTextView()
	tv.
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	return tv
}

func NewFlex(border bool) *tview.Flex {
	f := tview.NewFlex()
	f.SetBorder(border)
	return f
}
