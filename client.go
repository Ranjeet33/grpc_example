package main

import (
	"books"
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := books.NewAllBooksClient(conn)

	book, err := client.GetBook(context.Background(), &books.RequestID{ID: 2})
	if err != nil {
		log.Println(err)
	}
	log.Println(book)

	allBook, err := client.GetAllBook(context.Background(), &books.None{})
	if err != nil {
		log.Println(err)
	}
	log.Println(allBook)

	createBook, err := client.CreateBook(context.Background(), &books.ResponceBook{ID: 4, Name: "physics", Des: "for 10 std.", Date: "01/06/1998"})
	if err != nil {
		log.Println(err)
	}
	log.Println(createBook)

	replaceBook, err := client.ReplaceBook(context.Background(), &books.Updatebook{Id: 3, B: &books.ResponceBook{ID: 5, Name: "chemistry", Des: "for 11 std.", Date: "01/07/1999"}})
	if err != nil {
		log.Println(err)
	}
	log.Println(replaceBook)

	updateBook, err := client.UpdateBook(context.Background(), &books.Updatebook{Id: 2, B: &books.ResponceBook{Name: "hindi", Des: "for 9 std."}})
	if err != nil {
		log.Println(err)
	}
	log.Println(updateBook)

	deleteBook, err := client.DeleteBook(context.Background(), &books.RequestID{ID: 5})
	if err != nil {
		log.Println(err)
	}
	log.Println(deleteBook)

}
