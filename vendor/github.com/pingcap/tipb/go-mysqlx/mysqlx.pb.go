// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/pingcap/tipb/go-mysqlx/mysqlx.proto

/*
	Package Mysqlx is a generated protocol buffer package.

	It is generated from these files:
		github.com/pingcap/tipb/go-mysqlx/mysqlx.proto

	It has these top-level messages:
		ClientMessages
		ServerMessages
		Ok
		Error
*/
package Mysqlx

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientMessages_Type int32

const (
	ClientMessages_CON_CAPABILITIES_GET       ClientMessages_Type = 1
	ClientMessages_CON_CAPABILITIES_SET       ClientMessages_Type = 2
	ClientMessages_CON_CLOSE                  ClientMessages_Type = 3
	ClientMessages_SESS_AUTHENTICATE_START    ClientMessages_Type = 4
	ClientMessages_SESS_AUTHENTICATE_CONTINUE ClientMessages_Type = 5
	ClientMessages_SESS_RESET                 ClientMessages_Type = 6
	ClientMessages_SESS_CLOSE                 ClientMessages_Type = 7
	ClientMessages_SQL_STMT_EXECUTE           ClientMessages_Type = 12
	ClientMessages_CRUD_FIND                  ClientMessages_Type = 17
	ClientMessages_CRUD_INSERT                ClientMessages_Type = 18
	ClientMessages_CRUD_UPDATE                ClientMessages_Type = 19
	ClientMessages_CRUD_DELETE                ClientMessages_Type = 20
	ClientMessages_EXPECT_OPEN                ClientMessages_Type = 24
	ClientMessages_EXPECT_CLOSE               ClientMessages_Type = 25
	ClientMessages_CRUD_CREATE_VIEW           ClientMessages_Type = 30
	ClientMessages_CRUD_MODIFY_VIEW           ClientMessages_Type = 31
	ClientMessages_CRUD_DROP_VIEW             ClientMessages_Type = 32
)

var ClientMessages_Type_name = map[int32]string{
	1:  "CON_CAPABILITIES_GET",
	2:  "CON_CAPABILITIES_SET",
	3:  "CON_CLOSE",
	4:  "SESS_AUTHENTICATE_START",
	5:  "SESS_AUTHENTICATE_CONTINUE",
	6:  "SESS_RESET",
	7:  "SESS_CLOSE",
	12: "SQL_STMT_EXECUTE",
	17: "CRUD_FIND",
	18: "CRUD_INSERT",
	19: "CRUD_UPDATE",
	20: "CRUD_DELETE",
	24: "EXPECT_OPEN",
	25: "EXPECT_CLOSE",
	30: "CRUD_CREATE_VIEW",
	31: "CRUD_MODIFY_VIEW",
	32: "CRUD_DROP_VIEW",
}
var ClientMessages_Type_value = map[string]int32{
	"CON_CAPABILITIES_GET":       1,
	"CON_CAPABILITIES_SET":       2,
	"CON_CLOSE":                  3,
	"SESS_AUTHENTICATE_START":    4,
	"SESS_AUTHENTICATE_CONTINUE": 5,
	"SESS_RESET":                 6,
	"SESS_CLOSE":                 7,
	"SQL_STMT_EXECUTE":           12,
	"CRUD_FIND":                  17,
	"CRUD_INSERT":                18,
	"CRUD_UPDATE":                19,
	"CRUD_DELETE":                20,
	"EXPECT_OPEN":                24,
	"EXPECT_CLOSE":               25,
	"CRUD_CREATE_VIEW":           30,
	"CRUD_MODIFY_VIEW":           31,
	"CRUD_DROP_VIEW":             32,
}

func (x ClientMessages_Type) Enum() *ClientMessages_Type {
	p := new(ClientMessages_Type)
	*p = x
	return p
}
func (x ClientMessages_Type) String() string {
	return proto.EnumName(ClientMessages_Type_name, int32(x))
}
func (x *ClientMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientMessages_Type_value, data, "ClientMessages_Type")
	if err != nil {
		return err
	}
	*x = ClientMessages_Type(value)
	return nil
}
func (ClientMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{0, 0} }

