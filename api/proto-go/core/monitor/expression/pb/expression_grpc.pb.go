// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: expression.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// ExpressionServiceClient is the client API for ExpressionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpressionServiceClient interface {
	GetAlertExpressions(ctx context.Context, in *GetExpressionsRequest, opts ...grpc.CallOption) (*GetExpressionsResponse, error)
	GetMetricExpressions(ctx context.Context, in *GetMetricExpressionsRequest, opts ...grpc.CallOption) (*GetMetricExpressionsResponse, error)
	GetAlertNotifies(ctx context.Context, in *GetAlertNotifiesRequest, opts ...grpc.CallOption) (*GetAlertNotifiesResponse, error)
	GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*GetTemplatesResponse, error)
	GetOrgsLocale(ctx context.Context, in *GetOrgsLocaleRequest, opts ...grpc.CallOption) (*GetOrgsLocaleResponse, error)
}

type expressionServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewExpressionServiceClient(cc grpc1.ClientConnInterface) ExpressionServiceClient {
	return &expressionServiceClient{cc}
}

func (c *expressionServiceClient) GetAlertExpressions(ctx context.Context, in *GetExpressionsRequest, opts ...grpc.CallOption) (*GetExpressionsResponse, error) {
	out := new(GetExpressionsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.expression.ExpressionService/GetAlertExpressions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) GetMetricExpressions(ctx context.Context, in *GetMetricExpressionsRequest, opts ...grpc.CallOption) (*GetMetricExpressionsResponse, error) {
	out := new(GetMetricExpressionsResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.expression.ExpressionService/GetMetricExpressions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) GetAlertNotifies(ctx context.Context, in *GetAlertNotifiesRequest, opts ...grpc.CallOption) (*GetAlertNotifiesResponse, error) {
	out := new(GetAlertNotifiesResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.expression.ExpressionService/GetAlertNotifies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*GetTemplatesResponse, error) {
	out := new(GetTemplatesResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.expression.ExpressionService/GetTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressionServiceClient) GetOrgsLocale(ctx context.Context, in *GetOrgsLocaleRequest, opts ...grpc.CallOption) (*GetOrgsLocaleResponse, error) {
	out := new(GetOrgsLocaleResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.expression.ExpressionService/GetOrgsLocale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressionServiceServer is the server API for ExpressionService service.
// All implementations should embed UnimplementedExpressionServiceServer
// for forward compatibility
type ExpressionServiceServer interface {
	GetAlertExpressions(context.Context, *GetExpressionsRequest) (*GetExpressionsResponse, error)
	GetMetricExpressions(context.Context, *GetMetricExpressionsRequest) (*GetMetricExpressionsResponse, error)
	GetAlertNotifies(context.Context, *GetAlertNotifiesRequest) (*GetAlertNotifiesResponse, error)
	GetTemplates(context.Context, *GetTemplatesRequest) (*GetTemplatesResponse, error)
	GetOrgsLocale(context.Context, *GetOrgsLocaleRequest) (*GetOrgsLocaleResponse, error)
}

// UnimplementedExpressionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExpressionServiceServer struct {
}

func (*UnimplementedExpressionServiceServer) GetAlertExpressions(context.Context, *GetExpressionsRequest) (*GetExpressionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertExpressions not implemented")
}
func (*UnimplementedExpressionServiceServer) GetMetricExpressions(context.Context, *GetMetricExpressionsRequest) (*GetMetricExpressionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricExpressions not implemented")
}
func (*UnimplementedExpressionServiceServer) GetAlertNotifies(context.Context, *GetAlertNotifiesRequest) (*GetAlertNotifiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertNotifies not implemented")
}
func (*UnimplementedExpressionServiceServer) GetTemplates(context.Context, *GetTemplatesRequest) (*GetTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplates not implemented")
}
func (*UnimplementedExpressionServiceServer) GetOrgsLocale(context.Context, *GetOrgsLocaleRequest) (*GetOrgsLocaleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgsLocale not implemented")
}

func RegisterExpressionServiceServer(s grpc1.ServiceRegistrar, srv ExpressionServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ExpressionService_serviceDesc(srv, opts...), srv)
}

var _ExpressionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.monitor.expression.ExpressionService",
	HandlerType: (*ExpressionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "expression.proto",
}

func _get_ExpressionService_serviceDesc(srv ExpressionServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ExpressionService_GetAlertExpressions_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetAlertExpressions(ctx, req.(*GetExpressionsRequest))
	}
	var _ExpressionService_GetAlertExpressions_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ExpressionService_GetAlertExpressions_info = transport.NewServiceInfo("erda.core.monitor.expression.ExpressionService", "GetAlertExpressions", srv)
		_ExpressionService_GetAlertExpressions_Handler = h.Interceptor(_ExpressionService_GetAlertExpressions_Handler)
	}

	_ExpressionService_GetMetricExpressions_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetMetricExpressions(ctx, req.(*GetMetricExpressionsRequest))
	}
	var _ExpressionService_GetMetricExpressions_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ExpressionService_GetMetricExpressions_info = transport.NewServiceInfo("erda.core.monitor.expression.ExpressionService", "GetMetricExpressions", srv)
		_ExpressionService_GetMetricExpressions_Handler = h.Interceptor(_ExpressionService_GetMetricExpressions_Handler)
	}

	_ExpressionService_GetAlertNotifies_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetAlertNotifies(ctx, req.(*GetAlertNotifiesRequest))
	}
	var _ExpressionService_GetAlertNotifies_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ExpressionService_GetAlertNotifies_info = transport.NewServiceInfo("erda.core.monitor.expression.ExpressionService", "GetAlertNotifies", srv)
		_ExpressionService_GetAlertNotifies_Handler = h.Interceptor(_ExpressionService_GetAlertNotifies_Handler)
	}

	_ExpressionService_GetTemplates_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTemplates(ctx, req.(*GetTemplatesRequest))
	}
	var _ExpressionService_GetTemplates_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ExpressionService_GetTemplates_info = transport.NewServiceInfo("erda.core.monitor.expression.ExpressionService", "GetTemplates", srv)
		_ExpressionService_GetTemplates_Handler = h.Interceptor(_ExpressionService_GetTemplates_Handler)
	}

	_ExpressionService_GetOrgsLocale_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetOrgsLocale(ctx, req.(*GetOrgsLocaleRequest))
	}
	var _ExpressionService_GetOrgsLocale_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ExpressionService_GetOrgsLocale_info = transport.NewServiceInfo("erda.core.monitor.expression.ExpressionService", "GetOrgsLocale", srv)
		_ExpressionService_GetOrgsLocale_Handler = h.Interceptor(_ExpressionService_GetOrgsLocale_Handler)
	}

	var serviceDesc = _ExpressionService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetAlertExpressions",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetExpressionsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ExpressionServiceServer).GetAlertExpressions(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ExpressionService_GetAlertExpressions_info)
				}
				if interceptor == nil {
					return _ExpressionService_GetAlertExpressions_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.expression.ExpressionService/GetAlertExpressions",
				}
				return interceptor(ctx, in, info, _ExpressionService_GetAlertExpressions_Handler)
			},
		},
		{
			MethodName: "GetMetricExpressions",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetMetricExpressionsRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ExpressionServiceServer).GetMetricExpressions(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ExpressionService_GetMetricExpressions_info)
				}
				if interceptor == nil {
					return _ExpressionService_GetMetricExpressions_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.expression.ExpressionService/GetMetricExpressions",
				}
				return interceptor(ctx, in, info, _ExpressionService_GetMetricExpressions_Handler)
			},
		},
		{
			MethodName: "GetAlertNotifies",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetAlertNotifiesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ExpressionServiceServer).GetAlertNotifies(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ExpressionService_GetAlertNotifies_info)
				}
				if interceptor == nil {
					return _ExpressionService_GetAlertNotifies_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.expression.ExpressionService/GetAlertNotifies",
				}
				return interceptor(ctx, in, info, _ExpressionService_GetAlertNotifies_Handler)
			},
		},
		{
			MethodName: "GetTemplates",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTemplatesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ExpressionServiceServer).GetTemplates(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ExpressionService_GetTemplates_info)
				}
				if interceptor == nil {
					return _ExpressionService_GetTemplates_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.expression.ExpressionService/GetTemplates",
				}
				return interceptor(ctx, in, info, _ExpressionService_GetTemplates_Handler)
			},
		},
		{
			MethodName: "GetOrgsLocale",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetOrgsLocaleRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ExpressionServiceServer).GetOrgsLocale(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ExpressionService_GetOrgsLocale_info)
				}
				if interceptor == nil {
					return _ExpressionService_GetOrgsLocale_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.expression.ExpressionService/GetOrgsLocale",
				}
				return interceptor(ctx, in, info, _ExpressionService_GetOrgsLocale_Handler)
			},
		},
	}
	return &serviceDesc
}
