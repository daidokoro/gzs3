package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// root command (calls all other commands)
var gzs3Cmd = &cobra.Command{
	Use:     "gzs3",
	Short:   "Clone git Repo & ZIP to AWS S3...",
	Example: "gzs3 git@github.com/some/repo.git",
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	// define session
	//
	// },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		// create sessoin
		sess, err := manager.GetSess(profile)
		handleError(err)

		// define repo
		gituri = args[0]
		gitrepo, err := newrepo(gituri, gituser)
		handleError(err)

		// write file to s3
		if _, err = S3write(
			gitrepo.conf.Bucket,
			gitrepo.conf.Key,
			gitrepo.zipData,
			sess,
		); err != nil {
			handleError(err)
		}

		fmt.Printf("zip created in s3: %s\n",
			log.ColorString(
				fmt.Sprintf("s3://%s/%s", gitrepo.conf.Bucket, gitrepo.conf.Key),
				"red",
			),
		)

	},
}

// define flags
func init() {
	gzs3Cmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "configured AWS profile")
	gzs3Cmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "debug mode")
	gzs3Cmd.PersistentFlags().BoolVarP(&colors, "no-colors", "", false, "disable color output")
	gzs3Cmd.PersistentFlags().StringVarP(&gituser, "user", "u", "", "git username")
	gzs3Cmd.PersistentFlags().StringVarP(&gitpass, "password", "", "", "git password")
	gzs3Cmd.PersistentFlags().StringVarP(&gitrsa, "ssh-rsa", "i", filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa"), "path to git SSH id_rsa")
}
