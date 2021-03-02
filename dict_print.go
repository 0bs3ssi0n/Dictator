package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dict_header(){
	fmt.Println("Dictator V1.0 | Dictionary Generator")
}

func dict_checkpoint_progress() int{
	clearLine()
	checkPoint := calc_get_next_checkpoint()
	lastCheckPoint := calc_get_last_checkpoint()
	curSize := DICTFileSize(finalFileName)
	if curSize == -1 || len(checkPoint.name) == 0{
		return -1
	}
	endSize := checkPoint.endFileSize
	fmt.Print("Appending " + checkPoint.name[2:len(checkPoint.name)] + " | step [" + strconv.Itoa(checkPoint.id) + "/"+ strconv.Itoa(lastCheckPoint.id) +"]: ")
	progress_print(curSize, endSize)
	return 1
}

func progress_print(curProgress int, finProgress int){
	block := "▅"
	noblock := " "
	percentage := int(float64(curProgress) / (float64(finProgress) / 100.0))
	if percentage > 0 {
		msg := "["
		msg += strings.Repeat(block, percentage/2)
		msg += strings.Repeat(noblock, 50-(percentage/2))
		msg += "] " + strconv.Itoa(percentage) + "%"
		fmt.Print(msg)
	}
}

func clearLine()  {
	fmt.Print(strings.Repeat("\b", 100)+strings.Repeat(" ", 100)+strings.Repeat("\b", 100))
}
