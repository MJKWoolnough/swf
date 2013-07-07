package swf

import (
	"bytes"
	"github.com/MJKWoolnough/equaler"
	"io"
	"math"
	"testing"
)

type Tester interface {
	io.ReaderFrom
	io.WriterTo
	equaler.Equaler
}

type TesterBits interface {
	ReadBitsFrom(BitReader, uint8) error
	WriteBitsTo(BitWriter, uint8) error
	equaler.Equaler
}

type TesterSize interface {
	Size() int32
}

type bitsTest struct {
	data equaler.Equaler
	bits uint8
}

type sizeTest struct {
	Sizer
	size int32
}

func test(t *testing.T, target Tester, data []byte, units []equaler.Equaler) {
	bufFrom := bytes.NewBuffer(data)
	bufTo := new(bytes.Buffer)
	var (
		err    error
		br, bw int64
	)
	for n, test := range units {
		if br, err = target.ReadFrom(bufFrom); err != nil {
			t.Errorf("test %d: %q", n+1, err)
		} else if !target.Equal(test) {
			t.Errorf("test %d: expecting %s, got %s", n+1, test, target)
		} else if bw, err = target.WriteTo(bufTo); err != nil {
			t.Errorf("test %d: %q", n+1, err)
		} else if br != bw {
			t.Errorf("test %d: read %d bytes, but wrote %d bytes", n+1, br, bw)
		} else {
			continue
		}
		break
	}
	if bytes.Compare(bufTo.Bytes(), data) != 0 {
		t.Errorf("expecting %v, got %v", data, bufTo.Bytes())
	}
}

func testBits(t *testing.T, target TesterBits, data []byte, units []bitsTest) {
	bufFrom := &bitReader{Reader: bytes.NewBuffer(data)}
	buf := new(bytes.Buffer)
	bufTo := &bitWriter{Writer: buf}
	var err error
	for n, test := range units {
		if err = target.ReadBitsFrom(bufFrom, test.bits); err != nil {
			t.Errorf("test %d: %q", n+1, err)
			break
		}
		if !target.Equal(test.data) {
			t.Errorf("test %d: expecting %s, got %s", n+1, test.data, target)
			break
		}
		if err = target.WriteBitsTo(bufTo, test.bits); err != nil {
			t.Errorf("test %d: %q", n+1, err)
			break
		}
	}
	bufTo.Align()
	if bytes.Compare(buf.Bytes(), data) != 0 {
		t.Errorf("expecting %v, got %v", data, buf.Bytes())
	}
}

func testSize(t *testing.T, tests []sizeTest) {
	for n, test := range tests {
		if x := test.Sizer.Size(); x != test.size {
			t.Errorf("test %d: size mismatch, got %d, expected %d", n+1, x, test.size)
		}
	}
}

func TestUint8(t *testing.T) {
	test(t, new(Uint8), []byte{0, 255, 127}, []equaler.Equaler{
		NewUint8(0),
		NewUint8(255),
		NewUint8(127),
	})
}

func TestInt8(t *testing.T) {
	test(t, new(Int8), []byte{0, 255, 127}, []equaler.Equaler{
		NewInt8(0),
		NewInt8(-1),
		NewInt8(127),
	})
}

func TestUint16(t *testing.T) {
	test(t, new(Uint16), []byte{210, 76, 91, 44, 253, 202, 3, 240, 205, 55}, []equaler.Equaler{
		NewUint16(19666),
		NewUint16(11355),
		NewUint16(51965),
		NewUint16(61443),
		NewUint16(14285),
	})
}

func TestInt16(t *testing.T) {
	test(t, new(Int16), []byte{210, 76, 91, 44, 253, 202, 3, 240, 205, 55}, []equaler.Equaler{
		NewInt16(19666),
		NewInt16(11355),
		NewInt16(-13571),
		NewInt16(-4093),
		NewInt16(14285),
	})
}

func TestUint32(t *testing.T) {
	test(t, new(Uint32), []byte{121, 66, 185, 86, 37, 70, 199, 153, 52, 236, 146, 89, 201, 117, 242, 97, 202, 189, 10, 213}, []equaler.Equaler{
		NewUint32(1454981753),
		NewUint32(2579973669),
		NewUint32(1502800948),
		NewUint32(1643279817),
		NewUint32(3574250954),
	})
}

