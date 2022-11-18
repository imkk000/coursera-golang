package main

import (
	"fmt"
	"io"
	"os"
)

/*
1. create data file in names.txt
cat << EOF > names.txt
a b
c d
f g
h i
j k
l m
n o
p q
r s
t u
v w
x y
z z

EOF

2. compile and run go code and input names.txt
echo 'names.txt' | go run read.go
*/

type Name struct {
	fname string
	lname string
}

func main() {
	// input filename for reading data
	var filename string
	fmt.Print("enter Name of the text file: ")
	fmt.Scan(&filename)

	// open file
	fs, err := os.Open(filename)
	if err != nil {
		fatal(err, "failed to open file")
	}
	defer func() {
		if err = fs.Close(); err != nil {
			fatal(err, "failed to close file")
		}
	}()

	// read data into filename
	var names []Name
	var name Name
	for {
		_, err = fmt.Fscanf(fs, "%s %s", &name.fname, &name.lname)
		if err == io.EOF {
			break
		}
		names = append(names, name)
	}

	// print all data
	fmt.Printf("\nfound %d:\n", len(names))
	for i, n := range names {
		fmt.Printf("%2d %s %s\n", i+1, n.fname, n.lname)
	}
}

// fatal printing message and exit with status 1
func fatal(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		os.Exit(1)
	}
}
