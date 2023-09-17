// Code generated by thriftgo (0.3.1). DO NOT EDIT.

package base

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type UserInfo struct {
	Uid               int32  `thrift:"uid,1" frugal:"1,default,i32" json:"uid"`
	UserName          string `thrift:"user_name,2" frugal:"2,default,string" json:"user_name"`
	Points            int32  `thrift:"points,3" frugal:"3,default,i32" json:"points"`
	AvatarUrl         string `thrift:"avatar_url,4" frugal:"4,default,string" json:"avatar_url"`
	Email             string `thrift:"email,5" frugal:"5,default,string" json:"email"`
	PhoneNumber       string `thrift:"phone_number,6" frugal:"6,default,string" json:"phone_number"`
	PersonalSignature string `thrift:"personal_signature,7" frugal:"7,default,string" json:"personal_signature"`
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (p *UserInfo) InitDefault() {
	*p = UserInfo{}
}

func (p *UserInfo) GetUid() (v int32) {
	return p.Uid
}

func (p *UserInfo) GetUserName() (v string) {
	return p.UserName
}

func (p *UserInfo) GetPoints() (v int32) {
	return p.Points
}

func (p *UserInfo) GetAvatarUrl() (v string) {
	return p.AvatarUrl
}

func (p *UserInfo) GetEmail() (v string) {
	return p.Email
}

func (p *UserInfo) GetPhoneNumber() (v string) {
	return p.PhoneNumber
}

func (p *UserInfo) GetPersonalSignature() (v string) {
	return p.PersonalSignature
}
func (p *UserInfo) SetUid(val int32) {
	p.Uid = val
}
func (p *UserInfo) SetUserName(val string) {
	p.UserName = val
}
func (p *UserInfo) SetPoints(val int32) {
	p.Points = val
}
func (p *UserInfo) SetAvatarUrl(val string) {
	p.AvatarUrl = val
}
func (p *UserInfo) SetEmail(val string) {
	p.Email = val
}
func (p *UserInfo) SetPhoneNumber(val string) {
	p.PhoneNumber = val
}
func (p *UserInfo) SetPersonalSignature(val string) {
	p.PersonalSignature = val
}

var fieldIDToName_UserInfo = map[int16]string{
	1: "uid",
	2: "user_name",
	3: "points",
	4: "avatar_url",
	5: "email",
	6: "phone_number",
	7: "personal_signature",
}

func (p *UserInfo) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField7(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserInfo[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserInfo) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Uid = v
	}
	return nil
}

func (p *UserInfo) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.UserName = v
	}
	return nil
}

func (p *UserInfo) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Points = v
	}
	return nil
}

func (p *UserInfo) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.AvatarUrl = v
	}
	return nil
}

func (p *UserInfo) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Email = v
	}
	return nil
}

func (p *UserInfo) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PhoneNumber = v
	}
	return nil
}

func (p *UserInfo) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PersonalSignature = v
	}
	return nil
}

func (p *UserInfo) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserInfo"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
		if err = p.writeField7(oprot); err != nil {
			fieldId = 7
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("uid", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Uid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UserInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("points", thrift.I32, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Points); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *UserInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("avatar_url", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.AvatarUrl); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *UserInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("email", thrift.STRING, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Email); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *UserInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("phone_number", thrift.STRING, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PhoneNumber); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *UserInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("personal_signature", thrift.STRING, 7); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PersonalSignature); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 end error: ", p), err)
}

func (p *UserInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserInfo(%+v)", *p)
}

func (p *UserInfo) DeepEqual(ano *UserInfo) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Uid) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserName) {
		return false
	}
	if !p.Field3DeepEqual(ano.Points) {
		return false
	}
	if !p.Field4DeepEqual(ano.AvatarUrl) {
		return false
	}
	if !p.Field5DeepEqual(ano.Email) {
		return false
	}
	if !p.Field6DeepEqual(ano.PhoneNumber) {
		return false
	}
	if !p.Field7DeepEqual(ano.PersonalSignature) {
		return false
	}
	return true
}

func (p *UserInfo) Field1DeepEqual(src int32) bool {

	if p.Uid != src {
		return false
	}
	return true
}
func (p *UserInfo) Field2DeepEqual(src string) bool {

	if strings.Compare(p.UserName, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field3DeepEqual(src int32) bool {

	if p.Points != src {
		return false
	}
	return true
}
func (p *UserInfo) Field4DeepEqual(src string) bool {

	if strings.Compare(p.AvatarUrl, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field5DeepEqual(src string) bool {

	if strings.Compare(p.Email, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field6DeepEqual(src string) bool {

	if strings.Compare(p.PhoneNumber, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field7DeepEqual(src string) bool {

	if strings.Compare(p.PersonalSignature, src) != 0 {
		return false
	}
	return true
}
