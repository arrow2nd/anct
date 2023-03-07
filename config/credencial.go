package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/arrow2nd/anct/api"
)

const filename = ".cred.toml"

// Config : 設定
type Config struct {
	dir string
}

// New : 作成
func New() (*Config, error) {
	dir, err := getDefaultDir()
	if err != nil {
		return nil, err
	}

	return &Config{
		dir: dir,
	}, nil
}

// Save : 保存
func (c *Config) Save(t *api.Token) error {
	buf := &bytes.Buffer{}

	if err := toml.NewEncoder(buf).Encode(t); err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	path := filepath.Join(c.dir, filename)
	if err := os.WriteFile(path, buf.Bytes(), os.ModePerm); err != nil {
		return fmt.Errorf("failed to save (%s): %w", path, err)
	}

	return nil
}

// Load : 読み込み
func (c *Config) Load() (*api.Token, error) {
	path := filepath.Join(c.dir, filename)
	if _, err := os.Stat(path); err != nil {
		if err := c.createNewFile(); err != nil {
			return nil, err
		}
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load (%s): %w", path, err)
	}

	token := &api.Token{}
	if err := toml.Unmarshal(buf, token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal (%s): %w", path, err)
	}

	return token, nil
}

func (c *Config) createNewFile() error {
	return c.Save(&api.Token{
		Client: &api.ClientToken{},
		User: &api.UserToken{
			Bearer: "",
		},
	})
}

func getDefaultDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("failed to get home directory")
	}

	homeDir = filepath.Join(homeDir, ".config", "anct")

	// ディレクトリが無いなら作成
	if _, err := os.Stat(homeDir); err != nil {
		if err := os.MkdirAll(homeDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	return homeDir, nil
}
