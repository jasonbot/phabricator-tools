package phabricatortools

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/karlseguin/typed"
)

// CommandLineClient implements a very simple Call operator which shells out to `arc`.
type CommandLineClient interface {
	Call(method string, inputs interface{}, outputs interface{}) error
}

type commandLineClient struct {
}

func (*commandLineClient) Call(method string, inputs interface{}, outputs interface{}) error {
	cmd := exec.Command("arc", "call-conduit", method)

	// We want to pipe in our request JSON via stdin
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

	//fmt.Printf("Sending: %v\n", string(stdinContent))

	// Now wait for arc to run
	cmd.Wait()

	// Grab the JSON if the tool ran successfully
	commandOutput, err := cmd.Output()
	if err != nil {
		return err
	}
	//fmt.Printf("Getting: %v\n", string(commandOutput))

	jsonBody, err := typed.Json(commandOutput)
	if err != nil {
		return err
	}

	// If the call resulted in an error, fail correctly
	if jsonBody.String("error") != "" || jsonBody.String("errorMessage") != "" {
		return fmt.Errorf("Error calling %v: %v message :%v", method, jsonBody.String("error"), jsonBody.String("errorMessage"))
	}

	// Otherwise, treat response as whatever struct we got passed in to unmarshal to
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
