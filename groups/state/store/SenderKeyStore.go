package store

import (
	"github.com/StrongSquirrel/libsignal-protocol-go/groups/state/record"
	"github.com/StrongSquirrel/libsignal-protocol-go/protocol"
)

type SenderKey interface {
	StoreSenderKey(senderKeyName *protocol.SenderKeyName, keyRecord *record.SenderKey)
	LoadSenderKey(senderKeyName *protocol.SenderKeyName) *record.SenderKey
}
