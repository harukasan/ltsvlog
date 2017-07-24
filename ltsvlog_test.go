// Copyright 2017 Shunsuke Michii. All rights reserved.

package ltsvlog_test

import (
	"bytes"
	"testing"
	"time"

	. "github.com/harukasan/ltsvlog"
)

func TestWrite(t *testing.T) {
	ti := time.Unix(0, 0)

	tests := []struct {
		f    func()
		want []byte
	}{
		{
			f: func() {
				Logf(F("time", ti), F("int", 123), F("float", 3.2), F("quoted", "\""))
			},
			want: []byte("time:1970-01-01T09:00:00+09:00\tint:123\tfloat:3.2\tquoted:\"\\\"\"\n"),
		},
		{
			f: func() {
				Logf(F("hoge", "fuga"), F("piyo", "piyo"))
			},
			want: []byte("hoge:fuga\tpiyo:piyo\n"),
		},
		{
			f: func() {
				s := struct {
					Text   string
					Answer int
					Yo     string `ltsv:"-"`
				}{
					Text:   "the Answer",
					Answer: 42,
					Yo:     "should be ignored",
				}
				Log(s)
			},
			want: []byte("text:the Answer\tanswer:42\n"),
		},
	}

	w := bytes.NewBuffer(nil)
	SetOutput(w)
	for _, test := range tests {
		w.Reset()
		test.f()
		if got := w.Bytes(); !bytes.Equal(got, test.want) {
			t.Errorf("\ngot:  %swant: %s", string(got), string(test.want))
		}
	}
}
