package utils

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strconv"
	"unicode/utf8"
	"unsafe"
)

func BytesToRune(byteArray []byte) []rune {
	runeArray := make([]rune, 0)
	for _, b := range byteArray {
		if b != 0 {
			r, _ := utf8.DecodeRune([]byte{b})
			runeArray = append(runeArray, r)
		}
	}
	return runeArray
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func JSON2PB(formJsonStr string, toPb proto.Message) error {
	// json字符串转pb
	return json.Unmarshal([]byte(formJsonStr), &toPb)
}

func PB2JSON(fromPb proto.Message, toJsonStr string) error {
	// pb转json字符串
	jsonStr, err := json.Marshal(fromPb)
	if err == nil {
		toJsonStr = string(jsonStr)
	}

	return err
}

// ToArray slice转为数组结构体
func ToArray(slice interface{}, dataType interface{}) interface{} {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		return nil
	}
	sliceLen := value.Len()
	typeValue := reflect.ValueOf(dataType)

	code := int32(0)
	switch value.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		code = 1
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		code = 2
	case reflect.Float64, reflect.Float32:
		code = 3
	case reflect.String:
		code = 4
	case reflect.Interface:
		code = 5
	default:
		return value.Interface()
	}

	switch typeValue.Kind() {
	case reflect.Int:
		r := make([]int, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int(value.Index(i).Int())
			case 2:
				r[i] = int(value.Index(i).Uint())
			case 3:
				r[i] = int(value.Index(i).Float())
			case 4:
				if v, err := strconv.Atoi(value.Index(i).String()); err == nil {
					r[i] = v
				}
			case 5:
				r[i] = value.Index(i).Interface().(int)
			}
		}
		return r
	case reflect.Int8:
		r := make([]int8, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int8(value.Index(i).Int())
			case 2:
				r[i] = int8(value.Index(i).Uint())
			case 3:
				r[i] = int8(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int8(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int8)
			}
		}
		return r
	case reflect.Int16:
		r := make([]int16, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int16(value.Index(i).Int())
			case 2:
				r[i] = int16(value.Index(i).Uint())
			case 3:
				r[i] = int16(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int16(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int16)
			}
		}
		return r
	case reflect.Int32:
		r := make([]int32, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = int32(value.Index(i).Int())
			case 2:
				r[i] = int32(value.Index(i).Uint())
			case 3:
				r[i] = int32(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = int32(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(int32)
			}
		}
		return r
	case reflect.Int64:
		r := make([]int64, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = value.Index(i).Int()
			case 2:
				r[i] = int64(value.Index(i).Uint())
			case 3:
				r[i] = int64(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseInt(value.Index(i).String(), 10, 64); err == nil {
					r[i] = v
				}
			case 5:
				r[i] = value.Index(i).Interface().(int64)
			}
		}
		return r
	case reflect.Uint:
		r := make([]uint, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint(value.Index(i).Int())
			case 2:
				r[i] = uint(value.Index(i).Uint())
			case 3:
				r[i] = uint(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint)
			}
		}
		return r
	case reflect.Uint8:
		r := make([]uint8, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint8(value.Index(i).Int())
			case 2:
				r[i] = uint8(value.Index(i).Uint())
			case 3:
				r[i] = uint8(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint8(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint8)
			}
		}
		return r
	case reflect.Uint16:
		r := make([]uint16, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint16(value.Index(i).Int())
			case 2:
				r[i] = uint16(value.Index(i).Uint())
			case 3:
				r[i] = uint16(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint16(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint16)
			}
		}
		return r
	case reflect.Uint32:
		r := make([]uint32, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint32(value.Index(i).Int())
			case 2:
				r[i] = uint32(value.Index(i).Uint())
			case 3:
				r[i] = uint32(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = uint32(v)
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint32)
			}
		}
		return r
	case reflect.Uint64:
		r := make([]uint64, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = uint64(value.Index(i).Int())
			case 2:
				r[i] = value.Index(i).Uint()
			case 3:
				r[i] = uint64(value.Index(i).Float())
			case 4:
				if v, err := strconv.ParseUint(value.Index(i).String(), 10, 64); err == nil {
					r[i] = v
				}
			case 5:
				r[i] = value.Index(i).Interface().(uint64)
			}
		}
		return r
	case reflect.Float32:
		r := make([]float32, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = float32(value.Index(i).Int())
			case 2:
				r[i] = float32(value.Index(i).Uint())
			case 3:
				r[i] = float32(value.Index(i).Float())
			case 4:
				if f, err := strconv.ParseFloat(value.Index(i).String(), 32); err == nil {
					r[i] = float32(f)
				}
			case 5:
				r[i] = value.Index(i).Interface().(float32)
			}
		}
		return r
	case reflect.Float64:
		r := make([]float64, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = float64(value.Index(i).Int())
			case 2:
				r[i] = float64(value.Index(i).Uint())
			case 3:
				r[i] = value.Index(i).Float()
			case 4:
				r[i], _ = strconv.ParseFloat(value.Index(i).String(), 64)

			case 5:
				r[i] = value.Index(i).Interface().(float64)
			}
		}
		return r
	case reflect.String:
		r := make([]string, sliceLen)
		for i := 0; i < sliceLen; i++ {
			switch code {
			case 1:
				r[i] = strconv.FormatInt(value.Index(i).Int(), 10)
			case 2:
				r[i] = strconv.FormatUint(value.Index(i).Uint(), 10)
			case 3:
				r[i] = strconv.FormatFloat(value.Index(i).Float(), 'f', -1, 64)
			case 4:
				r[i] = value.Index(i).String()
			case 5:
				r[i] = value.Index(i).Interface().(string)
			default:
				panic("no interface default case")
			}
		}
		return r
	default:
		panic("unhandled default case")
	}
	return value.Interface()
}
