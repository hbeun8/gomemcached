package datastore

import (
	"fmt"
	"testing"
//	"time"
)

func TestSet(t *testing.T) {

	tests := []struct{
		key string 
		value string
		flags uint16
		exp int
		want string
	} {
		{"k", "v", 0, 0, "Stored\r\n"},
		{"", "", 0, 0, "Error\r\n"},
		{"k", "v", 10, 0, "Stored\r\n"},
		{"k", "v", 10, 10, "Stored\r\n"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v,%v,%v", tt.key, tt.value, tt.flags, tt.exp)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string]Results)
				d := Datastore{Datastore: m}
            	ans := d.Set(tt.key, tt.value, tt.flags, tt.exp)
            	if ans != tt.want {
                	t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}


func TestGet(t *testing.T) {

	tests := []struct{
		key string
		value string
		flags uint16
		exp int
		want Results
	} {
		{"k", "v", 0, 0, Results{value : "v", flags: 0, byte_count: 0, data_block: "k", expiry:0}},
		{"", "", 0, 0, Results{ value: "End\r\n"}},

		//make expiry into int64
		{"m", "p", 0, 10000, Results{value: "p", flags:0, data_block:"m", byte_count:0, expiry: 10000}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v,%v,%v", tt.key, tt.value, tt.flags, tt.exp)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string]Results)
				d := Datastore{Datastore: m}
				d.Set(tt.key, tt.value, tt.flags, tt.exp)
            	ans := d.Get(tt.key, tt.exp)
            	if ans != tt.want {
                	t.Errorf("got %v, want %v", ans, tt.want)
            }
        })
    }
}

