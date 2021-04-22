package settingsserver

import (
	"context"

	pb "github.com/Confialink/wallet-settings/rpc/proto/settings"
	"github.com/twitchtv/twirp"
)

// Get returns value by path
func (n *SettingsHandlerServer) Get(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	setting, err := n.SettingsProvider.FindOneByPath(req.Path)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	result := &pb.Response{
		Setting: &pb.Setting{
			Path:  string(setting.Path),
			Value: setting.Value,
		},
	}

	return result, nil
}
