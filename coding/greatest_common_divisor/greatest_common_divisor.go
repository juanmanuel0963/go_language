package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {

	fmt.Println(str1 + " " + str2)
	sGCDOverall := ""
	//Iterate len times str1 and split str1 and str2 according to the len

	sTest := ""

	if len(str1) < len(str2) {
		sTest = str1
	} else {
		sTest = str2
	}

	for iWordLen := 1; iWordLen <= len(sTest); iWordLen++ {

		iteration := 1
		sFirstWord := ""
		sPlit := ""
		breaked := false
		sGCD := ""

		//fmt.Println("WordLenght: " + strconv.Itoa(iWordLen))

		for iSplitStart := 0; iSplitStart < len(sTest); iSplitStart = iSplitStart + iWordLen {

			iSplitEnd := 0

			if (iSplitStart + iWordLen) <= len(sTest) {
				iSplitEnd = iSplitStart + iWordLen
			} else {
				iSplitEnd = len(sTest)
			}

			sPlit = sTest[iSplitStart:iSplitEnd]
			//fmt.Println("Splited word " + strconv.Itoa(iteration) + " : " + sPlit)

			if iSplitStart == 0 {
				sFirstWord = sPlit
			}

			if sPlit != sFirstWord {
				breaked = true
				break
			}
			iteration += 1
		}

		if !breaked {
			if sPlit != sTest {
				if len(sPlit) > len(sGCD) {
					sGCD = sPlit
					sGCDOverall = sGCD
				}
			}
		}

		//fmt.Println("sGCD: " + sGCD)
		//fmt.Println("---")
	}

	//fmt.Println("sGCDOverall: " + sGCDOverall)
	return sGCDOverall
}

func main() {
	got := gcdOfStrings("NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM", "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM")
	fmt.Println(got)
}
