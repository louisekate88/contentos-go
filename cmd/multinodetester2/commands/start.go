package commands

import (
	"fmt"
	"github.com/coschain/cobra"
	cmd "github.com/coschain/contentos-go/cmd/cosd/commands"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/pprof"
	"github.com/coschain/contentos-go/config"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/mylog"
	"github.com/coschain/contentos-go/node"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var StartCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start cosd-path count(default 3)",
		Short: "start multi cosd node",
		Run:   StartNode,
	}
	cmd.Flags().Int64VarP(&NodeCnt, "number", "n", 3, "number of cosd thread")
	return cmd
}

type GlobalObject struct {
	sync.Mutex
	arr []*node.Node
	dposList []iservices.IConsensus
}

var globalObj GlobalObject

func StartNode(cmd *cobra.Command, args []string) {
	for i:=0;i<int(NodeCnt);i++{
		fmt.Println("i: ", i," NodeCnt: ", NodeCnt)
		app, cfg := makeNode(i)

		//go startNode(app, cfg)
		startNode(app, cfg)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("mian func")
	for i:=0;i<len(globalObj.dposList);i++ {
		fmt.Println()
		fmt.Println()
		fmt.Println("main func active producers:   ", globalObj.dposList[i].ActiveProducers())
		fmt.Println()
		fmt.Println()
	}

	now := time.Now()
	globalObj.dposList[0].MaybeProduceBlock(now)
	globalObj.dposList[0].MaybeProduceBlock(now.Add( 3 * time.Second))
	globalObj.dposList[0].MaybeProduceBlock(now.Add( 6 * time.Second))
	time.Sleep(10*time.Second)
	fmt.Println("head block id:   ", globalObj.dposList[0].GetHeadBlockId())
	fmt.Println("head block id:   ", globalObj.dposList[1].GetHeadBlockId())
	fmt.Println("head block id:   ", globalObj.dposList[2].GetHeadBlockId())

	WaitSignal()
}

func makeNode(index int) (*node.Node, node.Config) {
	var cfg node.Config
	confdir := filepath.Join(config.DefaultDataDir(), fmt.Sprintf("%s_%d", TesterClientIdentifier, index))
	fmt.Println("config dir: ", confdir)
	viper.Reset()
	viper.AddConfigPath(confdir)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err == nil {
		_ = viper.Unmarshal(&cfg)
	} else {
		fmt.Printf("fatal: not be initialized (do `init` first)\n")
		os.Exit(1)
	}
	if cfg.DataDir != "" {
		dir, err := filepath.Abs(cfg.DataDir)
		if err != nil {
			common.Fatalf("DataDir in cfg cannot be converted to absolute path")
		}
		cfg.DataDir = dir
	}
	app, err := node.New(&cfg)
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}
	fmt.Println("Name: ", cfg.Name)
	fmt.Println("p2p node port: ", cfg.P2P.NodePort)
	fmt.Println("p2p consensus port: ", cfg.P2P.NodeConsensusPort)
	return app, cfg
}

func startNode(app *node.Node, cfg node.Config) {
	app.Log = mylog.Init(cfg.ResolvePath("logs"), mylog.DebugLevel, 0)

	pprof.StartPprof()

	cmd.RegisterService(app, cfg)

	if err := app.Start(); err != nil {
		common.Fatalf("start node failed, err: %v\n", err)
	}

	globalObj.Lock()
	globalObj.arr = append(globalObj.arr, app)
	it, err := app.Service(iservices.ConsensusServerName)
	if err != nil {
		panic(err)
	}
	Icons := it.(iservices.IConsensus)
	Icons.ResetProdTimer( 86400 * time.Second )
	globalObj.dposList = append(globalObj.dposList, Icons)
	fmt.Println("append one to list")
	globalObj.Unlock()

	go app.Wait()
}

func WaitSignal() {
	SIGSTOP := syscall.Signal(0x13) //for windows compile
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, SIGSTOP, syscall.SIGUSR1, syscall.SIGUSR2)
	for {
		s := <-sigc
		fmt.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, SIGSTOP, syscall.SIGINT:
			for i:=0;i<len(globalObj.arr);i++ {
				globalObj.arr[i].MainLoop.Stop()
				globalObj.arr[i].Stop()
			}
			fmt.Println("Got interrupt, shutting down...")
			return
		case syscall.SIGHUP:
			fmt.Println("syscall.SIGHUP custom operation")
		case syscall.SIGUSR1:
			fmt.Println("syscall.SIGUSR1 custom operation")
		case syscall.SIGUSR2:
			fmt.Println("syscall.SIGUSR2 custom operation")
		default:
			return
		}
	}
}
