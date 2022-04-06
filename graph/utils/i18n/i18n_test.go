package i18n

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	ki18n "github.com/kataras/i18n"
	"github.com/stretchr/testify/require"
)

func FindPathForFile(fileName string) (string, error) {
	_, progName, _, _ := runtime.Caller(0)
	lastDir := path.Dir(progName)
	for {
		tryPath := filepath.Join(lastDir, fileName)
		if fi, err := os.Stat(tryPath); err == nil {
			if mode := fi.Mode(); mode.IsRegular() {
				return lastDir, nil
			}
		}

		newDir := filepath.Dir(lastDir)
		if newDir == "/" || newDir == lastDir {
			return "", fmt.Errorf("file '%s' not found", fileName)
		}
		lastDir = newDir
	}
}

// https://pkg.go.dev/testing
func TestMain(m *testing.M) {
	// Write code here to run before tests
	root, err := FindPathForFile("main.go")
	if err != nil {
		panic(err)
	}

	err = os.Chdir(root)
	if err != nil {
		panic(err)
	}

	instance, err = ki18n.New(ki18n.Glob("./config/locales/**/*", ki18n.LoaderConfig{}), "en-US", "zh-CN")
	if err != nil {
		panic(err)
	}

	instance.SetDefault("zh-CN")

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestHelloI18N(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	ctx := context.Background()

	ctx = context.WithValue(ctx, LangKey, "en-US")
	en := T(ctx, "hello")
	require.Equal(t, "Aloha", en)

	ctx = context.WithValue(ctx, LangKey, "zh-CN")
	cn := T(ctx, "hello")
	require.Equal(t, "你好呀", cn)
}
