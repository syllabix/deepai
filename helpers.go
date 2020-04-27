package deepai

import (
	"io"
	"mime/multipart"
)

func imageForm(file io.Reader) (*multipart.Writer, io.Reader) {
	pr, pw := io.Pipe()
	mpw := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()

		formPart, err := mpw.CreateFormFile("image", "target")
		if err != nil {
			pw.CloseWithError(err)
		}

		_, err = io.Copy(formPart, file)
		if err != nil {
			pw.CloseWithError(err)
		}

		err = mpw.Close()
		if err != nil {
			pw.CloseWithError(err)
		}
	}()

	return mpw, pr
}
