package cast

// Suite of casting functions to speed up solutions
// This is NOT idiomatic Go... but AOC isn't about that...

import (
	"fmt"
	"strconv"
)

// ToInt will case a given arg into an int type.
// Supported types are:
//   - string
func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// ToInt will case a given arg into an int type.
// Supported types are:
//   - string
func ToInt64(arg interface{}) int64 {
	var val int64
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.ParseInt(arg.(string), 10, 64)
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

func ToUInt64(arg interface{}) uint64 {
	var val uint64
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.ParseUint(arg.(string), 10, 64)
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// ToString will case a given arg into an int type.
// Supported types are:
//   - int
//   - byte
//   - rune
func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case int64:
		str = strconv.FormatInt(arg.(int64), 10)
	case uint64:
		str = strconv.FormatUint(arg.(uint64), 10)
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

const (
	ASCIICodeCapA   = int('A') // 65
	ASCIICodeCapZ   = int('Z') // 65
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 97
)

// ToASCIICode returns the ascii code of a given input
func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}

	return asciiVal
}

// ASCIIIntToChar returns a one character string of the given int
func ASCIIIntToChar(code int) string {
	return string(rune(code))
}

func StringListToIntList(s []string) []int {

	var retVal = make([]int, len(s))

	for idx, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		retVal[idx] = j
	}
	return retVal
}

func StringListToInt64List(s []string) []int64 {

	var retVal = make([]int64, len(s))

	for idx, i := range s {
		j := ToInt64(i)
		retVal[idx] = j
	}
	return retVal
}

func StringListToUInt64List(s []string) []uint64 {

	var retVal = make([]uint64, len(s))

	for idx, i := range s {
		j := ToUInt64(i)
		retVal[idx] = j
	}
	return retVal
}
