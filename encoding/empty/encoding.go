package empty

import (
	"bufio"
	"io"

	"github.com/goadesign/goa"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) goa.ResettableDecoder {

	return Decoder{r}
}

func (d Decoder) Decode(dst interface{}) error {

	br := bufio.NewReader(d.r)

	_, err := br.Read(dst.([]byte))
	if err != nil {
		return err
	}

	return nil
}

func (d Decoder) Reset(r io.Reader) {

	d.r = r

	return
}
