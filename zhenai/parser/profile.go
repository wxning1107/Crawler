package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+岁)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(月收入:[^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+kg)</div>`)
var genderRe = regexp.MustCompile(`"genderString":"([^"]+)"`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^(]+)([0-9.-]*)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(中专|大专|高中及以下|大学本科|硕士|博士)</div>`)
var ocupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:[^<]+</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>(和家人同住|已购房|租房)</div>`)
var carRe = regexp.MustCompile(`div class="m-btn pink" data-v-8b1eac0c>(未买车|已买车)</div>`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
		profile.Age = age
	}

	if age, err := strconv.Atoi(extractString(contents, heightRe)); err != nil {
		profile.Height = age
	}

	if age, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
		profile.Weight = age
	}

	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, ocupationRe)
	profile.Education = extractString(contents, educationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	//profile.Name = extractString(contents, nameRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
