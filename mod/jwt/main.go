package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lwinmgmg/linux-http/utils"
)

func printEachLine(body []byte) {
	data := ResponseData{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	for _, v := range strings.Split(data.Output, "\n") {
		fmt.Println(v)
	}
	fmt.Println("Exit Code :", data.ExitCode)
}

type RequestData struct {
	Cmd   string `json:"cmd"`
	Input string `json:"input"`
}

type ResponseData struct {
	Output   string `json:"output"`
	ExitCode int    `json:"exit_code"`
}

func main() {
	issuer, ok := os.LookupEnv("JWT_ISSUER")
	if !ok {
		panic("JWT_ISSUER is required")
	}
	secret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		panic("JWT_SECRET is required")
	}
	url, ok := os.LookupEnv("JWT_URL")
	if !ok {
		panic("JWT_URL is required")
	}
	tkn := utils.GetToken(issuer, secret, time.Minute)

	cmd, ok := os.LookupEnv("LH_CMD")
	if !ok {
		cmd = "ls -ahl"
	}
	input, ok := os.LookupEnv("LH_INPUT")
	if !ok {
		input = ""
	}
	data, err := json.Marshal(RequestData{
		Cmd:   cmd,
		Input: input,
	})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+tkn)
	client := http.Client{
		Timeout: time.Minute,
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		fmt.Println(string(body))
		panic(res.StatusCode)
	}
	printEachLine(body)
}
