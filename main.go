/*input
*/
package main

import (
  "fmt"
  "io/ioutil"
  "regexp"
  "strings"
)

func toTrans(srcstr string) string {
	//fmt.Println(srcstr)
	substr := strings.Split(srcstr, " ")

	if substr[0] == "reduced" {
		substr[0] = "giảm"
		substr[1] = "từ"
		substr[len(substr)-1] = "xuống"
	}

	if substr[0] == "increased" {
		substr[0] = "tăng"
		substr[1] = "từ"
		substr[len(substr)-1] = "lên"
	}

	if substr[0] == "reduction" {
		substr[1] = "thay đổi từ"
		substr[len(substr)-1] = "thành"
	}

	return strings.Join(substr, " ");
}

func main() {
	//take input
	data, err := ioutil.ReadFile("src.html")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	text := string(data)

	srcdicttext, err := ioutil.ReadFile("basicdict.inp")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	srcdict := strings.Split(string(srcdicttext), ",")


	desdicttext, err := ioutil.ReadFile("basicdict.out")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	desdict := strings.Split(string(desdicttext), ",")
	//take input

	//processing
	var re *regexp.Regexp
	for i:=0; i<len(srcdict); i++ {
		re = regexp.MustCompile(srcdict[i])
		text = re.ReplaceAllString(text, desdict[i])
	}


	re = regexp.MustCompile("[a-z]+ from.*?to")
	text = re.ReplaceAllStringFunc(text, toTrans)
	//processing

	//write output
	err = ioutil.WriteFile("des.html", []byte(text), 0644)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//write output
}