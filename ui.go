package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var W fyne.Window
var MD5Checked = true
var SHA1Checked = true
var SHA256Checked = true

type EntryWithEnterKeyEvent struct{ widget.Entry }

var FileOpenButton *widget.Button
var FileEntry *EntryWithEnterKeyEvent
var FileHashButton *widget.Button
var MD5Check *widget.Check
var SHA1Check *widget.Check
var SHA256Check *widget.Check
var MD5ProgressBar *widget.ProgressBarInfinite
var SHA1ProgressBar *widget.ProgressBarInfinite
var SHA256ProgressBar *widget.ProgressBarInfinite
var MD5Hash *widget.RichText
var SHA1Hash *widget.RichText
var SHA256Hash *widget.RichText
var CopyMD5Button *widget.Button
var CopySHA1Button *widget.Button
var CopySHA256Button *widget.Button

func (entry *EntryWithEnterKeyEvent) KeyDown(key *fyne.KeyEvent) {
	if fyne.KeyReturn == key.Name {
		HashFile()
	}
}

func InitUI() {
	W = App.NewWindow("File Hasher")

	spacer := canvas.NewRectangle(color.Transparent)
	spacer.SetMinSize(fyne.NewSize(1, 1))

	FileEntry = &EntryWithEnterKeyEvent{}
	FileEntry.ExtendBaseWidget(FileEntry)
	FileEntry.SetPlaceHolder("Enter File Path")
	FileOpenButton = &widget.Button{Text: "Open", OnTapped: OpenFile,
		Icon: theme.FileIcon()}
	FileHashButton = &widget.Button{Text: "Hash", OnTapped: HashFile,
		Icon: theme.MailSendIcon()}
	fileHbox := container.NewBorder(
		nil, nil,
		container.NewHBox(spacer, FileOpenButton, spacer),
		container.NewHBox(spacer, FileHashButton, spacer),
		FileEntry,
	)

	MD5Check = &widget.Check{
		Text:    "MD5",
		Checked: MD5Checked,
		OnChanged: func(b bool) {
			MD5Checked = b
		},
	}
	SHA1Check = &widget.Check{
		Text:    "SHA1",
		Checked: SHA1Checked,
		OnChanged: func(b bool) {
			SHA1Checked = b
		},
	}
	SHA256Check = &widget.Check{
		Text:    "SHA256",
		Checked: SHA256Checked,
		OnChanged: func(b bool) {
			SHA256Checked = b
		},
	}
	MD5ProgressBar = widget.NewProgressBarInfinite()
	MD5ProgressBar.Hide()
	SHA1ProgressBar = widget.NewProgressBarInfinite()
	SHA1ProgressBar.Hide()
	SHA256ProgressBar = widget.NewProgressBarInfinite()
	SHA256ProgressBar.Hide()
	MD5Hash = widget.NewRichText()
	SHA1Hash = widget.NewRichText()
	SHA256Hash = widget.NewRichText()
	CopyMD5Button = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		W.Clipboard().SetContent(MD5Hash.String())
	})
	CopyMD5Button.Hide()
	CopySHA1Button = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		W.Clipboard().SetContent(SHA1Hash.String())
	})
	CopySHA1Button.Hide()
	CopySHA256Button = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		W.Clipboard().SetContent(SHA256Hash.String())
	})
	CopySHA256Button.Hide()
	hashVbox := container.NewVBox(
		container.NewBorder(nil, nil,
			container.NewGridWrap(fyne.Size{
				Width:  SHA256Check.MinSize().Width,
				Height: MD5Check.MinSize().Height,
			}, MD5Check),
			container.NewHBox(CopyMD5Button, spacer),
			MD5ProgressBar, MD5Hash),
		container.NewBorder(nil, nil,
			container.NewGridWrap(fyne.Size{
				Width:  SHA256Check.MinSize().Width,
				Height: SHA1Check.MinSize().Height,
			}, SHA1Check),
			container.NewHBox(CopySHA1Button, spacer),
			SHA1ProgressBar, SHA1Hash),
		container.NewBorder(nil, nil,
			container.NewGridWrap(fyne.Size{
				Width:  SHA256Check.MinSize().Width,
				Height: SHA256Check.MinSize().Height,
			}, SHA256Check),
			container.NewHBox(CopySHA256Button, spacer),
			SHA256ProgressBar, SHA256Hash))

	W.SetContent(container.NewVBox(
		fileHbox,
		spacer,
		hashVbox,
	))
	W.Resize(fyne.NewSize(710, 170))
	W.SetFixedSize(true)
	W.Show()
}
