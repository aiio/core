package ethx

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ETH struct {
	client *ethclient.Client
}

// NewETH eth nft
func NewETH(rawURL string) (*ETH, error) {
	var conn, err = ethclient.Dial(rawURL)
	if err != nil {
		return nil, err
	}
	return &ETH{client: conn}, nil
}

// GetName Erc721 查询合约名称
func (s *ETH) GetName(contractAddress string) (string, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721ABI))
	if err != nil {
		return "", err
	}

	contract, err := s.BoundContract(contractAddress, parsed)
	if err != nil {
		return "", err
	}

	// Erc721 查询合约名称
	var out []interface{}
	err = contract.Call(nil, &out, "name")
	if err != nil {
		return "", err
	}
	return out[0].(string), nil
}

// GetTokenURI Erc721 查询tokenURI
func (s *ETH) GetTokenURI(contractAddress string, tokenId int64) (string, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721ABI))
	if err != nil {
		return "", err
	}

	contract, err := s.BoundContract(contractAddress, parsed)
	if err != nil {
		return "", err
	}

	var out []interface{}
	err = contract.Call(nil, &out, "tokenURI", big.NewInt(tokenId))
	if err != nil {
		return "", err
	}
	return out[0].(string), nil
}

// GetURI Erc1155 查询URI
func (s *ETH) GetURI(contractAddress string, tokenId int64) (string, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155ABI))
	if err != nil {
		return "", err
	}

	contract, err := s.BoundContract(contractAddress, parsed)
	if err != nil {
		return "", err
	}

	var out []interface{}
	err = contract.Call(nil, &out, "uri", big.NewInt(tokenId))
	if err != nil {
		return "", err
	}
	return out[0].(string), nil
}

// BoundContract 绑定合约
func (s *ETH) BoundContract(contractAddress string, parsed abi.ABI) (*bind.BoundContract, error) {
	return bind.NewBoundContract(common.HexToAddress(contractAddress), parsed, s.client, s.client, s.client), nil
}

func DeployErc721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return bind.DeployContract(auth, parsed, common.FromHex(Erc721Bin), backend, name_, symbol_)
}

func DeployErc1155(auth *bind.TransactOpts, backend bind.ContractBackend, uri_ string) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc1155ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return bind.DeployContract(auth, parsed, common.FromHex(Erc1155Bin), backend, uri_)
}
