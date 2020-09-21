package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// unzip function will return string and error.
func unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)

	if err != nil {
		return filenames, err
	}

	// defer makes sure the file is closed
	// at the end of the program no matter what.
	defer r.Close()

	// this loop will run until there are files in the source directory.
	for _, f := range r.File {
		// store path/filename for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// check for invalid filetypes
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("[-]%s is an illegal filepath", fpath)
		}
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// create a new folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		// creates files in the directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}
		// the created file will be stored in outfile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}

	}
	return filenames, nil
}
