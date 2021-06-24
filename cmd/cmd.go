package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func generateCommand(options *terraform.Options, args ...string) shell.Command {
	cmd := shell.Command{
		Command:    options.TerraformBinary,
		Args:       args,
		WorkingDir: options.TerraformDir,
		Env:        options.EnvVars,
		Logger:     options.Logger,
	}

	return cmd
}

func RunTerraformCommand(additionalOptions *terraform.Options, args ...string) string {
	out, err := RunTerraformCommandE(additionalOptions, args...)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return out
}

func RunTerraformCommandE(additionalOptions *terraform.Options, additionalArgs ...string) (string, error) {
	options, args := terraform.GetCommonOptions(additionalOptions, additionalArgs...)

	cmd := generateCommand(options, args...)
	return RunCommandAndGetOutputE(cmd)
}

func RunCommandAndGetOutput(command shell.Command) string {
	output, err := RunCommandAndGetOutputE(command)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	return output
}

func RunCommandAndGetOutputE(command shell.Command) (string, error) {
	output, err := runCommand(command)
	if err != nil {
		return output.Combined(), &ErrWithCmdOutput{err, output}
	}
	return output.Combined(), nil
}

func GetExitCodeForTerraformCommand(additionalOptions *terraform.Options, args ...string) int {
	exitCode, err := GetExitCodeForTerraformCommandE(additionalOptions, args...)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	return exitCode
}

func GetExitCodeForTerraformCommandE(additionalOptions *terraform.Options, additionalArgs ...string) (int, error) {
	options, args := terraform.GetCommonOptions(additionalOptions, additionalArgs...)

	logger.Infof("Running %s with args %v", options.TerraformBinary, args)
	cmd := generateCommand(options, args...)
	_, err := RunCommandAndGetOutputE(cmd)
	if err == nil {
		return DefaultSuccessExitCode, nil
	}

	exitCode, getExitCodeErr := shell.GetExitCodeForRunCommandError(err)
	if getExitCodeErr == nil {
		return exitCode, nil
	}

	return DefaultErrorExitCode, getExitCodeErr
}

func runCommand(command shell.Command) (*output, error) {
	logger.Infof("Running command %s with args %s", command.Command, command.Args)

	cmd := exec.Command(command.Command, command.Args...)

	cmd.Dir = command.WorkingDir
	cmd.Stdin = os.Stdin
	cmd.Env = formatEnvVars(command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	output, err := readStdoutAndStderr(stdout, stderr)
	if err != nil {
		return output, err
	}

	return output, cmd.Wait()
}

type ErrWithCmdOutput struct {
	Underlying error
	Output     *output
}

func (e *ErrWithCmdOutput) Error() string {
	return fmt.Sprintf("error while running command: %v; %s", e.Underlying, e.Output.Stderr())
}

func readStdoutAndStderr(stdout, stderr io.ReadCloser) (*output, error) {
	out := newOutput()
	stdoutReader := bufio.NewReader(stdout)
	stderrReader := bufio.NewReader(stderr)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	var stdoutErr, stderrErr error
	go func() {
		defer wg.Done()
		stdoutErr = readData(stdoutReader, out.stdout)
	}()
	go func() {
		defer wg.Done()
		stderrErr = readData(stderrReader, out.stderr)
	}()
	wg.Wait()

	if stdoutErr != nil {
		return out, stdoutErr
	}

	if stderrErr != nil {
		return out, stderrErr
	}

	return out, nil
}

func readData(reader *bufio.Reader, writer io.StringWriter) error {
	var line string
	var readErr error
	for {
		line, readErr = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")

		if len(line) == 0 && readErr == io.EOF {
			break
		}

		logger.Infof(line)
		if _, err := writer.WriteString(line); err != nil {
			return err
		}

		if readErr != nil {
			break
		}
	}

	if readErr != io.EOF {
		return readErr
	}

	return nil
}

func formatEnvVars(command shell.Command) []string {
	env := os.Environ()
	for key, value := range command.Env {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	return env
}
