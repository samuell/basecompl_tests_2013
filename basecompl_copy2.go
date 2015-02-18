// --------------------------------------------------------------------
// Dummy processing DNA code (by taking base complement) 
// Sequential version, with a string copy, to imitate
// algorithms that produces a new string
// Author: Samuel Lampa, samuel.lampa@gmail.com
// Date: 2013-06-21
// --------------------------------------------------------------------

package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
    // Read file with DNA of Human Chromosome Y
	file, err := os.Open("Homo_sapiens.GRCh37.67.dna_rm.chromosome.Y.fa")
	if err != nil {
        log.Fatal(err)
	} else {
        scan := bufio.NewScanner(file)
        for scan.Scan() {
            line := scan.Bytes()
            // Set up a pipeline of sequential "dummy processing"
            // (Taking the base complement back and forth a few
            // times).
            lineNew1 := complement(line)
            lineNew2 := complement(lineNew1)
            lineNew3 := complement(lineNew2)
            lineNew4 := complement(lineNew3)
            fmt.Println(string(lineNew4))
        }
	}
}

func complement(line []byte) []byte {
    // Ugly and naive code to compute the DNA base complement
    for pos := range line {
        if line[pos] == 'A' {
            line[pos] = 'T'
        } else if line[pos] == 'T' {
            line[pos] = 'A'
        } else if line[pos] == 'C' {
            line[pos] = 'G'
        } else if line[pos] == 'G' {
            line[pos] = 'C'
        }
    }
    return append([]byte(nil), line...)
}
