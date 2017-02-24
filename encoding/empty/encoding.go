package empty

import (
	"bufio"
	"io"

	"github.com/goadesign/goa"
)

type ResettableDecoder struct {
	D Decoder
}

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) goa.ResettableDecoder {
	return ResettableDecoder{D.r: r}
}

func (d *Decoder) Decode(dst interface{}) error {

	br := bufio.NewReader(d.r)
	_, err := br.WriteTo(dst)
	if err != nil {
		return err
	}

	return nil
}

func (d *ResettableDecoder) Reset(r io.Reader) {

	d.D.r = r

	return
}
