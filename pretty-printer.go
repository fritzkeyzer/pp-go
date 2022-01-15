package pp_go

// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN

// https://en.wikipedia.org/wiki/ANSI_escape_code

import (
	"fmt"
	"github.com/fatih/color"
	"reflect"
	"strings"
)

var (
	a = color.New(color.FgGreen, color.Bold)
	b = color.New(color.FgRed, color.Bold)
)

func Dump(s interface{}) {

	title := reflect.ValueOf(s).Type().String()

	var fieldNames []string
	var fieldTypes []string
	var fieldVals []string

	sP := reflect.ValueOf(s).Elem()
	for i := 0; i < sP.NumField(); i++ {
		fieldName := sP.Type().Field(i).Name
		fieldType := fmt.Sprintf("%v", sP.Type().Field(i).Type)
		fieldVal := fmt.Sprintf("%v", sP.Field(i))

		fieldNames = append(fieldNames, fieldName)
		fieldTypes = append(fieldTypes, fieldType)
		fieldVals = append(fieldVals, fieldVal)
	}

	lnName := 0
	lnVal := 0
	for i, _ := range fieldNames {
		if len(fieldNames[i]) > lnName {
			lnName = len(fieldNames[i])
		}
		if len(fieldVals[i]) > lnVal {
			lnVal = len(fieldVals[i])
		}
	}

	a.Printf("%s\n", "- - - - - - - - - - - - - - - - - - ")
	a.Printf("%s\n", title)

	for i, _ := range fieldNames {
		name := fmt.Sprintf("%s", fieldNames[i])
		val := fmt.Sprintf("%-*s", lnVal, fieldVals[i])
		typ := fmt.Sprintf("%s", fieldTypes[i])

		padName := lnName - len(fieldNames[i])
		padVal := lnVal - len(fieldVals[i])

		a.Printf("\t%s: ", padLeft(name, " ", padName))
		color.Unset()
		b.Printf("%s ", padRight(val, "", padVal))
		color.Unset()
		a.Printf("\t%s\n", typ)
	}
	a.Printf("%s\n", "- - - - - - - - - - - - - - - - - - ")
}

func padLeft(s, p string, count int) string {
	ret := make([]byte, len(p)*count+len(s))

	b := ret[:len(p)*count]
	bp := copy(b, p)
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
	copy(ret[len(b):], s)
	return string(ret)
}

func padRight(s, p string, count int) string {

	return s + strings.Repeat(p, count)
}
