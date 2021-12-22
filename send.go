package module

import (
	"context"
)

// Sender ...
// @Description:
type Sender interface {
	SendRequest(ctx context.Context, id ID, msg []byte) error
	AnswerResponse(ctx context.Context, id ID, msg []byte) error
}
