package kgo

import (
	"strings"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	str := []byte("This is an string to encod")
	res := KEncr.Base64Encode(str)
	if !strings.HasSuffix(res, "=") {
		t.Error("Base64Encode fail")
		return
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	b.ResetTimer()
	str := []byte("This is an string to encod")
	for i := 0; i < b.N; i++ {
		KEncr.Base64Encode(str)
	}
}

func TestBase64Decode(t *testing.T) {
	str := "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="
	_, err := KEncr.Base64Decode(str)
	if err != nil {
		t.Error("Base64Decode fail")
		return
	}
	_, err = KEncr.Base64Decode("#iu3498r")
	if err == nil {
		t.Error("Base64Decode fail")
		return
	}
	_, err = KEncr.Base64Decode("VGhpcy")
	_, err = KEncr.Base64Decode("VGhpcyB")
}

func BenchmarkBase64Decode(b *testing.B) {
	b.ResetTimer()
	str := "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.Base64Decode(str)
	}
}

func TestBase64UrlEncodeDecode(t *testing.T) {
	str := []byte("This is an string to encod")
	res := KEncr.Base64UrlEncode(str)
	if strings.HasSuffix(res, "=") {
		t.Error("Base64UrlEncode fail")
		return
	}

	_, err := KEncr.Base64UrlDecode(res)
	if err != nil {
		t.Error("Base64UrlDecode fail")
		return
	}
}

func BenchmarkBase64UrlEncode(b *testing.B) {
	b.ResetTimer()
	str := []byte("This is an string to encod")
	for i := 0; i < b.N; i++ {
		KEncr.Base64UrlEncode(str)
	}
}

func BenchmarkBase64UrlDecode(b *testing.B) {
	b.ResetTimer()
	str := "VGhpcyBpcyBhbiBzdHJpbmcgdG8gZW5jb2Q"
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.Base64UrlDecode(str)
	}
}

func TestAuthCode(t *testing.T) {
	key := "123456"
	str := "hello world"

	res := KEncr.AuthCode(str, key, true, 0)
	if res == "" {
		t.Error("AuthCode Encode fail")
		return
	}

	res2 := KEncr.AuthCode(res, key, false, 0)
	if res2 == "" {
		t.Error("AuthCode Decode fail")
		return
	}
	KEncr.AuthCode(str, key, true, 1800)
	KEncr.AuthCode("", key, true, 0)
	KEncr.AuthCode("", "", true, 0)
}

func BenchmarkAuthCodeEncode(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "hello world"
	for i := 0; i < b.N; i++ {
		KEncr.AuthCode(str, key, true, 0)
	}
}

func BenchmarkAuthCodeDecode(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "a79b5do3C9nbaZsAz5j3NQRj4e/L6N+y5fs2U9r1mO0LinOWtxmscg=="
	for i := 0; i < b.N; i++ {
		KEncr.AuthCode(str, key, false, 0)
	}
}
