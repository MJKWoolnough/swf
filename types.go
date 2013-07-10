// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package swf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/MJKWoolnough/rwcount"
	"io"
	"math"
)

type Tag interface {
	io.ReaderFrom
	io.WriterTo
	MinVersion() uint8
	Name() string
	Size() int32
	TagId() uint16
}

type Upgradeable interface {
	MaxVersion() uint8
	Upgrade(uint8) Tag
}

type Sizer interface {
	Size() int32
}

type Int8 int8

func NewInt8(n int8) *Int8 {
	a := Int8(n)
	return &a
}

func (i *Int8) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, i)
	return
}

func (i *Int8) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, i)
	return
}

func (i *Int8) Size() int32 {
	return 1
}

func (i *Int8) String() string {
	return fmt.Sprintf("%d", *i)
}

type Int16 int16

func NewInt16(n int16) *Int16 {
	a := Int16(n)
	return &a
}

func (i *Int16) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, i)
	return
}

func (i *Int16) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, i)
	return
}

func (i *Int16) Size() int32 {
	return 2
}

func (i *Int16) String() string {
	return fmt.Sprintf("%d", *i)
}

type Int32 int32

func NewInt32(n int32) *Int32 {
	a := Int32(n)
	return &a
}

func (i *Int32) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, i)
	return
}

func (i *Int32) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, i)
	return
}

func (i *Int32) Size() int32 {
	return 4
}

func (i *Int32) String() string {
	return fmt.Sprintf("%d", *i)
}

type Int64 int64

func NewInt64(n int64) *Int64 {
	a := Int64(n)
	return &a
}

func (i *Int64) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, i)
	return
}

func (i *Int64) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, i)
	return
}

func (i *Int64) Size() int32 {
	return 8
}

func (i *Int64) String() string {
	return fmt.Sprintf("%d", *i)
}

type Uint8 uint8

func NewUint8(n uint8) *Uint8 {
	a := Uint8(n)
	return &a
}

func (u *Uint8) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, u)
	return
}

func (u *Uint8) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, u)
	return
}

func (u *Uint8) Size() int32 {
	return 1
}

func (u *Uint8) String() string {
	return fmt.Sprintf("%d", *u)
}

type Uint16 uint16

func NewUint16(n uint16) *Uint16 {
	a := Uint16(n)
	return &a
}

func (u *Uint16) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, u)
	return
}

func (u *Uint16) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, u)
	return
}

func (u *Uint16) Size() int32 {
	return 2
}

func (u *Uint16) String() string {
	return fmt.Sprintf("%d", *u)
}

type Uint32 uint32

func NewUint32(n uint32) *Uint32 {
	a := Uint32(n)
	return &a
}

func (u *Uint32) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, u)
	return
}

func (u *Uint32) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, u)
	return
}

func (u *Uint32) Size() int32 {
	return 4
}

func (u *Uint32) String() string {
	return fmt.Sprintf("%d", *u)
}

type Uint64 uint64

func NewUint64(n uint64) *Uint64 {
	a := Uint64(n)
	return &a
}

func (u *Uint64) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, u)
	return
}

func (u *Uint64) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, u)
	return
}

func (u *Uint64) Size() int32 {
	return 8
}

func (u *Uint64) String() string {
	return fmt.Sprintf("%d", *u)
}

type Float16 float32

func NewFloat16(n float32) *Float16 {
	a := Float16(n)
	return &a
}

func (i *Float16) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	var d uint16
	if err = binary.Read(c, binary.LittleEndian, &d); err == nil || err == io.EOF {
		var bits uint32
		if d&0x7FFF == 0 {
			bits = uint32(d) << 16
		} else {
			sign := d & 0x8000
			exponent := d & 0x7C00
			mantissa := d & 0x03FF
			if exponent == 0 {
				var e uint32
				mantissa <<= 1
				for e = 0; mantissa&0x400 == 0; e++ {
					mantissa <<= 1
				}
				newSign := uint32(sign) << 16
				newExponent := (uint32(exponent>>10) - 15 + 127 - e) << 23
				newMantissa := uint32(mantissa&0x03FF) << 13
				bits = newSign | newExponent | newMantissa
			} else if exponent == 0x7C00 {
				if mantissa == 0 {
					bits = uint32(sign)<<16 | 0x7F800000 // +/- Inf
				} else {
					bits = 0xFFC00000 //NaN
				}
			} else {
				newSign := uint32(sign) << 16
				newExponent := uint32(exponent>>10-15+127) << 23
				newMantissa := uint32(mantissa) << 13
				bits = newSign | newExponent | newMantissa
			}
		}
		*i = Float16(math.Float32frombits(bits))
	}
	return
}

