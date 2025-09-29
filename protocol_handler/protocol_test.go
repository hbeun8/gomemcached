package protocol

import (
	"fmt"
	"testing"
)

func TestSetProtocol(t *testing.T) {

	tests := []struct{
		command string 
		key string
		flags string
		expiry string
		bytecount string
		noreply string
		datablock []string
		want []string
	} {
		{"\r\n", "", "", "", "", "", []string{""}, []string{"Missing Command\r\n"}},
		{"set", "test", "0", "0", "4", "noreply\r\n", []string{"Datablock\r\n"}, []string{"set", "test", "0", "0", "4", "noreply", "Datablock"}},
		{"set", "test", "0", "100", "4\r\n", "", []string{"Datablock\r\n"}, []string{"set", "test", "0", "100", "4", "", "Datablock"}},	
		{"set", "test", "0", "0", "4\r\n", "", []string{"Datablock\r\n"}, []string{"set", "test", "0", "0", "4", "", "Datablock"}},
		{"set", "test", "", "0", "4", "\r\n", []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"set", "test", "", "", "4", "\r\n", []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"set", "test", "\r\n", "", "", "" , []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"set", "test", "0", "0", "4", "noreply\r\n", []string{"\r\n"}, []string{"Missing datablock\r\n"}},		
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v,%v,%v, %v, %v, %v", tt.command, tt.key, tt.flags, tt.expiry, tt.bytecount, tt.noreply, tt.datablock)
        t.Run(testname, func(t *testing.T) {
				c := Commands{
						command: tt.command,
						key: tt.key,
						flags: tt.flags, 
						expiry: tt.expiry,
						bytecount: tt.bytecount,
						noreply: tt.noreply,
					}
				p := Parser{CommandLine: c, Datablock: tt.datablock}

				ans := p.Protocol_Handler()
            	for i := range ans {
					if ans[i] != tt.want[i] {
                		t.Errorf("got %v, want %v", ans, tt.want)
					}
				}
        })
    }
}