func TestInt32(t *testing.T) {
	test(t, new(Int32), []byte{121, 66, 185, 86, 37, 70, 199, 153, 52, 236, 146, 89, 201, 117, 242, 97, 202, 189, 10, 213}, []equaler.Equaler{
		NewInt32(1454981753),
		NewInt32(-1714993627),
		NewInt32(1502800948),
		NewInt32(1643279817),
		NewInt32(-720716342),
	})
}

func TestUint64(t *testing.T) {
	test(t, new(Uint64), []byte{159, 170, 253, 129, 194, 127, 153, 19, 55, 3, 89, 117, 131, 10, 191, 182, 49, 19, 154, 1, 58, 55, 126, 92, 2, 203, 233, 210, 230, 24, 140, 44, 192, 44, 176, 179, 13, 246, 218, 132}, []equaler.Equaler{
		NewUint64(1412300431538629279),
		NewUint64(13168255395180380983),
		NewUint64(6664825220829418289),
		NewUint64(3209968014068402946),
		NewUint64(9573234496639085760),
	})
}

func TestInt64(t *testing.T) {
	test(t, new(Int64), []byte{159, 170, 253, 129, 194, 127, 153, 19, 55, 3, 89, 117, 131, 10, 191, 182, 49, 19, 154, 1, 58, 55, 126, 92, 2, 203, 233, 210, 230, 24, 140, 44, 192, 44, 176, 179, 13, 246, 218, 132}, []equaler.Equaler{
		NewInt64(1412300431538629279),
		NewInt64(-5278488678529170633),
		NewInt64(6664825220829418289),
		NewInt64(3209968014068402946),
		NewInt64(-8873509577070465856),
	})
}

func TestFloat16(t *testing.T) {
	test(t, new(Float16), []byte{0, 60, 1, 60, 255, 123, 0, 128, 0, 252}, []equaler.Equaler{
		NewFloat16(1),
		NewFloat16(1.0009765625),
		NewFloat16(65504),
		NewFloat16(-0),
		NewFloat16(float32(math.Inf(-1))),
	})
}

func TestFloat(t *testing.T) {
	test(t, new(Float), []byte{0, 0, 68, 65, 250, 238, 231, 195, 255, 255, 127, 127, 1, 0, 0, 0, 0, 0, 128, 255}, []equaler.Equaler{
		NewFloat(12.25),
		NewFloat(-463.867),
		NewFloat(math.MaxFloat32),
		NewFloat(math.SmallestNonzeroFloat32),
		NewFloat(float32(math.Inf(-1))),
	})
}

func TestDouble(t *testing.T) {
	test(t, new(Double), []byte{0, 0, 0, 0, 0, 0, 0, 192, 255, 255, 255, 255, 255, 255, 239, 127, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 240, 255, 0, 0, 0, 0, 0, 0, 0, 128}, []equaler.Equaler{
		NewDouble(-2),
		NewDouble(math.MaxFloat64),
		NewDouble(math.SmallestNonzeroFloat64),
		NewDouble(math.Inf(-1)),
		NewDouble(-0),
	})
}

func TestTwips(t *testing.T) {
	test(t, new(Twips), []byte{121, 66, 185, 86, 37, 70, 199, 153, 52, 236, 146, 89, 201, 117, 242, 97, 202, 189, 10, 213}, []equaler.Equaler{
		NewTwips(1454981753),
		NewTwips(-1714993627),
		NewTwips(1502800948),
		NewTwips(1643279817),
		NewTwips(-720716342),
	})
}

func TestFixed(t *testing.T) {
	test(t, new(Fixed), []byte{0, 0, 0, 0, 255, 255, 255, 255, 1, 0, 0, 0}, []equaler.Equaler{
		NewFixed(0),
		NewFixed(65535.99998474121),
		NewFixed(0.0000152587890625),
	})
}