func (f *Float16) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	var d uint16
	bits := math.Float32bits(float32(*f))
	sign := bits & 0x80000000
	exponent := bits & 0x7F800000
	mantissa := bits & 0x007FFFFF
	if exponent == 0 {
		d = uint16(bits >> 16)
	} else if exponent == 0x7F800000 {
		if mantissa == 0 {
			d = uint16((sign >> 16) | 0x7C00) // +/- Inf
		} else {
			d = 0xFE00 //NaN
		}
	} else {
		d = uint16(sign >> 16)
		unbiased := int32(exponent>>23) - 127 + 15
		if unbiased > 31 {
			d |= 0x7C00 // +/- inf
		} else if unbiased <= 0 {
			var newMantissa uint16
			if unbiased < -10 { //14 - unbiased > 24
				newMantissa = 0
			} else {
				mantissa |= 0x800000
				newMantissa = uint16(mantissa >> uint32(14-unbiased))
				if mantissa>>uint32(13-unbiased)&1 == 1 {
					newMantissa++
				}
			}
			d |= newMantissa
		} else {
			newExponent := uint16(unbiased << 10)
			newMantissa := uint16(mantissa >> 13)
			d |= newExponent | newMantissa
			if mantissa&0x1000 > 1 {
				d++
			}
		}

	}
	err = binary.Write(c, binary.LittleEndian, d)
	return
}

func (f *Float16) Size() int32 {
	return 2
}

func (f *Float16) String() string {
	return fmt.Sprintf("%g", *f)
}

type Float float32

func NewFloat(n float32) *Float {
	a := Float(n)
	return &a
}

func (f *Float) ReadFrom(fr io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: fr}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, f)
	return
}

func (f *Float) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, f)
	return
}

func (f *Float) Size() int32 {
	return 4
}

func (f *Float) String() string {
	return fmt.Sprintf("%g", *f)
}

type Double float64

func NewDouble(n float64) *Double {
	a := Double(n)
	return &a
}

func (d *Double) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, d)
	return
}

func (d *Double) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, d)
	return
}

func (d *Double) Size() int32 {
	return 8
}

func (d *Double) String() string {
	return fmt.Sprintf("%g", *d)
}

type BitReader interface {
	io.Reader
	Align()
	ReadBits([]bool) error
}

type bitReader struct {
	io.Reader
	data byte
	left uint8
}

func (b *bitReader) Read(d []byte) (int, error) {
	b.left = 0
	return b.Reader.Read(d)
}

func (b *bitReader) ReadBits(bits []bool) (err error) {
	for i := 0; i < len(bits); i++ {
		if b.left == 0 {
			err = binary.Read(b.Reader, binary.LittleEndian, &b.data)
			var t [8]bool
			for i := uint(0); i < 8; i++ {
				t[i] = b.data>>(8-i-1)&1 > 0
			}
			if err != nil {
				if err != io.EOF || len(bits)-i > 8 {
					return
				}
			}
			b.left = 8
		}
		bits[i] = b.data>>(b.left-1)&1 == 1
		b.left--
	}
	return
}

func (b *bitReader) Align() {
	b.left = 0
}

type BitWriter interface {
	io.Writer
	Align()
	WriteBits([]bool) error
}

type bitWriter struct {
	io.Writer
	data byte
	left uint8
}

func (b *bitWriter) Write(d []byte) (int, error) {
	b.Align()
	return b.Writer.Write(d)
}

func (b *bitWriter) WriteBits(bits []bool) (err error) {
	for _, bit := range bits {
		if bit {
			b.data |= 1 << (7 - b.left)
		}
		if b.left == 7 {
			b.left = 0
			_, err = b.Writer.Write([]byte{b.data})
			b.data = 0
			if err != nil {
				break
			}
		} else {
			b.left++
		}
	}
	return
}

func (b *bitWriter) Align() {
	if b.left > 0 {
		b.left = 0
		b.Writer.Write([]byte{b.data})
	}
	b.data = 0
}

type Twips int32

func NewTwips(n int32) *Twips {
	a := Twips(n)
	return &a
}

func (t *Twips) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, t)
	return
}

func (t *Twips) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, t)
	return
}

func (t *Twips) Size() int32 {
	return 4
}

func (t Twips) String() string {
	n := t / 20
	d := t % 20
	s := fmt.Sprintf("%d", n)
	if d > 0 {
		s += fmt.Sprintf(" %d/20", d)
	}
	return s
}

