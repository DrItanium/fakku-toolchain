package fakku

import (
	"fmt"
)

type UserApiFunction struct {
	Name string
}

func (c UserApiFunction) Construct() string {
	return fmt.Sprintf("%s/users/%s", ApiHeader, c.Name)
}

func GetUserProfile(name string) (*UserProfile, error) {
	var c User
	url := UserApiFunction{Name: name}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		//cheat a little bit :D
		return &(*c.Profile), nil
	}
}

type User struct {
	Profile *UserProfile `json:"user"`
}

// DAMN, I can't just have Go convert these to bools for me.
// I'll need to do the conversion manually
type UserProfile struct {
	Username          string `json:"user_username"`
	Url               string `json:"user_url"`
	Rank              string `json:"user_rank"`
	Avatar            string `json:"user_avatar"`
	AvatarWidth       uint   `json:"user_avatar_width"`
	AvatarHeight      uint   `json:"user_avatar_height"`
	RegistrationDate  uint   `json:"user_registration_date"`
	LastVisit         uint   `json:"user_last_visit"`
	Subscribed        uint   `json:"user_subscribed"`
	Timezone          int    `json:"user_timezone"`
	Posts             uint   `json:"user_posts"`
	Topics            uint   `json:"user_topics"`
	Comments          uint   `json:"user_comments"`
	Signature         string `json:"user_signature"`
	ForumReputation   int    `json:"user_forum_reputation"`
	CommentReputation int    `json:"user_comment_reputation"`
	Gold              uint   `json:"user_gold"`
	Online            uint   `json:"user_online"`
}

type UserFavoritesApiFunction struct {
	UserApiFunction
	SupportsPagination
}

func (c UserFavoritesApiFunction) Construct() string {
	base := fmt.Sprintf("%s/favorites", c.UserApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserFavoritesPage(user string, page uint) (*UserFavorites, error) {
	var c UserFavorites
	url := UserFavoritesApiFunction{
		UserApiFunction:    UserApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserFavorites(user string) (*UserFavorites, error) {
	return GetUserFavoritesPage(user, 0)
}

type UserFavorites struct {
	Favorites []*Content `json:"favorites"`
	Total     uint       `json:"total"`
	Pages     uint       `json:"pages"`
}

type UserAchievementsApiFunction struct {
	UserApiFunction
}

func (c UserAchievementsApiFunction) Construct() string {
	return fmt.Sprintf("%s/achievements", c.UserApiFunction.Construct())
}

func GetUserAchievements(user string) (*UserAchievements, error) {
	var c UserAchievements
	url := UserAchievementsApiFunction{
		UserApiFunction: UserApiFunction{Name: user},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type UserAchievements struct {
	Achievements []*UserAchievement `json:"achievements"`
	Total        uint               `json:"total"`
}
type UserAchievement struct {
	Name        string `json:"achievement_name"`
	Description string `json:"achievement_description"`
	Icon        string `json:"achievement_icon"`
	Class       string `json:"achievement_class"`
	Date        uint   `json:"achievement_date"`
}

type UserPostsApiFunction struct {
	UserApiFunction
	SupportsPagination
}

func (c UserPostsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/posts", c.UserApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserPostsPage(user string, page uint) (*UserPosts, error) {
	var c UserPosts
	url := UserPostsApiFunction{
		UserApiFunction:    UserApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserPosts(user string) (*UserPosts, error) {
	return GetUserPostsPage(user, 0)
}

type UserPosts struct {
	Posts []*UserPost `json:"posts"`
	Total uint        `json:"total"`
	Pages uint        `json:"pages"`
}

type UserPost struct {
	Id         uint   `json:"post_id"`
	Date       uint   `json:"post_date"`
	Text       string `json:"post_text"`
	Reputation int    `json:"post_reputation"`
	TopicTitle string `json:"post_topic_title"`
	TopicUrl   string `json:"post_topic_url"`
}

type UserTopicsApiFunction struct {
	UserApiFunction
	SupportsPagination
}

func (c UserTopicsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/topics", c.UserApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserTopicsPage(user string, page uint) (*UserTopics, error) {
	var c UserTopics
	url := UserTopicsApiFunction{
		UserApiFunction:    UserApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserTopics(user string) (*UserTopics, error) {
	return GetUserTopicsPage(user, 0)
}

type UserTopics struct {
	Topics []*UserTopic `json:"topics"`
	Total  uint         `json:"total"`
	Pages  uint         `json:"pages"`
}

type UserTopic struct {
	Title       string `json:"topic_title"`
	Url         string `json:"topic_url"`
	Time        uint   `json:"topic_time"`
	Replies     uint   `json:"topic_replies"`
	Status      uint   `json:"topic_status"`
	Poll        uint   `json:"topic_poll"`
	LastPostId  uint   `json:"topic_last_post_id"`
	PostPreview string `json:"topic_post_preview"`
	Poster      string `json:"topic_poster"`
	PosterUrl   string `json:"topic_poster_url"`
}
type UserCommentsApiFunction struct {
	UserApiFunction
	SupportsPagination
}

func (c UserCommentsApiFunction) Construct() string {
	base := fmt.Sprintf("%s/comments", c.UserApiFunction.Construct())
	return PaginateString(base, c.Page)
}

func GetUserCommentsPage(user string, page uint) (*UserComments, error) {
	var c UserComments
	url := UserCommentsApiFunction{
		UserApiFunction:    UserApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}

}

func GetUserComments(user string) (*UserComments, error) {
	return GetUserCommentsPage(user, 0)
}

type UserComments struct {
	Comments []*UserComment `json:"comments"`
	Total    uint           `json:"total"`
	Pages    uint           `json:"pages"`
}

type UserComment struct {
	Id          uint   `json:"comment_id"`
	AttachedId  uint   `json:"comment_attached_id"`
	Reputation  int    `json:"comment_reputation"`
	Text        string `json:"comment_string"`
	Date        uint   `json:"comment_date"`
	ContentName string `json:"comment_content_name"`
	ContentUrl  string `json:"comment_content_url"`
}
