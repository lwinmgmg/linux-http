package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func getInput() *bufio.Reader {
	rd := bufio.NewReader(os.Stdin)
	return rd
}

func main() {

	for {
		fmt.Print("Type your input : ")
		inputBytes := getInput()
		input, err := inputBytes.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(input) == "" {
			continue
		}
		valList := strings.Split(input, " ")
		var newList []string
		for _, v := range valList {
			newList = append(newList, strings.Trim(v, "\n"))
		}
		cmd := exec.Command(newList[0], newList[1:]...)
		w, _ := cmd.StdinPipe()
		var ch = make(chan struct{}, 1)
		go func() {
			var charch = make(chan string, 1)
			a := getInput()
			go func() {
				i, _ := a.ReadString('\n')
				charch <- i
			}()
			select {
			case dt := <-charch:
				w.Write([]byte(dt))
				close(a)
				return
			case <-ch:
				return
			}
		}()
		r, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		buffer := bufio.NewReader(r)
		for {
			output, err := buffer.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			fmt.Print(string(output))
		}
	}
}
