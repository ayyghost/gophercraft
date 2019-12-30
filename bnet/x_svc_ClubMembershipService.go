// generated by protoc-gen-gcraft : DO NOT EDIT
package bnet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	membership "github.com/superp00t/gophercraft/bnet/bgs/protocol/club/v1/membership"
	math "math"
)

// shut the compiler up
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = protocol.E_MethodOptions

const ClubMembershipServiceHash = 0x94B94786

type ClubMembershipService interface {
	Subscribe(*Conn, uint32, *membership.SubscribeRequest)
	Unsubscribe(*Conn, uint32, *membership.UnsubscribeRequest)
	GetState(*Conn, uint32, *membership.GetStateRequest)
	UpdateClubSharedSettings(*Conn, uint32, *membership.UpdateClubSharedSettingsRequest)
	GetStreamMentions(*Conn, uint32, *membership.GetStreamMentionsRequest)
	RemoveStreamMentions(*Conn, uint32, *membership.RemoveStreamMentionsRequest)
	AdvanceStreamMentionViewTime(*Conn, uint32, *membership.AdvanceStreamMentionViewTimeRequest)
}

func DispatchClubMembershipService(conn *Conn, svc ClubMembershipService, token uint32, method uint32, data []byte) error {
	switch method {
	case 1:
		var args membership.SubscribeRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.Subscribe(conn, token, &args)
	case 2:
		var args membership.UnsubscribeRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.Unsubscribe(conn, token, &args)
	case 3:
		var args membership.GetStateRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetState(conn, token, &args)
	case 4:
		var args membership.UpdateClubSharedSettingsRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.UpdateClubSharedSettings(conn, token, &args)
	case 5:
		var args membership.GetStreamMentionsRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetStreamMentions(conn, token, &args)
	case 6:
		var args membership.RemoveStreamMentionsRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.RemoveStreamMentions(conn, token, &args)
	case 7:
		var args membership.AdvanceStreamMentionViewTimeRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.AdvanceStreamMentionViewTime(conn, token, &args)
	}
	return nil
}

type EmptyClubMembershipService struct{}

func (e EmptyClubMembershipService) Subscribe(conn *Conn, token uint32, args *membership.SubscribeRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyClubMembershipService) Unsubscribe(conn *Conn, token uint32, args *membership.UnsubscribeRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyClubMembershipService) GetState(conn *Conn, token uint32, args *membership.GetStateRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyClubMembershipService) UpdateClubSharedSettings(conn *Conn, token uint32, args *membership.UpdateClubSharedSettingsRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyClubMembershipService) GetStreamMentions(conn *Conn, token uint32, args *membership.GetStreamMentionsRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyClubMembershipService) RemoveStreamMentions(conn *Conn, token uint32, args *membership.RemoveStreamMentionsRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyClubMembershipService) AdvanceStreamMentionViewTime(conn *Conn, token uint32, args *membership.AdvanceStreamMentionViewTimeRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (c *Conn) ClubMembershipService_Request_Subscribe(args *membership.SubscribeRequest) (*membership.SubscribeResponse, error) {
	header, bytes, err := c.Request(ClubMembershipServiceHash, 1, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out membership.SubscribeResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) ClubMembershipService_Request_Unsubscribe(args *membership.UnsubscribeRequest) error {
	header, _, err := c.Request(ClubMembershipServiceHash, 2, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) ClubMembershipService_Request_GetState(args *membership.GetStateRequest) (*membership.GetStateResponse, error) {
	header, bytes, err := c.Request(ClubMembershipServiceHash, 3, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out membership.GetStateResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) ClubMembershipService_Request_UpdateClubSharedSettings(args *membership.UpdateClubSharedSettingsRequest) error {
	header, _, err := c.Request(ClubMembershipServiceHash, 4, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) ClubMembershipService_Request_GetStreamMentions(args *membership.GetStreamMentionsRequest) (*membership.GetStreamMentionsResponse, error) {
	header, bytes, err := c.Request(ClubMembershipServiceHash, 5, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out membership.GetStreamMentionsResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) ClubMembershipService_Request_RemoveStreamMentions(args *membership.RemoveStreamMentionsRequest) error {
	header, _, err := c.Request(ClubMembershipServiceHash, 6, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) ClubMembershipService_Request_AdvanceStreamMentionViewTime(args *membership.AdvanceStreamMentionViewTimeRequest) error {
	header, _, err := c.Request(ClubMembershipServiceHash, 7, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}