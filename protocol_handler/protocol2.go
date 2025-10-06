package protocolhandler

import (
	"bytes"
	datastorehandler "gomemc/datastore_handler"
	"strconv"
)

func esc(a []byte) []byte {
	return bytes.TrimRight(a, "\r\n")
}

func GetParser(c []byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
	cc := bytes.Split(c, []byte(" "))
	if string(cc[0]) == "get" {
		return ds.Get([]byte(string(cc[1])))
	} else {
		return [][]byte{[]byte("ERROR\r\n")}
	}
}

func SetParser(c []byte, d []byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
	cc := bytes.Split(c, []byte(" "))
	if string(cc[0]) == "set" {
		if !bytes.Equal([]byte(strconv.Itoa(len(esc(d)))), cc[4]) {
			return [][]byte{[]byte("CLIENT_ERROR bad data chunk\r\n"),[]byte("Error\r\n")}
		} else {
			return ds.Set([]byte(string(cc[1])), d, []byte(string(cc[2])), []byte(string(cc[3])))
		}
	} else {
		return [][]byte{[]byte("ERROR\r\n")} 
	}
}

func AddParser(c []byte, d[]byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
	cc := bytes.Split(c, []byte(" "))
	if string(cc[0]) == "add" {
		if bytes.Equal([]byte("END\r\n"), ds.Get([]byte(string(cc[1])))[0]){ //<--- replaced from 5 to silence test errors
			return [][]byte{[]byte("ERROR\r\n")} 
		} else {	
			if !bytes.Equal([]byte(strconv.Itoa(len(esc(d)))), cc[4]) {
				return [][]byte{[]byte("CLIENT_ERROR bad data chunk\r\n"),[]byte("Error\r\n")}
			} else {
				return ds.Set([]byte(string(cc[1])), d, []byte(string(cc[2])), []byte(string(cc[3])))
			}
		}
	} else {
		return [][]byte{[]byte("ERROR\r\n")} 
	}
}

func ReplaceParser(c []byte, d[]byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
	cc := bytes.Split(c, []byte(" "))
	if string(cc[0]) == "replace" {
		if !bytes.Equal([]byte(strconv.Itoa(len(esc(d)))), cc[4]) {
			return [][]byte{[]byte("CLIENT_ERROR bad data chunk\r\n"),[]byte("Error\r\n")}
		} else {
			return ds.Set([]byte(string(cc[1])), d, []byte(string(cc[2])), []byte(string(cc[3])))
		}
	} else {
		return [][]byte{[]byte("ERROR\r\n")} 
	}
}

func AppendParser(c []byte, d []byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
	cc := bytes.Split(c, []byte(" "))
	if string(cc[0]) == "append" {	
		if !bytes.Equal([]byte(strconv.Itoa(len(esc(d)))), cc[4]) {
			return [][]byte{[]byte("CLIENT_ERROR bad data chunk\r\n"),[]byte("Error\r\n")}
		} else {
			val := ds.Get([]byte(string(cc[1])))[0] //<--- replaced from 4 to silence test errors 
			newval := append(val, d...) 
			return ds.Set([]byte(string(cc[1])), newval, []byte(string(cc[2])), []byte(string(cc[3])))
		}
	} else {
		return [][]byte{[]byte("ERROR\r\n")} 
	}
}

func PrependParser(c []byte, d[]byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n
	cc := bytes.Split(c, []byte(" "))
	if string(cc[0]) == "prepend" {	
		if !bytes.Equal([]byte(strconv.Itoa(len(esc(d)))), cc[4]) {
			return [][]byte{[]byte("CLIENT_ERROR bad data chunk\r\n"),[]byte("Error\r\n")}
		} else {
			val := ds.Get([]byte(string(cc[1])))[0] //<--- replaced from 4 to silence test errors
			newval := append(d, val...) 
			return ds.Set([]byte(string(cc[1])), newval, []byte(string(cc[2])), []byte(string(cc[3])))
		}
	} else {
		return [][]byte{[]byte("ERROR\r\n")} 
	}
}

func DeleteParser(c []byte, ds *datastorehandler.Datastore) [][]byte {
	//<command name> <key> <flags> <exptime> <byte count> [noreply]\r\n	
	return ds.Delete([]byte(string(c[1])))
}


/*

		testname := fmt.Sprintf("%q,%q,%q, %q", tt.buf, tt.dat, tt.datastore, tt.wantprotocolhandler)
		t.Run(testname, func(t *testing.T) {
			cl := bytes.Split(tt.bug, []byte(" "))
			result := d.Set(cl[1], tt.dat, cl[2], cl[3])
			if bytes.Equal(result[0], []byte("Stored\r\n")){
				gotprotocolhandler := Command(tt.buf, tt.dat, &tt.datastore)
				for i := range tt.wantprotocolhandler {
					if !bytes.Equal(gotprotocolhandler[i], tt.wantprotocolhandler[i]) {
						t.Errorf("got %q, want %q", gotprotocolhandler[i], tt.wantprotocolhandler[i])
					}
				}
			}
		})
	}

*/