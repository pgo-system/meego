// Code generated by Kitex v0.7.1. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *UserInfo) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserInfo[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserInfo) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Uid = v

	}
	return offset, nil
}

func (p *UserInfo) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UserName = v

	}
	return offset, nil
}

func (p *UserInfo) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Points = v

	}
	return offset, nil
}

func (p *UserInfo) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.AvatarUrl = v

	}
	return offset, nil
}

func (p *UserInfo) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Email = v

	}
	return offset, nil
}

func (p *UserInfo) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PhoneNumber = v

	}
	return offset, nil
}

func (p *UserInfo) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PersonalSignature = v

	}
	return offset, nil
}

// for compatibility
func (p *UserInfo) FastWrite(buf []byte) int {
	return 0
}

func (p *UserInfo) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "UserInfo")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *UserInfo) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("UserInfo")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *UserInfo) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "uid", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Uid)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "user_name", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.UserName)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "points", thrift.I32, 3)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Points)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "avatar_url", thrift.STRING, 4)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.AvatarUrl)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "email", thrift.STRING, 5)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Email)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "phone_number", thrift.STRING, 6)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.PhoneNumber)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "personal_signature", thrift.STRING, 7)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.PersonalSignature)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *UserInfo) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("uid", thrift.I32, 1)
	l += bthrift.Binary.I32Length(p.Uid)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *UserInfo) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("user_name", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.UserName)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *UserInfo) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("points", thrift.I32, 3)
	l += bthrift.Binary.I32Length(p.Points)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *UserInfo) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("avatar_url", thrift.STRING, 4)
	l += bthrift.Binary.StringLengthNocopy(p.AvatarUrl)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *UserInfo) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("email", thrift.STRING, 5)
	l += bthrift.Binary.StringLengthNocopy(p.Email)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *UserInfo) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("phone_number", thrift.STRING, 6)
	l += bthrift.Binary.StringLengthNocopy(p.PhoneNumber)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *UserInfo) field7Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("personal_signature", thrift.STRING, 7)
	l += bthrift.Binary.StringLengthNocopy(p.PersonalSignature)

	l += bthrift.Binary.FieldEndLength()
	return l
}
