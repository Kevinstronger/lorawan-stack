// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/join.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf3 "github.com/gogo/protobuf/types"

import github_com_TheThingsNetwork_ttn_pkg_types "github.com/TheThingsNetwork/ttn/pkg/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Join Request
type JoinReq struct {
	// Raw uplink bytes (PHYPayload)
	RawPayload []byte `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	// LoRaWAN Payload
	Payload Message `protobuf:"bytes,2,opt,name=payload" json:"payload"`
	// End device identifiers
	// - this includes the DevAddr assigned by the Network Server
	EndDeviceIdentifiers `protobuf:"bytes,3,opt,name=end_device,json=endDevice,embedded=end_device" json:"end_device"`
	// MAC version selected by the NS
	SelectedMacVersion string `protobuf:"bytes,4,opt,name=selected_mac_version,json=selectedMacVersion,proto3" json:"selected_mac_version,omitempty"`
	// NetID of the accepting network server
	NetID github_com_TheThingsNetwork_ttn_pkg_types.NetID `protobuf:"bytes,5,opt,name=net_id,json=netId,proto3,customtype=github.com/TheThingsNetwork/ttn/pkg/types.NetID" json:"net_id"`
	// Downlink Settings
	DownlinkSettings DLSettings `protobuf:"bytes,6,opt,name=downlink_settings,json=downlinkSettings" json:"downlink_settings"`
	// RX Delay in seconds
	RxDelay uint32 `protobuf:"varint,7,opt,name=rx_delay,json=rxDelay,proto3" json:"rx_delay,omitempty"`
	// Optional CFList
	CFList *CFList `protobuf:"bytes,8,opt,name=cf_list,json=cfList" json:"cf_list,omitempty"`
}

func (m *JoinReq) Reset()                    { *m = JoinReq{} }
func (*JoinReq) ProtoMessage()               {}
func (*JoinReq) Descriptor() ([]byte, []int) { return fileDescriptorJoin, []int{0} }

func (m *JoinReq) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinReq) GetPayload() Message {
	if m != nil {
		return m.Payload
	}
	return Message{}
}

func (m *JoinReq) GetSelectedMacVersion() string {
	if m != nil {
		return m.SelectedMacVersion
	}
	return ""
}

func (m *JoinReq) GetDownlinkSettings() DLSettings {
	if m != nil {
		return m.DownlinkSettings
	}
	return DLSettings{}
}

func (m *JoinReq) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *JoinReq) GetCFList() *CFList {
	if m != nil {
		return m.CFList
	}
	return nil
}

// Answer to the Join Request
type JoinAns struct {
	// Raw uplink bytes (PHYPayload)
	RawPayload []byte `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	// The session keys
	SessionKeys `protobuf:"bytes,2,opt,name=session_keys,json=sessionKeys,embedded=session_keys" json:"session_keys"`
	// Lifetime of the session
	Lifetime *google_protobuf3.Duration `protobuf:"bytes,3,opt,name=lifetime" json:"lifetime,omitempty"`
}

func (m *JoinAns) Reset()                    { *m = JoinAns{} }
func (*JoinAns) ProtoMessage()               {}
func (*JoinAns) Descriptor() ([]byte, []int) { return fileDescriptorJoin, []int{1} }

func (m *JoinAns) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinAns) GetLifetime() *google_protobuf3.Duration {
	if m != nil {
		return m.Lifetime
	}
	return nil
}

// Session keys for a LoRaWAN session.
// Only the components for which the keys were meant, will have the key-encryption-key (KEK) to decrypt the individual keys
type SessionKeys struct {
	// Join Server issued identifier for the session keys
	SessionKeyId string `protobuf:"bytes,1,opt,name=session_key_id,json=sessionKeyId,proto3" json:"session_key_id,omitempty"`
	// The (encrypted) Forwarding Network Session Integrity Key (or Network Session Key in 1.0 compatibility mode)
	FNwkSIntKey *KeyEnvelope `protobuf:"bytes,2,opt,name=f_nwk_s_int_key,json=fNwkSIntKey" json:"f_nwk_s_int_key,omitempty"`
	// The (encrypted) Serving Network Session Integrity Key
	SNwkSIntKey *KeyEnvelope `protobuf:"bytes,4,opt,name=s_nwk_s_int_key,json=sNwkSIntKey" json:"s_nwk_s_int_key,omitempty"`
	// The (encrypted) Network Session Encryption Key
	NwkSEncKey *KeyEnvelope `protobuf:"bytes,6,opt,name=nwk_s_enc_key,json=nwkSEncKey" json:"nwk_s_enc_key,omitempty"`
	// The (encrypted) Application Session Key
	AppSKey *KeyEnvelope `protobuf:"bytes,8,opt,name=app_s_key,json=appSKey" json:"app_s_key,omitempty"`
}

