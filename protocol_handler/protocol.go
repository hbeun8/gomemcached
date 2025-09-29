package protocol

import "strings"

type Commands struct {
	command string
	key string
	flags string
	expiry string
	bytecount string
	noreply string
}

type Parser struct {
	CommandLine Commands
	Datablock []string
}

func fl(s string, sep string) string {
	return strings.Trim(s, sep)
}

func (p *Parser) Protocol_Handler() []string {
	sep := "\r\n"
	b := strings.ToLower(fl(p.CommandLine.command, sep))
	switch b {
	case "" : 
			return []string{"Missing Command\r\n"}
	case "set", "add", "replace", "append", "prepend":
		if p.CommandLine.flags == "" || p.CommandLine.expiry == "" || p.CommandLine.bytecount == "" {
			return []string{"Invalid Command, missing key||flag||exptime||bytecount\r\n"}
		} 
		if p.CommandLine.key == "" {
			return []string{"Missing Command\r\n"}
		} else {
			if p.CommandLine.noreply != "noreply" {
				if strings.TrimSpace(p.Datablock[0]) == "" {
					return []string{"Missing datablock\r\n"}
				} else {
				return []string{
					fl(p.CommandLine.command, sep),
					fl(p.CommandLine.key, sep),
					fl(p.CommandLine.flags, sep),
					fl(p.CommandLine.expiry, sep),
					fl(p.CommandLine.bytecount, sep),
					fl(p.CommandLine.noreply, sep),
					strings.Join(strings.Split(p.Datablock[0], sep), ""),
					}
				}
			} else {
				if strings.TrimSpace(p.Datablock[0]) == "" {
					return []string{"Missing datablock\r\n"}
				} else {
					return []string{
						fl(p.CommandLine.command, sep),
						fl(p.CommandLine.key, sep),
						fl(p.CommandLine.flags, sep),
						fl(p.CommandLine.expiry, sep),
						fl(p.CommandLine.bytecount, sep),
						strings.Join(strings.Split(p.Datablock[0], sep), ""),
					}
				}
			}
		}
	case "get":
		{
			
			if fl(p.CommandLine.key, sep) == ""{
				return []string{"Missing Key\r\n"}
			} else {
				return []string{
					fl(p.CommandLine.command, sep),
					fl(p.CommandLine.key, sep),
				}
			}
		}
	default: 
		return []string{"Invalid Command"}
	}
	
}



