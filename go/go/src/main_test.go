package _go

import (
	"bytes"
	"io/ioutil"
	"testing"

	pb "github.com/bcarlog/gRPC/go/go/src/github.com/bcarlog/gRPC/go/go"
	"google.golang.org/protobuf/proto"
)

// func TestWritePersonWritesPerson(t *testing.T) {
// 	buf := new(bytes.Buffer)
// 	// [START populate_proto]
// 	p := pb.Person{
// 		Id:    1234,
// 		Name:  "John Doe",
// 		Email: "jdoe@example.com",
// 		Phones: []*pb.Person_PhoneNumber{
// 			{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
// 		},
// 	}
// 	// [END populate_proto]
// 	writePerson(buf, &p)
// 	got := buf.String()
// 	want := `Person ID: 1234
//   Name: John Doe
//   E-mail address: jdoe@example.com
//   Home phone #: 555-4321
// `
// 	if got != want {
// 		t.Errorf("writePerson(%s) =>\n\t%q, want %q", p.String(), got, want)
// 	}
// }

func TestCreateProtoData(t *testing.T) {
	protoFile := "./data_generated"
	bookForSave := &pb.AddressBook{
		People: []*pb.Person{
			{
				Id:    1234,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
				},
			},
		},
	}

	out, err := proto.Marshal(bookForSave)
	if err != nil {
		t.Errorf("Failed to encode address book: %s", err)
	}
	// File proto file in local storage
	if err := ioutil.WriteFile(protoFile, out, 0644); err != nil {
		t.Errorf("Failed to write address book: %s", err)
	}

	// Get proto file from local storage
	in, err := ioutil.ReadFile(protoFile)
	if err != nil {
		t.Errorf("Error reading file:%s", err)
	}

	bookForLoad := &pb.AddressBook{}
	if err := proto.Unmarshal(in, bookForLoad); err != nil {
		t.Errorf("Failed to parse address book: %s", err)
	}

	// Compare strings
	bufForSave := new(bytes.Buffer)
	bufForLoad := new(bytes.Buffer)

	writePerson(bufForSave, *&bookForSave.People[0])
	writePerson(bufForLoad, *&bookForLoad.People[0])

	t.Error(bookForLoad)

	if bufForSave.String() != bufForLoad.String() {
		t.Errorf("writePerson() =>\n\t%q, want %q", bufForSave, bufForLoad)
	}
}
