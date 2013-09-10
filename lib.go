package main

func trimNewline(msg string) (trimmedMsg string) {
	trimmedMsg = msg[0 : len(msg)-1]
	return
}
