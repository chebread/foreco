package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"

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

var koreanFoods = []string{
	// 배달 및 외식 메뉴
	"치킨",
	"마라탕",
	"엽기떡볶이",
	"피자",
	"족발·보쌈",
	"막창·곱창",
	"돈가스",
	"찜닭",
	"회",
	"삼겹살",
	"양꼬치",
	"감자탕",
	"짜장면",
	"짬뽕",
	"육회",
	"닭갈비",
	"양대창",
	"연어초밥",
	"초밥",
	"파스타",
	"햄버거",
	"스테이크",
	"브런치 메뉴",
	"마라샹궈",
	"훠궈",
	"김치찜",
	"부대찌개",
	"곱도리탕",
	"아귀찜",
	"쭈꾸미볶음",
	// 간편식 및 분식
	"샌드위치",
	"컵밥",
	"김밥",
	"컵라면",
	"즉석밥",
	"밀키트",
	"만두",
	"냉면",
	"쫄면",
	"닭강정",
	"컵떡볶이",
	"주먹밥",
	"튀김",
	"순대",
	"어묵",
	"핫도그",
	// 디저트 및 간식
	"탕후루",
	"크로플",
	"소금빵",
	"약과",
	"호두과자",
	"크림치즈",
	"베이글",
	"마카롱",
	"케이크",
	"아이스크림",
	"그릭 요거트",
	"젤리",
	"빵",
	"쿠키",
	"프레첼",
	"꽈배기",
	"찹쌀떡",
	"버블티",
	"스무디",
	"에이드",
	"커피",
	"밀크티",
	"감자튀김",
	"호떡",
	"붕어빵",
	"닭꼬치",
	"떡볶이와 함께 먹는 꼬북칩",
	// 한식 및 기타
	"김치찌개",
	"된장찌개",
	"국밥",
	"쌀국수",
	"비빔밥",
	"간장게장",
	"제육볶음",
	"갈비",
	"순두부찌개",
	"설렁탕",
	"갈비탕",
	"육개장",
	"샤브샤브",
	"콩국수",
	"냉모밀",
	"김치볶음밥",
	"오므라이스",
	"순두부",
	"비빔국수",
	"칼국수",
	"수제비",
	"해물파전",
	"김치전",
	"떡국",
	"라멘",
	"카레",
}

var veganFoods = []string{
	"비빔밥",
	"잡채",
	"김치찌개",
	"된장찌개",
	"두부조림",
	"순두부찌개",
	"버섯전골",
	"콩나물국",
	"야채전",
	"파래전",
	"김치전",
	"감자전",
	"비건 만두",
	"콩국수",
	"묵사발",
	"고구마순 볶음",
	"채개장",
	"떡볶이",
	"비건 떡국",
	"들깨 수제비",
	"우엉 들깨탕",
	"비건 김밥",
	"쌈밥",
	"알리오 올리오 파스타",
	"토마토 라구 파스타",
	"비건 크림 파스타",
	"마라 크림 떡볶이",
	"뇨끼",
	"라자냐",
	"피자",
	"콜리플라워 스테이크",
	"비건 버거",
	"비건 타코",
	"부리또 볼",
	"감바스 알 아히요",
	"비건 샌드위치",
	"비건 랩",
	"포케",
	"렌틸콩 수프",
	"버섯 리조또",
	"비건 카츠 커리",
	"비건 뇨끼",
	"토마토 스튜",
	"짜장면",
	"중화풍 오이 쫄면",
	"얌운센",
	"코코넛 커리",
	"팟타이",
	"마라탕",
	"채식 탄탄면",
	"비건 네기토로 돈",
	"마파두부",
	"마라샹궈",
	"비건 깐풍기",
	"비건 유린기",
	"량피",
	"고추잡채",
	"비건 반미 샌드위치",
	"볶음밥",
	"비건 초밥",
	"채소 덮밥",
	"쌀국수",
	"두부 탄두리",
	"찹쌀루니",
	"비건 케이크",
	"비건 쿠키",
	"비건 머핀",
	"비건 마카롱",
	"비건 젤라또",
	"비건 아이스크림",
	"아몬드 초코볼",
	"비건 도넛",
	"미니 약과",
	"길쭉이 보리과자",
	"미니 달고나",
	"비건 초콜릿",
	"비건 베이글",
	"식물성 요거트",
	"구운 고구마",
	"군밤",
	"호떡",
	"견과류",
	"아몬드",
	"호두",
	"캐슈너트",
	"제철 과일",
	"샐러드",
	"포테이토 샐러드",
	"두부 샐러드",
	"병아리콩 샐러드",
	"현미밥",
	"고구마밥",
	"버섯 떡갈비",
	"콩까스",
	"참나물 페스토 파스타",
	"채소구이",
	"콩고기",
	"렌틸콩 요리",
	"콩물",
	"오이 소박이",
	"버섯 강정",
	"두부 동그랑땡",
	"콜리플라워 볶음밥",
	"퀴노아 샐러드",
	"낫또",
}

var foods = map[string][]string{
	"korean": koreanFoods,
	"vegan":  veganFoods,
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
	fmt.Fprintln(w, "\t-h, --help\tShow this help message")
	fmt.Fprintln(w, "\t-V, --version\tShow version information")

	w.Flush()
}

func version() {
	fmt.Printf("%s %s\n", ProgramName, ProgramVersion)
}
