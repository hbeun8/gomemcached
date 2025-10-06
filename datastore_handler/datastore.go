package datastorehandler

//import commandhandler "gomemc/command_handler"

//var c commandhandler

import ( 
	"time"
	"bytes"
	"strconv"
	"log"
)



type Datastore struct {
	Datastore map[string][][]byte
}

func (d *Datastore) Set(key []byte, val []byte, flag []byte, expiry []byte) [][]byte {
	if bytes.Equal(key, []byte("")) {
		return [][]byte{[]byte("ERROR\r\n")}
	} else {
		var resultblock [][]byte
		t := time.Now()
		resultblock = append(resultblock, []byte("VALUE"))
		resultblock = append(resultblock, flag)
		i, e := strconv.Atoi(string(expiry))
		if e!=nil {
			log.Println("Expiry could not be used")
			return [][]byte{[]byte("ERROR\r\n")}
		} else {
			if i == 0 {
				resultblock = append(resultblock, expiry)
			} else {
				expiryvalue := int(t.UnixNano()) + i
				resultblock = append(resultblock, []byte(strconv.Itoa(expiryvalue)))
			}
		}
		resultblock = append(resultblock, []byte(strconv.Itoa(len(val))+"\r\n"))
		resultblock = append(resultblock, val)
		resultblock = append(resultblock, []byte("END\r\n"))
		d.Datastore[string(key)] = resultblock
		return [][]byte{[]byte("STORED\r\n")}
	}
	
}


func (d *Datastore) Get(k []byte) [][]byte {
	block := d.Datastore[string(k)]
	if len(block) !=6 {
		log.Print(string(k), block)
		log.Fatal("Block format incompatible")
		return block
	}
	if len(block) == 0 {
		return [][]byte{[]byte("END\r\n")}
	}
	//block[1] = value, [1]=flag, [2]=expiry
	if !bytes.Equal(block[2], []byte("0")) {
		return block
	}
	expiry, e :=strconv.Atoi(string(block[2]))
	if e!=nil {
		log.Println("Expiry could not be used")
		return [][]byte{[]byte("ERROR\r\n")}
	}
	if int(time.Now().UnixNano()) < expiry || expiry!=0  {
		block = append(block, []byte("END\r\n"))
		return block
	} else {
		return [][]byte{[]byte("EXPIRED\r\n")}
	}
}


func (d *Datastore) Delete(k []byte) [][]byte {
	
	d.Datastore[string(k)] = [][]byte{[]byte("")}
	return [][]byte{[]byte("END\r\n")}

}

