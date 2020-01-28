package main

type (
	Msg map[string]interface{}
)

func (m Msg) GetProtocol() string {
	return m["protocol"].(string)
}

func (m Msg) GetCommand() string {
	return m["command"].(string)
}