func TestFixed8(t *testing.T) {
	test(t, new(Fixed8), []byte{1, 0, 255, 255, 0, 0}, []equaler.Equaler{
		NewFixed8(0.00390625),
		NewFixed8(255.99609375),
		NewFixed8(0),
	})
}

func TestEncodedU32(t *testing.T) {
	test(t, new(EncodedU32), []byte{0, 127, 255, 1, 255, 255, 255, 255, 15}, []equaler.Equaler{
		NewEncodedU32(0),
		NewEncodedU32(127),
		NewEncodedU32(255),
		NewEncodedU32(4294967295),
	})
}

func TestEncodedU32Size(t *testing.T) {
	testSize(t, []sizeTest{
		{NewEncodedU32(1), 1},
		{NewEncodedU32(1024), 2},
		{NewEncodedU32(65536), 3},
	})
}

func TestBitUint(t *testing.T) {
	testBits(t, new(BitUint), []byte{67, 191, 240}, []bitsTest{
		{NewBitUint(0), 1},
		{NewBitUint(1), 1},
		{NewBitUint(1), 5},
		{NewBitUint(3), 2},
		{NewBitUint(1), 2},
		{NewBitUint(255), 8},
		{NewBitUint(16), 5},
	})
}

func TestBitUintSize(t *testing.T) {
	testSize(t, []sizeTest{
		{NewBitUint(1), 1},
		{NewBitUint(1024), 11},
		{NewBitUint(65536), 17},
	})
}

func TestBitInt(t *testing.T) {
	testBits(t, new(BitInt), []byte{67, 191, 240}, []bitsTest{
		{NewBitInt(0), 1},
		{NewBitInt(-1), 1},
		{NewBitInt(1), 5},
		{NewBitInt(-1), 2},
		{NewBitInt(1), 2},
		{NewBitInt(-1), 8},
		{NewBitInt(-16), 5},
	})
}

func TestBitIntSize(t *testing.T) {
	testSize(t, []sizeTest{
		{NewBitInt(1), 2},
		{NewBitInt(1024), 12},
		{NewBitInt(65536), 18},
		{NewBitInt(0), 1},
		{NewBitInt(-1), 1},
		{NewBitInt(-14), 5},
	})
}

func TestBitFixed(t *testing.T) {
	testBits(t, new(BitFixed), []byte{64, 0, 16, 0, 12, 20, 0}, []bitsTest{
		{NewBitFixed(0), 1},
		{NewBitFixed(-1), 17},
		{NewBitFixed(0.5), 17},
		{NewBitFixed(3.01953125), 19},
	})
}

func TestBitFixedSize(t *testing.T) {
	testSize(t, []sizeTest{
		{NewBitFixed(1), 18},
		{NewBitFixed(1024), 28},
		{NewBitFixed(30000), 32},
	})
}

func TestString(t *testing.T) {
	test(t, new(String), []byte{104, 101, 108, 108, 111, 44, 32, 228, 184, 150, 231, 149, 140, 0, 72, 111, 87, 32, 84, 104, 69, 32, 87, 101, 66, 32, 73, 115, 32, 119, 79, 118, 69, 110, 33, 63, 0}, []equaler.Equaler{
		NewString("hello, 世界"),
		NewString("HoW ThE WeB Is wOvEn!?"),
	})
}

func TestLanguageCode(t *testing.T) {
	test(t, new(LanguageCode), []byte{3, 4, 5, 1, 2}, []equaler.Equaler{
		NewLanguageCode(uint8(LANGUAGE_KOREAN)),
		NewLanguageCode(uint8(LANGUAGE_SIMPLIFIED_CHINESE)),
		NewLanguageCode(uint8(LANGUAGE_TRADITIONAL_CHINESE)),
		NewLanguageCode(uint8(LANGUAGE_LATIN)),
		NewLanguageCode(uint8(LANGUAGE_JAPENESE)),
	})
}

func TestRGB(t *testing.T) {
	test(t, new(RGB), []byte{153, 107, 154, 139, 30, 219, 247, 213, 239, 232, 255, 173}, []equaler.Equaler{
		NewRGB(153, 107, 154),
		NewRGB(139, 30, 219),
		NewRGB(247, 213, 239),
		NewRGB(232, 255, 173),
	})
}

