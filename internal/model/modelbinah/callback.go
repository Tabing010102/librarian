package modelbinah

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
)

type UploadCallBack struct {
	id           UploadCallbackID
	controlBlock *ControlBlock
}

type DownloadCallBack struct {
	id           DownloadCallbackID
	controlBlock *ControlBlock
}

func (u *UploadCallBack) GenerateUploadToken(ctx context.Context, meta FileMetadata,
	expire time.Duration) (string, error) {
	claims := libauth.FromContext(ctx)
	if claims == nil {
		return "", errors.New("token required")
	}
	return u.controlBlock.a.GenerateToken(
		claims.UserID,
		claims.PorterID,
		libauth.ClaimsTypeUploadToken,
		claims.UserType,
		expire,
		libauth.WithClaimsTransferExtra(&libauth.ClaimsTransferExtra{
			TransferMetadata: &uploadTokenPayload{
				meta,
				u.id,
			},
		}),
	)
}

func (u *DownloadCallBack) GenerateDownloadToken(ctx context.Context, meta FileMetadata,
	expire time.Duration) (string, error) {
	claims := libauth.FromContext(ctx)
	if claims == nil {
		return "", errors.New("token required")
	}
	return u.controlBlock.a.GenerateToken(
		claims.UserID,
		claims.PorterID,
		libauth.ClaimsTypeDownloadToken,
		claims.UserType,
		expire,
		libauth.WithClaimsTransferExtra(&libauth.ClaimsTransferExtra{
			TransferMetadata: &downloadTokenPayload{
				meta,
				u.id,
			},
		}),
	)
}
