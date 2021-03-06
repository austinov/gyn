package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/fatih/structs"
	"github.com/nguyenthenguyen/docx"
)

type FillDocxCallback func(appointment interface{}, doc *docx.Docx) error

func FillDocx(appointment interface{}, templatePath string, cb FillDocxCallback) (*os.File, error) {
	doc, err := docx.ReadDocxFile(templatePath)
	if err != nil {
		return nil, err
	}
	defer doc.Close()
	tmpDocx := doc.Editable()
	// retrieve all fields of struct
	fields := structs.Fields(appointment)
	for _, f := range fields {
		processField(f, tmpDocx)
	}
	// execute callback to fill some fields
	if cb != nil {
		if err = cb(appointment, tmpDocx); err != nil {
			return nil, err
		}
	}
	// create temporary file to save filled document
	file, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		return nil, err
	}
	tmpDocx.WriteToFile(file.Name())
	return file, nil
}

func processField(field *structs.Field, doc *docx.Docx) error {
	if field.IsEmbedded() {
		fields := structs.Fields(field.Value())
		for _, f := range fields {
			processField(f, doc)
		}
	}
	// fill document only for field with tag 'docx'
	if tag := field.Tag("docx"); tag != "" {
		val, err := decode(field.Value(), field.Kind())
		if err != nil {
			return fmt.Errorf("decode error of field %s: %#v", field.Name(), err)
		}
		doc.Replace(tag, val, -1)
	}
	return nil
}

func decode(data interface{}, dataKind reflect.Kind) (string, error) {
	switch dataKind {
	case reflect.String:
		return data.(string), nil
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		return fmt.Sprintf("%d", data), nil
	case reflect.Float32:
	case reflect.Float64:
		return fmt.Sprintf("%f", data), nil
	case reflect.Bool:
		return fmt.Sprintf("%t", data), nil
	}
	return "", fmt.Errorf("unsupported type: %s", dataKind)
}
