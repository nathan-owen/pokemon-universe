/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import (
	"net"
	pnet "network" // PU Network packet
)

type Connection struct {
	Socket net.Conn
}

func NewConnection(_socket net.Conn) *Connection {
	return &Connection { Socket : _socket }
}

func (c *Connection) HandleConnection() {
	for {
		packet := pnet.NewPacket()
		var headerbuffer [2]uint8 
		recv, err := c.Socket.Read(headerbuffer[0:])
		if err != nil || recv == 0 {
			break
		}
		copy(packet.Buffer[0:2], headerbuffer[0:2])
		packet.GetHeader()
		
		databuffer := make([]uint8, packet.MsgSize)
		recv, err = c.Socket.Read(databuffer[0:])
		if recv == 0 {	
			continue 
		} else if err != nil {
			g_logger.Printf("Connection read error: %v", err)
		}
		
		copy(packet.Buffer[2:], databuffer[:])
		c.ProcessPacket(packet)
	}
	
	g_logger.Println("Connection closed")
}

func (c *Connection) ProcessPacket(packet *pnet.Packet) {
	header := packet.ReadUint8()
	switch header {
		case pnet.HEADER_LOGIN:
	}
}