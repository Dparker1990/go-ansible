package util

func TrimNewline(msg string) string {
	return msg[0 : len(msg)-1]
}
