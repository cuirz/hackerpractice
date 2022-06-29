package main

import (
	"math/rand"
	"strconv"
	"strings"
)

//535. TinyURL 的加密与解密
//TinyURL 是一种 URL 简化服务， 比如：当你输入一个 URLhttps://leetcode.com/problems/design-tinyurl时，它将返回一个简化的URLhttp://tinyurl.com/4e9iAk 。请你设计一个类来加密与解密 TinyURL 。
//
//加密和解密算法如何设计和运作是没有限制的，你只需要保证一个 URL 可以被加密成一个 TinyURL ，并且这个 TinyURL 可以用解密方法恢复成原本的 URL 。
//
//实现 Solution 类：
//
//Solution() 初始化 TinyURL 系统对象。
//String encode(String longUrl) 返回 longUrl 对应的 TinyURL 。
//String decode(String shortUrl) 返回 shortUrl 原本的 URL 。题目数据保证给定的 shortUrl 是由同一个系统对象加密的。
//
//
//示例：
//
//输入：url = "https://leetcode.com/problems/design-tinyurl"
//输出："https://leetcode.com/problems/design-tinyurl"
//
//解释：
//Solution obj = new Solution();
//string tiny = obj.encode(url); // 返回加密后得到的 TinyURL 。
//string ans = obj.decode(tiny); // 返回解密后得到的原本的 URL 。
//
//提示：
//
//1 <= url.length <= 10^4
//题目数据保证 url 是一个有效的 URL
type Codec struct {
	urls map[int]string
}

func Constructor() Codec {
	return Codec{
		make(map[int]string),
	}

}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	for {
		key := rand.Int()
		if this.urls[key] == "" {
			this.urls[key] = longUrl
			return "http://tinyurl.com/" + strconv.Itoa(key)
		}
	}

}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	index := strings.LastIndex(shortUrl, "/")
	key, _ := strconv.Atoi(shortUrl[index+1:])
	return this.urls[key]

}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
