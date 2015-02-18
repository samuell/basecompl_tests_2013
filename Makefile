basecompl:
	go build -o basecompl basecompl.go

basecompl_par:
	go build -o basecompl_par basecompl_par.go

all: basecompl basecompl_par

clean:
	rm basecompl
	rm basecompl_par
