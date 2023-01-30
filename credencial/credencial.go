package credencial

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

// Save : 保存
func Save(t *api.Token) error {
	buf := &bytes.Buffer{}

	if err := toml.NewEncoder(buf).Encode(t); err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	path := filepath.Join(configDir, filename)
	if err := os.WriteFile(path, buf.Bytes(), os.ModePerm); err != nil {
		return fmt.Errorf("failed to save (%s): %w", path, err)
	}

	return nil
}

// Load : 読み込み
func Load() (*api.Token, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(configDir, filename)
	if _, err := os.Stat(path); err != nil {
		if err := createNewFile(); err != nil {
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

func createNewFile() error {
	return Save(&api.Token{
		Client: &api.ClientToken{
			ID:     "",
			Secret: "",
		},
		User: &api.UserToken{
			Bearer: "",
		},
	})
}

func getConfigDir() (string, error) {
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
