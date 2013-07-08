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

// A SWF decoder.
package swf

//http://www.m2osw.com/swf_alexref
//http://www.the-labs.com/MacromediaFlash/SWF-Spec/SWFfilereference.html

import (
	"code.google.com/p/lzma"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/MJKWoolnough/rwcount"
	"io"
	"io/ioutil"
)

type Error struct {
	Tag string
	Err error
}

func wrapError(tag string, err *error) error {
	if *err != nil {
		return &Error{tag, *err}
	}
	return nil
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Tag, e.Err)
}

type ErrMinVersion struct {
	Tag string
	Ver uint8
}

func (e ErrMinVersion) Error() string {
	return fmt.Sprintf("tag %q requires a SWF file of at least version %d.", e.Tag, e.Ver)
}

type InvalidTagCode struct {
	TagCode uint16
}

func (i InvalidTagCode) Error() string {
	return fmt.Sprintf("invalid/unknown tag code discovered: %d", i.TagCode)
}

type ParserError struct {
	Tag, Field, Found string
}

func (p ParserError) Error() string {
	return fmt.Sprintf("unable to parse tag %q, field %q - found %q", p.Tag, p.Field, p.Found)
}

type BadHeader struct {
	Code uint8
	Err  error
}

func (b BadHeader) Error() string {
	switch b.Code {
	case 0:
		fmt.Sprintf("error while reading header: %q", b.Err)
	case 1:
		fmt.Sprintf("invalid signature: %q", b.Err)
	case 2:
		fmt.Sprintf("invalid compression type")
	}
	return "unknown error"
}

type compression uint8

func (c compression) String() string {
	switch c {
	case COMPRESS_NONE:
		return "No compression"
	case COMPRESS_ZLIB:
		return "ZLIB compression"
	case COMPRESS_LZMA:
		return "LZMA compression"
	}
	return "Unknown compression type"
}

const MAX_VER uint8 = 11

const (
	COMPRESS_NONE compression = iota
	COMPRESS_ZLIB
	COMPRESS_LZMA
)

func TagFromId(id uint16) Tag {
	return TagFromIdVer(id, MAX_VER)
}

func TagFromIdVer(id uint16, ver uint8) Tag {
	return nil
}

type SWF struct {
	Compressed compression
	Version    uint8
	FrameSize  Rect
	FrameRate  uint16
	FrameCount uint16
	Tags       []Tag
}

func (s SWF) String() string {
	return fmt.Sprintf("SWF Version: %d\nFrame Size: %dx%d\nFrame Rate: %d\nFrame Count\nCompression: %s", s.Version, s.FrameSize.Xmax, s.FrameSize.Ymax, s.FrameRate, s.FrameCount, s.Compressed)
}

