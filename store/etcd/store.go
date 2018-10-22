package etcd

import (
	"context"
	rc "github.com/kihamo/runtime-config"
	"github.com/kihamo/runtime-config/config"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
	"strconv"
)

type Store struct {
	client *clientv3.Client
}

func NewStore(client *clientv3.Client) *Store {
	return &Store{
		client: client,
	}
}

func (s *Store) VersionList(context.Context) ([]config.Version, error) {
	return nil, config.ErrorNotImplemented
}

func (s *Store) VersionCreate(context.Context, config.Version) error {
	return config.ErrorNotImplemented
}

func (s *Store) VersionRead(ctx context.Context, version config.Version) (config.Version, error) {
	path, err := getPathForVersion(version.ProjectID(), version.ID())
	if err != nil {
		return nil, err
	}

	response, err := s.client.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	if response.Count == 0 {
		return nil, config.ErrorVersionNotFound
	}

	return version, nil
}

func (s *Store) VersionUpdate(context.Context, config.Version) error {
	return config.ErrorNotImplemented
}

func (s *Store) VersionDelete(context.Context, config.Version) error {
	return config.ErrorNotImplemented
}

func (s *Store) VersionSetChangeCallback(func(config.Version, config.Version)) {

}

func (s *Store) VariableList(ctx context.Context, version config.Version) ([]config.Variable, error) {
	path, err := getPathForVersion(version.ProjectID(), version.ID())
	if err != nil {
		return nil, err
	}

	response, err := s.client.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	variables := make([]config.Variable, response.Count, response.Count)

	for i, kv := range response.Kvs {
		variables[i] = rc.NewVariable(string(kv.Key), string(kv.Value))
	}

	return variables, nil
}

func (s *Store) VariableCreate(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) VariableRead(ctx context.Context, version config.Version, variable config.Variable) (config.Variable, error) {
	path, err := getPathForVariable(version.ProjectID(), version.ID(), variable.Name())
	if err != nil {
		return nil, err
	}

	response, err := s.client.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	if response.Count == 0 {
		return nil, config.ErrorVariableNotFound
	}

	return variable, err
}

func (s *Store) VariableUpdate(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) VariableDelete(context.Context, config.Version, config.Variable) error {
	return config.ErrorNotImplemented
}

func (s *Store) VariableSetAnyChangeCallback(func(config.Variable, config.Value, config.Value)) {
}

func (s *Store) VariableSetKeyChangeCallback(func(config.Value, config.Value)) {

}

func getPathForVersion(projectID uint64, versionID string) (string, error) {
	if projectID == 0 {
		return "", errors.New("Project ID is zero")
	}

	if versionID == "" {
		return "", errors.New("Version ID is empty")
	}

	return "/config/" + strconv.FormatUint(projectID, 10) + "/" + versionID, nil
}

func getPathForVariable(projectID uint64, versionID, variableName string) (string, error) {
	if variableName == "" {
		return "", errors.New("Variable name is empty")
	}

	pathVersion, err := getPathForVersion(projectID, versionID)
	if err != nil {
		return "", err
	}

	return pathVersion + "/" + variableName, nil
}
