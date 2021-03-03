package schema

import (
	"go/ast"
	"gooorm/dialect"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag string
}

// Schema represents a table of database
type Schema struct {
	Model interface{}
	Name string
	Fields []*Field
	FieldNames []string
	fieldMap map[string]*Field
}

func (s *Schema) GetField(name string) *Field{
	return s.fieldMap[name]
}

// Parse parses any given object to a Schema
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelTyp := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelTyp.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelTyp.NumField(); i++ {
		p := modelTyp.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("gooorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}

	return schema
}
