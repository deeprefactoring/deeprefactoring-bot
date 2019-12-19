package message

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"time"
)

// ShuffleStringSlice is a string slice with counter and shuffle methods
type ShuffleStringSlice struct {
	s []string
	i int
}

// UnmarshalYAML parses string array into a slice and initializes a counter
func (s *ShuffleStringSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str []string
	err := unmarshal(&str)
	s.s = str
	return err
}

// ShuffleWith does a shuffle based on a shuffle function
func (s *ShuffleStringSlice) ShuffleWith(sf func([]string)) {
	sf(s.s)
}

// GetNext returns a next string from a shuffled slice
// it will begin from the first element if the counter will reach the end of the slice
func (s *ShuffleStringSlice) GetNext() string {
	if len(s.s) == 0 {
		return ""
	}

	s.i++
	if s.i == len(s.s) {
		s.i = 0
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
	r   *rand.Rand
	msg StorageModel
}

// NewFileMessage creates new message repository access entity
func NewFileMessage(path string) (*FileMessage, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	msg := StorageModel{}

	err = yaml.Unmarshal(yamlFile, &msg)
	if err != nil {
		return nil, err
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	shuffleFunc := func(a []string) {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}

	// shuffle message slices
	msg.Greeting.ShuffleWith(shuffleFunc)
	msg.Curse.ShuffleWith(shuffleFunc)
	msg.Roll.ShuffleWith(shuffleFunc)

	return &FileMessage{
		r:   r,
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
