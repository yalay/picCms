package controllers

import (
	"testing"
	"bufio"
	"strings"
	"os"
	"fmt"
	"sync"
	"time"
)

func TestTranslateEng(t *testing.T) {
	f, err := os.Open("titles.txt")
	if err != nil {
		t.Error(err)
		return
	}
	var wg sync.WaitGroup
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		go func(text string) {
			wg.Add(1)
			engText, err := TranslateToEng(text)
			if err != nil {
				fmt.Println(text, err)
			} else {
				fmt.Printf("%s -> %s\n", text, engText)
			}
			wg.Done()
		}(line)
	}
	time.Sleep(2*time.Second)
	wg.Wait()
}


func TestTranslateCht(t *testing.T) {
	f, err := os.Open("titles.txt")
	if err != nil {
		t.Error(err)
		return
	}
	var wg sync.WaitGroup
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		go func(text string) {
			wg.Add(1)
			engText, err := TranslateToCht(text)
			if err != nil {
				fmt.Println(text, err)
			} else {
				fmt.Printf("%s -> %s\n", text, engText)
			}
			wg.Done()
		}(line)
	}
	time.Sleep(2*time.Second)
	wg.Wait()
}
