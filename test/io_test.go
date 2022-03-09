/**
 * @Description
 * @Author r0cky
 * @Date 2021/12/24 16:13
 **/
package test

import (
	"context"
	"fmt"
	"os"
	"sec-tools/bin/misc"
	"sec-tools/secio"
	"syscall"
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	//adapter := secio.NewAdapter()
}

func TestFragment(t *testing.T) {
	fmt.Println(secio.Fragment("D:\\Alienware\\桌面\\ip.txt", 51))
}

func TestProcess(t *testing.T) {
	p, _ := os.FindProcess(55536)
	_ = p.Signal(syscall.SIGINT)
}

func TestCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var finish = make(chan struct{})
	go func() {
		defer close(finish)
		time.Sleep(time.Second * 5)
	}()
	time.AfterFunc(time.Second*3, cancel)
	select {
	case <-ctx.Done(): //超时/被 cancel 结束
		fmt.Println(123123)
	case <-finish: //正常结束
		fmt.Println(123)
	}
}

func TestToMap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "111"
	fmt.Println(misc.ToMap(m))
}
