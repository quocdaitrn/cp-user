package codec

import (
	"context"
	"net/http"

	kithttp "github.com/quocdaitrn/golang-kit/http"

	"github.com/quocdaitrn/cp-user/domain/service"
)

// DecodeGetCurrentUserProfileRequest decodes GetCurrentUserProfileRequest from http.Request.
func DecodeGetCurrentUserProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &service.GetCurrentUserProfileRequest{}
	if err := kithttp.Bind(r, req); err != nil {
		return nil, err
	}
	return req, nil
}
