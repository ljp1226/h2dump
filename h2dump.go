package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"regexp"
	"flag"
	"os"
)

var host *string
var port *string
var user *string
var passwd *string
var dbName *string
var fileName *string
var useDbStr = "use "

func main() {
	host = flag.String("h", "127.0.0.1", "mysql host")
	port = flag.String("P", "3306", "mysql port")
	user = flag.String("u", "root", "mysql user")
	passwd = flag.String("p", "root", "mysql password")
	dbName = flag.String("d", "mysql", "mysql database")
	fileName = flag.String("f", "db.sql", "script dir")
	flag.Parse()

	useDbStr = useDbStr + *dbName + ";"

	handleDumpTask()
}

func handleDumpTask() {
	cmd := execMysqlCmd(useDbStr + "show tables")
	//run cmd
	outstr := runCMd(cmd)
	tableNames := strings.Split(outstr, "\n")
	if len(tableNames) <= 1 {
		fmt.Println("export fail!")
		return
	}

	newStrs := ""
	for idx, tableName := range tableNames {
		if idx != 0 {
			cmd := execMysqlCmd(useDbStr + "show create table " + tableName)
			outstr := runCMd(cmd)
			newline := strings.Split(outstr, "\n")
			for idx, v := range newline {
				if idx != 0 {
					newLine2 := strings.Split(v, "\t")
					if len(newLine2) > 1 {
						newStrs += regReplace(newLine2[1])
					}
				}
			}

		}
	}
	writeToFile2(newStrs)
	sedCmd := exec.Command("sed", "", "-i", "1d", *fileName)
	runCMd(sedCmd)
	fmt.Println("export complete!")
}

func writeToFile(tbarray []byte) {
	err2 := ioutil.WriteFile(*fileName, tbarray, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func writeToFile2(result string) {
	// Make test file
	testFile, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	// Remove the redirect from command
	cmd := exec.Command("echo", result)
	// Redirect the output here (this is the key part)
	cmd.Stdout = testFile

	if err = cmd.Start(); err != nil {
		panic(err)
	}
	cmd.Wait()
}

func runCMd(cmd *exec.Cmd) string {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	//读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	return string(opBytes)
}

func regReplace(newstr string) string {
	newstr = "mysql dump to h2 script\n" + newstr
	newstr = strings.Replace(newstr, "\\n", "\n", -1)
	newstr = strings.Replace(newstr, "`", "", -1)
	newstr = strings.Replace(newstr, "DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP", "AS CURRENT_TIMESTAMP", -1)
	newstr = strings.Replace(newstr, "DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP", "AS CURRENT_TIMESTAMP", -1)
	r := regexp.MustCompile(" COMMENT '.*'")
	newstr = r.ReplaceAllString(newstr, "")
	r = regexp.MustCompile(" ENGINE=.*")
	newstr = r.ReplaceAllString(newstr, ";\n")
	return newstr
}

func execMysqlCmd(cmd string) *exec.Cmd {
	return exec.Command("mysql", "-h" + *host, "-u" + *user, "-p" + *passwd, "-P" + *port, "-e", cmd)
}
