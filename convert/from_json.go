package convert

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

func FromJson(b []byte, baseStructName string) (string, error) {
	// trim space
	s := strings.TrimSpace(string(b))

	// unmarshal from json
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal from json: %v", err)
	}

	// convert
	var objs []structObject
	traverse(baseStructName, m, &objs)

	// printout
	buf := new(strings.Builder)
	for i, obj := range objs {
		if i != 0 {
			buf.WriteString("\n")
		}

		buf.WriteString(obj.String())
	}

	return strings.TrimSpace(buf.String()), nil
}

type structObject struct {
	name   string
	fields []field
}

func (obj structObject) String() string {
	buf := new(strings.Builder)

	// header
	buf.WriteString(fmt.Sprintf("type %s struct {\n", obj.name))

	// body
	for _, fi := range obj.fields {
		buf.WriteString(fmt.Sprintf("\t%s\n", fi.String()))
	}

	// footer
	buf.WriteString("}\n")

	return buf.String()
}

type field struct {
	name string
	typ  string
	skip bool
}

func (fi field) String() string {
	if fi.skip {
		return fmt.Sprintf("// skipped field '%s': insufficient clue for field ", fi.name)
	}

	return fmt.Sprintf("%s %s `json:\"%s\"`", strcase.ToCamel(fi.name), fi.typ, fi.name)
}

func traverse(baseName string, m map[string]interface{}, objs *[]structObject) {
	var so structObject
	so.name = strcase.ToCamel(baseName)

	for k, v := range m {
		var fieldType string
		var skip bool

		switch c := v.(type) {
		case int:
			fieldType = "int"
		case float64:
			fieldType = "float64"
		case string:
			fieldType = "string"
		case map[string]interface{}:
			subName := strcase.ToCamel(k)
			traverse(subName, c, objs)
			fieldType = subName
		case []interface{}:
			if len(c) < 1 {
				skip = true
			} else {
				first := c[0]
				switch first.(type) {
				case int:
					fieldType = "[]int"
				case float64:
					fieldType = "[]float64"
				case string:
					fieldType = "[]string"
				default:
					skip = true
				}
			}
		default:
			skip = true
		}

		so.fields = append(so.fields, field{
			name: k,
			typ:  fieldType,
			skip: skip,
		})
	}

	*objs = append(*objs, so)
}