type ServerMessages_Type int32

const (
	ServerMessages_OK                         ServerMessages_Type = 0
	ServerMessages_ERROR                      ServerMessages_Type = 1
	ServerMessages_CONN_CAPABILITIES          ServerMessages_Type = 2
	ServerMessages_SESS_AUTHENTICATE_CONTINUE ServerMessages_Type = 3
	ServerMessages_SESS_AUTHENTICATE_OK       ServerMessages_Type = 4
	// NOTICE has to stay at 11 forever
	ServerMessages_NOTICE                               ServerMessages_Type = 11
	ServerMessages_RESULTSET_COLUMN_META_DATA           ServerMessages_Type = 12
	ServerMessages_RESULTSET_ROW                        ServerMessages_Type = 13
	ServerMessages_RESULTSET_FETCH_DONE                 ServerMessages_Type = 14
	ServerMessages_RESULTSET_FETCH_SUSPENDED            ServerMessages_Type = 15
	ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS ServerMessages_Type = 16
	ServerMessages_SQL_STMT_EXECUTE_OK                  ServerMessages_Type = 17
	ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS ServerMessages_Type = 18
)

var ServerMessages_Type_name = map[int32]string{
	0:  "OK",
	1:  "ERROR",
	2:  "CONN_CAPABILITIES",
	3:  "SESS_AUTHENTICATE_CONTINUE",
	4:  "SESS_AUTHENTICATE_OK",
	11: "NOTICE",
	12: "RESULTSET_COLUMN_META_DATA",
	13: "RESULTSET_ROW",
	14: "RESULTSET_FETCH_DONE",
	15: "RESULTSET_FETCH_SUSPENDED",
	16: "RESULTSET_FETCH_DONE_MORE_RESULTSETS",
	17: "SQL_STMT_EXECUTE_OK",
	18: "RESULTSET_FETCH_DONE_MORE_OUT_PARAMS",
}
var ServerMessages_Type_value = map[string]int32{
	"OK":                                   0,
	"ERROR":                                1,
	"CONN_CAPABILITIES":                    2,
	"SESS_AUTHENTICATE_CONTINUE":           3,
	"SESS_AUTHENTICATE_OK":                 4,
	"NOTICE":                               11,
	"RESULTSET_COLUMN_META_DATA":           12,
	"RESULTSET_ROW":                        13,
	"RESULTSET_FETCH_DONE":                 14,
	"RESULTSET_FETCH_SUSPENDED":            15,
	"RESULTSET_FETCH_DONE_MORE_RESULTSETS": 16,
	"SQL_STMT_EXECUTE_OK":                  17,
	"RESULTSET_FETCH_DONE_MORE_OUT_PARAMS": 18,
}

func (x ServerMessages_Type) Enum() *ServerMessages_Type {
	p := new(ServerMessages_Type)
	*p = x
	return p
}
func (x ServerMessages_Type) String() string {
	return proto.EnumName(ServerMessages_Type_name, int32(x))
}
func (x *ServerMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServerMessages_Type_value, data, "ServerMessages_Type")
	if err != nil {
		return err
	}
	*x = ServerMessages_Type(value)
	return nil
}
func (ServerMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{1, 0} }

type Error_Severity int32

const (
	Error_ERROR Error_Severity = 0
	Error_FATAL Error_Severity = 1
)

var Error_Severity_name = map[int32]string{
	0: "ERROR",
	1: "FATAL",
}
var Error_Severity_value = map[string]int32{
	"ERROR": 0,
	"FATAL": 1,
}

func (x Error_Severity) Enum() *Error_Severity {
	p := new(Error_Severity)
	*p = x
	return p
}
func (x Error_Severity) String() string {
	return proto.EnumName(Error_Severity_name, int32(x))
}
func (x *Error_Severity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_Severity_value, data, "Error_Severity")
	if err != nil {
		return err
	}
	*x = Error_Severity(value)
	return nil
}
func (Error_Severity) EnumDescriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{3, 0} }

// IDs of messages that can be sent from client to the server
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ClientMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientMessages) Reset()                    { *m = ClientMessages{} }
func (m *ClientMessages) String() string            { return proto.CompactTextString(m) }
func (*ClientMessages) ProtoMessage()               {}
func (*ClientMessages) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{0} }

