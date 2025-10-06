package datastorehandler

import (
	"bytes"
	"fmt"
	"testing"
	// "time"
)

func TestSet(t *testing.T) {

	tests := []struct{
		key []byte 
		value []byte
		flag []byte
		expiry []byte
		want [][]byte
	} {
		{[]byte("k"), []byte("v"), []byte("0"), []byte("1"), [][]byte{[]byte("STORED\r\n")}},
		{[]byte(""), []byte(""), []byte(""), []byte(""),[][]byte{[]byte("ERROR\r\n")}},
		{[]byte("k"), []byte("v"), []byte("0"), []byte("Invlid Expiry"),[][]byte{[]byte("ERROR\r\n")}},
		{[]byte("k"), []byte("v"), []byte("10"), []byte("1"), [][]byte{[]byte("STORED\r\n")}},
		{[]byte("k"), []byte("v"), []byte("10"), []byte("1"), [][]byte{[]byte("STORED\r\n")}},
		{[]byte("k"), []byte("v"), []byte("10"), []byte("string"), [][]byte{[]byte("ERROR\r\n")}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q,%q,%q,%q", tt.key, tt.value, tt.flag, tt.expiry)
        t.Run(testname, func(t *testing.T) {
			M := make(map[string][][]byte)
			D := Datastore{Datastore: M}
            ans := D.Set(tt.key, tt.value, tt.flag, tt.expiry)
            for i := range tt.want {
				if !bytes.Equal(ans[i], tt.want[i]) {
                	t.Errorf("got %q, want %q", ans, tt.want)
            	}
			}
        })
    }
}


func TestGet(t *testing.T) {
	tests := []struct{
		key []byte
		want [][]byte
	} {
		{[]byte("k"), [][]byte{[]byte("value"), []byte("0"), []byte("0")}},
		{[]byte("notk"), [][]byte{[]byte("End\r\n")}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt.key)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string][][]byte)
				d := Datastore{Datastore: m}
				result := d.Set([]byte("key"), []byte("value"), []byte("0"), []byte("0"))
            	if bytes.Equal(result[0], []byte("Stored\r\n")){
					ans := d.Get(tt.key)
            		for i:=range ans{
						if !bytes.Equal(ans[i], tt.want[i]) {
                			t.Errorf("got %v, want %v", ans[i], tt.want[i])
					}
				}
            }
        })
    }
}

/*
func TestAdd(t *testing.T) {
	tests := []struct{
		key []byte 
		value []byte
		flags []byte
		exp []byte
		want [][]byte
	} {
		{[]byte("k"), []byte("v"), []byte("0"), []byte("0"), [][]byte{[]byte("Not_Stored\r\n")}},
		{[]byte("new_k"), []byte("v"), []byte("0"), []byte("0"), [][]byte{[]byte("Stored\r\n")}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%q,%q,%q,%q", tt.key, tt.value, tt.flags, tt.exp)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string][][]byte)
				d := Datastore{Datastore: m}
				result := d.Set([]byte("k"), []byte("value"), []byte("0"), []byte("0"))
            	if string(result[0]) == "Stored" {
					ans := d.Add(tt.key, tt.value, tt.flags, tt.exp)
					for i := range tt.want {
						if !bytes.Equal(ans[i], tt.want[i]) {
                			t.Errorf("got %q, want %q", ans, tt.want)
						}
					}
				}
        })
    }
}


func TestReplace(t *testing.T) {
	tests := []struct{
		key []byte 
		value []byte
		flags []byte
		exp []byte
		want [][]byte
	} {
		{[]byte("k"), []byte("v"), []byte("0"), []byte("0"), [][]byte{[]byte("Stored\r\n")}},
		{[]byte("new_k"), []byte("v"), []byte("0"), []byte("0"), [][]byte{[]byte("Stored\r\n")}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%q,%q,%q,%q", tt.key, tt.value, tt.flags, tt.exp)
        t.Run(testname, func(t *testing.T) {
				m := make(map[string][][]byte)
				d := Datastore{Datastore: m}
				result := d.Set([]byte("k"), []byte("value"), []byte("0"), []byte("0"))
            	if string(result[0]) == "Stored" {
					ans := d.Replace(tt.key, tt.value, tt.flags, tt.exp)
					for i:=range tt.want{
						if !bytes.Equal(ans[i], tt.want[i]) {
                			t.Errorf("got %q, want %q", ans[i], tt.want[i])
						}
					}
				}
        })
    }
}

*/