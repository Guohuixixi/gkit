package option

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApply(t *testing.T) {
	app := &App{}
	Apply[App](app, app.WithName("test"), app.WithVersion("1.0.0"))
	assert.Equal(t, app, &App{Name: "test", Version: "1.0.0"})
}

func TestApplyError(t *testing.T) {
	app := &App{}
	err := ApplyError[App](app, app.WithNameError("test"), app.WithVersionError("1.0.0"))
	require.NoError(t, err)
	assert.Equal(t, app, &App{Name: "test", Version: "1.0.0"})
	err = ApplyError[App](app, app.WithNameError(""), app.WithVersionError("1.0.0"))
	assert.Equal(t, nameError, err)
	fmt.Println(app)
}

type App struct {
	Name    string
	Version string
}

func (a *App) WithName(name string) Option[App] {
	return func(a *App) {
		a.Name = name
	}
}

func (a *App) WithVersion(version string) Option[App] {
	return func(a *App) {
		a.Version = version
	}
}

func (a *App) WithNameError(name string) OptionError[App] {
	return func(a *App) error {
		if name == "" {
			return fmt.Errorf("名称不能为空")
		}
		a.Name = name
		return nil
	}
}

func (a *App) WithVersionError(version string) OptionError[App] {
	return func(a *App) error {
		if version == "" {
			return fmt.Errorf("版本不能为空")
		}
		a.Version = version
		return nil
	}
}

var nameError = fmt.Errorf("名称不能为空")
var versionError = fmt.Errorf("版本不能为空")
