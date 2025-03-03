package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	nativeDialog "github.com/sqweek/dialog"
	"io"
	"os"
	"sync"
)

func OpenFile() {
	fileBuilder := nativeDialog.File().Title("Open File")
	filename, err := fileBuilder.Load()
	if err != nil {
		if err.Error() != "Cancelled" {
			panic(err)
		}
	} else {
		FileEntry.SetText(filename)
		HashFile()
	}
}

func HashFile() {
	FileEntry.Disable()
	FileOpenButton.Disable()
	FileHashButton.Disable()
	MD5Check.Disable()
	SHA1Check.Disable()
	SHA256Check.Disable()
	defer func() {
		FileEntry.Enable()
		FileOpenButton.Enable()
		FileHashButton.Enable()
		MD5Check.Enable()
		SHA1Check.Enable()
		SHA256Check.Enable()
	}()
	_, err := os.Stat(FileEntry.Text)
	if err != nil {
		nativeDialog.Message(err.Error()).Error()
		return
	}
	MD5Hash.ParseMarkdown("")
	SHA1Hash.ParseMarkdown("")
	SHA256Hash.ParseMarkdown("")
	CopyMD5Button.Hide()
	CopySHA1Button.Hide()
	CopySHA256Button.Hide()
	var wg sync.WaitGroup
	if MD5Checked {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f, err := os.Open(FileEntry.Text)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			MD5ProgressBar.Show()
			md5Hash := md5.New()
			if _, err := io.Copy(md5Hash, f); err != nil {
				panic(err)
			}
			MD5ProgressBar.Hide()
			MD5Hash.ParseMarkdown(fmt.Sprintf("`%x`", md5Hash.Sum(nil)))
			CopyMD5Button.Show()
		}()
	}
	if SHA1Checked {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f, err := os.Open(FileEntry.Text)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			SHA1ProgressBar.Show()
			sha1hash := sha1.New()
			if _, err := io.Copy(sha1hash, f); err != nil {
				panic(err)
			}
			SHA1ProgressBar.Hide()
			SHA1Hash.ParseMarkdown(fmt.Sprintf("`%x`", sha1hash.Sum(nil)))
			CopySHA1Button.Show()
		}()
	}
	if SHA256Checked {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f, err := os.Open(FileEntry.Text)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			SHA256ProgressBar.Show()
			sha256hash := sha256.New()
			if _, err := io.Copy(sha256hash, f); err != nil {
				panic(err)
			}
			SHA256ProgressBar.Hide()
			SHA256Hash.ParseMarkdown(fmt.Sprintf("`%x`", sha256hash.Sum(nil)))
			CopySHA256Button.Show()
		}()
	}
	wg.Wait()
}
