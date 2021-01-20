// Code generated by candi v1.3.1.

package usecase

import (
	"bytes"
	"context"
	"fmt"

	"monorepo/services/storage-service/internal/modules/storage/domain"
	"monorepo/services/storage-service/pkg/shared/repository"

	"github.com/minio/minio-go/v6"
	"pkg.agungdwiprasetyo.com/candi/candishared"
	"pkg.agungdwiprasetyo.com/candi/codebase/factory/dependency"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
	"pkg.agungdwiprasetyo.com/candi/logger"
	"pkg.agungdwiprasetyo.com/candi/tracer"
)

type storageUsecaseImpl struct {
	cache interfaces.Cache

	repoMongo   *repository.RepoMongo
	minioClient *minio.Client
}

// NewStorageUsecase usecase impl constructor
func NewStorageUsecase(deps dependency.Dependency) StorageUsecase {
	return &storageUsecaseImpl{
		cache: deps.GetRedisPool().Cache(),

		repoMongo: repository.GetSharedRepoMongo(),
	}
}

func (uc *storageUsecaseImpl) Hello(ctx context.Context) (msg string) {
	trace := tracer.StartTrace(ctx, "StorageUsecase:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	msg, _ = uc.repoMongo.StorageRepo.FindHello(ctx)
	return
}

func (uc *storageUsecaseImpl) Upload(ctx context.Context, buff []byte, metadata *domain.UploadMetadata) <-chan candishared.Result {
	output := make(chan candishared.Result)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				output <- candishared.Result{Error: fmt.Errorf("%v", r)}
			}
			close(output)
		}()

		n, err := uc.minioClient.PutObject("tong", metadata.Filename, bytes.NewReader(buff), -1,
			minio.PutObjectOptions{ContentType: metadata.ContentType})
		if err != nil {
			logger.LogE(err.Error())
			panic(err)
		}

		fmt.Println("Uploaded", " size: ", n, "Successfully.", "localhost:9000/tong/...")
	}()

	return output
}