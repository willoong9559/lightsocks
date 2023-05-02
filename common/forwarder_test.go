package common_test

import (
	"bytes"
	"testing"

	"github.com/willoong9559/lightsocks/common"
)

const psk = "/DonXhfDlabZ4JJ2mvs9ob5dZM54KhEsIu6se1oHowilQTSx0iGFUlDdPzd3H9qqmc26hDZo5Js8xvYbjrWMWbDPMpfppNhUIw+LuImgs6d1cPKQ8a1980vR//SRx8RP3uzhZRQSn2vnwQFzVuL6L0hGllV/iEA+OMyv18rCwO9Y6/gVvB2eArYt5SRFGlG3Ho0OoqiAhyvjbRnqyKnV37nbsmD1M+05ECiuW0n9BQlXbHFyDZQ7isvULl90AIYwnEIpq0d6TQrcIBzF8L9D0G9EC/7J1oF8fp0YFhNOYgT3Z1zmbpO9DFO003m7Jo9Mavlp6GZKBpiCg2ExAzUlYw=="

var forwarder *common.Forwarder

func init() {
	var err error
	forwarder, err = common.NewForwarderWithStr(psk)
	if err != nil {
		panic(err)
	}
}

func TestEncodeAndDecode(t *testing.T) {
	text := []byte("hello world     ")
	forwarder.Encode(text)
	forwarder.Decode(text)
	if !bytes.Equal(text, []byte("hello world     ")) {
		t.Errorf("encode and decode 验证失败，原%s, 解密后%s", "hello world     ", text)
	}
}

type TestWriterAndReader struct {
	Data string
}

func (t *TestWriterAndReader) Write(p []byte) (int, error) {
	t.Data = string(p)
	return len(p), nil
}

func (t *TestWriterAndReader) Read(p []byte) (int, error) {
	n := copy(p, []byte(t.Data))
	return n, nil
}

func TestEncodeWriteAndDecodeRead(t *testing.T) {
	testWriterAndReader := new(TestWriterAndReader)
	text, text1 := []byte("hello"), []byte("hello")
	text2 := make([]byte, 1024)
	nw, _ := forwarder.EncodeWrite(testWriterAndReader, text)
	dr, _ := forwarder.DecodeRead(testWriterAndReader, text2)
	if nw != dr {
		t.Errorf("encode len != decode len: nw = %d, dr = %d", nw, dr)
	}
	if !bytes.Equal(text1, text2[:dr]) {
		t.Error("TestEncodeWriteAndDecodeRead failed")
	}
}