type Fixed float64

func NewFixed(n float64) *Fixed {
	a := Fixed(n)
	return &a
}

func (i *Fixed) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	var d uint32
	if err = binary.Read(c, binary.LittleEndian, &d); err == nil || err == io.EOF {
		*i = Fixed(d) / 65536
	}
	return
}

func (f *Fixed) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, uint32(*f*65536))
	return
}

func (f *Fixed) Size() int32 {
	return 4
}

func (f *Fixed) String() string {
	return fmt.Sprintf("%f", *f)
}

type Fixed8 float32

func NewFixed8(n float32) *Fixed8 {
	a := Fixed8(n)
	return &a
}

func (i *Fixed8) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	var d uint16
	if err = binary.Read(c, binary.LittleEndian, &d); err == nil || err == io.EOF {
		*i = Fixed8(d) / 256
	}
	return
}

func (f *Fixed8) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, uint16(*f*256))
	return
}

func (f *Fixed8) Size() int32 {
	return 2
}

func (f *Fixed8) String() string {
	return fmt.Sprintf("%f", *f)
}

type EncodedU32 uint32

func NewEncodedU32(n float64) *EncodedU32 {
	a := EncodedU32(n)
	return &a
}

func (e *EncodedU32) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	shift := EncodedU32(0)
	b := []byte{255}
	read := 0
	for b[0]>>7 == 1 {
		_, err = c.Read(b)
		if err != nil && err != io.EOF {
			return
		}
		read++
		if read == 5 && b[0] > 15 {
			err = errors.New("encodedU32: malformed data")
			return
		}
		*e |= EncodedU32(b[0]) << shift
		shift += 7
	}
	return
}

func (e *EncodedU32) WriteTo(w io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: w}
	defer func() { total = c.BytesWritten() }()
	data := make([]byte, 5)
	f := uint32(*e)
	g := 0
	for i := 0; i < 5; i++ {
		data[i] = byte(f & 127)
		f >>= 7
		if data[i] != 0 || i == 0 {
			g = i
		}
	}
	for i := 0; i < g; i++ {
		data[i] |= 128
	}
	_, err = c.Write(data[:g+1])
	return
}

func (e *EncodedU32) Size() int32 {
	var i int32
	j := EncodedU32(1)
	for i = 1; i < 5; i++ {
		j *= 128
		if *e < j {
			break
		}
	}
	return i
}

func (e *EncodedU32) String() string {
	return fmt.Sprintf("%d", *e)
}

type BitUint uint32

func NewBitUint(n uint32) *BitUint {
	a := BitUint(n)
	return &a
}

func (b *BitUint) ReadBitsFrom(f BitReader, n uint8) (err error) {
	*b = 0
	bits := make([]bool, n, n)
	if err = f.ReadBits(bits); err != nil && err != io.EOF {
		return
	}
	for i := uint8(0); i < n; i++ {
		*b <<= 1
		if bits[i] {
			*b++
		}
	}
	return
}

func (b *BitUint) WriteBitsTo(f BitWriter, n uint8) (err error) {
	data := make([]bool, n)
	c := uint32(*b)
	for i := uint8(0); i < n; i++ {
		data[n-i-1] = c&1 == 1
		c >>= 1
	}
	err = f.WriteBits(data)
	return
}

func (b *BitUint) Size() int32 {
	bd := uint32(*b)
	pos := int32(32)
	for bit := uint32(1) << 31; bit > 0; bit >>= 1 {
		if bd&bit != 0 {
			return pos
		}
		pos--
	}
	return 1
}

func (b *BitUint) String() string {
	return fmt.Sprintf("%d", *b)
}

type BitInt int32

func NewBitInt(n int32) *BitInt {
	a := BitInt(n)
	return &a
}

func (b *BitInt) ReadBitsFrom(f BitReader, n uint8) (err error) {
	*b = 0
	if n == 0 || n > 32 {
		return
	}
	bits := make([]bool, 32)
	if err = f.ReadBits(bits[32-n:]); err != nil && err != io.EOF {
		return
	}
	if bits[32-n] {
		for i := uint8(0); i < 32-n; i++ {
			bits[i] = true
		}
	}
	for i := 0; i < 32; i++ {
		*b <<= 1
		if bits[i] {
			*b++
		}
	}
	return
}

func (b *BitInt) WriteBitsTo(f BitWriter, n uint8) (err error) {
	data := make([]bool, n)
	c := uint32(*b)
	for i := uint8(0); i < n; i++ {
		data[n-i-1] = c&1 == 1
		c >>= 1
	}
	err = f.WriteBits(data)
	return
}

