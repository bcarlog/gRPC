package _go

import (
	"fmt"
	"io"

	pb "github.com/bcarlog/gRPC/go/go/src/github.com/bcarlog/gRPC/go/go"
)

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.PhoneType_PHONE_TYPE_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case pb.PhoneType_PHONE_TYPE_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case pb.PhoneType_PHONE_TYPE_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}
