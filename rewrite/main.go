package main

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const licnese = `
<!--
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
-->
`

func main() {
	path := os.Args[1]
	file := readFile(path)
	file = removeToc(file)
	file = removeTop(file)
	file = rewriteCodeSample(file)
	file = rewriteFrontMatter(file)
	file = rewriteLink(file)
	file = rewriteTabs(file)
	file = rewriteCenter(file)

	fmt.Print(file)
}

func removeToc(file string) string {
	re := regexp.MustCompile(`(\* This will be replaced by the TOC\n){0,1}{:toc}`)
	return re.ReplaceAllString(file, "")
}

func removeTop(file string) string {
	re := regexp.MustCompile(`{%\s*top\s*%}`)
	return re.ReplaceAllString(file, "{{< top >}}")
}

func rewriteFrontMatter(file string) string {

	position := regexp.MustCompile(`nav-pos: (.+)`)
	weight := 0
	weights := position.FindStringSubmatch(file)

	if weights != nil {
		var err error
		if weight , err = strconv.Atoi(weights[1]); err != nil {
				panic(err)
		}

		weight += 1
	}

	file = position.ReplaceAllString(file, fmt.Sprintf("weight: %d\ntype: docs", weight))

	title := regexp.MustCompile(`title: (.+)`)
	titleMatch := title.FindStringSubmatch(file)
	if titleMatch == nil {
		panic("page does not have title")
	}

	pageTitle := strings.Trim(titleMatch[1], ` "`)
	file = strings.ReplaceAll(file, licnese, licnese+"\n"+"# " + pageTitle+"\n")

	navTitle := regexp.MustCompile(`nav-title: (.+)`)
	if menuTitle := navTitle.FindStringSubmatch(file); menuTitle != nil {
		file = title.ReplaceAllString(file, "title: " + menuTitle[1])
	}

	re := regexp.MustCompile(`nav-id: .+`)
	file = re.ReplaceAllString(file, "")


	file = re.ReplaceAllString(file, "")

	re = regexp.MustCompile(`nav-parent_id: .+`)
	file = re.ReplaceAllString(file, "")

	blank := regexp.MustCompile(`(?m)^\s*$`)

	parts := strings.Split(file, "\n")
	var inFrontMatter = false
	for i := 0; i < len(parts); i++ {
		if parts[i] == "---" {
			if inFrontMatter {
				break
			} else {
				inFrontMatter = true
			}
		}

		if inFrontMatter && blank.MatchString(parts[i]) {
			parts = append(parts[:i], parts[i+1:]...)
		}
	}

	return strings.Join(parts, "\n")
}

func rewriteCodeSample(file string) string {
	re := regexp.MustCompile(`{% highlight (.+) %}`)
	file = re.ReplaceAllString(file, "```$1")

	re = regexp.MustCompile(`{% endhighlight %}`)
	return re.ReplaceAllString(file, "```")
}

func rewriteLink(file string) string {
	re := regexp.MustCompile("{% link (.+)\\.md %}")
	return re.ReplaceAllString(file, `{{< ref "/$1" >}}`)
}

var tabsStart = `<div class="codetabs" markdown="1">[\s]*<div data-lang="(.+)" markdown="1">`

var tabsMiddle = `</div>[\s]*<div data-lang="(.+)" markdown="1">`

var tabsEnd = `</div>[\s]*</div>`

var newTabStart = `{{< tabs "#UUID" >}}
{{< tab "$1" >}}`

var newTabMiddle = `{{< /tab >}}
{{< tab "$1" >}}`

var newTabsEnd = `{{< /tab >}}
{{< /tabs >}}`

func rewriteTabs(file string) string {
	start := regexp.MustCompile(tabsStart)
	middle := regexp.MustCompile(tabsMiddle)
	end := regexp.MustCompile(tabsEnd)

	java := regexp.MustCompile(`{{< tab "java" >}}`)
	scala := regexp.MustCompile(`{{< tab "scala" >}}`)
	python := regexp.MustCompile(`{{< tab "python" >}}`)
	sql := regexp.MustCompile(`{{< tab "sql" >}}`)

	placeholder := regexp.MustCompile(`#UUID`)

	file = start.ReplaceAllString(file, newTabStart)
	file = middle.ReplaceAllString(file, newTabMiddle)
	file = end.ReplaceAllString(file, newTabsEnd)
	file = java.ReplaceAllString(file, `{{< tab "Java" >}}`)
	file = scala.ReplaceAllString(file, `{{< tab "Scala" >}}`)
	file = python.ReplaceAllString(file, `{{< tab "Python" >}}`)
	file = sql.ReplaceAllString(file, `{{< tab "SQL" >}}`)

	return placeholder.ReplaceAllStringFunc(file, func(s string) string {
		return uuid.New().String()
	})
}

func rewriteCenter(file string) string {
	return regexp.
		MustCompile(`<p class="text-center">(.+)?</p>`).
		ReplaceAllString(file, `{{< center >}}\n$1\n{{< /center >}}`)
}

// Read the file into a single string
// or fail hard in case of error.
func readFile(path string) string {
	if contents, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(contents)
	}
}
