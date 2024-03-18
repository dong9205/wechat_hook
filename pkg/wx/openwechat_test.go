package wx

import (
	"testing"
	"time"

	"github.com/dong9205/wechat_hook/pkg/logger"
)

func setup() {
	logger.InitLogger()
	newBotManager()
}

func TestOpenWechatLoginGetUsers(t *testing.T) {
	setup()
	// 获取所有的好友
	friends, err := botmgr.self.Friends()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("friends: %v\n", friends)
}

func TestOpenWechatLoginGetGroups(t *testing.T) {
	setup()
	// 获取所有的群组
	for {
		groups, err := botmgr.self.Groups()
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log("groups", groups)
		time.Sleep(time.Second * 5)
	}
}