func (b *BitInt) Size() int32 {
	bd := uint32(*b)
	hb := (bd >> 31) == 0
	pos := int32(32)
	for bit := uint32(1) << 30; bit > 0; bit >>= 1 {
		if (bd&bit == 0) != hb {
			return pos
		}
		pos--
	}
	return 1
}

func (b *BitInt) String() string {
	return fmt.Sprintf("%d", *b)
}

type BitFixed float64

func NewBitFixed(n float64) *BitFixed {
	a := BitFixed(n)
	return &a
}

func (b *BitFixed) ReadBitsFrom(f BitReader, n uint8) (err error) {
	var bI BitInt
	err = bI.ReadBitsFrom(f, n)
	*b = BitFixed(bI) / 65536
	return
}

func (b *BitFixed) WriteBitsTo(f BitWriter, n uint8) (err error) {
	c := BitInt(*b * 65536)
	err = c.WriteBitsTo(f, n)
	return
}

func (b *BitFixed) Size() int32 {
	c := BitInt(*b * 65536)
	return c.Size()
}

func (b *BitFixed) String() string {
	return fmt.Sprintf("%f", *b)
}

type String string

func NewString(s string) *String {
	a := String(s)
	return &a
}

func (s *String) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	data := make([]byte, 0)
	b := []byte{255}
	var n int
	for {
		n, err = c.Read(b)
		if err != nil && n == 0 {
			return
		}
		if b[0] == 0 {
			break
		}
		data = append(data, b[0])
	}
	*s = String(data)
	return
}

func (s *String) WriteTo(f io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	if err = binary.Write(c, binary.LittleEndian, []byte(*s)); err != nil {
		return
	}
	err = binary.Write(c, binary.LittleEndian, byte(0))
	return
}

func (s *String) Size() int32 {
	return int32(len(*s)) + 1
}

func (s *String) String() string {
	return string(*s)
}

const (
	LANGUAGE_LATIN LanguageCode = iota + 1
	LANGUAGE_JAPENESE
	LANGUAGE_KOREAN
	LANGUAGE_SIMPLIFIED_CHINESE
	LANGUAGE_TRADITIONAL_CHINESE
)

type LanguageCode uint8

func NewLanguageCode(n uint8) *LanguageCode {
	a := LanguageCode(n)
	return &a
}

func (l *LanguageCode) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, l)
	if *l == 0 || *l > 5 {
		err = errors.New("languageCode: unknown id")
	}
	return
}

func (l *LanguageCode) WriteTo(f io.Writer) (total int64, err error) {
	if *l == 0 || *l > 5 {
		err = errors.New("languageCode: unknown id")
		return
	}
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, l)
	return
}

func (l *LanguageCode) Size() int32 {
	return 1
}

func (l *LanguageCode) String() string {
	switch *l {
	case LANGUAGE_LATIN:
		return "Latin"
	case LANGUAGE_JAPENESE:
		return "Japenese"
	case LANGUAGE_KOREAN:
		return "Korean"
	case LANGUAGE_SIMPLIFIED_CHINESE:
		return "Simplified Chinese"
	case LANGUAGE_TRADITIONAL_CHINESE:
		return "Traditional Chinese"
	}
	return "Unknown language code"
}

type RGB struct {
	Red, Green, Blue uint8
}

func NewRGB(r, g, b uint8) *RGB {
	return &RGB{r, g, b}
}

func (r *RGB) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, r)
	return
}

func (r *RGB) WriteTo(f io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, r)
	return
}

func (r *RGB) Size() int32 {
	return 3
}

func (r *RGB) String() string {
	return fmt.Sprintf("Red: %d, Green: %d, Blue: %d", r.Red, r.Green, r.Blue)
}

type RGBA struct {
	RGB
	Alpha uint8
}

func NewRGBA(r, g, b, a uint8) *RGBA {
	return &RGBA{RGB{r, g, b}, a}
}

func (r *RGBA) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, r)
	return
}

func (r *RGBA) WriteTo(f io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, r)
	return
}

func (r *RGBA) Size() int32 {
	return 4
}

func (r *RGBA) String() string {
	return fmt.Sprintf("%s, Alpha: %d", r.RGB.String(), r.Alpha)
}

type ARGB struct {
	Alpha uint8
	RGB
}

func NewARGB(a, r, g, b uint8) *ARGB {
	return &ARGB{a, RGB{r, g, b}}
}

