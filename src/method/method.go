package method

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

func HeaderFileSlug(header string) string {
	header = strings.ToLower(header)
	header = strings.Replace(header, " ", "-", -1)
	header = strings.Replace(header, "/", "-", -1)
	return header + ".sh"
}

//ProjectNamebyEmail function used for generate the default project name
func ProjectNamebyEmail(email string) string {

	list := strings.Split(email, "@")
	if list[0] == "" {
		return slug.Make("default")
	}
	return slug.Make(list[0] + "-Default")
}

func CheckURL(url string) (string, error) {

	spliturl := strings.Split(url, "://")
	schema := "https://"

	if len(spliturl) > 1 {
		url = spliturl[1]
	} else {
		url = spliturl[0]
	}
	method := "GET"
	// splitdomain := strings.Split(url, ".")
	var domain string

	domain = schema + url

	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(method, domain, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//	req.Header.Set("Location", "Chandigarh")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		//	return "", err
	}

	if res == nil {

		if !strings.HasPrefix(url, "www") {
			domain = schema + "www." + url
		}
		req, err = http.NewRequest(method, domain, nil)
		if err != nil {
			log.Println(err)
			return "", err
		}
		res, err = client.Do(req)
		if err != nil {
			log.Println(err)
			//	return "", err
		}
		if res == nil {
			if strings.HasPrefix(spliturl[0], "https") {
				log.Println(err)
				return domain, err
			}
			domain = strings.Replace(domain, "https", "http", 1)
			req, err = http.NewRequest(method, domain, nil)
			if err != nil {
				log.Println(err)
				return "", err
			}
			res, err = client.Do(req)
			if err != nil {
				log.Println(err)
				return "", err
			}

			defer res.Body.Close()
			if res.StatusCode > 300 {
				return res.Request.URL.Scheme + "://" + res.Request.URL.Hostname(), nil
			}
			return res.Request.URL.Scheme + "://" + res.Request.URL.Hostname(), nil
		}
	}
	defer res.Body.Close()
	if res.StatusCode > 300 {
		log.Println(res.Request.URL.Scheme+"://"+res.Request.URL.Hostname(), nil)
		return res.Request.URL.Scheme + "://" + res.Request.URL.Hostname(), nil
	}
	return res.Request.URL.Scheme + "://" + res.Request.URL.Hostname(), nil
}
