package shared

type ServerStatus struct {
	IsClosed bool
}

type Message struct {
	Cmd     byte
	Content []byte
}

type Cmd byte

const (
	CmdOk Cmd = 200
)
