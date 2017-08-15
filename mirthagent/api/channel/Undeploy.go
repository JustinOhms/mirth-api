package channel

import (
	"fmt"

	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
	"github.com/parnurzeal/gorequest"
)

func (Ω *Channel) Undeploy(args ...string) (chan bool, chan error) {
	c := make(chan bool, 1)
	ec := make(chan error, 1)
	req := Ω.Session.NewRequest().Post(Ω.Session.Paths.Channels.Undeploy())

	if len(args) > 0 {
		for _, channelId := range args {
			req.Send(fmt.Sprintf("channelId=%s", channelId))
		}
	}
	req.Send(fmt.Sprintf("returnErrors=true"))
	go setEnable(req, c, ec)
	return c, ec
}

func undeploy(req *gorequest.SuperAgent, c chan bool, ec chan error) {
	defer close(c)
	defer close(ec)
	r, _, e := req.EndBytes()
	if f.ResponseOrStatusErrors(ec, r, e, "Error undeploying channel(s)") {
		return
	}
	c <- true
}