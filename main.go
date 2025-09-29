package main

import (
	"gomemc/datastore"
	"fmt"
)

func main() {

	m := make(map[string]datastore.Results)
	d := datastore.Datastore{Datastore: m}
	r := d.Set("name", "Go", 0, 100000)
	fmt.Println(r)
	s := d.Get("name", 100000)
	fmt.Println(s)

}
