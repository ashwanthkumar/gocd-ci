package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const SpecFile = ".gocd-ci.yml"

var ciCommand = &cobra.Command{
	Use:   "gocd-ci",
	Short: "Run your CI pipeline from " + SpecFile,
	Long:  `Run your CI pipeline from ` + SpecFile,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Trying to open %s\n", SpecFile)
		contents, err := ioutil.ReadFile(SpecFile)
		if err != nil {
			handleError(cmd, err)
		}
		fmt.Printf("Parsing the %s as YAML\n", SpecFile)
		spec, err := DecodeCISpecFromBytes(contents)
		if err != nil {
			handleError(cmd, err)
		}
		// Set all the Env values
		for key, value := range spec.Env {
			fmt.Printf("Setting Environment variable '%s' to '%s'\n", key, value)
			os.Setenv(key, value)
		}

		var hasAtleastOneCommandFailed bool
		for _, command := range spec.Cmd {
			fmt.Printf("Executing %s\n", command)
			actualCmd := exec.Command("sh", "-c", command)
			actualCmd.Stdout = os.Stdout
			actualCmd.Stderr = os.Stderr
			err = actualCmd.Start()
			if err != nil {
				hasAtleastOneCommandFailed = true
				fmt.Printf("%v\n", err)
				continue
			}
			err = actualCmd.Wait()
			if err != nil {
				hasAtleastOneCommandFailed = true
				fmt.Printf("%v\n", err)
				continue
			}
		}

		if hasAtleastOneCommandFailed {
			fmt.Printf("Some errors were found while building. Check the above logs.\n")
			os.Exit(1)
		}
	},
}

func handleError(cmd *cobra.Command, err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
