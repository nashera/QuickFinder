package window

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// func makeButtonTab() fyne.Widget {
// 	disabled := widget.NewButton("Disabled", func() {})
// 	disabled.Disable()

// 	return widget.NewVBox(
// 		widget.NewLabel("Text label"),
// 		widget.NewButton("Text button", func() { fmt.Println("tapped text button") }),
// 		widget.NewButtonWithIcon("With icon", theme.ConfirmIcon(), func() { fmt.Println("tapped icon button") }),
// 		disabled,
// 	)
// }

// func makeInputTab() fyne.Widget {
// 	entry := widget.NewEntry()
// 	entry.SetPlaceHolder("Entry")
// 	entryReadOnly := widget.NewEntry()
// 	entryReadOnly.SetText("Entry (disabled)")
// 	entryReadOnly.Disable()

// 	disabledCheck := widget.NewCheck("Disabled check", func(bool) {})
// 	disabledCheck.Disable()
// 	radio := widget.NewRadio([]string{"Radio Item 1", "Radio Item 2"}, func(s string) { fmt.Println("selected", s) })
// 	radio.Horizontal = true
// 	disabledRadio := widget.NewRadio([]string{"Disabled radio"}, func(string) {})
// 	disabledRadio.Disable()

// 	return widget.NewVBox(
// 		entry,
// 		entryReadOnly,
// 		widget.NewSelect([]string{"Option 1", "Option 2", "Option 3"}, func(s string) { fmt.Println("selected", s) }),
// 		widget.NewCheck("Check", func(on bool) { fmt.Println("checked", on) }),
// 		disabledCheck,
// 		radio,
// 		disabledRadio,
// 		widget.NewSlider(0, 100),
// 	)
// }

// func makeProgressTab() fyne.Widget {
// 	progress := widget.NewProgressBar()
// 	infProgress := widget.NewProgressBarInfinite()

// 	go func() {
// 		num := 0.0
// 		for num < 1.0 {
// 			time.Sleep(100 * time.Millisecond)
// 			progress.SetValue(num)
// 			num += 0.01
// 		}

// 		progress.SetValue(1)
// 	}()

// 	return widget.NewVBox(
// 		widget.NewLabel("Percent"), progress,
// 		widget.NewLabel("Infinite"), infProgress)
// }

// func makeFormTab() fyne.Widget {
// 	name := widget.NewEntry()
// 	name.SetPlaceHolder("John Smith")
// 	email := widget.NewEntry()
// 	email.SetPlaceHolder("test@example.com")
// 	password := widget.NewPasswordEntry()
// 	password.SetPlaceHolder("Password")
// 	largeText := widget.NewMultiLineEntry()

// 	form := &widget.Form{
// 		OnCancel: func() {
// 			fmt.Println("Cancelled")
// 		},
// 		OnSubmit: func() {
// 			fmt.Println("Form submitted")
// 			fmt.Println("Name:", name.Text)
// 			fmt.Println("Email:", email.Text)
// 			fmt.Println("Password:", password.Text)
// 			fmt.Println("Message:", largeText.Text)
// 		},
// 	}
// 	form.Append("Name", name)
// 	form.Append("Email", email)
// 	form.Append("Password", password)
// 	form.Append("Message", largeText)

// 	return form
// }

// func makeScrollTab() fyne.CanvasObject {
// 	logo := canvas.NewImageFromResource(theme.FyneLogo())
// 	logo.SetMinSize(fyne.NewSize(320, 320))
// 	list := widget.NewHBox()
// 	list2 := widget.NewVBox()

// 	for i := 1; i <= 20; i++ {
// 		index := i // capture
// 		list.Append(widget.NewButton(fmt.Sprintf("Button %d", index), func() {
// 			fmt.Println("Tapped", index)
// 		}))
// 		list2.Append(widget.NewButton(fmt.Sprintf("Button %d", index), func() {
// 			fmt.Println("Tapped", index)
// 		}))
// 	}

// 	horiz := widget.NewHScrollContainer(list)
// 	vert := widget.NewVScrollContainer(list2)

// 	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
// 		fyne.NewContainerWithLayout(layout.NewBorderLayout(horiz, nil, nil, nil), horiz, vert),
// 		makeScrollBothTab())
// }

// func makeScrollBothTab() fyne.CanvasObject {
// 	logo := canvas.NewImageFromResource(theme.FyneLogo())
// 	logo.SetMinSize(fyne.NewSize(800, 800))

// 	scroll := widget.NewScrollContainer(logo)
// 	scroll.Resize(fyne.NewSize(400, 400))

// 	return scroll
// }

// // WidgetScreen 搜索界面
// func WidgetScreen() fyne.CanvasObject {
// 	toolbar := widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() { fmt.Println("New") }),
// 		widget.NewToolbarSeparator(),
// 		widget.NewToolbarSpacer(),
// 		widget.NewToolbarAction(theme.ContentCutIcon(), func() { fmt.Println("Cut") }),
// 		widget.NewToolbarAction(theme.ContentCopyIcon(), func() { fmt.Println("Copy") }),
// 		widget.NewToolbarAction(theme.ContentPasteIcon(), func() { fmt.Println("Paste") }),
// 	)
// 	return fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil),
// 		toolbar,
// 		widget.NewTabContainer(
// 			widget.NewTabItem("Buttons", makeButtonTab()),
// 			widget.NewTabItem("Input", makeInputTab()),
// 			widget.NewTabItem("Progress", makeProgressTab()),
// 			widget.NewTabItem("Form", makeFormTab()),
// 			widget.NewTabItem("Scroll", makeScrollTab()),
// 		),
// 	)
// }

// SearchScreen 搜索界面
func SearchScreen() fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry")
	// list := widget.NewHBox()
	list2 := widget.NewVBox()

	for i := 1; i <= 20; i++ {
		index := i // capture
		// list.Append(widget.NewButton(fmt.Sprintf("Button %d", index), func() {
		// 	fmt.Println("Tapped", index)
		// }))
		list2.Append(widget.NewButton(fmt.Sprintf("Button %d", index), func() {
			fmt.Println("Tapped", index)
		}))
	}

	// horiz := widget.NewHScrollContainer(list)
	vert := widget.NewVScrollContainer(list2)
	vert.SetMinSize(fyne.NewSize(400, 400))

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), entry, fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), vert))
}
