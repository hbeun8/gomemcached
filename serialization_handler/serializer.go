package serializationhandler

import "bytes"
import "os"

func Serializer(s [][]byte) {
	sep := []byte("\r\n")
	for i:=range s {
	k :=bytes.Split(s[i], sep)
		for j:= range k {
		var b bytes.Buffer
			b.Write(k[j])
			b.WriteTo(os.Stdout)
		//return k[i]
		}
	} 
}

func Serializer2(s [][]byte) []byte{
	buf := bytes.Buffer{}
	for i:=range s {
		buf.Write(s[i])
		//os.Stdout.Write(buf.Bytes())
	}
	return buf.Bytes()
}

func SerializerforTest(s [][]byte) []byte{
	sep := []byte("\r\n")
	for i := range s {
		k :=bytes.Split(s[i], sep)
		for i:= range k {
			var b bytes.Buffer
			b.Write(k[i])
			//b.WriteTo(os.Stdout)
			return k[i]
		} 
	}
	return s[0]
}

/*
func main() {
s := [][]byte{[]byte("CLIENT_ERROR bad data chunk\r\n"), []byte("Error\r\n")}
//Serializer(s)
S(s)
}
*/

