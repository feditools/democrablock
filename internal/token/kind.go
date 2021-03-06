package token

// Kind represents the kind of model to encode a token for.
type Kind int64

// This order can not change else all external urls with tokens will become invalid.
const (
	// KindFediInstance is a token that represents a federated social instance.
	KindFediInstance Kind = 1 + iota
	// KindFediAccount is a token that represents a federated social account.
	KindFediAccount
	// KindBlock is a token that represents a blocked federated social instance.
	KindBlock
)

func (k Kind) String() string {
	switch k {
	case KindFediInstance:
		return "FediInstance"
	case KindFediAccount:
		return "FediAccount"
	case KindBlock:
		return "Block"
	default:
		return "unknown"
	}
}
