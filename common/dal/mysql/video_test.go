package mysql

import (
	"testing"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	InitDB()
}

func TestFindVideoListinVideo(t *testing.T) {
	InitDB()
	list, err := FindVideoListinVideo(111)
	assert.NoError(t, err)
	logs.Info("list: %s", list)
}

func TestDatabaseVideo_user(t *testing.T) {
	InitDB()
	user, err := FindUserInfoinUser(111)
	assert.NoError(t, err)
	logs.Info("user: %s", user)
}

func TestFindIDinLike(t *testing.T) {
	InitDB()
	like_count, err := FindCountinLikes(123, 111)
	assert.NoError(t, err)
	logs.Info("like_count: %s", like_count)
}

func TestCreateVideo(t *testing.T) {
	InitDB()
	err := CreateVideo(2, 111, "11", "11", 1, 2, "11", "11")
	assert.NoError(t, err)

}

func TestFindMaxIdinVideos(t *testing.T) {
	InitDB()
	video_id, err := FindMaxIdinVideos()
	assert.NoError(t, err)
	logs.Info("video_id: %s", video_id)

}

func TestFindCountinFollows(t *testing.T) {
	InitDB()
	follow_count, err := FindCountinFollows(111, 111)
	assert.NoError(t, err)
	logs.Info("follow_count: %s", follow_count)
}
