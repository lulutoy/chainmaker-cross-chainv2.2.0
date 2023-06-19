package main

import (
	"chainmaker.org/chainmaker-cross/sdk"
	"chainmaker.org/chainmaker-cross/sdk/builder"
	"context"
	"fmt"
	"time"
)

func main() {
	//chainmakerTochainmaker()
	chainmakerUser()
}

func chainmakerUser() {

	//用配置文件创建一个跨链SDK实例
	crossSDK, err := sdk.NewCrossSDK(sdk.WithConfigFile("/home/experiment/cross-chain/release/config/chainmaker/cross_chain_sdk.yml"))
	if err != nil {
		panic(err)
	}
	//require.NoError(t, err)
	//创建构造交易的上下文
	tx1Ctx := sdk.NewCrossTxBuildCtx(
		//链名字
		"chain1",
		//交易在跨链交易中的索引值
		0,
		//事务合约Excute时，要执行的业务合约信息
		builder.NewContract("balance_002", "Plus", builder.NewParamsWithMap(map[string]string{"number": "1"})),
		//事务合约Rollback时，要执行的业务合约信息
		builder.NewContract("balance_002", "Minus", builder.NewParamsWithMap(map[string]string{"number": "1"})),
	)
	tx2Ctx := sdk.NewCrossTxBuildCtx(
		//链名字
		"chain2",
		//交易在跨链交易中的索引值
		1,
		//事务合约Excute时，要执行的业务合约信息
		builder.NewContract("balance_002", "Minus", builder.NewParamsWithMap(map[string]string{"number": "1"})),
		//事务合约Rollback时，要执行的业务合约信息
		builder.NewContract("balance_002", "Plus", builder.NewParamsWithMap(map[string]string{"number": "1"})),
		//如果业务合约中使用proof，此可选项会将proofkey写入合约参数中，chainmaker中存放在key为proofKey的参数中, fabric写在参数列表的第一个参数
		builder.WithUseProofKey(true),
	)
	//生成跨链事件
	crossEvent, err := crossSDK.GenCrossEvent(tx1Ctx, tx2Ctx)
	//require.NoError(t, err)

	//发送跨链事件，参数syncResult代表是否同步等待跨链结果
	//SendCrossEvent(event *CrossEventContext, url string, syncResult bool, opts ...EventSendOption)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := crossSDK.SendCrossEvent(crossEvent, "http://192.168.30.128:8080", true, sdk.WithContextOpt(ctx), sdk.WithSyncStrategyOpt(sdk.SyncStrategy{
		0, 2 * time.Second, 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("SendCrossEvent->res:%v", res)
	//require.NoError(t, err)
	//获取crossID对应的跨链结果
	res, err = crossSDK.QueryCrossResult(crossEvent.GetCrossID(), "http://192.168.30.128:8080")
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("QueryCrossResult->res:%v", res)
	//require.NoError(t, err)

}

func chainmakerTofabric() {
	//生成CrossSDK实例
	crossSDK, err := sdk.NewCrossSDK(sdk.WithConfigFile("./cross_chain_sdk.yml"))
	if err != nil {
		panic(err)
	}
	//构造跨链事件
	crossEvent, err := crossSDK.GenCrossEvent(
		sdk.NewCrossTxBuildCtx(
			"baa1v210", 0,
			builder.NewContract("Balance", "Plus", builder.NewParamsWithMap(map[string]string{"number": "1"})),
			builder.NewContract("Balance", "Reset", nil),
		),
		sdk.NewCrossTxBuildCtx(
			"mychannel", 1,
			builder.NewContract("fabcar", "ChangeCarOwner", builder.NewParamsNoKeys("CAR0", "OWNER")),
			builder.NewContract("fabcar", "QueryAllCars", nil),
		),
	)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//发送跨链事件
	res, err := crossSDK.SendCrossEvent(crossEvent, "https://localhost:8080", true,
		sdk.WithContextOpt(ctx), sdk.WithSyncStrategyOpt(sdk.SyncStrategy{
			0, 20 * time.Second, 20 * time.Second,
		}))

	if err != nil {
		panic(err)
	}
	if res != nil {
		fmt.Printf("%+v\n", *res)
	}
}
