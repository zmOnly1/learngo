package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`<a data-[\w-]+="" target="_self" class="user f-cl" href="//www.zhenai.com/n/login\?channelId=\d+&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F(\d+)">`)

const guessData = `<a data-v-4a9ca87a="" target="_self" class="user f-cl" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1736961045">`

//var re = regexp.MustCompile(`<div data-[\w-]+="" class="m-btn purple">([^<]+)</div>`)
//const data = `<div data-v-8b1eac0c="" class="purple-btns"><div data-v-8b1eac0c="" class="m-btn purple">未婚</div><div data-v-8b1eac0c="" class="m-btn purple">32岁</div><div data-v-8b1eac0c="" class="m-btn purple">魔羯座(12.22-01.19)</div><div data-v-8b1eac0c="" class="m-btn purple">158cm</div><div data-v-8b1eac0c="" class="m-btn purple">工作地:广州天河区</div><div data-v-8b1eac0c="" class="m-btn purple">月收入:2-5万</div><div data-v-8b1eac0c="" class="m-btn purple">生物工程</div><div data-v-8b1eac0c="" class="m-btn purple">大学本科</div></div>`

func main() {
	matches := re.FindAllSubmatch([]byte(guessData), -1)

	for _, m := range matches {
		for _, subM := range m {
			fmt.Printf("node: %s ", subM)
		}
		fmt.Println()
	}
}
