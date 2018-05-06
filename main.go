package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
	"bufio"
	"strings"
	"strconv"
)

const (
	COMMAND      = "COMMAND"
	DAYS         = "DAYS"
	DAYOFTHEWEEK = "DAYOFTHEWEEK"
	MONTH        = "MONTH"
	HOURS        = "HOURS"
	MINUTES      = "MINUTES"
)

func main() {
	app := cli.NewApp()
	AppInfo(app)
	app.Action = func(context *cli.Context) error {
		if context.Bool("new") {
			GenerateCronCommand()
		} else {
			fmt.Println(context.Args().Get(0))
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "new, n",
			Usage: "generate new cron command",
		},
	}
	app.Run(os.Args)
}
func GenerateCronCommand() {
	question := [6]string{MINUTES, HOURS, DAYS, MONTH, DAYOFTHEWEEK, COMMAND}
	//[TODO]yamlに切り出し
	questionMessage := map[string]string{
		MINUTES:      "minutes?(e.g. 0 to 59)",
		HOURS:        "hours?(e.g. 0 to 23)",
		DAYS:         "days?(e.g. 1 to 31)",
		MONTH:        "month?(e.g. 1 to 12)",
		DAYOFTHEWEEK: "Which day of the week?(e.g sun, mon, etc.. or 0, 1, 2, etc...)",
		COMMAND:      "What's command?"}
	fmt.Println(
		"multiple specification(e.g. 1,5 )\n" +
			"interval specification(e.g. */5 )\n" +
			"range specification(e.g. 1-5 )")
	res := []string{}
	for _, question := range question {
		fmt.Println(questionMessage[question]);
		fmt.Print(">")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		res = append(res, ConversionCronCommand(question, stdin.Text()))
	}
	fmt.Println(strings.Join(res, " "))
}

func ConversionCronCommand(question string, answer string) string {
	if (answer == "") {
		return "*"
	}
	if question == DAYOFTHEWEEK {
		if iStr, ok := arrayContainsIndex([]string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}, answer); ok {
			return iStr
		}

	}
	return answer
}

func AppInfo(app *cli.App) {
	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"
}

func arrayContainsIndex(arr []string, str string) (string, bool) {
	for k, v := range arr {
		if v == str {
			return strconv.Itoa(k), true
		}
	}
	return "0", false
}