func TestRGBA(t *testing.T) {
	test(t, new(RGBA), []byte{153, 107, 154, 139, 30, 219, 247, 213, 239, 232, 255, 173}, []equaler.Equaler{
		NewRGBA(153, 107, 154, 139),
		NewRGBA(30, 219, 247, 213),
		NewRGBA(239, 232, 255, 173),
	})
}

func TestARGB(t *testing.T) {
	test(t, new(ARGB), []byte{153, 107, 154, 139, 30, 219, 247, 213, 239, 232, 255, 173}, []equaler.Equaler{
		NewARGB(153, 107, 154, 139),
		NewARGB(30, 219, 247, 213),
		NewARGB(239, 232, 255, 173),
	})
}

func TestRect(t *testing.T) {
	test(t, new(Rect), []byte{24, 41, 128, 75, 169, 126, 32, 54, 0}, []equaler.Equaler{
		NewRect(0, 1, 2, 3),
		NewRect(234, 191, 32, 108),
	})
}

func TestRectSize(t *testing.T) {
	testSize(t, []sizeTest{
		{NewRect(0, 1, 2, 3), 3},
		{NewRect(234, 191, 32, 108), 6},
	})
}

func TestMatrix(t *testing.T) {
	test(t, new(Matrix), []byte{205, 0, 0, 8, 0, 12, 194, 0, 3, 0, 0, 165, 145, 216, 217, 52, 0, 1, 32, 0, 49, 64, 0, 8, 0, 37, 103, 0, 236, 109, 128, 0, 32, 8, 0, 15, 131, 252, 128, 0, 127, 248, 128, 2, 224, 31, 47, 160}, []equaler.Equaler{
		NewMatrix(2, 0.5, 0.25, 3, 200, -40),
		NewMatrix(19.25, 4.5, 0.5, 0.125, 12, -4),
		NewMatrix(219, 512.5, 1020.5, 8190.125, 124, -4192),
	})
}

func TestMatrixSize(t *testing.T) {
	testSize(t, []sizeTest{
		{NewMatrix(2, 0.5, 0.25, 3, 200, -40), 14},
		{NewMatrix(19.25, 4.5, 0.5, 0.125, 12, -4), 14},
		{NewMatrix(219, 512.5, 1020.5, 8190.125, 124, -4192), 20},
	})
}

func TestCXForm(t *testing.T) {
	test(t, new(CXForm), []byte{229, 56, 247, 106, 177, 73, 230, 192, 229, 134, 173, 77, 21, 83, 238, 96}, []equaler.Equaler{
		NewCXForm(156, 247, 213, 197, 79, 108),
		NewCXForm(195, 173, 154, 85, 159, 230),
	})
}

func TestCXFormSize(t *testing.T) {
	testSize(t, []sizeTest{
		{ NewCXForm(156, 247, 213, 197, 79, 108), 8 },
		{ NewCXForm(195, 173, 154, 85, 159, 230), 8 },
		{ NewCXForm(1, 3, 14, 2, 9, 30), 6 },
	})
}

func TestCXFormWithAlpha(t *testing.T) {
	test(t, new(CXFormWithAlpha), []byte{229, 56, 247, 106, 160, 24, 164, 243, 98, 104, 229, 134, 173, 77, 53, 202, 169, 247, 48, 56}, []equaler.Equaler{
		NewCXFormWithAlpha(156, 247, 213, 128, 197, 79, 108, 154),
		NewCXFormWithAlpha(195, 173, 154, 215, 85, 159, 230, 14),
	})
}

func TestCXFormWithAlphaSize(t *testing.T) {
	testSize(t, []sizeTest{
		{ NewCXFormWithAlpha(156, 247, 213, 128, 197, 79, 108, 154), 10 },
		{ NewCXFormWithAlpha(195, 173, 154, 215, 85, 159, 230, 14), 10 },
		{ NewCXFormWithAlpha(1, 3, 14, 5, 2, 9, 30, 14), 7 },
	})
}

func (u *Uint8) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Uint8); ok && *eu == *u {
		return true
	}
	return false
}

