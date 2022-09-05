package main

import (
	"fmt"
)

func main() {
	//Example_Twitter()
	//Example_Twitter2()
	//Example_Twitter3()
	Example_Twitter4()
	//Example_Twitter5()
}

type Twitter struct {
	usersTweets   map[int]int
	allTweets     []int
	subscriptions map[int][]int
}

func Constructor() Twitter {
	var twitter Twitter
	twitter.usersTweets = map[int]int{}
	twitter.subscriptions = map[int][]int{}
	return twitter
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.usersTweets[tweetId] = userId
	this.allTweets = append(this.allTweets, tweetId)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	newsFeed := []int{}
	listOfFollowed := append(this.subscriptions[userId], userId)
	fmt.Println(listOfFollowed)
	for i := len(this.allTweets) - 1; i >= 0 || len(newsFeed) < 10; i-- {
		for _, userFollowed := range listOfFollowed {
			if this.usersTweets[this.allTweets[i]] == userFollowed {
				newsFeed = append(newsFeed, this.allTweets[i])
			}
		}
		if len(newsFeed) == 10 {
			return newsFeed
		}
	}
	return newsFeed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	for _, followee := range this.subscriptions[followerId] {
		if followee == followeeId {
			return
		}
	}
	this.subscriptions[followerId] = append(this.subscriptions[followerId], followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for i, followee := range this.subscriptions[followerId] {
		if followee == followeeId {
			this.subscriptions[followerId] = append(this.subscriptions[followerId][0:i], this.subscriptions[followerId][i+1:]...)
		}
	}
}

func Example_Twitter() {
	t := Constructor()
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	fmt.Println(t.GetNewsFeed(2))
	t.PostTweet(2, 22)
	fmt.Println(t.GetNewsFeed(2))
	t.PostTweet(3, 33)
	t.Follow(1, 3)
	t.PostTweet(1, 11)
	fmt.Println(t.GetNewsFeed(1))
	fmt.Println(t.GetNewsFeed(2))
	fmt.Println(t.GetNewsFeed(3))
	// OutPut:
	// []
	// []
	// [22]
	// [11 33 22]
	// [22]
	// [33]
}

func Example_Twitter2() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	fmt.Println(twitter.GetNewsFeed(1))
	// OutPut:
	// [5]
	// [6 5]
	// [5]
}

// need to pull tweets of followee when followed
func Example_Twitter3() {
	twitter := Constructor()
	twitter.PostTweet(1, 1)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(2, 1)
	fmt.Println(twitter.GetNewsFeed(2))
	twitter.Unfollow(2, 1)
	fmt.Println(twitter.GetNewsFeed(2))
	// OutPut:
	// [1]
	// [1]
	// []
}

// try follow twice
func Example_Twitter4() {
	twitter := Constructor()
	twitter.PostTweet(2, 5)
	twitter.Follow(1, 2)
	twitter.Follow(1, 2)
	fmt.Println(twitter.GetNewsFeed(1))

	// OutPut:
	// [5]
}

// test the order of tweets
func Example_Twitter5() {
	twitter := Constructor()
	twitter.PostTweet(2, 22)
	twitter.PostTweet(3, 33)
	twitter.PostTweet(4, 44)
	twitter.PostTweet(5, 55)

	twitter.Follow(1, 2)
	twitter.Follow(1, 3)
	twitter.Follow(1, 4)
	twitter.Follow(1, 5)

	twitter.Unfollow(1, 3)
	fmt.Println(twitter.GetNewsFeed(1))

	twitter.Follow(1, 3)
	fmt.Println(twitter.GetNewsFeed(1))

	// OutPut:
	// [55 44 22]
	// [55 44 33 22]
}
