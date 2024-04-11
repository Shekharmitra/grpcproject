// server.go
// controller.go
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    PostService "grpcproject/service/service" 

    pb "grpcproject/grpcmodels/blogpost.proto" 
)

type PostController struct {
    pb.UnimplementedPostServiceServer
    service PostService
}

func NewPostController(service PostService) *PostController {
    return &PostController{
        service: service,
    }
}

func (c *PostController) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
    return c.service.CreatePost(ctx, req)
}

func (c *PostController) ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.ReadPostResponse, error) {
    return c.service.ReadPost(ctx, req)
}

func (c *PostController) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
    return c.service.UpdatePost(ctx, req)
}

func (c *PostController) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
    return c.service.DeletePost(ctx, req)
}

func main() {
    repo := NewMemoryPostRepository()
    service := NewPostService(repo)
    controller := NewPostController(service)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterPostServiceServer(grpcServer, controller)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
