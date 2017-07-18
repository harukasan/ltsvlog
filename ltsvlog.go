// Copyright 2017 Shunsuke Michii. All rights reserved.

package ltsvlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

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
	buf        *bytes.Buffer
	w          io.Writer
	TimeFormat string
}

func (l *Logger) Logf(fields ...Field) {
	defer l.buf.Reset()
	fmt.Fprintf(l.buf, format1, fields[0].Key, l.format(fields[0].Value))
	for _, f := range fields[1:] {
		fmt.Fprintf(l.buf, format2, f.Key, l.format(f.Value))
	}
	fmt.Fprintf(l.buf, delim)
	l.buf.WriteTo(l.w)
}

func (l *Logger) Log(v interface{}) {
	defer l.buf.Reset()
	ltsv.MarshalTo(l.buf, v)
	fmt.Fprintf(l.buf, delim)
	l.buf.WriteTo(l.w)
}

func (l *Logger) format(v interface{}) string {
	var s string
	switch v := v.(type) {
	case time.Time:
		s = v.Format(l.TimeFormat)
	case string:
		s = v
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
	buf:        bytes.NewBuffer(nil),
	w:          os.Stdout,
	TimeFormat: time.RFC3339,
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
