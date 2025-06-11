package message
import (
	"encoding/binary"
	"fmt"
	"io"
)

type messageID uint8

const (
	// MsgChoke chokes the receiver
	MsgChoke messageID = 0
	// MsgUnchoke unchokes the receiver
	MsgUnchoke messageID = 1
	// MsgInterested expresses interest in receiving data
	MsgInterested messageID = 2
	// MsgNotInterested expresses disinterest in receiving data
	MsgNotInterested messageID = 3
	// MsgHave alerts the receiver that the sender has downloaded a piece
	MsgHave messageID = 4
	// MsgBitfield encodes which pieces that the sender has downloaded
	MsgBitfield messageID = 5
	// MsgRequest requests a block of data from the receiver
	MsgRequest messageID = 6
	// MsgPiece delivers a block of data to fulfill a request
	MsgPiece messageID = 7
	// MsgCancel cancels a request
	MsgCancel messageID = 8
)

type Message struct{
	ID messageID
	Payload []byte
}
// FormatRequest creates a request message
func FormatRequest (index,begin,length int) *Message{
	payload := make([]byte,12)
	binary.BigEndian.PutUint32(payload[0:4],uint32(index))
	binary.BigEndian.PutUint32(payload[4:8],uint32(begin))
	binary.BigEndian.PutUint32(payload[8:12],uint32(length))
	return &Message{
		ID:MsgRequest,
		Payload: payload,
	}
}

// Format have creates a Have mesage
func FormatHave(index int) *Message{
	payload := make([]byte,4)
	binary.BigEndian.PutUint32(payload, uint32(index))
	return &Message{ID:MsgHave,Payload:payload}
}

// copies the piece payload
func ParsePiece(index int,buf []byte,msg *Message) (int ,error){
	if msg.ID != MsgPiece{
		return 0,fmt.Errorf("Expected Piece ID:%d got ID %d",MsgPiece,msg.ID)
	}
	if len(msg.Payload) < 8{
		return 0, fmt.Errorf("Payload too short. %d < 8", len(msg.Payload))
	}
	parsedIndex := int(binary.BigEndian.Uint32(msg.Payload[0:4]))
	if parsedIndex != index{
		return 0,fmt.Errorf("Expected Index %d got %d",index,parsedIndex)
	}
	begin := int(binary.BigEndian.Uint32(msg.Payload[4:8]))
	if begin >= len(buf){
		return 0,fmt.Errorf("Begin offset too high %d >= %d",begin,len(buf))
	}
	data := msg.Payload[8:]
	if begin+len(data) > len(buf){
		return 0,fmt.Errorf("Data too long [%d] for offset %d with length %d",len(data),begin,len(buf))
	}
	copy(buf[begin:],data)
	return len(data),nil
}

// parses a Have message
func ParseHave (msg *Message)(int ,error){
	if msg.ID != MsgHave{
		return 0,fmt.Errorf("Expected Have ID %d got ID %d",MsgHave,msg.ID)
	}
	if len(msg.Payload) != 4{
		return 0,fmt.Errorf("Expected payload length 4, got length %d", len(msg.Payload))

	}
	index := int(binary.BigEndian.Uint32(msg.Payload))
	return index,nil
}
// convets the message into <prefix>->4bytes,<msgID>->1 byte,<payload>->rest of the bytes
func (m *Message) Serialize() []byte{
	if m == nil{
		return make([]byte,4)
	}
	length := uint32(len(m.Payload)+1)
	buf := make([]byte,4+length)
	binary.BigEndian.PutUint32(buf[0:4],length)
	buf[4] = byte(m.ID)
	copy(buf[5:],m.Payload)
	return buf
}

// read 
func Read(r io.Reader) (*Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)

	// keep-alive message
	if length == 0 {
		return nil, nil
	}

	messageBuf := make([]byte, length)
	_, err = io.ReadFull(r, messageBuf)
	if err != nil {
		return nil, err
	}

	m := Message{
		ID:      messageID(messageBuf[0]),
		Payload: messageBuf[1:],
	}

	return &m, nil
}

func (m *Message) name() string {
	if m == nil {
		return "KeepAlive"
	}
	switch m.ID {
	case MsgChoke:
		return "Choke"
	case MsgUnchoke:
		return "Unchoke"
	case MsgInterested:
		return "Interested"
	case MsgNotInterested:
		return "NotInterested"
	case MsgHave:
		return "Have"
	case MsgBitfield:
		return "Bitfield"
	case MsgRequest:
		return "Request"
	case MsgPiece:
		return "Piece"
	case MsgCancel:
		return "Cancel"
	default:
		return fmt.Sprintf("Unknown#%d", m.ID)
	}
}

func (m *Message) String() string {
	if m == nil {
		return m.name()
	}
	return fmt.Sprintf("%s [%d]", m.name(), len(m.Payload))
}