ipackage main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type file struct {
	name        string
	directory   *directory
	tags        string
	isDirectory bool
}

type directory struct {
	id     string
	name   string
	master *directory
	files  []file
}

type server struct {
	name        string
	description string
	url         string
	directories []directory
}

type files []file

func main() {
	servers := []server{
		server{
			"localhost",
			"localhost",
			"sync://127.0.0.1:50123",
			[]directory{
				directory{
					"Alice",
					"/home/alice",
					nil,
					[]file{},
				},
				directory{
					"Bob",
					"/home/bob",
					nil,
					[]file{},
				},
			},
		},
	}

	// Read the directories..
	for _, server := range servers {
		fmt.Printf("Server: %v\n", server.name)
		for _, directory := range server.directories {
			fmt.Printf(" > directory: %v\n", directory.name)
			directory.files = readDirectory(&directory, directory.name)
			for _, f := range directory.files {
				fmt.Printf(" > file: %v\n", f.name)
			}
		}
	}
}

func readDirectory(directory *directory, folder string) files {
	value := files{}

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal()
	}

	for _, f := range files {
		isDirectory := f.IsDir()
		value = append(value, file{folder + "/" + f.Name(), directory, "", isDirectory})
		if isDirectory {
			value = append(value, readDirectory(directory, folder+"/"+f.Name())...)
			continue
		}
		value = append(value, file{folder + "/" + f.Name(), directory, "", isDirectory})
	}

	return value

}

/*
 Tagy souborů budou jen na serveru..
*/

