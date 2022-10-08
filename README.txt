goes to site from list (currently in file "links.go")
1/2 runs tags, finding "a" type ("<a href="xxx.xx" >")
check every tag, looking for "href" in Key
			if a.Key == "href" {
				links = a.Val
			}
collect value:
a.Val
2/2 also can check by value in tag ("<xxx="chordsBlock">"):
			if a.Val == "chordsBlock" {
				}
*Dont collect text "<a href="xxx.xx">TEXT</a>"