func (s *SWF) ReadFrom(f io.Reader) (total int64, err error) {
	c := &rwcount.CountReader{Reader: f}
	defer func() { total = c.BytesRead() }()
	f = c
	var (
		signature  [3]byte
		fileLength int32
	)
	s.Tags = make([]Tag, 0)
	if err = binary.Read(f, binary.LittleEndian, &signature); err != nil {
		err = &BadHeader{0, err}
		return
	}
	switch signature[0] {
	case 'F':
		s.Compressed = COMPRESS_NONE
	case 'C':
		s.Compressed = COMPRESS_ZLIB
	case 'Z':
		s.Compressed = COMPRESS_LZMA
	default:
		err = &BadHeader{0, errors.New(string(signature[:]))}
		return
	}
	if signature[1] != 'W' || signature[2] != 'S' {
		err = &BadHeader{1, errors.New(string(signature[:]))}
		return
	}
	if err = binary.Read(f, binary.LittleEndian, &s.Version); err != nil {
		err = &BadHeader{0, err}
		return
	}
	if err = binary.Read(f, binary.LittleEndian, &fileLength); err != nil {
		err = &BadHeader{0, err}
		return
	}
	if s.Compressed == COMPRESS_ZLIB {
		var d io.ReadCloser
		d, err = zlib.NewReader(f)
		if err != nil {
			err = &BadHeader{0, err}
			return
		}
		defer d.Close()
		f = d
	} else if s.Compressed == COMPRESS_LZMA {
		d := lzma.NewReader(f)
		defer d.Close()
		f = d
	}
	if _, err = s.FrameSize.ReadFrom(f); err != nil {
		err = &BadHeader{0, err}
		return
	}
	if err = binary.Read(f, binary.LittleEndian, &s.FrameRate); err != nil {
		err = &BadHeader{0, err}
		return
	}
	if err = binary.Read(f, binary.LittleEndian, &s.FrameCount); err != nil {
		err = &BadHeader{0, err}
		return
	}
	fmt.Println(s)
	// 	s.frames = make([]frame, 0, s.frameCount)
	// 	s.dictionary = make(dictionary)
	var (
		tagCode uint16
		length  uint32
	)
	if s.Version >= 8 {
		if err = binary.Read(f, binary.LittleEndian, &tagCode); err != nil {
			err = &BadHeader{0, err}
			return
		}
		length = uint32(tagCode & 63)
		tagCode >>= 6
		if length == 63 {
			if err = binary.Read(f, binary.LittleEndian, &length); err != nil {
				err = &BadHeader{0, err}
				return
			}
		}
		fmt.Printf("Tag: %d Length: %d\n", tagCode, length)
		lr := io.LimitReader(f, int64(length))
		if tagCode != 69 {
			err = &InvalidTagCode{tagCode}
			return
		}
		// 		if err = fileAttributes(s, lr); err != nil {
		// 			return
		// 		}
		_, err = io.Copy(ioutil.Discard, lr)
	}
	for {
		if err = binary.Read(f, binary.LittleEndian, &tagCode); err != nil {
			return
		}
		if tagCode == 0 {
			fmt.Println("END")
			break
		}
		length = uint32(tagCode & 63)
		tagCode >>= 6
		if length == 63 {
			if err = binary.Read(f, binary.LittleEndian, &length); err != nil {
				return
			}
		}
		fmt.Printf("Tag: %d Length: %d\n", tagCode, length)
		lr := io.LimitReader(f, int64(length))
		tag := TagFromIdVer(tagCode, s.Version)
		if tag == nil {
			err = &InvalidTagCode{tagCode}
			return
		}
		s.Tags = append(s.Tags, tag)
		if _, err = tag.ReadFrom(lr); err != nil {
			return
		}
		_, err = io.Copy(ioutil.Discard, lr)
		if err != nil && err != io.EOF {
			return
		}
	}
	return
}

func (s *SWF) WriteTo(f io.Writer) (total int64, err error) {
	if s.Version == 0 {
		var v uint8
		for _, tag := range s.Tags {
			if v = tag.MinVersion(); v > s.Version {
				s.Version = v
			}
		}
	} else {
		var v uint8
		for n, tag := range s.Tags {
			if v = tag.MinVersion(); v < s.Version {
				err = &ErrMinVersion{tag.Name(), v}
				return
			} else if u, ok := tag.(Upgradeable); ok && u.MaxVersion() < s.Version {
				s.Tags[n] = u.Upgrade(s.Version)
			}
		}
	}
	c := &rwcount.CountWriter{Writer: f}
	defer func() { total = c.BytesWritten() }()
	switch s.Compressed {
	case COMPRESS_NONE:
		if err = binary.Write(c, binary.LittleEndian, 'F'); err != nil {
			return
		}
	case COMPRESS_ZLIB:
		if err = binary.Write(c, binary.LittleEndian, 'C'); err != nil {
			return
		}
	case COMPRESS_LZMA:
		if err = binary.Write(c, binary.LittleEndian, 'Z'); err != nil {
			return
		}
	default:
		err = &BadHeader{Code: 2}
		return
	}
	if err = binary.Write(c, binary.LittleEndian, s.Version); err != nil {
		return
	}
	length := int32(3 + 1 + 4 + 16 + 2 + 2)
	for _, tag := range s.Tags {
		length += 2
		l := tag.Size()
		length += l
		if l >= 63 {
			length += 4
		}
	}
	if err = binary.Write(c, binary.LittleEndian, length); err != nil {
		return
	}
	if _, err = s.FrameSize.WriteTo(c); err != nil {
		return
	}
	if err = binary.Write(c, binary.LittleEndian, s.FrameRate); err != nil {
		return
	}
	if err = binary.Write(c, binary.LittleEndian, s.FrameCount); err != nil {
		return
	}
	for _, tag := range s.Tags {
		l := tag.Size()
		if l >= 63 {
			if err = binary.Write(c, binary.LittleEndian, tag.TagId()<<6|63); err != nil {
				return
			} else if err = binary.Write(c, binary.LittleEndian, l); err != nil {
				return
			}
		} else if err = binary.Write(c, binary.LittleEndian, tag.TagId()<<6|uint16(l)); err != nil {
			return
		}
		if _, err = tag.WriteTo(c); err != nil {
			return
		}
	}
	return
}
