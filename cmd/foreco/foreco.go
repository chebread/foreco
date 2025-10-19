package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/chebread/foreco/internal/db"
	"github.com/chebread/foreco/internal/lib"
)

// foreco는 음식 추천기임.
// map 공부하기
// kong pkg 공부하기 -> 일단 os.Args로 구현함

// subcommand: brew install 이나 git commit 처럼 메인 명령어 뒤에 따라오는 추가적인 명령어는 서브커맨드(Subcommand)라고 부름.
// 프로그램이 제공하는 여러 기능 중 어떤 기능을 실행할지 결정함. 마치 하나의 프로그램 안에 여러 개의 작은 프로그램이 있는 것과 같음.

// flag, option: 하이픈으로 시작하는 인자는 플래그(Flag) 또는 옵션(Option)이라고 부름.
// 명령어의 동작 방식을 바꾸거나 추가 정보를 제공하는 역할을 함.

var ProgramName = os.Args[0]
var ProgramVersion string = "development"

var foods = map[string][]string{
	"korean": db.KoreanFoods,
	"vegan":  db.VeganFoods,
}
var categories = map[string]string{
	"korean": "한식",
	"vegan":  "비건",
}

func main() {
	// 인수가 제공되지 않음 (모든 음식 렌덤)
	// 방어 코드
	if len(os.Args) < 2 {
		// os.Args is []string
		var allFoods []string                              // nil slice도 append가 가능함.
		var koreanFoodsToAppend []string = foods["korean"] // 1...100
		var veganFoodsToAppend []string = foods["vegan"]   // 1...50

		allFoods = append(allFoods, koreanFoodsToAppend...) // oods["korean"]...는 불가능함.
		allFoods = append(allFoods, veganFoodsToAppend...)

		var randOffset int = rand.Intn(len(allFoods) - 1) // 1...150
		var food string = allFoods[randOffset]

		// 음식의 카테고리 확인
		var category string
		if len(koreanFoodsToAppend) >= randOffset+1 {
			category = "한식"
		} else {
			category = "비건"
		}

		lib.BoldCyanPrintf("%s: %s\n", category, food)

		return
	}

	// 방어 코드를 사용했기 때문에 여기서 부터는 len(os.Args) >= 2임.

	// flag
	// 방어 코드
	flagIdx := strings.Index(os.Args[1], "-")
	if flagIdx == 0 {
		flagLastIdx := strings.LastIndex(os.Args[1], "-")
		flagVal := os.Args[1][flagLastIdx+1:] // string slice는 [m:n] 이면 m부터 n-1까지 자른 다는 거임. 즉, m이 포함됨.

		switch flagVal {
		case "help", "h":
			help()
		case "version", "V":
			version()
		default:
			lib.RedCyanPrintf("Error: '%s' you entered is not supported\n", flagVal)
		}

		return
	}

	var inputedFood string = os.Args[1]

	// 해당하는 key가 없음.
	var _, ok = foods[inputedFood]
	// 방어 코드
	if !ok {
		lib.RedCyanPrintf("Error: '%s' you entered is not in the category\n", inputedFood)
		return
	}

	var randOffset int = rand.Intn(len(foods[inputedFood]) - 1)
	var food string = foods[inputedFood][randOffset]
	var category string = categories[inputedFood]

	lib.BoldCyanPrintf("%s: %s\n", category, food)
}

func help() {
	w := tabwriter.NewWriter(os.Stdout, 0, 10, 2, ' ', 0)

	fmt.Fprintf(w, "%s is food recommender.\n", ProgramName)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "USAGE:")
	fmt.Fprintf(w, "\t%s [OPTIONS] [ARGUMENTS]...\n", ProgramName)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "ARGUMENTS:")
	fmt.Fprintln(w, "\t<none>\tRecommend a random food.")
	fmt.Fprintln(w, "\tkorean\tRecommend random Korean food.")
	fmt.Fprintln(w, "\tvegan\tRecommend random vegan food.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "OPTIONS:")
	fmt.Fprintln(w, "\t-h, --help\tShow this help message.")
	fmt.Fprintln(w, "\t-V, --version\tShow version information.")

	w.Flush()
}

func version() {
	fmt.Printf("%s %s\n", ProgramName, ProgramVersion)
}
