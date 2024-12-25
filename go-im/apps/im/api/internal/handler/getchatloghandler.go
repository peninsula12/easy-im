package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/types"
)

// 根据用户获取聊天记录
func getChatLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetChatLogLogic(r.Context(), svcCtx)
		resp, err := l.GetChatLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
