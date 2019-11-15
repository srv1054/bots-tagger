package tagger

// manages all of the redis connectivity and validation for storage or

import (
	"os"

	"github.com/go-redis/redis"
)

func redisClient(myBot MyBot, db int) (rediscon *redis.Client, err error) {
	redislocale := os.Getenv("redislocale")
	client := redis.NewClient(&redis.Options{
		Addr:     myBot.RedisCon,
		Password: myBot.RedisPass,
		DB:       db,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func redisCheck(myBot MyBot, slackname string) (gitname string, worked bool) {
	client, err := redisClient(myBot.RedisDB)
	if err != nil {
		return "redis error", false
	}

	gn, err := client.Get(slackname).Result()
	if err != nil {
		//nil will mean not found
		return "", false
	}
	return gn, true
}

func redisSet(myBot MyBot, slackname string, gitname string) (status bool) {
	//set redis db1 for slackname->githubid
	client, err := redisClient(myBot.RedisDB)
	if err != nil {
		return false
	}
	err = client.Set(slackname, gitname, 0).Err()
	if err != nil {
		return false
	}

	//set redis db2 for githubid->slackname
	client, err = redisClient(2)
	if err != nil {
		return false
	}
	err = client.Set(gitname, slackname, 0).Err()
	if err != nil {
		return false
	}
	//for now while small initiating save to disc on each transaction.
	//this should be replaced by a timed process if this thing needs to scale at all
	_ = client.BgSave()

	return true
}
