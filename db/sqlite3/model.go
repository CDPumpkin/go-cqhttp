//go:build !((windows || arm) && (windows || arm64))
// +build !windows,!arm !windows,!arm64

package sqlite3

const (
	Sqlite3GroupMessageTableName          = "grpmsg"
	Sqlite3MessageAttributeTableName      = "msgattr"
	Sqlite3GuildMessageAttributeTableName = "gmsgattr"
	Sqlite3QuotedInfoTableName            = "quoinf"
	Sqlite3PrivateMessageTableName        = "privmsg"
	Sqlite3GuildChannelMessageTableName   = "guildmsg"
)

// StoredMessageAttribute 持久化消息属性
type StoredMessageAttribute struct {
	ID         int64 // ID is the crc64 of 字段s below
	MessageSeq int32
	InternalID int32
	SenderUin  int64
	SenderName string
	Timestamp  int64
}

// StoredGuildMessageAttribute 持久化频道消息属性
type StoredGuildMessageAttribute struct {
	ID           int64 // ID is the crc64 of 字段s below
	MessageSeq   int64
	InternalID   int64
	SenderTinyID int64
	SenderName   string
	Timestamp    int64
}

// QuotedInfo 引用回复
type QuotedInfo struct {
	ID            int64 // ID is the crc64 of 字段s below
	PrevID        string
	PrevGlobalID  int32
	QuotedContent string // QuotedContent is json of original content
}

// StoredGroupMessage 持久化群消息
type StoredGroupMessage struct {
	GlobalID     int32
	ID           string
	AttributeID  int64
	SubType      string
	QuotedInfoID int64
	GroupCode    int64
	AnonymousID  string
	Content      string // Content is json of original content
}

// StoredPrivateMessage 持久化私聊消息
type StoredPrivateMessage struct {
	GlobalID     int32
	ID           string
	AttributeID  int64
	SubType      string
	QuotedInfoID int64
	SessionUin   int64
	TargetUin    int64
	Content      string // Content is json of original content
}

// StoredGuildChannelMessage 持久化频道消息
type StoredGuildChannelMessage struct {
	ID           string
	AttributeID  int64
	GuildID      int64
	ChannelID    int64
	QuotedInfoID int64
	Content      string // Content is json of original content
}
