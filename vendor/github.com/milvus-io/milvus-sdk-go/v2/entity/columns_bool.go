// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate at 2021-07-19 17:30:03.688864552 +0800 CST m=+0.002944243

package entity 

import (
	"errors"

	"github.com/milvus-io/milvus-sdk-go/v2/internal/proto/schema"
)

// ColumnBool generated columns type for Bool
type ColumnBool struct {
	name   string
	values []bool
}

// Name returns column name
func (c *ColumnBool) Name() string {
	return c.name
}

// Type returns column FieldType
func (c *ColumnBool) Type() FieldType {
	return FieldTypeBool
}

// Len returns column values length
func (c *ColumnBool) Len() int {
	return len(c.values)
}

// FieldData return column data mapped to schema.FieldData
func (c *ColumnBool) FieldData() *schema.FieldData {
	fd := &schema.FieldData{
		Type: schema.DataType_Bool,
		FieldName: c.name,
	}
	data := make([]bool, 0, c.Len())
	for i := 0 ;i < c.Len(); i++ {
		data = append(data, bool(c.values[i]))
	}
	fd.Field = &schema.FieldData_Scalars{
		Scalars: &schema.ScalarField{
			Data: &schema.ScalarField_BoolData{
				BoolData: &schema.BoolArray{
					Data: data,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnBool) ValueByIdx(idx int) (bool, error) {
	var r bool // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// Data returns column data
func (c *ColumnBool) Data() []bool {
	return c.values
}

// NewColumnBool auto generated constructor
func NewColumnBool(name string, values []bool) *ColumnBool {
	return &ColumnBool {
		name: name,
		values: values,
	}
}
