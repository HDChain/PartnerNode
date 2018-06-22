package programfind

import (
	"flag"
	"fmt"
	//"glog-master"
	"net/http"
	//"os"
	"os"
	"os/exec"
	"syscall"
)

//2.执行进程

func init() {
	// 初始化 日志系统
	flag.Set("alsologtostderr", "true") // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", "./log")        // 日志文件保存目录
	flag.Set("v", "3")                  // 配置V输出的等级。
	flag.Parse()
	return
}

// 主函数
func Startprogram() {
	//fmt.Print(os.Args[1:])
	//glog.Info("Entry main!!!")
	http.HandleFunc("/DuliServer", DuLiServer) // 独立验证
	err := http.ListenAndServe(":8365", nil)
	if err != nil {
		fmt.Println("Entry nil", err.Error())
		return
	}
	return
}

//func main() {
// fmt.Println("启动游戏成功")
// glog.Info("Entry main!!!")
// fmt.Print(os.Args[1:])
// fmt.Printf(flag.Arg(1))
// glog.Info("....", flag.Arg(1))
// glog.Flush()
// // 正式的
// http.ListenAndServe(":8364", nil)
// return
//}

// 独立服务器

func DuLiServer(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		// 获取函数
		strLoginName, bLoginName := req.Form["LoginName"]
		strLoginPW, bLoginPW := req.Form["LoginPW"]
		strGameName, bGameName := req.Form["GameName"]
		strXCName, bXCName := req.Form["XCName"]
		if bLoginPW && bLoginName && bGameName && bXCName {
			fmt.Println("strLoginName:", strLoginName)
			fmt.Println("strLoginPW", strLoginPW)
			fmt.Println("strGameName:", strGameName)
			fmt.Println("strXCName", strXCName)
			fmt.Fprint(w, "发送成功！！！")
			// 发送给游戏
			CallEXE("F:gocodecodeopenCVbinGame.exe")
			return
		}
		fmt.Fprint(w, "启动失败，参数不对！！！")
		return
	}
}

func CallExeEx() {
	//在我们的例子中，我们将执行 ls 命令。Go 需要提供我们需要执行的可执行文件的绝对路径，所以我们将使用exec.LookPath 来得到它（大概是 /bin/ls）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	//Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}
	//Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量。
	env := os.Environ()
	//这里是 os.Exec 调用。如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// 启动exe
func CallEXE(strGameName string) {
	fmt.Println("开始启动游戏。。。")
	arg := []string{"参数一", "参数二"}
	fmt.Println("------------", arg)
	// cmd := exec.Command("F:最新版本游戏test1testtest.exe", arg...)
	cmd := exec.Command("Game.exe", arg...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}
