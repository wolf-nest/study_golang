## CSP vs Actor
* 和Actor的直接通讯不同，CSP模式则是通过Channel进行通讯的，更松耦合一些
* Go中channel是有容量限制并且独立于处理Groutine而如 Erlang，Actor 模式中的mailbox 容量是无限的，接收进程也总是被动的处理消息。

## select
* 多渠道的选择
    ```$go 
    select {
        case ret := <- retCh1:
            t.Logf("result %s",ret)
        case ret := <- retCh2:
            t.Logf("result %s",ret)
        default:
            t.Error("No one returned")
    }
    ```
* 超时控制
    ```$go
    select {
        case ret := <-retCh:
            t.Logf("result %s",ret)
        case <-time.After(time.Second * 1):
            t.Error("time out")
    }

