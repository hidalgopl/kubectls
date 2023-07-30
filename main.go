package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var Colors = []string{Red, Gray, Blue, Purple, Cyan, White, Green, Yellow}
var NotSupportedSubCommands = map[string]struct{}{"edit": {}, "port-forward": {}, "attach": {}, "exec": {}}

func main() {

	var contextFlag string
	flag.StringVar(&contextFlag, "context", "", "")
	flag.Parse()
	kubectlArgsLen := len(os.Args) - 2
	kubectlArgs := make([]string, kubectlArgsLen)
	k := 0
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--context") {
			continue
		}
		if _, notSupported := NotSupportedSubCommands[arg]; notSupported {
			fmt.Printf("subcommand %s is currently not supported.\n", arg)
			os.Exit(1)
		}
		kubectlArgs[k] = arg
		k++
	}

	for i, k8sCtx := range strings.Split(contextFlag, ",") {
		currentColor := Colors[i]
		cmd := buildKubectlCommandForContext(k8sCtx, kubectlArgs)
		fmt.Printf("%s---------------------------------------------------------------------------------------\n", currentColor)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("got errors: %w", err)
		}
		fmt.Println(string(output), "\n---------------------------------------------------------------------------------------", Reset)
	}
	os.Exit(0)
}

func buildKubectlCommandForContext(k8sCtx string, kubectlArgs []string) *exec.Cmd {
	kubectlCmdArgs := []string{"--context", k8sCtx}
	kubectlCmdArgs = append(kubectlCmdArgs, kubectlArgs...)
	cmd := exec.Command("kubectl", kubectlCmdArgs...)
	cmd.Env = os.Environ()
	return cmd
}