func TestGetProtocol(t *testing.T) {

	tests := []struct{
		command string 
		key string
		want []string
	} {
		{"get", "validkey\r\n", []string{"get", "validkey"}},
		{"get", "validkey\r\n", []string{"get", "validkey"}},
		{"get", "\r\n", []string{"Missing Key\r\n"}},
		{"get", "invalidkey\r\n", []string{"get", "invalidkey"}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v", tt.command, tt.key)
        t.Run(testname, func(t *testing.T) {
			c := Commands{
				command: tt.command,
				key: tt.key,
			}
			p := Parser{CommandLine: c}
			ans := p.Protocol_Handler()
            for i := range ans {
				if ans[i] != tt.want[i] {
                	t.Errorf("got %v, want %v", ans, tt.want)
				}
			}
        })
    }
}


func TestAddProtocol(t *testing.T) {

	tests := []struct{
		command string 
		key string
		flags string
		expiry string
		bytecount string
		noreply string
		datablock []string
		want []string
	} {
		{"\r\n", "", "", "", "", "", []string{""}, []string{"Missing Command\r\n"}},
		{"add", "test", "0", "0", "4", "noreply\r\n", []string{"Datablock\r\n"}, []string{"add", "test", "0", "0", "4", "noreply", "Datablock"}},
		{"add", "test", "0", "0", "4\r\n", "", []string{"Datablock\r\n"}, []string{"add", "test", "0", "0", "4", "", "Datablock"}},
		{"add", "test", "", "0", "4", "\r\n", []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"add", "test", "", "", "4", "\r\n", []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"add", "test", "\r\n", "", "", "" , []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"add", "test", "0", "0", "4", "noreply\r\n", []string{"\r\n"}, []string{"Missing datablock\r\n"}},		
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v,%v,%v, %v, %v, %v", tt.command, tt.key, tt.flags, tt.expiry, tt.bytecount, tt.noreply, tt.datablock)
        t.Run(testname, func(t *testing.T) {
				c := Commands{
						command: tt.command,
						key: tt.key,
						flags: tt.flags, 
						expiry: tt.expiry,
						bytecount: tt.bytecount,
						noreply: tt.noreply,
					}
				p := Parser{CommandLine: c, Datablock: tt.datablock}

				ans := p.Protocol_Handler()
            	for i := range ans {
					if ans[i] != tt.want[i] {
                		t.Errorf("got %v, want %v", ans, tt.want)
					}
				}
        })
    }
}

func TestReplaceProtocol(t *testing.T) {

	tests := []struct{
		command string 
		key string
		flags string
		expiry string
		bytecount string
		noreply string
		datablock []string
		want []string
	} {
		{"\r\n", "", "", "", "", "", []string{""}, []string{"Missing Command\r\n"}},
		{"replace", "test", "0", "0", "4", "noreply\r\n", []string{"Datablock\r\n"}, []string{"replace", "test", "0", "0", "4", "noreply", "Datablock"}},
		{"replace", "test", "0", "0", "4\r\n", "", []string{"Datablock\r\n"}, []string{"replace", "test", "0", "0", "4", "", "Datablock"}},
		{"replace", "test", "", "0", "4", "\r\n", []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"replace", "test", "", "", "4", "\r\n", []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"replace", "test", "\r\n", "", "", "" , []string{"Datablock\r\n"}, []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}},
		{"replace", "test", "0", "0", "4", "noreply\r\n", []string{"\r\n"}, []string{"Missing datablock\r\n"}},		
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v,%v,%v, %v, %v, %v", tt.command, tt.key, tt.flags, tt.expiry, tt.bytecount, tt.noreply, tt.datablock)
        t.Run(testname, func(t *testing.T) {
				c := Commands{
						command: tt.command,
						key: tt.key,
						flags: tt.flags, 
						expiry: tt.expiry,
						bytecount: tt.bytecount,
						noreply: tt.noreply,
					}
				p := Parser{CommandLine: c, Datablock: tt.datablock}

				ans := p.Protocol_Handler()
            	for i := range ans {
					if ans[i] != tt.want[i] {
                		t.Errorf("got %v, want %v", ans, tt.want)
					}
				}
        })
    }
}

func TestAppendPrependProtocol(t *testing.T) {

	tests := []struct{
		command string 
		key string
		flags string
		expiry string
		bytecount string
		noreply string
		datablock []string
		want []string
	} {
		{"set", "akey", "0", "0", "4\r\n", "", []string{"John\r\n"}, []string{"set", "akey", "0", "0", "4", "", "John"}},
		{"append", "akey", "0", "0", "4\r\n", "", []string{"More\r\n"}, []string{"append", "akey", "0", "0", "4", "", "More"}},
		{"prepend", "akey", "0", "0", "4\r\n", "", []string{"Send\r\n"}, []string{"prepend", "akey", "0", "0", "4", "", "Send"}},
		{"append", "invalidkey", "0", "0", "4\r\n", "", []string{"test\r\n"}, []string{"append", "invalidkey", "0", "0", "4", "", "test"}},
		{"prepend", "invalidkey", "0", "0", "4\r\n", "", []string{"test\r\n"}, []string{"prepend", "invalidkey", "0", "0", "4", "", "test"}},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v,%v,%v,%v, %v, %v, %v", tt.command, tt.key, tt.flags, tt.expiry, tt.bytecount, tt.noreply, tt.datablock)
        t.Run(testname, func(t *testing.T) {
				c := Commands{
						command: tt.command,
						key: tt.key,
						flags: tt.flags, 
						expiry: tt.expiry,
						bytecount: tt.bytecount,
						noreply: tt.noreply,
					}
				p := Parser{CommandLine: c, Datablock: tt.datablock}

				ans := p.Protocol_Handler()
            	for i := range ans {
					if ans[i] != tt.want[i] {
                		t.Errorf("got %v, want %v", ans, tt.want)
					}
				}
        })
    }
}
