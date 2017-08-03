//file svc4.go
package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/sycdtk/gostudy/serviceexit/svc"
)

//代码运行的时候，可以：
//通过输入”s-“或者”start-“+服务名，来启动一个服务 用”c-“或”cancel-“+服务名，来退出指定服务 可以用 “exit”或者Control+C、kill来退出程序（除了kill -9）。

//在此基础上，还可以利用context包实现服务超时退出，利用for range限制服务数量，利用HTTP实现微服务RestFUL信息驱动。由于扩展之后代码增加，显得冗余，这里不再赘述。
//</cancel||c></start||s>
type Program struct {
	ctx        context.Context
	exitFunc   context.CancelFunc
	cancelFunc map[string]context.CancelFunc
	wg         WaitGroupWrapper
}

func main() {
	p := &Program{
		cancelFunc: make(map[string]context.CancelFunc),
	}
	p.ctx, p.exitFunc = context.WithCancel(context.Background())
	svc.Run(p)

}
func (p *Program) Init() error {
	//just demon,do nothing
	return nil
}
func (p *Program) Start() error {
	fmt.Println("本程序将会根据输入,启动或终止服务。")

	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			fmt.Println("程序退出命令:exit;服务启动命令:<start||s>-[name];服务停止命令:<cancel||c>-[name]。请注意大小写!")
			input, _, _ := reader.ReadLine()
			command := string(input)
			switch command {
			case "exit":
				goto OutLoop
			default:
				command, name, err := splitInput(input)
				if err != nil {
					fmt.Println(err)
					continue
				}
				switch command {
				case "start", "s":
					newctx, cancelFunc := context.WithCancel(p.ctx)
					p.cancelFunc[name] = cancelFunc

					p.wg.Wrap(func() {
						Func(newctx, name)
					})

				case "cancel", "c":
					cancelFunc, founded := p.cancelFunc[name]
					if founded {
						cancelFunc()
					}
				}
			}
		}
	OutLoop:
		//由于程序退出被Run的os.Notify阻塞，因此调用以下方法通知退出代码执行。
		svc.Interrupt()
	}()
	return nil
}
func (p *Program) Stop() error {
	p.exitFunc()
	p.wg.Wait()
	fmt.Println("所有服务终止,程序退出!")
	return nil
}

//用来转换输入字符串为输入命令
func splitInput(input []byte) (command, name string, err error) {
	line := string(input)
	strs := strings.Split(line, "-")
	if strs == nil || len(strs) != 2 {
		err = errors.New("输入不符合规则。")
		return
	}
	command = strs[0]
	name = strs[1]
	return
}

// 一个简单的循环方法,模拟被加载、释放的微服务
func Func(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			goto OutLoop
		case <-time.Tick(time.Second * 2):
			fmt.Printf("%s is running.\n", name)
		}
	}
OutLoop:
	fmt.Printf("%s is end.\n", name)
}

//WaitGroup封装结构
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(f func()) {
	w.Add(1)
	go func() {
		f()
		w.Done()
	}()
}
