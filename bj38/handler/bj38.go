package handler

import (
	"context"
	pb "golandProject/bj38/pb/bj38"
)

type SayHandler struct{}

// Hello 实现 proto 中的 Say.Hello
func (h *SayHandler) Hello(
	ctx context.Context,
	req *pb.Request,
	rsp *pb.Response,
) error {

	rsp.Message = "hello " + req.Name
	return nil
}
