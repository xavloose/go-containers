package containers

import "reflect"

// We can't declare these as constants so in order to keep them as such
// we will create functions to return these internal types
var (
	typeInt        = reflect.TypeOf((*int)(nil)).Elem()
	typeInt8       = reflect.TypeOf((*int8)(nil)).Elem()
	typeInt16      = reflect.TypeOf((*int16)(nil)).Elem()
	typeInt32      = reflect.TypeOf((*int32)(nil)).Elem()
	typeInt64      = reflect.TypeOf((*int64)(nil)).Elem()
	typeUint       = reflect.TypeOf((*uint)(nil)).Elem()
	typeUint8      = reflect.TypeOf((*uint8)(nil)).Elem()
	typeUint16     = reflect.TypeOf((*uint16)(nil)).Elem()
	typeUint32     = reflect.TypeOf((*uint32)(nil)).Elem()
	typeUint64     = reflect.TypeOf((*uint64)(nil)).Elem()
	typeFloat32    = reflect.TypeOf((*float32)(nil)).Elem()
	typeFloat64    = reflect.TypeOf((*float64)(nil)).Elem()
	typeComplex64  = reflect.TypeOf((*complex64)(nil)).Elem()
	typeComplex128 = reflect.TypeOf((*complex128)(nil)).Elem()
	typeByte       = reflect.TypeOf((*byte)(nil)).Elem()
)

func TypeInt() reflect.Type {
	return typeInt
}

func TypeInt8() reflect.Type {
	return typeInt8
}

func TypeInt16() reflect.Type {
	return typeInt16
}

func TypeInt32() reflect.Type {
	return typeInt32
}

func TypeInt64() reflect.Type {
	return typeInt64
}

func TypeUint() reflect.Type {
	return typeUint
}

func TypeUint8() reflect.Type {
	return typeUint8
}

func TypeUint16() reflect.Type {
	return typeUint16
}

func TypeUint32() reflect.Type {
	return typeUint32
}

func TypeUint64() reflect.Type {
	return typeUint64
}

func TypeFloat32() reflect.Type {
	return typeFloat32
}

func TypeFloat64() reflect.Type {
	return typeFloat64
}

func TypeComplex64() reflect.Type {
	return typeComplex64
}

func TypeComplex128() reflect.Type {
	return typeComplex128
}

func TypeByte() reflect.Type {
	return typeByte
}
