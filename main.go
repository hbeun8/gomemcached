package main

import (
	"gomemc/datastore"
	"fmt"
)

func main() {

	m := make(map[string]datastore.Results)
	d := datastore.Datastore{Datastore: m}
	r := d.Set("name", "Go")
	fmt.Println(r)
	s := d.Get("name")
	fmt.Println(s)

}
