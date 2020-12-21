package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	flg struct {
		n                int
		duration         time.Duration
		randomRace       bool
		randomGoMaxProcs bool
		randomParallel   bool
	}
	goMaxProcs     = []int{2, 4, 8, 16, 32, 64}
	goTestParallel = []int{2, 4, 8}
)

type params struct {
	args []string
	envs []string
}

func init() {
	flag.IntVar(&flg.n, "n", -1, "number of go test executions: -1 means infinite loop")
	flag.DurationVar(&flg.duration, "d", 0, "duration")
	flag.BoolVar(&flg.randomRace, "random-race", true, `randomly add "go test -race"`)
	flag.BoolVar(&flg.randomGoMaxProcs, "random-go-max-procs", true, `randomly change "GOMAXPROCS", 2,4,8,16,32,64`)
	flag.BoolVar(&flg.randomParallel, "random-paralell", true, `randomly change "go test -p", 2,4,8`)
	rand.Seed(time.Now().Unix())
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	log.Println("starting long test")

	ctx := context.Background()
	timer := time.NewTimer(flg.duration)
	defer timer.Stop()

	for i := 0; ; i++ {
		if 0 <= flg.n && flg.n <= i {
			break
		}

		p := createParams(args)
		log.Printf("test %d, %s go %s", i, strings.Join(p.envs, " "), strings.Join(p.args, " "))
		cmd := exec.CommandContext(ctx, "go", p.args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = append(os.Environ(), p.envs...)
		if err := cmd.Run(); err != nil {
			return err
		}

		select {
		case <-timer.C:
			if flg.duration != 0 {
				return nil
			}
		default:
		}
	}
	return nil
}

func createParams(args []string) *params {
	envs := []string{}
	testArgs := []string{"test", "-count=1"}
	if flg.randomRace && rand.Intn(2) == 0 {
		testArgs = append(testArgs, "-race")
	}
	if flg.randomGoMaxProcs {
		envs = append(envs, fmt.Sprintf("GOMAXPROCS=%d", goMaxProcs[rand.Intn(len(goMaxProcs))]))
	}
	if flg.randomParallel {
		testArgs = append(testArgs, fmt.Sprintf("-p=%d", goTestParallel[rand.Intn(len(goTestParallel))]))
	}
	for i, a := range args {
		if a == "--" && i+1 != len(args) {
			testArgs = append(testArgs, args[i+1:]...)
			break
		}
	}
	return &params{args: testArgs, envs: envs}
}
