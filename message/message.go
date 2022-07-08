package message

type Message struct {
	Msg string `bson:"msg" json:"msg"`
}

func NewMessage(msg string) Message {
	return Message{
		Msg: msg,
	}
}

func ListMessages(msgs []string) []Message {
	messages := make([]Message, len(msgs))

	for k, msg := range msgs {
		messages[k] = NewMessage(msg)
	}

	return messages
}

type MessagesRepository interface {
	InsertMessage(message Message) error
	InsertManyMessages(messages []Message) error
}
