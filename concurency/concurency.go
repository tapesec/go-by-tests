package concurency

type (
	WebsiteChecker func(url string) bool
	result         struct {
		string
		bool
	}
)

func CheckWebSites(wc WebsiteChecker, urls []string) map[string]bool {
	checkedUrl := make(map[string]bool)
	resultChan := make(chan result)
	for _, url := range urls {
		go func() {
			resultChan <- result{url, wc(url)}
		}()
	}
	urlsLength := len(urls)
	for i := 0; i < urlsLength; i++ {
		r := <-resultChan
		checkedUrl[r.string] = r.bool
	}
	return checkedUrl
}
