// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package contracts

import (
	"math/big"
	"strings"
	"errors"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)





	// ProductMetaData contains all meta data concerning the Product contract.
	var ProductMetaData = &bind.MetaData{
		ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"AcceptProduct\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"DelegateProduct\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"productId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NewProduct\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"}],\"name\":\"acceptProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"createProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_productId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"delegateProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"productToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
		Bin: "0x608060405234801561001057600080fd5b5061105a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806302ec06be14610067578063420fadc8146100835780634c5bb4dd146100b357806365078a0c146100cf5780637acc0b20146100eb578063949d225d1461011e575b600080fd5b610081600480360381019061007c9190610ab4565b61013c565b005b61009d60048036038101906100989190610b33565b6103ce565b6040516100aa9190610ba1565b60405180910390f35b6100cd60048036038101906100c89190610be8565b610401565b005b6100e960048036038101906100e49190610b33565b6105ba565b005b61010560048036038101906101009190610b33565b610796565b6040516101159493929190610ccc565b60405180910390f35b6101266108ab565b6040516101339190610d27565b60405180910390f35b600a600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054111561018957600080fd5b60006040518060800160405280838152602001600060ff1681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000190805190602001906102229291906108b7565b5060208201518160010160006101000a81548160ff021916908360ff16021790555060408201518160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600060016000805490506102e89190610d71565b9050336001600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919061038c90610da5565b91905055507ff21dbf9e399a0713acf6eda2b3cd910118a9989bd6e58e7c143405fc4ff2121881836040516103c2929190610dee565b60405180910390a15050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461046c57600080fd5b600080838154811061048157610480610e1e565b5b906000526020600020906003020160010160009054906101000a900460ff1660ff16146104e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104da90610e99565b60405180910390fd5b60008083815481106104f8576104f7610e1e565b5b9060005260206000209060030201905060018160010160006101000a81548160ff021916908360ff160217905550818160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd6f28d7dae1dcd8daebfd9cc968f71165e5063ea796de46f92a95a4da155e5e183838360010160009054906101000a900460ff166040516105ad93929190610eb9565b60405180910390a1505050565b6001600082815481106105d0576105cf610e1e565b5b906000526020600020906003020160010160009054906101000a900460ff1660ff16146105fc57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166000828154811061062757610626610e1e565b5b906000526020600020906003020160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461067957600080fd5b600080828154811061068e5761068d610e1e565b5b9060005260206000209060030201905060008160010160006101000a81548160ff021916908360ff16021790555060008160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550338160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f1ad74a28593362b47a0a3377ee29ac1d4c8473a8a120047b05fe6ef99ef5867582826000018360010160009054906101000a900460ff1660405161078a93929190610fe6565b60405180910390a15050565b600081815481106107a657600080fd5b90600052602060002090600302016000915090508060000180546107c990610f1f565b80601f01602080910402602001604051908101604052809291908181526020018280546107f590610f1f565b80156108425780601f1061081757610100808354040283529160200191610842565b820191906000526020600020905b81548152906001019060200180831161082557829003601f168201915b5050505050908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905084565b60008080549050905090565b8280546108c390610f1f565b90600052602060002090601f0160209004810192826108e5576000855561092c565b82601f106108fe57805160ff191683800117855561092c565b8280016001018555821561092c579182015b8281111561092b578251825591602001919060010190610910565b5b509050610939919061093d565b5090565b5b8082111561095657600081600090555060010161093e565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6109c182610978565b810181811067ffffffffffffffff821117156109e0576109df610989565b5b80604052505050565b60006109f361095a565b90506109ff82826109b8565b919050565b600067ffffffffffffffff821115610a1f57610a1e610989565b5b610a2882610978565b9050602081019050919050565b82818337600083830152505050565b6000610a57610a5284610a04565b6109e9565b905082815260208101848484011115610a7357610a72610973565b5b610a7e848285610a35565b509392505050565b600082601f830112610a9b57610a9a61096e565b5b8135610aab848260208601610a44565b91505092915050565b600060208284031215610aca57610ac9610964565b5b600082013567ffffffffffffffff811115610ae857610ae7610969565b5b610af484828501610a86565b91505092915050565b6000819050919050565b610b1081610afd565b8114610b1b57600080fd5b50565b600081359050610b2d81610b07565b92915050565b600060208284031215610b4957610b48610964565b5b6000610b5784828501610b1e565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b8b82610b60565b9050919050565b610b9b81610b80565b82525050565b6000602082019050610bb66000830184610b92565b92915050565b610bc581610b80565b8114610bd057600080fd5b50565b600081359050610be281610bbc565b92915050565b60008060408385031215610bff57610bfe610964565b5b6000610c0d85828601610b1e565b9250506020610c1e85828601610bd3565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610c62578082015181840152602081019050610c47565b83811115610c71576000848401525b50505050565b6000610c8282610c28565b610c8c8185610c33565b9350610c9c818560208601610c44565b610ca581610978565b840191505092915050565b600060ff82169050919050565b610cc681610cb0565b82525050565b60006080820190508181036000830152610ce68187610c77565b9050610cf56020830186610cbd565b610d026040830185610b92565b610d0f6060830184610b92565b95945050505050565b610d2181610afd565b82525050565b6000602082019050610d3c6000830184610d18565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d7c82610afd565b9150610d8783610afd565b925082821015610d9a57610d99610d42565b5b828203905092915050565b6000610db082610afd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610de357610de2610d42565b5b600182019050919050565b6000604082019050610e036000830185610d18565b8181036020830152610e158184610c77565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f697320616c72656164792064656c656761746564000000000000000000000000600082015250565b6000610e83601483610c33565b9150610e8e82610e4d565b602082019050919050565b60006020820190508181036000830152610eb281610e76565b9050919050565b6000606082019050610ece6000830186610d18565b610edb6020830185610b92565b610ee86040830184610cbd565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610f3757607f821691505b60208210811415610f4b57610f4a610ef0565b5b50919050565b60008190508160005260206000209050919050565b60008154610f7381610f1f565b610f7d8186610c33565b94506001821660008114610f985760018114610faa57610fdd565b60ff1983168652602086019350610fdd565b610fb385610f51565b60005b83811015610fd557815481890152600182019150602081019050610fb6565b808801955050505b50505092915050565b6000606082019050610ffb6000830186610d18565b818103602083015261100d8185610f66565b905061101c6040830184610cbd565b94935050505056fea2646970667358221220f577fd166cb09634bd4c32eb2a3a814748e47dfa52b4cc2a07e663ce2eb5b6af64736f6c634300080b0033",
		
	}
	// ProductABI is the input ABI used to generate the binding from.
	// Deprecated: Use ProductMetaData.ABI instead.
	var ProductABI = ProductMetaData.ABI

	

	
		// ProductBin is the compiled bytecode used for deploying new contracts.
		// Deprecated: Use ProductMetaData.Bin instead.
		var ProductBin = ProductMetaData.Bin

		// DeployProduct deploys a new Ethereum contract, binding an instance of Product to it.
		func DeployProduct(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Product, error) {
		  parsed, err := ProductMetaData.GetAbi()
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  if parsed == nil {
			return common.Address{}, nil, nil, errors.New("GetABI returned nil")
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProductBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &Product{ ProductCaller: ProductCaller{contract: contract}, ProductTransactor: ProductTransactor{contract: contract}, ProductFilterer: ProductFilterer{contract: contract} }, nil
		}
	

	// Product is an auto generated Go binding around an Ethereum contract.
	type Product struct {
	  ProductCaller     // Read-only binding to the contract
	  ProductTransactor // Write-only binding to the contract
	  ProductFilterer   // Log filterer for contract events
	}

	// ProductCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ProductCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ProductTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ProductTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ProductFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ProductFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ProductSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ProductSession struct {
	  Contract     *Product        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ProductCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ProductCallerSession struct {
	  Contract *ProductCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ProductTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ProductTransactorSession struct {
	  Contract     *ProductTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ProductRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ProductRaw struct {
	  Contract *Product // Generic contract binding to access the raw methods on
	}

	// ProductCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ProductCallerRaw struct {
		Contract *ProductCaller // Generic read-only contract binding to access the raw methods on
	}

	// ProductTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ProductTransactorRaw struct {
		Contract *ProductTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewProduct creates a new instance of Product, bound to a specific deployed contract.
	func NewProduct(address common.Address, backend bind.ContractBackend) (*Product, error) {
	  contract, err := bindProduct(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &Product{ ProductCaller: ProductCaller{contract: contract}, ProductTransactor: ProductTransactor{contract: contract}, ProductFilterer: ProductFilterer{contract: contract} }, nil
	}

	// NewProductCaller creates a new read-only instance of Product, bound to a specific deployed contract.
	func NewProductCaller(address common.Address, caller bind.ContractCaller) (*ProductCaller, error) {
	  contract, err := bindProduct(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ProductCaller{contract: contract}, nil
	}

	// NewProductTransactor creates a new write-only instance of Product, bound to a specific deployed contract.
	func NewProductTransactor(address common.Address, transactor bind.ContractTransactor) (*ProductTransactor, error) {
	  contract, err := bindProduct(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ProductTransactor{contract: contract}, nil
	}

	// NewProductFilterer creates a new log filterer instance of Product, bound to a specific deployed contract.
 	func NewProductFilterer(address common.Address, filterer bind.ContractFilterer) (*ProductFilterer, error) {
 	  contract, err := bindProduct(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ProductFilterer{contract: contract}, nil
 	}

	// bindProduct binds a generic wrapper to an already deployed contract.
	func bindProduct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ProductABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Product *ProductRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Product.Contract.ProductCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Product *ProductRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Product.Contract.ProductTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Product *ProductRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Product.Contract.ProductTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Product *ProductCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _Product.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Product *ProductTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Product.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Product *ProductTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Product.Contract.contract.Transact(opts, method, params...)
	}

	
		// ProductToOwner is a free data retrieval call binding the contract method 0x420fadc8.
		//
		// Solidity: function productToOwner(uint256 ) view returns(address)
		func (_Product *ProductCaller) ProductToOwner(opts *bind.CallOpts , arg0 *big.Int ) (common.Address, error) {
			var out []interface{}
			err := _Product.contract.Call(opts, &out, "productToOwner" , arg0)
			
			if err != nil {
				return *new(common.Address),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
			
			return out0,  err
			
		}

		// ProductToOwner is a free data retrieval call binding the contract method 0x420fadc8.
		//
		// Solidity: function productToOwner(uint256 ) view returns(address)
		func (_Product *ProductSession) ProductToOwner( arg0 *big.Int ) ( common.Address,  error) {
		  return _Product.Contract.ProductToOwner(&_Product.CallOpts , arg0)
		}

		// ProductToOwner is a free data retrieval call binding the contract method 0x420fadc8.
		//
		// Solidity: function productToOwner(uint256 ) view returns(address)
		func (_Product *ProductCallerSession) ProductToOwner( arg0 *big.Int ) ( common.Address,  error) {
		  return _Product.Contract.ProductToOwner(&_Product.CallOpts , arg0)
		}
	
		// Products is a free data retrieval call binding the contract method 0x7acc0b20.
		//
		// Solidity: function products(uint256 ) view returns(string name, uint8 status, address owner, address newOwner)
		func (_Product *ProductCaller) Products(opts *bind.CallOpts , arg0 *big.Int ) (struct{ Name string;Status uint8;Owner common.Address;NewOwner common.Address; }, error) {
			var out []interface{}
			err := _Product.contract.Call(opts, &out, "products" , arg0)
			
			outstruct := new(struct{  Name string;  Status uint8;  Owner common.Address;  NewOwner common.Address;  })
			if err != nil {
				return *outstruct, err
			}
			 
			outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string) 
			outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8) 
			outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address) 
			outstruct.NewOwner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

			return *outstruct, err
			
		}

		// Products is a free data retrieval call binding the contract method 0x7acc0b20.
		//
		// Solidity: function products(uint256 ) view returns(string name, uint8 status, address owner, address newOwner)
		func (_Product *ProductSession) Products( arg0 *big.Int ) (struct{ Name string;Status uint8;Owner common.Address;NewOwner common.Address; },  error) {
		  return _Product.Contract.Products(&_Product.CallOpts , arg0)
		}

		// Products is a free data retrieval call binding the contract method 0x7acc0b20.
		//
		// Solidity: function products(uint256 ) view returns(string name, uint8 status, address owner, address newOwner)
		func (_Product *ProductCallerSession) Products( arg0 *big.Int ) (struct{ Name string;Status uint8;Owner common.Address;NewOwner common.Address; },  error) {
		  return _Product.Contract.Products(&_Product.CallOpts , arg0)
		}
	
		// Size is a free data retrieval call binding the contract method 0x949d225d.
		//
		// Solidity: function size() view returns(uint256 count)
		func (_Product *ProductCaller) Size(opts *bind.CallOpts ) (*big.Int, error) {
			var out []interface{}
			err := _Product.contract.Call(opts, &out, "size" )
			
			if err != nil {
				return *new(*big.Int),  err
			}
			
			out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			
			return out0,  err
			
		}

		// Size is a free data retrieval call binding the contract method 0x949d225d.
		//
		// Solidity: function size() view returns(uint256 count)
		func (_Product *ProductSession) Size() ( *big.Int,  error) {
		  return _Product.Contract.Size(&_Product.CallOpts )
		}

		// Size is a free data retrieval call binding the contract method 0x949d225d.
		//
		// Solidity: function size() view returns(uint256 count)
		func (_Product *ProductCallerSession) Size() ( *big.Int,  error) {
		  return _Product.Contract.Size(&_Product.CallOpts )
		}
	

	
		// AcceptProduct is a paid mutator transaction binding the contract method 0x65078a0c.
		//
		// Solidity: function acceptProduct(uint256 _productId) returns()
		func (_Product *ProductTransactor) AcceptProduct(opts *bind.TransactOpts , _productId *big.Int ) (*types.Transaction, error) {
			return _Product.contract.Transact(opts, "acceptProduct" , _productId)
		}

		// AcceptProduct is a paid mutator transaction binding the contract method 0x65078a0c.
		//
		// Solidity: function acceptProduct(uint256 _productId) returns()
		func (_Product *ProductSession) AcceptProduct( _productId *big.Int ) (*types.Transaction, error) {
		  return _Product.Contract.AcceptProduct(&_Product.TransactOpts , _productId)
		}

		// AcceptProduct is a paid mutator transaction binding the contract method 0x65078a0c.
		//
		// Solidity: function acceptProduct(uint256 _productId) returns()
		func (_Product *ProductTransactorSession) AcceptProduct( _productId *big.Int ) (*types.Transaction, error) {
		  return _Product.Contract.AcceptProduct(&_Product.TransactOpts , _productId)
		}
	
		// CreateProduct is a paid mutator transaction binding the contract method 0x02ec06be.
		//
		// Solidity: function createProduct(string _name) returns()
		func (_Product *ProductTransactor) CreateProduct(opts *bind.TransactOpts , _name string ) (*types.Transaction, error) {
			return _Product.contract.Transact(opts, "createProduct" , _name)
		}

		// CreateProduct is a paid mutator transaction binding the contract method 0x02ec06be.
		//
		// Solidity: function createProduct(string _name) returns()
		func (_Product *ProductSession) CreateProduct( _name string ) (*types.Transaction, error) {
		  return _Product.Contract.CreateProduct(&_Product.TransactOpts , _name)
		}

		// CreateProduct is a paid mutator transaction binding the contract method 0x02ec06be.
		//
		// Solidity: function createProduct(string _name) returns()
		func (_Product *ProductTransactorSession) CreateProduct( _name string ) (*types.Transaction, error) {
		  return _Product.Contract.CreateProduct(&_Product.TransactOpts , _name)
		}
	
		// DelegateProduct is a paid mutator transaction binding the contract method 0x4c5bb4dd.
		//
		// Solidity: function delegateProduct(uint256 _productId, address _newOwner) returns()
		func (_Product *ProductTransactor) DelegateProduct(opts *bind.TransactOpts , _productId *big.Int , _newOwner common.Address ) (*types.Transaction, error) {
			return _Product.contract.Transact(opts, "delegateProduct" , _productId, _newOwner)
		}

		// DelegateProduct is a paid mutator transaction binding the contract method 0x4c5bb4dd.
		//
		// Solidity: function delegateProduct(uint256 _productId, address _newOwner) returns()
		func (_Product *ProductSession) DelegateProduct( _productId *big.Int , _newOwner common.Address ) (*types.Transaction, error) {
		  return _Product.Contract.DelegateProduct(&_Product.TransactOpts , _productId, _newOwner)
		}

		// DelegateProduct is a paid mutator transaction binding the contract method 0x4c5bb4dd.
		//
		// Solidity: function delegateProduct(uint256 _productId, address _newOwner) returns()
		func (_Product *ProductTransactorSession) DelegateProduct( _productId *big.Int , _newOwner common.Address ) (*types.Transaction, error) {
		  return _Product.Contract.DelegateProduct(&_Product.TransactOpts , _productId, _newOwner)
		}
	

	

	

	
		// ProductAcceptProductIterator is returned from FilterAcceptProduct and is used to iterate over the raw logs and unpacked data for AcceptProduct events raised by the Product contract.
		type ProductAcceptProductIterator struct {
			Event *ProductAcceptProduct // Event containing the contract specifics and raw log

			contract *bind.BoundContract // Generic contract to use for unpacking event data
			event    string              // Event name to use for unpacking event data

			logs chan types.Log        // Log channel receiving the found contract events
			sub  ethereum.Subscription // Subscription for errors, completion and termination
			done bool                  // Whether the subscription completed delivering logs
			fail error                 // Occurred error to stop iteration
		}
		// Next advances the iterator to the subsequent event, returning whether there
		// are any more events found. In case of a retrieval or parsing error, false is
		// returned and Error() can be queried for the exact failure.
		func (it *ProductAcceptProductIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ProductAcceptProduct)
					if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
						it.fail = err
						return false
					}
					it.Event.Raw = log
					return true

				default:
					return false
				}
			}
			// Iterator still in progress, wait for either a data or an error event
			select {
			case log := <-it.logs:
				it.Event = new(ProductAcceptProduct)
				if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
					it.fail = err
					return false
				}
				it.Event.Raw = log
				return true

			case err := <-it.sub.Err():
				it.done = true
				it.fail = err
				return it.Next()
			}
		}
		// Error returns any retrieval or parsing error occurred during filtering.
		func (it *ProductAcceptProductIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ProductAcceptProductIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ProductAcceptProduct represents a AcceptProduct event raised by the Product contract.
		type ProductAcceptProduct struct { 
			ProductId *big.Int; 
			Name string; 
			Status uint8; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterAcceptProduct is a free log retrieval operation binding the contract event 0x1ad74a28593362b47a0a3377ee29ac1d4c8473a8a120047b05fe6ef99ef58675.
		//
		// Solidity: event AcceptProduct(uint256 productId, string name, uint8 status)
 		func (_Product *ProductFilterer) FilterAcceptProduct(opts *bind.FilterOpts) (*ProductAcceptProductIterator, error) {
			
			
			
			

			logs, sub, err := _Product.contract.FilterLogs(opts, "AcceptProduct")
			if err != nil {
				return nil, err
			}
			return &ProductAcceptProductIterator{contract: _Product.contract, event: "AcceptProduct", logs: logs, sub: sub}, nil
 		}

		// WatchAcceptProduct is a free log subscription operation binding the contract event 0x1ad74a28593362b47a0a3377ee29ac1d4c8473a8a120047b05fe6ef99ef58675.
		//
		// Solidity: event AcceptProduct(uint256 productId, string name, uint8 status)
		func (_Product *ProductFilterer) WatchAcceptProduct(opts *bind.WatchOpts, sink chan<- *ProductAcceptProduct) (event.Subscription, error) {
			
			
			
			

			logs, sub, err := _Product.contract.WatchLogs(opts, "AcceptProduct")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ProductAcceptProduct)
						if err := _Product.contract.UnpackLog(event, "AcceptProduct", log); err != nil {
							return err
						}
						event.Raw = log

						select {
						case sink <- event:
						case err := <-sub.Err():
							return err
						case <-quit:
							return nil
						}
					case err := <-sub.Err():
						return err
					case <-quit:
						return nil
					}
				}
			}), nil
		}

		// ParseAcceptProduct is a log parse operation binding the contract event 0x1ad74a28593362b47a0a3377ee29ac1d4c8473a8a120047b05fe6ef99ef58675.
		//
		// Solidity: event AcceptProduct(uint256 productId, string name, uint8 status)
		func (_Product *ProductFilterer) ParseAcceptProduct(log types.Log) (*ProductAcceptProduct, error) {
			event := new(ProductAcceptProduct)
			if err := _Product.contract.UnpackLog(event, "AcceptProduct", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// ProductDelegateProductIterator is returned from FilterDelegateProduct and is used to iterate over the raw logs and unpacked data for DelegateProduct events raised by the Product contract.
		type ProductDelegateProductIterator struct {
			Event *ProductDelegateProduct // Event containing the contract specifics and raw log

			contract *bind.BoundContract // Generic contract to use for unpacking event data
			event    string              // Event name to use for unpacking event data

			logs chan types.Log        // Log channel receiving the found contract events
			sub  ethereum.Subscription // Subscription for errors, completion and termination
			done bool                  // Whether the subscription completed delivering logs
			fail error                 // Occurred error to stop iteration
		}
		// Next advances the iterator to the subsequent event, returning whether there
		// are any more events found. In case of a retrieval or parsing error, false is
		// returned and Error() can be queried for the exact failure.
		func (it *ProductDelegateProductIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ProductDelegateProduct)
					if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
						it.fail = err
						return false
					}
					it.Event.Raw = log
					return true

				default:
					return false
				}
			}
			// Iterator still in progress, wait for either a data or an error event
			select {
			case log := <-it.logs:
				it.Event = new(ProductDelegateProduct)
				if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
					it.fail = err
					return false
				}
				it.Event.Raw = log
				return true

			case err := <-it.sub.Err():
				it.done = true
				it.fail = err
				return it.Next()
			}
		}
		// Error returns any retrieval or parsing error occurred during filtering.
		func (it *ProductDelegateProductIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ProductDelegateProductIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ProductDelegateProduct represents a DelegateProduct event raised by the Product contract.
		type ProductDelegateProduct struct { 
			ProductId *big.Int; 
			NewOwner common.Address; 
			Status uint8; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterDelegateProduct is a free log retrieval operation binding the contract event 0xd6f28d7dae1dcd8daebfd9cc968f71165e5063ea796de46f92a95a4da155e5e1.
		//
		// Solidity: event DelegateProduct(uint256 productId, address newOwner, uint8 status)
 		func (_Product *ProductFilterer) FilterDelegateProduct(opts *bind.FilterOpts) (*ProductDelegateProductIterator, error) {
			
			
			
			

			logs, sub, err := _Product.contract.FilterLogs(opts, "DelegateProduct")
			if err != nil {
				return nil, err
			}
			return &ProductDelegateProductIterator{contract: _Product.contract, event: "DelegateProduct", logs: logs, sub: sub}, nil
 		}

		// WatchDelegateProduct is a free log subscription operation binding the contract event 0xd6f28d7dae1dcd8daebfd9cc968f71165e5063ea796de46f92a95a4da155e5e1.
		//
		// Solidity: event DelegateProduct(uint256 productId, address newOwner, uint8 status)
		func (_Product *ProductFilterer) WatchDelegateProduct(opts *bind.WatchOpts, sink chan<- *ProductDelegateProduct) (event.Subscription, error) {
			
			
			
			

			logs, sub, err := _Product.contract.WatchLogs(opts, "DelegateProduct")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ProductDelegateProduct)
						if err := _Product.contract.UnpackLog(event, "DelegateProduct", log); err != nil {
							return err
						}
						event.Raw = log

						select {
						case sink <- event:
						case err := <-sub.Err():
							return err
						case <-quit:
							return nil
						}
					case err := <-sub.Err():
						return err
					case <-quit:
						return nil
					}
				}
			}), nil
		}

		// ParseDelegateProduct is a log parse operation binding the contract event 0xd6f28d7dae1dcd8daebfd9cc968f71165e5063ea796de46f92a95a4da155e5e1.
		//
		// Solidity: event DelegateProduct(uint256 productId, address newOwner, uint8 status)
		func (_Product *ProductFilterer) ParseDelegateProduct(log types.Log) (*ProductDelegateProduct, error) {
			event := new(ProductDelegateProduct)
			if err := _Product.contract.UnpackLog(event, "DelegateProduct", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	
		// ProductNewProductIterator is returned from FilterNewProduct and is used to iterate over the raw logs and unpacked data for NewProduct events raised by the Product contract.
		type ProductNewProductIterator struct {
			Event *ProductNewProduct // Event containing the contract specifics and raw log

			contract *bind.BoundContract // Generic contract to use for unpacking event data
			event    string              // Event name to use for unpacking event data

			logs chan types.Log        // Log channel receiving the found contract events
			sub  ethereum.Subscription // Subscription for errors, completion and termination
			done bool                  // Whether the subscription completed delivering logs
			fail error                 // Occurred error to stop iteration
		}
		// Next advances the iterator to the subsequent event, returning whether there
		// are any more events found. In case of a retrieval or parsing error, false is
		// returned and Error() can be queried for the exact failure.
		func (it *ProductNewProductIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ProductNewProduct)
					if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
						it.fail = err
						return false
					}
					it.Event.Raw = log
					return true

				default:
					return false
				}
			}
			// Iterator still in progress, wait for either a data or an error event
			select {
			case log := <-it.logs:
				it.Event = new(ProductNewProduct)
				if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
					it.fail = err
					return false
				}
				it.Event.Raw = log
				return true

			case err := <-it.sub.Err():
				it.done = true
				it.fail = err
				return it.Next()
			}
		}
		// Error returns any retrieval or parsing error occurred during filtering.
		func (it *ProductNewProductIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ProductNewProductIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ProductNewProduct represents a NewProduct event raised by the Product contract.
		type ProductNewProduct struct { 
			ProductId *big.Int; 
			Name string; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterNewProduct is a free log retrieval operation binding the contract event 0xf21dbf9e399a0713acf6eda2b3cd910118a9989bd6e58e7c143405fc4ff21218.
		//
		// Solidity: event NewProduct(uint256 productId, string name)
 		func (_Product *ProductFilterer) FilterNewProduct(opts *bind.FilterOpts) (*ProductNewProductIterator, error) {
			
			
			

			logs, sub, err := _Product.contract.FilterLogs(opts, "NewProduct")
			if err != nil {
				return nil, err
			}
			return &ProductNewProductIterator{contract: _Product.contract, event: "NewProduct", logs: logs, sub: sub}, nil
 		}

		// WatchNewProduct is a free log subscription operation binding the contract event 0xf21dbf9e399a0713acf6eda2b3cd910118a9989bd6e58e7c143405fc4ff21218.
		//
		// Solidity: event NewProduct(uint256 productId, string name)
		func (_Product *ProductFilterer) WatchNewProduct(opts *bind.WatchOpts, sink chan<- *ProductNewProduct) (event.Subscription, error) {
			
			
			

			logs, sub, err := _Product.contract.WatchLogs(opts, "NewProduct")
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ProductNewProduct)
						if err := _Product.contract.UnpackLog(event, "NewProduct", log); err != nil {
							return err
						}
						event.Raw = log

						select {
						case sink <- event:
						case err := <-sub.Err():
							return err
						case <-quit:
							return nil
						}
					case err := <-sub.Err():
						return err
					case <-quit:
						return nil
					}
				}
			}), nil
		}

		// ParseNewProduct is a log parse operation binding the contract event 0xf21dbf9e399a0713acf6eda2b3cd910118a9989bd6e58e7c143405fc4ff21218.
		//
		// Solidity: event NewProduct(uint256 productId, string name)
		func (_Product *ProductFilterer) ParseNewProduct(log types.Log) (*ProductNewProduct, error) {
			event := new(ProductNewProduct)
			if err := _Product.contract.UnpackLog(event, "NewProduct", log); err != nil {
				return nil, err
			}
			event.Raw = log
			return event, nil
		}

 	


