package entity

const (
	TemplateServiceImpl = `

// %s 
func (s svcImpl) %s(
  ctx context.Context, 
  req *pb.%sRequest,
) (*pb.%sResponse, error) {

}

`
	TemplateServiceInterface = `
// %s 
%s(ctx context.Context, req *pb.%sRequest) (*pb.%sResponse, error)
`

	BackendTemplateEndpointMake = `

func make%sEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(*entity.%sReq)
		resp := s.%s(ctx, req)
		return resp, nil
	}
}

`
	ServiceTemplateEndpointMake = `

func make%sEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(*pb.%sReq)
		resp := s.%s(ctx, req)
		return resp, nil
	}
}

`
	TemplateEndpointMakeCall = `
%s:         make%sEndpoint(svc),
`
	TemplateEndpointStructMember = `
%s endpoint.Endpoint
`
	BackendTemplateTransportHandler = `
	%sHandler := httptransport.NewServer(
		endpoints.%s,
		decodeHTTP%sRequest,
		encodeHTTP%sResponse,
		options...,
	)
`

	ServiceTemplateTransportHandler = `
	%sTransport := grpctransport.NewServer(
		endpoints.%s,
		decode%sRequest,
		encode%sResponse,
	)
`

	BackendTemplateTransportHandle = `
	r.Handle("/", %sHandler).Methods("")
`

	ServiceTemplateTransportHandle = `
		%s:         %sTransport,
`
	BackendTemplateTransportImpl = `

func decodeHTTP%sRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request entity.%sReq
	if e := json.NewDecoder(r.Body).Decode(&request); e != nil {
		return nil, e
	}
	return &request, nil
}

func encodeHTTP%sResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp, _ := response.(*entity.%s)
	middleware.LogAndAddRequestID(ctx, &resp.CommonRet)
	return json.NewEncoder(w).Encode(resp)
}
`

	ServiceTemplateTransportImpl = `

func (n *grpcServer)%s(ctx context.Context,
	req *pb.%sRequest) (*pb.%sResponse, error) {
	_, resp, err := n.%s.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.%sResponse), nil
}

func decode%sRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, _ := request.(*pb.%sRequest)
	return req, nil
}

func encode%sResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, _ := response.(*pb.%sResponse)
	return resp, nil
}
`
)
