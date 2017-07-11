package gcsstore

//redefing methods that are delegated for tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tus/tusd/gcsstore"
)

func NewTestGCSService(filename string) (*GCSService, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccount.File((filename)))
	if err != nil {
		return nil, err
	}

	service := &GCSService{
		Client: client,
		Ctx:    ctx,
		getObjectAttrs: func(params GCSObjectParams) (*storage.ObjectAttrs, error) {
			return nil, nil
		},
		ReadObject: func(params GCSObjectParams) (GCSReader, error) {
			return nil, nil
		},
		SetObjectMetadata: func(params GCSObjectParams, metadata map[string]string) error {
			return nil
		},

		DeleteObject: func(params GCSObjectParams) error {
			return nil
		},
		WriteObject: func(params GCSObjectParams, r io.Reader) (int64, error) {
			return 0, nil
		},
	}

	return service, nil
}

func TestGCSCompose(t *testing.T) {
	service = NewTestGCSService("testing")
	//service.compose("bucket", ["t1", "t2", "t3"], "destination")
}
