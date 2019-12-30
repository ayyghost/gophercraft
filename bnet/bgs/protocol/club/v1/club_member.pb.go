// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: club_member.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v2 "github.com/superp00t/gophercraft/bnet/bgs/protocol/v2"
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

type Member struct {
	Id                   *MemberId         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	BattleTag            *string           `protobuf:"bytes,2,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	Role                 []uint32          `protobuf:"varint,3,rep,packed,name=role" json:"role,omitempty"`
	Attribute            []*v2.Attribute   `protobuf:"bytes,4,rep,name=attribute" json:"attribute,omitempty"`
	JoinTime             *uint64           `protobuf:"varint,5,opt,name=join_time,json=joinTime" json:"join_time,omitempty"`
	PresenceLevel        *PresenceLevel    `protobuf:"varint,6,opt,name=presence_level,json=presenceLevel,enum=bgs.protocol.club.v1.PresenceLevel" json:"presence_level,omitempty"`
	ModeratorMute        *bool             `protobuf:"varint,7,opt,name=moderator_mute,json=moderatorMute" json:"moderator_mute,omitempty"`
	WhisperLevel         *WhisperLevel     `protobuf:"varint,8,opt,name=whisper_level,json=whisperLevel,enum=bgs.protocol.club.v1.WhisperLevel" json:"whisper_level,omitempty"`
	Note                 *string           `protobuf:"bytes,9,opt,name=note" json:"note,omitempty"`
	Active               *bool             `protobuf:"varint,50,opt,name=active" json:"active,omitempty"`
	Voice                *MemberVoiceState `protobuf:"bytes,51,opt,name=voice" json:"voice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{0}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetId() *MemberId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Member) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

func (m *Member) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *Member) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Member) GetJoinTime() uint64 {
	if m != nil && m.JoinTime != nil {
		return *m.JoinTime
	}
	return 0
}

func (m *Member) GetPresenceLevel() PresenceLevel {
	if m != nil && m.PresenceLevel != nil {
		return *m.PresenceLevel
	}
	return PresenceLevel_PRESENCE_LEVEL_NONE
}

func (m *Member) GetModeratorMute() bool {
	if m != nil && m.ModeratorMute != nil {
		return *m.ModeratorMute
	}
	return false
}

func (m *Member) GetWhisperLevel() WhisperLevel {
	if m != nil && m.WhisperLevel != nil {
		return *m.WhisperLevel
	}
	return WhisperLevel_WHISPER_LEVEL_OPEN
}

func (m *Member) GetNote() string {
	if m != nil && m.Note != nil {
		return *m.Note
	}
	return ""
}

func (m *Member) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

func (m *Member) GetVoice() *MemberVoiceState {
	if m != nil {
		return m.Voice
	}
	return nil
}

type MemberResult struct {
	MemberId             *MemberId `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Status               *uint32   `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberResult) Reset()         { *m = MemberResult{} }
func (m *MemberResult) String() string { return proto.CompactTextString(m) }
func (*MemberResult) ProtoMessage()    {}
func (*MemberResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{1}
}

func (m *MemberResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberResult.Unmarshal(m, b)
}
func (m *MemberResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberResult.Marshal(b, m, deterministic)
}
func (m *MemberResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberResult.Merge(m, src)
}
func (m *MemberResult) XXX_Size() int {
	return xxx_messageInfo_MemberResult.Size(m)
}
func (m *MemberResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberResult.DiscardUnknown(m)
}

var xxx_messageInfo_MemberResult proto.InternalMessageInfo

func (m *MemberResult) GetMemberId() *MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *MemberResult) GetStatus() uint32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

