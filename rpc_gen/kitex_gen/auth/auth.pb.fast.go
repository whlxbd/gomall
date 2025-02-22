// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package auth

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Rule) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Rule[number], err)
}

func (x *Rule) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Rule) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Role, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Rule) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Touter, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliverTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeliverTokenReq[number], err)
}

func (x *DeliverTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *VerifyTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VerifyTokenReq[number], err)
}

func (x *VerifyTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeliveryResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeliveryResp[number], err)
}

func (x *DeliveryResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VerifyResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VerifyResp[number], err)
}

func (x *VerifyResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Res, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetPayloadReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPayloadReq[number], err)
}

func (x *GetPayloadReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetPayloadResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPayloadResp[number], err)
}

func (x *GetPayloadResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetPayloadResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AuthenticateReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AuthenticateReq[number], err)
}

func (x *AuthenticateReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Role, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AuthenticateReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Router, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AuthenticateResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AuthenticateResp[number], err)
}

func (x *AuthenticateResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Ok, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *CreateReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateReq[number], err)
}

func (x *CreateReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Role, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Router, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *ListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListReq[number], err)
}

func (x *ListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Pagesize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListResp[number], err)
}

func (x *ListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Rule
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Rule = append(x.Rule, &v)
	return offset, nil
}

func (x *DeleteReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteReq[number], err)
}

func (x *DeleteReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *DeleteResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *GetReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetReq[number], err)
}

func (x *GetReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetResp[number], err)
}

func (x *GetResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Rule
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Rule = &v
	return offset, nil
}

func (x *UpdateReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateReq[number], err)
}

func (x *UpdateReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Rule
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Rule = &v
	return offset, nil
}

func (x *UpdateResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *Rule) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *Rule) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Rule) fastWriteField2(buf []byte) (offset int) {
	if x.Role == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRole())
	return offset
}

func (x *Rule) fastWriteField3(buf []byte) (offset int) {
	if x.Touter == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetTouter())
	return offset
}

func (x *DeliverTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeliverTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *VerifyTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *VerifyTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *DeliveryResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeliveryResp) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *VerifyResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *VerifyResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Res {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetRes())
	return offset
}

func (x *GetPayloadReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetPayloadReq) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *GetPayloadResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetPayloadResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetPayloadResp) fastWriteField2(buf []byte) (offset int) {
	if x.Type == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetType())
	return offset
}

func (x *AuthenticateReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AuthenticateReq) fastWriteField1(buf []byte) (offset int) {
	if x.Role == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetRole())
	return offset
}

func (x *AuthenticateReq) fastWriteField2(buf []byte) (offset int) {
	if x.Router == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRouter())
	return offset
}

func (x *AuthenticateResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AuthenticateResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Ok {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetOk())
	return offset
}

func (x *CreateReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateReq) fastWriteField1(buf []byte) (offset int) {
	if x.Role == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetRole())
	return offset
}

func (x *CreateReq) fastWriteField2(buf []byte) (offset int) {
	if x.Router == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRouter())
	return offset
}

func (x *CreateResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *ListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListReq) fastWriteField1(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetPage())
	return offset
}

func (x *ListReq) fastWriteField2(buf []byte) (offset int) {
	if x.Pagesize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetPagesize())
	return offset
}

func (x *ListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ListResp) fastWriteField1(buf []byte) (offset int) {
	if x.Rule == nil {
		return offset
	}
	for i := range x.GetRule() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRule()[i])
	}
	return offset
}

func (x *DeleteReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *DeleteResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetResp) fastWriteField1(buf []byte) (offset int) {
	if x.Rule == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRule())
	return offset
}

func (x *UpdateReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateReq) fastWriteField1(buf []byte) (offset int) {
	if x.Rule == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRule())
	return offset
}

func (x *UpdateResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *Rule) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *Rule) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetId())
	return n
}

func (x *Rule) sizeField2() (n int) {
	if x.Role == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRole())
	return n
}

func (x *Rule) sizeField3() (n int) {
	if x.Touter == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetTouter())
	return n
}

func (x *DeliverTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeliverTokenReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *VerifyTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *VerifyTokenReq) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *DeliveryResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeliveryResp) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *VerifyResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *VerifyResp) sizeField1() (n int) {
	if !x.Res {
		return n
	}
	n += fastpb.SizeBool(1, x.GetRes())
	return n
}

func (x *GetPayloadReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetPayloadReq) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *GetPayloadResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetPayloadResp) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUserId())
	return n
}

func (x *GetPayloadResp) sizeField2() (n int) {
	if x.Type == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetType())
	return n
}

func (x *AuthenticateReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AuthenticateReq) sizeField1() (n int) {
	if x.Role == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetRole())
	return n
}

func (x *AuthenticateReq) sizeField2() (n int) {
	if x.Router == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRouter())
	return n
}

func (x *AuthenticateResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AuthenticateResp) sizeField1() (n int) {
	if !x.Ok {
		return n
	}
	n += fastpb.SizeBool(1, x.GetOk())
	return n
}

func (x *CreateReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateReq) sizeField1() (n int) {
	if x.Role == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetRole())
	return n
}

func (x *CreateReq) sizeField2() (n int) {
	if x.Router == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRouter())
	return n
}

func (x *CreateResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *ListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListReq) sizeField1() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetPage())
	return n
}

func (x *ListReq) sizeField2() (n int) {
	if x.Pagesize == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetPagesize())
	return n
}

func (x *ListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ListResp) sizeField1() (n int) {
	if x.Rule == nil {
		return n
	}
	for i := range x.GetRule() {
		n += fastpb.SizeMessage(1, x.GetRule()[i])
	}
	return n
}

func (x *DeleteReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetId())
	return n
}

func (x *DeleteResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetId())
	return n
}

func (x *GetResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetResp) sizeField1() (n int) {
	if x.Rule == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRule())
	return n
}

func (x *UpdateReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateReq) sizeField1() (n int) {
	if x.Rule == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRule())
	return n
}

func (x *UpdateResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_Rule = map[int32]string{
	1: "Id",
	2: "Role",
	3: "Touter",
}

var fieldIDToName_DeliverTokenReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_VerifyTokenReq = map[int32]string{
	1: "Token",
}

var fieldIDToName_DeliveryResp = map[int32]string{
	1: "Token",
}

var fieldIDToName_VerifyResp = map[int32]string{
	1: "Res",
}

var fieldIDToName_GetPayloadReq = map[int32]string{
	1: "Token",
}

var fieldIDToName_GetPayloadResp = map[int32]string{
	1: "UserId",
	2: "Type",
}

var fieldIDToName_AuthenticateReq = map[int32]string{
	1: "Role",
	2: "Router",
}

var fieldIDToName_AuthenticateResp = map[int32]string{
	1: "Ok",
}

var fieldIDToName_CreateReq = map[int32]string{
	1: "Role",
	2: "Router",
}

var fieldIDToName_CreateResp = map[int32]string{}

var fieldIDToName_ListReq = map[int32]string{
	1: "Page",
	2: "Pagesize",
}

var fieldIDToName_ListResp = map[int32]string{
	1: "Rule",
}

var fieldIDToName_DeleteReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_DeleteResp = map[int32]string{}

var fieldIDToName_GetReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_GetResp = map[int32]string{
	1: "Rule",
}

var fieldIDToName_UpdateReq = map[int32]string{
	1: "Rule",
}

var fieldIDToName_UpdateResp = map[int32]string{}
