package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/playwright-community/playwright-go"
)

func main() {
	err := playwright.Install()

	pw, err := playwright.Run()
	if err != nil {
		log.Fatal(err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v\n", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
	}
	if _, err = page.Goto("file:///home/judas/Work/code-examples/specific-tasks/yandex-maps-parser/index.html"); err != nil {
		log.Fatalf("could not goto: %v\n", err)
	}

	resultsCount, err := page.Evaluate(`async () => {
		var resultsCount = 0

		async function init() {
			await ymaps.ready()
	
			var myMap = new ymaps.Map('map', {
				bounds: [[55.585079, 37.324313], [55.759817, 37.875966]], 
				controls: [],
			});
	
			var searchControl = new ymaps.control.SearchControl({
				options: {
					provider: 'yandex#search',
					results: 2000
				}
			});
			await myMap.controls.add(searchControl)
			await searchControl.search('школы')
	
			resultsCount = await searchControl.getResultsCount()
			console.log("count before: ", resultsCount)
			console.log("zoom: ", myMap.getZoom())
			console.log("bounds: ", myMap.getBounds())
			console.log("center: ", myMap.getCenter())
		}
	
		await init()
	
		return resultsCount
	  }
	`)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(resultsCount)

	time.Sleep(15 * time.Second)

	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v\n", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop playwright: %v\n", err)
	}
}
