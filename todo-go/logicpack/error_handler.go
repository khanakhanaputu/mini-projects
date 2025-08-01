package logicpack

type notFoundErr struct {
	Message string
}

func (n *notFoundErr) Error() string {
	return n.Message
}
