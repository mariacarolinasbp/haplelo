package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Haplotype struct {
	id        int
	sequence  string
	frequency int
}

type Patient struct {
	id string
	h1 int
	h2 int
}

func main() {
	// var err error
	const INP string = "inp/acre.inp"
	const OUT string = "out/acre"
	// cmd := exec.Command("phase", "-MS", "-f1", "-S666", INP, OUT, "8000", "300", "1000")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stdout

	// if err = cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	content, _ := ioutil.ReadFile(OUT)
	buff := strings.Split(string(content), "\n")

	var lsb, lse, bpb, bpe int
	for i := 0; i < len(buff); i++ {
		if buff[i] == "BEGIN LIST_SUMMARY" {
			lsb = i
		}
		if buff[i] == "END LIST_SUMMARY" {
			lse = i
		}
		if buff[i] == "BEGIN BESTPAIRS_SUMMARY" {
			bpb = i
		}
		if buff[i] == "END BESTPAIRS_SUMMARY" {
			bpe = i
		}
	}

	haplotypes := make([]Haplotype, 0)
	for i := lsb + 1; i < lse; i++ {
		var h Haplotype
		fields := strings.Fields(buff[i])
		id, _ := strconv.Atoi(fields[0])
		h.id = id
		h.sequence = fields[1]
		freq, _ := strconv.Atoi(strings.Split(fields[2], ".")[0])
		h.frequency = freq
		haplotypes = append(haplotypes, h)
	}

	patients := make([]Patient, 0)
	for i := bpb + 1; i < bpe; i++ {
		var p Patient
		replacer := strings.NewReplacer(":", " ", "(", " ", ",", " ", ")", " ")
		fields := strings.Fields(replacer.Replace(buff[i]))
		p.id = fields[0]
		h1, _ := strconv.Atoi(fields[1])
		p.h1 = h1
		h2, _ := strconv.Atoi(fields[2])
		p.h2 = h2
		patients = append(patients, p)
	}

	log.Println(haplotypes, patients)
}