type RemoveMemberOptions struct {
	Id                   *MemberId          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Reason               *ClubRemovedReason `protobuf:"varint,2,opt,name=reason,enum=bgs.protocol.club.v1.ClubRemovedReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RemoveMemberOptions) Reset()         { *m = RemoveMemberOptions{} }
func (m *RemoveMemberOptions) String() string { return proto.CompactTextString(m) }
func (*RemoveMemberOptions) ProtoMessage()    {}
func (*RemoveMemberOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{2}
}

func (m *RemoveMemberOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveMemberOptions.Unmarshal(m, b)
}
func (m *RemoveMemberOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveMemberOptions.Marshal(b, m, deterministic)
}
func (m *RemoveMemberOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMemberOptions.Merge(m, src)
}
func (m *RemoveMemberOptions) XXX_Size() int {
	return xxx_messageInfo_RemoveMemberOptions.Size(m)
}
func (m *RemoveMemberOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMemberOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMemberOptions proto.InternalMessageInfo

func (m *RemoveMemberOptions) GetId() *MemberId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RemoveMemberOptions) GetReason() ClubRemovedReason {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ClubRemovedReason_CLUB_REMOVED_REASON_NONE
}

type MemberRemovedAssignment struct {
	Id                   *MemberId          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Reason               *ClubRemovedReason `protobuf:"varint,2,opt,name=reason,enum=bgs.protocol.club.v1.ClubRemovedReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MemberRemovedAssignment) Reset()         { *m = MemberRemovedAssignment{} }
func (m *MemberRemovedAssignment) String() string { return proto.CompactTextString(m) }
func (*MemberRemovedAssignment) ProtoMessage()    {}
func (*MemberRemovedAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{3}
}

func (m *MemberRemovedAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberRemovedAssignment.Unmarshal(m, b)
}
func (m *MemberRemovedAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberRemovedAssignment.Marshal(b, m, deterministic)
}
func (m *MemberRemovedAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberRemovedAssignment.Merge(m, src)
}
func (m *MemberRemovedAssignment) XXX_Size() int {
	return xxx_messageInfo_MemberRemovedAssignment.Size(m)
}
func (m *MemberRemovedAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberRemovedAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_MemberRemovedAssignment proto.InternalMessageInfo

func (m *MemberRemovedAssignment) GetId() *MemberId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MemberRemovedAssignment) GetReason() ClubRemovedReason {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ClubRemovedReason_CLUB_REMOVED_REASON_NONE
}

type MemberVoiceOptions struct {
	StreamId             *uint64               `protobuf:"varint,1,opt,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	Joined               *bool                 `protobuf:"varint,2,opt,name=joined" json:"joined,omitempty"`
	Microphone           *VoiceMicrophoneState `protobuf:"varint,3,opt,name=microphone,enum=bgs.protocol.club.v1.VoiceMicrophoneState" json:"microphone,omitempty"`
	Active               *bool                 `protobuf:"varint,4,opt,name=active" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MemberVoiceOptions) Reset()         { *m = MemberVoiceOptions{} }
func (m *MemberVoiceOptions) String() string { return proto.CompactTextString(m) }
func (*MemberVoiceOptions) ProtoMessage()    {}
func (*MemberVoiceOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{4}
}

func (m *MemberVoiceOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberVoiceOptions.Unmarshal(m, b)
}
func (m *MemberVoiceOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberVoiceOptions.Marshal(b, m, deterministic)
}
func (m *MemberVoiceOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberVoiceOptions.Merge(m, src)
}
func (m *MemberVoiceOptions) XXX_Size() int {
	return xxx_messageInfo_MemberVoiceOptions.Size(m)
}
func (m *MemberVoiceOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberVoiceOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MemberVoiceOptions proto.InternalMessageInfo

func (m *MemberVoiceOptions) GetStreamId() uint64 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *MemberVoiceOptions) GetJoined() bool {
	if m != nil && m.Joined != nil {
		return *m.Joined
	}
	return false
}

func (m *MemberVoiceOptions) GetMicrophone() VoiceMicrophoneState {
	if m != nil && m.Microphone != nil {
		return *m.Microphone
	}
	return VoiceMicrophoneState_MICROPHONE_STATE_NORMAL
}

func (m *MemberVoiceOptions) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

type MemberVoiceState struct {
	Id                   *string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	StreamId             *uint64               `protobuf:"varint,2,opt,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	Joined               *bool                 `protobuf:"varint,3,opt,name=joined" json:"joined,omitempty"`
	Microphone           *VoiceMicrophoneState `protobuf:"varint,4,opt,name=microphone,enum=bgs.protocol.club.v1.VoiceMicrophoneState" json:"microphone,omitempty"`
	Active               *bool                 `protobuf:"varint,5,opt,name=active" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MemberVoiceState) Reset()         { *m = MemberVoiceState{} }
func (m *MemberVoiceState) String() string { return proto.CompactTextString(m) }
func (*MemberVoiceState) ProtoMessage()    {}
func (*MemberVoiceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{5}
}

func (m *MemberVoiceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberVoiceState.Unmarshal(m, b)
}
func (m *MemberVoiceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberVoiceState.Marshal(b, m, deterministic)
}
func (m *MemberVoiceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberVoiceState.Merge(m, src)
}
func (m *MemberVoiceState) XXX_Size() int {
	return xxx_messageInfo_MemberVoiceState.Size(m)
}
func (m *MemberVoiceState) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberVoiceState.DiscardUnknown(m)
}

