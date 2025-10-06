package commandhandler

import (
	"bytes"
	"fmt"
	"testing"
	datastorehandler "gomemc/datastore_handler"
	protocolhandler "gomemc/protocol_handler"
)

func TestSifter(t *testing.T) {

	tests := []struct {
		//incomplete buffers will not be addressed here.
		buf                            []byte
		wantcommandline, wantdatablock []byte
	}{
		{[]byte("set k 0 0 flag\r\n"), []byte("set k 0 0 flag\r\n"), []byte("")},
		{[]byte("get k\r\n"), []byte("get k\r\n"), []byte("")},
		{[]byte("append k 0 0 flag\r\n"), []byte("append k 0 0 flag\r\n"), []byte("")},
		{[]byte("prepend k 0 0 flag\r\n"), []byte("prepend k 0 0 flag\r\n"), []byte("")},
		{[]byte("add k 0 0 flag\r\n"), []byte("add k 0 0 flag\r\n"), []byte("")},
		{[]byte("replace k 0 0 flag\r\n"), []byte("replace k 0 0 flag\r\n"), []byte("")},
		{[]byte("delete k 0 0 flag\r\n"), []byte("delete k 0 0 flag\r\n"), []byte("")},
		{[]byte("datablock\r\n"), []byte(""), []byte("datablock\r\n")},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q,%q,%q", tt.buf, tt.wantcommandline, tt.wantdatablock)
		t.Run(testname, func(t *testing.T) {
			gotcommandline, gotdatablock := Sifter(tt.buf)
			if !bytes.Equal(gotcommandline, tt.wantcommandline) || !bytes.Equal(gotdatablock, tt.wantdatablock) {
				t.Errorf("got %v - %v, want %v - %v", gotcommandline, tt.wantcommandline, gotdatablock, tt.wantdatablock)
			}
		})
	}
}

func TestCommandwDat(t *testing.T) {
	d := datastorehandler.Datastore{Datastore: (make(map[string] [][]byte))} 
	tests := []struct {
		//incomplete buffers will not be addressed here.
		buf, dat 			[]byte
		datastore 			datastorehandler.Datastore
		wantprotocolhandler [][]byte
	}{
		{[]byte("set k 0 0 7 flag\r\n"), []byte("namaste\r\n"),  d, 
			protocolhandler.SetParser(
				[]byte("set k 0 0 7 flag\r\n"), 
				[]byte("namaste\r\n"), 
				&d,
			),
		},
		{[]byte("append k 0 0 9 flag\r\n"), []byte("datablock\r\n"), d,
			protocolhandler.AppendParser(
				[]byte("append k 0 0 9 flag\r\n"),
				[]byte("datablock\r\n"),
				&d,
			),
		},
		{[]byte("prepend k 0 0 9 flag\r\n"), []byte("datablock\r\n"), d,
			protocolhandler.PrependParser(
				[]byte("prepend k 0 0 9 flag\r\n"),
				[]byte("datablock\r\n"),
				&d,
			),
		},
		{[]byte("add k 0 0 5 flag\r\n"), []byte("hello\r\n"), d,
				protocolhandler.AddParser(
				[]byte("add k 0 0 5 flag\r\n"),
				[]byte("hello\r\n"),
				&d,
			),
		},
		{[]byte("replace k 0 0 5 flag\r\n"), []byte("hello\r\n"), d,
				protocolhandler.ReplaceParser(
				[]byte("replace k 0 0 5 flag\r\n"),
				[]byte("hello\r\n"),
				&d,
			),
		},
		{[]byte("unknowncommand\r\n"), []byte(""), d, [][]byte{[]byte("Error\r\n")}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q,%q,%q, %q", tt.buf, tt.dat, tt.datastore, tt.wantprotocolhandler)
		t.Run(testname, func(t *testing.T) {
			gotprotocolhandler := Command(tt.buf, tt.dat, &tt.datastore)
			//fmt.Printf("%q", gotprotocolhandler)
			for i := range tt.wantprotocolhandler {
				if !bytes.Equal(gotprotocolhandler[i], tt.wantprotocolhandler[i]) {
					t.Errorf("got %q, want %q", gotprotocolhandler[i], tt.wantprotocolhandler[i])
				}

			}
		})
	}
}

func TestCommandwoDat(t *testing.T) {
	d := datastorehandler.Datastore{Datastore: (make(map[string] [][]byte))} 
	tests := []struct {
		//incomplete buffers will not be addressed here.
		buf		 			[]byte
		datastore 			datastorehandler.Datastore
		wantprotocolhandler [][]byte
	}{
	{[]byte("get k\r\n"), d, protocolhandler.GetParser([]byte("get k\r\n"),&d)},
	{[]byte("delete k\r\n"), d, protocolhandler.DeleteParser([]byte("delete k\r\n"), &d)},
	{[]byte("unknowncommand\r\n"), d, [][]byte{[]byte("ERROR\r\n")}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%q,%q,%q", tt.buf, tt.datastore, tt.wantprotocolhandler)
		t.Run(testname, func(t *testing.T) {
			cl := bytes.Split(tt.buf, []byte(" "))
			result := d.Set(cl[1], []byte("namaste"), cl[2], cl[3])
			if bytes.Equal(result[0], []byte("Stored\r\n")){
				gotprotocolhandler := Command(tt.buf, []byte("namaste"), &tt.datastore)
				for i := range tt.wantprotocolhandler {
					if !bytes.Equal(gotprotocolhandler[i], tt.wantprotocolhandler[i]) {
						t.Errorf("got %q, want %q", gotprotocolhandler[i], tt.wantprotocolhandler[i])
					}
				}
			}
		})
	}
}