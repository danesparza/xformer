package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//	Set up our flags
	inputfile := flag.String("inputfile", "yourdatafile.csv", "The input (data) file")
	outputfile := flag.String("outputfile", "youroutput.dat", "The ouput (result) file")
	templatefile := flag.String("templatefile", "template.dat", "The template file to use on every data row")

	//	Parse the command line for flags:
	flag.Parse()

	//	Print out our args:
	fmt.Printf("\nInput file: %v\n", *inputfile)
	fmt.Printf("Template file: %v\n", *templatefile)
	fmt.Printf("Output file: %v\n", *outputfile)

	//	Read the input file in:
	inputbuffer, err := ioutil.ReadFile(*inputfile)
	if err != nil {
		log.Fatal(err)
	}

	//	Convert it to a string and then split into an array of lines:
	inputlines := strings.Split(string(inputbuffer), "\r\n")

	//	Read the template file in:
	templatebuffer, err := ioutil.ReadFile(*templatefile)
	if err != nil {
		log.Fatal(err)
	}

	//	Convert it to a string:
	templatedata := strings.TrimSpace(string(templatebuffer))

	//	Loop through each input line (if it actually has something)
	outputLines := []string{}
	for _, inputline := range inputlines {
		trimmedLine := strings.TrimSpace(inputline)
		if len(trimmedLine) > 0 {
			outputLines = append(outputLines, fmt.Sprintf(templatedata, inputline))
		}
	}

	//	Dump the buffer to the output file:
	outputBuffer := strings.Join(outputLines, "\n")
	ba := []byte(outputBuffer)
	err = ioutil.WriteFile(*outputfile, ba, 0755)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nWrote %v bytes to the output file: %v\n", len(ba), *outputfile)
	}
}
