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

func hideSensitiveData(ptrRef reflect.Value) interface{} {
	//if v == nil {
	//	return nil
	//}
	//ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Ptr {
		v := ptrRef.Interface()
		ptrRef = reflect.New(reflect.TypeOf(v))
		ptrRef.Elem().Set(reflect.ValueOf(v))
		fmt.Println("making a copy", ptrRef.Elem())
	}
	ref := ptrRef.Elem()
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		tag := refType.Field(i).Tag.Get("log")
		currField := ref.Field(i)
		if tag == "hide" || tag == "false" {
			if currField.CanSet() {
				redactField(&ref, i)
			}
		}
		fieldType := currField.Type()
		fieldKind := fieldType.Kind()

		if fieldKind == reflect.Struct {
			hideSensitiveData(currField.Addr())
		} else if fieldKind == reflect.Ptr && currField.Elem().Kind() == reflect.Struct {
			// making a copy so that original data do not get changed
			//newCopy := reflect.New(fieldType.Elem()).Elem()
			//newCopy.Set(currField.Elem())
			//currField.Set()
			hideSensitiveData(currField.Elem())
			//currField.Set(newCopy.Addr())
		}
	}
	return ref.Interface()
}

func (e *HideSensitiveFieldsEncoder) EncodeEntry(
	entry zapcore.Entry,
	fields []zapcore.Field,
) (*buffer.Buffer, error) {
	for idx, field := range fields {
		currInterface := field.Interface
		if field.Type == 23 && currInterface != nil {
			value := reflect.ValueOf(currInterface)
			kind := value.Kind()
			if kind == reflect.Struct {
				fields[idx].Interface = hideSensitiveData(value)
			} else if value.Elem().Kind() == reflect.Struct { // in case ptr is passed in the log
				// passes only value so that original struct do not get changed
				fields[idx].Interface = hideSensitiveData(value.Elem())
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

func CustomZap(l *zap.SugaredLogger, a GitRegistry, b *DockerArtifactStoreBean, c *GitHostRequest, d *Test, e *ComprehensiveStruct, f StructOne) {
	//l.Infow("Info", "A :", a, "B :", b, "C :", c, "D :", d, "E :", e, "F: ", f, "Test 2 :", Test2{MyMap: nil, Test3: getNil()})
	l.Infow("Info", "A : ", a)
	l.Infow("Info", "E: ", e)
	fmt.Println("A is : ", a)
	fmt.Println("E is : ", e)
	fmt.Println("Nested Struct in E is : ", (*e).StructPointer)
	// l.Warnw("Warning", "A :", a, "B :", b, "C :", c, "D :", d, "E :", nil)
	// l.Errorw("Error", "A :", a, "B :", b, "C :", c, "D :", d, "E :", nil)
}

// func main() {
// 	l, _ := NewSugardLogger("custom")
// 	a := GitRegistry{}
// 	faker.FakeData(&a)
// 	l.Infow("Error", "KEY", a)
// }
