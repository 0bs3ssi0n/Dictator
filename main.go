package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DICTCommand struct {
	name string
	longname string
	descr string
	function func(data string)
	cfunction func(data string)int
}

var DICTCommandList []DICTCommand

var finalFileName string
var tmpFileName string

func addDICTCommand(name string, longname string, descr string, function func(data string), cfunction func(data string)int){
	DICTCommandList = append(DICTCommandList,
		DICTCommand{
			name:     name,
			longname: longname,
			descr:    descr,
			function: function,
			cfunction: cfunction,
		})
}

func DICTUsage(arg string){
	msg := "Usage: Dictator <arguments>\n"
	msg += "DICTIONARY GENERATOR FLAGS:\n"
	for _, cmd := range DICTCommandList{
		msg += "  " + cmd.name + "  " + cmd.longname + ": " + cmd.descr + "\n"
	}
	msg += "EXAMPLES:\n"
	msg += "  " + "Dictator --file ./list.txt --case --chars \"!@#%\"" + "\n"
	msg += "  " + "Dictator --file ./list.txt --chars '_-' --file ./list.txt" + "\n"
	msg += "  " + "Dictator --nums nano --file ./list.txt --nums nano --chars '!@#%' --chars '!@#%'" + "\n"

	fmt.Println(msg)
}

func DICTPrefunc(){
	DICT_args := os.Args[1:]
	DICT_argsSize := len(DICT_args)
	if DICT_argsSize == 0{
		DICTUsage("")
		os.Exit(0)
	}
	for _, cmd := range DICT_args {
		if ((cmd == "--help")||(cmd == "-h")){
			DICTUsage("")
			os.Exit(0)
		}
	}
	for i, cmd := range DICT_args {
		for _, Dcmd := range DICTCommandList{
			if cmd == Dcmd.name || cmd == Dcmd.longname{
				if (i+1) < DICT_argsSize{
					nextCmd := DICT_args[i+1]
					passArgs := true
					for _, CheckDcmd := range DICTCommandList{
						if CheckDcmd.name == nextCmd || CheckDcmd.longname == nextCmd{
							passArgs = false
							eSize := Dcmd.cfunction("")
							if eSize != -1{
								calc_add_checkpoint(CheckDcmd.longname, eSize)
							}
							break
						}
					}
					if passArgs{
						eSize := Dcmd.cfunction(nextCmd)
						if eSize != -1{
							calc_add_checkpoint(Dcmd.longname, eSize)
						}
					}
				}else {
					eSize := Dcmd.cfunction("")
					if eSize != -1{
						calc_add_checkpoint(Dcmd.longname, eSize)
					}
				}
			}
		}
	}

	fmt.Printf("Estimated file size: %s\n", calc_num_to_string(estSize))
	//fmt.Printf("Estimated file size: %d\n", (estSize))

}

func DICTator(){
	DICT_args := os.Args[1:]
	DICT_argsSize := len(DICT_args)
	go func() {
		for {
			time.Sleep(100000000)
			dict_checkpoint_progress()
		}
	}()
	for i, cmd := range DICT_args {
		for _, Dcmd := range DICTCommandList{
			if cmd == Dcmd.name || cmd == Dcmd.longname{
				if (i+1) < DICT_argsSize{
					nextCmd := DICT_args[i+1]
					passArgs := true
					for _, CheckDcmd := range DICTCommandList{
						if CheckDcmd.name == nextCmd || CheckDcmd.longname == nextCmd{
							passArgs = false
							Dcmd.function("")
							break
						}
					}
					if passArgs{
						Dcmd.function(nextCmd)
					}
				}else {
					Dcmd.function("")
				}
			}
		}
	}
	if DICT_argsSize == 0{
		DICTClear(true)
	}else{
		DICTClear(false)
		fmt.Println("\nFile has been created! ("+finalFileName+")")
	}
}

func main(){
	dict_header()
	SetupCloseHandler()
	// Add num range, -n 500-1000
	// Add l33t mode
	// Check for speed improvements, shits slow af
	addDICTCommand("-f", "--file", "Specify a file where every line will be appended to each line in current dictionary.", dict_file, calc_file)
	addDICTCommand("-C", "--case", "Specify if it should change between lower/uppercase. <firstupper/upper/lower/all>", dict_case, calc_case)
	addDICTCommand("-n", "--nums", "Specify numbers to append. <nano>(0-9) <small>(0-99) <medium>(0-999) <big>(0-9999)", dict_nums, calc_nums)
	addDICTCommand("-y", "--years", "Specify years to append. <small>(1995-2025) <medium>(1950-2030) <big>(1900-2030)", dict_years, calc_years)
	addDICTCommand("-c", "--chars", "Specify characters it should append after each line. <chars>", dict_chars, calc_chars)
	addDICTCommand("-cs", "--charset", "Specify pre-defined charset it should append after each line. <special/uppercase/lowercase>", dict_charset, calc_charset)
	addDICTCommand("-t", "--text", "Specify a piece of text it should append after each line.", dict_string, calc_text)
	addDICTCommand("-h", "--help", "Show help menu.", DICTUsage, calc_none)

	DICTGenerateOutFiles()
	DICTPrefunc()

	DICTator()

}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		DICTClear(true)
		fmt.Println("\nInterrupted...")
		os.Exit(0)
	}()
}