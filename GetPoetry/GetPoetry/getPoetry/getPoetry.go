package getPoetry

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/mongo"
)

var Count int

type Poetry struct {
	Title   string
	Dynasty string
	Author  string
	Content string
}

func GetContent(body string) []string {
	reg := regexp.MustCompile(`<div class="conview conview_main show"><div>(.*?)</div></div>`)
	regChinese := regexp.MustCompile(`[^\x00-\xff]+`)
	var res []string

	for _, v := range reg.FindAllString(body, -1) {
		var tmpChinese string
		for _, vChinese := range regChinese.FindAllString(v, -1) {
			tmpChinese = fmt.Sprintf("%v%v\n", tmpChinese, vChinese)
			//fmt.Println(vChinese)
		}
		res = append(res, tmpChinese)
		//fmt.Printf(tmpChinese)
	}
	return res
}

func GetTitle(body string) []string {
	reg := regexp.MustCompile(`<a href="/poetry/[0-9]+.html" target=_blank>(.*?)</a>`)
	regChinese := regexp.MustCompile(`[^\x00-\xff]+`)
	var res []string

	for _, v := range reg.FindAllString(body, -1) {
		var tmpChinese string
		for _, vChinese := range regChinese.FindAllString(v, -1) {
			tmpChinese = fmt.Sprintf("%v%v\n", tmpChinese, vChinese)
			//fmt.Println(vChinese)
		}
		res = append(res, tmpChinese)
		//fmt.Printf(vtmpChinese)
	}
	return res
}

func GetAuthor(body string) []string {
	reg := regexp.MustCompile(`<A href="/poets/[0-9]+.html" target=_blank>([^\x00-\xff]+)</A>`)
	regChinese := regexp.MustCompile(`[^\x00-\xff]+`)
	var res []string

	for _, v := range reg.FindAllString(body, -1) {
		var tmpChinese string
		for _, vChinese := range regChinese.FindAllString(v, -1) {
			tmpChinese = fmt.Sprintf("%v%v\n", tmpChinese, vChinese)
			//fmt.Println(vChinese)
		}
		res = append(res, tmpChinese)
		//fmt.Printf(tmpChinese)
	}
	return res
}

func GetDynasty(body string) []string {
	reg := regexp.MustCompile(`<div class="zz">(\s*.*?)</span>`)
	regChinese := regexp.MustCompile(`[^\x00-\xff]+`)
	var res []string

	for _, v := range reg.FindAllString(body, -1) {
		var tmpChinese string
		for _, vChinese := range regChinese.FindAllString(v, -1) {
			tmpChinese = fmt.Sprintf("%v%v\n", tmpChinese, vChinese)
			//fmt.Println(vChinese)
		}
		res = append(res, tmpChinese)
		//fmt.Printf(tmpChinese)
	}
	return res
}

func GetPoetryToFile(body, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	titles := GetTitle(body)
	dynastys := GetDynasty(body)
	authors := GetAuthor(body)
	contents := GetContent(body)
	if len(titles) < 10 || len(dynastys) < 10 || len(authors) < 10 || len(contents) < 10 {
		return
	}

	for i := 0; i < 10; i++ {
		info := fmt.Sprintf("----------------------------number %v----------------------------\n", Count)
		file.WriteString(info)
		file.WriteString(titles[i])
		file.WriteString(dynastys[i])
		file.WriteString(authors[i])
		file.WriteString(contents[i])
		Count++
	}
	file.Close()
}

func GetPoetryToMongo(body string, client mongo.Client) {
	collection := client.Database("poetrys").Collection("poetrydata")
	var ins []interface{}

	titles := GetTitle(body)
	dynastys := GetDynasty(body)
	authors := GetAuthor(body)
	contents := GetContent(body)
	if len(titles) < 10 || len(dynastys) < 10 || len(authors) < 10 || len(contents) < 10 {
		return
	}

	for i := 0; i < 10; i++ {
		tmp := Poetry{Title: titles[i], Dynasty: dynastys[i], Author: authors[i], Content: contents[i]}
		ins = append(ins, tmp)
	}

	_, err := collection.InsertMany(context.TODO(), ins)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted documents successfully.")
}
