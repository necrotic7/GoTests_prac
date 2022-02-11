
package main

type WebsiteChecker func(string) bool
type result struct{
	string
	bool
}


func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool{
	result := make(map[string]bool)
	
	for _, url := range urls{
		go func(u string){
			result[u] = wc(u)
		}(url)
	}
	return result
}