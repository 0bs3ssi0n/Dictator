package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

var dict_defaults = map[string] string{
	"case":		"all",
	"nums":		"medium",
	"years":	"small",
	"charset":	"special",

}

var charsets = map[string] string{
	"special": "!\"#$%&\\'()*+,-./:;<=>?@[\\\\]^_`{|}~",
	"upper": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"lower": "abcdefghijklmnopqrstuvwxyz",
}

func dict_afterfunc(){
	calc_del_next_checkpoint()
}

func dict_movetmp()(*os.File, *bufio.Scanner, *os.File){
	DICTMoveFinalToTmp()
	tmpF := DICTOpenTmpFile()
	finalScanner := DICTFileScanner(tmpF)
	finalF := DICTOpenFinalFile()
	return tmpF, finalScanner, finalF
}

func dict_checkargs(funcname string, arg string, boolList []*bool, stringList []string) bool{
	if arg == ""{
		arg = dict_defaults[funcname]
	}else if arg == "default"{
		arg = dict_defaults[funcname]
	}

	checkfunc := false
	for i, listItem := range stringList{
		if arg == listItem{
			*boolList[i] = true
			checkfunc = true
		}
	}
	return checkfunc
}

func dict_file(arg string){
	if arg == ""{
		return
	}

	tmpF, _, finalF := dict_movetmp()

	if DICTFileSize(tmpFileName) == 0{
		f, err := os.Open(arg)
		check(err)
		fReader := newFileReader(f)
		fWriter := newFileWriter(finalF)
		for {
			chunks, err:= fReader.ReadLines()
			if err != nil{
				break
			}
			for i, chunk := range(chunks){
				chunks[i] = append(chunk, []byte("\n")...)
			}
			fWriter.WriteLines(chunks)
		}
		f.Close()
	}else {
		tmpReader := newFileReader(tmpF)
		f, err := os.Open(arg)
		check(err)
		fReader := newFileReader(f)
		fWriter := newFileWriter(finalF)
		for {
			chunks, err := tmpReader.ReadLines()
			if err != nil{
				break
			}
			for i, chunk := range(chunks){
				for {
					fchunks, err := fReader.ReadLines()
					if err != nil {
						break
					}
					for fi, fchunk := range (fchunks) {
						fchunk = append(fchunk, []byte("\n")...)
						fchunks[fi] = append(chunk, fchunk...)
					}
					fWriter.WriteLines(fchunks)
				}
				fReader.ResetPointer()
				chunks[i] = append(chunk, []byte("\n")...)
			}
			fWriter.WriteLines(chunks)
		}
		f.Close()
	}
	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}

func dict_case(arg string){
	fupper := false
	upper := false
	lower := false
	all := false

	res := dict_checkargs("case", arg, []*bool{&fupper, &upper, &lower, &all}, []string{"firstupper", "upper", "lower", "all"})
	if !res{
		return
	}

	tmpF, _, finalF := dict_movetmp()


	fReader := newFileReader(tmpF)
	fWriter := newFileWriter(finalF)

	for {
		cChunks := [][]byte{}
		chunks, err := fReader.ReadLines()
		if err != nil {
			break
		}
		for i, chunk := range (chunks) {
			chunks[i] = append(chunk, []byte("\n")...)
			cChunks = append(cChunks, chunks[i])
			if fupper || all {
				nchunk := bytes.ToUpper(chunks[i][0:1])
				nchunk = append(nchunk, chunks[i][1:]...)
				cChunks = append(cChunks, nchunk)
			}
			if upper || all{
				nchunk := bytes.ToUpper(chunks[i])
				cChunks = append(cChunks, nchunk)
			}
			if lower || all{
				nchunk := bytes.ToLower(chunks[i])
				cChunks = append(cChunks, nchunk)
			}
		}
		fWriter.WriteLines(cChunks)
	}

	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}

func dict_nums(arg string){
	big := false
	medium := false
	small := false
	nano := false

	res := dict_checkargs("nums", arg, []*bool{&big, &medium, &small, &nano}, []string{"big", "medium", "small", "nano"})
	if !res{
		return
	}

	tmpF, _, finalF := dict_movetmp()

	fReader := newFileReader(tmpF)
	fWriter := newFileWriter(finalF)

	if DICTFileSize(tmpFileName) == 0{
		chunks := [][]byte{}
		if big {
			for i := 0; i <= 9999; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		} else if medium {
			for i := 0; i <= 999; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		} else if small {
			for i := 0; i <= 99; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		} else if nano {
			for i := 0; i <= 9; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		}
		fWriter.WriteLines(chunks)
	}else {
		for {
			cChunks := [][]byte{}
			chunks, err := fReader.ReadLines()
			if err != nil {
				break
			}
			for i, _ := range (chunks) {
				if big {
					for c := 0; c <= 9999; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)
					}
				} else if medium {
					for c := 0; c <= 999; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)				}
				} else if small {
					for c := 0; c <= 99; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)				}
				} else if nano {
					for c := 0; c <= 9; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)
					}
				}
				nChunk := []byte{}
				nChunk = append(nChunk, chunks[i]...)
				nChunk = append(nChunk, []byte("\n")...)
				cChunks = append(cChunks, nChunk)
			}

			fWriter.WriteLines(cChunks)
		}
	}

	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}

