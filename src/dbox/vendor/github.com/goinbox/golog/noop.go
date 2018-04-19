package golog

type NoopLogger struct {
}

func (n *NoopLogger) Debug(msg []byte) {
}

func (n *NoopLogger) Info(msg []byte) {
}

func (n *NoopLogger) Notice(msg []byte) {
}

func (n *NoopLogger) Warning(msg []byte) {
}

func (n *NoopLogger) Error(msg []byte) {
}

func (n *NoopLogger) Critical(msg []byte) {
}

func (n *NoopLogger) Alert(msg []byte) {
}

func (n *NoopLogger) Emergency(msg []byte) {
}

func (n *NoopLogger) Log(level int, msg []byte) error {
	return nil
}

func (n *NoopLogger) Flush() error {
	return nil
}

func (n *NoopLogger) Free() {
}

type NoopFormater struct {
}

func (n *NoopFormater) Format(level int, msg []byte) []byte {
	return msg
}

type NoopWriter struct {
}

func (n *NoopWriter) Write(msg []byte) (int, error) {
	return 0, nil
}

func (n *NoopWriter) Flush() error {
	return nil
}

func (n *NoopWriter) Free() {
}
