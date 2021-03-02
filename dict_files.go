package main

import (
	"bufio"
	"bytes"
	"io"
	"math/rand"
	"os"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func DICTGenerateOutFiles(){
	tmpFileName = "DICTator_" + RandStringRunes(10)
	finalFileName = RandStringRunes(10)
	DICTCreateFiles()

}

func DICTClear(all bool){
	os.Remove(tmpFileName)
	if all{
		os.Remove(finalFileName)
	}
}

func DICTFileSize(path string)int{
	f, err := os.Stat(path)
	if err != nil{
		return -1
	}
	return int(f.Size())
}

func DICTLineCount(path string) (int, error) {
	f, err := os.Open(path)
	check(err)
	r := bufio.NewReader(f)
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func DICTMoveFinalToTmp(){
	err := os.Rename(finalFileName, tmpFileName)
	check(err)
}

func DICTOpenTmpFile() *os.File{
	f, err := os.Open(tmpFileName)
	check(err)
	return f
}

func DICTCreateFiles(){
	f, err := os.OpenFile(finalFileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	f.Close()
	f, err = os.OpenFile(tmpFileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	f.Close()
}

func DICTOpenFinalFile() *os.File{
	f, err := os.OpenFile(finalFileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	return f
}

func DICTFileScanner(f *os.File) *bufio.Scanner{
	scanner := bufio.NewScanner(f)
	return scanner
}


func DICTCloseFile(f *os.File){
	f.Close()
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
