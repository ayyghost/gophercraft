// generated by protoc-gen-gcraft : DO NOT EDIT
package bnet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v1 "github.com/superp00t/gophercraft/bnet/bgs/protocol/presence/v1"
	math "math"
)

// shut the compiler up
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = protocol.E_MethodOptions

const PresenceListenerHash = 0x890AB85F

type PresenceListener interface {
	OnSubscribe(*Conn, uint32, *v1.SubscribeNotification)
	OnStateChanged(*Conn, uint32, *v1.StateChangedNotification)
}

func DispatchPresenceListener(conn *Conn, svc PresenceListener, token uint32, method uint32, data []byte) error {
	switch method {
	case 1:
		var args v1.SubscribeNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnSubscribe(conn, token, &args)
	case 2:
		var args v1.StateChangedNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnStateChanged(conn, token, &args)
	}
	return nil
}

type EmptyPresenceListener struct{}

func (e EmptyPresenceListener) OnSubscribe(conn *Conn, token uint32, args *v1.SubscribeNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyPresenceListener) OnStateChanged(conn *Conn, token uint32, args *v1.StateChangedNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (c *Conn) PresenceListener_Request_OnSubscribe(args *v1.SubscribeNotification) error {
	header, _, err := c.Request(PresenceListenerHash, 1, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) PresenceListener_Request_OnStateChanged(args *v1.StateChangedNotification) error {
	header, _, err := c.Request(PresenceListenerHash, 2, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}