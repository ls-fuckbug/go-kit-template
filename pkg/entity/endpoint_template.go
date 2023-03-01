package entity

const (
	BackendTemplateEndpointMake = `
func make%sEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(*entity.%sRequest)
		resp := s.%s(ctx, req)
		return resp, nil
	}
}
`

	ServiceTemplateEndpointMake = `
func make%sEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, _ := request.(*pb.%sRequest)
		resp, err := s.%s(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
`

	TemplateEndpointMakeCall = `
%s:         make%sEndpoint(svc),`

	TemplateEndpointStructMember = `
%s endpoint.Endpoint`
)
