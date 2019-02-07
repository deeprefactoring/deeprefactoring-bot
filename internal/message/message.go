package message

import (
	"io/ioutil"
	"math/rand"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// StorageModel format of message file
type StorageModel struct {
	Greeting []string `yaml:"greeting"`
	Curse    []string `yaml:"curse"`
	Roll     []string `yaml:"roll"`
}

// FileMessage provides messages from internal storage
type FileMessage struct {
	r   *rand.Rand
	msg StorageModel
}

// NewFileMessage create new mesage repository access entity
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

	return &FileMessage{
		r:   r,
		msg: msg,
	}, nil
}

// randomMessage returns a random message from a slice
func (m *FileMessage) randomMessage(messages []string) string {
	// avoid empty slice crash
	if len(messages) == 0 {
		return ""
	}
	return messages[m.r.Intn(len(messages))]
}

// GetGreeting returns a greeting message from file
func (m *FileMessage) GetGreeting() string {
	return m.randomMessage(m.msg.Greeting)
}

// GetCurse returns a farewell message from file
func (m *FileMessage) GetCurse() string {
	return m.randomMessage(m.msg.Curse)
}

// GetRoll returns a topic message from file
func (m *FileMessage) GetRoll() string {
	return m.randomMessage(m.msg.Roll)
}
