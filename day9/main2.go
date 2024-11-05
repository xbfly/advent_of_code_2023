package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("day9input.txt")
	scanner := bufio.NewScanner(file)

	var nums [][]int
	for scanner.Scan() {
		numString := scanner.Text()
		numStrSplit := strings.Split(numString, " ")

		ns := []int{}
		for _, v := range numStrSplit {
			f, _ := strconv.Atoi(string(v))
			ns = append(ns, f)
		}

		nums = append(nums, ns)
	}

	sum := 0
	for _, n := range nums {
		sum += differences(n)
	}

	fmt.Println(sum)
}

func differences(n []int) int {

	diff := [][]int{n}

	ds := n
	notZero := false
	for !notZero {
		newDiff := []int{}
		for i, j := 0, 1; i < len(ds)-1 && j < len(ds); i, j = i+1, j+1 {
			newDiff = append(newDiff, ds[j]-ds[i])
		}

		diff = append([][]int{newDiff}, diff...)
		ds = newDiff

		notZero = true
		for _, r := range newDiff {
			if r != 0 {
				notZero = false
				break
			}
		}

	}

	for i, j := 0, 1; i < len(diff)-1 && j < len(diff); i, j = i+1, j+1 {
		x := diff[i][0]
		y := diff[j][0]

		z := y - x
		diff[j] = append([]int{z}, diff[j]...)
	}

	w := len(diff)

	return diff[w-1][0]
}
