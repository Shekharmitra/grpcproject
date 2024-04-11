// persistent_test.go
package main

import (
    "testing"

    pb "grpcproject/grpcmodels/blogpost.proto"  
)

func TestMemoryPostRepository(t *testing.T) {
    repo := NewMemoryPostRepository()

    // Test CreatePost
    post := &pb.Post{
        PostId:          "1",
        Title:           "Test Post",
        Content:         "This is a test post.",
        Author:          "Test Author",
        PublicationDate: "2024-04-12",
        Tags:            []string{"test", "unit-test"},
    }
    createdPost, err := repo.CreatePost(post)
    if err != nil {
        t.Errorf("Failed to create post: %v", err)
    }
    if createdPost == nil {
        t.Errorf("Expected created post, got nil")
    }
    if createdPost.PostId != "1" {
        t.Errorf("Expected post ID to be 1, got %s", createdPost.PostId)
    }

    // Test ReadPost
    retrievedPost, err := repo.ReadPost("1")
    if err != nil {
        t.Errorf("Failed to read post: %v", err)
    }
    if retrievedPost == nil {
        t.Errorf("Expected retrieved post, got nil")
    }
    if retrievedPost.PostId != "1" {
        t.Errorf("Expected post ID to be 1, got %s", retrievedPost.PostId)
    }

    // Test UpdatePost
    updatedPost := &pb.Post{
        PostId:          "1",
        Title:           "Updated Test Post",
        Content:         "This is an updated test post.",
        Author:          "Updated Test Author",
        PublicationDate: "2024-04-13",
        Tags:            []string{"test", "unit-test", "update"},
    }
    updatedPost, err = repo.UpdatePost(updatedPost)
    if err != nil {
        t.Errorf("Failed to update post: %v", err)
    }
    if updatedPost == nil {
        t.Errorf("Expected updated post, got nil")
    }
    if updatedPost.Title != "Updated Test Post" {
        t.Errorf("Expected post title to be 'Updated Test Post', got %s", updatedPost.Title)
    }

    // Test DeletePost
    err = repo.DeletePost("1")
    if err != nil {
        t.Errorf("Failed to delete post: %v", err)
    }
    _, err = repo.ReadPost("1")
    if err == nil {
        t.Errorf("Expected post not found after deletion")
    }
}
