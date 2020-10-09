package main

import (
	"books"
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type abook struct {
	ID   int32
	Name string
	Des  string
	Date string
}

var allBook []abook

type server struct {
}

func (s *server) GetBook(ctx context.Context, request *books.RequestID) (*books.ResponceBook, error) {
	id := request.GetID()
	for i := 0; i < len(allBook); i++ {
		if id == allBook[i].ID {
			return &books.ResponceBook{ID: allBook[i].ID, Name: allBook[i].Name, Des: allBook[i].Des, Date: allBook[i].Date}, nil
		}
	}

	return &books.ResponceBook{ID: 0, Name: "", Des: "", Date: ""}, status.Errorf(codes.InvalidArgument, "ID not found")
}

func (s *server) GetAllBook(ctx context.Context, request *books.None) (*books.AllResponceBook, error) {
	a := &books.AllResponceBook{B: []*books.ResponceBook{
		{ID: 1, Name: "maths", Des: "for 8 std", Date: "01/01/1992"},
		{ID: 2, Name: "science", Des: "for 7 std", Date: "10/08/2010"},
		{ID: 3, Name: "english", Des: "for 6 std", Date: "01/01/1990"},
		{ID: 5, Name: "so. sci.", Des: "for 9 std", Date: "01/10/2006"},
	},
	}
	return a, nil
}

func (s *server) CreateBook(ctx context.Context, request *books.ResponceBook) (*books.None, error) {
	a := abook{ID: request.GetID(), Name: request.GetName(), Des: request.GetDes(), Date: request.GetDate()}
	allBook = append(allBook, a)

	return &books.None{}, nil
}

func (s *server) DeleteBook(ctx context.Context, request *books.RequestID) (*books.None, error) {
	id := request.GetID()

	for index, b := range allBook {
		if id == b.ID {
			allBook = append(allBook[:index], allBook[index+1:]...)
			return &books.None{}, nil
		}
	}

	return &books.None{}, status.Errorf(codes.InvalidArgument, "ID not found")
}

func (s *server) UpdateBook(ctx context.Context, request *books.Updatebook) (*books.ResponceBook, error) {
	id := request.GetId()
	for i := 0; i < len(allBook); i++ {
		if id == allBook[i].ID {

			if allBook[i].Name != request.GetB().Name {
				if request.GetB().Name != "" {
					allBook[i].Name = request.GetB().Name
				}
			}

			if allBook[i].Des != request.GetB().Des {
				if request.GetB().Des != "" {
					allBook[i].Des = request.GetB().Des
				}
			}

			if allBook[i].Date != request.GetB().Date {
				if request.GetB().Date != "" {
					allBook[i].Date = request.GetB().Date
				}
			}
			return &books.ResponceBook{ID: allBook[i].ID, Name: allBook[i].Name, Des: allBook[i].Des, Date: allBook[i].Date}, nil
		}
	}
	return &books.ResponceBook{ID: 0, Name: "", Des: "", Date: ""}, status.Errorf(codes.InvalidArgument, "ID not found")
}

func (s *server) ReplaceBook(ctx context.Context, request *books.Updatebook) (*books.ResponceBook, error) {
	id := request.GetId()
	for i := 0; i < len(allBook); i++ {
		if id == allBook[i].ID {
			allBook[i].Name = request.GetB().Name
			allBook[i].Des = request.GetB().Des
			allBook[i].Date = request.GetB().Date

			return &books.ResponceBook{ID: allBook[i].ID, Name: allBook[i].Name, Des: allBook[i].Des, Date: allBook[i].Date}, nil
		}
	}
	return &books.ResponceBook{ID: 0, Name: "", Des: "", Date: ""}, status.Errorf(codes.InvalidArgument, "ID not found")
}

func main() {
	allBook = []abook{{ID: 1, Name: "maths", Des: "for 8 std", Date: "01/01/1992"}, {ID: 2, Name: "science", Des: "for 7 std", Date: "10/08/2010"}, {ID: 3, Name: "english", Des: "for 6 std", Date: "01/01/1990"}, {ID: 5, Name: "so. sci.", Des: "for 9 std", Date: "01/10/2006"}}

	flag.Parse()
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Println(err)
	}

	serv := grpc.NewServer()
	books.RegisterAllBooksServer(serv, &server{})
	reflection.Register(serv)

	if e := serv.Serve(lis); e != nil {
		log.Println(err)
	}

}
