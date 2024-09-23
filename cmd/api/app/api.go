package app

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/yjwj-3/content-release/cmd/api/app/config"
	"github.com/yjwj-3/content-release/cmd/api/app/options"
	"github.com/yjwj-3/content-release/internal/api"
	"k8s.io/klog/v2"
)

// NewApiCommand creates a *cobra.Command object with default parameters
func NewApiCommand() *cobra.Command {
	// options
	o, err := options.NewOptions()
	if err != nil {
		klog.Background().Error(err, "Unable to initialize command options")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
	// cmd
	cmd := &cobra.Command{
		Use:  "content-release-api",
		Long: `content-release-api aggregates the api interfaces provided by this project.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := o.Config()
			if err != nil {
				return err
			}
			return run(context.Background(), c)
		},
		Args: cobra.NoArgs,
	}

	// 读取flag
	fs := cmd.Flags()
	flagSets := o.Flags()
	for _, f := range flagSets {
		fs.AddFlagSet(f)
	}
	return cmd
}

func run(ctx context.Context, c *config.ApiConfig) error {
	s := api.NewApiServer(c)
	err := s.Serve()
	return err
}
