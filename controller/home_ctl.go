package controller

import (
	"fmt"
	"github.com/slyrx/lobste_golang/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (h *HomeController) Active(c *gin.Context) {
	stories, showMore, err := h.getActiveStories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(200, gin.H{
	//	"message":  stories,
	//	"message2": showMore,
	//})

	// 渲染视图
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Active Discussions",
		"stories":  stories,
		"showMore": showMore,
	})
}

func (h *HomeController) getActiveStories(c *gin.Context) ([]models.Story, bool, error) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 20
	}

	stories := h.GetActiveStories(page)

	showMore := len(stories) == models.ActiveStoriesPerPage

	return stories, showMore, nil
}

func (h *HomeController) getActivePartial() string {
	return `<div class="active-discussions">
        <h2>Active Discussions</h2>
        <p>These are the most active discussions on the site.</p>
    </div>`
}

func (h *HomeController) GetActiveStories(count int) []models.Story {
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
