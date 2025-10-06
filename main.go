package main

import connectionhandler "gomemc/connection_handler"

func main() {

	connectionhandler.Connect()
	/*
	m := make(map[string][][]byte)
	d := datastorehandler.Datastore{Datastore: m}
	r := d.Set("name", "Go", 0, 100000)
	fmt.Println(r)
	s := d.Get("name", 100000)
	fmt.Println(s)
	*/
}
