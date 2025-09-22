package datastore

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {

	tests := []struct{
		key string 
		value string
		want string
	} {
		{"k", "v", "Stored\r\n"},
		{"", "", "Error\r\n"},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v", tt.key, tt.value)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string]Results)
				d := Datastore{Datastore: m}
            	ans := d.Set(tt.key, tt.value)
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
		want Results
	} {
		{"k", "v", Results{value : "v", data_block: "k"}},
		{"", "", Results{ value: "End\r\n"}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt.key)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string]Results)
				d := Datastore{Datastore: m}
				d.Set(tt.key, tt.value)
            	ans := d.Get(tt.key)
            	if ans != tt.want {
                	t.Errorf("got %+v, want %+v", ans, tt.want)
            }
        })
    }
}

