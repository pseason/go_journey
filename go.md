### Go Journey
<https://www.processon.com/view/link/5a9ba4c8e4b0a9d22eb3bdf0#map>

<http://interview.wzcu.com/>

<http://c.biancheng.net>

<https://mojotv.cn/404#Golang>

<http://www.topgoer.com/>
##### 1. go 安装、配置
```
1. 安装基础
    1. 开启gomodule：   (go env -w GO111MODULE=on) 大于 go 1.14 默认开启 
    2. 开启代理：       (go env -w GOPROXY=https://goproxy.cn,direct)
2. vscode settings.json 配置
{
    // 保存代码时自动编译
    "go.buildOnSave": "off",
    // 保存代码时优化
    "go.lintOnSave": "workspace",
    // 保存代码时检查潜在错误
    "go.vetOnSave": "workspace",
    // 保存代码时执行测试
    "go.coverOnSave": false,
    "go.useCodeSnippetsOnFunctionSuggest": true,
    "go.useCodeSnippetsOnFunctionSuggestWithoutType": true,
    // 代码格式化
    "go.formatTool": "gofmt",
    "go.gocodePackageLookupMode": "go",
    "go.gotoSymbol.includeImports": true,
    "go.docsTool": "godoc",
    "go.inferGopath": false,
    "go.useLanguageServer": true,
    "go.autocompleteUnimportedPackages": true,
}

3. 安装go 工具
    command+shift+P，然后键入：go:install/update tools，将所有 16 个插件都勾选上

```
##### 2. go 简介
```
1. go语言简介
    1. Go语言也称为 Golang，是由 Google 公司开发的一种静态强类型、编译型、并发型、
        并具有垃圾回收功能的编程语言
    2. Go 使用编译器来编译代码。编译器将源代码编译成二进制（或字节码）格式；在编译代码时，
        编译器检查错误、优化性能并输出可在不同平台上运行的二进制文件.
    3. Go语言支持交叉编译，比如说你可以在运行 Linux 系统的计算机上开发可以在Windows
        上运行的应用程序
    4.Go语言快速编译，高效执行，易于开发，可以进行网络编程、系统编程、并发编程、分布式编程

2. Go语言的特性有哪些
    1. 语法简单
    2. 并发模型
    3. 内存分配
    4. 垃圾回收
    5. 静态链接
    6. 标准库
    7. 工具链
```
##### 3. go 基础语法
```
1. go基础类型
    Go语言的基本类型有：
        bool
        string
        int、int8、int16、int32、int64
        uint、uint8、uint16、uint32、uint64、uintptr
        byte // uint8 的别名
        rune // int32 的别名 代表一个 Unicode 码
        float32、float64 // 浮点类型
        complex64、complex128 // 复数

2.声明变量
    标准声明： var 变量名 变量类型
    批量声明： var {
        a int
        b string
    }
    简短声明： 变量名 := 表达式

3.变量初始化
    标准格式：      var 变量名 变量类型 = 表达式 
    编译器推导：    var 变量名 = 表达式
    声明并初始化:   变量名 := 表达式

4. 匿名变量
    在编码过程中，可能会遇到没有名称的变量、类型或方法。
    虽然这不是必须的，但有时候这样做可以极大地增强代码的灵活性，这些变量被统称为匿名变量

    a,_ = func(){
        return 100,200   
    }

5. 变量作用域
    根据变量定义位置的不同，可以分为以下三个类型：
        函数内定义的变量称为局部变量
        函数外定义的变量称为全局变量
        函数定义中的变量称为形式参数

6. 字符串及其操作
    字符串定义
    字符串拼接 “+”
    定义多行字符串 ``
    字符串长度 
        len()、utf8.RuneCountInString
    字符串遍历
        ASCII 使用for循环
        Unicode 使用range
    字符串截取
        字符串索引比较常用的有如下几种方法：
        strings.Index：正向搜索子字符串。
        strings.LastIndex：反向搜索子字符串。
        搜索的起始位置可以通过切片偏移制作。
    字符串修改
        修改字符串时，可以将字符串转换为 []byte 进行修改。
        []byte 和 string 可以通过强制类型转换互转
    字符串拼接
        使用类java StringBuffer, bytes.Buffer
    base64 编解码

7. 字符格式化输出
    %v	按值的本来值输出
    %+v	在 %v 基础上，对结构体字段名和值进行展开
    %#v	输出 Go 语言语法格式的值
    %T	输出 Go 语言语法格式的类型和值
    %%	输出 % 本体
    %b	整型以二进制方式显示
    %o	整型以八进制方式显示
    %d	整型以十进制方式显示
    %x	整型以十六进制方式显示
    %X	整型以十六进制、字母大写方式显示
    %U	Unicode 字符
    %f	浮点数
    %p	指针，十六进制方式显示

8. 指针类型
    Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作，
    当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指使用*操作符，也就是指针取值
    数据取内存地址：ptr x := &y
    内存地址转取值：z := *x

9. 变量逃逸分析
    1.栈（Stack）是一种拥有特殊规则的线性表数据结构
    2.栈只允许从线性表的同一端放入和取出数据，按照后进先出（LIFO，Last InFirst Out）的顺序
    3.栈可用于内存分配，栈的分配和回收速度非常快
    4.堆分配内存和栈分配内存相比，堆适合不可预知大小的内存分配。
        但是为此付出的代价是分配速度较慢，而且会形成内存碎片
    5.Go语言将这个过程整合到了编译器中，命名为“变量逃逸分析”。通过编译器分析代码的特征和代码的生命周期，
        决定应该使用堆还是栈来进行内存分配
    6.编译器觉得变量应该分配在堆和栈上的原则是：
        变量是否被取地址；
        变量是否发生逃逸。
    7.在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法——分析 在程序的哪些地方可以访问到指针。
        也是就是说逃逸分析是解决指针作用范围的编译优化方法。编程中常见的两种逃逸情景：
        1，函数中局部对象指针被返回（不确定被谁访问）
        2，对象指针被多个子程序（如线程 协程）共享使用
    8.逃逸分析的场景
        指针逃逸(函数返回局部变量的指针)
        栈空间不足逃逸
        闭包引用逃逸
        动态类型逃逸
        切片或map赋值
    go build -gcflags "-m -l" file.go 

10. 变量生命周期
    a.变量的生命周期指的是在程序运行期间变量有效存在的时间间隔
    b.变量的生命周期与变量的作用域有着不可分割的联系：
        1.全局变量：它的生命周期和整个程序的运行周期是一致的
        2.局部变量：它的生命周期则是动态的，从创建这个变量的声明语句开始，
            到这个变量不再被引用为止
        3.形式参数和函数返回值：它们都属于局部变量，在函数被调用的时候创建，
            函数调用结束后被销毁
    c.栈和堆的区别在于：
        1.堆（heap）：堆是用于存放进程执行中被动态分配的内存段。它的大小并不固定，可动态扩张或缩减。
            当进程调用 malloc 等函数分配内存时，新分配的内存就被动态加入到堆上（堆被扩张）
            。当利用 free 等函数释放内存时，被释放的内存从堆中被剔除（堆被缩减）
        2.栈(stack)：栈又称堆栈， 用来存放程序暂时创建的局部变量，也就是我们函数的大括号{ }中定义的局部变量

11. go语言make和new关键字的区别及实现原理
    a. 其实他们的规则很简单，new 只分配内存，
        而 make 只能用于 slice、map 和 channel 的初始化
    b. make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，
        而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，
        所以就没有必要返回他们的指针
    c. Go语言中的 new 和 make 主要区别如下：
        make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
        new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
        new 分配的空间被清零。make 分配空间后，会进行初始化
    d. make原理
        在编译期的类型检查阶段，Go语言其实就将代表 make 关键字的 OMAKE 节点根据参数类型的不同
        转换成了 OMAKESLICE、OMAKEMAP 和 OMAKECHAN 三种不同类型的节点，
        这些节点最终也会调用不同的运行时函数来初始化数据结构
    e. new 原理
        内置函数 new 会在编译期的 SSA 代码生成阶段经过 callnew 函数的处理，
         如果请求创建的类型大小是 0，那么就会返回一个表示空指针的 zerobase 变量，在遇到其他情况时会将关键字转换成 newobject
    f. 总结
        make 关键字的主要作用是创建 slice、map 和 Channel 等内置的数据结构，
        而 new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针
    g. SSA概念
        SSA即静态单一赋值，意思就是某个计算流程中，某些元素可能会出现被反复赋值的问题，
        那么之前这个元素存在的值就消失了，这种情况就不便于回溯，而且容易在多线程的情况下产生错误。
        因此如果采用静态单一赋值，每一个过程量都是一个新的变量，每个变量都只有一个定值，这种方法就是SSA静态单一赋值方法

12. go 语言常量和const关键字
    1. 定义常量的表达式必须为能被编译器求值的表达式，并且常量只能是布尔型、
        数字型（整数型、浮点型和复数）和字符串型。
    2. 常量的声明语法为： 
        const name [type] = value\express
        const name = value\express
        const (
            name [type] = value\express
            name  = value\express
        )
        [type] 可不书写，可由编译器根据值推断类型
    3. iota 常量生成器
        常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，
        但是不用每行都写一遍初始化表达式。在一个 const 声明语句中，
        在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加1
    4. 无类型常量
        编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，
        可以认为至少有 256bit 的运算精度。
        这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、
        无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。

13. Go语言模拟枚举（const和iota模拟枚举）
    1. Go语言现阶段没有枚举类型
    2. 可以使用 const 常量配合 iota 来模拟枚举类型
    3. 枚举转字符，可使用定义String()来实现

14. Go语言type关键字（类型别名）
    1. 类型定义：
        使用方法 type name 基础类型 -> type Type int
    2. 类型别名：
        使用方法 type name = 基础类型 -> type Alias = int

15. Go语言strconv包：字符串和数值类型的相互转换
    1. Itoa()：整型转字符串
    2. Atoi()：字符串转整型
    3. Parse 系列函数
        ParseBool()
        ParseInt()
        ParseUnit()
        ParseFloat()
    4. Format 系列函数
        FormatBool()
        FormatInt()
        FormatUint()
        FormatFloat()
    5. Append 系列函数
        AppendBool()
        AppendFloat()
        AppendInt()
        AppendUint()
16. Go语言注释及文档工具
    1. 单行 //， 多行 /*    */
    2. go doc 工具 godoc 
        go doc  // 查看项目注释
        godoc -http=:6060   // 使用工具构建web服务查看

17. Go语言中的关键字与标识符
    Go语言的词法元素包括 5 种，分别是标识符（identifier）、关键字（keyword）、
    操作符（operator）、分隔符（delimiter）、字面量（literal），
    它们是组成Go语言代码和程序的最基本单位
    1. 关键字即是被Go语言赋予了特殊含义的单词，也可以称为保留字,关键字一共有 25 个：
        break	default 	func	interface	select
        case	defer	go	map	struct
        chan	else	goto	package	switch
        const	fallthrough	if	range	type
        continue	for	import	return	var
    2. 标识符是指Go语言对各种变量、方法、函数等命名时使用的字符序列，
        1.标识符由若干个字母、下划线_、和数字组成，
            且第一个字符必须是字母。通俗的讲就是凡可以自己定义的名称都可以叫做标识符
        2.标识符的命名需要遵守以下规则：
            由 26 个英文字母、0~9、_组成；
                a.不能以数字开头，例如 var 1num int 是错误的；
                b.Go语言中严格区分大小写；
                c.标识符不能包含空格；
                d.不能以系统保留关键字作为标识符，比如 break，if 等等。
            命名标识符时还需要注意以下几点：
                a.标识符的命名要尽量采取简短且有意义；
                b.不能和标准库中的包名重复；
                c.为变量、函数、常量命名时采用驼峰命名法，例如 stuName、getVal；
        3.当然Go语言中的变量、函数、常量名称的首字母也可以大写，如果首字母大写，
            则表示它可以被其它的包访问（类似于 Java 中的 public）；
            如果首字母小写，则表示它只能在本包中使用 (类似于 Java 中 private）

18. Go语言运算符的优先级
    所谓优先级，就是当多个运算符出现在同一个表达式中时，先执行哪个运算符
    Go语言有几十种运算符，被分成十几个级别，有的运算符优先级不同，有的
    运算符优先级相同，请看下表。
        Go语言运算符优先级和结合性一览表
            优先级	分类				结合性		运算符	
            1		逗号运算符		从左到右		,
            2		赋值运算符		从右到左		=、+=、-=、*=、/=、 %=、 >=、 <<=、&=、^=、|=	
            3		逻辑或			从左到右		||
            4		逻辑与			从左到右		&&
            5		按位或			从左到右		|
            6		按位异或		从左到右		^
            7		按位与			从左到右		&
            8		相等/不等		从左到右		==、!=
            9		关系运算符		从左到右		<、<=、>、>=
            10		位移运算符		从左到右		<<、>>
            11		加法/减法		从左到右		+、-	
            12		乘法/除法/取余	从左到右		*（乘号）、/、%
            13		单目运算符		从右到左		!、*（指针）、& 、++、--、+（正号）、-（负号）	
            14		后缀运算符		从左到右		( )、[ ]、->
        注意：优先级值越大，表示优先级越高
                括号的优先级是最高的，括号中的表达式会优先执行
    
```
##### 4. go 语言容器
```
1. Go语言 数组
    a. 数组的声明语法如下：
        var 数组变量名 [元素数量]Type
        语法说明如下所示：
            - 数组变量名：数组声明及使用时的变量名。
            - 元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，
            元素数量不能含有到运行时才能确认大小的数值。
            - Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组
    b. 多维数组
        声明多维数组的语法如下所示：
            var array_name [size1][size2]...[sizen] array_type
        其中，array_name 为数组的名字，array_type 为数组的类型，size1、size2 等等为数组每一维度的长度
    c. 数据类型转换之后结果为
        %T -> [5]int...

2. Go语言 切片（slice）
    a. 切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型,这个片段可以是整个数组，
      也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内.
    b. Go语言中切片的内部结构包含地址、大小和容量,切片一般用于快速地操作一块数据集合,如果将数据集合比作切糕的话,
      切片就是你要的“那一块”，切的过程包含从哪里开始（切片的起始位置）及切多大（切片的大小），
      容量可以理解为装切片的口袋大小
    c. 从数组生成切片
        slice [开始位置 : 结束位置]
        语法说明如下：
            slice：表示目标切片对象；
            开始位置：对应目标切片对象的索引；
            结束位置：对应目标切片的结束索引
    d. 从数组或切片生成新的切片拥有如下特性
        取出的元素数量为：结束位置 - 开始位置；
        取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
        当缺省开始位置时，表示从连续区域开头到结束位置；
        当缺省结束位置时，表示从开始位置到整个连续区域末尾；
        两者同时缺省时，与切片本身等效；
        两者同时为 0 时，等效于空切片，一般用于切片复位。
    e. 直接声明新切片
        var name []type -> var ints []int
        其中 name 表示切片的变量名，Type 表示切片对应的元素类型
    f. 使用 make() 函数构造切片
        如果需要动态地创建一个切片，可以使用 make() 内建函数
            make( []Type, size, cap )
        其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，
        这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题
    g. 使用append()方法为切片添加元素
        slice = append(slice,value)
        在使用 append() 函数为切片动态添加元素时，
        如果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变
        
        切片默认cap为6，容量的扩展规律是按,长度少于1024时,
        容量的2倍数进行扩充,超过1024的，1.25倍扩容
    h. 切片拷贝复制(copy)
        copy() 函数的使用格式如下
            copy( destSlice, srcSlice )
        其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice），
        目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，
        copy() 函数的返回值表示实际发生复制的元素个数
    i. 切片元素删除
        1.从开头位置删除
            利用切片特性              a = a[N:]
            利用append()函数          append(a[:0], a[N:]...)  
            利用copy()函数            a[:copy(a,a[N:])]
        2.从中间位置删除
            利用append()函数          append(a[:i],a[i+n:]...)
            利用copy()函数            a[:i+copy(a[:i],a[i+N:])]
        3.从尾部删除
            利用切片特性              a = a[:len(a) - N]

3. Go语言 map（Go语言映射）
    a. 声明map
        1.var Name map[keyType]valueType
            其中
                name 为声明变量名称
                keyType为map，key声明类型
                keyValue为map,value声明类型
        2.使用make函数创建
            map := make(map[keyType]valueType)
            map := make(map[keyType]valueType,cap) cap作为扩容参数
    b. map 删除元素、清空
        删除元素：delete(map,key)
        清空数据：使用make()创建新map
    c. map多键索引
        1. 基于哈希值的多键索引及查询
            传统的数据索引过程是将输入的数据做特征值。这种特征值有几种常见做法：
                - 将特征使用某种算法转为整数，即哈希值，使用整型值做索引。
                将特征转为字符串，使用字符串做索引。
                - 数据都基于特征值构建好索引后，就可以进行查询。查询时，重复这个过程，
                将查询条件转为特征值，使用特征值进行查询得到结果
        2. 利用 map 特性的多键索引及查询
            使用结构体进行多键索引和查询比传统的写法更为简单，
            最主要的区别是无须准备哈希函数及相应的字段无须做哈希合并
        总结：
            基于哈希值的多键索引查询和利用 map 特性的多键索引查询的代码复杂程度显而易见。
            聪明的程序员都会利用 Go语言的特性进行快速的多键索引查询。
            代码量大大减少的关键是：
                Go语言的底层会为 map 的键自动构建哈希值。能够构建哈希值的类型必须是非动态类型、非指针、函数、闭包。
                非动态类型：可用数组，不能用切片。
                非指针：每个指针数值都不同，失去哈希意义。
                函数、闭包不能作为 map 的键。
    d. map 在并发情况下，只读是线程安全的，同时读写是线程不安全的

4. Go语言sync.Map
    a. 声明sync.Map变量
        var xxx sync.Map

    sync.Map 有以下特性：
        - 无须初始化，直接声明即可。
        - sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，
            Store 表示存储，Load 表示获取，Delete 表示删除。
        - 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，
            Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false

5. Go语言list（列表）
    a. 特性
        - 在Go语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，
            列表能够高效地进行任意位置的元素插入和删除操作
        - 列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，
            这既带来了便利，也引来一些问题，例如给列表中放入了一个 interface{} 类型的值，
            取出值后，如果要将 interface{} 转换为其他类型将会发生宕机
    b. 初始化列表
        list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的
            name := list.New()
            var name list.List
    c. 列表操作
        list.PushFront()        向列表首部添加
        list.pushBack()         向列表尾端添加
        InsertAfter()       	在 mark 点之后插入元素，mark 点由其他插入函数提供
        InsertBefore()          在 mark 点之前插入元素，mark 点由其他插入函数提供
        PushBackList()      	添加 other 列表元素到尾部
        PushFrontList()         添加 other 列表元素到头部
```
##### 4. go 流程控制
```
1. Go语言if else（分支结构）
    a. 基本： if 条件 {} else if 条件 {} else{}
    b. 特殊： if 执行语句; 条件 {} else if 条件 {} else{}

2. Go语言for（循环结构）
    a. 使用循环语句时，需要注意的有以下几点：
        - 左花括号{必须与 for 处于同一行。
        - Go语言中的 for 循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别是，
            Go语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量。
        - Go语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break

3. Go语言for range（键值循环）
    for key, val := range coll {
        ...
    }
    a. 通过 for range 遍历的返回值有一定的规律：
        - 数组、切片、字符串返回索引和值。
        - map 返回键和值。
        - 通道（channel）只返回通道内的值。
    b. val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，
       对它所做的任何修改都不会影响到集合中原有的值

4. Go语言函数（Go语言func）
    a. defer panic recover 关键字
    b. 闭包
    c. 链式方法
    d. 函数实现接口
    e. 匿名函数

5. Go语言结构体（struct）
    a. 为任意类型添加方法
        type MyName Type:
            type MyInt int
            func (i MyInt) IsZero() bool{
                return i == 0
            }

    b. 初始化结构体
        t := new(T)
        t := &T{}
    
    c. 使用父子关系的结构体实现继承

6. Go语言接口（interface）
    a. Go语言实现接口的条件
        - 接口的方法与实现接口的类型方法格式一致
        - 接口中所有方法均被实现
    b. 类与接口的关系
        - 一个类可以实现多个接口
        - 多个类型可以实现相同接口
    c. 接口类型转换
        var ii interface{} = 1
        num := ii.(int)
    d. 内嵌接口

7. Go语言包（package）
    a.Go语言包的初始化有如下特点：
        - 包初始化程序从 main 函数引用的包开始，逐级查找包的引用，
          直到找到没有引用其他包的包，最终生成一个包引用的有向无环图。
        - Go 编译器会将有向无环图转换为一棵树，然后从树的叶子节点开始逐层向上对包进行初始化。
        - 单个包的初始化过程如上图所示，先初始化常量，然后是全局变量，最后执行包的 init 函数。
    
8. Go 游戏服务器框架
    <https://blog.csdn.net/weixin_34256074/article/details/86029237>

    <https://link.zhihu.com/?target=https%3A//github.com/duanhf2012/origin>

    <https://blog.csdn.net/Nankys/article/details/111556935>

    <https://github.com/topfreegames/pitaya>

    pomelo

    pitaya go-nano goworld leaf cellnet

    自己在学习的同时对源码随后做了些注释。后续会对整个框架做详细的结构性的分析。注释版本的github仓库如下。
    代码会及时更新合并官方的master. 仓库如下：
    https://github.com/bytemode/pitaya-notes
    同时见了qq群用于技术交流，正在使用go的朋友欢迎加入交流。
    QQ群：621275137

9. Go语言的编译与工具
    a. 相关命令
        go help command(mod build get install...)
        go get              拉取网络依赖包 
        go mod              go module 相关
        go build 工程构建
        go build -o myapp.exe
        go build -a 
        go build -gcflags '-m -l' main.go   构建分析变量逃逸，堆栈信息
        go clean 清除构建内容
        go fmt 格式化代码文件
        go install 编译并安装
        go pprof 性能分析命令

```