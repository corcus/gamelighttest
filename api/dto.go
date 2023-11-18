package main

type IncomingRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func (incoming IncomingRequest) Validate() bool {
	return incoming.Sender != "" && incoming.Receiver != "" && incoming.Message != ""
}
