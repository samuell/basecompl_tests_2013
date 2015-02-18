/// --------------------------------------------------------------------
// Dummy processing DNA code (by taking base complement)
// Threaded version
// Author: Samuel Lampa, samuel.lampa@gmail.com
// Date: 2013-06-21
// --------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	BUFSIZE    = 512
	NUMTHREADS = 3 // (NumCPU - 1)
)

var baseConv = [256]byte{
	'A': 'T',
	'T': 'A',
	'C': 'G',
	'G': 'C',
	'N': 'N',
}

func main() {
	// Set the number of Operating System-threads to use
	fmt.Println("Starting ", NUMTHREADS, " threads ...")
	runtime.GOMAXPROCS(NUMTHREADS)

	// Read DNA file
	inFileName := "Homo_sapiens.GRCh37.67.dna_rm.chromosome.Y.fa"
	fileReadChan := fileReaderGen(inFileName)

	// Setting up a sequence of dummy processing, where
	// each operation is done in a separate goroutine
	// and thus possibly will be multiplexed upon an OS
	// thread), and communication is done through channels
	// between each pair of goroutines.
	complChan1 := baseComplementGen(fileReadChan)
	complChan2 := baseComplementGen(complChan1)
	complChan3 := baseComplementGen(complChan2)
	complChan4 := baseComplementGen(complChan3)

	// Read the last channel in the "pipeline"
	for line := range complChan4 {
		if len(line) > 0 {
			fmt.Println(string(line))
		}
	}
}

// create a goroutine that will read the dna file
// line by line. it returns a channel which can be
// read line by line, to "draw" the goroutine's execution.
func fileReaderGen(filename string) chan []byte {
	fileReadChan := make(chan []byte, BUFSIZE)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		} else {
			scan := bufio.NewScanner(file)
			for scan.Scan() {
				fileReadChan <- scan.Bytes()
			}
			close(fileReadChan)
			fmt.Println("Closed file reader channel")
		}
	}()
	return fileReadChan
}

// Initiates a goroutine that will take the base complement
// of each line in the input channel specified, and returns
// an output channel that can be read line by line.
func baseComplementGen(inChan chan []byte) chan []byte {
	returnChan := make(chan []byte, BUFSIZE)
	go func() {
		for line := range inChan {
			for pos := range line {
				line[pos] = baseConv[line[pos]]
			}
			returnChan <- append([]byte(nil), line...)
		}
		close(returnChan)
		fmt.Println("Closed base complement generator channel")
	}()
	return returnChan
}
