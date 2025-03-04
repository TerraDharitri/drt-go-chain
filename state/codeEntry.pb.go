// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: codeEntry.proto

package state

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CodeEntry struct {
	Code          []byte `protobuf:"bytes,1,opt,name=Code,proto3" json:"code,omitempty"`
	NumReferences uint32 `protobuf:"varint,2,opt,name=NumReferences,proto3" json:"numReferences"`
}

func (m *CodeEntry) Reset()      { *m = CodeEntry{} }
func (*CodeEntry) ProtoMessage() {}
func (*CodeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_12dd9a05813f72a9, []int{0}
}
func (m *CodeEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CodeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeEntry.Merge(m, src)
}
func (m *CodeEntry) XXX_Size() int {
	return m.Size()
}
func (m *CodeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CodeEntry proto.InternalMessageInfo

func (m *CodeEntry) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *CodeEntry) GetNumReferences() uint32 {
	if m != nil {
		return m.NumReferences
	}
	return 0
}

func init() {
	proto.RegisterType((*CodeEntry)(nil), "proto.CodeEntry")
}

func init() { proto.RegisterFile("codeEntry.proto", fileDescriptor_12dd9a05813f72a9) }

var fileDescriptor_12dd9a05813f72a9 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0x4f, 0x49,
	0x75, 0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52,
	0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9,
	0xfa, 0x60, 0xe1, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba, 0x94, 0x72, 0xb8,
	0x38, 0x9d, 0x61, 0x06, 0x09, 0xa9, 0x71, 0xb1, 0x80, 0x38, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x4e, 0x42, 0xaf, 0xee, 0xc9, 0xf3, 0x81, 0x6c, 0xd1, 0xc9, 0xcf, 0xcd, 0x2c, 0x49, 0xcd, 0x2d,
	0x28, 0xa9, 0x0c, 0x02, 0xcb, 0x0b, 0x99, 0x73, 0xf1, 0xfa, 0x95, 0xe6, 0x06, 0xa5, 0xa6, 0xa5,
	0x16, 0xa5, 0xe6, 0x25, 0xa7, 0x16, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x3a, 0x09, 0xbe, 0xba,
	0x27, 0xcf, 0x9b, 0x87, 0x2c, 0x11, 0x84, 0xaa, 0xce, 0xc9, 0xfe, 0xc2, 0x43, 0x39, 0x86, 0x1b,
	0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0xde, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xac, 0xc5, 0x25, 0x89, 0x25, 0xa9, 0x49, 0x6c, 0x60,
	0x57, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0x44, 0x30, 0x6d, 0xfe, 0x00, 0x00, 0x00,
}

func (this *CodeEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodeEntry)
	if !ok {
		that2, ok := that.(CodeEntry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Code, that1.Code) {
		return false
	}
	if this.NumReferences != that1.NumReferences {
		return false
	}
	return true
}
func (this *CodeEntry) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&state.CodeEntry{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "NumReferences: "+fmt.Sprintf("%#v", this.NumReferences)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCodeEntry(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CodeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumReferences != 0 {
		i = encodeVarintCodeEntry(dAtA, i, uint64(m.NumReferences))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintCodeEntry(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCodeEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodeEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CodeEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCodeEntry(uint64(l))
	}
	if m.NumReferences != 0 {
		n += 1 + sovCodeEntry(uint64(m.NumReferences))
	}
	return n
}

func sovCodeEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodeEntry(x uint64) (n int) {
	return sovCodeEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CodeEntry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CodeEntry{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`NumReferences:` + fmt.Sprintf("%v", this.NumReferences) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCodeEntry(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CodeEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodeEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CodeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodeEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumReferences", wireType)
			}
			m.NumReferences = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumReferences |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodeEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCodeEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCodeEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodeEntry
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
					return 0, ErrIntOverflowCodeEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodeEntry
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
			if length < 0 {
				return 0, ErrInvalidLengthCodeEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodeEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodeEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodeEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodeEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodeEntry = fmt.Errorf("proto: unexpected end of group")
)
