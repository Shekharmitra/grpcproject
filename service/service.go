// service.go
package main

import (
    "context"
    "fmt"

	pb "grpcproject/grpcmodels/blogpost.proto" 
	PostRepository "grpcproject/store/store" 
)

type PostService interface {
    CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error)
    ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.ReadPostResponse, error)
    UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error)
    DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error)
}

type postService struct {
    repo PostRepository
}

func NewPostService(repo PostRepository) PostService {
    return &postService{repo: repo}
}

func (s *postService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
    post, err := s.repo.CreatePost(req.GetPost())
    if err != nil {
        return nil, err
    }
    return &pb.CreatePostResponse{Post: post}, nil
}

func (s *postService) ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.ReadPostResponse, error) {
    post, err := s.repo.ReadPost(req.GetPostId())
    if err != nil {
        return &pb.ReadPostResponse{Error: fmt.Sprintf("Post not found: %v", err)}, nil
    }
    return &pb.ReadPostResponse{Post: post}, nil
}

func (s *postService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
    post, err := s.repo.UpdatePost(req.GetPost())
    if err != nil {
        return &pb.UpdatePostResponse{Error: fmt.Sprintf("Update failed: %v", err)}, nil
    }
    return &pb.UpdatePostResponse{Post: post}, nil
}

func (s *postService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
    err := s.repo.DeletePost(req.GetPostId())
    if err != nil {
        return &pb.DeletePostResponse{Success: false, ErrorMessage: fmt.Sprintf("Delete failed: %v", err)}, nil
    }
    return &pb.DeletePostResponse{Success: true}, nil
}
