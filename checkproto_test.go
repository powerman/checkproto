package checkproto_test

import (
	"testing"
	"time"

	"github.com/powerman/check"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDeepEqualProto(tt *testing.T) {
	tt.Parallel()
	t := check.T(tt)

	// Two proto messages: equal via checker (proto.Equal).
	t.DeepEqual(&emptypb.Empty{}, &emptypb.Empty{})

	// Same proto message with itself.
	now := timestamppb.Now()
	t.DeepEqual(now, now)

	// Two different proto messages.
	time.Sleep(time.Millisecond)
	later := timestamppb.Now()
	todo := t.TODO()
	todo.DeepEqual(now, later)

	// Mixed: proto vs non-proto are not equal (checker claims as unequal).
	todo.DeepEqual(now, "string")
	t.NotDeepEqual(now, "string")

	// Mixed: proto value (not pointer) is NOT detected as ProtoMessage,
	// falls through to deepequal. Both sides are the same zero value,
	// so deepequal reports equality.
	t.DeepEqual(emptypb.Empty{}, emptypb.Empty{})

	// Non-proto values still work via deepequal.
	t.DeepEqual(42, 42)
	todo.DeepEqual(42, 43)
}

func TestNotDeepEqualProto(tt *testing.T) {
	tt.Parallel()
	t := check.T(tt)

	// Two equal protos are not not-equal.
	todo := t.TODO()
	todo.NotDeepEqual(&emptypb.Empty{}, &emptypb.Empty{})

	// Two different proto messages are not equal.
	t.NotDeepEqual(timestamppb.Now(), &emptypb.Empty{})

	// Mixed: proto vs non-proto are not equal.
	t.NotDeepEqual(timestamppb.Now(), 42)

	// Non-proto values still work.
	todo.NotDeepEqual(42, 42)
	t.NotDeepEqual(42, 43)
}

func TestNonProtoValues(tt *testing.T) {
	tt.Parallel()
	t := check.T(tt)

	// Verify non-proto types are unaffected by the checker.
	t.DeepEqual(nil, nil)
	t.DeepEqual("hello", "hello")
	t.DeepEqual(42, 42)
	t.DeepEqual([]int{1, 2, 3}, []int{1, 2, 3})

	todo := t.TODO()
	todo.DeepEqual("hello", "world")
	todo.DeepEqual(42, 43)
}