func dict_years(arg string){ // <small>(1995-2025) <medium>(1950-2030) <big>(1900-2030)
	big := false
	medium := false
	small := false

	res := dict_checkargs("years", arg, []*bool{&big, &medium, &small}, []string{"big", "medium", "small"})
	if !res{
		return
	}

	tmpF, _, finalF := dict_movetmp()

	fReader := newFileReader(tmpF)
	fWriter := newFileWriter(finalF)

	if DICTFileSize(tmpFileName) == 0{
		chunks := [][]byte{}
		if big {
			for i := 1900; i <= 2030; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		} else if medium {
			for i := 1950; i <= 2030; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		} else if small{
			for i := 1995; i <= 2025; i++ {
				chunks = append(chunks, []byte(strconv.Itoa(i)+"\n"))
			}
		}
		fWriter.WriteLines(chunks)
	}else {
		for {
			cChunks := [][]byte{}
			chunks, err := fReader.ReadLines()
			if err != nil {
				break
			}
			for i, _ := range (chunks) {
				if big {
					for c := 1900; c <= 2030; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)
					}
				} else if medium {
					for c := 1950; c <= 2030; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)					}
				} else {
					for c := 1995; c <= 2025; c++ {
						nChunk := []byte{}
						nChunk = append(nChunk, chunks[i]...)
						nChunk = append(nChunk, []byte(strconv.Itoa(c)+"\n")...)
						cChunks = append(cChunks, nChunk)					}
				}
				nChunk := []byte{}
				nChunk = append(nChunk, chunks[i]...)
				nChunk = append(nChunk, []byte("\n")...)
				cChunks = append(cChunks, nChunk)
			}
			fWriter.WriteLines(cChunks)
		}
	}

	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}

func dict_chars(arg string){
	if arg == ""{
		return
	}

	tmpF, _, finalF := dict_movetmp()

	fReader := newFileReader(tmpF)
	fWriter := newFileWriter(finalF)

	if DICTFileSize(tmpFileName) == 0{
		chunks := [][]byte{}
		for _, c := range (strings.Split(arg, "")) {
			nChunk := []byte{}
			nChunk = append(nChunk, []byte(c+"\n")...)
			chunks = append(chunks, nChunk)
		}
		fWriter.WriteLines(chunks)
	}else {
		for {
			cChunks := [][]byte{}
			chunks, err := fReader.ReadLines()
			if err != nil {
				break
			}
			for i, _ := range(chunks){
				for _, c := range (strings.Split(arg, "")) {
					nChunk := []byte{}
					nChunk = append(nChunk, chunks[i]...)
					nChunk = append(nChunk, []byte(c+"\n")...)
					cChunks = append(cChunks, nChunk)
				}
				nChunk := []byte{}
				nChunk = append(nChunk, chunks[i]...)
				nChunk = append(nChunk, []byte("\n")...)
				cChunks = append(cChunks, nChunk)
			}
			fWriter.WriteLines(cChunks)
		}
	}

	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}

func dict_charset(arg string){
	lower := false
	upper := false
	special := false

	res := dict_checkargs("charset", arg, []*bool{&lower, &upper, &special}, []string{"lower", "upper", "special"})
	if !res{
		return
	}

	tmpF, _, finalF := dict_movetmp()

	fReader := newFileReader(tmpF)
	fWriter := newFileWriter(finalF)

	if DICTFileSize(tmpFileName) == 0{
		chunks := [][]byte{}
		for _, c := range (strings.Split(charsets[arg], "")) {
			nChunk := []byte{}
			nChunk = append(nChunk, []byte(c+"\n")...)
			chunks = append(chunks, nChunk)
		}
		fWriter.WriteLines(chunks)
	}else {
		for {
			cChunks := [][]byte{}
			chunks, err := fReader.ReadLines()
			if err != nil {
				break
			}
			for i, _ := range(chunks) {
				for _, c := range (strings.Split(charsets[arg], "")) {
					nChunk := []byte{}
					nChunk = append(nChunk, chunks[i]...)
					nChunk = append(nChunk, []byte(c+"\n")...)
					cChunks = append(cChunks, nChunk)
				}
				nChunk := []byte{}
				nChunk = append(nChunk, chunks[i]...)
				nChunk = append(nChunk, []byte("\n")...)
				cChunks = append(cChunks, nChunk)
			}
			fWriter.WriteLines(cChunks)
		}
	}

	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}

func dict_string(arg string){
	if arg == ""{
		return
	}

	tmpF, _, finalF := dict_movetmp()

	fReader := newFileReader(tmpF)
	fWriter := newFileWriter(finalF)


	if DICTFileSize(tmpFileName) == 0{
		chunks := [][]byte{}
		nChunk := []byte{}
		nChunk = append(nChunk, []byte(arg+"\n")...)
		chunks = append(chunks, nChunk)
		fWriter.WriteLines(chunks)
	}else {
		for {
			cChunks := [][]byte{}
			chunks, err := fReader.ReadLines()
			if err != nil {
				break
			}
			for i, _ := range(chunks) {

				nChunk := []byte{}
				nChunk = append(nChunk, chunks[i]...)
				nChunk = append(nChunk, []byte("\n")...)
				cChunks = append(cChunks, nChunk)

				oChunk := []byte{}
				oChunk = append(oChunk, chunks[i]...)
				oChunk = append(oChunk, []byte(arg+"\n")...)
				cChunks = append(cChunks, oChunk)

			}
			fWriter.WriteLines(cChunks)
		}
	}

	DICTCloseFile(tmpF)
	DICTCloseFile(finalF)

	dict_afterfunc()
}
