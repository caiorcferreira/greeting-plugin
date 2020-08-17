package greeting_plugin

import (
	"context"
	"github.com/b2wdigital/restQL-golang/pkg/restql"
)

const pluginName = "greeter"

func init() {
	restql.RegisterPlugin(func() restql.PluginInfo {
		return restql.PluginInfo{
			Name: pluginName,
			Type: restql.Lifecycle,
			New: func(logger restql.Logger) (restql.Plugin, error) {
				return NewGreetPlugin(logger)
			},
		}
	})
}

func NewGreetPlugin(logger restql.Logger) (restql.LifecyclePlugin, error) {
	return &GreetPlugin{log: logger}, nil
}

type GreetPlugin struct {
	log restql.Logger
}

func (g *GreetPlugin) BeforeTransaction(ctx context.Context, tr restql.TransactionRequest) context.Context {
	g.log.Info("WELCOME TO RESTQL")
	return ctx
}

func (g *GreetPlugin) AfterTransaction(ctx context.Context, tr restql.TransactionResponse) context.Context {
	return ctx
}

func (g *GreetPlugin) BeforeQuery(ctx context.Context, query string, queryCtx restql.QueryContext) context.Context {
	return ctx
}

func (g *GreetPlugin) AfterQuery(ctx context.Context, query string, result map[string]interface{}) context.Context {
	return ctx
}

func (g *GreetPlugin) BeforeRequest(ctx context.Context, request restql.HttpRequest) context.Context {
	return ctx
}

func (g *GreetPlugin) AfterRequest(ctx context.Context, request restql.HttpRequest, response restql.HttpResponse, err error) context.Context {
	return ctx
}

func (g *GreetPlugin) Name() string {
	return pluginName
}
