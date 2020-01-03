package message

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"sync/atomic"
	"time"
)

// ShuffleStringSlice is a string slice with counter and shuffle methods
type ShuffleStringSlice struct {
	s []string
	i uint64
}

// UnmarshalYAML parses string array into a slice and initializes a counter
func (s *ShuffleStringSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str []string
	err := unmarshal(&str)
	shuffleStringSlice(str)
	s.s = str
	return err
}

// GetNext returns a next string from a shuffled slice
// it will begin from the first element if the counter will reach the end of the slice
func (s *ShuffleStringSlice) GetNext() string {
	if len(s.s) == 0 {
		return ""
	}

	atomic.AddUint64(&s.i, 1)
	if s.i == uint64(len(s.s)) {
		s.i = 0
		shuffleStringSlice(s.s)
	}
	return s.s[s.i]
}

// StorageModel format of message file
type StorageModel struct {
	Greeting ShuffleStringSlice `yaml:"greeting"`
	Curse    ShuffleStringSlice `yaml:"curse"`
	Roll     ShuffleStringSlice `yaml:"roll"`
}

// FileMessage provides messages from internal storage
type FileMessage struct {
	msg StorageModel
}

// NewFileMessage creates new message repository access entity
func NewFileMessage(path string) (*FileMessage, error) {
	// set random generator seed
	rand.Seed(time.Now().UTC().UnixNano())

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	msg := StorageModel{}

	err = yaml.Unmarshal(yamlFile, &msg)
	if err != nil {
		return nil, err
	}

	return &FileMessage{
		msg: msg,
	}, nil
}

// GetGreeting returns a greeting message from file
func (m *FileMessage) GetGreeting() string {
	return m.msg.Greeting.GetNext()
}

// GetCurse returns a farewell message from file
func (m *FileMessage) GetCurse() string {
	return m.msg.Curse.GetNext()
}

// GetRoll returns a topic message from file
func (m *FileMessage) GetRoll() string {
	return m.msg.Roll.GetNext()
}

// shuffleStringSlice does a slice shuffle
func shuffleStringSlice(a []string) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}
