// Code generated by protoc-gen-gogo.
// source: file.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		file.proto

	It has these top-level messages:
		File
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type File struct {
	Path     *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Key      []byte  `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Hash     []byte  `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	FileSize *int64  `protobuf:"varint,4,req,name=file_size" json:"file_size,omitempty"`
	Kind     *int32  `protobuf:"varint,5,req,name=kind" json:"kind,omitempty"`
	// Timestamp formated as RFC 3339
	ModTime          []byte `protobuf:"bytes,6,req,name=mod_time" json:"mod_time,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto1.CompactTextString(m) }
func (*File) ProtoMessage()    {}

func (m *File) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *File) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *File) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *File) GetFileSize() int64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *File) GetKind() int32 {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return 0
}

func (m *File) GetModTime() []byte {
	if m != nil {
		return m.ModTime
	}
	return nil
}

func init() {
	proto1.RegisterType((*File)(nil), "bolt.protocol.File")
}
func (m *File) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *File) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Path == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintFile(data, i, uint64(len(*m.Path)))
		i += copy(data[i:], *m.Path)
	}
	if m.Key != nil {
		data[i] = 0x12
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if m.Hash != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Hash)))
		i += copy(data[i:], m.Hash)
	}
	if m.FileSize == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x20
		i++
		i = encodeVarintFile(data, i, uint64(*m.FileSize))
	}
	if m.Kind == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x28
		i++
		i = encodeVarintFile(data, i, uint64(*m.Kind))
	}
	if m.ModTime == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x32
		i++
		i = encodeVarintFile(data, i, uint64(len(m.ModTime)))
		i += copy(data[i:], m.ModTime)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64File(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32File(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFile(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *File) Size() (n int) {
	var l int
	_ = l
	if m.Path != nil {
		l = len(*m.Path)
		n += 1 + l + sovFile(uint64(l))
	}
	if m.Key != nil {
		l = len(m.Key)
		n += 1 + l + sovFile(uint64(l))
	}
	if m.Hash != nil {
		l = len(m.Hash)
		n += 1 + l + sovFile(uint64(l))
	}
	if m.FileSize != nil {
		n += 1 + sovFile(uint64(*m.FileSize))
	}
	if m.Kind != nil {
		n += 1 + sovFile(uint64(*m.Kind))
	}
	if m.ModTime != nil {
		l = len(m.ModTime)
		n += 1 + l + sovFile(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFile(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFile(x uint64) (n int) {
	return sovFile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *File) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Path = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSize", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FileSize = &v
			hasFields[0] |= uint64(0x00000002)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Kind = &v
			hasFields[0] |= uint64(0x00000004)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModTime", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModTime = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipFile(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
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
	if hasFields[0]&uint64(0x00000008) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFile(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFile
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowFile
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowFile
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFile
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFile
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipFile(data[start:])
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
	ErrInvalidLengthFile = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFile   = fmt.Errorf("proto: integer overflow")
)