func (u *Uint16) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Uint16); ok && *eu == *u {
		return true
	}
	return false
}

func (u *Uint32) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Uint32); ok && *eu == *u {
		return true
	}
	return false
}

func (u *Uint64) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Uint64); ok && *eu == *u {
		return true
	}
	return false
}

func (i *Int8) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Int8); ok && *eu == *i {
		return true
	}
	return false
}

func (i *Int16) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Int16); ok && *eu == *i {
		return true
	}
	return false
}

func (i *Int32) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Int32); ok && *eu == *i {
		return true
	}
	return false
}

func (i *Int64) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Int64); ok && *eu == *i {
		return true
	}
	return false
}

func (f *Float16) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Float16); ok && *eu == *f {
		return true
	}
	return false
}

func (f *Float) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Float); ok && *eu == *f {
		return true
	}
	return false
}

func (d *Double) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Double); ok && *eu == *d {
		return true
	}
	return false
}

func (t *Twips) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Twips); ok && *eu == *t {
		return true
	}
	return false
}

func (f *Fixed) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Fixed); ok && *eu == *f {
		return true
	}
	return false
}

func (f *Fixed8) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Fixed8); ok && *eu == *f {
		return true
	}
	return false
}

func (en *EncodedU32) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*EncodedU32); ok && *eu == *en {
		return true
	}
	return false
}

func (b *BitUint) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*BitUint); ok && *eu == *b {
		return true
	}
	return false
}

func (b *BitInt) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*BitInt); ok && *eu == *b {
		return true
	}
	return false
}

func (b *BitFixed) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*BitFixed); ok && *eu == *b {
		return true
	}
	return false
}

func (s *String) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*String); ok && *eu == *s {
		return true
	}
	return false
}

func (l *LanguageCode) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*LanguageCode); ok && *eu == *l {
		return true
	}
	return false
}

func (r *RGB) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*RGB); ok {
		return eu.Red == r.Red && eu.Green == r.Green && eu.Blue == r.Blue
	}
	return false
}

func (r *RGBA) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*RGBA); ok && *eu == *r {
		return eu.Red == r.Red && eu.Green == r.Green && eu.Blue == r.Blue && eu.Alpha == r.Alpha
	}
	return false
}

func (a *ARGB) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*ARGB); ok && *eu == *a {
		return eu.Alpha == a.Alpha && eu.Red == a.Red && eu.Green == a.Green && eu.Blue == a.Blue
	}
	return false
}

func (r *Rect) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Rect); ok && *eu == *r {
		return eu.Xmin == r.Xmin && eu.Xmax == r.Xmax && eu.Ymin == r.Ymin && eu.Ymax == r.Ymax
	}
	return false
}

func (m *Matrix) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*Matrix); ok && *eu == *m {
		return eu.ScaleX == m.ScaleX && eu.ScaleY == m.ScaleY && eu.RotateSkew0 == m.RotateSkew0 && eu.RotateSkew1 == m.RotateSkew1 && eu.TranslateX == m.TranslateX && eu.TranslateY == m.TranslateY
	}
	return false
}

func (c *CXForm) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*CXForm); ok && *eu == *c {
		return eu.RedMultTerm == c.RedMultTerm && eu.GreenMultTerm == c.GreenMultTerm && eu.BlueMultTerm == c.BlueMultTerm && eu.RedAddTerm == c.RedAddTerm && eu.GreenAddTerm == c.GreenAddTerm && eu.BlueAddTerm == c.BlueAddTerm
	}
	return false
}

func (c *CXFormWithAlpha) Equal(e equaler.Equaler) bool {
	if eu, ok := e.(*CXFormWithAlpha); ok && *eu == *c {
		return eu.RedMultTerm == c.RedMultTerm && eu.GreenMultTerm == c.GreenMultTerm && eu.BlueMultTerm == c.BlueMultTerm && eu.AlphaMultTerm == c.AlphaMultTerm && eu.RedAddTerm == c.RedAddTerm && eu.GreenAddTerm == c.GreenAddTerm && eu.BlueAddTerm == c.BlueAddTerm && eu.AlphaAddTerm == c.AlphaAddTerm
	}
	return false
}
