package message

import (
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	testSingleMSG = ShuffleStringSlice{
		s: []string{
			"test message",
		},
	}

	testManyMSG = ShuffleStringSlice{
		s: []string{
			"test message",
			"test message 2",
			"test message 3",
			"test message 100500",
		},
	}
	testNoMSG ShuffleStringSlice
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
			if (!tt.empty && !testStringInSlice(got, tt.fields.msg.Greeting.s)) || (tt.empty && got != "") {
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
			if (!tt.empty && !testStringInSlice(got, tt.fields.msg.Curse.s)) || (tt.empty && got != "") {
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
			if (!tt.empty && !testStringInSlice(got, tt.fields.msg.Roll.s)) || (tt.empty && got != "") {
				t.Errorf("GetRoll() returns wrong message '%s'", got)
			}
		})
	}
}

func TestNewFileMessage(t *testing.T) {
	msgText := `---
greeting:
  - "greeting1"
  - "greeting2"
  - "greeting3"

curse:
  - "curse1"
  - "curse2"
  - "curse3"

roll:
  - "roll1"
  - "roll2"
  - "roll3"
`
	file, err := ioutil.TempFile(os.TempDir(), "*.yml")
	if err != nil {
		t.Fatal(err)
	}
	file.Write([]byte(msgText))
	defer os.Remove(file.Name())

	fm, err := NewFileMessage(file.Name())
	if err != nil {
		t.Errorf("unexpected error %v", err)
	} else {
		if !testStringInSlice("greeting2", fm.msg.Greeting.s) {
			t.Errorf("wrong file parsing: rolls %v must contain %s", fm.msg.Greeting.s, "greeting2")
		}
		if !testStringInSlice("curse2", fm.msg.Curse.s) {
			t.Errorf("wrong file parsing: rolls %v must contain %s", fm.msg.Curse.s, "curse2")
		}
		if !testStringInSlice("roll2", fm.msg.Roll.s) {
			t.Errorf("wrong file parsing: rolls %v must contain %s", fm.msg.Roll.s, "roll2")
		}
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
