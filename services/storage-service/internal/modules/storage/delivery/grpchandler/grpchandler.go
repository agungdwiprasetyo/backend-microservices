// Code generated by candi v1.3.1.

package grpchandler

import (
	"bytes"
	"errors"
	"io"
	proto "monorepo/sdk/storage-service/proto/storage"
	"monorepo/services/storage-service/internal/modules/storage/domain"
	"monorepo/services/storage-service/internal/modules/storage/usecase"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"pkg.agungdp.dev/candi/codebase/factory/types"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

// GRPCHandler rpc handler
type GRPCHandler struct {
	mw        interfaces.Middleware
	uc        usecase.StorageUsecase
	validator interfaces.Validator
}

// NewGRPCHandler func
func NewGRPCHandler(mw interfaces.Middleware, uc usecase.StorageUsecase, validator interfaces.Validator) *GRPCHandler {
	return &GRPCHandler{
		mw: mw, uc: uc, validator: validator,
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server, mwGroup *types.MiddlewareGroup) {
	proto.RegisterStorageServiceServer(server, h)

	// register middleware for method
	mwGroup.AddProto(proto.File_storage_storage_proto, "Upload", h.mw.GRPCBearerAuth)
}

// Upload method
func (h *GRPCHandler) Upload(stream proto.StorageService_UploadServer) (err error) {
	trace := tracer.StartTrace(stream.Context(), "StorageDeliveryGRPC:Upload")
	defer trace.Finish()
	ctx := trace.Context()

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	fields := meta.Get("filename")
	if len(fields) == 0 {
		return errors.New("missing filename field")
	}
	fileName := fields[0]

	fields = meta.Get("folder")
	if len(fields) == 0 {
		return errors.New("missing folder field")
	}
	folder := fields[0]

	var contentType string
	if u := meta.Get("contentType"); len(u) > 0 {
		contentType = u[0]
	}

	var size int64
	if u := meta.Get("size"); len(u) > 0 {
		s, _ := strconv.Atoi(u[0])
		size = int64(s)
	}

	buff := new(bytes.Buffer)
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		buff.Write(res.Content)
	}

	err = h.uc.Upload(ctx,
		buff,
		&domain.UploadMetadata{
			ContentType: contentType,
			FileSize:    size,
			Filename:    fileName,
		})
	if err != nil {
		return grpc.Errorf(codes.Internal, "%v", err)
	}

	return stream.SendAndClose(&proto.UploadStatus{
		Message: "Stream file success",
		Code:    proto.StatusCode_Ok,
		File:    "url" + "/" + folder + "/" + fileName,
	})
}