// IDs of messages that can be sent from server to client
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ServerMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServerMessages) Reset()                    { *m = ServerMessages{} }
func (m *ServerMessages) String() string            { return proto.CompactTextString(m) }
func (*ServerMessages) ProtoMessage()               {}
func (*ServerMessages) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{1} }

// generic Ok message
type Ok struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Ok) Reset()                    { *m = Ok{} }
func (m *Ok) String() string            { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()               {}
func (*Ok) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{2} }

func (m *Ok) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// generic Error message
//
// A ``severity`` of ``ERROR`` indicates the current message sequence is
// aborted for the given error and the session is ready for more.
//
// In case of a ``FATAL`` error message the client should not expect
// the server to continue handling any further messages and should
// close the connection.
//
// :param severity: severity of the error message
// :param code: error-code
// :param sql_state: SQL state
// :param msg: human readable error message
type Error struct {
	Severity         *Error_Severity `protobuf:"varint,1,opt,name=severity,enum=Mysqlx.Error_Severity,def=0" json:"severity,omitempty"`
	Code             *uint32         `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	SqlState         *string         `protobuf:"bytes,4,req,name=sql_state,json=sqlState" json:"sql_state,omitempty"`
	Msg              *string         `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorMysqlx, []int{3} }

const Default_Error_Severity Error_Severity = Error_ERROR

func (m *Error) GetSeverity() Error_Severity {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return Default_Error_Severity
}

