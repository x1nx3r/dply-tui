package main

import (
	"os"

	cli "github.com/dionisius77/dply/dply/handler/cli"
	cli_config "github.com/dionisius77/dply/dply/handler/cli/config"
	cli_deploy "github.com/dionisius77/dply/dply/handler/cli/deploy"
	cli_image "github.com/dionisius77/dply/dply/handler/cli/image"
	cli_project "github.com/dionisius77/dply/dply/handler/cli/project"
	cli_spec "github.com/dionisius77/dply/dply/handler/cli/spec"
	"github.com/spf13/cobra"
)

// for bump version lagi
func main() {

	rootCmd := &cobra.Command{
		Use:   "dplyon",
		Short: "CLI client for dply - k8s service deployment management",
		Long:  "dplyon is the CLI client for dply, a custom k8s service deployment management tool. It registers container images, manages deployment specs (environment variables, ports, scaling, affinity) and deploys services into Kubernetes namespaces.",
	}
	rootCmd.AddCommand(cli.NewCmdStatus().Command)
	rootCmd.AddCommand(cli.NewCmdLogin().Command)
	rootCmd.AddCommand(cli.NewCmdLogout().Command)
	rootCmd.AddCommand(cli_config.New().Command)
	rootCmd.AddCommand(cli_project.New().Command)

	rootCmd.AddCommand(cli_image.New().Command)
	rootCmd.AddCommand(cli_spec.New().Command)
	rootCmd.AddCommand(cli_deploy.New().Command)

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}

}
