package common

import (
	"io"
)

type Forwarder struct {
	encodePassword *Lspasswd
	decodePassword *Lspasswd
}

func NewForwarder(password *Lspasswd) *Forwarder {
	return &Forwarder{
		encodePassword: password,
		decodePassword: GetDecodePasswd(password),
	}
}

func NewForwarderWithStr(passwordStr string) (*Forwarder, error) {
	password, err := Atp(passwordStr)
	if err != nil {
		return nil, err
	}
	return &Forwarder{
		encodePassword: password,
		decodePassword: GetDecodePasswd(password),
	}, nil
}

func (f *Forwarder) Encode(bs []byte) {
	for i, v := range bs {
		bs[i] = f.encodePassword[v]
	}
}

func (f *Forwarder) Decode(bs []byte) {
	for i, v := range bs {
		bs[i] = f.decodePassword[v]
	}
}

func (f *Forwarder) EncodeWrite(dst io.Writer, buf []byte) (int, error) {
	f.Encode(buf)
	writeCount, errWrite := dst.Write(buf)
	if errWrite != nil {
		return 0, errWrite
	}
	return writeCount, nil
}

func (f *Forwarder) EncodeCopy(dst io.Writer, rst io.Reader) error {
	buf := BufferPoolGet()
	defer BufferPoolPut(buf)
	for {
		readCount, errRead := rst.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, err := f.EncodeWrite(dst, buf[:readCount])
			if err != nil {
				return err
			}
			if writeCount != readCount {
				return io.ErrShortWrite
			}
		}
	}
}

func (f *Forwarder) DecodeRead(rst io.Reader, buf []byte) (int, error) {
	readCount, errRead := rst.Read(buf)
	if errRead != nil {
		if errRead != io.EOF {
			return 0, errRead
		} else {
			return 0, nil
		}
	}
	f.Decode(buf)
	return readCount, nil
}

func (f *Forwarder) DecodeCopy(dst io.Writer, rst io.Reader) error {
	buf := BufferPoolGet()
	defer BufferPoolPut(buf)

	for {
		readCount, errRead := f.DecodeRead(rst, buf)
		if errRead != nil {
			return errRead
		}
		if readCount > 0 {
			writeCount, errWrite := dst.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}
