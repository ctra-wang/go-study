package services

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

// 用于解决缓存gpt的secret-key使用

type CacheHistory struct {
	db *cache.Cache
}

func (c *CacheHistory) get(key string) (interface{}, bool) {
	res, flag := c.db.Get(key)
	if flag {
		return res, true
	}
	return nil, false
}

func (c *CacheHistory) set(key string, value interface{}) {
	// 默认5分钟
	c.db.Set(key, value, time.Minute*5)
}

func (c *CacheHistory) clear() {
	// 全部清空（单节点用户）
	c.db.Flush()
}

// CacheHistoryInterface 接口定义
type CacheHistoryInterface interface {
	// chatGPT秘钥相关
	get(key string) (interface{}, bool)
	set(key string, value interface{})
	clear()

	// SetQACache 缓存 chatGpt聊天记录 ------ chatGPT 对话相关
	SetQACache(q string, a string)
	// GetQACache 缓存 chatGpt聊天记录
	GetQACache() (string, bool)
	ClearQACache()
}

const (
	userSessionContextKey = "userSessionContext"
)

// NewCacheHistory 实例化一个 CacheHistory对象，加上过期时间和销毁时间
func NewCacheHistory() CacheHistoryInterface {
	return &CacheHistory{
		db: cache.New(time.Second*60, time.Minute*10),
	}
}

func (c *CacheHistory) SetQACache(q string, a string) {
	value := fmt.Sprintf("%s /n %s", q, a)
	c.set(userSessionContextKey, value)
}

func (c *CacheHistory) GetQACache() (string, bool) {

	value, b := c.get(userSessionContextKey)
	if b {
		return value.(string), true
	}
	return "", false
}

func (c *CacheHistory) ClearQACache() {
	c.clear()
}
