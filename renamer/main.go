package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

type RenameTarget struct {
	Path string
	Basename string
}

func main() {
	var found []RenameTarget
	path := flag.String("path", "samples", "A path where files will be renamed recursively")
	pattern := flag.String("pattern", "", "A pattern that rename a specific subset of files. It takes all the files that end in _NNN.txt and rename them to instead read (1 of 4).txt")
	flag.Parse()

	if i := strings.Index(*pattern, "_NNN.txt"); i != -1 {
		needle := (*pattern)[:i]

		err := filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}

			if j := strings.Index(path, needle); !info.IsDir() && j != -1 {
				found = append(found, RenameTarget{
					Path:     path,
					Basename: path[:j+len(needle)],
				})
			}
			return nil
		})

		if err != nil {
			panic(err)
		}

		for k, rt := range found {
			err = os.Rename(rt.Path, fmt.Sprintf("%v_(%d of %d).txt", rt.Basename, k+1, len(found)))
			if err != nil {
				panic(err)
			}
		}
	}


}
