# 有限状态机FSM的原理与GO的实现

[![img](https://upload.jianshu.io/users/upload_avatars/5709266/ad5afb07-d947-4ed2-9c94-f27f8cdd46d5?imageMogr2/auto-orient/strip|imageView2/1/w/96/h/96/format/webp)](https://www.jianshu.com/u/6078656f6748)

[陈康stozen](https://www.jianshu.com/u/6078656f6748)关注

0.5792017.07.23 02:26:00字数 698阅读 7,225

![img](https://upload-images.jianshu.io/upload_images/5709266-36eff66b41d11c07.png?imageMogr2/auto-orient/strip|imageView2/2/w/900/format/webp)

有限状态机(Finite-state machine, 简写FSM)又可以称作有限状态自动机。它必须是可以附着在某种事物上的，且该事物的状态是有限的，通过某些触发事件，会让其状态发生转换。为此，有限状态机就是描述这些有限的状态和触发事件及转换行为的数学模型。

## 有限状态机组成

**有限状态机有两个必要的特点，一是离散的，二是有限的。基于这两点，现实世界上绝大多数事物因为复杂的状态而无法用有限状态机表示。**

而描述事物的有限状态机模型的元素由以下组成：

- 状态(State)：事物的状态，包括初始状态和所有事件触发后的状态
- 事件(Event)：触发状态变化或者保持原状态的事件
- 行为或转换(Action/Transition)：执行状态转换的过程
- 检测器(Guard)：检测某种状态要转换成另一种状态的条件是否满足

## 应用领域

除了刚刚介绍的数学模型应用，有限状态机在许多不同领域都有重要应用，包括电气工程、语言学、计算机科学、哲学、生物学、数学和逻辑学。有限状态机归属于自动机理论，下面的自动机理论的领域分层图中就可以看出，越是外层的概念越复杂。

![img](https://upload-images.jianshu.io/upload_images/5709266-791b738cda6aa5f2.png?imageMogr2/auto-orient/strip|imageView2/2/w/440/format/webp)

## 有限状态机的举例

我们就拿身边最经典的电风扇来举例。假如电风扇有4个按钮，分别是关、1档、2档和3档，关按钮负责关闭电风扇，也就是停止电风扇的转动；而1、2、3档都可以让电风扇开启，且风扇转动的速度不一样，产生的风力也不一样。

这时我们判断出风扇的4个状态，分别是关闭(poweroff)、1档(1st gear)、2档(2nd gear)、3档(3rd gear)。而4个按钮的按下操作可以影响电风扇的状态。下面用状态图来说明：

![img](https://upload-images.jianshu.io/upload_images/5709266-b2eab09221fb9152.png?imageMogr2/auto-orient/strip|imageView2/2/w/569/format/webp)

如果看不清楚，还有状态转移表

![img](https://upload-images.jianshu.io/upload_images/5709266-d02202dfe551e897.png?imageMogr2/auto-orient/strip|imageView2/2/w/677/format/webp)

为了更直观的让程序员了解FSM具体有什么用，我将电风扇的有限状态机用程序来演示。

## Go语言下的有限状态机

一共2个文件，fsm.go是有限状态机的抽象定义，main.go里是有限状态机在电风扇上的具体状态呈现，代码如下：



```go
// fsm.go
package main

import (
    "fmt"
    "sync"
)

type FSMState string            // 状态
type FSMEvent string            // 事件
type FSMHandler func() FSMState // 处理方法，并返回新的状态

// 有限状态机
type FSM struct {
    mu       sync.Mutex                           // 排他锁
    state    FSMState                             // 当前状态
    handlers map[FSMState]map[FSMEvent]FSMHandler // 处理地图集，每一个状态都可以出发有限个事件，执行有限个处理
}

// 获取当前状态
func (f *FSM) getState() FSMState {
    return f.state
}

// 设置当前状态
func (f *FSM) setState(newState FSMState) {
    f.state = newState
}

// 某状态添加事件处理方法
func (f *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
    if _, ok := f.handlers[state]; !ok {
        f.handlers[state] = make(map[FSMEvent]FSMHandler)
    }
    if _, ok := f.handlers[state][event]; ok {
        fmt.Printf("[警告] 状态(%s)事件(%s)已定义过", state, event)
    }
    f.handlers[state][event] = handler
    return f
}

// 事件处理
func (f *FSM) Call(event FSMEvent) FSMState {
    f.mu.Lock()
    defer f.mu.Unlock()
    events := f.handlers[f.getState()]
    if events == nil {
        return f.getState()
    }
    if fn, ok := events[event]; ok {
        oldState := f.getState()
        f.setState(fn())
        newState := f.getState()
        fmt.Println("状态从 [", oldState, "] 变成 [", newState, "]")
    }
    return f.getState()
}

// 实例化FSM
func NewFSM(initState FSMState) *FSM {
    return &FSM{
        state:    initState,
        handlers: make(map[FSMState]map[FSMEvent]FSMHandler),
    }
}
```



```go
// main.go
package main

import (
    "fmt"
)

var (
    Poweroff        = FSMState("关闭")
    FirstGear       = FSMState("1档")
    SecondGear      = FSMState("2档")
    ThirdGear       = FSMState("3档")
    PowerOffEvent   = FSMEvent("按下关闭按钮")
    FirstGearEvent  = FSMEvent("按下1档按钮")
    SecondGearEvent = FSMEvent("按下2档按钮")
    ThirdGearEvent  = FSMEvent("按下3档按钮")
    PowerOffHandler = FSMHandler(func() FSMState {
        fmt.Println("电风扇已关闭")
        return Poweroff
    })
    FirstGearHandler = FSMHandler(func() FSMState {
        fmt.Println("电风扇开启1档，微风徐来！")
        return FirstGear
    })
    SecondGearHandler = FSMHandler(func() FSMState {
        fmt.Println("电风扇开启2档，凉飕飕！")
        return SecondGear
    })
    ThirdGearHandler = FSMHandler(func() FSMState {
        fmt.Println("电风扇开启3档，发型被吹乱了！")
        return ThirdGear
    })
)

// 电风扇
type ElectricFan struct {
    *FSM
}

// 实例化电风扇
func NewElectricFan(initState FSMState) *ElectricFan {
    return &ElectricFan{
        FSM: NewFSM(initState),
    }
}

// 入口函数
func main() {

    efan := NewElectricFan(Poweroff) // 初始状态是关闭的
    // 关闭状态
    efan.AddHandler(Poweroff, PowerOffEvent, PowerOffHandler)
    efan.AddHandler(Poweroff, FirstGearEvent, FirstGearHandler)
    efan.AddHandler(Poweroff, SecondGearEvent, SecondGearHandler)
    efan.AddHandler(Poweroff, ThirdGearEvent, ThirdGearHandler)
    // 1档状态
    efan.AddHandler(FirstGear, PowerOffEvent, PowerOffHandler)
    efan.AddHandler(FirstGear, FirstGearEvent, FirstGearHandler)
    efan.AddHandler(FirstGear, SecondGearEvent, SecondGearHandler)
    efan.AddHandler(FirstGear, ThirdGearEvent, ThirdGearHandler)
    // 2档状态
    efan.AddHandler(SecondGear, PowerOffEvent, PowerOffHandler)
    efan.AddHandler(SecondGear, FirstGearEvent, FirstGearHandler)
    efan.AddHandler(SecondGear, SecondGearEvent, SecondGearHandler)
    efan.AddHandler(SecondGear, ThirdGearEvent, ThirdGearHandler)
    // 3档状态
    efan.AddHandler(ThirdGear, PowerOffEvent, PowerOffHandler)
    efan.AddHandler(ThirdGear, FirstGearEvent, FirstGearHandler)
    efan.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
    efan.AddHandler(ThirdGear, ThirdGearEvent, ThirdGearHandler)

    // 开始测试状态变化
    efan.Call(ThirdGearEvent)  // 按下3档按钮
    efan.Call(FirstGearEvent)  // 按下1档按钮
    efan.Call(PowerOffEvent)   // 按下关闭按钮
    efan.Call(SecondGearEvent) // 按下2档按钮
    efan.Call(PowerOffEvent)   // 按下关闭按钮
}
```

执行后返回：



```bash
电风扇开启3档，发型被吹乱了！
状态从 [ 关闭 ] 变成 [ 3档 ]
电风扇开启1档，微风徐来！
状态从 [ 3档 ] 变成 [ 1档 ]
电风扇已关闭
状态从 [ 1档 ] 变成 [ 关闭 ]
电风扇开启2档，凉飕飕！
状态从 [ 关闭 ] 变成 [ 2档 ]
电风扇已关闭
状态从 [ 2档 ] 变成 [ 关闭 ]
```



27人点赞



[Go语言](https://www.jianshu.com/nb/14641699)