package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"learngo2/chat_server/proto"
	"net"
)

type Client struct {
	conn net.Conn
	buf  [8192]byte
}

func (p *Client) readPackage() (*proto.Message, error) {
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		return nil, errors.New("read header failed")
	}
	buffer := bytes.NewBuffer(p.buf[0:4])
	var packLen uint32
	err = binary.Read(buffer, binary.BigEndian, &packLen)
	if err != nil {
		fmt.Println("read package len failed")
		return nil, err
	}

	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		return nil, errors.New("read body failed")
	}
	var msg proto.Message
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err: ", err)
		return nil, err
	}
	return &msg, err
}

func (p *Client) writePackage(data []byte) error {
	packLen := uint32(len(data))
	buffer := bytes.NewBuffer(p.buf[0:4])
	err := binary.Write(buffer, binary.BigEndian, packLen)
	if err != nil {
		fmt.Println("read package len failed")
		return err
	}

	n, err := p.conn.Write(p.buf[0:4])
	if n != 4 {
		return errors.New("write data failed")
	}

	n, err = p.conn.Write(data)
	if n != 4 {
		return errors.New("write data failed")
	}
	if n != int(packLen) {
		fmt.Println("write data not finished")
		return errors.New("write data not finished")
	}
	return nil
}

func (p *Client) Process() error {
	for {
		msg, err := p.readPackage()
		if err != nil {
			return err
		}
		err = p.processMsg(msg)
		if err != nil {
			return err
		}
	}
}

func (p *Client) processMsg(msg *proto.Message) error {
	switch msg.Cmd {
	case proto.UserLogin:
		return p.login(msg)
	case proto.UserRegister:
		return p.register(msg)
	default:
		return errors.New("unSupport message")
	}
}

func (p *Client) loginResp(err error) {
	var respMsg proto.Message
	respMsg.Cmd = proto.UserLoginRes

	var loginRes proto.LoginCmdRes
	loginRes.Code = 200
	if err != nil {
		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%v", err)
	}
	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed, ", err)
		return
	}
}

func (p *Client) login(msg *proto.Message) error {
	var err error
	defer func() {
		p.loginResp(err)
	}()

	var cmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return err
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return err
	}
	return nil
}

func (p *Client) register(msg *proto.Message) error {
	var cmd proto.RegisterCmd
	err := json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return err
	}
	err = mgr.Register(&cmd.User)
	if err != nil {
		return err
	}
	return nil
}
