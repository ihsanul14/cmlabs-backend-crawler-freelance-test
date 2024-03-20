package usecase

import (
	"cmlabs-backend-crawler-freelance-test/entity/http"
	"cmlabs-backend-crawler-freelance-test/entity/memory"
	"cmlabs-backend-crawler-freelance-test/framework/utils"
	"context"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type IUsecase interface {
	Crawl(context.Context, []string) error
}

type Usecase struct {
	Http    http.IHttp
	Memory  memory.IMemory
	Visited map[string]bool
}

type Request struct {
	Url    string
	Domain string
	Node   *html.Node
	Key    int
}

func NewUsecase(http http.IHttp, memory memory.IMemory) IUsecase {
	return &Usecase{
		Http:    http,
		Memory:  memory,
		Visited: make(map[string]bool),
	}
}

func (u *Usecase) Crawl(ctx context.Context, req []string) error {
	for i, v := range req {
		go func(ctx context.Context, req []string) {
			r := Request{
				Url:    v,
				Domain: utils.GetDomain(v),
				Key:    i,
			}
			u.crawl(ctx, r)
			u.Visited = make(map[string]bool)
			log.Printf("%s completed... \n", v)
		}(ctx, req)
	}
	return nil
}

func (u *Usecase) crawl(ctx context.Context, req Request) error {
	res, err := u.Http.Call(ctx, req.Url)
	if err != nil {
		log.Println(err)
		return err
	}
	data := memory.SaveRequest{
		Body:     res.Body,
		Domain:   req.Domain,
		FileName: utils.GenerateFileName(req.Url),
	}
	u.Memory.Save(ctx, data)

	for _, l := range res.Link {
		lUrl, err := url.Parse(l)
		if err != nil {
			log.Printf("usecase.crawl: %v \n", err)
			return err
		}
		if strings.Contains(lUrl.Hostname(), req.Domain) {
			if _, ok := u.Visited[l]; !ok {
				u.Visited[l] = false
			}
		}
	}

	for k, v := range u.Visited {
		if !v && strings.Contains(k, req.Domain) {
			u.Visited[k] = true
			req.Url = k
			u.crawl(ctx, req)
		}
	}
	return nil
}
