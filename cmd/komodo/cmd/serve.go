package cmd

import (
	"github.com/kruemelmann/komodo/pkg/watcher"
	"github.com/kruemelmann/komodo/pkg/web"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringP("filename", "f", "", "Specify the filename of the LaTex File")
}

var (
	host = "0.0.0.0"
	port = "9090"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the webserver to serve the pdf file to a browser",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fname, _ := cmd.Flags().GetString("filename")

		go watcher.WatchFile(fname, buildCommand)
		web.StartServer(host, port, fname)
	},
}