var xxx_messageInfo_MemberVoiceState proto.InternalMessageInfo

func (m *MemberVoiceState) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *MemberVoiceState) GetStreamId() uint64 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *MemberVoiceState) GetJoined() bool {
	if m != nil && m.Joined != nil {
		return *m.Joined
	}
	return false
}

func (m *MemberVoiceState) GetMicrophone() VoiceMicrophoneState {
	if m != nil && m.Microphone != nil {
		return *m.Microphone
	}
	return VoiceMicrophoneState_MICROPHONE_STATE_NORMAL
}

func (m *MemberVoiceState) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

type CreateMemberOptions struct {
	Id                   *MemberId       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Attribute            []*v2.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateMemberOptions) Reset()         { *m = CreateMemberOptions{} }
func (m *CreateMemberOptions) String() string { return proto.CompactTextString(m) }
func (*CreateMemberOptions) ProtoMessage()    {}
func (*CreateMemberOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{6}
}

func (m *CreateMemberOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMemberOptions.Unmarshal(m, b)
}
func (m *CreateMemberOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMemberOptions.Marshal(b, m, deterministic)
}
func (m *CreateMemberOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMemberOptions.Merge(m, src)
}
func (m *CreateMemberOptions) XXX_Size() int {
	return xxx_messageInfo_CreateMemberOptions.Size(m)
}
func (m *CreateMemberOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMemberOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMemberOptions proto.InternalMessageInfo

func (m *CreateMemberOptions) GetId() *MemberId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CreateMemberOptions) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type MemberDescription struct {
	Id                   *MemberId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	BattleTag            *string   `protobuf:"bytes,2,opt,name=battle_tag,json=battleTag" json:"battle_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MemberDescription) Reset()         { *m = MemberDescription{} }
func (m *MemberDescription) String() string { return proto.CompactTextString(m) }
func (*MemberDescription) ProtoMessage()    {}
func (*MemberDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{7}
}

func (m *MemberDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberDescription.Unmarshal(m, b)
}
func (m *MemberDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberDescription.Marshal(b, m, deterministic)
}
func (m *MemberDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberDescription.Merge(m, src)
}
func (m *MemberDescription) XXX_Size() int {
	return xxx_messageInfo_MemberDescription.Size(m)
}
func (m *MemberDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberDescription.DiscardUnknown(m)
}

var xxx_messageInfo_MemberDescription proto.InternalMessageInfo

func (m *MemberDescription) GetId() *MemberId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MemberDescription) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

type RoleOptions struct {
	MemberId             *MemberId `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Role                 []uint32  `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoleOptions) Reset()         { *m = RoleOptions{} }
func (m *RoleOptions) String() string { return proto.CompactTextString(m) }
func (*RoleOptions) ProtoMessage()    {}
func (*RoleOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{8}
}

func (m *RoleOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleOptions.Unmarshal(m, b)
}
func (m *RoleOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleOptions.Marshal(b, m, deterministic)
}
func (m *RoleOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleOptions.Merge(m, src)
}
func (m *RoleOptions) XXX_Size() int {
	return xxx_messageInfo_RoleOptions.Size(m)
}
func (m *RoleOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RoleOptions proto.InternalMessageInfo

func (m *RoleOptions) GetMemberId() *MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *RoleOptions) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type RoleAssignment struct {
	MemberId             *MemberId `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Role                 []uint32  `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoleAssignment) Reset()         { *m = RoleAssignment{} }
func (m *RoleAssignment) String() string { return proto.CompactTextString(m) }
func (*RoleAssignment) ProtoMessage()    {}
func (*RoleAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{9}
}

func (m *RoleAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAssignment.Unmarshal(m, b)
}
func (m *RoleAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAssignment.Marshal(b, m, deterministic)
}
func (m *RoleAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAssignment.Merge(m, src)
}
func (m *RoleAssignment) XXX_Size() int {
	return xxx_messageInfo_RoleAssignment.Size(m)
}
func (m *RoleAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAssignment proto.InternalMessageInfo

func (m *RoleAssignment) GetMemberId() *MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *RoleAssignment) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type MemberAttributeAssignment struct {
	MemberId             *MemberId       `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Attribute            []*v2.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MemberAttributeAssignment) Reset()         { *m = MemberAttributeAssignment{} }
