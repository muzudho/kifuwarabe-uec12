package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"

	c "github.com/muzudho/kifuwarabe-uec12/controller"
	e "github.com/muzudho/kifuwarabe-uec12/entities"
)

// KifuwarabeV1 - きふわらべバージョン１。
// NNGSへの接続を試みる。
func KifuwarabeV1() {
	e.G.Chat.Trace("# きふわらべv1プログラム開始☆（＾～＾）\n")

	config := c.LoadGameConf("resources/kifuwarabe-v1.gameConf.toml")

	e.G.Chat.Trace("# Config読んだ☆（＾～＾）\n")
	e.G.Chat.Trace("# Komi=%f\n", config.Game.Komi)
	e.G.Chat.Trace("# BoardSize=%d\n", config.Game.BoardSize)
	e.G.Chat.Trace("# MaxMoves=%d\n", config.Game.MaxMoves)
	e.G.Chat.Trace("# BoardData=%s\n", config.Game.BoardData)
	e.G.Chat.Trace("# SentinelBoardMax()=%d\n", config.SentinelBoardMax())

	board := e.NewBoardV9a(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardMax(), config.Komi(), config.MaxMoves())
	// presenter := p.NewPresenterV9a()

	e.G.Chat.Trace("# 盤を新規作成した☆（＾～＾）\n")

	rand.Seed(time.Now().UnixNano())

	e.G.Chat.Trace("# (^q^) ランダムの種を設定したぜ☆\n")

	board.InitBoard()

	e.G.Chat.Trace("# NNGSへの接続を試みるぜ☆（＾～＾）\n")

	nngsConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Nngs.Server, config.Nngs.Port))
	if err != nil {
		panic(err)
	}

	e.G.Chat.Trace("# NNGSへ接続でけた☆（＾～＾）\n")

	e.G.Chat.Trace("# NNGSへユーザー名 %s を送ったろ……☆（＾～＾）\n", config.Nngs.User)

	e.G.Chat.Send(nngsConn, fmt.Sprintf("%s\n", config.Nngs.User))

	e.G.Chat.Trace("# NNGSからの返信を待と……☆（＾～＾）\n")

	status, err := bufio.NewReader(nngsConn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%s", status)

}