func (a *ARGB) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	err = binary.Read(c, binary.LittleEndian, a)
	return
}

func (a *ARGB) WriteTo(f io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	err = binary.Write(c, binary.LittleEndian, a)
	return
}

func (a *ARGB) Size() int32 {
	return 4
}

func (a *ARGB) String() string {
	return fmt.Sprintf("Alpha: %d, %s", a.Alpha, a.RGB.String())
}

type Rect struct {
	Xmin, Xmax, Ymin, Ymax Twips
}

func NewRect(Xmin, Xmax, Ymin, Ymax Twips) *Rect {
	return &Rect{Xmin, Xmax, Ymin, Ymax}
}

func (r *Rect) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	b := &bitReader{Reader: c}
	var (
		nBits BitUint
		t     BitInt
	)
	err = nBits.ReadBitsFrom(b, 5)
	if err != nil {
		return
	}
	nBits8 := uint8(nBits)
	err = t.ReadBitsFrom(b, uint8(nBits8))
	r.Xmin = Twips(t)
	if err != nil {
		return
	}
	err = t.ReadBitsFrom(b, uint8(nBits8))
	r.Xmax = Twips(t)
	if err != nil {
		return
	}
	err = t.ReadBitsFrom(b, uint8(nBits8))
	r.Ymin = Twips(t)
	if err == nil || err == io.EOF {
		t.ReadBitsFrom(b, uint8(nBits8))
		r.Ymax = Twips(t)
	}
	return
}

func (r *Rect) WriteTo(f io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	b := &bitWriter{Writer: c}
	defer b.Align()
	biXmin := BitInt(r.Xmin)
	biXmax := BitInt(r.Xmax)
	biYmin := BitInt(r.Ymin)
	biYmax := BitInt(r.Ymax)
	nBits := uint8(max(biXmin.Size(), biXmax.Size(), biYmin.Size(), biYmax.Size()))
	bits := BitUint(nBits)
	if err = bits.WriteBitsTo(b, 5); err != nil {
		return
	}
	if err = biXmin.WriteBitsTo(b, nBits); err != nil {
		return
	}
	if err = biXmax.WriteBitsTo(b, nBits); err != nil {
		return
	}
	if err = biYmin.WriteBitsTo(b, nBits); err != nil {
		return
	}
	err = biYmax.WriteBitsTo(b, nBits)
	return
}

func (r *Rect) Size() int32 {
	biXmin := BitInt(r.Xmin)
	biXmax := BitInt(r.Xmax)
	biYmin := BitInt(r.Ymin)
	biYmax := BitInt(r.Ymax)
	total := (5 + 4*max(biXmin.Size(), biXmax.Size(), biYmin.Size(), biYmax.Size()))
	if total%8 == 0 {
		return total / 8
	}
	return 1 + total/8
}

func (r *Rect) String() string {
	return fmt.Sprintf("(%s, %s), (%s, %s), (%s, %s), (%s, %s)", r.Xmin.String(), r.Ymin.String(), r.Xmax.String(), r.Ymin.String(), r.Xmax.String(), r.Ymax.String(), r.Xmin.String(), r.Ymin.String())
}

type Matrix struct {
	ScaleX, ScaleY, RotateSkew0, RotateSkew1 float32
	TranslateX, TranslateY                   Twips
}

func NewMatrix(ScaleX, ScaleY, RotateSkew0, RotateSkew1 float32, TranslateX, TranslateY int32) *Matrix {
	return &Matrix{ScaleX, ScaleY, RotateSkew0, RotateSkew1, Twips(TranslateX), Twips(TranslateY)}
}

func (m *Matrix) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	var (
		d BitUint
		bf BitFixed
	)
	b := &bitReader{Reader: c}
	if err = d.ReadBitsFrom(b, 1); err != nil {
		return
	}
	if d == 1 {
		if err = d.ReadBitsFrom(b, 5); err != nil {
			return
		}
		if err = bf.ReadBitsFrom(b, uint8(d)); err != nil {
			return
		}
		m.ScaleX = float32(bf)
		if err = bf.ReadBitsFrom(b, uint8(d)); err != nil {
			return
		}
		m.ScaleY = float32(bf)
	} else {
		m.ScaleX = 1
		m.ScaleY = 1
	}
	if err = d.ReadBitsFrom(b, 1); err != nil {
		return
	}
	if d == 1 {
		if err = d.ReadBitsFrom(b, 5); err != nil {
			return
		}
		if err = bf.ReadBitsFrom(b, uint8(d)); err != nil {
			return
		}
		m.RotateSkew0 = float32(bf)
		if err = bf.ReadBitsFrom(b, uint8(d)); err != nil {
			return
		}
		m.RotateSkew1 = float32(bf)
	} else {
		m.RotateSkew0 = 1
		m.RotateSkew1 = 1
	}
	if err = d.ReadBitsFrom(b, 1); err != nil {
		return
	}
	if d == 1 {
		if err = d.ReadBitsFrom(b, 5); err != nil {
			return
		}
		var sB BitInt
		if err = sB.ReadBitsFrom(b, uint8(d)); err != nil {
			return
		}
		m.TranslateX = Twips(sB)
		if err = sB.ReadBitsFrom(b, uint8(d)); err != nil {
			return
		}
		m.TranslateY = Twips(sB)
	}
	return
}

