package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arrow2nd/anct/api"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	d := t.TempDir()
	cfg := &Config{dir: d}

	token := &api.Token{
		Client: &api.ClientToken{
			ID:     "id",
			Secret: "secret",
		},
		User: &api.UserToken{
			Bearer: "bearer",
		},
	}

	err := cfg.Save(token)
	assert.NoError(t, err)

	bytes, err := os.ReadFile(filepath.Join(d, filename))
	assert.NoError(t, err)

	want := `[Client]
  ID = "id"
  Secret = "secret"

[User]
  Bearer = "bearer"
`

	assert.Equal(t, want, string(bytes))
}

func TestLoad(t *testing.T) {
	t.Run("読み込めるか", func(t *testing.T) {
		toml := `[Client]
  ID = "id"
  Secret = "secret"

[User]
  Bearer = "bearer"
`

		d := t.TempDir()
		os.WriteFile(filepath.Join(d, filename), []byte(toml), os.ModePerm)

		cfg := &Config{dir: d}
		token, err := cfg.Load()
		assert.NoError(t, err)

		assert.Equal(t, "id", token.Client.ID)
		assert.Equal(t, "secret", token.Client.Secret)
		assert.Equal(t, "bearer", token.User.Bearer)
	})

	t.Run("ファイルがない場合に作成されるか", func(t *testing.T) {
		d := t.TempDir()
		cfg := &Config{dir: d}

		_, err := cfg.Load()
		assert.NoError(t, err)

		files, err := os.ReadDir(d)
		assert.NoError(t, err)

		assert.Len(t, files, 1)

		for i, f := range files {
			if n := f.Name(); n != filename {
				t.Errorf("files[%d] : %s != %s", i, n, filename)
			}
		}
	})
}
