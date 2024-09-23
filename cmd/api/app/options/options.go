package options

import (
	"github.com/spf13/pflag"
	"github.com/yjwj-3/content-release/cmd/api/app/config"
)

type ApiOptions struct {
	App string
}

func NewOptions() (*ApiOptions, error) {
	return &ApiOptions{}, nil
}

func (o *ApiOptions) Flags() map[string]*pflag.FlagSet {
	flags := make(map[string]*pflag.FlagSet)
	flags["app"] = o.AddFlags("app")
	return flags
}

func (o *ApiOptions) Config() (*config.ApiConfig, error) {
	c := &config.ApiConfig{}
	c.App = o.App
	return c, nil
}

func (o *ApiOptions) AddFlags(name string) *pflag.FlagSet {
	set := pflag.NewFlagSet(name, pflag.ExitOnError)
	set.StringVar(&o.App, "app", "app", "web port")
	return set
}
