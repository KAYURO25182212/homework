package main 
import (
	"fmt"
)
func main(){
	
	book1 := Book{Title: "1984", Author: "George Orwell"}
	book2 := Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"}
	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)
	for _, book := range library.Books {
	  fmt.Println(book.GetDetails())
	}
	
}
type Book struct {
	Title  string
	Author string
  }
  
  func (b Book) GetDetails() string {
	return  fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
  }
  
  type Library struct {
	Books []Book
  }
  
  func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
  }

  type Student struct {
	grades string
  }
  