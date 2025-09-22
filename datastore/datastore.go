package datastore

//import commandhandler "gomemc/command_handler"

//var c commandhandler

type Results struct {
	value string
	flags int
	data_block string
	byte_count int
}

type Datastore struct {
	Datastore map[string] Results
}

func (d *Datastore) Set(k string, v string) string {
	if k ==  "" {
		return "Error\r\n"
	} else {
		r:=Results{value : v, data_block: k}
		d.Datastore[k] = r
		return "Stored\r\n"
	}
}


func (d *Datastore) Get(k string) Results {
	//d.Datastore = make(map[string]Results)
	r := d.Datastore[k]
	if r.value != "" {
		return  r
	} else {
		return Results{ value: "End\r\n"}
	}
}


