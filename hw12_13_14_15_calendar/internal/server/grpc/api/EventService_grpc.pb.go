// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: EventService.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EventService_NewEvent_FullMethodName    = "/event.EventService/newEvent"
	EventService_GetEvent_FullMethodName    = "/event.EventService/getEvent"
	EventService_UpdateEvent_FullMethodName = "/event.EventService/updateEvent"
	EventService_DeleteEvent_FullMethodName = "/event.EventService/deleteEvent"
	EventService_All_FullMethodName         = "/event.EventService/all"
	EventService_Next_FullMethodName        = "/event.EventService/next"
	EventService_Day_FullMethodName         = "/event.EventService/day"
	EventService_Week_FullMethodName        = "/event.EventService/week"
	EventService_Month_FullMethodName       = "/event.EventService/month"
	EventService_SinceDay_FullMethodName    = "/event.EventService/sinceDay"
	EventService_SinceWeek_FullMethodName   = "/event.EventService/sinceWeek"
	EventService_SinceMonth_FullMethodName  = "/event.EventService/sinceMonth"
)

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	NewEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventIdResponse, error)
	GetEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	All(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Next(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Day(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Week(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Month(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error)
	SinceDay(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*EventResponse, error)
	SinceWeek(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*EventResponse, error)
	SinceMonth(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) NewEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventIdResponse)
	err := c.cc.Invoke(ctx, EventService_NewEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_GetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EventService_UpdateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) DeleteEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EventService_DeleteEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) All(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_All_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Next(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_Next_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Day(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_Day_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Week(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_Week_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Month(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_Month_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) SinceDay(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_SinceDay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) SinceWeek(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_SinceWeek_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) SinceMonth(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventService_SinceMonth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	NewEvent(context.Context, *Event) (*EventIdResponse, error)
	GetEvent(context.Context, *EventRequest) (*EventResponse, error)
	UpdateEvent(context.Context, *Event) (*emptypb.Empty, error)
	DeleteEvent(context.Context, *EventRequest) (*emptypb.Empty, error)
	All(context.Context, *UserRequest) (*EventResponse, error)
	Next(context.Context, *UserRequest) (*EventResponse, error)
	Day(context.Context, *UserRequest) (*EventResponse, error)
	Week(context.Context, *UserRequest) (*EventResponse, error)
	Month(context.Context, *UserRequest) (*EventResponse, error)
	SinceDay(context.Context, *ListRequest) (*EventResponse, error)
	SinceWeek(context.Context, *ListRequest) (*EventResponse, error)
	SinceMonth(context.Context, *ListRequest) (*EventResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) NewEvent(context.Context, *Event) (*EventIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEvent not implemented")
}
func (UnimplementedEventServiceServer) GetEvent(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventServiceServer) UpdateEvent(context.Context, *Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventServiceServer) DeleteEvent(context.Context, *EventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventServiceServer) All(context.Context, *UserRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedEventServiceServer) Next(context.Context, *UserRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedEventServiceServer) Day(context.Context, *UserRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Day not implemented")
}
func (UnimplementedEventServiceServer) Week(context.Context, *UserRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Week not implemented")
}
func (UnimplementedEventServiceServer) Month(context.Context, *UserRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Month not implemented")
}
func (UnimplementedEventServiceServer) SinceDay(context.Context, *ListRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SinceDay not implemented")
}
func (UnimplementedEventServiceServer) SinceWeek(context.Context, *ListRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SinceWeek not implemented")
}
func (UnimplementedEventServiceServer) SinceMonth(context.Context, *ListRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SinceMonth not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_NewEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).NewEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_NewEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).NewEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).DeleteEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).All(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_Next_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Next(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Day_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Day(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_Day_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Day(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Week_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Week(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_Week_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Week(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Month_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Month(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_Month_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Month(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_SinceDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SinceDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_SinceDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SinceDay(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_SinceWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SinceWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_SinceWeek_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SinceWeek(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_SinceMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SinceMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_SinceMonth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SinceMonth(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "newEvent",
			Handler:    _EventService_NewEvent_Handler,
		},
		{
			MethodName: "getEvent",
			Handler:    _EventService_GetEvent_Handler,
		},
		{
			MethodName: "updateEvent",
			Handler:    _EventService_UpdateEvent_Handler,
		},
		{
			MethodName: "deleteEvent",
			Handler:    _EventService_DeleteEvent_Handler,
		},
		{
			MethodName: "all",
			Handler:    _EventService_All_Handler,
		},
		{
			MethodName: "next",
			Handler:    _EventService_Next_Handler,
		},
		{
			MethodName: "day",
			Handler:    _EventService_Day_Handler,
		},
		{
			MethodName: "week",
			Handler:    _EventService_Week_Handler,
		},
		{
			MethodName: "month",
			Handler:    _EventService_Month_Handler,
		},
		{
			MethodName: "sinceDay",
			Handler:    _EventService_SinceDay_Handler,
		},
		{
			MethodName: "sinceWeek",
			Handler:    _EventService_SinceWeek_Handler,
		},
		{
			MethodName: "sinceMonth",
			Handler:    _EventService_SinceMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EventService.proto",
}