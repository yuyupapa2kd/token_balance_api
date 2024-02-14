package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/the-medium/token-balance-api/internal/config"
	"github.com/the-medium/token-balance-api/internal/contract"
	"github.com/the-medium/token-balance-api/internal/resource"
)

func GetCoinBalance(address string) (resource.ResBalance, error) {
	var result resource.ResBalance

	client, err := ethclient.Dial(config.RuntimeConf.RpcEndpoint)
	if err != nil {
		fmt.Println("err : ", err)
		return result, err
	}

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println("err : ", err)
		return result, err
	}

	result.Balance = balance.String()

	return result, nil
}

func GetTokenBalance(symbol, address string) (resource.ResBalance, error) {
	var result resource.ResBalance
	var ca string

	switch symbol {
	case "SOP":
		ca = config.RuntimeConf.TokenCA.SOP
	case "LOUI":
		ca = config.RuntimeConf.TokenCA.LOUI
	case "ksETH":
		ca = config.RuntimeConf.TokenCA.KsETH
	case "ksUSDT":
		ca = config.RuntimeConf.TokenCA.KsUSDT
	case "ksXRP":
		ca = config.RuntimeConf.TokenCA.KsXRP
	case "ksBNB":
		ca = config.RuntimeConf.TokenCA.KsBNB
	case "ksKLAY":
		ca = config.RuntimeConf.TokenCA.KsKLAY
	case "inKSTA":
		ca = config.RuntimeConf.TokenCA.InKSTA
	case "DLT":
		ca = config.RuntimeConf.TokenCA.DLT
	case "XABT":
		ca = config.RuntimeConf.TokenCA.XABT
	case "BOM":
		ca = config.RuntimeConf.TokenCA.BOM
	default:
		return result, errors.New("unsupported token symbol")
	}

	client, err := ethclient.Dial(config.RuntimeConf.RpcEndpoint)
	if err != nil {
		fmt.Println("err : ", err)
		return result, err
	}

	tokenContract, err := contract.NewERC20(common.HexToAddress(ca), client)
	if err != nil {
		return result, err
	}

	balance, err := tokenContract.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return result, err
	}

	result.Balance = balance.String()

	return result, nil
}
