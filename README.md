A web server that allows the user to add and search for people from a map in Golang.

Add new people to the map by sending a POST request to /people with the following JSON:
```
type Person struct {
    name string
    age int
    profession string
    hairColor string
}
```
Display a list of all the people in the map by visiting /people

Search of a specific person by name by visiting /people/{name}

Start the server using:
```
go run main.go
```
