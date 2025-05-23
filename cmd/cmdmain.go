package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/pprof"
	"syscall"
	"time"

	"k8s.io/klog/v2"

	"github.com/kubeovn/kube-ovn/cmd/controller"
	"github.com/kubeovn/kube-ovn/cmd/daemon"
	"github.com/kubeovn/kube-ovn/cmd/ovn_ic_controller"
	"github.com/kubeovn/kube-ovn/cmd/ovn_leader_checker"
	"github.com/kubeovn/kube-ovn/cmd/ovn_monitor"
	"github.com/kubeovn/kube-ovn/cmd/pinger"
	"github.com/kubeovn/kube-ovn/cmd/speaker"
	"github.com/kubeovn/kube-ovn/pkg/util"
)

const (
	CmdController       = "kube-ovn-controller"
	CmdDaemon           = "kube-ovn-daemon"
	CmdMonitor          = "kube-ovn-monitor"
	CmdPinger           = "kube-ovn-pinger"
	CmdSpeaker          = "kube-ovn-speaker"
	CmdOvnLeaderChecker = "kube-ovn-leader-checker"
	CmdOvnICController  = "kube-ovn-ic-controller"
)

const timeFormat = "2006-01-02_15:04:05"

func dumpProfile() {
	ch1 := make(chan os.Signal, 1)
	ch2 := make(chan os.Signal, 1)
	signal.Notify(ch1, syscall.SIGUSR1)
	signal.Notify(ch2, syscall.SIGUSR2)
	go func() {
		for {
			<-ch1
			name := fmt.Sprintf("cpu-profile-%s.pprof", time.Now().Format(timeFormat))
			f, err := os.Create(filepath.Join(os.TempDir(), name)) // #nosec G303,G304
			if err != nil {
				klog.Errorf("failed to create cpu profile file: %v", err)
				return
			}
			if err = pprof.StartCPUProfile(f); err != nil {
				klog.Errorf("failed to start cpu profile: %v", err)
				return
			}
			defer f.Close()
			time.Sleep(30 * time.Second)
			pprof.StopCPUProfile()
		}
	}()
	go func() {
		for {
			<-ch2
			name := fmt.Sprintf("mem-profile-%s.pprof", time.Now().Format(timeFormat))
			f, err := os.Create(filepath.Join(os.TempDir(), name)) // #nosec G303,G304
			if err != nil {
				klog.Errorf("failed to create memory profile file: %v", err)
				return
			}
			if err = pprof.WriteHeapProfile(f); err != nil {
				klog.Errorf("failed to write memory profile file: %v", err)
			}
			defer f.Close()
		}
	}()
}

func main() {
	cmd := filepath.Base(os.Args[0])
	switch cmd {
	case CmdController:
		dumpProfile()
		controller.CmdMain()
	case CmdDaemon:
		dumpProfile()
		daemon.CmdMain()
	case CmdMonitor:
		dumpProfile()
		ovn_monitor.CmdMain()
	case CmdPinger:
		dumpProfile()
		pinger.CmdMain()
	case CmdSpeaker:
		dumpProfile()
		speaker.CmdMain()
	case CmdOvnLeaderChecker:
		ovn_leader_checker.CmdMain()
	case CmdOvnICController:
		ovn_ic_controller.CmdMain()
	default:
		util.LogFatalAndExit(nil, "%s is an unknown command", cmd)
	}
}
