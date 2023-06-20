package entity

const (
	TemplateRepoServiceCallImpl = `
func (s *serviceImpl) %s(
	ctx context.Context,
	req *pb.%sRequest,
) (*pb.%sResponse, error) {
	logSugar := s.logger.With(logtool.WithCtx(ctx)).Sugar()
	logSugar.Infow("call rpc func [%s] begin", "req", req)
	resp, err := s.client.%s(ctx, req)
	if err != nil {
		logSugar.Warnw("%s 失败", "err", err)
		return nil, err
	}
	logSugar.Infow("call rpc func [%s] end", "resp", resp)
	if res.Code != pb.RetCode_OK {
		return nil, errors.New("rpc调用失败")
	}
	return resp, err
}
`
	TemplateRepoSampleFuncImpl = `
func (s serviceImpl) %s(ctx context.Context) error {
	logSugar := s.logger.With(logtool.WithCtx(ctx)).Sugar()
	logSugar.Infow("call %s begin")

	logSugar.Infow("call %s end")
	return nil
}
`
)
