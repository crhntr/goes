package goes

type Runtimer interface {
	Global() Value
	Null() Value
	Undefined() Value
	ValueOf(x interface{}) Value
}

type (
	Booler interface {
		Bool() bool
	}
	Caller interface {
		Call(m string, args ...interface{}) Value
	}
	Floater interface {
		Float() float64
	}
	Getter interface {
		Get(p string) Value
	}
	Indexer interface {
		Index(i int) Value
	}
	InstanceOfer interface {
		InstanceOf(t Value) bool
	}
	Inter interface {
		Int() int
	}
	Invoker interface {
		Invoke(args ...interface{}) Value
	}
	Lengther interface {
		Length() int
	}
	Newer interface {
		New(args ...interface{}) Value
	}
	Setter interface {
		Set(p string, x interface{})
	}
	SetIndexer interface {
		SetIndex(i int, x interface{})
	}
	Stringer interface {
		String() string
	}
	Truther interface {
		Truthy() bool
	}
	Typer interface {
		Type() Type
	}
	Value interface {
		Booler
		Caller
		Floater
		Getter
		Indexer
		InstanceOfer
		Inter
		Invoker
		Lengther
		Newer
		Setter
		SetIndexer
		Stringer
		Truther
		Typer
	}
)

type Func interface {
	Releaser
}

type Function func(this Value, args []Value) interface{}

type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "undefined"
	case TypeNull:
		return "null"
	case TypeBoolean:
		return "boolean"
	case TypeNumber:
		return "number"
	case TypeString:
		return "string"
	case TypeSymbol:
		return "symbol"
	case TypeObject:
		return "object"
	case TypeFunction:
		return "function"
	default:
		panic("bad type")
	}
}

type Releaser interface {
	Release()
}

type TypedArray interface {
	Releaser
}

// func TypedArrayOf(slice interface{}) TypedArray {
// 	return nil
// }
