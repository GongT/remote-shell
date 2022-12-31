package action_base

type Message interface {
	Handle() (bool, error)
	GetTypeId() uint32
	GetId() uint32
	SetId(uint32)
}

type MessageBase struct {
	TypeId uint32 `json:"type"`
	Id     uint32 `json:"id"`
}

func (mb *MessageBase) GetTypeId() uint32 { return mb.TypeId }
func (mb *MessageBase) GetId() uint32     { return mb.Id }
func (mb *MessageBase) SetId(u uint32)    { mb.Id = u }

const InvalidMessage = 0

const (
	TypeCallback = 1
	TypeOpen     = 2
	TypeUrl      = 3
	TypeMagnet   = 4
)
