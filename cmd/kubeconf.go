package cmd

import (
	"fmt"
	"os"

	"github.com/rwxrob/kubeconf/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	cli "k8s.io/cli-runtime/pkg/genericclioptions"
)

// set by goreleaser via ldflags
var version = `(version unknown)`

var rootCmd = &cobra.Command{

	// TODO handoff options and arguments

	Use:   `kubeconf`,
	Short: `Print api.Config from KUBECONFIG as JSON`,
	RunE: func(c *cobra.Command, args []string) error {
		conf, err := cli.NewConfigFlags(true).ToRawKubeConfigLoader().RawConfig()
		if err != nil {
			return err
		}
		internal.PrintAsJSON(conf)
		return nil
	},
}

var versionCmd = &cobra.Command{
	Use:   `version`,
	Short: `Print current tagged version`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(version)
	},
}

func Execute() {
	confflags := cli.NewConfigFlags(true)
	ownflags := pflag.NewFlagSet("kubeconf", pflag.ExitOnError)
	confflags.AddFlags(ownflags)
	pflag.CommandLine = ownflags
	rootCmd.AddCommand(versionCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