func (m *MemberAttributeAssignment) String() string { return proto.CompactTextString(m) }
func (*MemberAttributeAssignment) ProtoMessage()    {}
func (*MemberAttributeAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{10}
}

func (m *MemberAttributeAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberAttributeAssignment.Unmarshal(m, b)
}
func (m *MemberAttributeAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberAttributeAssignment.Marshal(b, m, deterministic)
}
func (m *MemberAttributeAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberAttributeAssignment.Merge(m, src)
}
func (m *MemberAttributeAssignment) XXX_Size() int {
	return xxx_messageInfo_MemberAttributeAssignment.Size(m)
}
func (m *MemberAttributeAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberAttributeAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_MemberAttributeAssignment proto.InternalMessageInfo

func (m *MemberAttributeAssignment) GetMemberId() *MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *MemberAttributeAssignment) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type SubscriberStateOptions struct {
	Voice                *MemberVoiceOptions `protobuf:"bytes,1,opt,name=voice" json:"voice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubscriberStateOptions) Reset()         { *m = SubscriberStateOptions{} }
func (m *SubscriberStateOptions) String() string { return proto.CompactTextString(m) }
func (*SubscriberStateOptions) ProtoMessage()    {}
func (*SubscriberStateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{11}
}

func (m *SubscriberStateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberStateOptions.Unmarshal(m, b)
}
func (m *SubscriberStateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberStateOptions.Marshal(b, m, deterministic)
}
func (m *SubscriberStateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberStateOptions.Merge(m, src)
}
func (m *SubscriberStateOptions) XXX_Size() int {
	return xxx_messageInfo_SubscriberStateOptions.Size(m)
}
func (m *SubscriberStateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberStateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberStateOptions proto.InternalMessageInfo

func (m *SubscriberStateOptions) GetVoice() *MemberVoiceOptions {
	if m != nil {
		return m.Voice
	}
	return nil
}

type SubscriberStateAssignment struct {
	MemberId             *MemberId         `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Active               *bool             `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
	Voice                *MemberVoiceState `protobuf:"bytes,3,opt,name=voice" json:"voice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SubscriberStateAssignment) Reset()         { *m = SubscriberStateAssignment{} }
func (m *SubscriberStateAssignment) String() string { return proto.CompactTextString(m) }
func (*SubscriberStateAssignment) ProtoMessage()    {}
func (*SubscriberStateAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{12}
}

func (m *SubscriberStateAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberStateAssignment.Unmarshal(m, b)
}
func (m *SubscriberStateAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberStateAssignment.Marshal(b, m, deterministic)
}
func (m *SubscriberStateAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberStateAssignment.Merge(m, src)
}
func (m *SubscriberStateAssignment) XXX_Size() int {
	return xxx_messageInfo_SubscriberStateAssignment.Size(m)
}
func (m *SubscriberStateAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberStateAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberStateAssignment proto.InternalMessageInfo

func (m *SubscriberStateAssignment) GetMemberId() *MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *SubscriberStateAssignment) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

func (m *SubscriberStateAssignment) GetVoice() *MemberVoiceState {
	if m != nil {
		return m.Voice
	}
	return nil
}

type MemberStateOptions struct {
	Attribute            []*v2.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	PresenceLevel        *PresenceLevel  `protobuf:"varint,2,opt,name=presence_level,json=presenceLevel,enum=bgs.protocol.club.v1.PresenceLevel" json:"presence_level,omitempty"`
	ModeratorMute        *bool           `protobuf:"varint,3,opt,name=moderator_mute,json=moderatorMute" json:"moderator_mute,omitempty"`
	WhisperLevel         *WhisperLevel   `protobuf:"varint,4,opt,name=whisper_level,json=whisperLevel,enum=bgs.protocol.club.v1.WhisperLevel" json:"whisper_level,omitempty"`
	Note                 *string         `protobuf:"bytes,5,opt,name=note" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MemberStateOptions) Reset()         { *m = MemberStateOptions{} }
func (m *MemberStateOptions) String() string { return proto.CompactTextString(m) }
func (*MemberStateOptions) ProtoMessage()    {}
func (*MemberStateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{13}
}

func (m *MemberStateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberStateOptions.Unmarshal(m, b)
}
func (m *MemberStateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberStateOptions.Marshal(b, m, deterministic)
}
func (m *MemberStateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberStateOptions.Merge(m, src)
}
func (m *MemberStateOptions) XXX_Size() int {
	return xxx_messageInfo_MemberStateOptions.Size(m)
}
func (m *MemberStateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberStateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MemberStateOptions proto.InternalMessageInfo

func (m *MemberStateOptions) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *MemberStateOptions) GetPresenceLevel() PresenceLevel {
	if m != nil && m.PresenceLevel != nil {
		return *m.PresenceLevel
	}
	return PresenceLevel_PRESENCE_LEVEL_NONE
}

func (m *MemberStateOptions) GetModeratorMute() bool {
	if m != nil && m.ModeratorMute != nil {
		return *m.ModeratorMute
	}
	return false
}

func (m *MemberStateOptions) GetWhisperLevel() WhisperLevel {
	if m != nil && m.WhisperLevel != nil {
		return *m.WhisperLevel
	}
	return WhisperLevel_WHISPER_LEVEL_OPEN
}

func (m *MemberStateOptions) GetNote() string {
	if m != nil && m.Note != nil {
		return *m.Note
	}
	return ""
}

type MemberStateAssignment struct {
	MemberId             *MemberId       `protobuf:"bytes,1,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
	Attribute            []*v2.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	PresenceLevel        *PresenceLevel  `protobuf:"varint,3,opt,name=presence_level,json=presenceLevel,enum=bgs.protocol.club.v1.PresenceLevel" json:"presence_level,omitempty"`
	ModeratorMute        *bool           `protobuf:"varint,4,opt,name=moderator_mute,json=moderatorMute" json:"moderator_mute,omitempty"`
	WhisperLevel         *WhisperLevel   `protobuf:"varint,5,opt,name=whisper_level,json=whisperLevel,enum=bgs.protocol.club.v1.WhisperLevel" json:"whisper_level,omitempty"`
	Note                 *string         `protobuf:"bytes,6,opt,name=note" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MemberStateAssignment) Reset()         { *m = MemberStateAssignment{} }
func (m *MemberStateAssignment) String() string { return proto.CompactTextString(m) }
func (*MemberStateAssignment) ProtoMessage()    {}
func (*MemberStateAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d38a4243e5ce23bb, []int{14}
}

func (m *MemberStateAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberStateAssignment.Unmarshal(m, b)
}
func (m *MemberStateAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberStateAssignment.Marshal(b, m, deterministic)
}
func (m *MemberStateAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberStateAssignment.Merge(m, src)
}
func (m *MemberStateAssignment) XXX_Size() int {
	return xxx_messageInfo_MemberStateAssignment.Size(m)
}
func (m *MemberStateAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberStateAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_MemberStateAssignment proto.InternalMessageInfo

func (m *MemberStateAssignment) GetMemberId() *MemberId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *MemberStateAssignment) GetAttribute() []*v2.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *MemberStateAssignment) GetPresenceLevel() PresenceLevel {
	if m != nil && m.PresenceLevel != nil {
		return *m.PresenceLevel
	}
	return PresenceLevel_PRESENCE_LEVEL_NONE
}

func (m *MemberStateAssignment) GetModeratorMute() bool {
	if m != nil && m.ModeratorMute != nil {
		return *m.ModeratorMute
	}
	return false
}

func (m *MemberStateAssignment) GetWhisperLevel() WhisperLevel {
	if m != nil && m.WhisperLevel != nil {
		return *m.WhisperLevel
	}
	return WhisperLevel_WHISPER_LEVEL_OPEN
}

func (m *MemberStateAssignment) GetNote() string {
	if m != nil && m.Note != nil {
		return *m.Note
	}
	return ""
}

func init() {
	proto.RegisterType((*Member)(nil), "bgs.protocol.club.v1.Member")
	proto.RegisterType((*MemberResult)(nil), "bgs.protocol.club.v1.MemberResult")
	proto.RegisterType((*RemoveMemberOptions)(nil), "bgs.protocol.club.v1.RemoveMemberOptions")
	proto.RegisterType((*MemberRemovedAssignment)(nil), "bgs.protocol.club.v1.MemberRemovedAssignment")
	proto.RegisterType((*MemberVoiceOptions)(nil), "bgs.protocol.club.v1.MemberVoiceOptions")
	proto.RegisterType((*MemberVoiceState)(nil), "bgs.protocol.club.v1.MemberVoiceState")
	proto.RegisterType((*CreateMemberOptions)(nil), "bgs.protocol.club.v1.CreateMemberOptions")
	proto.RegisterType((*MemberDescription)(nil), "bgs.protocol.club.v1.MemberDescription")
	proto.RegisterType((*RoleOptions)(nil), "bgs.protocol.club.v1.RoleOptions")
	proto.RegisterType((*RoleAssignment)(nil), "bgs.protocol.club.v1.RoleAssignment")
	proto.RegisterType((*MemberAttributeAssignment)(nil), "bgs.protocol.club.v1.MemberAttributeAssignment")
	proto.RegisterType((*SubscriberStateOptions)(nil), "bgs.protocol.club.v1.SubscriberStateOptions")
	proto.RegisterType((*SubscriberStateAssignment)(nil), "bgs.protocol.club.v1.SubscriberStateAssignment")
	proto.RegisterType((*MemberStateOptions)(nil), "bgs.protocol.club.v1.MemberStateOptions")
	proto.RegisterType((*MemberStateAssignment)(nil), "bgs.protocol.club.v1.MemberStateAssignment")
}

func init() { proto.RegisterFile("club_member.proto", fileDescriptor_d38a4243e5ce23bb) }

var fileDescriptor_d38a4243e5ce23bb = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xc1, 0x8e, 0xdc, 0x44,
	0x10, 0xc5, 0x1e, 0xcf, 0x30, 0x53, 0x9b, 0x19, 0x92, 0x4e, 0x58, 0x9c, 0x8d, 0x40, 0x23, 0x47,
	0x10, 0x8b, 0x83, 0x9d, 0x0c, 0x17, 0x04, 0x11, 0x28, 0x09, 0x12, 0x22, 0x62, 0x45, 0xd4, 0x89,
	0x00, 0x71, 0xb1, 0x6c, 0x4f, 0xe1, 0x6d, 0x64, 0xbb, 0xad, 0xee, 0xb6, 0x03, 0x27, 0x6e, 0x1c,
	0x38, 0xf2, 0x01, 0xfc, 0x01, 0xfc, 0x00, 0xff, 0xc2, 0x17, 0xf0, 0x11, 0xc8, 0xdd, 0xf6, 0xac,
	0x27, 0x0c, 0x61, 0x37, 0x3b, 0xc0, 0xcd, 0xd5, 0x7e, 0xe5, 0xf7, 0xfc, 0xaa, 0xba, 0x0a, 0xae,
	0xa4, 0x79, 0x9d, 0x44, 0x05, 0x16, 0x09, 0x8a, 0xa0, 0x12, 0x5c, 0x71, 0x72, 0x2d, 0xc9, 0xa4,
	0x79, 0x4c, 0x79, 0x1e, 0xb4, 0xef, 0x83, 0xe6, 0xce, 0xd1, 0xb5, 0x01, 0x30, 0x62, 0x6b, 0x03,
	0x38, 0xba, 0x95, 0xe5, 0x3c, 0x89, 0xf3, 0x08, 0xbf, 0x55, 0x58, 0x4a, 0xc6, 0x4b, 0x19, 0x16,
	0x28, 0x65, 0x9c, 0x61, 0xc4, 0x2b, 0xd5, 0xc6, 0x1d, 0xf0, 0x66, 0x5c, 0xb1, 0x30, 0xcd, 0x19,
	0x96, 0x2a, 0x6c, 0x56, 0x61, 0xac, 0x94, 0x60, 0x49, 0xad, 0x30, 0x52, 0xdf, 0x55, 0xd8, 0x83,
	0x5e, 0xd1, 0x1c, 0x58, 0xd6, 0x45, 0x7f, 0x20, 0xaa, 0x74, 0x88, 0xf0, 0xfe, 0x18, 0xc1, 0xe4,
	0x58, 0x6b, 0x20, 0x01, 0xd8, 0x6c, 0xed, 0x5a, 0x4b, 0xcb, 0x3f, 0x58, 0xbd, 0x11, 0xec, 0xd2,
	0x1c, 0x18, 0xe4, 0x27, 0x6b, 0x6a, 0xb3, 0x35, 0x79, 0x1d, 0x20, 0x89, 0x95, 0xca, 0x31, 0x52,
	0x71, 0xe6, 0xda, 0x4b, 0xcb, 0x9f, 0xd1, 0x99, 0x39, 0x79, 0x12, 0x67, 0xe4, 0x10, 0x1c, 0xc1,
	0x73, 0x74, 0x47, 0xcb, 0x91, 0x3f, 0xbf, 0x6f, 0x5f, 0xb6, 0xa8, 0x8e, 0xc9, 0xbb, 0x30, 0xdb,
	0x88, 0x75, 0x9d, 0xe5, 0xc8, 0x3f, 0x58, 0x1d, 0x6d, 0xb3, 0x35, 0xab, 0xe0, 0x5e, 0x8f, 0xa0,
	0xa7, 0x60, 0x72, 0x03, 0x66, 0xdf, 0x70, 0x56, 0x46, 0x8a, 0x15, 0xe8, 0x8e, 0x97, 0x96, 0xef,
	0xd0, 0x69, 0x7b, 0xf0, 0x84, 0x15, 0x48, 0x1e, 0xc2, 0xa2, 0x12, 0x28, 0xb1, 0x4c, 0x31, 0xca,
	0xb1, 0xc1, 0xdc, 0x9d, 0x2c, 0x2d, 0x7f, 0xb1, 0xba, 0xb9, 0xfb, 0x4f, 0x1e, 0x75, 0xd8, 0x4f,
	0x5b, 0x28, 0x9d, 0x57, 0xc3, 0x90, 0xbc, 0x09, 0x8b, 0x82, 0xaf, 0x51, 0xc4, 0x8a, 0x8b, 0xa8,
	0x68, 0x75, 0xbe, 0xbc, 0xb4, 0xfc, 0x29, 0x9d, 0x6f, 0x4e, 0x8f, 0x5b, 0x3d, 0x1f, 0xc3, 0xfc,
	0xe9, 0x09, 0x93, 0x15, 0x8a, 0x8e, 0x71, 0xaa, 0x19, 0xbd, 0xdd, 0x8c, 0x5f, 0x18, 0xa8, 0x21,
	0xbc, 0xf4, 0x74, 0x10, 0x11, 0x02, 0x4e, 0xc9, 0x15, 0xba, 0x33, 0xed, 0xa1, 0x7e, 0x26, 0x87,
	0x30, 0x89, 0x53, 0xc5, 0x1a, 0x74, 0x57, 0x9a, 0xbb, 0x8b, 0xc8, 0x5d, 0x18, 0x37, 0x9c, 0xa5,
	0xe8, 0xbe, 0xa3, 0x0b, 0xf5, 0xd6, 0xf3, 0x0a, 0xf5, 0x79, 0x0b, 0x7c, 0xac, 0x62, 0x85, 0xd4,
	0x24, 0x79, 0x29, 0x5c, 0x32, 0xaf, 0x28, 0xca, 0x3a, 0x57, 0xe4, 0x7d, 0x98, 0x6d, 0x3a, 0xf0,
	0x8c, 0xa5, 0x9f, 0x16, 0xdd, 0x53, 0x2b, 0x51, 0xaa, 0x58, 0xd5, 0x52, 0x17, 0x7f, 0x4e, 0xbb,
	0xc8, 0xfb, 0xc1, 0x82, 0xab, 0x14, 0x0b, 0xde, 0xa0, 0x49, 0xfa, 0xcc, 0x34, 0xee, 0xb9, 0x1b,
	0xec, 0x43, 0x98, 0x08, 0x8c, 0x25, 0x2f, 0xf5, 0xf7, 0x17, 0xab, 0x5b, 0xbb, 0x73, 0x1e, 0xe4,
	0x75, 0x62, 0xe8, 0xd6, 0x54, 0xc3, 0x69, 0x97, 0xe6, 0xfd, 0x68, 0xc1, 0x6b, 0xfd, 0xef, 0xea,
	0xf7, 0xf7, 0xa4, 0x64, 0x59, 0x59, 0x60, 0xa9, 0xfe, 0x7b, 0x31, 0xbf, 0x58, 0x40, 0x06, 0x65,
	0xe9, 0x4d, 0xb9, 0x01, 0x33, 0xa9, 0x04, 0xc6, 0x45, 0x5f, 0x01, 0x87, 0x4e, 0xcd, 0x81, 0x71,
	0xb8, 0x6d, 0x70, 0x5c, 0x6b, 0xd2, 0x29, 0xed, 0x22, 0xf2, 0x10, 0xa0, 0x60, 0xa9, 0xe0, 0xd5,
	0x09, 0x2f, 0xdb, 0x1b, 0xd6, 0x0a, 0x7a, 0x7b, 0xb7, 0x20, 0x4d, 0x76, 0xbc, 0x01, 0x9b, 0x6e,
	0x18, 0x64, 0x0f, 0x1a, 0xcd, 0x19, 0x36, 0x9a, 0xf7, 0x9b, 0x05, 0x97, 0x9f, 0x6d, 0x23, 0xb2,
	0xd8, 0xb8, 0x36, 0xd3, 0xae, 0x6c, 0xa9, 0xb7, 0xff, 0x56, 0xfd, 0xe8, 0x39, 0xea, 0x9d, 0x3d,
	0xa9, 0x1f, 0x6f, 0xa9, 0xff, 0x1e, 0xae, 0x3e, 0x10, 0x18, 0xab, 0x0b, 0xb6, 0xe0, 0xd6, 0xb0,
	0xb2, 0xcf, 0x31, 0xac, 0xbc, 0x04, 0xae, 0x98, 0x2f, 0x7d, 0x84, 0x32, 0x15, 0x4c, 0xf3, 0xef,
	0x79, 0xc4, 0x7a, 0x09, 0x1c, 0x50, 0x9e, 0x6f, 0x5a, 0xe9, 0x82, 0x97, 0xd9, 0x8c, 0x6b, 0x7b,
	0x7b, 0x5c, 0x7b, 0x08, 0x8b, 0x96, 0x63, 0x70, 0x73, 0xfe, 0x15, 0x9a, 0x9f, 0x2c, 0xb8, 0x6e,
	0xe0, 0x1b, 0x37, 0xf7, 0x45, 0xf9, 0xe2, 0x35, 0xfc, 0x12, 0x0e, 0x1f, 0xd7, 0x49, 0x5b, 0xbe,
	0x04, 0x85, 0xee, 0xbd, 0xde, 0xea, 0x0f, 0xfa, 0x29, 0x6c, 0xc4, 0xf8, 0xff, 0x38, 0x85, 0xbb,
	0xc4, 0x7e, 0x0e, 0xff, 0x6a, 0xc1, 0xf5, 0x67, 0x3e, 0xbd, 0x3f, 0x87, 0xfb, 0x1b, 0x61, 0xef,
	0x5e, 0x1c, 0xa3, 0x17, 0x59, 0x1c, 0x3f, 0xdb, 0xfd, 0xf4, 0xda, 0xf2, 0x61, 0xcb, 0x5b, 0xeb,
	0x3c, 0xcb, 0xfc, 0xaf, 0xfb, 0xda, 0xde, 0xe3, 0xbe, 0x1e, 0x9d, 0x69, 0x5f, 0x3b, 0x17, 0xdc,
	0xd7, 0xe3, 0xd3, 0x7d, 0xed, 0xfd, 0x6e, 0xc3, 0xab, 0x03, 0x83, 0xfe, 0xf7, 0xe6, 0xdd, 0x61,
	0xf0, 0x68, 0x8f, 0x06, 0x3b, 0x67, 0x32, 0x78, 0x7c, 0x41, 0x83, 0x27, 0xa7, 0x06, 0xdf, 0xbf,
	0xfb, 0xd5, 0x7b, 0x19, 0x53, 0x27, 0x75, 0x12, 0xa4, 0xbc, 0x08, 0x65, 0x5d, 0xa1, 0xa8, 0x6e,
	0xdf, 0x56, 0x61, 0xc6, 0xab, 0x13, 0x14, 0xa9, 0x88, 0xbf, 0x56, 0x61, 0x52, 0xa2, 0x0a, 0x93,
	0x4c, 0x86, 0x3d, 0x59, 0xd8, 0x92, 0x85, 0xcd, 0x9d, 0x47, 0x2f, 0xfd, 0x19, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x4e, 0xd8, 0xcd, 0x99, 0x0b, 0x00, 0x00,
}