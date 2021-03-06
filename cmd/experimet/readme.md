## Tools
	$ go tool compile stdIN.go 
    -> stdIN.o

    $ go tool compile -pack stdIN.go 
    -> stdIN.a - ar архив
    $ ar t stdIN.a   - просмотреть архив

    $ go tool compile -S stdIN.go 
    -> вывод всего процесса
        os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0
        0x0000 00000 (<autogenerated>:1)        TEXT    os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0000 00000 (<autogenerated>:1)        MOVQ    ""..this+8(SP), AX
        0x0005 00005 (<autogenerated>:1)        MOVQ    (AX), AX
        0x0008 00008 (<autogenerated>:1)        MOVQ    AX, ""..this+8(SP)
        0x000d 00013 (<autogenerated>:1)        XORPS   X0, X0
        0x0010 00016 (<autogenerated>:1)        MOVUPS  X0, "".~r0+16(SP)
        0x0015 00021 (<autogenerated>:1)        JMP     os.(*file).close(SB)
        0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
        0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
        rel 22+4 t=8 os.(*file).close+0

    $ GODEBUG=gctrace=1 go run garbageCollection.go
    -> показ размера кучи
        gc 1 @0.099s 0%: 0+0.99+0 ms clock, 0+0/0/0.99+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
        gc 2 @0.177s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
        gc 3 @0.262s 0%: 0+2.0+0 ms clock, 0+0.99/0.99/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
        gc 4 @0.343s 0%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
        gc 5 @0.357s 0%: 0+1.0+0 ms clock, 0+0/1.0/2.0+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
        gc 6 @0.376s 0%: 0+1.0+0 ms clock, 0+1.0/1.0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
        gc 7 @0.390s 0%: 0+1.0+0 ms clock, 0+0/1.0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 4 P

#### трехцветный алгоритм сборщика мусора

#### отладка
    $ go build panicRecover.go
    $ strace panicRecover.exe 
        --- Process 15612 created
        --- Process 15612 loaded C:\Windows\System32\ntdll.dll at 00007fffc6a50000
        --- Process 15612 loaded C:\Windows\System32\kernel32.dll at 00007fffc5e80000
        --- Process 15612 loaded C:\Windows\System32\KernelBase.dll at 00007fffc4190000

    For Linux
    $ dtrace ...
    $ dtruss ..

#### ENV
    $ go env
        set GOARCH=amd64
        set GOCACHE=C:\Users\Yuri\AppData\Local\go-build
        set GOENV=C:\Users\Yuri\AppData\Roaming\go\env
        set GOEXE=.exe
        set GOHOSTARCH=amd64
        set GOHOSTOS=windows
        set GOMODCACHE=C:\Users\Yuri\go\pkg\mod
        set GOOS=windows
        set GOPATH=C:\Users\Yuri\go
        set GOPROXY=https://proxy.golang.org,direct
        set GOROOT=c:\go
        set GOSUMDB=sum.golang.org
        set GOTOOLDIR=c:\go\pkg\tool\windows_amd64
        set GCCGO=gccgo
        set AR=ar
        set CC=gcc
        set CXX=g++
        set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\Yuri\AppData\Local\Temp\go-build747506618=/tmp/go-build -gno-record-gcc-switches

#### ASM
    $ GOOS=darwin GOARCH=amd64 go tool compile -S goEnv.go

    разбор синтаксического дерева
    $ go tool compile -W goEnv.go

    WASM
    $ go build -o main.wasm goEnv.go
    получаем 2Мб файлик - отлично

    OS	    $GOOS
    Linux	linux
    MacOS X	darwin
    Windows	windows
    Android	android