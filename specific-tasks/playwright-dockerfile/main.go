/*
Copyright 2020 Matt Hamilton

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/playwright-community/playwright-go"
)

var Alpine string

const xOffset20 = "0.000273033894595"
const yOffset20 = "0.000000000000041"

// const xOffset13 = "0.03494834106878"
// const yOffset13 = "0.109863281249998"

const xOffset18 = "0.001089390546725"
const yOffset18 = "0.006866455078125"

const xOffset13 = "0.03494834106878"
const yOffset13 = "0.109863281249998"
const zOffset13 = "0.000015659770345"

const res = "qwe"

func main() {
	arr := [4]string{"1", "2", "3", "4"}

	qq := "qwe"
	arr[3] = qq
	fmt.Println(arr)

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("unable to launch playwright: %s\n", err.Error())
	}
	opts := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args: []string{
			"--no-zygote",
			"--single-process",
			"--disable-default-apps",
			"--disable-filesystem",
		},
		ChromiumSandbox: playwright.Bool(false),
	}
	if Alpine != "" {
		opts.ExecutablePath = playwright.String("/usr/bin/chromium-browser")
	}
	browser, err := pw.Chromium.Launch(opts)
	if err != nil {
		log.Fatalf("unable to lauch Chromium: %s\n", err.Error())
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// time.Sleep(time.Minute * 1)
	page.SetDefaultTimeout(150 * 1000)
	if _, err = page.Goto("file:///app/index.html"); err != nil {
		log.Fatalf("could not goto: %v\n", err)
	}
	fmt.Println("success")
	time.Sleep(time.Minute * 10)

	// 10 km
	// offset x 0,0174780845214
	// offset y 0,0549316406251

	// center1x := "55.645374" // "55.743337"
	// center1y := "37.704212" // "37.510063"

	center1x := "55.743337"
	center1y := "37.510063"
	// manager is {yaMap, searchControl} object
	manager, err := page.EvaluateHandle(`async ({centerX, centerY}) => {
		await ymaps.ready()
	
		let yaMap = new ymaps.Map("map", {
			center: [centerX, centerY],
			zoom: 20,

			controls: [],
		})
	
		let searchControl = new ymaps.control.SearchControl({
			options: {
				provider: "yandex#search",
				results: 2000,
			},
		})
		await yaMap.controls.add(searchControl)
	
		return {
			yaMap: yaMap,
			searchControl: searchControl,
		}
	}
	`, map[string]interface{}{
		"centerX": center1x,
		"centerY": center1y,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	// 55.61042565893122 37.59434871875
	// VM18:13 55.68032234106877 37.81407528125

	// 0: 55.743337057262931: 37.51006299999995length: 2[[Prototype]]: Array(0)

	// VM18:14 55.70838865893122 37.40019971875
	// VM18:15 55.778285341068774 37.61992628125

	// VM18:17 55.70837299916087 37.40019971875
	// VM18:18 55.778269681298426 37.61992628125

	// VM18:22 1012

	//manager.yaMap.getBounds()
	storeData, err := page.EvaluateHandle(`async ({manager, xOffset, yOffset, zOffset, centerX, centerY}) => {
		xOffset = Number(xOffset)
		yOffset = Number(yOffset)
		centerX = Number(centerX)
		centerY = Number(centerY)

		await manager.yaMap.setBounds([
			[centerX - xOffset, centerY - yOffset],
			[centerX + xOffset, centerY + yOffset],
		])

		console.log(manager.yaMap.getCenter())

		console.log(centerX - xOffset, centerY - yOffset)
		console.log(centerX + xOffset, centerY + yOffset)

		console.log(centerX - xOffset - zOffset, centerY - yOffset)
		console.log(centerX + xOffset - zOffset, centerY + yOffset)

		await manager.searchControl.search("магазины")

		resultsCount = await manager.searchControl.getResultsCount()

		console.log(resultsCount)
	
		return manager.yaMap.getBounds()
	}
	`, map[string]interface{}{
		"manager": manager,
		"xOffset": xOffset20,
		"yOffset": yOffset20,
		"zOffset": zOffset13,
		"centerX": center1x,
		"centerY": center1y,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	fmt.Println(storeData)
	// prop, _ := storeData.GetProperties()
	// // fmt.Printf("%#v\n", prop2)

	// prop3, _ := prop["0"].GetProperties()
	// prop4, _ := prop["1"].GetProperties()

	// fmt.Printf("[%s, %s] ; [%s, %s]\n", prop3["0"].String(), prop3["1"].String(), prop4["0"].String(), prop4["1"].String())

	// prop, _ := storeData.GetProperties()
	// // fmt.Printf("%#v\n", prop2)

	// prop3 := prop["0"].String()
	// prop4 := prop["1"].String()

	// fmt.Printf("[%s, %s]\n", prop3, prop4)

	// jj := make(map[string]interface{})
	// jj["5"] = 2
	// fmt.Printf("%#v\n", jj)

	// fmt.Printf("%#v\n", storeData)
	// fmt.Println()

	// fmt.Println(storeData.String())
	// fmt.Println()

	// prop, _ := storeData.GetProperties()
	// fmt.Printf("%#v\n", prop)

	// fmt.Println()

	// prop2, _ := prop["1"].GetProperties()
	// fmt.Printf("%#v\n", prop2)

	// fmt.Println(storeData.GetProperty("0"))
	// fmt.Println(storeData.GetProperty("c"))
	// fmt.Println()

	// for i := 0; i < 100; i++ {
	// 	time.Sleep(4 * time.Second)

	// 	res, err := page.EvaluateHandle(`async ({manager, i}) => {
	// 		await manager.searchControl.search("магазины")

	// 		let resultsCount = await manager.searchControl.getResultsCount()

	// 		return resultsCount + i
	// 	}
	// 	`, map[string]interface{}{
	// 		"manager": manager,
	// 		"i":       i,
	// 	})
	// 	if err != nil {
	// 		fmt.Fprint(os.Stderr, err)
	// 	}

	// 	fmt.Println(res)
	// }

	time.Sleep(10500 * time.Second)

	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v\n", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop playwright: %v\n", err)
	}
}
