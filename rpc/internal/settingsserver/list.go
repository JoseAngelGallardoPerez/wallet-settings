package settingsserver

import (
	"context"

	pb "github.com/Confialink/wallet-settings/rpc/proto/settings"
	"github.com/twitchtv/twirp"
)

// List returns list of settings by path
func (n *SettingsHandlerServer) List(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	settings, err := n.SettingsProvider.FindByPath(req.Path)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	responseSettings := make([]*pb.Setting, len(settings))
	for i, s := range settings {
		responseSettings[i] = &pb.Setting{
			Path:  string(s.Path),
			Value: s.Value,
		}
	}

	result := &pb.Response{
		Settings: responseSettings,
	}

	return result, nil
}
