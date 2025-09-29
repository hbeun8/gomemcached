package datastore

//import commandhandler "gomemc/command_handler"

//var c commandhandler

import ( "time")

type Results struct {
	value string
	flags uint16
	data_block string
	byte_count int
	expiry int
}

type Datastore struct {
	Datastore map[string] Results
}

func (d *Datastore) Set(k string, v string, f uint16, e int) string {
	if k ==  "" {
		return "Error\r\n"
	} else {
		t := time.Now()
		if e == 0 {
			r:=Results{value : v, data_block: k, flags: f, expiry: 0}	
			d.Datastore[k] = r
		} else {
			r:=Results{value : v, data_block: k, flags: f, expiry: int(t.UnixNano()) + e}
			d.Datastore[k] = r
		}
		return "Stored\r\n"
	}
}


func (d *Datastore) Get(k string, e int) Results {
	
	r := d.Datastore[k]
	if r.value != "" {
		if e == 0 {
			return r
		}
		if int(time.Now().UnixNano()) < r.expiry || r.expiry!=0  {
			r.expiry = e
			return r
		} else {
			return Results{ value: "Expired\r\n"}
		}
	} else {
		return Results{ value: "End\r\n"}
	}
}


