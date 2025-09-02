package models

import (
	"time"

	"github.com/google/uuid"
)

// USERS
type User struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Handle      string    `json:"handle"` // necessary ?
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	ProfileJSON string    `json:"profile_json"`
	Password    string    `json:"password"` // hashed password
}

// DEVICE
type Device struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	DeviceType string    `json:"device_type"`
	PushToken  string    `json:"push_token"`
	LastSeen   time.Time `json:"last_seen"`
}

// CHAT
type Chat struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Metadata  string    `json:"metadata"` // jsonb as string
}

// CHAT_MEMBER
type ChatMember struct {
	ChatID   uuid.UUID `json:"chat_id"`
	UserID   uuid.UUID `json:"user_id"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
	LeftAt   time.Time `json:"left_at"`
}

// MESSAGE
type Message struct {
	ID               uuid.UUID `gorm:"primaryKey" json:"id"`
	ChatID           uuid.UUID `json:"chat_id"`
	SenderDeviceID   uuid.UUID `json:"sender_device_id"`
	Ciphertext       []byte    `json:"ciphertext"`
	ContentType      string    `json:"content_type"`
	ExtMeta          string    `json:"ext_meta"` // jsonb as string
	SentAt           time.Time `json:"sent_at"`
	ServerReceivedAt time.Time `json:"server_received_at"`
	ExpiresAt        time.Time `json:"expires_at"`
}

// DELIVERY
type Delivery struct {
	MessageID uuid.UUID `json:"message_id"`
	DeviceID  uuid.UUID `json:"device_id"`
	Status    string    `json:"status"`
	TS        time.Time `json:"ts"`
}

// REACTION
type Reaction struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	MessageID uuid.UUID `json:"message_id"`
	UserID    uuid.UUID `json:"user_id"`
	Emoji     string    `json:"emoji"`
	TS        time.Time `json:"ts"`
}

// MEDIA
type Media struct {
	ID               uuid.UUID `gorm:"primaryKey" json:"id"`
	UploaderDeviceID uuid.UUID `json:"uploader_device_id"`
	BlobURL          string    `json:"blob_url"`
	ThumbURL         string    `json:"thumb_url"`
	Bytes            int       `json:"bytes"`
	SHA256           string    `json:"sha256"`
	Mime             string    `json:"mime"`
	CreatedAt        time.Time `json:"created_at"`
}

// MESSAGE_MEDIA
type MessageMedia struct {
	MessageID uuid.UUID `json:"message_id"`
	MediaID   uuid.UUID `json:"media_id"`
}

// IDENTITY_KEY
type IdentityKey struct {
	DeviceID  uuid.UUID `json:"device_id"`
	PublicKey []byte    `json:"public_key"`
	CreatedAt time.Time `json:"created_at"`
}

// PREKEY
type PreKey struct {
	DeviceID  uuid.UUID `json:"device_id"`
	PrekeyID  int       `json:"prekey_id"`
	PublicKey []byte    `json:"public_key"`
	Consumed  bool      `json:"consumed"`
	CreatedAt time.Time `json:"created_at"`
}
