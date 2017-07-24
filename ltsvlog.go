// Copyright 2017 Shunsuke Michii. All rights reserved.

package ltsvlog

import (
	"bytes"
	"encoding"
	"fmt"
	"io"
	"os"
	"strconv"

	ltsv "github.com/Songmu/go-ltsv"
)

type Field struct {
	Key   string
	Value interface{}
}

func F(key string, value interface{}) Field {
	return Field{key, value}
}

const (
	format1 = "%s:%s"
	format2 = "\t%s:%s"
	delim   = "\n"
)

type Logger struct {
	w io.Writer
}

func (l *Logger) Logf(fields ...Field) {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, format1, fields[0].Key, l.format(fields[0].Value))
	for _, f := range fields[1:] {
		fmt.Fprintf(buf, format2, f.Key, l.format(f.Value))
	}
	fmt.Fprintf(buf, delim)
	buf.WriteTo(l.w)
}

func (l *Logger) Log(v interface{}) {
	buf := bytes.NewBuffer(nil)
	ltsv.MarshalTo(buf, v)
	fmt.Fprintf(buf, delim)
	buf.WriteTo(l.w)
}

func (l *Logger) format(v interface{}) string {
	var s string
	switch v := v.(type) {
	case string:
		s = v
	case encoding.TextMarshaler:
		b, err := v.MarshalText()
		if err != nil {
			// TODO: handling error
			return "(failed to marshal)"
		}
		s = string(b)
	default:
		s = fmt.Sprint(v)
	}
	if needQuote(s) {
		s = strconv.Quote(s)
	}
	return s
}

func needQuote(s string) bool {
	for _, c := range s {
		if !(0x21 <= c && c <= 0x7f && c != '"' && c != '\\') {
			return true
		}
	}
	return false
}

func (l *Logger) SetOutput(w io.Writer) {
	l.w = w
}

var DefaultLogger = &Logger{
	w: os.Stdout,
}

func Log(v interface{}) {
	DefaultLogger.Log(v)
}

func Logf(fields ...Field) {
	DefaultLogger.Logf(fields...)
}

func SetOutput(w io.Writer) {
	DefaultLogger.SetOutput(w)
}
