package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Progress struct{}

func (l *Progress) Write(p []byte) (n int, err error) {
	fmt.Println("out>", string(p))
	return len(p), nil
}

func main() {
	cmd := exec.Command("phase", "-MS", "-f1", "-S666", "acre.inp", "out/acre", "8000", "300", "1000")

	var prgs Progress
	cmd.Stderr = &prgs

	err := cmd.Run()
	log.Printf("exit error: %v", err)
}
