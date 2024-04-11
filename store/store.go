// store.go
package store

import (
	pb "grpcproject/grpcmodels/blogpost.proto"  
	
)

type PostRepository interface {
    CreatePost(post *pb.Post) (*pb.Post, error)
    ReadPost(postID string) (*pb.Post, error)
    UpdatePost(post *pb.Post) (*pb.Post, error)
    DeletePost(postID string) error
}

type MemoryPostRepository struct {
    posts map[string]*pb.Post
}

func NewMemoryPostRepository() *MemoryPostRepository {
    return &MemoryPostRepository{
        posts: make(map[string]*pb.Post),
    }
}

func (r *MemoryPostRepository) CreatePost(post *pb.Post) (*pb.Post, error) {
    r.posts[post.PostId] = post
    return post, nil
}

func (r *MemoryPostRepository) ReadPost(postID string) (*pb.Post, error) {
    post, ok := r.posts[postID]
    if !ok {
        return nil, fmt.Errorf("post not found")
    }
    return post, nil
}

func (r *MemoryPostRepository) UpdatePost(post *pb.Post) (*pb.Post, error) {
    r.posts[post.PostId] = post
    return post, nil
}

func (r *MemoryPostRepository) DeletePost(postID string) error {
    delete(r.posts, postID)
    return nil
}
