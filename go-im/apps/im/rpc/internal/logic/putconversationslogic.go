package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/peninsula12/easy-im/go-im/apps/im/model"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/im"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/pkg/status"
	"github.com/peninsula12/easy-im/go-im/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PutConversations 更新会话
func (l *PutConversationsLogic) PutConversations(in *im.PutConversationsReq) (*im.PutConversationsResp, error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "conversationModel.FindByUserId err %v,req %v", err, in.UserId)
	}
	if data.ConversationList == nil {
		data.ConversationList = make(map[string]*immodels.Conversation)
	}
	for s, conversation := range in.ConversationList {
		var oldTotal int
		if data.ConversationList[s] != nil {
			oldTotal = data.ConversationList[s].Total
		}
		data.ConversationList[s] = &immodels.Conversation{
			ConversationId: conversation.ConversationId,
			ChatType:       status.ChatType(conversation.ChatType),
			IsShow:         conversation.IsShow,
			Total:          int(conversation.Read) + oldTotal,
			Seq:            conversation.Seq,
		}
	}
	_, err = l.svcCtx.ConversationsModel.Update(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "conversationModel.Update err %v,req %v", err, data)
	}
	return &im.PutConversationsResp{}, nil
}
