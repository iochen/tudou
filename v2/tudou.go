package tudou

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/saracen/go7z"

	"github.com/iochen/tudou/common/crypto/aes"
)

var (
	KEY = []byte("XDXDtudou@KeyFansClub^_^Encode!!")
	IV  = []byte("Potato@Key@_@=_=")
	Map = []rune{
		'謹', '穆', '僧', '室', '藝', '瑟', '彌', '提', '蘇', '醯', '盧', '呼', '舍', '參', '沙', '伊',
		'隸', '麼', '遮', '闍', '度', '蒙', '孕', '薩', '夷', '他', '姪', '豆', '特', '逝', '輸', '楞',
		'栗', '寫', '數', '曳', '諦', '羅', '故', '實', '訶', '知', '三', '藐', '耨', '依', '槃', '涅',
		'竟', '究', '想', '夢', '倒', '顛', '遠', '怖', '恐', '礙', '以', '亦', '智', '盡', '老', '至',
		'吼', '足', '幽', '王', '告', '须', '弥', '灯', '护', '金', '刚', '游', '戏', '宝', '胜', '通',
		'药', '师', '琉', '璃', '普', '功', '德', '山', '善', '住', '过', '去', '七', '未', '来', '贤',
		'劫', '千', '五', '百', '万', '花', '亿', '定', '六', '方', '名', '号', '东', '月', '殿', '妙',
		'尊', '树', '根', '西', '皂', '焰', '北', '清', '数', '精', '进', '首', '下', '寂', '量', '诸',
		'多', '释', '迦', '牟', '尼', '勒', '阿', '閦', '陀', '中', '央', '众', '生', '在', '界', '者',
		'行', '于', '及', '虚', '空', '慈', '忧', '各', '令', '安', '稳', '休', '息', '昼', '夜', '修',
		'持', '心', '求', '诵', '此', '经', '能', '灭', '死', '消', '除', '毒', '害', '高', '开', '文',
		'殊', '利', '凉', '如', '念', '即', '说', '曰', '帝', '毘', '真', '陵', '乾', '梭', '哈', '敬',
		'禮', '奉', '祖', '先', '孝', '雙', '親', '守', '重', '師', '愛', '兄', '弟', '信', '朋', '友',
		'睦', '宗', '族', '和', '鄉', '夫', '婦', '教', '孫', '時', '便', '廣', '積', '陰', '難', '濟',
		'急', '恤', '孤', '憐', '貧', '創', '廟', '宇', '印', '造', '經', '捨', '藥', '施', '茶', '戒',
		'殺', '放', '橋', '路', '矜', '寡', '拔', '困', '粟', '惜', '福', '排', '解', '紛', '捐', '資',
	}
)

func Decode(str string) ([]byte, error) {
	tudouMap := make(map[rune]byte)
	for k, v := range Map {
		tudouMap[v] = byte(k)
	}

	rStr := []rune(str)
	if string(rStr[:5]) != "如是我闻：" {
		return nil, errors.New("not valid tudou code")
	}
	rStr = rStr[5:]

	var en []byte
	for i := 0; i < len(rStr); i++ {
		en = append(en, tudouMap[rStr[i]])
	}

	dec, err := aes.Decode(en, KEY, IV)
	if err != nil {
		return nil, err
	}

	depressed, err := go7z.NewReader(bytes.NewReader(dec), int64(len(dec)))
	if err != nil {
		return nil, err
	}
	_, err = depressed.Next()
	if err != nil {
		return nil, err
	}

	all, err := ioutil.ReadAll(depressed)
	if err != nil {
		return nil, err
	}

	return all, nil
}
