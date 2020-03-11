package stream

import (
	"fmt"
	"net/url"
	"time"
)

var (
	// OptionHardDelete tells the API to do a hard delete instead of a
	// normal soft delete.
	OptionHardDelete = NewOption("hard_delete", true)

	// OptionMarkMessagesDeleted tells the API to mark all messages belonging to
	// the user as deleted in addition to deleting the user.
	OptionMarkMessagesDeleted = NewOption("mark_messages_deleted", true)
)

const (
	optionKeyType         = "type"
	optionKeyID           = "id"
	optionKeyUserID       = "user_id"
	optionKeyTargetUserID = "target_user_id"
	optionKeyTimeout      = "timeout"
	optionKeyLocation     = "url"
	optionKeyReason       = "reason"
)

func compileOptions(opts ...Option) url.Values {
	val := url.Values{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		val.Add(opt.Key(), optionAsString(opt))
	}

	return val
}

func compileOptionsMap(options []Option) map[string]interface{} {
	results := map[string]interface{}{}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		results[opt.Key()] = optionAsString(opt)
	}

	return results
}

func optionAsString(o Option) string {
	switch v := o.Value().(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Option represents a optional value that can be sent with an API call.
type Option interface {
	Key() string
	Value() interface{}
}

// NewOption is provided as a trapdoor for easily adding options that aren't
// explicitly supported by the library.
func NewOption(key string, value interface{}) Option {
	return option{key: key, value: value}
}

// makeTargetID is a helper function for making a target ID option.
func makeTargetID(targetID string) Option {
	return NewOption(optionKeyTargetUserID, targetID)
}

// makeTargetID is a helper function for making a userID option.
func makeUserID(userID string) Option {
	return NewOption(optionKeyUserID, userID)
}

// OptionTimeout returns a timeout option for banning or muting a user. If the
// time is less than one second then
func OptionTimeout(duration time.Duration) Option {
	// TODO: something here!!!! fatal or round up?
	if duration < time.Second {
		duration = time.Second
	}

	return NewOption(optionKeyTimeout, int(duration.Seconds()))
}

// OptionBanReason allows you to specify an optional reason for banning a user
// from chat.
func OptionBanReason(reason string) Option {
	return NewOption(optionKeyReason, reason)
}

// optionLocation returns an option for a location.
func optionURL(location string) Option {
	return NewOption(optionKeyLocation, location)
}

type option struct {
	key   string
	value interface{}
}

func (o option) Key() string {
	return o.key
}

func (o option) Value() interface{} {
	return o.value
}

// PaginateGreaterThan returns an Option for paginating.
func PaginateGreaterThan(id int) Option {
	return NewOption("id_gt", id)
}

// PaginateOffset returns an Option for paginating.
func PaginateOffset(id int) Option {
	return NewOption("offset", id)
}

// PaginateLimit returns an Option for paginating.
func PaginateLimit(id int) Option {
	return NewOption("limit", id)
}

// PaginateGreaterThanOrEqual returns an Option for paginating.
func PaginateGreaterThanOrEqual(id int) Option {
	return NewOption("id_gte", id)
}

// PaginateLessThan returns an Option for paginating.
func PaginateLessThan(id int) Option {
	return NewOption("id_lt", id)
}

// PaginateLessThanOrEqual returns an Option for paginating.
func PaginateLessThanOrEqual(id int) Option {
	return NewOption("id_lte", id)
}
