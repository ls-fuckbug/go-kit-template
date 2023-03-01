package entity

const (
	TemplateRepoServiceCallImpl = `
func (s *serviceImpl) %s(
	ctx context.Context,
	req *pb.%sRequest,
) (*pb.%sResponse, error) {
	logSugar := p.logger.With(logtool.WithCtx(ctx)).Sugar()
	logSugar.Infow("call rpc func [%s] begin", "req", req)
	resp, err := s.client.%s(ctx, req)
	if err != nil {
		logSugar.Warnw("%s 失败", "err", err)
		return nil, err
	}
	logSugar.Infow("call rpc func [%s] end", "resp", resp)
	return resp, err
}
`
	TemplateRepoSampleFuncImpl = `
func (s serviceImpl) %s(ctx context.Context) error {
	logSugar := r.logger.With(logtool.WithCtx(ctx)).Sugar()
	logSugar.Infow("call %s begin")

	logSugar.Infow("call %s end")
	return nil
}
`
)
