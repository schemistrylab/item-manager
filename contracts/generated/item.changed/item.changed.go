// Code generated by github.com/schemistrylab/schemistry. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avro
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Simple struct {
	Uuid string `json:"uuid"`
}

const SimpleAvroCRC64Fingerprint = "\xd5e\xee\x9a\xdee\xf6\xbf"

func NewSimple() Simple {
	r := Simple{}
	return r
}

func DeserializeSimple(r io.Reader) (Simple, error) {
	t := NewSimple()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSimpleFromSchema(r io.Reader, schema string) (Simple, error) {
	t := NewSimple()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSimple(r Simple, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Uuid, w)
	if err != nil {
		return err
	}
	return err
}

func (r Simple) Serialize(w io.Writer) error {
	return writeSimple(r, w)
}

func (r Simple) Schema() string {
	return "{\"fields\":[{\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"org.itemmanager.avro.simple\",\"type\":\"record\"}"
}

func (r Simple) SchemaName() string {
	return "org.itemmanager.avro.simple"
}

func (_ Simple) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Simple) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Simple) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Simple) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Simple) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Simple) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Simple) SetString(v string)   { panic("Unsupported operation") }
func (_ Simple) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Simple) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Uuid}

		return w

	}
	panic("Unknown field index")
}

func (r *Simple) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Simple) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Simple) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Simple) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Simple) HintSize(int)                     { panic("Unsupported operation") }
func (_ Simple) Finalize()                        {}

func (_ Simple) AvroCRC64Fingerprint() []byte {
	return []byte(SimpleAvroCRC64Fingerprint)
}

func (r Simple) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["uuid"], err = json.Marshal(r.Uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Simple) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["uuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Uuid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for uuid")
	}
	return nil
}
