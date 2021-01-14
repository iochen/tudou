package tudou_test

import (
	"fmt"
	"testing"

	"github.com/iochen/tudou/v2"
)

func TestDecode(t *testing.T) {
	s, err := tudou.Decode("如是我闻：护多重各茶告重路灯蘇究央彌此路多除根万毒栗普福雙夷時貧紛根積尼清參游千贤知诵沙以智央利怖令贤至万矜持寡万释山经僧教孤宝憐进幽三如逝及矜敬數通未多拔資众沙特住持路者栗宝根孤即众橋牟廟胜刚婦倒經令五僧数敬梭藝婦寡度首沙帝依槃陀困足闍闍夢智牟害即央行普琉謹捨修祖")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(s))
}
