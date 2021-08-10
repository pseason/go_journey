1. 安装 protoc

    windows、linux、mac 下载对应版本，放入 GoPath bin目录
    
    
2. 安装protoc-gen-go
    a. Run the following command to install the Go protocol buffers plugin:

        go install google.golang.org/protobuf/cmd/protoc-gen-go

    b. The compiler plugin protoc-gen-go will be installed in $GOBIN, defaulting to $GOPATH/bin. 
        It must be in your $PATH for the protocol compiler protoc to find it.
        Now run the compiler, specifying the source directory 
        (where your application's source code lives – the current directory is used if you don't provide a value),
        the destination directory (where you want the generated code to go; often the same as $SRC_DIR), and the path to your .proto. 
     

        protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/file.proto

    c. Because you want Go code, you use the --go_out option – similar options are provided for other supported languages.