// client.go
package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    pb "grpcproject/grpcmodels/blogpost.proto" 
)

func main() {
    conn, err := grpc.Dial(":50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewPostServiceClient(conn)

    // Create Post
    createPostReq := &pb.CreatePostRequest{
        Title:           "New Post",
        Content:         "This is a new post.",
        Author:          "John Doe",
        PublicationDate: "2024-04-12",
        Tags:            []string{"gRPC", "Go", "Blogging"},
    }
    createPostRes, err := c.CreatePost(context.Background(), createPostReq)
    if err != nil {
        log.Fatalf("Failed to create post: %v", err)
    }
    log.Printf("Created Post: %v", createPostRes.GetPost())

    // Read Post
    readPostReq := &pb.ReadPostRequest{
        PostId: createPostRes.GetPost().GetPostId(),
    }
    readPostRes, err := c.ReadPost(context.Background(), readPostReq)
    if err != nil {
        log.Fatalf("Failed to read post: %v", err)
    }
    log.Printf("Read Post: %v", readPostRes.GetPost())

    // Update Post
    updatePostReq := &pb.UpdatePostRequest{
        PostId: createPostRes.GetPost().GetPostId(),
        Title:  "Updated Post",
        Content: "This is an updated post.",
        Author:  "Jane Doe",
        Tags:    []string{"gRPC", "Go", "Blogging", "Update"},
    }
    updatePostRes, err := c.UpdatePost(context.Background(), updatePostReq)
    if err != nil {
        log.Fatalf("Failed to update post: %v", err)
    }
    log.Printf("Updated Post: %v", updatePostRes.GetPost())

    // Delete Post
    deletePostReq := &pb.DeletePostRequest{
        PostId: createPostRes.GetPost().GetPostId(),
    }
    deletePostRes, err := c.DeletePost(context.Background(), deletePostReq)
    if err != nil {
        log.Fatalf("Failed to delete post: %v", err)
    }
    log.Printf("Delete Post: Success=%t", deletePostRes.GetSuccess())
}
