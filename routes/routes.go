package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slyrx/lobste_golang/controller"
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
	homeController := &controller.HomeController{}
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/active", homeController.Active)

	//// 首页相关路由
	//homeController := new(controllers.HomeController)
	//r.GET("/", homeController.Index)
	//r.GET("/404", homeController.FourOhFour)
	//r.GET("/rss", homeController.IndexRSS)
	//r.GET("/hottest", homeController.IndexJSON)
	//r.GET("/page/:page", homeController.IndexPage)
	//r.GET("/active", homeController.Active)
	//r.GET("/active/page/:page", homeController.ActivePage)
	//r.GET("/newest", homeController.Newest)
	//r.GET("/newest/page/:page", homeController.NewestPage)
	//r.GET("/recent", homeController.Recent)
	//r.GET("/recent/page/:page", homeController.RecentPage)
	//r.GET("/hidden", homeController.Hidden)
	//r.GET("/hidden/page/:page", homeController.HiddenPage)
	//r.GET("/saved", homeController.Saved)
	//r.GET("/saved/page/:page", homeController.SavedPage)
	//r.GET("/upvoted/stories", homeController.UpvotedStories)
	//r.GET("/upvoted/stories/page/:page", homeController.UpvotedStoriesPage)
	//r.GET("/upvoted", homeController.RedirectUpvotedStories)
	//r.GET("/upvoted/page/:page", homeController.RedirectUpvotedStoriesPage)
	//r.GET("/top", homeController.Top)
	//r.GET("/top/rss", homeController.TopRSS)
	//r.GET("/top/page/:page", homeController.TopPage)
	//r.GET("/top/:length", homeController.TopLength)
	//r.GET("/top/:length/page/:page", homeController.TopLengthPage)
	//
	//// 用户相关路由
	//userController := new(controllers.UserController)
	//r.GET("/threads", userController.UserThreads)
	//r.GET("/replies", userController.AllReplies)
	//r.GET("/replies/page/:page", userController.AllRepliesPage)
	//r.GET("/replies/comments", userController.CommentReplies)
	//r.GET("/replies/comments/page/:page", userController.CommentRepliesPage)
	//r.GET("/replies/stories", userController.StoryReplies)
	//r.GET("/replies/stories/page/:page", userController.StoryRepliesPage)
	//r.GET("/replies/unread", userController.UnreadReplies)
	//r.GET("/replies/unread/page/:page", userController.UnreadRepliesPage)
	//r.GET("/login", userController.Login)
	//r.POST("/login", userController.LoginPost)
	//r.POST("/logout", userController.Logout)
	//r.GET("/login/2fa", userController.TwoFALogin)
	//r.POST("/login/2fa_verify", userController.TwoFAVerify)
	//r.GET("/signup", userController.Signup)
	//r.POST("/signup", userController.SignupPost)
	//r.GET("/signup/invite", userController.SignupInvite)
	//r.GET("/login/forgot_password", userController.ForgotPassword)
	//r.POST("/login/reset_password", userController.ResetPassword)
	//r.GET("/login/set_new_password", userController.SetNewPassword)
	//r.POST("/login/set_new_password", userController.SetNewPasswordPost)
	//r.GET("/users", userController.UserTree)
	//r.GET("/~:username", userController.ShowUser)
	//r.GET("/~:username/standing", userController.UserStanding)
	//r.GET("/~:user/stories", userController.NewestByUser)
	//r.GET("/~:user/stories/page/:page", userController.NewestByUserPage)
	//r.GET("/~:user/threads", userController.UserThreads)
	//r.POST("/~:username/ban", userController.BanUser)
	//r.POST("/~:username/unban", userController.UnbanUser)
	//r.POST("/~:username/disable_invitation", userController.DisableInvitation)
	//r.POST("/~:username/enable_invitation", userController.EnableInvitation)

	//// 标签相关路由
	//tagController := new(controllers.TagController)
	//r.GET("/t/:tag", tagController.SingleTag)
	//r.GET("/t/:tag", tagController.MultiTag)
	//r.GET("/t/:tag/page/:page", tagController.TaggedPage)
	//
	//// 域名相关路由
	//domainController := new(controllers.DomainController)
	//r.GET("/domain/:id", domainController.RedirectDomain)
	//r.GET("/domain/:id/page/:page", domainController.RedirectDomainPage)
	//r.GET("/domains/:id", domainController.ForDomain)
	//r.GET("/domains/:id/page/:page", domainController.ForDomainPage)
	//r.POST("/domains", domainController.Create)
	//r.PATCH("/domains/:id", domainController.Update)
	//
	//// 搜索相关路由
	//searchController := new(controllers.SearchController)
	//r.GET("/search", searchController.Index)
	//r.GET("/search/:q", searchController.Search)
	//
	//// 故事相关路由
	//storyController := new(controllers.StoryController)
	//r.GET("/stories/:short_id", storyController.RedirectShortID)
	//r.POST("/stories/:id/upvote", storyController.Upvote)
	//r.POST("/stories/:id/flag", storyController.Flag)
	//r.POST("/stories/:id/unvote", storyController.Unvote)
	//r.PATCH("/stories/:id/destroy", storyController.Destroy)
	//r.PATCH("/stories/:id/undelete", storyController.Undelete)
	//r.POST("/stories/:id/hide", storyController.Hide)
	//r.POST("/stories/:id/unhide", storyController.Unhide)
	//r.POST("/stories/:id/save", storyController.Save)
	//r.POST("/stories/:id/unsave", storyController.Unsave)
	//r.GET("/stories/:id/suggest", storyController.Suggest)
	//r.POST("/stories/:id/suggest", storyController.SubmitSuggestions)
	//r.POST("/stories/fetch_url_attributes", storyController.FetchURLAttributes)
	//r.POST("/stories/preview", storyController.Preview)
	//r.POST("/stories/check_url_dupe", storyController.CheckURLDupe)
	//
	//// 评论相关路由
	//commentController := new(controllers.CommentController)
	//r.GET("/comments/:id", commentController.RedirectFromShortID)
	//r.GET("/comments/:id/reply", commentController.Reply)
	//r.POST("/comments/:id/upvote", commentController.Upvote)
	//r.POST("/comments/:id/flag", commentController.Flag)
	//r.POST("/comments/:id/unvote", commentController.Unvote)
	//r.POST("/comments/:id/delete", commentController.Delete)
	//r.POST("/comments/:id/undelete", commentController.Undelete)
	//r.POST("/comments/:id/disown", commentController.Disown)
	//r.GET("/comments/page/:page", commentController.Index)
	//r.GET("/comments", commentController.Index)
	//
	//// 消息相关路由
	//messageController := new(controllers.MessageController)
	//r.GET("/messages/sent", messageController.Sent)
	//r.GET("/messages", messageController.Index)
	//r.POST("/messages/batch_delete", messageController.BatchDelete)
	//r.POST("/messages/:id/keep_as_new", messageController.KeepAsNew)
	//r.POST("/messages/:id/mod_note", messageController.ModNote)
	//
	//// 其他路由
	//r.GET("/inbox", controllers.InboxHandler)
	//r.GET("/c/:id", commentController.RedirectFromShortID)
	//r.GET("/c/:id.json", commentController.ShowShortID)
	//r.GET("/s/:story_id/:title/comments/:id", commentController.RedirectFromShortID)
	//r.GET("/s/:id/:title", storyController.Show)

	//// 重定向路由
	//r.GET("/u", controllers.RedirectToUsers)
	//r.GET("/u/:username", controllers.RedirectToUserProfile)
	//r.GET("/@:username", controllers.RedirectToUserProfile)
	//r.GET("/u/:username/standing", controllers.RedirectToUserStanding)
	//r.GET("/newest/:user", controllers.RedirectToNewestByUser)
	//r.GET("/newest/:user/page/:page", controllers.RedirectToNewestByUserPage)
	//r.GET("/threads/:user", controllers.RedirectToUserThreads)
	//
	//// 头像相关路由
	//r.GET("/avatars/:username_size.png", controllers.ShowAvatar)
	//r.POST("/avatars/expire", controllers.ExpireAvatar)
	//
	//// 设置相关路由
	//r.GET("/settings", controllers.Settings)
	//r.POST("/settings", controllers.UpdateSettings)
	//r.POST("/settings/delete_account", controllers.DeleteAccount)
	//r.GET("/settings/2fa", controllers.TwoFASettings)
	//r.POST("/settings/2fa_auth", controllers.TwoFAAuth)
	//r.GET("/settings/2fa_enroll", controllers.TwoFAEnroll)
	//r.GET("/settings/2fa_verify", controllers.TwoFAVerify)
	//r.POST("/settings/2fa_update", controllers.TwoFAUpdate)
	//r.POST("/settings/pushover_auth", controllers.PushoverAuth)
	//r.GET("/settings/pushover_callback", controllers.PushoverCallback)
	//r.GET("/settings/mastodon_authentication", controllers.MastodonAuthentication)
	//r.GET("/settings/mastodon_auth", controllers.MastodonAuth)
	//r.GET("/settings/mastodon_callback", controllers.MastodonCallback)
	//r.POST("/settings/mastodon_disconnect", controllers.MastodonDisconnect)
	//r.GET("/settings/github_auth", controllers.GithubAuth)
	//r.GET("/settings/github_callback", controllers.GithubCallback)
	//r.POST("/settings/github_disconnect", controllers.GithubDisconnect)
	//
	//// Keybase 相关路由
	//r.GET("/.well-known/keybase-proof-config", controllers.KeybaseConfig)
	//r.POST("/keybase_proofs", controllers.CreateKeybaseProof)
	//r.DELETE("/keybase_proofs/:id", controllers.DestroyKeybaseProof)
	//
	//// 过滤器相关路由
	//r.GET("/filters", controllers.Filters)
	//r.POST("/filters", controllers.UpdateFilters)
	//
	//// 标签相关路由
	//r.GET("/tags", controllers.Tags)
	//r.GET("/tags.json", controllers.TagsJSON)
	//r.GET("/tags/new", controllers.NewTag)
	//r.GET("/tags/:tag_name/edit", controllers.EditTag)
	//r.POST("/tags", controllers.CreateTag)
	//r.POST("/tags/:tag_name", controllers.UpdateTag)
	//
	//// 分类相关路由
	//r.GET("/categories/new", controllers.NewCategory)
	//r.GET("/categories/:category_name/edit", controllers.EditCategory)
	//r.GET("/categories/:category", controllers.Category)
	//r.POST("/categories", controllers.CreateCategory)
	//r.POST("/categories/:category_name", controllers.UpdateCategory)
	//
	//// 邀请相关路由
	//r.POST("/invitations", controllers.CreateInvitation)
	//r.GET("/invitations", controllers.Invitations)
	//r.GET("/invitations/request", controllers.RequestInvitation)
	//r.POST("/invitations/create_by_request", controllers.CreateInvitationByRequest)
	//r.GET("/invitations/confirm/:code", controllers.ConfirmInvitationEmail)
	//r.POST("/invitations/send_for_request", controllers.SendInvitationForRequest)
	//r.GET("/invitations/:invitation_code", controllers.InvitedSignup)
	//r.POST("/invitations/delete_request", controllers.DeleteInvitationRequest)
	//
	//// 帽子相关路由
	//r.GET("/hats", controllers.Hats)
	//r.GET("/hats/build_request", controllers.BuildHatRequest)
	//r.POST("/hats/create_request", controllers.CreateHatRequest)
	//r.GET("/hats/requests", controllers.HatRequests)
	//r.POST("/hats/approve_request/:id", controllers.ApproveHatRequest)
	//r.POST("/hats/reject_request/:id", controllers.RejectHatRequest)
	//
	//// 版主相关路由
	//r.GET("/moderations", controllers.Moderations)
	//r.GET("/moderations/page/:page", controllers.ModerationsPaginated)
	//r.GET("/moderators", controllers.Moderators)
	//
	//// 管理员相关路由
	//r.GET("/mod", controllers.ModIndex)
	//r.GET("/mod/flagged_stories/:period", controllers.FlaggedStories)
	//r.GET("/mod/flagged_comments/:period", controllers.FlaggedComments)
	//r.GET("/mod/commenters/:period", controllers.Commenters)
	//r.GET("/mod/notes", controllers.ModNotes)
	//r.GET("/mod/notes/:period", controllers.ModNotesPeriod)
	//r.POST("/mod/notes", controllers.CreateModNote)
	//
	//// 其他路由
	//r.GET("/privacy", controllers.Privacy)
	//r.GET("/about", controllers.About)
	//r.GET("/chat", controllers.Chat)
	//r.GET("/stats", controllers.Stats)
	//r.POST("/csp-violation-report", controllers.CSPViolationReport)
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
