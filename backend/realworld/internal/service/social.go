package service

import (
	"context"
	pb "realworld/api/realworld/v1"
)

func (s *RealWorldService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.MultipleArticlesReplay, error) {
	return &pb.MultipleArticlesReplay{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticlesReplay, error) {
	return &pb.MultipleArticlesReplay{}, nil
}
func (s *RealWorldService) GetArticles(ctx context.Context, req *pb.GetArticlesRequest) (*pb.MultipleArticlesReplay, error) {
	return &pb.MultipleArticlesReplay{}, nil
}
func (s *RealWorldService) CreateArticles(ctx context.Context, req *pb.CreateArticlesRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) UpdateArticles(ctx context.Context, req *pb.UpdateArticlesRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) DeleteArticles(ctx context.Context, req *pb.DeleteArticlesRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}

func (s *RealWorldService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.SingleCommentReplay, error) {
	return &pb.SingleCommentReplay{}, nil
}
func (s *RealWorldService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.MultipleCommentReplay, error) {
	return &pb.MultipleCommentReplay{}, nil
}
func (s *RealWorldService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.SingleCommentReplay, error) {
	return &pb.SingleCommentReplay{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.TagsReplay, error) {
	return &pb.TagsReplay{}, nil
}
