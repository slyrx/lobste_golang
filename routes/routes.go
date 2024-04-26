// routes/routes.go
package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slyrx/lobste_golang/models"
	"math/rand"
	"net/http"
	"time"
)

func RegisterRoutes(r *gin.Engine) {
	// 在这里注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		// 从缓存或数据库获取最热门的故事列表,并进行分页处理
		// 生成假数据
		stories := generateFakeStories(10)
		//showMore := true

		// 将假数据以 JSON 格式返回
		c.JSON(http.StatusOK, stories)
	})

	// 功能列表
}

func generateFakeStories(count int) []models.Story {
	stories := make([]models.Story, count)
	for i := 0; i < count; i++ {
		stories[i] = models.Story{
			ID:            int64(i + 1),
			Title:         fmt.Sprintf("Fake Story %d", i+1),
			URL:           fmt.Sprintf("https://example.com/story-%d", i+1),
			Description:   fmt.Sprintf("This is a fake story number %d.", i+1),
			Score:         rand.Intn(1000),
			CommentsCount: rand.Intn(100),
			Upvotes:       rand.Intn(500),
			Downvotes:     rand.Intn(100),
			Submitter: models.User{
				ID:       int64(rand.Intn(1000)),
				Username: fmt.Sprintf("user%d", rand.Intn(1000)),
			},
			SubmittedAt: time.Now().Add(-time.Duration(rand.Intn(86400*30)) * time.Second),
		}
	}
	return stories
}
