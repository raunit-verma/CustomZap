package main

import (
	"fmt"
	"net/http"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func GetLogger() *zap.SugaredLogger {
	return logger
}

type LogConfig struct {
	Level   int  `env:"LOG_LEVEL" envDefault:"0"` // default info
	DevMode bool `env:"LOGGER_DEV_MODE" envDefault:"false"`
}

type HideSensitiveFieldsEncoder struct {
	zapcore.Encoder
	cfg zapcore.EncoderConfig
}

func redactField(ref *reflect.Value, i int) {
	refField := ref.Field(i)
	newValue := reflect.New(refField.Type()).Elem()
	fieldType := ref.Field(i).Type().Kind()
	switch fieldType {
	case reflect.String:
		newValue.SetString("[REDACTED]")
	}
	ref.Field(i).Set(newValue)
}

func hideSensitiveData(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Ptr {
		ptrRef = reflect.New(reflect.TypeOf(v))
		ptrRef.Elem().Set(reflect.ValueOf(v))
	}
	ref := ptrRef.Elem()
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		tag := refType.Field(i).Tag.Get("log")
		if tag == "hide" || tag == "false" {
			if ref.Field(i).CanSet() {
				redactField(&ref, i)
			}
		}
		fieldType := ref.Field(i).Type().Kind()
		if fieldType == reflect.Struct {
			hideSensitiveData(ref.Field(i).Addr().Interface())
		} else if fieldType == reflect.Ptr && ref.Field(i).Elem().Kind() == reflect.Struct {
			// making a copy so that original data do not get changed
			newCopy := reflect.New(ref.Field(i).Type().Elem()).Elem()
			newCopy.Set(ref.Field(i).Elem())
			hideSensitiveData(newCopy.Addr().Interface())
			ref.Field(i).Set(newCopy.Addr())
		}
	}
	return ref.Interface()
}

func (e *HideSensitiveFieldsEncoder) EncodeEntry(
	entry zapcore.Entry,
	fields []zapcore.Field,
) (*buffer.Buffer, error) {
	for idx, field := range fields {
		if field.Type == 23 && field.Interface != nil {
			value := reflect.ValueOf(field.Interface)
			kind := value.Kind()
			if kind == reflect.Struct {
				fields[idx].Interface = hideSensitiveData(field.Interface)
			} else if value.Elem().Kind() == reflect.Struct {
				// passes only value so that original struct do not get changed
				fields[idx].Interface = hideSensitiveData(value.Elem().Interface())
			}
		}
	}
	// return e.Encoder.EncodeEntry(entry, fields)
	return e.Encoder.EncodeEntry(entry, fields)
}

func newHideSensitiveFieldsEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	encoder := zapcore.NewConsoleEncoder(config)
	return &HideSensitiveFieldsEncoder{encoder, config}
}

func InitLogger(s string) (*zap.SugaredLogger, error) {
	_ = zap.RegisterEncoder("hideSensitiveData", func(config zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return newHideSensitiveFieldsEncoder(config), nil
	})
	config := zap.NewProductionConfig()
	// Will be added in else section for Prod mode only
	if s == "custom" {
		config.Encoding = "hideSensitiveData"
	}

	l, err := config.Build()
	if err != nil {
		fmt.Println("failed to create the default logger: " + err.Error())
		return nil, err
	}
	logger = l.Sugar()
	return logger, nil
}

func NewSugardLogger(s string) (*zap.SugaredLogger, error) {
	return InitLogger(s)
}

func NewHttpClient() *http.Client {
	return http.DefaultClient
}

func getNil() *Test3 {
	return nil
}

func CustomZap(l *zap.SugaredLogger, a *GitRegistry, b *DockerArtifactStoreBean, c *GitHostRequest, d *Test) {
	l.Infow("Info", "A :", a, "B :", b, "C :", c, "D :", d, "E :", nil, "F: ", Test2{MyMap: nil, Test3: getNil()})
	l.Warnw("Warning", "A :", a, "B :", b, "C :", c, "D :", d, "E :", nil)
	l.Errorw("Error", "A :", a, "B :", b, "C :", c, "D :", d, "E :", nil)
}

// func main() {
// 	l, _ := NewSugardLogger("custom")
// 	a := GitRegistry{}
// 	faker.FakeData(&a)
// 	l.Infow("Error", "KEY", a)
// }
