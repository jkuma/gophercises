/*
	This package exposes a command line to massively rename files
	based on the pattern bellow:
		- xxxxxx_NNN.ext

	The matched files are renamed to xxxxxx (index of nbfiles).ext

	Flags available:
		- path: the complete path to directory of file to be scanned.
		- dry: if true, run the commands without modifying files.
*/

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func init() {
	os.Mkdir("samples", os.ModePerm)
	os.Create("./samples/birthday_001.txt")
	os.Create("./samples/birthday_002.txt")
	os.Create("./samples/birthday_003.txt")
	os.Create("./samples/birthday_004.txt")
	os.Create("./samples/christmas 2016 (1 of 100).txt")
	os.Create("./samples/christmas 2016 (2 of 100).txt")
	os.Create("./samples/christmas 2016 (3 of 100).txt")
	os.Create("./samples/christmas 2016 (4 of 100).txt")
	os.Create("./samples/christmas 2016 (5 of 100).txt")
	os.Mkdir("samples/nested", os.ModePerm)
	os.Create("./samples/nested/birthday_005.txt")
	os.Create("./samples/nested/n_008.txt")
	os.Create("./samples/nested/n_009.txt")
	os.Create("./samples/nested/n_010.txt")
}

func main() {
	rename := make(map[string][]MatchResult)
	p := flag.String("path", "samples", "A path where files will be renamed recursively")
	d := flag.Bool("dry", false, "Do everything apart from actually renaming the files")
	flag.Parse()

	err := filepath.Walk(*p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		mr, err := match(path)

		if err == nil {
			rename[mr.Basename] = append(rename[mr.Basename], mr)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, mrs := range rename {
		for _, mr := range mrs {
			n := fmt.Sprintf("%s/%s (%d of %d).%s", filepath.Dir(mr.Filename), mr.Basename, mr.Index, len(mrs), mr.Extension)
			fmt.Printf("File %v is renamed to %v\n", mr.Filename, n)

			if *d {
				os.Rename(mr.Filename, n)
			}
		}
	}
}

type MatchResult struct {
	Filename  string
	Basename  string
	Index     int
	Extension string
}

func match(filename string) (MatchResult, error) {
	results := make([][][]byte, 1)
	re := regexp.MustCompile(`(\w+)_(\d{3})\.(\w+)`)
	results = re.FindAllSubmatch([]byte(filename), -1)

	if len(results) > 0 {
		i, err := strconv.Atoi(string(results[0][2]))

		if err != nil {
			panic(err)
		}

		return MatchResult{
			Filename:  filename,
			Basename:  string(results[0][1]),
			Index:     i,
			Extension: string(results[0][3]),
		}, nil
	}

	return MatchResult{}, errors.New(`filename does not match our pattern (\w+)_(\d{3})\.(\w+)`)
}
