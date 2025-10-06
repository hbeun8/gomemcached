package commandhandler

import (
	"bytes"
	datastorehandler "gomemc/datastore_handler"
	protocolhandler "gomemc/protocol_handler"
)

func Sifter(buf []byte) ([]byte, []byte){
	//incomplete buffers are not addressed here but rather
	//in the protocol.
	sep := bytes.Index(buf, []byte(" "))
	if sep == -1 {
		return []byte(""), buf
	}
	switch string(bytes.ToLower(buf[:sep])) {
	case "get", "set", "append", "prepend", "add", "replace", "delete":
		return buf, []byte("")
	default:
		return []byte(""), buf
	}
}

func Command(buf []byte, dat []byte, d *datastorehandler.Datastore) [][]byte{

	sep := bytes.Index(buf, []byte(" "))
	if sep == -1 {
		return [][]byte{[]byte("Error\r\n")}
	} else {
		switch string(bytes.ToLower(buf[:sep])) {
		case "get":
			return protocolhandler.GetParser(buf, d)
		case "set": 
			return protocolhandler.SetParser(buf, dat, d)
		case "append":
			return protocolhandler.AppendParser(buf, dat, d)
		case "prepend":
			return protocolhandler.PrependParser(buf, dat, d)
		case "add":
			return protocolhandler.AddParser(buf, dat, d)
		case "replace":
			return protocolhandler.ReplaceParser(buf, dat, d)
		case "delete":
			return protocolhandler.DeleteParser(buf, d)
		default:
			return [][]byte{[]byte("Unknown Error\r\n")}
		}
	}
}