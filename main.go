package main

import (
	"fmt"
	"image/png"
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
)

func main() {
	gui()
}

func gui() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetFullScreen(true)
	x := widget.NewButton("My button", func() {
	})
	x.Alignment = widget.ButtonAlignTrailing

	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewHBox(
				widget.NewLabel("Fuck this shit Fyne!"),
				x,
			),
		),
	)

	w.ShowAndRun()
}

func listener() {
	kind := map[uint8]string{3: "KeyDown", 4: "KeyHold", 5: "KeyUp", 10: "MouseDrag"}
	when := []uint8{hook.KeyDown, hook.KeyHold, hook.KeyUp, hook.MouseDrag}

	fmt.Println("Press any key")
	for _, w := range when {
		hook.Register(w, nil, func(e hook.Event) {
			char := hook.RawcodetoKeychar(e.Rawcode)
			fmt.Printf("Kind: %7v, Rawcode: %3v, Keychar: %5v, Keycode: %5v, Char: %v\n", kind[e.Kind], e.Rawcode, e.Keychar, e.Keycode, char)
		})
	}

	s := hook.Start()
	<-hook.Process(s)
}

func ss() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}
func oddNaturalNumberGenerator(x int) []int {
	//a for loop with x as maximum value
	//for each value of i in the loop, increment by 2, if value is odd add to the list.

	oddnumber := 1
	arrayOddNums := []int{1}

	for i := 1; i < x; i++ {
		oddnumber += 2
		if oddnumber%2 != 0 {
			//add oddnumber to list
			arrayOddNums = append(arrayOddNums, oddnumber)
		}
	}
	return arrayOddNums
}

func pointers() {
	helloWorld := "helloworld"
	var pointerToHelloWorld *string
	pointerToHelloWorld = &helloWorld
	fmt.Println("Pointer to helloWorld")
	fmt.Println(pointerToHelloWorld)
}
