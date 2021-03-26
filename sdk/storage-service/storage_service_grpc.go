package storageservice

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	proto "monorepo/sdk/storage-service/proto/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/metadata"
	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/logger"
	"pkg.agungdp.dev/candi/tracer"
)

const defaultStreamLimitSize = 50 * candihelper.MByte

type storageGRPCImpl struct {
	client          proto.StorageServiceClient
	authKey         string
	streamLimitSize int64
}

// NewStorageServiceGRPC constructor for storage service GRPC stream
func NewStorageServiceGRPC(host string, authKey string, streamLimitSize uint64) (StorageService, error) {
	conn, err := grpc.Dial(host,
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  50 * time.Millisecond,
				Multiplier: 5,
				MaxDelay:   50 * time.Millisecond,
			},
			MinConnectTimeout: 1 * time.Second,
		}),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(int(100*candihelper.MByte)),
			grpc.MaxCallSendMsgSize(int(100*candihelper.MByte)),
		),
	)
	if err != nil {
		return nil, err
	}

	if streamLimitSize <= 0 {
		streamLimitSize = defaultStreamLimitSize
	}

	return &storageGRPCImpl{
		client:          proto.NewStorageServiceClient(conn),
		authKey:         authKey,
		streamLimitSize: int64(streamLimitSize),
	}, nil
}

func (u *storageGRPCImpl) Upload(ctx context.Context, file io.Reader, header Header) (res Response, err error) {
	trace := tracer.StartTrace(ctx, "StorageGRPCClient:Upload")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
		trace.Finish()
	}()
	ctx = trace.Context()

	md := metadata.Pairs("authorization", u.authKey,
		"filename", header.Filename,
		"folder", header.Folder,
		"contentType", header.ContentType,
		"size", strconv.Itoa(int(header.Size)))
	trace.InjectGRPCMetadata(md)

	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := u.client.Upload(ctx)
	if err != nil {
		logger.LogE(err.Error())
		return res, err
	}
	defer stream.CloseSend()

	i := int64(0)
	buff := make([]byte, u.streamLimitSize)
	// stream send file with grpc
	for {
		n, err := file.Read(buff)
		if err != nil && err != io.EOF {
			return res, err
		}
		if n == 0 {
			break
		}

		i += u.streamLimitSize
		if err = stream.Send(&proto.Chunk{
			Content:   buff,
			TotalSize: header.Size,
			Received:  i,
		}); err != nil {
			logger.LogE(err.Error())
			return res, err
		}
	}

	status, err := stream.CloseAndRecv()
	if err != nil {
		logger.LogE(err.Error())
		return res, err
	}

	if status.Code != proto.StatusCode_Ok {
		logger.LogE("not success")
		return res, errors.New("Status code is not success")
	}

	res = Response{
		Location: status.File, Size: status.Size,
	}
	return
}
