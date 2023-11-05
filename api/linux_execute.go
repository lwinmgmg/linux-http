package api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/linux-http/middleware"
	"golang.org/x/exp/slices"
)

type LinuxCmd struct {
	Cmd   string `json:"cmd"`
	Input string `json:"input"`
}

type LinuxCmdResp struct {
	Output   string `json:"output"`
	ExitCode int    `json:"exit_code"`
}

func (v1 *ControllerV1) LinuxExecute(ctx *gin.Context) {
	cmdInput := LinuxCmd{}
	if err := ctx.ShouldBindJSON(&cmdInput); err != nil {
		panic(middleware.NewPanic(http.StatusUnprocessableEntity, 1, fmt.Sprintf("Can't parse json %v", err)))
	}
	if len(Env.LH_ALLOW_CMDS) > 0 {
		if !slices.Contains(Env.LH_ALLOW_CMDS, cmdInput.Cmd) {
			panic(middleware.NewPanic(http.StatusNotAcceptable, 1, "the input cmd is not allowed"))
		}
	}
	cmdNArgs := strings.Split(cmdInput.Cmd, " ")
	cmdName := cmdNArgs[0]
	args := []string{}
	if len(cmdNArgs) > 1 {
		args = cmdNArgs[1:]
	}
	cmd := exec.Command(cmdName, args...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(middleware.NewPanic(http.StatusBadRequest, 1, err.Error()))
	}
	stdOut := bufio.NewReader(out)
	var stdIn *bufio.Writer
	if cmdInput.Input != "" {
		in, err := cmd.StdinPipe()
		if err != nil {
			panic(middleware.NewPanic(http.StatusBadRequest, 2, err.Error()))
		}
		stdIn = bufio.NewWriter(in)
	}
	if err := cmd.Start(); err != nil {
		panic(middleware.NewPanic(http.StatusBadRequest, 3, err.Error()))
	}
	if stdIn != nil {
		go stdIn.Write([]byte(cmdInput.Input))
	}
	buffer := bufio.NewReader(stdOut)
	var outputStr string = ""
	for {
		output, err := buffer.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(middleware.NewPanic(http.StatusBadRequest, 4, err.Error()))
		}
		outputStr += string(output)
	}
	pState, err := cmd.Process.Wait()
	exitCode := pState.ExitCode()
	if exitCode == 0 {
		ctx.JSON(http.StatusOK, LinuxCmdResp{
			Output:   outputStr,
			ExitCode: exitCode,
		})
		return
	}
	panic(middleware.NewPanic(http.StatusAccepted, 1, "", map[string]any{
		"output":    outputStr,
		"exit_code": exitCode,
	}))
}
