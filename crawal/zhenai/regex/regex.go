package main

import (
	"fmt"
	"regexp"
)

const ageRe = `<div data-[\w-]+="" class="m-btn purple">([^<]+)</div>`
const data = `<div data-v-8b1eac0c="" class="purple-btns"><div data-v-8b1eac0c="" class="m-btn purple">未婚</div><div data-v-8b1eac0c="" class="m-btn purple">32岁</div><div data-v-8b1eac0c="" class="m-btn purple">魔羯座(12.22-01.19)</div><div data-v-8b1eac0c="" class="m-btn purple">158cm</div><div data-v-8b1eac0c="" class="m-btn purple">工作地:广州天河区</div><div data-v-8b1eac0c="" class="m-btn purple">月收入:2-5万</div><div data-v-8b1eac0c="" class="m-btn purple">生物工程</div><div data-v-8b1eac0c="" class="m-btn purple">大学本科</div></div>`

func main() {
	re := regexp.MustCompile(ageRe)
	matches := re.FindAllSubmatch([]byte(data), -1)

	for _, m := range matches {
		for _, subM := range m {
			fmt.Printf("node: %s ", subM)
		}
		fmt.Println()
	}
}
