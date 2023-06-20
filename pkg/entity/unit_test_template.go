package entity

const (
	TemplateServiceUnitTestServiceFuncImpl = `
func Test_svcImpl_%s(t *testing.T) {
    type args struct {
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
			got, err := s.%s(testCtx, tt.args.req)
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
			got := s.%s(testCtx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s() got = %%v, want %%v", got, tt.want)
			}
		})
	}
}
`
)
