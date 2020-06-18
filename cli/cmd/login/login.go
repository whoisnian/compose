package login

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/docker/api/cli/mobycli"
	"github.com/docker/api/client"
	"github.com/docker/api/errdefs"
)

// Command returns the login command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login [OPTIONS] [SERVER] | login azure",
		Short: "Log in to a Docker registry",
		Long:  "Log in to a Docker registry or cloud backend.\nIf no registry server is specified, the default is defined by the daemon.",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runLogin,
	}
	// define flags for backward compatibility with com.docker.cli
	flags := cmd.Flags()
	flags.StringP("username", "u", "", "Username")
	flags.StringP("password", "p", "", "Password")
	flags.BoolP("password-stdin", "", false, "Take the password from stdin")

	return cmd
}

func runLogin(cmd *cobra.Command, args []string) error {
	if len(args) == 1 && !strings.Contains(args[0], ".") {
		backend := args[0]
		switch backend {
		case "azure":
			return cloudLogin(cmd, "aci")
		default:
			return errors.New("unknown backend type for cloud login: " + backend)
		}
	}
	return mobycli.ExecCmd(cmd)
}

func cloudLogin(cmd *cobra.Command, backendType string) error {
	ctx := cmd.Context()
	cs, err := client.GetCloudService(ctx, backendType)
	if err != nil {
		return errors.Wrap(errdefs.ErrLoginFailed, "cannot connect to backend")
	}
	err = cs.Login(ctx, nil)
	if errors.Is(err, context.Canceled) {
		return errors.New("login canceled")
	}
	if err != nil {
		return err
	}
	fmt.Println("login succeeded")
	return nil
}