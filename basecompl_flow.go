package main

import (
	"bufio"
	"fmt"
	"github.com/trustmaster/goflow"
	"log"
	"os"
	"runtime"
)

const (
	BUFSIZE    = 16
	NUMTHREADS = 4
)

var baseConv = [256]byte{
	'A': 'T',
	'T': 'A',
	'C': 'G',
	'G': 'C',
	'N': 'N',
}

// -------------------------------------------------------------
// Basecomlementer Component
// -------------------------------------------------------------
type BaseComplementer struct {
	flow.Component               // Embedding "superclass"
	Dna            <-chan []byte // Input port
	DnaBaseCompl   chan<- []byte // Output port
}

func (bc *BaseComplementer) OnDna(sequence []byte) {
	for pos := range sequence {
		sequence[pos] = baseConv[sequence[pos]]
	}
	bc.DnaBaseCompl <- sequence
}

// -------------------------------------------------------------
// Printer
// -------------------------------------------------------------
type Printer struct {
	flow.Component
	Line <-chan []byte // Input
}

// Prints a line when it gets it
func (p *Printer) OnLine(line []byte) {
	fmt.Println(string(line))
}

// ---------------------------------------------------------
// BaseComplementer network
// ---------------------------------------------------------
type BaseComplementerApp struct {
	flow.Graph
}

func NewBaseCompelenterApp() *BaseComplementerApp {
	network := new(BaseComplementerApp)
	network.InitGraphState()

	// Add components
	network.Add(new(BaseComplementer), "basecompl1")
	network.Add(new(BaseComplementer), "basecompl2")
	network.Add(new(BaseComplementer), "basecompl3")
	network.Add(new(BaseComplementer), "basecompl4")
	network.Add(new(Printer), "printer")

	// Connect components
	network.Connect("basecompl1", "DnaBaseCompl", "basecompl2", "Dna", make(chan []byte, BUFSIZE))
	network.Connect("basecompl2", "DnaBaseCompl", "basecompl3", "Dna", make(chan []byte, BUFSIZE))
	network.Connect("basecompl3", "DnaBaseCompl", "basecompl4", "Dna", make(chan []byte, BUFSIZE))
	network.Connect("basecompl4", "DnaBaseCompl", "printer", "Line", make(chan []byte, BUFSIZE))
	network.MapInPort("In", "basecompl1", "Dna")

	return network
}

// ---------------------------------------------------------
// Main method
// ---------------------------------------------------------
var finish chan bool

// Use this handler to let main() know when the network terminates
func (a *BaseComplementerApp) Finish() {
	finish <- true
}

func main() {
	// Set the number of Operating System-threads to use
	fmt.Println("Starting ", NUMTHREADS, " threads ...")
	runtime.GOMAXPROCS(NUMTHREADS)

	// Termination signal channel
	finish = make(chan bool)
	// Create network
	net := NewBaseCompelenterApp()

	// Create the "In" channel
	in := make(chan []byte, BUFSIZE)
	net.SetInPort("In", in)

	// Run net
	flow.RunNet(net)

	file, err := os.Open("Homo_sapiens.GRCh37.67.dna_rm.chromosome.Y.fa")
	if err != nil {
		log.Fatal(err)
	} else {
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			line := scan.Bytes()
			in <- line
		}
	}

	// Close the input to shut the network down
	close(in)
	// ... wait until done ...
	<-finish
}
