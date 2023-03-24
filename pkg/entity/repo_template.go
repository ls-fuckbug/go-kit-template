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
	TemplateServiceUnitTestServiceFuncImpl = `
func Test_svcImpl_%s(t *testing.T) {
    type args struct {
		ctx context.Context
		req *pb.%sRequest
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		args    args
		want    *pb.%sResponse
		wantErr bool
	}{
      {
      },
    }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &svcImpl{
				logger:     testFields.logger,
			}
			got, err := s.%s(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s() error = %%v, wantErr %%v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s() got = %%v, want %%v", got, tt.want)
			}
		})
	}
}
`
	TemplateBackendUnitTestServiceFuncImpl = `
func Test_svcImpl_%s(t *testing.T) {
    type args struct {
		ctx context.Context
		req *entity.%sReq
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		args    args
		want    *entity.%sResp
	}{
      {
      },
    }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &svcImpl{
				logger:     testFields.logger,
			}
			got := s.%s(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s() got = %%v, want %%v", got, tt.want)
			}
		})
	}
}
`
)
