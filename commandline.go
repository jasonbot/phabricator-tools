package phabricatortools

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/karlseguin/typed"
)

// CommandLineClient implements a very small subset of the gonduit Conn object's calls
type CommandLineClient interface {
	Call(method string, inputs interface{}, outputs interface{}) error
}

type commandLineClient struct {
}

func (*commandLineClient) Call(method string, inputs interface{}, outputs interface{}) error {
	cmd := exec.Command("arc", "call-conduit", method)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stdinContent, err := json.Marshal(inputs)
	if err != nil {
		return err
	}
	_, err = stdin.Write(stdinContent)
	if err != nil {
		return err
	}
	stdin.Close()

	//cmd.Start()
	cmd.Wait()

	commandOutput, err := cmd.Output()
	if err != nil {
		return err
	}

	jsonBody, err := typed.Json(commandOutput)
	if err != nil {
		return err
	}

	if jsonBody.String("error") != "" || jsonBody.String("errorMessage") != "" {
		return fmt.Errorf("Error code: %v message :%v", jsonBody.String("error"), jsonBody.String("errorMessage"))
	}

	resultBytes, err := jsonBody.ToBytes("response")

	if err != nil {
		return err
	}

	err = json.Unmarshal(resultBytes, outputs)

	if err != nil {
		return err
	}

	return nil
}

func dialViaCmdLine() (CommandLineClient, error) {
	return &commandLineClient{}, nil
}
