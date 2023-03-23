package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Javlopez/kube-user/pkg/crypt"
	"github.com/Javlopez/kube-user/pkg/models"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csr",
	Short: "Create Certificate Signing Request(CSR)",
	Long:  `Will generate a  Certificate Signing Request (CSR)`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO this should be something like
		// opts := models.NewOptions(cmd)
		// opts.User()
		// opts.WithSomething()
		name, _ := cmd.Flags().GetString("name")
		key := crypt.New()
		opts := models.Options{
			User: name,
		}
		err := key.Build(opts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().String("name", "", "Name of the user")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// kubeuser
