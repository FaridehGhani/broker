package repository

import "github.com/FaridehGhani/broker/message"

func (m MongoDb) InsertMessage(message message.Message) error {
	return m.insert(message)
}

func (m MongoDb) InsertManyMessages(messages []message.Message) error {
	documents := make([]interface{}, len(messages))
	for k, msg := range messages {
		documents[k] = msg
	}

	return m.insertMany(documents)
}
