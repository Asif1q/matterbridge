// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Conversation undocumented
type Conversation struct {
	// Entity is the base model of Conversation
	Entity
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// LastDeliveredDateTime undocumented
	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`
	// UniqueSenders undocumented
	UniqueSenders []string `json:"uniqueSenders,omitempty"`
	// Preview undocumented
	Preview *string `json:"preview,omitempty"`
	// Threads undocumented
	Threads []ConversationThread `json:"threads,omitempty"`
}

// ConversationMember undocumented
type ConversationMember struct {
	// Entity is the base model of ConversationMember
	Entity
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// ConversationMemberRoles undocumented
type ConversationMemberRoles struct {
	// Object is the base model of ConversationMemberRoles
	Object
}

// ConversationThread undocumented
type ConversationThread struct {
	// Entity is the base model of ConversationThread
	Entity
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"toRecipients,omitempty"`
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// LastDeliveredDateTime undocumented
	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`
	// UniqueSenders undocumented
	UniqueSenders []string `json:"uniqueSenders,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// Preview undocumented
	Preview *string `json:"preview,omitempty"`
	// IsLocked undocumented
	IsLocked *bool `json:"isLocked,omitempty"`
	// Posts undocumented
	Posts []Post `json:"posts,omitempty"`
}
