// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: social.proto

package server

import (
	"context"

	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
)

type SocialServer struct {
	svcCtx *svc.ServiceContext
	social.UnimplementedSocialServer
}

func NewSocialServer(svcCtx *svc.ServiceContext) *SocialServer {
	return &SocialServer{
		svcCtx: svcCtx,
	}
}

func (s *SocialServer) FriendPutIn(ctx context.Context, in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	l := logic.NewFriendPutInLogic(ctx, s.svcCtx)
	return l.FriendPutIn(in)
}

func (s *SocialServer) FriendPutInHandle(ctx context.Context, in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	l := logic.NewFriendPutInHandleLogic(ctx, s.svcCtx)
	return l.FriendPutInHandle(in)
}

func (s *SocialServer) FriendPutInList(ctx context.Context, in *social.FriendPutInListReq) (*social.FriendPutInListResp, error) {
	l := logic.NewFriendPutInListLogic(ctx, s.svcCtx)
	return l.FriendPutInList(in)
}

func (s *SocialServer) FriendList(ctx context.Context, in *social.FriendListReq) (*social.FriendListResp, error) {
	l := logic.NewFriendListLogic(ctx, s.svcCtx)
	return l.FriendList(in)
}

// 群要求
func (s *SocialServer) GroupCreate(ctx context.Context, in *social.GroupCreateReq) (*social.GroupCreateResp, error) {
	l := logic.NewGroupCreateLogic(ctx, s.svcCtx)
	return l.GroupCreate(in)
}

func (s *SocialServer) GroupPutin(ctx context.Context, in *social.GroupPutinReq) (*social.GroupPutinResp, error) {
	l := logic.NewGroupPutinLogic(ctx, s.svcCtx)
	return l.GroupPutin(in)
}

func (s *SocialServer) GroupPutinList(ctx context.Context, in *social.GroupPutinListReq) (*social.GroupPutinListResp, error) {
	l := logic.NewGroupPutinListLogic(ctx, s.svcCtx)
	return l.GroupPutinList(in)
}

func (s *SocialServer) GroupPutInHandle(ctx context.Context, in *social.GroupPutInHandleReq) (*social.GroupPutInHandleResp, error) {
	l := logic.NewGroupPutInHandleLogic(ctx, s.svcCtx)
	return l.GroupPutInHandle(in)
}

func (s *SocialServer) GroupList(ctx context.Context, in *social.GroupListReq) (*social.GroupListResp, error) {
	l := logic.NewGroupListLogic(ctx, s.svcCtx)
	return l.GroupList(in)
}

func (s *SocialServer) GroupUsers(ctx context.Context, in *social.GroupUsersReq) (*social.GroupUsersResp, error) {
	l := logic.NewGroupUsersLogic(ctx, s.svcCtx)
	return l.GroupUsers(in)
}
