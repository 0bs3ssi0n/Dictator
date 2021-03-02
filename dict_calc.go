package main

import (
	"strconv"
	"strings"
)

type progress_checkpoint struct {
	id int
	name string
	counter int
	endFileSize int
}

var checkpoints []progress_checkpoint

var estSize = 0
var estLines = 0

func calc_add_checkpoint(name string, endFileSize int){
	highestCount := 0
	highestId := 0
	for _, checkpoint := range(checkpoints){
		if name == checkpoint.name && highestCount < checkpoint.counter{
			highestCount = checkpoint.counter
		}
		if highestId < checkpoint.id{
			highestId = checkpoint.id
		}
	}
	highestId++
	highestCount++
	checkpoints = append(checkpoints, progress_checkpoint{
		id:		     highestId,
		name:        name,
		counter:     highestCount,
		endFileSize: endFileSize,
	})
}

func calc_del_next_checkpoint(){
	nID := calc_get_next_checkpoint().id
	ncheckpoints := []progress_checkpoint{}
	for _, checkpoint := range(checkpoints){
		if nID != checkpoint.id{
			ncheckpoints = append(ncheckpoints, checkpoint)
		}
	}
	checkpoints = ncheckpoints
}

func calc_get_last_checkpoint() progress_checkpoint{
	var sID int
	for _, checkpoint := range(checkpoints){
		if sID < checkpoint.id || sID == 0{
			sID = checkpoint.id
		}
	}
	if sID != 0{
		return calc_get_checkpoint(sID)
	}else{
		return progress_checkpoint{}
	}
}

func calc_get_next_checkpoint() progress_checkpoint{
	var sID int
	for _, checkpoint := range(checkpoints){
		if sID > checkpoint.id || sID == 0{
			sID = checkpoint.id
		}
	}
	if sID != 0{
		return calc_get_checkpoint(sID)
	}else{
		return progress_checkpoint{}
	}
}

func calc_get_checkpoint(id int) progress_checkpoint{
	for _, checkpoint := range(checkpoints){
		if id == checkpoint.id{
			return checkpoint
		}
	}
	return progress_checkpoint{}
}

func calc_num_to_string(byts int)string{
	rbyts := ""
	if byts < 1000{
		rbyts += strconv.Itoa(byts) + " bytes"
	}else if byts >= 1000 && byts < 1000000{
		rbyts += strconv.Itoa(byts/1000) + " Kb"
	}else if byts >= 1000000 && byts < 1000000000{
		rbyts += strconv.Itoa(byts/1000000) + " Mb"
	}else{
		rbyts += strconv.Itoa(byts/1000000000) + " Gb"
	}
	return rbyts
}

func calc_file(arg string) int{
	if arg == ""{
		return -1
	}
	fileLines, err := DICTLineCount(arg)
	check(err)
	fileSize := DICTFileSize(arg)

	if  estSize == 0{
		estSize = fileSize
		estLines = fileLines
	}else{
		newSize := 0
		newSize += estSize
		newSize += (fileSize * estLines)
		newSize += ((estSize - estLines) * fileLines)
		estSize = newSize
		estLines += (fileLines * estLines)
	}
	return estSize
}

func calc_case(arg string) int{
	fupper := false
	upper := false
	lower := false
	all := false

	res := dict_checkargs("case", arg, []*bool{&fupper, &upper, &lower, &all}, []string{"firstupper", "upper", "lower", "all"})
	if !res{
		return -1
	}

	if  estSize == 0{
		return -1
	}else{
		newSize := 0
		newLines := 0
		newSize += estSize
		newLines += estLines
		if fupper || all{
			newSize += estSize
			newLines += estLines
		}
		if upper || all{
			newSize += estSize
			newLines += estLines
		}
		if lower || all{
			newSize += estSize
			newLines += estLines
		}
		estSize = newSize
		estLines = newLines
	}
	return estSize
}

