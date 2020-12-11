package command

import (
	"flag"
	"fmt"
	"github.com/pefish/go-coin-eth"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	go_logger "github.com/pefish/go-logger"
	go_reflect "github.com/pefish/go-reflect"
	telegram_robot "github.com/pefish/telegram-bot-manager/pkg/telegram-robot"
	"github.com/pefish/telegram-bot-manager/pkg/telegram-sender"
	"github.com/pkg/errors"
	"strings"
)

type DefaultCommand struct {

}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{

	}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("ws-server", "wss://bsc-ws-node.nariox.org:443", "ws server")
	flagSet.String("contract-address", "0xA1f482Dc58145Ba2210bC21878Ca34000E2e8fE4", "contract address")
	flagSet.String("telegram-token", "", "telegram token")
	flagSet.String("chat-id", "", "chat id of group")
	return nil
}

func (dc *DefaultCommand) OnExited() error {
	return nil
}

func (dc *DefaultCommand) Start(data commander.StartData) error {
	wsServer, err := go_config.Config.GetString("ws-server")
	if err != nil {
		return err
	}
	contractAddress, err := go_config.Config.GetString("contract-address")
	if err != nil {
		return err
	}
	telegramToken, err := go_config.Config.GetString("telegram-token")
	if err != nil {
		return err
	}
	if telegramToken == "" {
		return errors.New("telegram token must be set")
	}
	chatIdStr, err := go_config.Config.GetString("chat-id")
	if err != nil {
		return err
	}
	if chatIdStr == "" {
		return errors.New("chatId must be set")
	}
	chatId, err := go_reflect.Reflect.ToInt64(chatIdStr)
	if err != nil {
		return err
	}
	telegramRobot := telegram_robot.NewRobot("", telegramToken)
	telegramRobot.SetLogger(go_logger.Logger)
	wallet, err := go_coin_eth.NewWallet(wsServer)
	if err != nil {
		return err
	}
	resultChan := make(chan map[string]interface{})
	errChan := make(chan error)
	sub, err := wallet.WatchLogs(
		resultChan,
		errChan,
		contractAddress,
		`[{"inputs":[{"internalType":"address","name":"admin_","type":"address"},{"internalType":"uint256","name":"delay_","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"txHash","type":"bytes32"},{"indexed":true,"internalType":"address","name":"target","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"},{"indexed":false,"internalType":"string","name":"signature","type":"string"},{"indexed":false,"internalType":"bytes","name":"data","type":"bytes"},{"indexed":false,"internalType":"uint256","name":"eta","type":"uint256"}],"name":"CancelTransaction","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"txHash","type":"bytes32"},{"indexed":true,"internalType":"address","name":"target","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"},{"indexed":false,"internalType":"string","name":"signature","type":"string"},{"indexed":false,"internalType":"bytes","name":"data","type":"bytes"},{"indexed":false,"internalType":"uint256","name":"eta","type":"uint256"}],"name":"ExecuteTransaction","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newAdmin","type":"address"}],"name":"NewAdmin","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"newDelay","type":"uint256"}],"name":"NewDelay","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newPendingAdmin","type":"address"}],"name":"NewPendingAdmin","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"txHash","type":"bytes32"},{"indexed":true,"internalType":"address","name":"target","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"},{"indexed":false,"internalType":"string","name":"signature","type":"string"},{"indexed":false,"internalType":"bytes","name":"data","type":"bytes"},{"indexed":false,"internalType":"uint256","name":"eta","type":"uint256"}],"name":"QueueTransaction","type":"event"},{"inputs":[],"name":"GRACE_PERIOD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"MAXIMUM_DELAY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"MINIMUM_DELAY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"acceptAdmin","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"admin","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"admin_initialized","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"target","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"string","name":"signature","type":"string"},{"internalType":"bytes","name":"data","type":"bytes"},{"internalType":"uint256","name":"eta","type":"uint256"}],"name":"cancelTransaction","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"delay","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"target","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"string","name":"signature","type":"string"},{"internalType":"bytes","name":"data","type":"bytes"},{"internalType":"uint256","name":"eta","type":"uint256"}],"name":"executeTransaction","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"pendingAdmin","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"target","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"string","name":"signature","type":"string"},{"internalType":"bytes","name":"data","type":"bytes"},{"internalType":"uint256","name":"eta","type":"uint256"}],"name":"queueTransaction","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"queuedTransactions","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"delay_","type":"uint256"}],"name":"setDelay","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"pendingAdmin_","type":"address"}],"name":"setPendingAdmin","outputs":[],"stateMutability":"nonpayable","type":"function"},{"stateMutability":"payable","type":"receive"}]`,
		"QueueTransaction",
		nil)
	if err != nil {
		return err
	}
	err = telegramRobot.TelegramSender().SendMsg(telegram_sender.MsgStruct{
		ChatId: chatId,
		Msg:    []byte(fmt.Sprintf("时间锁监控已启动... 监控地址为 %s", contractAddress)),
	}, 0)
	if err != nil {
		return err
	}
	go_logger.Logger.Info("watching...")
	for {
		select {
		case result := <- resultChan:
			methodStr := result["signature"].(string)
			if strings.Contains(methodStr, "setMigrator") {
				methodStr = "⚠️：setMigrator被调用，请核实"
			}
			if strings.Contains(methodStr, "setDelay") {
				methodStr = "⚠️：setDelay被调用，请核实"
			}
			err := telegramRobot.TelegramSender().SendMsg(telegram_sender.MsgStruct{
				ChatId: chatId,
				Msg:    []byte(methodStr),
			}, 0)
			if err != nil {
				return err
			}
		case err := <- errChan:
			sub.Unsubscribe()
			return err
		}
	}
}

