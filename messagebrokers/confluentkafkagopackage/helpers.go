package confluentkafkagopackage

// *** для счётчика ***
type Counting struct {
	ch chan struct {
		Message string
		Count   int
	}
}

func NewCounting() *Counting {
	return &Counting{ch: make(chan struct {
		Message string
		Count   int
	})}
}

func (c *Counting) SendMessage(msg string, count int) {
	c.ch <- struct {
		Message string
		Count   int
	}{Message: msg, Count: count}
}

func (c *Counting) GetChan() <-chan struct {
	Message string
	Count   int
} {
	return c.ch
}

// *** для сообщений ***
type Message struct {
	Type, Message string
}

func (m *Message) GetType() string {
	return m.Type
}

func (m *Message) SetType(v string) {
	m.Type = v
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) SetMessage(v string) {
	m.Message = v
}

// *** для логирования ***
type Logging struct {
	Ch chan Messager
}

func NewLogging() *Logging {
	return &Logging{
		Ch: make(chan Messager),
	}
}

func (l *Logging) GetChan() <-chan Messager {
	return l.Ch
}

func (l *Logging) Send(msgType, msgData string) {
	l.Ch <- &Message{Type: msgType, Message: msgData}
}
