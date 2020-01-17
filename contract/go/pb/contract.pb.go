// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract/pb/contract.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ArgPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArgPair) Reset()         { *m = ArgPair{} }
func (m *ArgPair) String() string { return proto.CompactTextString(m) }
func (*ArgPair) ProtoMessage()    {}
func (*ArgPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{0}
}

func (m *ArgPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArgPair.Unmarshal(m, b)
}
func (m *ArgPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArgPair.Marshal(b, m, deterministic)
}
func (m *ArgPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArgPair.Merge(m, src)
}
func (m *ArgPair) XXX_Size() int {
	return xxx_messageInfo_ArgPair.Size(m)
}
func (m *ArgPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ArgPair.DiscardUnknown(m)
}

var xxx_messageInfo_ArgPair proto.InternalMessageInfo

func (m *ArgPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ArgPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CallArgs struct {
	Method               string     `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Args                 []*ArgPair `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Caller               string     `protobuf:"bytes,3,opt,name=caller,proto3" json:"caller,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CallArgs) Reset()         { *m = CallArgs{} }
func (m *CallArgs) String() string { return proto.CompactTextString(m) }
func (*CallArgs) ProtoMessage()    {}
func (*CallArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{1}
}

func (m *CallArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallArgs.Unmarshal(m, b)
}
func (m *CallArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallArgs.Marshal(b, m, deterministic)
}
func (m *CallArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallArgs.Merge(m, src)
}
func (m *CallArgs) XXX_Size() int {
	return xxx_messageInfo_CallArgs.Size(m)
}
func (m *CallArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_CallArgs.DiscardUnknown(m)
}

var xxx_messageInfo_CallArgs proto.InternalMessageInfo

func (m *CallArgs) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CallArgs) GetArgs() []*ArgPair {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *CallArgs) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

type SyscallHeader struct {
	Ctxid                int64    `protobuf:"varint,1,opt,name=ctxid,proto3" json:"ctxid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyscallHeader) Reset()         { *m = SyscallHeader{} }
func (m *SyscallHeader) String() string { return proto.CompactTextString(m) }
func (*SyscallHeader) ProtoMessage()    {}
func (*SyscallHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{2}
}

func (m *SyscallHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyscallHeader.Unmarshal(m, b)
}
func (m *SyscallHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyscallHeader.Marshal(b, m, deterministic)
}
func (m *SyscallHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyscallHeader.Merge(m, src)
}
func (m *SyscallHeader) XXX_Size() int {
	return xxx_messageInfo_SyscallHeader.Size(m)
}
func (m *SyscallHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SyscallHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SyscallHeader proto.InternalMessageInfo

func (m *SyscallHeader) GetCtxid() int64 {
	if m != nil {
		return m.Ctxid
	}
	return 0
}

type PutRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Key                  []byte         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte         `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{3}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PutRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{4}
}

func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (m *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(m, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

type GetRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Key                  []byte         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type DeleteRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Key                  []byte         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type TransferRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	From                 string         `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string         `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount               string         `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{9}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransferRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransferRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransferRequest) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type TransferResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{10}
}

func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (m *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(m, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

type ContractCallRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Module               string         `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	Contract             string         `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	Method               string         `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Args                 []*ArgPair     `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ContractCallRequest) Reset()         { *m = ContractCallRequest{} }
func (m *ContractCallRequest) String() string { return proto.CompactTextString(m) }
func (*ContractCallRequest) ProtoMessage()    {}
func (*ContractCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{11}
}

func (m *ContractCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCallRequest.Unmarshal(m, b)
}
func (m *ContractCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCallRequest.Marshal(b, m, deterministic)
}
func (m *ContractCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallRequest.Merge(m, src)
}
func (m *ContractCallRequest) XXX_Size() int {
	return xxx_messageInfo_ContractCallRequest.Size(m)
}
func (m *ContractCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallRequest proto.InternalMessageInfo

func (m *ContractCallRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractCallRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ContractCallRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ContractCallRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ContractCallRequest) GetArgs() []*ArgPair {
	if m != nil {
		return m.Args
	}
	return nil
}

type ContractCallResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContractCallResponse) Reset()         { *m = ContractCallResponse{} }
func (m *ContractCallResponse) String() string { return proto.CompactTextString(m) }
func (*ContractCallResponse) ProtoMessage()    {}
func (*ContractCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{12}
}

func (m *ContractCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCallResponse.Unmarshal(m, b)
}
func (m *ContractCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCallResponse.Marshal(b, m, deterministic)
}
func (m *ContractCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallResponse.Merge(m, src)
}
func (m *ContractCallResponse) XXX_Size() int {
	return xxx_messageInfo_ContractCallResponse.Size(m)
}
func (m *ContractCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallResponse proto.InternalMessageInfo

func (m *ContractCallResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type Response struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{13}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SetOutputRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Response             *Response      `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetOutputRequest) Reset()         { *m = SetOutputRequest{} }
func (m *SetOutputRequest) String() string { return proto.CompactTextString(m) }
func (*SetOutputRequest) ProtoMessage()    {}
func (*SetOutputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{14}
}

func (m *SetOutputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOutputRequest.Unmarshal(m, b)
}
func (m *SetOutputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOutputRequest.Marshal(b, m, deterministic)
}
func (m *SetOutputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOutputRequest.Merge(m, src)
}
func (m *SetOutputRequest) XXX_Size() int {
	return xxx_messageInfo_SetOutputRequest.Size(m)
}
func (m *SetOutputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOutputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetOutputRequest proto.InternalMessageInfo

func (m *SetOutputRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SetOutputRequest) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type SetOutputResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOutputResponse) Reset()         { *m = SetOutputResponse{} }
func (m *SetOutputResponse) String() string { return proto.CompactTextString(m) }
func (*SetOutputResponse) ProtoMessage()    {}
func (*SetOutputResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{15}
}

func (m *SetOutputResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOutputResponse.Unmarshal(m, b)
}
func (m *SetOutputResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOutputResponse.Marshal(b, m, deterministic)
}
func (m *SetOutputResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOutputResponse.Merge(m, src)
}
func (m *SetOutputResponse) XXX_Size() int {
	return xxx_messageInfo_SetOutputResponse.Size(m)
}
func (m *SetOutputResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOutputResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetOutputResponse proto.InternalMessageInfo

type GetCallArgsRequest struct {
	Header               *SyscallHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetCallArgsRequest) Reset()         { *m = GetCallArgsRequest{} }
func (m *GetCallArgsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCallArgsRequest) ProtoMessage()    {}
func (*GetCallArgsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea6d8c13449a4cc, []int{16}
}

func (m *GetCallArgsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCallArgsRequest.Unmarshal(m, b)
}
func (m *GetCallArgsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCallArgsRequest.Marshal(b, m, deterministic)
}
func (m *GetCallArgsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCallArgsRequest.Merge(m, src)
}
func (m *GetCallArgsRequest) XXX_Size() int {
	return xxx_messageInfo_GetCallArgsRequest.Size(m)
}
func (m *GetCallArgsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCallArgsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCallArgsRequest proto.InternalMessageInfo

func (m *GetCallArgsRequest) GetHeader() *SyscallHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ArgPair)(nil), "contract.ArgPair")
	proto.RegisterType((*CallArgs)(nil), "contract.CallArgs")
	proto.RegisterType((*SyscallHeader)(nil), "contract.SyscallHeader")
	proto.RegisterType((*PutRequest)(nil), "contract.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "contract.PutResponse")
	proto.RegisterType((*GetRequest)(nil), "contract.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "contract.GetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "contract.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "contract.DeleteResponse")
	proto.RegisterType((*TransferRequest)(nil), "contract.TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "contract.TransferResponse")
	proto.RegisterType((*ContractCallRequest)(nil), "contract.ContractCallRequest")
	proto.RegisterType((*ContractCallResponse)(nil), "contract.ContractCallResponse")
	proto.RegisterType((*Response)(nil), "contract.Response")
	proto.RegisterType((*SetOutputRequest)(nil), "contract.SetOutputRequest")
	proto.RegisterType((*SetOutputResponse)(nil), "contract.SetOutputResponse")
	proto.RegisterType((*GetCallArgsRequest)(nil), "contract.GetCallArgsRequest")
}

func init() { proto.RegisterFile("contract/pb/contract.proto", fileDescriptor_dea6d8c13449a4cc) }

var fileDescriptor_dea6d8c13449a4cc = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0xec, 0x24, 0x4d, 0x27, 0xa4, 0xa4, 0xdb, 0xaa, 0x58, 0x3d, 0x45, 0x8b, 0x2a, 0xe5,
	0x80, 0x62, 0x28, 0x5f, 0xd0, 0xa6, 0xd0, 0xde, 0x1a, 0x6d, 0x39, 0x71, 0x5b, 0xdb, 0x53, 0x27,
	0xc2, 0xce, 0x9a, 0xdd, 0x75, 0x21, 0x17, 0xfe, 0x8b, 0xbf, 0x43, 0xde, 0x5d, 0x3b, 0x8e, 0x84,
	0x10, 0x4a, 0x7b, 0x9b, 0x67, 0xcf, 0xcc, 0x7b, 0x4f, 0x33, 0x3b, 0x70, 0x1e, 0x8b, 0xb5, 0x96,
	0x3c, 0xd6, 0x61, 0x11, 0x85, 0x75, 0x3c, 0x2b, 0xa4, 0xd0, 0x82, 0x0c, 0x6a, 0x4c, 0x3f, 0xc0,
	0xc1, 0x95, 0x4c, 0x17, 0x7c, 0x25, 0xc9, 0x18, 0xfc, 0x6f, 0xb8, 0x09, 0x3a, 0x93, 0xce, 0xf4,
	0x90, 0x55, 0x21, 0x39, 0x85, 0xde, 0x13, 0xcf, 0x4a, 0x0c, 0xbc, 0x49, 0x67, 0xfa, 0x8a, 0x59,
	0x40, 0x39, 0x0c, 0xe6, 0x3c, 0xcb, 0xae, 0x64, 0xaa, 0xc8, 0x19, 0xf4, 0x73, 0xd4, 0x4b, 0x91,
	0xb8, 0x32, 0x87, 0xc8, 0x05, 0x74, 0xb9, 0x4c, 0x55, 0xe0, 0x4d, 0xfc, 0xe9, 0xf0, 0xf2, 0x78,
	0xd6, 0xf0, 0x3b, 0x32, 0x66, 0x7e, 0x57, 0xe5, 0x31, 0xcf, 0x32, 0x94, 0x81, 0x6f, 0xcb, 0x2d,
	0xa2, 0x17, 0x30, 0x7a, 0xd8, 0xa8, 0x0a, 0xdc, 0x21, 0x4f, 0x50, 0x56, 0x4a, 0x62, 0xfd, 0x73,
	0x65, 0x69, 0x7c, 0x66, 0x01, 0x45, 0x80, 0x45, 0xa9, 0x19, 0x7e, 0x2f, 0x51, 0x69, 0x12, 0x42,
	0x7f, 0x69, 0xb2, 0x4d, 0xd2, 0xf0, 0xf2, 0xcd, 0x96, 0x75, 0xa7, 0x19, 0x73, 0x69, 0xb5, 0x61,
	0x6b, 0x6e, 0xd7, 0xb0, 0xdf, 0x36, 0x3c, 0x82, 0xa1, 0xa1, 0x51, 0x85, 0x58, 0x2b, 0xa4, 0xf7,
	0x00, 0xb7, 0xf8, 0x82, 0xac, 0xf4, 0x2d, 0x0c, 0x4d, 0x43, 0xdb, 0x7f, 0x2b, 0xa2, 0xd3, 0x16,
	0xc1, 0x60, 0x74, 0x83, 0x19, 0x6a, 0x7c, 0x41, 0xe2, 0x31, 0x1c, 0xd5, 0x3d, 0x9d, 0xb7, 0x5f,
	0xf0, 0xfa, 0x8b, 0xe4, 0x6b, 0xf5, 0x88, 0x72, 0x6f, 0x1e, 0x02, 0xdd, 0x47, 0x29, 0x72, 0x43,
	0x74, 0xc8, 0x4c, 0x4c, 0x8e, 0xc0, 0xd3, 0xc2, 0x0d, 0xd9, 0xd3, 0xa2, 0x1a, 0x3c, 0xcf, 0x45,
	0xb9, 0xd6, 0x41, 0xd7, 0x0e, 0xde, 0x22, 0x4a, 0x60, 0xbc, 0xe5, 0x77, 0x9a, 0x7e, 0x77, 0xe0,
	0x64, 0xee, 0x28, 0xab, 0xc5, 0xdb, 0x5b, 0x58, 0xb5, 0xac, 0x22, 0x29, 0x33, 0x74, 0xd2, 0x1c,
	0x22, 0xe7, 0xd0, 0xbc, 0x07, 0x27, 0xb1, 0xc1, 0xad, 0x05, 0xef, 0xfe, 0x75, 0xc1, 0x7b, 0xff,
	0x5c, 0x70, 0xfa, 0x19, 0x4e, 0x77, 0xa5, 0xbb, 0x19, 0xcf, 0x60, 0x20, 0x5d, 0xec, 0xd4, 0x93,
	0x6d, 0x8b, 0x3a, 0x8b, 0x35, 0x39, 0x74, 0x01, 0x83, 0xa6, 0xf6, 0x0c, 0xfa, 0x4a, 0x73, 0x5d,
	0x2a, 0x53, 0xd9, 0x63, 0x0e, 0x91, 0x00, 0x0e, 0x72, 0x54, 0x8a, 0xa7, 0xb5, 0xbf, 0x1a, 0x56,
	0x13, 0x89, 0x44, 0xb2, 0x71, 0x5b, 0x6d, 0x62, 0xaa, 0x60, 0xfc, 0x80, 0xfa, 0xbe, 0xd4, 0xc5,
	0x33, 0x5e, 0x50, 0xdb, 0x86, 0xf7, 0x1f, 0x36, 0x4e, 0xe0, 0xb8, 0x45, 0xea, 0x3e, 0x7e, 0x02,
	0x72, 0x8b, 0xba, 0x3e, 0x29, 0xfb, 0x6a, 0xb9, 0x7e, 0x7f, 0xe7, 0x7f, 0x7d, 0x97, 0xae, 0xf4,
	0xb2, 0x8c, 0x66, 0xb1, 0xc8, 0xc3, 0x6b, 0xbc, 0x91, 0xc8, 0xf3, 0xb9, 0x48, 0x50, 0x86, 0xe5,
	0x0f, 0xfe, 0x94, 0x37, 0x57, 0x30, 0x4c, 0x45, 0x58, 0x44, 0x51, 0xdf, 0x1c, 0xc3, 0x8f, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x78, 0x13, 0xd7, 0xab, 0x2a, 0x05, 0x00, 0x00,
}
