// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/http_status.proto

package envoy_type

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// HTTP response codes supported in Envoy.
// For more details: http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int32

const (
	// Empty - This code not part of the HTTP status code specification, but it is needed for proto
	// `enum` type.
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}

var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}

// HTTP status.
type HttpStatus struct {
	// Supplies HTTP response code.
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.type.StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpStatus) Reset()         { *m = HttpStatus{} }
func (m *HttpStatus) String() string { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()    {}
func (*HttpStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}
func (m *HttpStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStatus.Merge(m, src)
}
func (m *HttpStatus) XXX_Size() int {
	return m.Size()
}
func (m *HttpStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStatus proto.InternalMessageInfo

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterEnum("envoy.type.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.HttpStatus")
}

func init() { proto.RegisterFile("envoy/type/http_status.proto", fileDescriptor_7544d7adacd3389b) }

var fileDescriptor_7544d7adacd3389b = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x6f, 0x1c, 0x45,
	0x14, 0x4e, 0x4f, 0xc5, 0x49, 0x5c, 0x71, 0x9c, 0x97, 0x8a, 0x13, 0x87, 0x10, 0x2c, 0x2b, 0x27,
	0xc4, 0xc1, 0x96, 0xe0, 0x0f, 0x60, 0x3b, 0x76, 0x6c, 0xe1, 0x99, 0x8c, 0xc6, 0x33, 0xb9, 0xa2,
	0x72, 0xd7, 0x9b, 0x99, 0x52, 0x7a, 0xea, 0x75, 0xaa, 0xdf, 0x8c, 0xdd, 0x1c, 0xf9, 0x05, 0xec,
	0xfb, 0x7a, 0x60, 0x11, 0x4a, 0x08, 0x08, 0xb8, 0x70, 0xe2, 0x18, 0xf6, 0xfc, 0x04, 0xe4, 0x1b,
	0x77, 0x76, 0x10, 0xa0, 0xaa, 0x59, 0x92, 0x4b, 0x6e, 0x55, 0xaf, 0xde, 0xf2, 0x2d, 0xaf, 0x5b,
	0x5e, 0x40, 0x37, 0xa0, 0x72, 0x99, 0xcb, 0x1c, 0x97, 0xbb, 0xcc, 0xf9, 0x93, 0x05, 0x6b, 0xee,
	0x17, 0x4b, 0xb9, 0x27, 0x26, 0x25, 0xe3, 0xeb, 0x52, 0x78, 0x3d, 0x3f, 0x3f, 0xd0, 0x99, 0x35,
	0x9a, 0x71, 0x79, 0x7c, 0x18, 0x26, 0x5d, 0xac, 0x49, 0xb9, 0xc9, 0x9c, 0xef, 0xc4, 0x42, 0xf5,
	0xb8, 0x3c, 0x9c, 0x92, 0xc1, 0x73, 0xc9, 0x62, 0xf2, 0xf0, 0xec, 0xa3, 0x67, 0x97, 0xee, 0x76,
	0x58, 0x1a, 0x66, 0xac, 0x91, 0xc1, 0xd5, 0xb9, 0x2f, 0x7f, 0xfe, 0x4a, 0x4c, 0x3d, 0x9d, 0x54,
	0x16, 0x0f, 0x8d, 0x4f, 0x90, 0x34, 0x62, 0xe5, 0x23, 0x5f, 0x4c, 0x4b, 0x79, 0x37, 0x55, 0x4d,
	0xcb, 0xa9, 0xf5, 0x5e, 0xce, 0x25, 0x1c, 0x52, 0x33, 0xf2, 0xd8, 0x1a, 0x39, 0xb6, 0xae, 0x8f,
	0x60, 0xd4, 0x51, 0x59, 0xb9, 0xf2, 0x04, 0xdc, 0x4e, 0xd4, 0x8c, 0x3c, 0xba, 0xe6, 0x51, 0x33,
	0x1a, 0xf8, 0x3a, 0x51, 0x27, 0xe4, 0xb1, 0x95, 0x34, 0xc5, 0x3c, 0x5c, 0xbf, 0x49, 0xd4, 0xa2,
	0x7c, 0xb0, 0x46, 0x6e, 0xa5, 0xcf, 0x5d, 0xf2, 0x96, 0x35, 0xdb, 0x01, 0x6e, 0xb9, 0x36, 0xf9,
	0x9e, 0x66, 0x4b, 0x0e, 0xbe, 0x4d, 0xd4, 0xac, 0x9c, 0xae, 0x51, 0xe8, 0x8b, 0x8e, 0xe1, 0xbb,
	0x44, 0x9d, 0x92, 0x33, 0x0d, 0x2c, 0x90, 0xc7, 0xa1, 0xef, 0x13, 0x75, 0x5a, 0xce, 0xd6, 0xb5,
	0x67, 0xab, 0xb3, 0x71, 0xf0, 0x87, 0x44, 0x81, 0x3c, 0x5e, 0xed, 0x67, 0x6c, 0x87, 0x58, 0xe1,
	0xc7, 0x44, 0xcd, 0xc9, 0x93, 0x2b, 0x99, 0x47, 0x6d, 0xca, 0x06, 0xe6, 0xe4, 0x03, 0x82, 0x3b,
	0x89, 0x3a, 0x2e, 0x8f, 0x6c, 0x55, 0x5b, 0x05, 0x1a, 0x38, 0x88, 0x29, 0xb1, 0x28, 0xcf, 0x70,
	0xad, 0x4b, 0x36, 0xc5, 0x02, 0x6e, 0x54, 0xd4, 0x19, 0x09, 0x55, 0x1a, 0xa0, 0xa9, 0xa3, 0xef,
	0x69, 0x87, 0x8e, 0xb3, 0x12, 0x6e, 0x56, 0x94, 0x94, 0x53, 0x1b, 0xd4, 0x77, 0x06, 0x3e, 0xae,
	0x04, 0x5a, 0x3b, 0x88, 0x57, 0xb8, 0x8b, 0x1e, 0x6e, 0x55, 0xc2, 0xf0, 0x1a, 0x71, 0x95, 0x8c,
	0x6d, 0x5b, 0x34, 0xf0, 0x49, 0x4c, 0x68, 0x15, 0x58, 0xf7, 0xb4, 0x5f, 0xc2, 0xa7, 0x15, 0x75,
	0x56, 0x9e, 0x6a, 0x62, 0x2f, 0x27, 0xaf, 0x7d, 0xd9, 0x40, 0x63, 0x3d, 0xa6, 0x0c, 0x9f, 0xc5,
	0xf8, 0x64, 0xca, 0x24, 0xfe, 0x79, 0x45, 0x9d, 0x94, 0x72, 0x55, 0x9b, 0x06, 0x5e, 0xef, 0x63,
	0xc1, 0xf0, 0x8c, 0x08, 0x32, 0xb4, 0x9c, 0x1e, 0xea, 0xf6, 0x14, 0x1a, 0x78, 0x56, 0x04, 0xf0,
	0x75, 0x5d, 0xf6, 0x62, 0xe5, 0xf5, 0xbe, 0xf5, 0x68, 0xe0, 0x39, 0x11, 0xf4, 0xdb, 0x20, 0xbf,
	0x6b, 0x8d, 0x41, 0x07, 0xcf, 0x8b, 0x00, 0xa4, 0x46, 0x3c, 0x04, 0xfe, 0x82, 0x88, 0xdc, 0x90,
	0xbb, 0x64, 0x6a, 0xc4, 0x2b, 0x59, 0x46, 0x7b, 0x68, 0xe0, 0x45, 0xa1, 0x94, 0x3c, 0x11, 0x02,
	0xd1, 0x29, 0xbd, 0x9b, 0x21, 0xbc, 0x24, 0x82, 0x57, 0x11, 0x7f, 0x70, 0x0b, 0x1d, 0xdb, 0x34,
	0x7a, 0x34, 0x99, 0xf5, 0xb2, 0x08, 0x46, 0x8c, 0x20, 0x36, 0x6d, 0x0f, 0xa9, 0xcf, 0xf0, 0x4a,
	0x1c, 0xb8, 0x46, 0xae, 0x9d, 0xd9, 0x94, 0xe1, 0x55, 0xa1, 0xa6, 0xe5, 0xe1, 0xcb, 0xe4, 0x10,
	0x5e, 0x8b, 0xe9, 0xdb, 0xe8, 0x3a, 0xdc, 0x9d, 0xf4, 0x78, 0x5d, 0xa8, 0x79, 0xa9, 0xea, 0x1e,
	0x53, 0x72, 0xc6, 0x86, 0xf6, 0x1b, 0xda, 0x66, 0x68, 0xe0, 0x8d, 0x31, 0xbd, 0x8c, 0xb4, 0x69,
	0x12, 0x6d, 0x6b, 0xdf, 0x41, 0x78, 0x53, 0x04, 0x61, 0x5a, 0x8d, 0xad, 0x10, 0x21, 0xd7, 0x81,
	0xb7, 0x84, 0x7a, 0x40, 0xce, 0xb5, 0x5c, 0xd1, 0xcf, 0x87, 0x0e, 0x57, 0xd1, 0x58, 0xdd, 0x2c,
	0x73, 0x84, 0xb7, 0x85, 0x3a, 0x27, 0x4f, 0x37, 0xb4, 0xeb, 0x60, 0x8d, 0x78, 0x47, 0xb3, 0x2d,
	0xda, 0x36, 0x52, 0x7b, 0x47, 0x04, 0xd9, 0xd7, 0xf7, 0x73, 0x4c, 0x59, 0xdf, 0x33, 0xf3, 0xdd,
	0x08, 0xa6, 0x6a, 0x8b, 0xa1, 0x0d, 0x38, 0x91, 0xff, 0xbd, 0xd8, 0xaa, 0xe5, 0x72, 0x4f, 0x29,
	0x16, 0x45, 0x68, 0xb2, 0xee, 0xd8, 0x72, 0x09, 0xef, 0x8b, 0xb0, 0x4f, 0xdb, 0x94, 0x5e, 0x43,
	0x03, 0x1f, 0x44, 0x75, 0x87, 0xcd, 0x2e, 0x61, 0x8e, 0xce, 0xa0, 0x4b, 0x4b, 0xf8, 0x30, 0x52,
	0x69, 0xe5, 0x1d, 0xaf, 0x0d, 0x4e, 0x98, 0x7f, 0x14, 0x91, 0xdf, 0xcb, 0x7c, 0xf2, 0x74, 0x23,
	0x16, 0x34, 0x89, 0xaa, 0xda, 0x95, 0x23, 0x0c, 0x05, 0xdc, 0x8c, 0x86, 0x8c, 0xae, 0x9b, 0xa8,
	0x0d, 0xfa, 0x0d, 0x8b, 0x99, 0x29, 0x26, 0xea, 0xdc, 0x8a, 0x30, 0xb7, 0x1c, 0xa3, 0x77, 0x3a,
	0xdb, 0x41, 0x3f, 0x40, 0xbf, 0xee, 0x3d, 0x79, 0xf8, 0x25, 0x6a, 0x5f, 0x23, 0xde, 0xea, 0xe5,
	0x19, 0x86, 0x8d, 0x41, 0x03, 0xbf, 0x8a, 0xd1, 0x96, 0x5d, 0xd6, 0x8c, 0x7b, 0xba, 0x84, 0xdf,
	0x22, 0xff, 0x50, 0x67, 0x53, 0x6c, 0x39, 0x3d, 0xd0, 0x36, 0x8b, 0x82, 0xfd, 0x1e, 0xcb, 0x47,
	0x69, 0x63, 0xa7, 0xff, 0x10, 0xea, 0x82, 0x9c, 0xdf, 0x6c, 0x36, 0xeb, 0x57, 0xd1, 0x17, 0x96,
	0x5c, 0x50, 0x79, 0x6c, 0x03, 0xfc, 0x29, 0xd4, 0x79, 0x79, 0xe6, 0xaa, 0xf6, 0x56, 0x3b, 0x5e,
	0xc9, 0x0a, 0xaa, 0x61, 0x87, 0xd8, 0x6a, 0xc6, 0x02, 0xfe, 0x1a, 0xe1, 0x2c, 0xfa, 0xed, 0xb6,
	0x4d, 0x2d, 0x3a, 0xde, 0x61, 0xf2, 0xba, 0x83, 0xf0, 0x77, 0xdc, 0xf3, 0x6d, 0xa2, 0xfc, 0x12,
	0x72, 0xb4, 0x00, 0xfe, 0x11, 0xa3, 0x8f, 0x6b, 0x7d, 0x9f, 0x83, 0xa2, 0x06, 0xfe, 0x15, 0xea,
	0xa2, 0x7c, 0xa8, 0x86, 0xbc, 0x47, 0xfe, 0xda, 0x7d, 0x76, 0xf3, 0x3f, 0xb1, 0x3a, 0x73, 0xfb,
	0x60, 0x21, 0xb9, 0x73, 0xb0, 0x90, 0xfc, 0x74, 0xb0, 0x90, 0xec, 0x1e, 0x89, 0x3f, 0xc7, 0xc7,
	0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xfd, 0x12, 0x85, 0x61, 0x05, 0x00, 0x00,
}

func (m *HttpStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHttpStatus(dAtA, i, uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHttpStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HttpStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovHttpStatus(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttpStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttpStatus(x uint64) (n int) {
	return sovHttpStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpStatus
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
			return fmt.Errorf("proto: HttpStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= StatusCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHttpStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpStatus
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
func skipHttpStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpStatus
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
					return 0, ErrIntOverflowHttpStatus
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
					return 0, ErrIntOverflowHttpStatus
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
				return 0, ErrInvalidLengthHttpStatus
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHttpStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpStatus
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
				next, err := skipHttpStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHttpStatus
				}
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
	ErrInvalidLengthHttpStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpStatus   = fmt.Errorf("proto: integer overflow")
)
