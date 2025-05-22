package db

import (
	"context"
	"testing"

	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomNews(t *testing.T) News {
	title := utils.RandomString(10)
	arg := CreateNewsParams{
		Title:       title,
		Slug:        utils.Slugify(title),
		ImageUrl:    "https://picsum.photos/200/300",
		Description: utils.RandomString(100),
		Author:      utils.RandomName(),
		Content:     utils.RandomString(500),
	}
	news, err := testStore.CreateNews(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, news)
	require.NotEmpty(t, news.NewsID)
	require.NotEmpty(t, news.Slug)
	require.NotEmpty(t, news.ImageUrl)
	require.NotEmpty(t, news.Title)
	require.NotEmpty(t, news.Description)
	require.NotEmpty(t, news.Author)
	require.NotEmpty(t, news.Content)
	require.NotEmpty(t, news.CreatedAt)
	return news
}

func TestCreateNews(t *testing.T) {
	createRandomNews(t)
}

func TestGetNews(t *testing.T) {
	news1 := createRandomNews(t)
	news2, err := testStore.GetNews(context.Background(), news1.NewsID)
	require.NoError(t, err)
	require.NotEmpty(t, news2)
	require.Equal(t, news1.NewsID, news2.NewsID)
	require.Equal(t, news1.Slug, news2.Slug)
	require.Equal(t, news1.ImageUrl, news2.ImageUrl)
	require.Equal(t, news1.Title, news2.Title)
	require.Equal(t, news1.Description, news2.Description)
	require.Equal(t, news1.Author, news2.Author)
	require.Equal(t, news1.Content, news2.Content)
}

func TestListNews(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomNews(t)
	}

	arg := ListNewsParams{
		Limit:  5,
		Offset: 5,
	}

	news, err := testStore.ListNews(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, news, 5)

	for _, Flight := range news {
		require.NotEmpty(t, Flight)
	}
}
