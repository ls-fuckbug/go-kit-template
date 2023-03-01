package entity

const (
	BackendTemplateServiceImpl = `
// %s 
func (s svcImpl) %s(
  ctx context.Context, 
  req *entity.%sRequest,
) *entity.%sResponse {

    return &entity.%sResponse{}
}

`

	ServiceTemplateServiceImpl = `
// %s 
func (s svcImpl) %s(
  ctx context.Context, 
  req *pb.%sRequest,
) (*pb.%sResponse, error) {

    return &pb.%sResponse{}, nil
}

`
	ServiceTemplateServiceInterface = `
// %s 
%s(ctx context.Context, req *pb.%sRequest) (*pb.%sResponse, error)`
	BackendTemplateServiceInterface = `
// %s 
%s(ctx context.Context, req *entity.%sRequest) *entity.%sResponse`
)