func (m *Matrix) WriteTo(f io.Writer) (total int64, err error) {
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	b := &bitWriter{Writer: c}
	defer b.Align()
	var size BitUint
	zero := BitUint(0)
	one := BitUint(1)
	if m.ScaleX != 1 || m.ScaleY != 1 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		scaleX, scaleY := BitFixed(m.ScaleX), BitFixed(m.ScaleY)
		size = BitUint(max(scaleX.Size(), scaleY.Size()))
		if err = size.WriteBitsTo(b, 5); err != nil {
			return
		}
		
		if err = scaleX.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = scaleY.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	if m.RotateSkew0 != 1 || m.RotateSkew1 != 1 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		rotateSkew0, rotateSkew1 := BitFixed(m.RotateSkew0), BitFixed(m.RotateSkew1)
		size = BitUint(max(rotateSkew0.Size(), rotateSkew1.Size()))
		if err = size.WriteBitsTo(b, 5); err != nil {
			return
		}
		if err = rotateSkew0.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = rotateSkew1.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	if m.TranslateX != 0 || m.TranslateY != 0 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		x, y := BitInt(m.TranslateX), BitInt(m.TranslateY)
		size = BitUint(max(x.Size(), y.Size()))
		if err = size.WriteBitsTo(b, 5); err != nil {
			return
		}
		if err = x.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = y.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	return
}

func (m *Matrix) Size() int32 {
	total := int32(0)
	if m.ScaleX != 1 || m.ScaleY != 1 {
		scaleX, scaleY := BitFixed(m.ScaleX), BitFixed(m.ScaleY)
		total += 1 + 5 + 2*max(scaleX.Size(), scaleY.Size())
	}
	if m.RotateSkew0 != 0 || m.RotateSkew1 != 0 {
		rotateSkew0, rotateSkew1 := BitFixed(m.RotateSkew0), BitFixed(m.RotateSkew1)
		total += 1 + 5 + 2*max(rotateSkew0.Size(), rotateSkew1.Size())
	}
	if m.TranslateX != 0 || m.TranslateY != 0 {
		x, y := BitInt(m.TranslateX), BitInt(m.TranslateY)
		total += 1 + 5 + 2*max(x.Size(), y.Size())
	}
	if total%8 == 0 {
		return total / 8
	}
	return 1 + total/8
}

func (m *Matrix) String() string {
	return fmt.Sprintf("MATRIX: [ [ %f, %f ], [ %f, %f ], [ %s, %s ] ]", m.ScaleX, m.RotateSkew0, m.RotateSkew1, m.ScaleY, m.TranslateX.String(), m.TranslateY.String())
}

type CXForm struct {
	RedMultTerm, GreenMultTerm, BlueMultTerm, RedAddTerm, GreenAddTerm, BlueAddTerm int16
}

func NewCXForm(RedMultTerm, GreenMultTerm, BlueMultTerm, RedAddTerm, GreenAddTerm, BlueAddTerm int16) *CXForm {
	return &CXForm{RedMultTerm, GreenMultTerm, BlueMultTerm, RedAddTerm, GreenAddTerm, BlueAddTerm}
}

func (c *CXForm) ReadFrom(f io.Reader) (total int64, err error) {
	cr := &rwcount.CountReader{Reader: f}
	defer func() { total = cr.BytesRead() }()
	var a, m, n BitUint
	b := &bitReader{Reader: cr}
	if err = m.ReadBitsFrom(b, 1); err == nil {
		if err = a.ReadBitsFrom(b, 1); err == nil {
			err = n.ReadBitsFrom(b, 4)
		}
	}
	if err != nil {
		return
	}
	bits := uint8(n)
	if m == 1 {
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.RedMultTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.GreenMultTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil && !(a == 0 && err == io.EOF) {
			return
		}
		c.BlueMultTerm = int16(n)
	} else {
		c.RedMultTerm, c.GreenMultTerm, c.BlueMultTerm = 256, 256, 256
	}
	if a == 1 {
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.RedAddTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.GreenAddTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.BlueAddTerm = int16(n)
	}
	return
}

