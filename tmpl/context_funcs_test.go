package tmpl

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	expected := `foo:
  bar: BAR
`
	expectedFilename := "values.yaml"
	ctx := &Context{basePath: ".", readFile: func(filename string) ([]byte, error) {
		if filename != expectedFilename {
			return nil, fmt.Errorf("unexpected filename: expected=%v, actual=%s", expectedFilename, filename)
		}
		return []byte(expected), nil
	}}
	actual, err := ctx.ReadFile(expectedFilename)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}

func TestReadFile_PassAbsPath(t *testing.T) {
	expected := `foo:
  bar: BAR
`
	expectedFilename, _ := filepath.Abs("values.yaml")
	ctx := &Context{basePath: ".", readFile: func(filename string) ([]byte, error) {
		if filename != expectedFilename {
			return nil, fmt.Errorf("unexpected filename: expected=%v, actual=%s", expectedFilename, filename)
		}
		return []byte(expected), nil
	}}
	actual, err := ctx.ReadFile(expectedFilename)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}

func TestToYaml(t *testing.T) {
	expected := `foo:
  bar: BAR
`
	vals := Values(map[string]interface{}{
		"foo": map[interface{}]interface{}{
			"bar": "BAR",
		},
	})
	actual, err := ToYaml(vals)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}

func TestFromYaml(t *testing.T) {
	raw := `foo:
  bar: BAR
`
	expected := Values(map[string]interface{}{
		"foo": map[interface{}]interface{}{
			"bar": "BAR",
		},
	})
	actual, err := FromYaml(raw)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}

func TestSetValueAtPath_OneComponent(t *testing.T) {
	input := map[string]interface{}{
		"foo": "",
	}
	expected := map[string]interface{}{
		"foo": "FOO",
	}
	actual, err := SetValueAtPath("foo", "FOO", input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}

func TestSetValueAtPath_TwoComponents(t *testing.T) {
	input := map[string]interface{}{
		"foo": map[interface{}]interface{}{
			"bar": "",
		},
	}
	expected := map[string]interface{}{
		"foo": map[interface{}]interface{}{
			"bar": "FOO_BAR",
		},
	}
	actual, err := SetValueAtPath("foo.bar", "FOO_BAR", input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}

func TestTpl(t *testing.T) {
	text := `foo: {{ .foo }}
`
	expected := `foo: FOO
`
	ctx := &Context{basePath: "."}
	actual, err := ctx.Tpl(text, map[string]interface{}{"foo": "FOO"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: expected=%v, actual=%v", expected, actual)
	}
}
