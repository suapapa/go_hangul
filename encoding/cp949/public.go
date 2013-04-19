// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp949

import (
	"io"
)

// convert utf-8 stream to cp949 stream
func From(in []byte) ([]byte, error) {
	tr, err := fromCp949()
	if err != nil {
		return nil, err
	}

	_, out, err := tr.Translate(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// convert cp949 stream to utf-8 stream
func To(in []byte) ([]byte, error) {
	tr, err := toCp949()
	if err != nil {
		return nil, err
	}

	_, out, err := tr.Translate(in)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type cp949Reader struct {
	r     io.Reader
	tr    translator
	cdata []byte // unconsumed data from converter.
	rdata []byte // unconverted data from reader.
	err   error  // final error from reader.
}

// create io.Reader which read cp949 source to utf-8
func NewReader(r io.Reader) (io.Reader, error) {
	tr, err := fromCp949()
	if err != nil {
		return nil, err
	}

	return &cp949Reader{r: r, tr: tr}, nil
}

func (r *cp949Reader) Read(buf []byte) (int, error) {
	for {
		if len(r.cdata) > 0 {
			n := copy(buf, r.cdata)
			r.cdata = r.cdata[n:]
			return n, nil
		}
		if r.err == nil {
			r.rdata = ensureCap(r.rdata, len(r.rdata)+len(buf))
			n, err := r.r.Read(r.rdata[len(r.rdata):cap(r.rdata)])
			// Guard against non-compliant Readers.
			if n == 0 && err == nil {
				err = io.EOF
			}
			r.rdata = r.rdata[0 : len(r.rdata)+n]
			r.err = err
		} else if len(r.rdata) == 0 {
			break
		}
		nc, cdata, cvterr := r.tr.Translate(r.rdata)
		if cvterr != nil {
			// TODO
		}
		r.cdata = cdata

		// Ensure that we consume all bytes at eof
		// if the converter refuses them.
		if nc == 0 && r.err != nil {
			nc = len(r.rdata)
		}

		// Copy unconsumed data to the start of the rdata buffer.
		r.rdata = r.rdata[0:copy(r.rdata, r.rdata[nc:])]
	}
	return 0, r.err
}

type cp949Writer struct {
	w   io.Writer
	tr  translator
	buf []byte // unconsumed data from writer.
}

// creater io.Writer which write utf8 input src to cp949
func NewWriter(w io.Writer) (io.Writer, error) {
	tr, err := toCp949()
	if err != nil {
		return nil, err
	}

	return &cp949Writer{w: w, tr: tr}, nil
}

func (w *cp949Writer) Write(data []byte) (int, error) {
	wdata := data
	if len(w.buf) > 0 {
		w.buf = append(w.buf, data...)
		wdata = w.buf
	}
	n, cdata, err := w.tr.Translate(wdata)
	if err != nil {
		// TODO
	}
	if n > 0 {
		_, err = w.w.Write(cdata)
		if err != nil {
			return 0, err
		}
	}
	w.buf = w.buf[:0]
	if n < len(wdata) {
		w.buf = append(w.buf, wdata[n:]...)
	}
	return len(data), nil
}
