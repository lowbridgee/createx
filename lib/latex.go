package lib

import (
	"os"
	"io"
)

func LatexTemplateCopy(filename string) {
	srcName := "/Users/lowbridge/Documents/go/createx/template/latex.tex"
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if  err != nil {
		panic(err)
	}
	//_ = os.Link(src, filename)
}
