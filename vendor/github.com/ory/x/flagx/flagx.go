package flagx

import (
	"time"

	"github.com/ory/x/cmdx"
	"github.com/spf13/cobra"
)

func MustGetBool(cmd *cobra.Command, name string) bool {
	ok, err := cmd.Flags().GetBool(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return ok
}

func MustGetString(cmd *cobra.Command, name string) string {
	s, err := cmd.Flags().GetString(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return s
}

func MustGetDuration(cmd *cobra.Command, name string) time.Duration {
	d, err := cmd.Flags().GetDuration(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return d
}

func MustGetStringSlice(cmd *cobra.Command, name string) []string {
	ss, err := cmd.Flags().GetStringSlice(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return ss
}
