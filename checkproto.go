// Package checkproto provides protobuf message comparison for github.com/powerman/check.
//
// Import this package as a blank import in your test file or TestMain
// to enable proto.Equal comparison via check.DeepEqual/NotDeepEqual:
//
//	import _ "github.com/powerman/checkproto"
package checkproto

import (
	"github.com/powerman/check"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoinits // Required to register EqualChecker.
func init() {
	check.RegisterEqualChecker(EqualProto)
}

// EqualProto compares two values using [proto.Equal]
// when both implement [protoreflect.ProtoMessage].
// If only one side is a proto message they are reported as unequal (claimed).
// Otherwise the checker does not apply.
func EqualProto(actual, expected any) (equal, ok bool) {
	msg1, ok1 := actual.(protoreflect.ProtoMessage)
	msg2, ok2 := expected.(protoreflect.ProtoMessage)
	switch {
	case ok1 && ok2:
		return proto.Equal(msg1, msg2), true
	case ok1 || ok2:
		return false, true
	default:
		return false, false
	}
}
