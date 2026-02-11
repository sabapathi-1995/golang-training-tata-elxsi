package api

import (
	"context"

	"google.golang.org/grpc"
)

// Service/Method names (match proto reference).
const (
	NotesServiceName          = "notes.v1.NotesService"
	methodCreateNote          = "/notes.v1.NotesService/CreateNote"
	methodListNotes           = "/notes.v1.NotesService/ListNotes"
	methodUploadNotes         = "/notes.v1.NotesService/UploadNotes"
	methodChatNotes           = "/notes.v1.NotesService/ChatNotes"
)

// NotesServiceServer is the server API.
type NotesServiceServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	ListNotes(*ListNotesRequest, NotesService_ListNotesServer) error
	UploadNotes(NotesService_UploadNotesServer) error
	ChatNotes(NotesService_ChatNotesServer) error
}

// UnimplementedNotesServiceServer can be embedded for forward compatibility.
type UnimplementedNotesServiceServer struct{}

func (UnimplementedNotesServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, grpc.Errorf(grpc.Code(grpc.ErrServerStopped), "method CreateNote not implemented")
}
func (UnimplementedNotesServiceServer) ListNotes(*ListNotesRequest, NotesService_ListNotesServer) error {
	return grpc.Errorf(grpc.Code(grpc.ErrServerStopped), "method ListNotes not implemented")
}
func (UnimplementedNotesServiceServer) UploadNotes(NotesService_UploadNotesServer) error {
	return grpc.Errorf(grpc.Code(grpc.ErrServerStopped), "method UploadNotes not implemented")
}
func (UnimplementedNotesServiceServer) ChatNotes(NotesService_ChatNotesServer) error {
	return grpc.Errorf(grpc.Code(grpc.ErrServerStopped), "method ChatNotes not implemented")
}

// ---- Server stream interfaces ----

type NotesService_ListNotesServer interface {
	Send(*Note) error
	grpc.ServerStream
}

type NotesService_UploadNotesServer interface {
	Recv() (*CreateNoteRequest, error)
	SendAndClose(*UploadNotesResponse) error
	grpc.ServerStream
}

type NotesService_ChatNotesServer interface {
	Recv() (*CreateNoteRequest, error)
	Send(*Ack) error
	grpc.ServerStream
}

// RegisterNotesServiceServer registers service.
func RegisterNotesServiceServer(s grpc.ServiceRegistrar, srv NotesServiceServer) {
	s.RegisterService(&grpc.ServiceDesc{
		ServiceName: NotesServiceName,
		HandlerType: (*NotesServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CreateNote",
				Handler:    _NotesService_CreateNote_Handler,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "ListNotes",
				Handler:       _NotesService_ListNotes_Handler,
				ServerStreams: true,
			},
			{
				StreamName:    "UploadNotes",
				Handler:       _NotesService_UploadNotes_Handler,
				ClientStreams: true,
			},
			{
				StreamName:    "ChatNotes",
				Handler:       _NotesService_ChatNotes_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "api/notes.proto",
	}, srv)
}

func _NotesService_CreateNote_Handler(srv any, ctx context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NotesServiceServer).CreateNote(ctx, in)
}

func _NotesService_ListNotes_Handler(srv any, stream grpc.ServerStream) error {
	in := new(ListNotesRequest)
	if err := stream.RecvMsg(in); err != nil {
		return err
	}
	return srv.(NotesServiceServer).ListNotes(in, &notesServiceListNotesServer{stream})
}

type notesServiceListNotesServer struct{ grpc.ServerStream }
func (x *notesServiceListNotesServer) Send(m *Note) error { return x.ServerStream.SendMsg(m) }

func _NotesService_UploadNotes_Handler(srv any, stream grpc.ServerStream) error {
	return srv.(NotesServiceServer).UploadNotes(&notesServiceUploadNotesServer{stream})
}

