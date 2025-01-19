// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package aiorder

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *AIOrder) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AIOrder[number], err)
}

func (x *AIOrder) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *AIOrder) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *AIOrder) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ModelName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AIOrder) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Prompt, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AIOrder) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Status = AIOrderStatus(v)
	return offset, nil
}

func (x *AIOrder) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Result, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AIOrder) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.CreateTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AIOrder) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.CompleteTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateAIOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateAIOrderReq[number], err)
}

func (x *CreateAIOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CreateAIOrderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ModelName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateAIOrderReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Prompt, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateAIOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateAIOrderResp[number], err)
}

func (x *CreateAIOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetAIOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetAIOrderReq[number], err)
}

func (x *GetAIOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetAIOrderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *GetAIOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetAIOrderResp[number], err)
}

func (x *GetAIOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v AIOrder
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Order = &v
	return offset, nil
}

func (x *CancelAIOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CancelAIOrderReq[number], err)
}

func (x *CancelAIOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *CancelAIOrderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CancelAIOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CancelAIOrderResp[number], err)
}

func (x *CancelAIOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *AIOrder) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *AIOrder) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *AIOrder) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *AIOrder) fastWriteField3(buf []byte) (offset int) {
	if x.ModelName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetModelName())
	return offset
}

func (x *AIOrder) fastWriteField4(buf []byte) (offset int) {
	if x.Prompt == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPrompt())
	return offset
}

func (x *AIOrder) fastWriteField5(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, int32(x.GetStatus()))
	return offset
}

func (x *AIOrder) fastWriteField6(buf []byte) (offset int) {
	if x.Result == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetResult())
	return offset
}

func (x *AIOrder) fastWriteField7(buf []byte) (offset int) {
	if x.CreateTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetCreateTime())
	return offset
}

func (x *AIOrder) fastWriteField8(buf []byte) (offset int) {
	if x.CompleteTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.GetCompleteTime())
	return offset
}

func (x *CreateAIOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateAIOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CreateAIOrderReq) fastWriteField2(buf []byte) (offset int) {
	if x.ModelName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetModelName())
	return offset
}

func (x *CreateAIOrderReq) fastWriteField3(buf []byte) (offset int) {
	if x.Prompt == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPrompt())
	return offset
}

func (x *CreateAIOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateAIOrderResp) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *GetAIOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetAIOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *GetAIOrderReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *GetAIOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetAIOrderResp) fastWriteField1(buf []byte) (offset int) {
	if x.Order == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetOrder())
	return offset
}

func (x *CancelAIOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CancelAIOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *CancelAIOrderReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *CancelAIOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CancelAIOrderResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *AIOrder) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *AIOrder) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetOrderId())
	return n
}

func (x *AIOrder) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetUserId())
	return n
}

func (x *AIOrder) sizeField3() (n int) {
	if x.ModelName == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetModelName())
	return n
}

func (x *AIOrder) sizeField4() (n int) {
	if x.Prompt == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPrompt())
	return n
}

func (x *AIOrder) sizeField5() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, int32(x.GetStatus()))
	return n
}

func (x *AIOrder) sizeField6() (n int) {
	if x.Result == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetResult())
	return n
}

func (x *AIOrder) sizeField7() (n int) {
	if x.CreateTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetCreateTime())
	return n
}

func (x *AIOrder) sizeField8() (n int) {
	if x.CompleteTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.GetCompleteTime())
	return n
}

func (x *CreateAIOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateAIOrderReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *CreateAIOrderReq) sizeField2() (n int) {
	if x.ModelName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetModelName())
	return n
}

func (x *CreateAIOrderReq) sizeField3() (n int) {
	if x.Prompt == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPrompt())
	return n
}

func (x *CreateAIOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateAIOrderResp) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetOrderId())
	return n
}

func (x *GetAIOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetAIOrderReq) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetOrderId())
	return n
}

func (x *GetAIOrderReq) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetUserId())
	return n
}

func (x *GetAIOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetAIOrderResp) sizeField1() (n int) {
	if x.Order == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetOrder())
	return n
}

func (x *CancelAIOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CancelAIOrderReq) sizeField1() (n int) {
	if x.OrderId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetOrderId())
	return n
}

func (x *CancelAIOrderReq) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetUserId())
	return n
}

func (x *CancelAIOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CancelAIOrderResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

var fieldIDToName_AIOrder = map[int32]string{
	1: "OrderId",
	2: "UserId",
	3: "ModelName",
	4: "Prompt",
	5: "Status",
	6: "Result",
	7: "CreateTime",
	8: "CompleteTime",
}

var fieldIDToName_CreateAIOrderReq = map[int32]string{
	1: "UserId",
	2: "ModelName",
	3: "Prompt",
}

var fieldIDToName_CreateAIOrderResp = map[int32]string{
	1: "OrderId",
}

var fieldIDToName_GetAIOrderReq = map[int32]string{
	1: "OrderId",
	2: "UserId",
}

var fieldIDToName_GetAIOrderResp = map[int32]string{
	1: "Order",
}

var fieldIDToName_CancelAIOrderReq = map[int32]string{
	1: "OrderId",
	2: "UserId",
}

var fieldIDToName_CancelAIOrderResp = map[int32]string{
	1: "Success",
}