func calc_nums(arg string) int{
	big := false
	medium := false
	small := false
	nano := false

	res := dict_checkargs("nums", arg, []*bool{&big, &medium, &small, &nano}, []string{"big", "medium", "small", "nano"})
	if !res{
		return -1
	}

	if  estSize == 0{
		if big { // 9999
			estSize += (1*10 + 10) + (2*90 + 90) + (3 * 900 + 900) + (4 * 9000 + 9000)
			estLines += 10000
		} else if medium { // 999
			estSize += (1*10 + 10) + (2*90 + 90) + (3 * 900 + 900)
			estLines += 1000
		} else if small { // 99
			estSize += (1*10 + 10) + (2*90 + 90)
			estLines += 100
		} else if nano { // 9
			estSize += (1*10 + 1*10)
			estLines += 10
		}
	}else{
		newSize := 0
		newLines := 0
		if big { // 9999
			newSize += estSize
			newLines += estLines
			newSize += (((1*10) + (2*90) + (3 * 900) + (4 * 9000)) * estLines)
			newSize += (10000 * estSize)
			newLines += (10000 * estLines)
			estSize = newSize
			estLines = newLines
		} else if medium { // 999
			newSize += estSize
			newLines += estLines
			newSize += (((1*10) + (2*90) + (3 * 900)) * estLines)
			newSize += (1000 * estSize)
			newLines += (1000 * estLines)
			estSize = newSize
			estLines = newLines
		} else if small { // 99
			newSize += estSize
			newLines += estLines
			newSize += (((1*10) + (2*90)) * estLines)
			newSize += (100 * estSize)
			newLines += (100 * estLines)
			estSize = newSize
			estLines = newLines
		} else if nano { // 9
			newSize += estSize
			newLines += estLines
			newSize += (((1*10)) * estLines)
			newSize += (10 * estSize)
			newLines += (10 * estLines)
			estSize = newSize
			estLines = newLines
		}
	}
	return estSize
}

func calc_years(arg string) int{ // <small>(1995-2025) (31) <medium>(1950-2030) (81) <big>(1900-2030) (131)
	big := false
	medium := false
	small := false

	res := dict_checkargs("years", arg, []*bool{&big, &medium, &small}, []string{"big", "medium", "small"})
	if !res{
		return -1
	}

	if  estSize == 0{
		if big {
			estSize += (4 * 131 + 131)
			estLines += 131
		} else if medium {
			estSize += (4 * 81 + 81)
			estLines += 81
		} else if small {
			estSize += (4 * 31 + 31)
			estLines += 31
		}
	}else{
		newSize := 0
		newLines := 0
		if big {
			newSize += estSize
			newLines += estLines
			newSize += ((4 * 131) * estLines)
			newSize += (131 * estSize)
			newLines += (131 * estLines)
			estSize = newSize
			estLines = newLines
		} else if medium {
			newSize += estSize
			newLines += estLines
			newSize += ((4 * 81) * estLines)
			newSize += (81 * estSize)
			newLines += (81 * estLines)
			estSize = newSize
			estLines = newLines
		} else if small {
			newSize += estSize
			newLines += estLines
			newSize += ((4 * 31) * estLines)
			newSize += (31 * estSize)
			newLines += (31 * estLines)
			estSize = newSize
			estLines = newLines
		}
	}
	return estSize
}

func calc_chars(arg string) int{
	if arg == ""{
		return -1
	}

	charsetSize := len(strings.Split(arg, ""))

	newSize := 0
	newLines := 0
	if estSize == 0{
		newSize += (charsetSize*2)
		newLines += charsetSize
	}else{
		newSize += estSize
		newLines += estLines
		newSize += (charsetSize * estSize) + (estLines * charsetSize)
		newLines += (charsetSize * estLines)
	}
	estSize = newSize
	estLines = newLines

	return estSize
}

func calc_charset(arg string) int{
	lower := false
	upper := false
	special := false

	res := dict_checkargs("charset", arg, []*bool{&lower, &upper, &special}, []string{"lower", "upper", "special"})
	if !res{
		return -1
	}

	charsetSize := len(strings.Split(charsets[arg], ""))

	newSize := 0
	newLines := 0
	if estSize == 0{
		newSize += (charsetSize*2)
		newLines += charsetSize
	}else{
		newSize += estSize
		newLines += estLines
		newSize += (charsetSize * estSize) + (estLines * charsetSize)
		newLines += (charsetSize * estLines)
	}
	estSize = newSize
	estLines = newLines
	return estSize
}

func calc_text(arg string) int{
	if arg == ""{
		return -1
	}

	newSize := 0
	newLines := 0
	newSize += estSize
	newLines += estLines

	if estSize == 0{
		newSize += (len(arg)+1)
		newLines += 1
	}else {
		newSize += (len(arg) * estLines) + estSize
		newLines += estLines
	}
	estSize = newSize
	estLines = newLines

	return estSize
}

func calc_none(arg string) int{
	return -1
}