type notesServiceUploadNotesServer struct{ grpc.ServerStream }
func (x *notesServiceUploadNotesServer) Recv() (*CreateNoteRequest, error) {
	m := new(CreateNoteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func (x *notesServiceUploadNotesServer) SendAndClose(m *UploadNotesResponse) error { return x.ServerStream.SendMsg(m) }

func _NotesService_ChatNotes_Handler(srv any, stream grpc.ServerStream) error {
	return srv.(NotesServiceServer).ChatNotes(&notesServiceChatNotesServer{stream})
}

type notesServiceChatNotesServer struct{ grpc.ServerStream }
func (x *notesServiceChatNotesServer) Recv() (*CreateNoteRequest, error) {
	m := new(CreateNoteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func (x *notesServiceChatNotesServer) Send(m *Ack) error { return x.ServerStream.SendMsg(m) }

// ---- Client stub ----

type NotesServiceClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	ListNotes(ctx context.Context, in *ListNotesRequest, opts ...grpc.CallOption) (NotesService_ListNotesClient, error)
	UploadNotes(ctx context.Context, opts ...grpc.CallOption) (NotesService_UploadNotesClient, error)
	ChatNotes(ctx context.Context, opts ...grpc.CallOption) (NotesService_ChatNotesClient, error)
}

type notesServiceClient struct{ cc grpc.ClientConnInterface }

func NewNotesServiceClient(cc grpc.ClientConnInterface) NotesServiceClient { return &notesServiceClient{cc} }

func (c *notesServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	if err := c.cc.Invoke(ctx, methodCreateNote, in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

type NotesService_ListNotesClient interface {
	Recv() (*Note, error)
	grpc.ClientStream
}

func (c *notesServiceClient) ListNotes(ctx context.Context, in *ListNotesRequest, opts ...grpc.CallOption) (NotesService_ListNotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotesService_serviceDesc.Streams[0], methodListNotes, opts...)
	if err != nil {
		return nil, err
	}
	x := &notesServiceListNotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type notesServiceListNotesClient struct{ grpc.ClientStream }
func (x *notesServiceListNotesClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type NotesService_UploadNotesClient interface {
	Send(*CreateNoteRequest) error
	CloseAndRecv() (*UploadNotesResponse, error)
	grpc.ClientStream
}

func (c *notesServiceClient) UploadNotes(ctx context.Context, opts ...grpc.CallOption) (NotesService_UploadNotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotesService_serviceDesc.Streams[1], methodUploadNotes, opts...)
	if err != nil {
		return nil, err
	}
	return &notesServiceUploadNotesClient{stream}, nil
}

type notesServiceUploadNotesClient struct{ grpc.ClientStream }
func (x *notesServiceUploadNotesClient) Send(m *CreateNoteRequest) error { return x.ClientStream.SendMsg(m) }
func (x *notesServiceUploadNotesClient) CloseAndRecv() (*UploadNotesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	out := new(UploadNotesResponse)
	if err := x.ClientStream.RecvMsg(out); err != nil {
		return nil, err
	}
	return out, nil
}

type NotesService_ChatNotesClient interface {
	Send(*CreateNoteRequest) error
	Recv() (*Ack, error)
	CloseSend() error
	grpc.ClientStream
}

func (c *notesServiceClient) ChatNotes(ctx context.Context, opts ...grpc.CallOption) (NotesService_ChatNotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotesService_serviceDesc.Streams[2], methodChatNotes, opts...)
	if err != nil {
		return nil, err
	}
	return &notesServiceChatNotesClient{stream}, nil
}

type notesServiceChatNotesClient struct{ grpc.ClientStream }
func (x *notesServiceChatNotesClient) Send(m *CreateNoteRequest) error { return x.ClientStream.SendMsg(m) }
func (x *notesServiceChatNotesClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceDesc reference for client streams.
var _NotesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: NotesServiceName,
	HandlerType: (*NotesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateNote", Handler: _NotesService_CreateNote_Handler},
	},
	Streams: []grpc.StreamDesc{
		{StreamName: "ListNotes", Handler: _NotesService_ListNotes_Handler, ServerStreams: true},
		{StreamName: "UploadNotes", Handler: _NotesService_UploadNotes_Handler, ClientStreams: true},
		{StreamName: "ChatNotes", Handler: _NotesService_ChatNotes_Handler, ServerStreams: true, ClientStreams: true},
	},
	Metadata: "api/notes.proto",
}
