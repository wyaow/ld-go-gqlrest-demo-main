package i18n

import (
	"context"
	"errors"
	"fmt"

	ki18n "github.com/kataras/i18n"
)

var instance *ki18n.I18n

func init() {
	var err error
	instance, err = ki18n.New(ki18n.Glob("./config/locales/**/*", ki18n.LoaderConfig{}), "en-US", "zh-CN")
	if err != nil {
		panic(err)
	}

	instance.SetDefault("zh-CN")
}

func T(ctx context.Context, format string, args ...interface{}) string {
	lang, ok := ctx.Value(LangKey).(string)
	if !ok || lang == "" {
		lang = "zh-CN"
	}

	translated := instance.Tr(lang, format, args...)
	if translated == "" {
		translated = fmt.Sprintf(format, args...)
	}
	return translated
}

func ErrorT(ctx context.Context, format string, args ...interface{}) error {
	return errors.New(T(ctx, format, args...))
}
