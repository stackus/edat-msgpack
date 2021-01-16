package msgpack

import (
	"reflect"
	"sync"

	"github.com/shamaton/msgpack"

	"github.com/stackus/edat/core"
	"github.com/stackus/edat/core/register_types"
)

func init() {
	core.RegisterDefaultMarshaller(newMsgPackMarshaller())
	registertypes.RegisterTypes()
}

type msgPackMarshaler struct {
	items map[string]reflect.Type
	mu    sync.Mutex
}

func newMsgPackMarshaller() *msgPackMarshaler {
	return &msgPackMarshaler{
		items: map[string]reflect.Type{},
		mu:    sync.Mutex{},
	}
}

func (*msgPackMarshaler) Marshal(v interface{}) ([]byte, error)      { return msgpack.Encode(v) }
func (*msgPackMarshaler) Unmarshal(data []byte, v interface{}) error { return msgpack.Decode(data, v) }
func (m *msgPackMarshaler) GetType(typeName string) reflect.Type     { return m.items[typeName] }
func (m *msgPackMarshaler) RegisterType(typeName string, v reflect.Type) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[typeName] = v
}
