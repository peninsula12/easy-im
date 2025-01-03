package test

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/logic"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"testing"
)

func TestFriendPutInListLogic_FriendPutInList(t *testing.T) {
	type args struct {
		in *social.FriendPutInListReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "name1",
			args: args{
				in: &social.FriendPutInListReq{
					UserId: "1842843962742673408",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewFriendPutInListLogic(context.Background(), svcCtx)
			got, err := l.FriendPutInList(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendPutInList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				l.Logger.Infof("FriendPutInList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
