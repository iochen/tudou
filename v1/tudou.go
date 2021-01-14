package tudou

import (
	"errors"
	"strings"

	"golang.org/x/text/encoding/unicode"

	"github.com/iochen/tudou/common/crypto/random"

	"github.com/iochen/tudou/common/crypto/aes"
)

var (
	KEY = []byte("XDXDtudou@KeyFansClub^_^Encode!!")
	IV  = []byte("Potato@Key@_@=_=")
	Map = []rune{
		'滅', '苦', '婆', '娑', '耶', '陀', '跋', '多', '漫', '都', '殿', '悉', '夜', '爍', '帝', '吉',
		'利', '阿', '無', '南', '那', '怛', '喝', '羯', '勝', '摩', '伽', '謹', '波', '者', '穆', '僧',
		'室', '藝', '尼', '瑟', '地', '彌', '菩', '提', '蘇', '醯', '盧', '呼', '舍', '佛', '參', '沙',
		'伊', '隸', '麼', '遮', '闍', '度', '蒙', '孕', '薩', '夷', '迦', '他', '姪', '豆', '特', '逝',
		'朋', '輸', '楞', '栗', '寫', '數', '曳', '諦', '羅', '曰', '咒', '即', '密', '若', '般', '故',
		'不', '實', '真', '訶', '切', '一', '除', '能', '等', '是', '上', '明', '大', '神', '知', '三',
		'藐', '耨', '得', '依', '諸', '世', '槃', '涅', '竟', '究', '想', '夢', '倒', '顛', '離', '遠',
		'怖', '恐', '有', '礙', '心', '所', '以', '亦', '智', '道', '。', '集', '盡', '死', '老', '至'}
	Keywords = []rune{'冥', '奢', '梵', '呐', '俱', '哆', '怯', '諳', '罰', '侄', '缽', '皤'}
)

func Encode(b []byte) (string, error) {
	enc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	u16, err := enc.Bytes(b)
	if err != nil {
		return "", err
	}
	data := []byte(u16)
	en, err := aes.Encode(data, KEY, IV)
	buf := new(strings.Builder)
	for _, i := range en {
		if i < 128 {
			buf.WriteString(string(Map[i]))
		} else {
			buf.WriteString(string(random.RandChoose(Keywords)))
			buf.WriteString(string(Map[i-128]))
		}
	}
	return "佛曰：" + buf.String(), nil
}

func Decode(str string) ([]byte, error) {
	keywordsMap := make(map[rune]bool)
	for _, v := range Keywords {
		keywordsMap[v] = true
	}

	tudouMap := make(map[rune]byte)
	for k, v := range Map {
		tudouMap[v] = byte(k)
	}

	rStr := []rune(str)
	if string(rStr[:3]) != "佛曰：" {
		return nil, errors.New("not valid tudou code")
	}
	rStr = rStr[3:]

	var en []byte
	for i := 0; i < len(rStr); i++ {
		if keywordsMap[rStr[i]] {
			i++
			en = append(en, tudouMap[rStr[i]]+128)
		} else {
			en = append(en, tudouMap[rStr[i]])
		}
	}
	u16, err := aes.Decode(en, KEY, IV)
	if err != nil {
		return nil, err
	}
	dec := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	b, err := dec.Bytes(u16)
	if err != nil {
		return nil, err
	}
	return b, nil
}
