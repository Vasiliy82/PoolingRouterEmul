package utils

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type ctxKey int8

const (
	CtxKeyReqId ctxKey = iota
)

var ErrNoReqID = errors.New("requestId not found")

func NewRequestID() string {
	return uuid.NewString()
}

func SetRequestID(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, CtxKeyReqId, requestId)
}

func GetRequestID(ctx context.Context) (string, error) {
	reqId, ok := ctx.Value(CtxKeyReqId).(string)
	if !ok {
		return "", ErrNoReqID
	}

	return reqId, nil
}
