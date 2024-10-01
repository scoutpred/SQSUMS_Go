package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var vals []string
var initial int
var sum int = 0

func recursiveinput(max int, curr int) {

	if max == 1 {

		vals = append(vals, inp(curr))
		finalsum(vals, 0)

	} else {
		if curr == 1 {
			curr = increment(curr)
			recursiveinput(max, curr)
		} else {
			if curr <= max {
				curr = increment(curr)
				recursiveinput(max, curr)
			} else {
				finalsum(vals, 0)
			}
		}
	}
}

func finalsum(arr []string, startpos int) {

	var arrlen int = len(arr)

	if arrlen == 1 {
		numlist := strings.Split(arr[startpos], " ")
		perindex(numlist, 1)
	} else {
		if (startpos + 1) <= arrlen {
			numlist := strings.Split(arr[startpos], " ")
			perindex(numlist, 1)
			startpos = startpos + 1
			finalsum(arr, startpos)
		} else {
		}

	}

}

func perindex(arr []string, pos int) {

	initial, _ = strconv.Atoi(arr[0])
	if pos <= initial {
		num, _ := strconv.Atoi(arr[pos])
		if num >= 0 {
			sum = sum + (num * num)
			pos = pos + 1
			perindex(arr, pos)
		} else {
			pos = pos + 1
			perindex(arr, pos)
		}
	} else {
		fmt.Println(sum)
		sum = 0
	}

}

func increment(curr int) int {
	vals = append(vals, inp(curr))
	curr = curr + 1
	return curr
}

func inp(curr int) string {
	var in, lst, res string
	fmt.Scanln(&in)
	var check, err = strconv.ParseInt(in, 10, 0)
	if err != nil {
		fmt.Println("Invalid Input. Code stopped!")
		os.Exit(0)
	} else {

		if check <= 0 {
			fmt.Println(check, " is less than or equal to 0. Code stopped!")
			os.Exit(0)
		} else {
			if check > 100 {
				fmt.Println(check, " is greater than 100. Code Stopped!")
				os.Exit(0)
			} else {
				scanner := bufio.NewScanner(os.Stdin)

				if scanner.Scan() {
					lst = scanner.Text()
					checklist := strings.Split(lst, " ")
					if int64(len(checklist)) > check {
						fmt.Println("Number list exceeds case size. Code Stopped!")
						os.Exit(0)
					} else {
						if int64(len(checklist)) < check {
							fmt.Println("Number list is less than case size. Code Stopped!")
							os.Exit(0)
						} else {
							listcheck(checklist, 0)
						}

					}

				}

			}
			res = fmt.Sprintf("%s %s", in, lst)
		}

	}
	return res

}

func listcheck(check []string, startpos int) {

	if len(check)-1 >= startpos {
		var checknum, err = strconv.ParseInt(check[startpos], 10, 0)

		if err != nil {
			fmt.Println("Invalid Input detected in the array! Code stopped!")
			os.Exit(0)
		} else {
			if checknum < -100 {
				fmt.Println(check[startpos], " is lesser than -100. Code Stopped!")
				os.Exit(0)
			} else {
				if checknum > 100 {
					fmt.Println(check[startpos], " is greater than 100. Code Stopped!")
					os.Exit(0)
				} else {
					startpos = startpos + 1
					listcheck(check, startpos)
				}
			}

		}

	}

}

func main() {
	var i int

	fmt.Scanln(&i)
	if i == 0 {
		fmt.Print("Invalid input for the number of test cases. Code stops.")
	} else {
		if 1 <= i && i <= 100 {
			recursiveinput(i, 1)
		} else {
			fmt.Println("Cannot have more than 100 test cases")

		}
	}
}