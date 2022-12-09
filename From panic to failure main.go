package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// не меняйте импорты, они нужны для проверки
// import (
//     "bufio"
//     "errors"
//     "fmt"
//     "io/ioutil"
//     "os"
//     "reflect"
//     "runtime"
//     "strconv"
//     "strings"
// )

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	var err error
	acc, trans, err = parseInput()
	defer func() {
		fmt.Print("-> ")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(acc, trans)
		}
	}()
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	var trans []int
	accSrc, transSrc := readInput()
	err, acc := parseAccount(accSrc)
	if err != nil {
		return acc, trans, err
	}
	trans, err = parseTransactions(transSrc)
	if err != nil {
		return acc, trans, err
	}
	return acc, trans, err
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (error, account) {
	var overdraft int
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return err, account{balance, overdraft}
	}
	overdraft, err = strconv.Atoi(parts[1])
	if err != nil {
		return err, account{balance, overdraft}
	}
	if overdraft < 0 {
		err = errors.New("expect overdraft >= 0")
		return err, account{balance, overdraft}
	}
	if balance < -overdraft {
		err = errors.New("balance cannot exceed overdraft")
		return err, account{balance, overdraft}
	}
	return err, account{balance, overdraft}
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	var err error
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return trans, err
		}
		trans[idx] = t
	}
	return trans, err
}
