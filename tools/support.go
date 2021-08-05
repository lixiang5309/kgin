package tools

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Support struct {
	Key string
}

var redisPool redis.Conn

func SupportNew(param ...string) (*Support, error) {
	var err error
	if len(param) < 3 {
		return nil, errors.New("缺少参数")
	}
	redisPool, err = redisConn(param[0], param[1])
	if err != nil {
		return nil, errors.New("初始化redis失败")
	}
	return &Support{Key: param[2]}, err
}

func redisConn(url, password string) (redis.Conn, error) {
	var (
		conn    redis.Conn
		options []redis.DialOption
		err     error
	)
	options = append(options, redis.DialDatabase(0))
	options = append(options, redis.DialPassword(password))
	conn, err = redis.Dial("tcp", url, options...) //连接
	if err != nil {
		return conn, err
	}
	if _, err := conn.Do("AUTH", password); err != nil {
		return conn, err
	}
	return conn, err
}

//点赞/取消点赞 key 为这场点赞活动的唯一key  user为人
func (p *Support) SetSupport(user string) error {
	script := redis.NewScript(1, `local value = redis.call("HEXISTS",KEYS[1],ARGV[1]) if (value == 1) then
		redis.call("HDEL",KEYS[1],ARGV[1])
	else
		redis.call("HSET",KEYS[1],ARGV[1],'1')
	end`)
	_, err := script.Do(redisPool, p.Key, user)
	if err != nil {
		fmt.Println("this2:", err.Error())
	}
	return nil
}

//获取点赞总数
func (p *Support) GetSupportTotal() int64 {
	count, err := redis.Int64(redisPool.Do("HLEN", p.Key))
	if err != nil {
		return 0
	}
	return count
}

//获取该用户是否点赞
func (p *Support) GetUserSupport(user string) (bool, error) {
	exists, err := redis.Int(redisPool.Do("HEXISTS", p.Key, user))
	if err != nil {
		return false, errors.New("redis连接出错")
	}
	if exists != 1 {
		return false, nil
	}
	return true, nil
}

func (p *Support) Close() {
	redisPool.Close()
}
