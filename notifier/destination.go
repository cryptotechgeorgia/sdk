package notifier

import "fmt"

// The set of types that can be used.
var (
	Telegram = newType("telegram")
	SMS      = newType("sms")
	Email    = newType("email")
)

// Set of known destination types.
var destinationTypes = make(map[string]DestinationType)

type DestinationType struct {
	value string
}

func newType(dstType string) DestinationType {
	dt := DestinationType{dstType}
	destinationTypes[dstType] = dt
	return dt
}

// name of the type.
func (dt DestinationType) String() string {
	return dt.value
}

// Equal provides support for the go-cmp package and testing.
func (dt DestinationType) Equal(ht2 DestinationType) bool {
	return dt.value == ht2.value
}

// for logging
func (ht DestinationType) MarshalText() ([]byte, error) {
	return []byte(ht.value), nil
}

func Parse(value string) (DestinationType, error) {
	typ, exists := destinationTypes[value]
	if !exists {
		return DestinationType{}, fmt.Errorf("invalid destination type %q", value)
	}

	return typ, nil
}
