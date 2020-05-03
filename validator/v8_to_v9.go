package validator

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sync"
)

type NewValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (n *NewValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {

		n.lazyInit()

		if err := n.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (n *NewValidator) Engine() interface{} {
	n.lazyInit()
	return n.validate
}

func (n *NewValidator) lazyInit() {
	n.once.Do(func() {
		n.validate = validator.New()
		n.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

var _ binding.StructValidator = &NewValidator{}