func (m *SessionKeys) Reset()                    { *m = SessionKeys{} }
func (*SessionKeys) ProtoMessage()               {}
func (*SessionKeys) Descriptor() ([]byte, []int) { return fileDescriptorJoin, []int{2} }

func (m *SessionKeys) GetSessionKeyId() string {
	if m != nil {
		return m.SessionKeyId
	}
	return ""
}

func (m *SessionKeys) GetFNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.FNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetSNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.SNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetNwkSEncKey() *KeyEnvelope {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func (m *SessionKeys) GetAppSKey() *KeyEnvelope {
	if m != nil {
		return m.AppSKey
	}
	return nil
}

type KeyEnvelope struct {
	// The (encrypted) key
	Key *github_com_TheThingsNetwork_ttn_pkg_types.AES128Key `protobuf:"bytes,1,opt,name=key,proto3,customtype=github.com/TheThingsNetwork/ttn/pkg/types.AES128Key" json:"key,omitempty"`
	// The label of the RFC 3394 key-encryption-key (KEK) that was used to encrypt the key
	KekLabel string `protobuf:"bytes,2,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
}

func (m *KeyEnvelope) Reset()                    { *m = KeyEnvelope{} }
func (*KeyEnvelope) ProtoMessage()               {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) { return fileDescriptorJoin, []int{3} }

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinReq)(nil), "ttn.v3.JoinReq")
	proto.RegisterType((*JoinAns)(nil), "ttn.v3.JoinAns")
	proto.RegisterType((*SessionKeys)(nil), "ttn.v3.SessionKeys")
	proto.RegisterType((*KeyEnvelope)(nil), "ttn.v3.KeyEnvelope")
}
func (m *JoinReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RawPayload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i += copy(dAtA[i:], m.RawPayload)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.Payload.Size()))
	n1, err := m.Payload.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.EndDeviceIdentifiers.Size()))
	n2, err := m.EndDeviceIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.SelectedMacVersion) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.SelectedMacVersion)))
		i += copy(dAtA[i:], m.SelectedMacVersion)
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.NetID.Size()))
	n3, err := m.NetID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.DownlinkSettings.Size()))
	n4, err := m.DownlinkSettings.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.RxDelay != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.RxDelay))
	}
	if m.CFList != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.CFList.Size()))
		n5, err := m.CFList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *JoinAns) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinAns) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RawPayload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i += copy(dAtA[i:], m.RawPayload)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.SessionKeys.Size()))
	n6, err := m.SessionKeys.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.Lifetime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.Lifetime.Size()))
		n7, err := m.Lifetime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *SessionKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionKeys) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SessionKeyId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.SessionKeyId)))
		i += copy(dAtA[i:], m.SessionKeyId)
	}
	if m.FNwkSIntKey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.FNwkSIntKey.Size()))
		n8, err := m.FNwkSIntKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.SNwkSIntKey != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.SNwkSIntKey.Size()))
		n9, err := m.SNwkSIntKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.NwkSEncKey != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.NwkSEncKey.Size()))
		n10, err := m.NwkSEncKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.AppSKey != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.AppSKey.Size()))
		n11, err := m.AppSKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	return i, nil
}

func (m *KeyEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyEnvelope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.Key.Size()))
		n12, err := m.Key.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	if len(m.KekLabel) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.KekLabel)))
		i += copy(dAtA[i:], m.KekLabel)
	}
	return i, nil
}

func encodeFixed64Join(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Join(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintJoin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *JoinReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.Payload.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = m.EndDeviceIdentifiers.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = len(m.SelectedMacVersion)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.NetID.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = m.DownlinkSettings.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.RxDelay != 0 {
		n += 1 + sovJoin(uint64(m.RxDelay))
	}
	if m.CFList != nil {
		l = m.CFList.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	return n
}

func (m *JoinAns) Size() (n int) {
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.SessionKeys.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.Lifetime != nil {
		l = m.Lifetime.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	return n
}

func (m *SessionKeys) Size() (n int) {
	var l int
	_ = l
	l = len(m.SessionKeyId)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.FNwkSIntKey != nil {
		l = m.FNwkSIntKey.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.SNwkSIntKey != nil {
		l = m.SNwkSIntKey.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.NwkSEncKey != nil {
		l = m.NwkSEncKey.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.AppSKey != nil {
		l = m.AppSKey.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	return n
}

func (m *KeyEnvelope) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	l = len(m.KekLabel)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	return n
}

func sovJoin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJoin(x uint64) (n int) {
	return sovJoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *JoinReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinReq{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`Payload:` + strings.Replace(strings.Replace(this.Payload.String(), "Message", "Message", 1), `&`, ``, 1) + `,`,
		`EndDeviceIdentifiers:` + strings.Replace(strings.Replace(this.EndDeviceIdentifiers.String(), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1), `&`, ``, 1) + `,`,
		`SelectedMacVersion:` + fmt.Sprintf("%v", this.SelectedMacVersion) + `,`,
		`NetID:` + fmt.Sprintf("%v", this.NetID) + `,`,
		`DownlinkSettings:` + strings.Replace(strings.Replace(this.DownlinkSettings.String(), "DLSettings", "DLSettings", 1), `&`, ``, 1) + `,`,
		`RxDelay:` + fmt.Sprintf("%v", this.RxDelay) + `,`,
		`CFList:` + strings.Replace(fmt.Sprintf("%v", this.CFList), "CFList", "CFList", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JoinAns) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinAns{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`SessionKeys:` + strings.Replace(strings.Replace(this.SessionKeys.String(), "SessionKeys", "SessionKeys", 1), `&`, ``, 1) + `,`,
		`Lifetime:` + strings.Replace(fmt.Sprintf("%v", this.Lifetime), "Duration", "google_protobuf3.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SessionKeys) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SessionKeys{`,
		`SessionKeyId:` + fmt.Sprintf("%v", this.SessionKeyId) + `,`,
		`FNwkSIntKey:` + strings.Replace(fmt.Sprintf("%v", this.FNwkSIntKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`SNwkSIntKey:` + strings.Replace(fmt.Sprintf("%v", this.SNwkSIntKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`NwkSEncKey:` + strings.Replace(fmt.Sprintf("%v", this.NwkSEncKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`AppSKey:` + strings.Replace(fmt.Sprintf("%v", this.AppSKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyEnvelope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyEnvelope{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`KekLabel:` + fmt.Sprintf("%v", this.KekLabel) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJoin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *JoinReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: JoinReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedMacVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelectedMacVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownlinkSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DownlinkSettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxDelay", wireType)
			}
			m.RxDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RxDelay |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CFList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CFList == nil {
				m.CFList = &CFList{}
			}
			if err := m.CFList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
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
func (m *JoinAns) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: JoinAns: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinAns: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SessionKeys.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lifetime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lifetime == nil {
				m.Lifetime = &google_protobuf3.Duration{}
			}
			if err := m.Lifetime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
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
func (m *SessionKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: SessionKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FNwkSIntKey == nil {
				m.FNwkSIntKey = &KeyEnvelope{}
			}
			if err := m.FNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SNwkSIntKey == nil {
				m.SNwkSIntKey = &KeyEnvelope{}
			}
			if err := m.SNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSEncKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NwkSEncKey == nil {
				m.NwkSEncKey = &KeyEnvelope{}
			}
			if err := m.NwkSEncKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppSKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AppSKey == nil {
				m.AppSKey = &KeyEnvelope{}
			}
			if err := m.AppSKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
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
func (m *KeyEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: KeyEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_pkg_types.AES128Key
			m.Key = &v
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KekLabel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KekLabel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
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
func skipJoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJoin
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
					return 0, ErrIntOverflowJoin
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
					return 0, ErrIntOverflowJoin
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
				return 0, ErrInvalidLengthJoin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJoin
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
				next, err := skipJoin(dAtA[start:])
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
	ErrInvalidLengthJoin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJoin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/join.proto", fileDescriptorJoin) }

var fileDescriptorJoin = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x6e, 0xd3, 0x4a,
	0x18, 0xc7, 0xeb, 0x36, 0xcd, 0x65, 0xd2, 0xdb, 0xf1, 0x39, 0x8b, 0xb4, 0x07, 0x39, 0x51, 0xc4,
	0x22, 0x2c, 0xb0, 0x69, 0xc3, 0x75, 0xd7, 0x86, 0x04, 0x94, 0xa6, 0xad, 0x90, 0x53, 0x81, 0xc4,
	0xc6, 0x72, 0xec, 0x2f, 0xee, 0x60, 0x77, 0xc6, 0x78, 0x26, 0x71, 0xbd, 0xe3, 0x19, 0x78, 0x05,
	0x5e, 0xa6, 0x62, 0xd5, 0x25, 0xea, 0x22, 0x02, 0xaf, 0x58, 0xf2, 0x08, 0xc8, 0xe3, 0x98, 0x44,
	0x95, 0x4a, 0xb3, 0xb2, 0x67, 0xfe, 0xf3, 0xfb, 0x6e, 0xff, 0x0f, 0xa9, 0x0e, 0xe6, 0x67, 0xa3,
	0x81, 0x6a, 0xd1, 0x73, 0xed, 0xf4, 0x0c, 0x4e, 0xcf, 0x30, 0x71, 0xd8, 0x09, 0xf0, 0x90, 0x06,
	0xae, 0xc6, 0x39, 0xd1, 0x4c, 0x1f, 0x6b, 0x1f, 0x28, 0x26, 0xaa, 0x1f, 0x50, 0x4e, 0xe5, 0x3c,
	0xe7, 0x44, 0x1d, 0x37, 0x77, 0x1e, 0xce, 0x71, 0x0e, 0x75, 0xa8, 0x26, 0xe4, 0xc1, 0x68, 0x28,
	0x4e, 0xe2, 0x20, 0xfe, 0x52, 0x6c, 0xe7, 0xf1, 0x22, 0x69, 0x80, 0xd8, 0x86, 0x0d, 0x63, 0x6c,
	0xc1, 0x94, 0xda, 0x5d, 0x84, 0xf2, 0x68, 0x60, 0x86, 0xe6, 0xb4, 0xbe, 0x1d, 0xc5, 0xa1, 0xd4,
	0xf1, 0x60, 0x56, 0x8e, 0x3d, 0x0a, 0x4c, 0x8e, 0xe9, 0x54, 0xaf, 0x7f, 0x5d, 0x41, 0x85, 0x43,
	0x8a, 0x89, 0x0e, 0x1f, 0xe5, 0x2a, 0x2a, 0x07, 0x66, 0x68, 0xf8, 0x66, 0xe4, 0x51, 0xd3, 0xae,
	0x48, 0x35, 0xa9, 0xb1, 0xa6, 0xa3, 0xc0, 0x0c, 0xdf, 0xa4, 0x37, 0xb2, 0x86, 0x0a, 0x99, 0xb8,
	0x5c, 0x93, 0x1a, 0xe5, 0xbd, 0x4d, 0x35, 0x6d, 0x5f, 0x3d, 0x06, 0xc6, 0x4c, 0x07, 0x5a, 0xb9,
	0xcb, 0x49, 0x75, 0x49, 0xcf, 0x5e, 0xc9, 0x1d, 0x84, 0x66, 0x4d, 0x54, 0x56, 0x04, 0x73, 0x2f,
	0x63, 0x3a, 0xc4, 0x6e, 0x0b, 0xa1, 0x6b, 0x03, 0xe1, 0x78, 0x88, 0x21, 0x60, 0xad, 0x62, 0x12,
	0xe0, 0x6a, 0x52, 0x95, 0xf4, 0x12, 0x64, 0xba, 0xfc, 0x08, 0xfd, 0xc7, 0xc0, 0x03, 0x8b, 0x83,
	0x6d, 0x9c, 0x9b, 0x96, 0x31, 0x86, 0x80, 0x61, 0x4a, 0x2a, 0xb9, 0x9a, 0xd4, 0x28, 0xe9, 0x72,
	0xa6, 0x1d, 0x9b, 0xd6, 0xdb, 0x54, 0x91, 0xdf, 0xa1, 0x3c, 0x01, 0x6e, 0x60, 0xbb, 0xb2, 0x9a,
	0x74, 0xd1, 0xda, 0x4f, 0xc2, 0x5e, 0x4f, 0xaa, 0xda, 0x5d, 0x13, 0xf4, 0x5d, 0x47, 0xe3, 0x91,
	0x0f, 0x4c, 0x3d, 0x01, 0xde, 0x6d, 0xc7, 0x93, 0xea, 0xaa, 0xf8, 0xd1, 0x57, 0x09, 0xf0, 0x6e,
	0xd2, 0xd1, 0x3f, 0x36, 0x0d, 0x89, 0x87, 0x89, 0x6b, 0x30, 0xe0, 0x3c, 0xe1, 0x2b, 0x79, 0xd1,
	0x98, 0x9c, 0x35, 0xd6, 0x3e, 0xea, 0x4f, 0x95, 0xe9, 0x3c, 0xb6, 0x32, 0x24, 0xbb, 0x97, 0xb7,
	0x51, 0x31, 0xb8, 0x30, 0x6c, 0xf0, 0xcc, 0xa8, 0x52, 0xa8, 0x49, 0x8d, 0x75, 0xbd, 0x10, 0x5c,
	0xb4, 0x93, 0xa3, 0xdc, 0x44, 0x05, 0x6b, 0x68, 0x78, 0x98, 0xf1, 0x4a, 0x51, 0xc4, 0xdd, 0xc8,
	0xe2, 0xbe, 0x7c, 0x75, 0x84, 0x19, 0x6f, 0xa1, 0x78, 0x52, 0xcd, 0xa7, 0xff, 0x7a, 0xde, 0x1a,
	0x26, 0xdf, 0xc3, 0x5c, 0xb1, 0xb4, 0x85, 0xea, 0x5f, 0xa4, 0xd4, 0xcc, 0x03, 0xc2, 0xee, 0x36,
	0x73, 0x1f, 0xad, 0x31, 0x60, 0xc9, 0xb4, 0x0c, 0x17, 0x22, 0x36, 0x75, 0xf4, 0xdf, 0x2c, 0x59,
	0x3f, 0xd5, 0x7a, 0x10, 0xcd, 0x9b, 0x52, 0x66, 0xb3, 0x6b, 0xf9, 0x09, 0x2a, 0x7a, 0x78, 0x08,
	0x1c, 0x9f, 0x67, 0xde, 0x6e, 0xab, 0xe9, 0xba, 0xa9, 0xd9, 0xba, 0xa9, 0xed, 0xe9, 0xba, 0xe9,
	0x7f, 0x9e, 0xd6, 0x3f, 0x2f, 0xa3, 0xf2, 0x5c, 0x74, 0xf9, 0x3e, 0xda, 0x98, 0x2b, 0x24, 0xf1,
	0x4c, 0x12, 0xbe, 0xae, 0xcd, 0x72, 0x75, 0x6d, 0xf9, 0x05, 0xda, 0x1c, 0x1a, 0x24, 0x74, 0x0d,
	0x66, 0x60, 0xc2, 0x93, 0x97, 0x37, 0x2b, 0xee, 0x41, 0xd4, 0x21, 0x63, 0xf0, 0xa8, 0x0f, 0x7a,
	0x79, 0x78, 0x12, 0xba, 0xfd, 0x2e, 0xe1, 0x3d, 0x88, 0x12, 0x94, 0xdd, 0x40, 0x73, 0x7f, 0x41,
	0xd9, 0x1c, 0xfa, 0x14, 0xad, 0xa7, 0x20, 0x10, 0x4b, 0x80, 0xf9, 0xdb, 0x41, 0x44, 0x42, 0xb7,
	0xdf, 0x21, 0x56, 0xc2, 0x69, 0xa8, 0x64, 0xfa, 0xbe, 0xc1, 0x04, 0x53, 0xbc, 0x9d, 0x29, 0x98,
	0xbe, 0xdf, 0xef, 0x41, 0x54, 0x1f, 0xa1, 0xf2, 0xdc, 0xbd, 0xdc, 0x45, 0x2b, 0x09, 0x29, 0x5c,
	0x6b, 0x3d, 0xbb, 0x9e, 0x54, 0x9b, 0x8b, 0x2f, 0xee, 0x41, 0xa7, 0xbf, 0xbb, 0xf7, 0xbc, 0x07,
	0x91, 0x9e, 0xc4, 0x90, 0xff, 0x47, 0x25, 0x17, 0x5c, 0xc3, 0x33, 0x07, 0xe0, 0x89, 0x91, 0x95,
	0xf4, 0xa2, 0x0b, 0xee, 0x51, 0x72, 0x6e, 0xbd, 0xfe, 0xf6, 0x43, 0x59, 0xfa, 0x14, 0x2b, 0xd2,
	0x65, 0xac, 0x48, 0x57, 0xb1, 0x22, 0x7d, 0x8f, 0x15, 0xe9, 0x67, 0xac, 0x2c, 0xfd, 0x8a, 0x15,
	0xe9, 0xfd, 0x83, 0x85, 0x92, 0x72, 0xe2, 0x0f, 0x06, 0x79, 0xe1, 0x78, 0xf3, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0xa3, 0xca, 0xf8, 0x40, 0x05, 0x00, 0x00,
}
