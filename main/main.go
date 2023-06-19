/*
 Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
   SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"chainmaker.org/chainmaker-cross/main/cmd"
	"github.com/spf13/cobra"
)

func main() {
	mainCmd := &cobra.Command{Use: "main start"}
	fmt.Println("------初始化开始--------")
	mainCmd.AddCommand(cmd.StartCMD())
	fmt.Println("------初始化结束--------")
	fmt.Println("------执行开始--------")
	err := mainCmd.Execute()
	fmt.Println("------执行结束--------")
	if err != nil {
		_ = fmt.Errorf("cross proxy start error, %v", err)
	}
}
