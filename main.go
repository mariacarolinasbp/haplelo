package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	f, err := os.OpenFile("out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	cmd := exec.Command("phase", "-MS", "-f1", "-S666", "acre.inp", "out/acre", "8000", "300", "1000")
	log.Printf("Running: %v", cmd.String())

	mwriter := io.MultiWriter(f, os.Stdout)
	cmd.Stderr = mwriter
	cmd.Stdout = mwriter

	err = cmd.Run()
	log.Printf("exit error: %v", err)
}
