// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package evmv1alpha1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Params_2_list)(nil)

type _Params_2_list struct {
	list *[]int64
}

func (x *_Params_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Params_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfInt64((*x.list)[i])
}

func (x *_Params_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Params_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Int()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Params_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Params at list field ExtraEips as it is not of Message kind"))
}

func (x *_Params_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Params_2_list) NewElement() protoreflect.Value {
	v := int64(0)
	return protoreflect.ValueOfInt64(v)
}

func (x *_Params_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Params            protoreflect.MessageDescriptor
	fd_Params_evm_denom  protoreflect.FieldDescriptor
	fd_Params_extra_eips protoreflect.FieldDescriptor
)

func init() {
	file_polaris_evm_v1alpha1_params_proto_init()
	md_Params = File_polaris_evm_v1alpha1_params_proto.Messages().ByName("Params")
	fd_Params_evm_denom = md_Params.Fields().ByName("evm_denom")
	fd_Params_extra_eips = md_Params.Fields().ByName("extra_eips")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_polaris_evm_v1alpha1_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.EvmDenom != "" {
		value := protoreflect.ValueOfString(x.EvmDenom)
		if !f(fd_Params_evm_denom, value) {
			return
		}
	}
	if len(x.ExtraEips) != 0 {
		value := protoreflect.ValueOfList(&_Params_2_list{list: &x.ExtraEips})
		if !f(fd_Params_extra_eips, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "polaris.evm.v1alpha1.Params.evm_denom":
		return x.EvmDenom != ""
	case "polaris.evm.v1alpha1.Params.extra_eips":
		return len(x.ExtraEips) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: polaris.evm.v1alpha1.Params"))
		}
		panic(fmt.Errorf("message polaris.evm.v1alpha1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "polaris.evm.v1alpha1.Params.evm_denom":
		x.EvmDenom = ""
	case "polaris.evm.v1alpha1.Params.extra_eips":
		x.ExtraEips = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: polaris.evm.v1alpha1.Params"))
		}
		panic(fmt.Errorf("message polaris.evm.v1alpha1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "polaris.evm.v1alpha1.Params.evm_denom":
		value := x.EvmDenom
		return protoreflect.ValueOfString(value)
	case "polaris.evm.v1alpha1.Params.extra_eips":
		if len(x.ExtraEips) == 0 {
			return protoreflect.ValueOfList(&_Params_2_list{})
		}
		listValue := &_Params_2_list{list: &x.ExtraEips}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: polaris.evm.v1alpha1.Params"))
		}
		panic(fmt.Errorf("message polaris.evm.v1alpha1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "polaris.evm.v1alpha1.Params.evm_denom":
		x.EvmDenom = value.Interface().(string)
	case "polaris.evm.v1alpha1.Params.extra_eips":
		lv := value.List()
		clv := lv.(*_Params_2_list)
		x.ExtraEips = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: polaris.evm.v1alpha1.Params"))
		}
		panic(fmt.Errorf("message polaris.evm.v1alpha1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "polaris.evm.v1alpha1.Params.extra_eips":
		if x.ExtraEips == nil {
			x.ExtraEips = []int64{}
		}
		value := &_Params_2_list{list: &x.ExtraEips}
		return protoreflect.ValueOfList(value)
	case "polaris.evm.v1alpha1.Params.evm_denom":
		panic(fmt.Errorf("field evm_denom of message polaris.evm.v1alpha1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: polaris.evm.v1alpha1.Params"))
		}
		panic(fmt.Errorf("message polaris.evm.v1alpha1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "polaris.evm.v1alpha1.Params.evm_denom":
		return protoreflect.ValueOfString("")
	case "polaris.evm.v1alpha1.Params.extra_eips":
		list := []int64{}
		return protoreflect.ValueOfList(&_Params_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: polaris.evm.v1alpha1.Params"))
		}
		panic(fmt.Errorf("message polaris.evm.v1alpha1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in polaris.evm.v1alpha1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.EvmDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.ExtraEips) > 0 {
			l = 0
			for _, e := range x.ExtraEips {
				l += runtime.Sov(uint64(e))
			}
			n += 1 + runtime.Sov(uint64(l)) + l
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ExtraEips) > 0 {
			var pksize2 int
			for _, num := range x.ExtraEips {
				pksize2 += runtime.Sov(uint64(num))
			}
			i -= pksize2
			j1 := i
			for _, num1 := range x.ExtraEips {
				num := uint64(num1)
				for num >= 1<<7 {
					dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
					num >>= 7
					j1++
				}
				dAtA[j1] = uint8(num)
				j1++
			}
			i = runtime.EncodeVarint(dAtA, i, uint64(pksize2))
			i--
			dAtA[i] = 0x12
		}
		if len(x.EvmDenom) > 0 {
			i -= len(x.EvmDenom)
			copy(dAtA[i:], x.EvmDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.EvmDenom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EvmDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.EvmDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType == 0 {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					x.ExtraEips = append(x.ExtraEips, v)
				} else if wireType == 2 {
					var packedLen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
						}
						if iNdEx >= l {
							return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						packedLen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if packedLen < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					postIndex := iNdEx + packedLen
					if postIndex < 0 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
					}
					if postIndex > l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					var elementCount int
					var count int
					for _, integer := range dAtA[iNdEx:postIndex] {
						if integer < 128 {
							count++
						}
					}
					elementCount = count
					if elementCount != 0 && len(x.ExtraEips) == 0 {
						x.ExtraEips = make([]int64, 0, elementCount)
					}
					for iNdEx < postIndex {
						var v int64
						for shift := uint(0); ; shift += 7 {
							if shift >= 64 {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
							}
							if iNdEx >= l {
								return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
							}
							b := dAtA[iNdEx]
							iNdEx++
							v |= int64(b&0x7F) << shift
							if b < 0x80 {
								break
							}
						}
						x.ExtraEips = append(x.ExtraEips, v)
					}
				} else {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ExtraEips", wireType)
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: polaris/evm/v1alpha1/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `Params` defines the parameters for the x/evm module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `evm_denom` represents the token denomination used as the native token
	// within the EVM.
	EvmDenom string `protobuf:"bytes,1,opt,name=evm_denom,json=evmDenom,proto3" json:"evm_denom,omitempty"`
	// `extra_eips` defines a list of additional EIPs for the vm.Config
	ExtraEips []int64 `protobuf:"varint,2,rep,packed,name=extra_eips,json=extraEips,proto3" json:"extra_eips,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polaris_evm_v1alpha1_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_polaris_evm_v1alpha1_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetEvmDenom() string {
	if x != nil {
		return x.EvmDenom
	}
	return ""
}

func (x *Params) GetExtraEips() []int64 {
	if x != nil {
		return x.ExtraEips
	}
	return nil
}

var File_polaris_evm_v1alpha1_params_proto protoreflect.FileDescriptor

var file_polaris_evm_v1alpha1_params_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x65, 0x76, 0x6d, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7e, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x65, 0x76, 0x6d,
	0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xf2, 0xde,
	0x1f, 0x10, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x65, 0x76, 0x6d, 0x5f, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x22, 0x52, 0x08, 0x65, 0x76, 0x6d, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x41, 0x0a, 0x0a,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x65, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03,
	0x42, 0x22, 0xe2, 0xde, 0x1f, 0x09, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x49, 0x50, 0x73, 0xf2,
	0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x65,
	0x69, 0x70, 0x73, 0x22, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x45, 0x69, 0x70, 0x73, 0x42,
	0xcc, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e,
	0x65, 0x76, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x65, 0x76, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x65, 0x76, 0x6d, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x45, 0x58, 0xaa, 0x02, 0x14, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x45,
	0x76, 0x6d, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x50, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x5c, 0x45, 0x76, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x20, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x5c, 0x45, 0x76, 0x6d,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x3a,
	0x3a, 0x45, 0x76, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_polaris_evm_v1alpha1_params_proto_rawDescOnce sync.Once
	file_polaris_evm_v1alpha1_params_proto_rawDescData = file_polaris_evm_v1alpha1_params_proto_rawDesc
)

func file_polaris_evm_v1alpha1_params_proto_rawDescGZIP() []byte {
	file_polaris_evm_v1alpha1_params_proto_rawDescOnce.Do(func() {
		file_polaris_evm_v1alpha1_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_polaris_evm_v1alpha1_params_proto_rawDescData)
	})
	return file_polaris_evm_v1alpha1_params_proto_rawDescData
}

var file_polaris_evm_v1alpha1_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_polaris_evm_v1alpha1_params_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: polaris.evm.v1alpha1.Params
}
var file_polaris_evm_v1alpha1_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_polaris_evm_v1alpha1_params_proto_init() }
func file_polaris_evm_v1alpha1_params_proto_init() {
	if File_polaris_evm_v1alpha1_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_polaris_evm_v1alpha1_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_polaris_evm_v1alpha1_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_polaris_evm_v1alpha1_params_proto_goTypes,
		DependencyIndexes: file_polaris_evm_v1alpha1_params_proto_depIdxs,
		MessageInfos:      file_polaris_evm_v1alpha1_params_proto_msgTypes,
	}.Build()
	File_polaris_evm_v1alpha1_params_proto = out.File
	file_polaris_evm_v1alpha1_params_proto_rawDesc = nil
	file_polaris_evm_v1alpha1_params_proto_goTypes = nil
	file_polaris_evm_v1alpha1_params_proto_depIdxs = nil
}
