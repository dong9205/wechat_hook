package wx

import (
	"errors"
	"sync"
	"time"

	"github.com/dong9205/wechat_hook/pkg/logger"
	"github.com/eatmoreapple/openwechat"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
)

type BotMgr struct {
	bot        *openwechat.Bot
	frendCache *cache.Cache
	groupCache *cache.Cache
	self       *openwechat.Self
	logger     zerolog.Logger
	mu         sync.Mutex
}

var botmgr *BotMgr

func GetBotMgr() *BotMgr {
	if botmgr == nil {
		botmgr = newBotManager()
	}
	return botmgr
}

func newBotManager() *BotMgr {
	logger := logger.GetLogger()
	bot := openwechat.DefaultBot(openwechat.Desktop)
	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		logger.Fatal().Err(err).Msg("登陆失败")
		return nil
	}
	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		logger.Fatal().Err(err).Msg("获取登录用户失败")
		return nil
	}
	return &BotMgr{
		bot:        bot,
		self:       self,
		frendCache: cache.New(1*time.Hour, 1*time.Minute),
		groupCache: cache.New(1*time.Hour, 1*time.Minute),
		logger:     logger,
	}
}

func (botmgr *BotMgr) GetFirend(friendName string) (friend *openwechat.Friend, err error) {
	friendAnd, ok := botmgr.frendCache.Get(friendName)
	if ok {
		if friend, ok = friendAnd.(*openwechat.Friend); ok {
			return
		}
	}
	friends, err := botmgr.self.Friends()
	if err != nil {
		return nil, err
	}
	for _, tmpFriend := range friends {
		if tmpFriend.NickName == friendName {
			friend = tmpFriend
		}
		botmgr.frendCache.Set(tmpFriend.NickName, tmpFriend, cache.DefaultExpiration)
	}
	if friend == nil {
		return nil, errors.New("friend not found")
	}
	return friend, err
}

func (botmgr *BotMgr) GetGroup(groupName string) (group *openwechat.Group, err error) {
	groupAny, ok := botmgr.groupCache.Get(groupName)
	if ok {
		if group, ok = groupAny.(*openwechat.Group); ok {
			return
		}
	}
	groups, err := botmgr.self.Groups()
	if err != nil {
		return nil, err
	}
	for _, tmpgroup := range groups {
		if tmpgroup.NickName == groupName {
			group = tmpgroup
		}
		botmgr.groupCache.Set(tmpgroup.NickName, tmpgroup, cache.DefaultExpiration)
	}
	if group == nil {
		return nil, errors.New("group not found")
	}
	return group, err
}

func (botmgr *BotMgr) Send(msg, dest, destType string) error {
	switch destType {
	case "friend":
		friend, err := botmgr.GetFirend(dest)
		if err != nil {
			botmgr.logger.Error().Err(err).Msg("获取好友失败")
			return err
		}
		botmgr.mu.Lock()
		defer botmgr.mu.Unlock()
		msgRes, err := friend.SendText(msg)
		if err != nil {
			botmgr.logger.Error().Err(err).Msg("发送消息失败")
			return err
		}
		botmgr.logger.Info().Any("msgRes", msgRes).Msg("发送消息成功")
	case "group":
		group, err := botmgr.GetGroup(dest)
		if err != nil {
			botmgr.logger.Error().Err(err).Msg("获取群组失败")
			return err
		}
		botmgr.mu.Lock()
		defer botmgr.mu.Unlock()
		msgRes, err := group.SendText(msg)
		if err != nil {
			botmgr.logger.Error().Err(err).Msg("发送消息失败")
			return err
		}
		botmgr.logger.Info().Any("msgRes", msgRes).Msg("发送消息成功")
	default:
		return errors.New("dest type error")
	}
	return nil
}
