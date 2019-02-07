package message

import (
	"math/rand"
	"testing"
	"time"
)

var (
	testSingleMSG = []string{
		"test message",
	}
	testManyMSG = []string{
		"test message",
		"test message 2",
		"test message 3",
		"test message 100500",
	}
	testNoMSG = []string{}
)

func TestFileMessageGetGreeting(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	type fields struct {
		r   *rand.Rand
		msg StorageModel
	}
	tests := []struct {
		name   string
		fields fields
		empty  bool
	}{
		{
			name: "single message",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Greeting: testSingleMSG,
				},
			},
			empty: false,
		},
		{
			name: "several messages",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Greeting: testManyMSG,
				},
			},
			empty: false,
		},
		{
			name: "no messages",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Greeting: testNoMSG,
				},
			},
			empty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FileMessage{
				r:   tt.fields.r,
				msg: tt.fields.msg,
			}

			got := m.GetGreeting()
			if (!tt.empty && !testStringInSlice(got, tt.fields.msg.Greeting)) || (tt.empty && got != "") {
				t.Errorf("GetGreeting() returns wrong message %s", got)
			}
		})
	}
}

func TestFileMessageGetCurse(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	type fields struct {
		r   *rand.Rand
		msg StorageModel
	}
	tests := []struct {
		name   string
		fields fields
		empty  bool
	}{
		{
			name: "single message",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Curse: testSingleMSG,
				},
			},
			empty: false,
		},
		{
			name: "several messages",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Curse: testManyMSG,
				},
			},
			empty: false,
		},
		{
			name: "no messages",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Curse: testNoMSG,
				},
			},
			empty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FileMessage{
				r:   tt.fields.r,
				msg: tt.fields.msg,
			}

			got := m.GetCurse()
			if (!tt.empty && !testStringInSlice(got, tt.fields.msg.Curse)) || (tt.empty && got != "") {
				t.Errorf("GetCurse() returns wrong message %s", got)
			}
		})
	}
}

func TestFileMessageGetRoll(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	type fields struct {
		r   *rand.Rand
		msg StorageModel
	}
	tests := []struct {
		name   string
		fields fields
		empty  bool
	}{
		{
			name: "single message",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Roll: testSingleMSG,
				},
			},
			empty: false,
		},
		{
			name: "several messages",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Roll: testManyMSG,
				},
			},
			empty: false,
		},
		{
			name: "no messages",
			fields: fields{
				r: rnd,
				msg: StorageModel{
					Roll: testNoMSG,
				},
			},
			empty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &FileMessage{
				r:   tt.fields.r,
				msg: tt.fields.msg,
			}

			got := m.GetRoll()
			if (!tt.empty && !testStringInSlice(got, tt.fields.msg.Roll)) || (tt.empty && got != "") {
				t.Errorf("GetRoll() returns wrong message '%s'", got)
			}
		})
	}
}

func testStringInSlice(s string, ss []string) bool {
	for _, rs := range ss {
		if s == rs {
			return true
		}
	}
	return false
}
