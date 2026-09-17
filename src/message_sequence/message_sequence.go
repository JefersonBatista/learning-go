package message_sequence

type Message interface {
	Bytes() []byte
	String() string
}

type Sequence interface {
	Next() Message
}

type condition func(m Message) bool
type transformation func(m Message) Message

type filter struct {
	Sequence
	condition condition
}

func (filter filter) Next() (next Message) {
	next = filter.Sequence.Next()
	for !filter.condition(next) {
		next = filter.Sequence.Next()
	}
	return
}

func Filter(sequence Sequence, condition condition) Sequence {
	return filter{sequence, condition}
}

type transform struct {
	Sequence
	transformation transformation
}

func (transform transform) Next() Message {
	next := transform.Sequence.Next()
	return transform.transformation(next)
}

func Transform(sequence Sequence, transformation transformation) Sequence {
	return transform{sequence, transformation}
}
