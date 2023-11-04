package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/linux-http/env"
	"github.com/lwinmgmg/linux-http/middleware"
)

// func getInput() *bufio.Reader {
// 	rd := bufio.NewReader(os.Stdin)
// 	return rd
// }

var Env env.Settings = env.NewEnv()

func main() {
	app := gin.New()
	app.Use(gin.CustomRecovery(middleware.PanicMiddleware))
	app.Run(fmt.Sprintf("%v:%v", Env.LH_HOST, Env.LH_PORT))
}

// for {
// 	fmt.Print("Type your input : ")
// 	inputBytes := getInput()
// 	input, err := inputBytes.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}
// 	if strings.TrimSpace(input) == "" {
// 		continue
// 	}
// 	valList := strings.Split(input, " ")
// 	var newList []string
// 	for _, v := range valList {
// 		newList = append(newList, strings.Trim(v, "\n"))
// 	}
// 	cmd := exec.Command(newList[0], newList[1:]...)
// 	cmd.Stdin = os.Stdin
// 	cmd.StdinPipe()
// 	r, err := cmd.StdoutPipe()
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := cmd.Start(); err != nil {
// 		panic(err)
// 	}
// 	buffer := bufio.NewReader(r)
// 	for {
// 		output, err := buffer.ReadByte()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			panic(err)
// 		}
// 		fmt.Print(string(output))
// 	}
// }
// }