func (c *CXForm) WriteTo(f io.Writer) (total int64, err error) {
	cw := &rwcount.CountWriter{Writer: f}
	defer func() { total = cw.BytesWritten() }()
	b := &bitWriter{Writer: cw}
	defer b.Align()
	var (
		hasAdd, hasMult bool
		size            int32
	)
	one := BitUint(1)
	zero := BitUint(0)
	ra := BitInt(c.RedAddTerm)
	ga := BitInt(c.GreenAddTerm)
	ba := BitInt(c.BlueAddTerm)
	rm := BitInt(c.RedMultTerm)
	gm := BitInt(c.GreenMultTerm)
	bm := BitInt(c.BlueMultTerm)
	if ra != 0 || ga != 0 || ba != 0 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		hasAdd = true
		size = max(ra.Size(), ga.Size(), ba.Size())
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	if rm != 256 || gm != 256 || bm != 256 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		hasMult = true
		size = max(size, rm.Size(), gm.Size(), bm.Size())
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	bSize := BitUint(size)
	if err = bSize.WriteBitsTo(b, 4); err != nil {
		return
	}
	if hasMult {
		if err = rm.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = gm.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = bm.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	}
	if hasAdd {
		if err = ra.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = ga.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = ba.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	}
	return
}

func (c *CXForm) Size() int32 {
	var total, size, mult int32
	total = 6
	ra := BitInt(c.RedAddTerm)
	ga := BitInt(c.GreenAddTerm)
	ba := BitInt(c.BlueAddTerm)
	rm := BitInt(c.RedMultTerm)
	gm := BitInt(c.GreenMultTerm)
	bm := BitInt(c.BlueMultTerm)
	if ra != 0 || ga != 0 || ba != 0 {
		mult = 3
		size = max(ra.Size(), ga.Size(), ba.Size())
	}
	if rm != 256 || gm != 256 || bm != 256 {
		mult += 3
		size = max(size, rm.Size(), gm.Size(), bm.Size())
	}
	total += mult * size
	if total%8 == 0 {
		return total / 8
	}
	return 1 + total/8
}

func (c *CXForm) String() string {
	s := "R"
	if c.RedMultTerm != 256 {
		s += fmt.Sprintf(" * %d/256", c.RedMultTerm)
	}
	if c.RedAddTerm != 0 {
		s += fmt.Sprintf(" + %d", c.RedAddTerm)
	}
	s += ", G"
	if c.GreenMultTerm != 256 {
		s += fmt.Sprintf(" * %d/256", c.GreenMultTerm)
	}
	if c.GreenAddTerm != 0 {
		s += fmt.Sprintf(" + %d", c.GreenAddTerm)
	}
	s += ", B"
	if c.BlueMultTerm != 256 {
		s += fmt.Sprintf(" * %d/256", c.BlueMultTerm)
	}
	if c.BlueAddTerm != 0 {
		s += fmt.Sprintf(" + %d", c.BlueAddTerm)
	}
	return s
}

type CXFormWithAlpha struct {
	CXForm
	AlphaMultTerm, AlphaAddTerm int16
}

func NewCXFormWithAlpha(RedMultTerm, GreenMultTerm, BlueMultTerm, AlphaMultTerm, RedAddTerm, GreenAddTerm, BlueAddTerm, AlphaAddTerm int16) *CXFormWithAlpha {
	return &CXFormWithAlpha{CXForm{RedMultTerm, GreenMultTerm, BlueMultTerm, RedAddTerm, GreenAddTerm, BlueAddTerm}, AlphaMultTerm, AlphaAddTerm}
}

func (c *CXFormWithAlpha) ReadFrom(f io.Reader) (total int64, err error) {
	cr := &rwcount.CountReader{Reader: f}
	defer func() { total = cr.BytesRead() }()
	var a, m, n BitUint
	b := &bitReader{Reader: cr}
	if err = m.ReadBitsFrom(b, 1); err == nil {
		err = a.ReadBitsFrom(b, 1)
		if err == nil {
			err = n.ReadBitsFrom(b, 4)
		}
	}
	if err != nil {
		return
	}
	bits := uint8(n)
	if m == 1 {
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.RedMultTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.GreenMultTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.BlueMultTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil && !(a == 0 && err == io.EOF) {
			return
		}
		c.AlphaMultTerm = int16(n)
	} else {
		c.RedMultTerm, c.GreenMultTerm, c.BlueMultTerm, c.AlphaMultTerm = 256, 256, 256, 256
	}
	if a == 1 {
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.RedAddTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.GreenAddTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.BlueAddTerm = int16(n)
		if err = n.ReadBitsFrom(b, bits); err != nil {
			return
		}
		c.AlphaAddTerm = int16(n)
	}
	return
}

