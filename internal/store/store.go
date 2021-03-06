package store

import (
	"context"

	"github.com/kihamo/runtime-config/config"
)

// Store interface describes storage provider for configs
type Store interface {
	Versions(context.Context, string) ([]config.Version, error)
	Variables(context.Context, config.Version) ([]config.Variable, error)
	VariableCreate(context.Context, config.Version, config.Variable) error
	VariableRead(context.Context, config.Version, config.Variable) (config.Variable, error)
	VariableUpdate(context.Context, config.Version, config.Variable) error
	VariableDelete(context.Context, config.Version, config.Variable) error
	SetVersionChangeCallback(config.VersionChangeCallback) error
	SetVariableChangeCallback(config.Version, config.VariableChangeCallback) error
	SetVariableChangeByNameCallback(config.Version, string, config.VariableChangeCallback) error
}
