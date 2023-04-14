package main

import (
	"auth/pkg/server"
	"auth/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	port string
	l    utils.Logger
)

func getServeCommand() *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the auth service",
		Long:  `Serve the auth service:`,
		Run: func(cmd *cobra.Command, args []string) {
			s := server.NewServer(port, server.NewKeycloak())
			s.Listen(l)

			l.Infof("Serving auth service on port %s", port)
			quit := make(chan os.Signal)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit
		},
	}
	serveCmd.PersistentFlags().BoolVarP(&l.Debug, "debug", "", true, "Enable debug logging")
	serveCmd.PersistentFlags().StringVarP(&port, "port", "", "8080", "Port to serve the auth service on")
	return serveCmd
}
