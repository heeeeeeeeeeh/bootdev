package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	s := mToSend.sender
	r := mToSend.recipient

	if s.name == "" || s.number == 0 {
		return false
	}

	if r.name == "" || r.number == 0 {
		return false
	}
	return true
}
