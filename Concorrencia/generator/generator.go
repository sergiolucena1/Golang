package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// google i/o 2012 - Go concurrency Patterns

// <- chan  - canal somente-leitura

// pegar os titulos baseados nas urls
func titulo(urls ... string) <- chan string{
	c := make (chan string )// criei o channel
	for _ , url := range urls{
		go func(url string) { // funcao anonima

			resp, _ := http.Get(url)// resposta: ignora o erro(_) , recebe http.get passando uma url

			html, _ := ioutil.ReadAll(resp.Body) // resposta do body

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
				// regex : pega o titulo de uma pagina
			c <- r.FindStringSubmatch(string(html))[1]

		}(url)
	}
	return c
}

func main (){
	//generator
	t1:= titulo("https://www.cod3r.com.br", "https://www.google.com.br")
	t2:= titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("primeiros :", <-t1, "|", <-t2)
	fmt.Println("Segundos :", <-t1, "|", <-t2)

}