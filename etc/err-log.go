// Author : Eric Kim
// Build Date : 23 Jul 2018  Last Update 02 Aug 2018
// End-Agent for Passcon Multi OS go binding with Windows, MacOS, iOS, and Android
// All rights are reserved.

package etc

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func OkLog(s string) {
	log.Println("[OK]: " + s)
}

func AokLog(s string) {
	log.Println("[Abango-OK]: " + s)
}

func ErrLog(point string, err error) {
	log.Println("[ERROR]: " + point)
	if err != nil {
		fmt.Println("[err.Error()]:", err)
	}
}

func ChkLog(point string, x ...interface{}) {
	log.Println("[CHECK:" + point + "] " + fmt.Sprintf("%v", x))
}

// func FatalLog(point string, err error) {
// 	fmt.Println("[FATAL-ERROR]: "+point, err)
// 	os.Exit(1000)
// }

func MyErr(s string, e error, eout bool) error {
	fmt.Println("[MyErr] Position -> ", s, strings.Repeat("=", 40))

	emsg := ""
	if e != nil {
		emsg = "Error: " + e.Error()
	} else {
		emsg = "ERROR is Nil: Wrong Error Check: Check err != OR err == is correct !  "
	}
	fmt.Println(emsg, "\n")
	whereami(2)
	whereami(3)
	whereami(4)
	fmt.Println(strings.Repeat("=", 80))

	if e != nil && eout == true { // quit running if it is FATAL ERROR
		log.Println("[FATAL-ERROR] : EXIT 100")
		os.Exit(100)
	}
	return errors.New(emsg)
}

func Tp(a ...interface{}) {
	fmt.Println(a)
}

func Atp(a ...interface{}) {
	fmt.Println("[Abango]->", a)
}

func agErr(s string, e error, amsg *string) string {
	fmt.Println("== agErr ", strings.Repeat("=", 90))
	// fpcs := make([]uintptr, 1)
	// n := runtime.Callers(2, fpcs)
	// if n == 0 {
	// 	fmt.Println("MSG: NO CALLER")
	// }
	// // caller := runtime.FuncForPC(fpcs[0] - 1)
	// caller := runtime.FuncForPC(fpcs[0])
	// // fmt.Println(caller.FileLine(fpcs[0] - 1))
	// fmt.Println(caller.FileLine(fpcs[0]))
	// fmt.Println(caller.Name())
	emsg := ""
	if e != nil {
		emsg = "Error: " + e.Error() + " in " + s
	} else {
		emsg = "Error: error is nil" + " in " + s // e 가 nil 인 상태에서 Error() 인용시 runtime error
	}
	fmt.Println(emsg, "\n")
	whereami(2)
	whereami(3)
	fmt.Println(strings.Repeat("=", 100))
	return emsg
}

func whereami(i int) {
	function, file, line, _ := runtime.Caller(i)
	fmt.Printf("  %d.File: %s - %d  %s\n   func: %s \n", i, chopPath(file), line, file, runtime.FuncForPC(function).Name())
}

func WhereAmI(depthList ...int) {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	// function, file, line, _ := runtime.Caller(depth)

	for i := 0; i < depth+1; i++ {

		function, file, line, _ := runtime.Caller(i)
		fmt.Printf("==Level %d==\n", i)
		fmt.Printf("File: %s - %d  %s\nFunction: %s \n", chopPath(file), line, file, runtime.FuncForPC(function).Name())
	}
	fmt.Printf("==End==\n")

	return
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