func (m *Error) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Error) GetSqlState() string {
	if m != nil && m.SqlState != nil {
		return *m.SqlState
	}
	return ""
}

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientMessages)(nil), "Mysqlx.ClientMessages")
	proto.RegisterType((*ServerMessages)(nil), "Mysqlx.ServerMessages")
	proto.RegisterType((*Ok)(nil), "Mysqlx.Ok")
	proto.RegisterType((*Error)(nil), "Mysqlx.Error")
	proto.RegisterEnum("Mysqlx.ClientMessages_Type", ClientMessages_Type_name, ClientMessages_Type_value)
	proto.RegisterEnum("Mysqlx.ServerMessages_Type", ServerMessages_Type_name, ServerMessages_Type_value)
	proto.RegisterEnum("Mysqlx.Error_Severity", Error_Severity_name, Error_Severity_value)
}
func (m *ClientMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMessages) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ServerMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerMessages) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Ok) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ok) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(len(*m.Msg)))
		i += copy(dAtA[i:], *m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Severity != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(*m.Severity))
	}
	if m.Code == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(*m.Code))
	}
	if m.Msg == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(len(*m.Msg)))
		i += copy(dAtA[i:], *m.Msg)
	}
	if m.SqlState == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMysqlx(dAtA, i, uint64(len(*m.SqlState)))
		i += copy(dAtA[i:], *m.SqlState)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMysqlx(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClientMessages) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServerMessages) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Ok) Size() (n int) {
	var l int
	_ = l
	if m.Msg != nil {
		l = len(*m.Msg)
		n += 1 + l + sovMysqlx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Error) Size() (n int) {
	var l int
	_ = l
	if m.Severity != nil {
		n += 1 + sovMysqlx(uint64(*m.Severity))
	}
	if m.Code != nil {
		n += 1 + sovMysqlx(uint64(*m.Code))
	}
	if m.Msg != nil {
		l = len(*m.Msg)
		n += 1 + l + sovMysqlx(uint64(l))
	}
	if m.SqlState != nil {
		l = len(*m.SqlState)
		n += 1 + l + sovMysqlx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMysqlx(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMysqlx(x uint64) (n int) {
	return sovMysqlx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServerMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServerMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ok) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ok: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ok: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Msg = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Error) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMysqlx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			var v Error_Severity
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (Error_Severity(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Severity = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Code = &v
			hasFields[0] |= uint64(0x00000001)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Msg = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SqlState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMysqlx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.SqlState = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipMysqlx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMysqlx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMysqlx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMysqlx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMysqlx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMysqlx
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMysqlx
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMysqlx(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMysqlx = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMysqlx   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/pingcap/tipb/go-mysqlx/mysqlx.proto", fileDescriptorMysqlx)
}

var fileDescriptorMysqlx = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x52, 0xd3, 0x5c,
	0x10, 0xc7, 0x49, 0x5a, 0xfa, 0xd1, 0x85, 0x96, 0xed, 0x81, 0x0f, 0x8a, 0x68, 0xed, 0x74, 0xbc,
	0xe0, 0x2a, 0xce, 0x78, 0xa7, 0x77, 0x21, 0xd9, 0x4a, 0x86, 0x24, 0x27, 0x9e, 0x73, 0x22, 0x78,
	0x75, 0x06, 0x31, 0x32, 0x68, 0xb1, 0x90, 0x54, 0x46, 0xde, 0xc4, 0x37, 0xf0, 0x55, 0xbc, 0xd3,
	0x47, 0xd0, 0xfa, 0x0a, 0x3e, 0x80, 0x93, 0x04, 0x52, 0x05, 0x46, 0xaf, 0x92, 0xfd, 0xfd, 0xf7,
	0xfc, 0x77, 0x67, 0x77, 0x16, 0x56, 0x82, 0x8b, 0xec, 0x6c, 0xf4, 0xe1, 0xe1, 0x49, 0xf1, 0xb1,
	0x4e, 0xd3, 0xf1, 0x64, 0xcc, 0x1a, 0x25, 0x1c, 0xfc, 0x34, 0xa1, 0xed, 0x8c, 0x8e, 0x93, 0x77,
	0x93, 0x20, 0xc9, 0xb2, 0x83, 0xa3, 0x24, 0x1b, 0x7c, 0x37, 0xa1, 0xae, 0x2e, 0x4e, 0x13, 0xd6,
	0x85, 0x55, 0x87, 0x87, 0xda, 0xb1, 0x23, 0x7b, 0xdb, 0xf3, 0x3d, 0xe5, 0x91, 0xd4, 0x4f, 0x49,
	0xa1, 0x71, 0xab, 0x22, 0x49, 0xa1, 0xc9, 0x5a, 0xd0, 0x2c, 0x14, 0x9f, 0x4b, 0xc2, 0x1a, 0xdb,
	0x84, 0x75, 0x49, 0x52, 0x6a, 0x3b, 0x56, 0x3b, 0x14, 0x2a, 0xcf, 0xb1, 0x15, 0x69, 0xa9, 0x6c,
	0xa1, 0xb0, 0xce, 0x7a, 0x70, 0xe7, 0xa6, 0xe8, 0xf0, 0x50, 0x79, 0x61, 0x4c, 0x38, 0xcf, 0xda,
	0x00, 0x85, 0x2e, 0x28, 0xf7, 0x6e, 0x54, 0x71, 0x69, 0xfe, 0x1f, 0x5b, 0x05, 0x94, 0xcf, 0x7c,
	0x2d, 0x55, 0xa0, 0x34, 0xed, 0x93, 0x13, 0x2b, 0xc2, 0xa5, 0xa2, 0x03, 0x11, 0xbb, 0x7a, 0xe8,
	0x85, 0x2e, 0x76, 0xd8, 0x32, 0x2c, 0x16, 0xa1, 0x17, 0x4a, 0x12, 0x0a, 0x59, 0x05, 0xe2, 0xc8,
	0xb5, 0x15, 0xe1, 0x4a, 0x05, 0x5c, 0xf2, 0x49, 0x11, 0xae, 0xe6, 0x80, 0xf6, 0x23, 0x72, 0x94,
	0xe6, 0x11, 0x85, 0xd8, 0x65, 0x08, 0x4b, 0x97, 0xa0, 0x2c, 0xbd, 0x91, 0x97, 0x2e, 0xde, 0x38,
	0x82, 0xf2, 0xa6, 0x9f, 0x7b, 0xb4, 0x87, 0xbd, 0x8a, 0x06, 0xdc, 0xf5, 0x86, 0x2f, 0x4a, 0x7a,
	0x9f, 0x31, 0x68, 0x97, 0xfe, 0x82, 0x47, 0x25, 0xeb, 0x0f, 0xa6, 0x26, 0xb4, 0x65, 0x92, 0x9e,
	0x27, 0x69, 0x35, 0xf6, 0x2f, 0x57, 0x63, 0x6f, 0x80, 0xc9, 0x77, 0x71, 0x8e, 0x35, 0x61, 0x9e,
	0x84, 0xe0, 0x02, 0x0d, 0xf6, 0x3f, 0x74, 0x1c, 0x1e, 0xfe, 0x39, 0x70, 0x34, 0xff, 0x31, 0xc0,
	0x5a, 0xbe, 0xa6, 0x9b, 0x3a, 0xdf, 0xc5, 0x3a, 0x03, 0x68, 0x84, 0x5c, 0x79, 0x0e, 0xe1, 0x62,
	0xee, 0x22, 0x48, 0xc6, 0xbe, 0x92, 0xa4, 0xb4, 0xc3, 0xfd, 0x38, 0x08, 0x75, 0x40, 0xca, 0xd6,
	0xae, 0xad, 0x6c, 0x5c, 0x62, 0x1d, 0x68, 0xcd, 0x74, 0xc1, 0xf7, 0xb0, 0x95, 0x1b, 0xcf, 0xd0,
	0x90, 0x94, 0xb3, 0xa3, 0x5d, 0x1e, 0x12, 0xb6, 0xd9, 0x3d, 0xd8, 0xb8, 0xae, 0xc8, 0x58, 0x46,
	0x14, 0xba, 0xe4, 0xe2, 0x32, 0xdb, 0x82, 0x07, 0xb7, 0x3d, 0xd4, 0x01, 0x17, 0xa4, 0x2b, 0x45,
	0x22, 0xb2, 0x75, 0x58, 0xb9, 0xbe, 0xdc, 0xbc, 0xf5, 0xce, 0xdf, 0x2d, 0x78, 0xac, 0x74, 0x64,
	0x0b, 0x3b, 0x90, 0xc8, 0x06, 0x6b, 0x60, 0xf2, 0xb7, 0x0c, 0xa1, 0x76, 0x92, 0x1d, 0x75, 0x8d,
	0xbe, 0xb1, 0xd5, 0x14, 0xf9, 0xef, 0xe0, 0x93, 0x01, 0xf3, 0x94, 0xa6, 0xe3, 0x94, 0x3d, 0x86,
	0x85, 0x2c, 0x39, 0x4f, 0xd2, 0xe3, 0xc9, 0x45, 0x91, 0xd0, 0x7e, 0xb4, 0x66, 0x95, 0x87, 0x61,
	0x15, 0x09, 0x96, 0xbc, 0x54, 0x9f, 0x94, 0xab, 0x10, 0x55, 0x3a, 0x63, 0x50, 0x3f, 0x1c, 0xbf,
	0x4a, 0xba, 0x66, 0xdf, 0xdc, 0x6a, 0x89, 0xe2, 0xff, 0xaa, 0x54, 0xad, 0x6f, 0x5e, 0x96, 0x62,
	0x9b, 0xd0, 0xcc, 0xce, 0x46, 0x3a, 0x9b, 0x1c, 0x4c, 0x92, 0x6e, 0xbd, 0xe0, 0x0b, 0xd9, 0xd9,
	0x48, 0xe6, 0xf1, 0xa0, 0x0f, 0x0b, 0x57, 0xfe, 0xb3, 0x65, 0x17, 0x7b, 0x1f, 0xda, 0xca, 0xf6,
	0xd1, 0xd8, 0xb6, 0x3e, 0x4f, 0x7b, 0xc6, 0xd7, 0x69, 0xcf, 0xf8, 0x36, 0xed, 0x19, 0x1f, 0x7f,
	0xf4, 0xe6, 0xe0, 0xee, 0xe1, 0xf8, 0xc4, 0x2a, 0x2e, 0xd9, 0x3a, 0x7c, 0x63, 0xfd, 0x7e, 0xd2,
	0x2f, 0xdf, 0xbf, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x03, 0xb4, 0x38, 0x6c, 0xeb, 0x03, 0x00,
	0x00,
}