func (c *CXFormWithAlpha) WriteTo(f io.Writer) (total int64, err error) {
	cw := &rwcount.CountWriter{Writer: f}
	defer func() { total = cw.BytesWritten() }()
	b := &bitWriter{Writer: cw}
	defer b.Align()
	var (
		hasAdd, hasMult bool
		size            int32
	)
	one := BitUint(1)
	zero := BitUint(0)
	ra := BitInt(c.RedAddTerm)
	ga := BitInt(c.GreenAddTerm)
	ba := BitInt(c.BlueAddTerm)
	aa := BitInt(c.AlphaAddTerm)
	rm := BitInt(c.RedMultTerm)
	gm := BitInt(c.GreenMultTerm)
	bm := BitInt(c.BlueMultTerm)
	am := BitInt(c.AlphaMultTerm)
	if ra != 0 || ga != 0 || ba != 0 || aa != 0 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		hasAdd = true
		size = max(ra.Size(), ga.Size(), ba.Size(), aa.Size())
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	if rm != 256 || ga != 256 || bm != 256 || am != 256 {
		if err = one.WriteBitsTo(b, 1); err != nil {
			return
		}
		hasMult = true
		size = max(size, rm.Size(), gm.Size(), bm.Size(), am.Size())
	} else if err = zero.WriteBitsTo(b, 1); err != nil {
		return
	}
	bSize := BitUint(size)
	if err = bSize.WriteBitsTo(b, 4); err != nil {
		return
	}
	if hasMult {
		if err = rm.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = gm.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = bm.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = am.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	}
	if hasAdd {
		if err = ra.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = ga.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = ba.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
		if err = aa.WriteBitsTo(b, uint8(size)); err != nil {
			return
		}
	}
	return
}

func (c *CXFormWithAlpha) Size() int32 {
	var total, size, mult int32
	total = 6
	ra := BitInt(c.RedAddTerm)
	ga := BitInt(c.GreenAddTerm)
	ba := BitInt(c.BlueAddTerm)
	aa := BitInt(c.AlphaAddTerm)
	rm := BitInt(c.RedMultTerm)
	gm := BitInt(c.GreenMultTerm)
	bm := BitInt(c.BlueMultTerm)
	am := BitInt(c.AlphaMultTerm)
	if ra != 0 || ga != 0 || ba != 0 || aa != 0 {
		mult = 4
		size = max(ra.Size(), ga.Size(), ba.Size(), aa.Size())
	}
	if rm != 256 || gm != 256 || bm != 256 || am != 256 {
		mult += 4
		size = max(size, rm.Size(), gm.Size(), bm.Size(), am.Size())
	}
	total += mult * size
	if total%8 == 0 {
		return total / 8
	}
	return 1 + total/8
}

func (c *CXFormWithAlpha) String() string {
	s := c.CXForm.String()
	s += ", A"
	if c.AlphaMultTerm != 256 {
		s += fmt.Sprintf(" * %d/256", c.AlphaMultTerm)
	}
	if c.AlphaAddTerm != 0 {
		s += fmt.Sprintf(" + %d", c.AlphaAddTerm)
	}
	return s
}

func ReadAll(f io.Reader, fs []io.ReaderFrom) (err error) {
	for _, r := range fs {
		if _, err = r.ReadFrom(f); err != nil {
			return
		}
	}
	return
}

func WriteAll(w io.Writer, fs []io.WriterTo) (err error) {
	for _, r := range fs {
		if _, err = r.WriteTo(w); err != nil {
			return
		}
	}
	return
}

func SizeAll(ss []Sizer) (total int32) {
	for _, s := range ss {
		total += s.Size()
	}
	return
}

func min(sizes ...int32) int32 {
	if len(sizes) == 0 {
		return 0
	}
	min := sizes[0]
	for _, i := range sizes[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

func max(sizes ...int32) int32 {
	if len(sizes) == 0 {
		return 0
	}
	max := sizes[0]
	for _, i := range sizes[1:] {
		if i > max {
			max = i
		}
	}
	return max
}
