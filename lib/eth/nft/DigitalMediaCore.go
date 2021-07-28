// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DigitalMediaCoreABI is the input ABI used to generate the binding from.
const DigitalMediaCoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"currentStartingDigitalMediaId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_metadataPath\",\"type\":\"string\"}],\"name\":\"createCollection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_totalSupply\",\"type\":\"uint32\"},{\"name\":\"_digitalMediaMetadataPath\",\"type\":\"string\"},{\"name\":\"_collectionMetadataPath\",\"type\":\"string\"},{\"name\":\"_numReleases\",\"type\":\"uint32\"}],\"name\":\"oboCreateDigitalMediaAndReleasesInNewCollection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"singleCreatorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDigitalMediaStore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digitalMediaId\",\"type\":\"uint256\"}],\"name\":\"burnDigitalMedia\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToDigitalMediaRelease\",\"outputs\":[{\"name\":\"printEdition\",\"type\":\"uint32\"},{\"name\":\"digitalMediaId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_creatorAddress\",\"type\":\"address\"}],\"name\":\"removeApprovedTokenCreator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedCreators\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDigitalMedia\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"totalSupply\",\"type\":\"uint32\"},{\"name\":\"printIndex\",\"type\":\"uint32\"},{\"name\":\"collectionId\",\"type\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"metadataPath\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getCollection\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"metadataPath\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_creator\",\"type\":\"address\"},{\"name\":\"_newCreator\",\"type\":\"address\"}],\"name\":\"changeCreator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint32\"},{\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"name\":\"_metadataPath\",\"type\":\"string\"}],\"name\":\"createDigitalMedia\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_totalSupply\",\"type\":\"uint32\"},{\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"name\":\"_metadataPath\",\"type\":\"string\"},{\"name\":\"_numReleases\",\"type\":\"uint32\"}],\"name\":\"oboCreateDigitalMediaAndReleases\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setOboApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burnToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_digitalMediaId\",\"type\":\"uint256\"},{\"name\":\"_numReleases\",\"type\":\"uint32\"}],\"name\":\"createDigitalMediaReleases\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDigitalMediaRelease\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"printEdition\",\"type\":\"uint32\"},{\"name\":\"digitalMediaId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint32\"},{\"name\":\"_digitalMediaMetadataPath\",\"type\":\"string\"},{\"name\":\"_collectionMetadataPath\",\"type\":\"string\"},{\"name\":\"_numReleases\",\"type\":\"uint32\"}],\"name\":\"createDigitalMediaAndReleasesInNewCollection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dmsAddress\",\"type\":\"address\"}],\"name\":\"setV1DigitalMediaStoreAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oboAddress\",\"type\":\"address\"}],\"name\":\"disableOboAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1DigitalMediaStore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCreatorAddress\",\"type\":\"address\"}],\"name\":\"changeSingleCreator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_digitalMediaId\",\"type\":\"uint256\"},{\"name\":\"_numReleases\",\"type\":\"uint32\"}],\"name\":\"oboCreateDigitalMediaReleases\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creatorRegistryStore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"resetApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTokenCreators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"disabledOboOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint32\"},{\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"name\":\"_metadataPath\",\"type\":\"string\"},{\"name\":\"_numReleases\",\"type\":\"uint32\"}],\"name\":\"createDigitalMediaAndReleases\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_creatorAddress\",\"type\":\"address\"}],\"name\":\"addApprovedTokenCreator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenName\",\"type\":\"string\"},{\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"name\":\"_tokenIdStartingCounter\",\"type\":\"uint256\"},{\"name\":\"_dmsAddress\",\"type\":\"address\"},{\"name\":\"_crsAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"OboApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"OboDisabledForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"printEdition\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"digitalMediaId\",\"type\":\"uint256\"}],\"name\":\"DigitalMediaReleaseCreateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"storeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"printIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metadataPath\",\"type\":\"string\"}],\"name\":\"DigitalMediaCreateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"storeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"metadataPath\",\"type\":\"string\"}],\"name\":\"DigitalMediaCollectionCreateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"storeContractAddress\",\"type\":\"address\"}],\"name\":\"DigitalMediaBurnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DigitalMediaReleaseBurnEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"digitalMediaId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"printEdition\",\"type\":\"uint32\"}],\"name\":\"UpdateDigitalMediaPrintIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"ChangedCreator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousCreatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newCreatorAddress\",\"type\":\"address\"}],\"name\":\"SingleCreatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DigitalMediaCoreBin is the compiled bytecode used for deploying new contracts.
var DigitalMediaCoreBin = "0x608060405260008060146101000a81548160ff02191690831515021790555060006013553480156200003057600080fd5b506040516200753b3803806200753b83398101806040528101908080518201929190602001805182019291906020018051906020019092919080519060200190929190805190602001909291905050508484848282336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160099080519060200190620000dd929190620006d0565b5080600a9080519060200190620000f6929190620006d0565b505050806013819055505050506001600060146101000a81548160ff02191690831515021790555062000138826200015d640100000000026401000000009004565b62000152816200042b640100000000026401000000009004565b50505050506200077f565b600081905060028173ffffffffffffffffffffffffffffffffffffffff16636c669f256040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620001c957600080fd5b505af1158015620001de573d6000803e3d6000fd5b505050506040513d6020811015620001f557600080fd5b81019080805190602001909291905050501415156200027c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e636f72726563742076657273696f6e2e000000000000000000000000000081525060200191505060405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a60800b86040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156200034457600080fd5b505af115801562000359573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b9905d116040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620003e457600080fd5b505af1158015620003f9573d6000803e3d6000fd5b505050506040513d60208110156200041057600080fd5b81019080805190602001909291905050506003819055505050565b600081905060018173ffffffffffffffffffffffffffffffffffffffff16630d8e6e2c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200049757600080fd5b505af1158015620004ac573d6000803e3d6000fd5b505050506040513d6020811015620004c357600080fd5b8101908080519060200190929190505050141515620004e157600080fd5b60405180807f617070726f76656443726561746f72526567697374727900000000000000000081525060170190506040518091039020600019168173ffffffffffffffffffffffffffffffffffffffff1663c20a03826040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b1580156200058057600080fd5b505af115801562000595573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015620005c057600080fd5b810190808051640100000000811115620005d957600080fd5b82810190506020810184811115620005f057600080fd5b81518560018202830111640100000000821117156200060e57600080fd5b50509291905050506040518082805190602001908083835b6020831015156200064d578051825260208201915060208101905060208303925062000626565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019161415156200068b57600080fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200071357805160ff191683800117855562000744565b8280016001018555821562000744579182015b828111156200074357825182559160200191906001019062000726565b5b50905062000753919062000757565b5090565b6200077c91905b80821115620007785760008160009055506001016200075e565b5090565b90565b616dac806200078f6000396000f30060806040526004361061027d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630154788d1461028257806301ffc9a7146102ad578063059dfe131461031157806306a186d11461034c57806306fdde03146103df578063081812fc1461046f578063095ea7b3146104dc578063147ca2af1461052957806318160ddd146105805780631b284e75146105ab57806323077f581461060257806323b872dd1461062f5780632f745c591461069c5780633db57cbe146106fd5780633f4ba83a1461075157806342842e0e1461076857806342966c68146107d55780634751ae99146108025780634f558e79146108455780634f6ccce71461088a57806355232467146108cb57806355df42751461094e5780635a1f3c2814610a5b5780635c975abb14610b3b5780635cf09fee14610b6a578063603417fb14610bcd5780636352211e14610c225780636ee17a7814610c8f57806370a0823114610d14578063774f99d014610d6b5780637b47ec1a14610dba5780637fb08c9b14610de75780638456cb5914610e245780638a603bdf14610e3b5780638b40e8aa14610e965780638da5cb5b14610f0957806391c6078814610f6057806392b8bde114610fa357806395d89b4114610fe6578063a0f01e0814611076578063a22cb465146110cd578063a7df572c1461111c578063b36f0faa1461115f578063b39ba60a146111bc578063b88d4fde14611213578063c87b56dd146112c6578063cf8dae051461136c578063d70f575f14611399578063e158386d146113f4578063e985e9c51461144f578063eca211e3146114ca578063f2fde38b1461152f578063fafa8a4b14611572575b600080fd5b34801561028e57600080fd5b506102976115b5565b6040518082815260200191505060405180910390f35b3480156102b957600080fd5b506102f760048036038101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506115bb565b604051808215151515815260200191505060405180910390f35b34801561031d57600080fd5b5061034a600480360381019080803590602001908201803590602001919091929391929390505050611827565b005b34801561035857600080fd5b506103dd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff169060200190929190803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803563ffffffff169060200190929190505050611884565b005b3480156103eb57600080fd5b506103f4611a5c565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610434578082015181840152602081019050610419565b50505050905090810190601f1680156104615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561047b57600080fd5b5061049a60048036038101908080359060200190929190505050611afe565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104e857600080fd5b50610527600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611b3b565b005b34801561053557600080fd5b5061053e611d01565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561058c57600080fd5b50610595611d27565b6040518082815260200191505060405180910390f35b3480156105b757600080fd5b506105c0611d34565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561060e57600080fd5b5061062d60048036038101908080359060200190929190505050611d5a565b005b34801561063b57600080fd5b5061069a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611d83565b005b3480156106a857600080fd5b506106e7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611e9a565b6040518082815260200191505060405180910390f35b34801561070957600080fd5b5061072860048036038101908080359060200190929190505050611f11565b604051808363ffffffff1663ffffffff1681526020018281526020019250505060405180910390f35b34801561075d57600080fd5b50610766611f45565b005b34801561077457600080fd5b506107d3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612003565b005b3480156107e157600080fd5b506108006004803603810190808035906020019092919050505061203b565b005b34801561080e57600080fd5b50610843600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612048565b005b34801561085157600080fd5b50610870600480360381019080803590602001909291905050506120f5565b604051808215151515815260200191505060405180910390f35b34801561089657600080fd5b506108b560048036038101908080359060200190929190505050612167565b6040518082815260200191505060405180910390f35b3480156108d757600080fd5b5061090c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061219f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561095a57600080fd5b50610979600480360381019080803590602001909291905050506121d2565b604051808781526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a1b578082015181840152602081019050610a00565b50505050905090810190601f168015610a485780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b348015610a6757600080fd5b50610a86600480360381019080803590602001909291905050506122c8565b604051808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610afe578082015181840152602081019050610ae3565b50505050905090810190601f168015610b2b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b348015610b4757600080fd5b50610b506123a2565b604051808215151515815260200191505060405180910390f35b348015610b7657600080fd5b50610bcb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506123b5565b005b348015610bd957600080fd5b50610c20600480360381019080803563ffffffff169060200190929190803590602001909291908035906020019082018035906020019190919293919293905050506123c4565b005b348015610c2e57600080fd5b50610c4d60048036038101908080359060200190929190505050612425565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610c9b57600080fd5b50610d12600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff16906020019092919080359060200190929190803590602001908201803590602001919091929391929390803563ffffffff1690602001909291905050506124a3565b005b348015610d2057600080fd5b50610d55600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612639565b6040518082815260200191505060405180910390f35b348015610d7757600080fd5b50610db8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506126bd565b005b348015610dc657600080fd5b50610de560048036038101908080359060200190929190505050612a4c565b005b348015610df357600080fd5b50610e2260048036038101908080359060200190929190803563ffffffff169060200190929190505050612a59565b005b348015610e3057600080fd5b50610e39612a84565b005b348015610e4757600080fd5b50610e6660048036038101908080359060200190929190505050612b44565b604051808481526020018363ffffffff1663ffffffff168152602001828152602001935050505060405180910390f35b348015610ea257600080fd5b50610f07600480360381019080803563ffffffff169060200190929190803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803563ffffffff169060200190929190505050612b9c565b005b348015610f1557600080fd5b50610f1e612c4e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610f6c57600080fd5b50610fa1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612c73565b005b348015610faf57600080fd5b50610fe4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612f7c565b005b348015610ff257600080fd5b50610ffb6131a5565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561103b578082015181840152602081019050611020565b50505050905090810190601f1680156110685780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561108257600080fd5b5061108b613247565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156110d957600080fd5b5061111a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080351515906020019092919050505061326d565b005b34801561112857600080fd5b5061115d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506133a9565b005b34801561116b57600080fd5b506111ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803563ffffffff169060200190929190505050613590565b005b3480156111c857600080fd5b506111d16136e1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561121f57600080fd5b506112c4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050613707565b005b3480156112d257600080fd5b506112f160048036038101908080359060200190929190505050613746565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015611331578082015181840152602081019050611316565b50505050905090810190601f16801561135e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561137857600080fd5b506113976004803603810190808035906020019092919050505061380f565b005b3480156113a557600080fd5b506113da600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061381c565b604051808215151515815260200191505060405180910390f35b34801561140057600080fd5b50611435600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061383c565b604051808215151515815260200191505060405180910390f35b34801561145b57600080fd5b506114b0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061385c565b604051808215151515815260200191505060405180910390f35b3480156114d657600080fd5b5061152d600480360381019080803563ffffffff16906020019092919080359060200190929190803590602001908201803590602001919091929391929390803563ffffffff16906020019092919050505061398a565b005b34801561153b57600080fd5b50611570600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506139fa565b005b34801561157e57600080fd5b506115b3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613b4f565b005b60035481565b600060405180807f737570706f727473496e74657266616365286279746573342900000000000000815250601901905060405180910390207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480611820575060405180807f736166655472616e7366657246726f6d28616464726573732c6164647265737381526020017f2c75696e743235362900000000000000000000000000000000000000000000008152506029019050604051809103902060405180807f617070726f766528616464726573732c75696e743235362900000000000000008152506018019050604051809103902060405180807f6f776e65724f662875696e7432353629000000000000000000000000000000008152506010019050604051809103902060405180807f62616c616e63654f6628616464726573732900000000000000000000000000008152506012019050604051809103902060405180807f746f74616c537570706c79282900000000000000000000000000000000000000815250600d019050604051809103902060405180807f73796d626f6c28290000000000000000000000000000000000000000000000008152506008019050604051809103902060405180807f6e616d6528290000000000000000000000000000000000000000000000000000815250600601905060405180910390201818181818187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b600060149054906101000a900460ff1615151561184357600080fd5b61187f3383838080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613cce565b505050565b600080600060149054906101000a900460ff161515156118a357600080fd5b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515148015611954575060011515601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514155b15156119c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f556e617070726f766564204f424f20616464726573732e00000000000000000081525060200191505060405180910390fd5b611a048986868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613cce565b9150611a448989848a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613f70565b9050611a5189828561430e565b505050505050505050565b606060098054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611af45780601f10611ac957610100808354040283529160200191611af4565b820191906000526020600020905b815481529060010190602001808311611ad757829003601f168201915b5050505050905090565b60006006600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000611b4682612425565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515611b8357600080fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611bc35750611bc2813361385c565b5b1515611bce57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16611bef83611afe565b73ffffffffffffffffffffffffffffffffffffffff16141580611c3f5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b15611cfc57826006600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a35b505050565b601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600d80549050905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff16151515611d7657600080fd5b611d808133614899565b50565b80611d8e3382614a38565b1515611d9957600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614151515611dd557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515611e1157600080fd5b611e1b8483614acd565b611e258483614c36565b611e2f8383614e4e565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b6000611ea583612639565b82101515611eb257600080fd5b600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481101515611efe57fe5b9060005260206000200154905092915050565b60116020528060005260406000206000915090508060000160009054906101000a900463ffffffff16908060010154905082565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611fa057600080fd5b600060149054906101000a900460ff161515611fbb57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b8061200e3382614a38565b151561201957600080fd5b6120358484846020604051908101604052806000815250613707565b50505050565b6120458133614f25565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156120a357600080fd5b601460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff021916905550565b6000806005600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415915050919050565b6000612171611d27565b8210151561217e57600080fd5b600d8281548110151561218d57fe5b90600052602060002001549050919050565b60126020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600080600060606121e4616bb6565b6121ed886150f9565b9050600073ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff1614151515612298576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4469676974616c4d65646961206e6f7420666f756e642e00000000000000000081525060200191505060405180910390fd5b879650806020015195508060400151945080606001519350806080015192508060a0015191505091939550919395565b60008060606122d5616c0f565b6122de85615388565b9050600073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614151515612389576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f436f6c6c656374696f6e206e6f7420666f756e642e000000000000000000000081525060200191505060405180910390fd5b8493508060200151925080604001519150509193909250565b600060149054906101000a900460ff1681565b6123c03383836155db565b5050565b600060149054906101000a900460ff161515156123e057600080fd5b61241e33858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613f70565b5050505050565b6000806005600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561249a57600080fd5b80915050919050565b60008060149054906101000a900460ff161515156124c057600080fd5b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515148015612571575060011515601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514155b15156125e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f556e617070726f766564204f424f20616464726573732e00000000000000000081525060200191505060405180910390fd5b61262387878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613f70565b905061263087828461430e565b50505050505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561267657600080fd5b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515612787576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001807f417070726f76616c20616464726573732069732073616d65206173206170707281526020017f6f7665722e00000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b601460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515612848576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f556e7265636f676e697a6564204f424f20616464726573732e0000000000000081525060200191505060405180910390fd5b60011515601660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151515612911576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f417070726f76616c20616464726573732069732064697361626c65642e00000081525060200191505060405180910390fd5b80601560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f017e8a478826a4348bfb695968246edfab885f8a76b03279cf4630ac073945c9338383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182151515158152602001935050505060405180910390a15050565b612a568133614f25565b50565b600060149054906101000a900460ff16151515612a7557600080fd5b612a8033838361430e565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612adf57600080fd5b600060149054906101000a900460ff16151515612afb57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600080600080612b53856120f5565b1515612b5e57600080fd5b6011600086815260200190815260200160002090508493508060000160009054906101000a900463ffffffff16925080600101549150509193909250565b600080600060149054906101000a900460ff16151515612bbb57600080fd5b612bf73386868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613cce565b9150612c373389848a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613f70565b9050612c4433828561430e565b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612cd057600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515612d80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f5631206d656469612073746f726520616c7265616479207365742e000000000081525060200191505060405180910390fd5b81905060018173ffffffffffffffffffffffffffffffffffffffff16636c669f256040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015612de957600080fd5b505af1158015612dfd573d6000803e3d6000fd5b505050506040513d6020811015612e1357600080fd5b8101908080519060200190929190505050141515612e99576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e636f72726563742076657273696f6e2e000000000000000000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a60800b86040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015612f6057600080fd5b505af1158015612f74573d6000803e3d6000fd5b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612fd757600080fd5b601460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515613098576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f556e7265636f676e697a6564204f424f20616464726573732e0000000000000081525060200191505060405180910390fd5b6001601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690557ffd0e0c743dbdd84ef4e7c513db9b7e085970164787288791343fda28575652dd81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b6060600a8054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561323d5780601f106132125761010080835404028352916020019161323d565b820191906000526020600020905b81548152906001019060200180831161322057829003601f168201915b5050505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156132a857600080fd5b80600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051808215151515815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156133e557600080fd5b601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156134d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f4e6f7420617070726f76656420746f206368616e67652073696e676c6520637281526020017f6561746f722e000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b80601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f384c948063df3740539b4b000658c1a22348e7f18c87f808085662e461e48e7160405160405180910390a350565b600060149054906101000a900460ff161515156135ac57600080fd5b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514801561365d575060011515601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514155b15156136d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f556e617070726f766564204f424f20616464726573732e00000000000000000081525060200191505060405180910390fd5b6136dc83838361430e565b505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b816137123382614a38565b151561371d57600080fd5b613728858585611d83565b61373485858585615a93565b151561373f57600080fd5b5050505050565b6060613751826120f5565b151561375c57600080fd5b600f60008381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156138035780601f106137d857610100808354040283529160200191613803565b820191906000526020600020905b8154815290600101906020018083116137e657829003601f168201915b50505050509050919050565b6138193382614acd565b50565b60146020528060005260406000206000915054906101000a900460ff1681565b60166020528060005260406000206000915054906101000a900460ff1681565b600060011515601660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514156138c05760009050613984565b600115156138ce8385615c81565b151514156138df5760019050613984565b601560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156139775760019050613984565b6139818383615e14565b90505b92915050565b60008060149054906101000a900460ff161515156139a757600080fd5b6139e533878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050613f70565b90506139f233828461430e565b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613a5557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515613a9157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613baa57600080fd5b60011515601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151515613c73576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f416464726573732064697361626c65642e00000000000000000000000000000081525060200191505060405180910390fd5b6001601460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b577fd7385856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015613dae578082015181840152602081019050613d93565b50505050905090810190601f168015613ddb5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015613dfb57600080fd5b505af1158015613e0f573d6000803e3d6000fd5b505050506040513d6020811015613e2557600080fd5b810190808051906020019092919050505090507f01e2312dcdafe7cd3f82579d8c121fdb930d46ef2eb231953a521ac62093e27781600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168686604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015613f29578082015181840152602081019050613f0e565b50505050905090810190601f168015613f565780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a18091505092915050565b600080613f7d8487615ea8565b1515614017576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f43726561746f7220666f7220636f6c6c656374696f6e206e6f7420617070726f81526020017f7665642e0000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166309242ba28760008888886040518663ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018563ffffffff1681526020018463ffffffff1663ffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561411c578082015181840152602081019050614101565b50505050905090810190601f1680156141495780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561416c57600080fd5b505af1158015614180573d6000803e3d6000fd5b505050506040513d602081101561419657600080fd5b810190808051906020019092919050505090507f794c5cd70604d9d8dc2cbca1f8be65f167e4147b6512541d41e8e410594098a081600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16888860008989604051808881526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018563ffffffff1663ffffffff1681526020018463ffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156142c25780820151818401526020810190506142a7565b50505050905090810190601f1680156142ef5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a180915050949350505050565b614316616bb6565b60006060600080614325616c47565b6000808863ffffffff161115156143ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f4661696c6564207072696e742065646974696f6e2e20204372656174696f6e2081526020017f636f756e74206d757374206265203e20302e000000000000000000000000000081525060400191505060405180910390fd5b6127108863ffffffff1610151561446f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f43616e6e6f74207072696e74206d6f7265207468616e2031304b20746f6b656e81526020017f73206174206f6e6365000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b614478896150f9565b96508660400151955061448f87608001518b615ee7565b1515614503576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f43726561746f72206e6f7420617070726f7665642e000000000000000000000081525060200191505060405180910390fd5b61450c8a615ff1565b15156145a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001807f43726561746f72206d757374206d617463682073696e676c652063726561746f81526020017f7220616464726573732e0000000000000000000000000000000000000000000081525060400191505060405180910390fd5b866020015163ffffffff1686890163ffffffff1611151515614630576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f546f74616c20737570706c792065786365656465642e0000000000000000000081525060200191505060405180910390fd5b6146736040805190810160405280600c81526020017f697066733a2f2f697066732f00000000000000000000000000000000000000008152508860a0015161616e565b9450600093505b8763ffffffff168463ffffffff16101561483657836001870101925060408051908101604052808463ffffffff1681526020018a81525091506146bb61633f565b9050816011600083815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff160217905550602082015181600101559050507f775f53e4c75ce0c74e611f7f0bb660e4cd647e0522ef0f8aefd4ecef373c5df9818b85888d604051808681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018463ffffffff1663ffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156147bb5780820151818401526020810190506147a0565b50505050905090810190601f1680156147e85780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a16148038a8261635c565b61480d81866163b3565b61482360016013546163f390919063ffffffff16565b601381905550838060010194505061467a565b614840878961640f565b7f12d99f5e49ef761c52953e4f9a109827fc3540292ba88c10d309fef47068525989898801604051808381526020018263ffffffff1663ffffffff1681526020019250505060405180910390a150505050505050505050565b6148a1616bb6565b6000806148ad856150f9565b92506148bd836080015185615ee7565b806148d257506148d183608001518561385c565b5b151561496c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001807f4661696c6564206469676974616c206d65646961206275726e2e202043616c6c81526020017f6572206e6f7420617070726f7665642e0000000000000000000000000000000081525060400191505060405180910390fd5b82604001518360200151039150614983838361640f565b614990836000015161657f565b90507f327ecc068f1b41267f69376098f6a50da487e4a4d762d53c01197d6a2f294b3e858583604051808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a15050505050565b600080614a4483612425565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480614ab357508373ffffffffffffffffffffffffffffffffffffffff16614a9b84611afe565b73ffffffffffffffffffffffffffffffffffffffff16145b80614ac45750614ac3818561385c565b5b91505092915050565b8173ffffffffffffffffffffffffffffffffffffffff16614aed82612425565b73ffffffffffffffffffffffffffffffffffffffff16141515614b0f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166006600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515614c325760006006600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a35b5050565b6000806000614c45858561669b565b600c6000858152602001908152602001600020549250614cb16001600b60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490506167ca90919063ffffffff16565b9150600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481101515614cff57fe5b9060005260206000200154905080600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084815481101515614d5957fe5b90600052602060002001819055506000600b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083815481101515614db557fe5b9060005260206000200181905550600b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480919060019003614e159190616c67565b506000600c60008681526020019081526020016000208190555082600c6000838152602001908152602001600020819055505050505050565b6000614e5a83836167e3565b600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082908060018154018082558091505090600182039060005260206000200160009091929091909150555080600c600084815260200190815260200160002081905550505050565b6000614f3083612425565b90508073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480614f9f57508173ffffffffffffffffffffffffffffffffffffffff16614f8784611afe565b73ffffffffffffffffffffffffffffffffffffffff16145b80614fb05750614faf818361385c565b5b151561504a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b8152602001807f4661696c656420746f6b656e206275726e2e202043616c6c6572206973206e6f81526020017f7420617070726f7665642e00000000000000000000000000000000000000000081525060400191505060405180910390fd5b615054818461693d565b60116000848152602001908152602001600020600080820160006101000a81549063ffffffff0219169055600182016000905550507f1e8df141f42ed659a8fe7e7c5966cbdf2d240d0c45f4c30cbe02526c618075ef8382604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505050565b615101616bb6565b600080600080600060606000615115616bb6565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561517357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156151d157600080fd5b6151da8a61657f565b91508173ffffffffffffffffffffffffffffffffffffffff166355df42758b6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561524b57600080fd5b505af115801561525f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060c081101561528957600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080516401000000008111156152d357600080fd5b828101905060208101848111156152e957600080fd5b815185600182028301116401000000008211171561530657600080fd5b5050929190505050809850819950829a50839b50849c50859d5050505050505060c0604051908101604052808981526020018863ffffffff1681526020018763ffffffff1681526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481525090508098505050505050505050919050565b615390616c0f565b600080606061539d616c0f565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156153fb57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561545957600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635a1f3c28876040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b1580156154ea57600080fd5b505af11580156154fe573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250606081101561552857600080fd5b810190808051906020019092919080519060200190929190805164010000000081111561555457600080fd5b8281019050602081018481111561556a57600080fd5b815185600182028301116401000000008211171561558757600080fd5b50509291905050508094508195508296505050506060604051908101604052808581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815250905080945050505050919050565b6000601260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141580156156a95750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b1515615743576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f43726561746f72206d7573742062652076616c6964206e6f6e2030783020616481526020017f64726573732e000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806157a857508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b151561581c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f556e617574686f72697a65642063616c6c65722e00000000000000000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156158d45781601260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506159f6565b8073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141515615977576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f556e617574686f72697a65642063616c6c65722e00000000000000000000000081525060200191505060405180910390fd5b81601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b7fde6cfdf21fe76bcb45258138e27bcd332b76941b24d226b5da8dc5f9cd531c3e8383604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a150505050565b600080615ab58573ffffffffffffffffffffffffffffffffffffffff16616a75565b1515615ac45760019150615c78565b8473ffffffffffffffffffffffffffffffffffffffff1663f0b9e5ba8786866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015615b86578082015181840152602081019050615b6b565b50505050905090810190601f168015615bb35780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015615bd457600080fd5b505af1158015615be8573d6000803e3d6000fd5b505050506040513d6020811015615bfe57600080fd5b8101908080519060200190929190505050905063f0b9e5ba7c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161491505b50949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515615ce057600080fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166367d6a7dc84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015615dd157600080fd5b505af1158015615de5573d6000803e3d6000fd5b505050506040513d6020811015615dfb57600080fd5b8101908080519060200190929190505050905092915050565b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000615eb2616c0f565b6000841415615ec45760019150615ee0565b615ecd84615388565b9050615edd816020015184615ee7565b91505b5092915050565b600080601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515615fb8578273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16149150615fea565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161491505b5092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156160bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f3078302063726561746f722061646472657373657320617265206e6f7420616c81526020017f6c6f7765642e000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061616757508173ffffffffffffffffffffffffffffffffffffffff16601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b9050919050565b606080606080606060008088955087945084518651016040519080825280601f01601f1916602001820160405280156161b65781602001602082028038833980820191505090505b50935083925060009150600090505b85518110156162785785818151811015156161dc57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561623b57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506161c5565b600090505b845181101561633057848181518110151561629457fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156162f357fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808060010191505061627d565b82965050505050505092915050565b600061635760016013546163f390919063ffffffff16565b905090565b6163668282616a88565b600d80549050600e600083815260200190815260200160002081905550600d8190806001815401808255809150509060018203906000526020600020016000909192909190915055505050565b6163bc826120f5565b15156163c757600080fd5b80600f600084815260200190815260200160002090805190602001906163ee929190616c93565b505050565b6000818301905082811015151561640657fe5b80905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561646e57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156164cc57600080fd5b6164d9836000015161657f565b90508073ffffffffffffffffffffffffffffffffffffffff16635160a1248460000151846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018263ffffffff1663ffffffff16815260200192505050600060405180830381600087803b15801561656257600080fd5b505af1158015616576573d6000803e3d6000fd5b50505050505050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156165de57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561663c57600080fd5b60035482101561667057600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050616696565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b919050565b8173ffffffffffffffffffffffffffffffffffffffff166166bb82612425565b73ffffffffffffffffffffffffffffffffffffffff161415156166dd57600080fd5b6167306001600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546167ca90919063ffffffff16565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060006005600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008282111515156167d857fe5b818303905092915050565b600073ffffffffffffffffffffffffffffffffffffffff166005600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561685157600080fd5b816005600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506168f66001600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546163f390919063ffffffff16565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b600080600061694c8585616b38565b6000600f600086815260200190815260200160002080546001816001161561010002031660029004905014151561699d57600f6000858152602001908152602001600020600061699c9190616d13565b5b600e60008581526020019081526020016000205492506169cc6001600d805490506167ca90919063ffffffff16565b9150600d828154811015156169dd57fe5b9060005260206000200154905080600d848154811015156169fa57fe5b90600052602060002001819055506000600d83815481101515616a1957fe5b9060005260206000200181905550600d805480919060019003616a3c9190616c67565b506000600e60008681526020019081526020016000208190555082600e6000838152602001908152602001600020819055505050505050565b600080823b905060008111915050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515616ac457600080fd5b616ace8282614e4e565b8173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b616b428282614acd565b616b4c8282614c36565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60c06040519081016040528060008152602001600063ffffffff168152602001600063ffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b60606040519081016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6040805190810160405280600063ffffffff168152602001600081525090565b815481835581811115616c8e57818360005260206000209182019101616c8d9190616d5b565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10616cd457805160ff1916838001178555616d02565b82800160010185558215616d02579182015b82811115616d01578251825591602001919060010190616ce6565b5b509050616d0f9190616d5b565b5090565b50805460018160011615610100020316600290046000825580601f10616d395750616d58565b601f016020900490600052602060002090810190616d579190616d5b565b5b50565b616d7d91905b80821115616d79576000816000905550600101616d61565b5090565b905600a165627a7a72305820a73734b6ffba0ffe09b7bc4426e9026c7bf29eeae994f42cbf274a5445d57b990029"

// DeployDigitalMediaCore deploys a new Ethereum contract, binding an instance of DigitalMediaCore to it.
func DeployDigitalMediaCore(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenName string, _tokenSymbol string, _tokenIdStartingCounter *big.Int, _dmsAddress common.Address, _crsAddress common.Address) (common.Address, *types.Transaction, *DigitalMediaCore, error) {
	parsed, err := abi.JSON(strings.NewReader(DigitalMediaCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DigitalMediaCoreBin), backend, _tokenName, _tokenSymbol, _tokenIdStartingCounter, _dmsAddress, _crsAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DigitalMediaCore{DigitalMediaCoreCaller: DigitalMediaCoreCaller{contract: contract}, DigitalMediaCoreTransactor: DigitalMediaCoreTransactor{contract: contract}, DigitalMediaCoreFilterer: DigitalMediaCoreFilterer{contract: contract}}, nil
}

// DigitalMediaCore is an auto generated Go binding around an Ethereum contract.
type DigitalMediaCore struct {
	DigitalMediaCoreCaller     // Read-only binding to the contract
	DigitalMediaCoreTransactor // Write-only binding to the contract
	DigitalMediaCoreFilterer   // Log filterer for contract events
}

// DigitalMediaCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type DigitalMediaCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalMediaCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DigitalMediaCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalMediaCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DigitalMediaCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DigitalMediaCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DigitalMediaCoreSession struct {
	Contract     *DigitalMediaCore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DigitalMediaCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DigitalMediaCoreCallerSession struct {
	Contract *DigitalMediaCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DigitalMediaCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DigitalMediaCoreTransactorSession struct {
	Contract     *DigitalMediaCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DigitalMediaCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type DigitalMediaCoreRaw struct {
	Contract *DigitalMediaCore // Generic contract binding to access the raw methods on
}

// DigitalMediaCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DigitalMediaCoreCallerRaw struct {
	Contract *DigitalMediaCoreCaller // Generic read-only contract binding to access the raw methods on
}

// DigitalMediaCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DigitalMediaCoreTransactorRaw struct {
	Contract *DigitalMediaCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDigitalMediaCore creates a new instance of DigitalMediaCore, bound to a specific deployed contract.
func NewDigitalMediaCore(address common.Address, backend bind.ContractBackend) (*DigitalMediaCore, error) {
	contract, err := bindDigitalMediaCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCore{DigitalMediaCoreCaller: DigitalMediaCoreCaller{contract: contract}, DigitalMediaCoreTransactor: DigitalMediaCoreTransactor{contract: contract}, DigitalMediaCoreFilterer: DigitalMediaCoreFilterer{contract: contract}}, nil
}

// NewDigitalMediaCoreCaller creates a new read-only instance of DigitalMediaCore, bound to a specific deployed contract.
func NewDigitalMediaCoreCaller(address common.Address, caller bind.ContractCaller) (*DigitalMediaCoreCaller, error) {
	contract, err := bindDigitalMediaCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreCaller{contract: contract}, nil
}

// NewDigitalMediaCoreTransactor creates a new write-only instance of DigitalMediaCore, bound to a specific deployed contract.
func NewDigitalMediaCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*DigitalMediaCoreTransactor, error) {
	contract, err := bindDigitalMediaCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreTransactor{contract: contract}, nil
}

// NewDigitalMediaCoreFilterer creates a new log filterer instance of DigitalMediaCore, bound to a specific deployed contract.
func NewDigitalMediaCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*DigitalMediaCoreFilterer, error) {
	contract, err := bindDigitalMediaCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreFilterer{contract: contract}, nil
}

// bindDigitalMediaCore binds a generic wrapper to an already deployed contract.
func bindDigitalMediaCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DigitalMediaCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigitalMediaCore *DigitalMediaCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DigitalMediaCore.Contract.DigitalMediaCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigitalMediaCore *DigitalMediaCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.DigitalMediaCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigitalMediaCore *DigitalMediaCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.DigitalMediaCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DigitalMediaCore *DigitalMediaCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DigitalMediaCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DigitalMediaCore *DigitalMediaCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DigitalMediaCore *DigitalMediaCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.contract.Transact(opts, method, params...)
}

// ApprovedCreators is a free data retrieval call binding the contract method 0x55232467.
//
// Solidity: function approvedCreators(address ) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) ApprovedCreators(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "approvedCreators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovedCreators is a free data retrieval call binding the contract method 0x55232467.
//
// Solidity: function approvedCreators(address ) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) ApprovedCreators(arg0 common.Address) (common.Address, error) {
	return _DigitalMediaCore.Contract.ApprovedCreators(&_DigitalMediaCore.CallOpts, arg0)
}

// ApprovedCreators is a free data retrieval call binding the contract method 0x55232467.
//
// Solidity: function approvedCreators(address ) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) ApprovedCreators(arg0 common.Address) (common.Address, error) {
	return _DigitalMediaCore.Contract.ApprovedCreators(&_DigitalMediaCore.CallOpts, arg0)
}

// ApprovedTokenCreators is a free data retrieval call binding the contract method 0xd70f575f.
//
// Solidity: function approvedTokenCreators(address ) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCaller) ApprovedTokenCreators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "approvedTokenCreators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTokenCreators is a free data retrieval call binding the contract method 0xd70f575f.
//
// Solidity: function approvedTokenCreators(address ) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreSession) ApprovedTokenCreators(arg0 common.Address) (bool, error) {
	return _DigitalMediaCore.Contract.ApprovedTokenCreators(&_DigitalMediaCore.CallOpts, arg0)
}

// ApprovedTokenCreators is a free data retrieval call binding the contract method 0xd70f575f.
//
// Solidity: function approvedTokenCreators(address ) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) ApprovedTokenCreators(arg0 common.Address) (bool, error) {
	return _DigitalMediaCore.Contract.ApprovedTokenCreators(&_DigitalMediaCore.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DigitalMediaCore.Contract.BalanceOf(&_DigitalMediaCore.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DigitalMediaCore.Contract.BalanceOf(&_DigitalMediaCore.CallOpts, _owner)
}

// CreatorRegistryStore is a free data retrieval call binding the contract method 0xb39ba60a.
//
// Solidity: function creatorRegistryStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) CreatorRegistryStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "creatorRegistryStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorRegistryStore is a free data retrieval call binding the contract method 0xb39ba60a.
//
// Solidity: function creatorRegistryStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) CreatorRegistryStore() (common.Address, error) {
	return _DigitalMediaCore.Contract.CreatorRegistryStore(&_DigitalMediaCore.CallOpts)
}

// CreatorRegistryStore is a free data retrieval call binding the contract method 0xb39ba60a.
//
// Solidity: function creatorRegistryStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) CreatorRegistryStore() (common.Address, error) {
	return _DigitalMediaCore.Contract.CreatorRegistryStore(&_DigitalMediaCore.CallOpts)
}

// CurrentDigitalMediaStore is a free data retrieval call binding the contract method 0x1b284e75.
//
// Solidity: function currentDigitalMediaStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) CurrentDigitalMediaStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "currentDigitalMediaStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentDigitalMediaStore is a free data retrieval call binding the contract method 0x1b284e75.
//
// Solidity: function currentDigitalMediaStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) CurrentDigitalMediaStore() (common.Address, error) {
	return _DigitalMediaCore.Contract.CurrentDigitalMediaStore(&_DigitalMediaCore.CallOpts)
}

// CurrentDigitalMediaStore is a free data retrieval call binding the contract method 0x1b284e75.
//
// Solidity: function currentDigitalMediaStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) CurrentDigitalMediaStore() (common.Address, error) {
	return _DigitalMediaCore.Contract.CurrentDigitalMediaStore(&_DigitalMediaCore.CallOpts)
}

// CurrentStartingDigitalMediaId is a free data retrieval call binding the contract method 0x0154788d.
//
// Solidity: function currentStartingDigitalMediaId() view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCaller) CurrentStartingDigitalMediaId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "currentStartingDigitalMediaId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentStartingDigitalMediaId is a free data retrieval call binding the contract method 0x0154788d.
//
// Solidity: function currentStartingDigitalMediaId() view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreSession) CurrentStartingDigitalMediaId() (*big.Int, error) {
	return _DigitalMediaCore.Contract.CurrentStartingDigitalMediaId(&_DigitalMediaCore.CallOpts)
}

// CurrentStartingDigitalMediaId is a free data retrieval call binding the contract method 0x0154788d.
//
// Solidity: function currentStartingDigitalMediaId() view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) CurrentStartingDigitalMediaId() (*big.Int, error) {
	return _DigitalMediaCore.Contract.CurrentStartingDigitalMediaId(&_DigitalMediaCore.CallOpts)
}

// DisabledOboOperators is a free data retrieval call binding the contract method 0xe158386d.
//
// Solidity: function disabledOboOperators(address ) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCaller) DisabledOboOperators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "disabledOboOperators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisabledOboOperators is a free data retrieval call binding the contract method 0xe158386d.
//
// Solidity: function disabledOboOperators(address ) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreSession) DisabledOboOperators(arg0 common.Address) (bool, error) {
	return _DigitalMediaCore.Contract.DisabledOboOperators(&_DigitalMediaCore.CallOpts, arg0)
}

// DisabledOboOperators is a free data retrieval call binding the contract method 0xe158386d.
//
// Solidity: function disabledOboOperators(address ) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) DisabledOboOperators(arg0 common.Address) (bool, error) {
	return _DigitalMediaCore.Contract.DisabledOboOperators(&_DigitalMediaCore.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "exists", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreSession) Exists(_tokenId *big.Int) (bool, error) {
	return _DigitalMediaCore.Contract.Exists(&_DigitalMediaCore.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _DigitalMediaCore.Contract.Exists(&_DigitalMediaCore.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _DigitalMediaCore.Contract.GetApproved(&_DigitalMediaCore.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _DigitalMediaCore.Contract.GetApproved(&_DigitalMediaCore.CallOpts, _tokenId)
}

// GetCollection is a free data retrieval call binding the contract method 0x5a1f3c28.
//
// Solidity: function getCollection(uint256 _id) view returns(uint256 id, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreCaller) GetCollection(opts *bind.CallOpts, _id *big.Int) (struct {
	Id           *big.Int
	Creator      common.Address
	MetadataPath string
}, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "getCollection", _id)

	outstruct := new(struct {
		Id           *big.Int
		Creator      common.Address
		MetadataPath string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.MetadataPath = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetCollection is a free data retrieval call binding the contract method 0x5a1f3c28.
//
// Solidity: function getCollection(uint256 _id) view returns(uint256 id, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreSession) GetCollection(_id *big.Int) (struct {
	Id           *big.Int
	Creator      common.Address
	MetadataPath string
}, error) {
	return _DigitalMediaCore.Contract.GetCollection(&_DigitalMediaCore.CallOpts, _id)
}

// GetCollection is a free data retrieval call binding the contract method 0x5a1f3c28.
//
// Solidity: function getCollection(uint256 _id) view returns(uint256 id, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) GetCollection(_id *big.Int) (struct {
	Id           *big.Int
	Creator      common.Address
	MetadataPath string
}, error) {
	return _DigitalMediaCore.Contract.GetCollection(&_DigitalMediaCore.CallOpts, _id)
}

// GetDigitalMedia is a free data retrieval call binding the contract method 0x55df4275.
//
// Solidity: function getDigitalMedia(uint256 _id) view returns(uint256 id, uint32 totalSupply, uint32 printIndex, uint256 collectionId, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreCaller) GetDigitalMedia(opts *bind.CallOpts, _id *big.Int) (struct {
	Id           *big.Int
	TotalSupply  uint32
	PrintIndex   uint32
	CollectionId *big.Int
	Creator      common.Address
	MetadataPath string
}, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "getDigitalMedia", _id)

	outstruct := new(struct {
		Id           *big.Int
		TotalSupply  uint32
		PrintIndex   uint32
		CollectionId *big.Int
		Creator      common.Address
		MetadataPath string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PrintIndex = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.CollectionId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.MetadataPath = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// GetDigitalMedia is a free data retrieval call binding the contract method 0x55df4275.
//
// Solidity: function getDigitalMedia(uint256 _id) view returns(uint256 id, uint32 totalSupply, uint32 printIndex, uint256 collectionId, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreSession) GetDigitalMedia(_id *big.Int) (struct {
	Id           *big.Int
	TotalSupply  uint32
	PrintIndex   uint32
	CollectionId *big.Int
	Creator      common.Address
	MetadataPath string
}, error) {
	return _DigitalMediaCore.Contract.GetDigitalMedia(&_DigitalMediaCore.CallOpts, _id)
}

// GetDigitalMedia is a free data retrieval call binding the contract method 0x55df4275.
//
// Solidity: function getDigitalMedia(uint256 _id) view returns(uint256 id, uint32 totalSupply, uint32 printIndex, uint256 collectionId, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) GetDigitalMedia(_id *big.Int) (struct {
	Id           *big.Int
	TotalSupply  uint32
	PrintIndex   uint32
	CollectionId *big.Int
	Creator      common.Address
	MetadataPath string
}, error) {
	return _DigitalMediaCore.Contract.GetDigitalMedia(&_DigitalMediaCore.CallOpts, _id)
}

// GetDigitalMediaRelease is a free data retrieval call binding the contract method 0x8a603bdf.
//
// Solidity: function getDigitalMediaRelease(uint256 _id) view returns(uint256 id, uint32 printEdition, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreCaller) GetDigitalMediaRelease(opts *bind.CallOpts, _id *big.Int) (struct {
	Id             *big.Int
	PrintEdition   uint32
	DigitalMediaId *big.Int
}, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "getDigitalMediaRelease", _id)

	outstruct := new(struct {
		Id             *big.Int
		PrintEdition   uint32
		DigitalMediaId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrintEdition = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DigitalMediaId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDigitalMediaRelease is a free data retrieval call binding the contract method 0x8a603bdf.
//
// Solidity: function getDigitalMediaRelease(uint256 _id) view returns(uint256 id, uint32 printEdition, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreSession) GetDigitalMediaRelease(_id *big.Int) (struct {
	Id             *big.Int
	PrintEdition   uint32
	DigitalMediaId *big.Int
}, error) {
	return _DigitalMediaCore.Contract.GetDigitalMediaRelease(&_DigitalMediaCore.CallOpts, _id)
}

// GetDigitalMediaRelease is a free data retrieval call binding the contract method 0x8a603bdf.
//
// Solidity: function getDigitalMediaRelease(uint256 _id) view returns(uint256 id, uint32 printEdition, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) GetDigitalMediaRelease(_id *big.Int) (struct {
	Id             *big.Int
	PrintEdition   uint32
	DigitalMediaId *big.Int
}, error) {
	return _DigitalMediaCore.Contract.GetDigitalMediaRelease(&_DigitalMediaCore.CallOpts, _id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _DigitalMediaCore.Contract.IsApprovedForAll(&_DigitalMediaCore.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _DigitalMediaCore.Contract.IsApprovedForAll(&_DigitalMediaCore.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreSession) Name() (string, error) {
	return _DigitalMediaCore.Contract.Name(&_DigitalMediaCore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) Name() (string, error) {
	return _DigitalMediaCore.Contract.Name(&_DigitalMediaCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) Owner() (common.Address, error) {
	return _DigitalMediaCore.Contract.Owner(&_DigitalMediaCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) Owner() (common.Address, error) {
	return _DigitalMediaCore.Contract.Owner(&_DigitalMediaCore.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _DigitalMediaCore.Contract.OwnerOf(&_DigitalMediaCore.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _DigitalMediaCore.Contract.OwnerOf(&_DigitalMediaCore.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreSession) Paused() (bool, error) {
	return _DigitalMediaCore.Contract.Paused(&_DigitalMediaCore.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) Paused() (bool, error) {
	return _DigitalMediaCore.Contract.Paused(&_DigitalMediaCore.CallOpts)
}

// SingleCreatorAddress is a free data retrieval call binding the contract method 0x147ca2af.
//
// Solidity: function singleCreatorAddress() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) SingleCreatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "singleCreatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SingleCreatorAddress is a free data retrieval call binding the contract method 0x147ca2af.
//
// Solidity: function singleCreatorAddress() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) SingleCreatorAddress() (common.Address, error) {
	return _DigitalMediaCore.Contract.SingleCreatorAddress(&_DigitalMediaCore.CallOpts)
}

// SingleCreatorAddress is a free data retrieval call binding the contract method 0x147ca2af.
//
// Solidity: function singleCreatorAddress() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) SingleCreatorAddress() (common.Address, error) {
	return _DigitalMediaCore.Contract.SingleCreatorAddress(&_DigitalMediaCore.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _DigitalMediaCore.Contract.SupportsInterface(&_DigitalMediaCore.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _DigitalMediaCore.Contract.SupportsInterface(&_DigitalMediaCore.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreSession) Symbol() (string, error) {
	return _DigitalMediaCore.Contract.Symbol(&_DigitalMediaCore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) Symbol() (string, error) {
	return _DigitalMediaCore.Contract.Symbol(&_DigitalMediaCore.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "tokenByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _DigitalMediaCore.Contract.TokenByIndex(&_DigitalMediaCore.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _DigitalMediaCore.Contract.TokenByIndex(&_DigitalMediaCore.CallOpts, _index)
}

// TokenIdToDigitalMediaRelease is a free data retrieval call binding the contract method 0x3db57cbe.
//
// Solidity: function tokenIdToDigitalMediaRelease(uint256 ) view returns(uint32 printEdition, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreCaller) TokenIdToDigitalMediaRelease(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PrintEdition   uint32
	DigitalMediaId *big.Int
}, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "tokenIdToDigitalMediaRelease", arg0)

	outstruct := new(struct {
		PrintEdition   uint32
		DigitalMediaId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrintEdition = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DigitalMediaId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenIdToDigitalMediaRelease is a free data retrieval call binding the contract method 0x3db57cbe.
//
// Solidity: function tokenIdToDigitalMediaRelease(uint256 ) view returns(uint32 printEdition, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreSession) TokenIdToDigitalMediaRelease(arg0 *big.Int) (struct {
	PrintEdition   uint32
	DigitalMediaId *big.Int
}, error) {
	return _DigitalMediaCore.Contract.TokenIdToDigitalMediaRelease(&_DigitalMediaCore.CallOpts, arg0)
}

// TokenIdToDigitalMediaRelease is a free data retrieval call binding the contract method 0x3db57cbe.
//
// Solidity: function tokenIdToDigitalMediaRelease(uint256 ) view returns(uint32 printEdition, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) TokenIdToDigitalMediaRelease(arg0 *big.Int) (struct {
	PrintEdition   uint32
	DigitalMediaId *big.Int
}, error) {
	return _DigitalMediaCore.Contract.TokenIdToDigitalMediaRelease(&_DigitalMediaCore.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _DigitalMediaCore.Contract.TokenOfOwnerByIndex(&_DigitalMediaCore.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _DigitalMediaCore.Contract.TokenOfOwnerByIndex(&_DigitalMediaCore.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DigitalMediaCore.Contract.TokenURI(&_DigitalMediaCore.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DigitalMediaCore.Contract.TokenURI(&_DigitalMediaCore.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreSession) TotalSupply() (*big.Int, error) {
	return _DigitalMediaCore.Contract.TotalSupply(&_DigitalMediaCore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) TotalSupply() (*big.Int, error) {
	return _DigitalMediaCore.Contract.TotalSupply(&_DigitalMediaCore.CallOpts)
}

// V1DigitalMediaStore is a free data retrieval call binding the contract method 0xa0f01e08.
//
// Solidity: function v1DigitalMediaStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCaller) V1DigitalMediaStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DigitalMediaCore.contract.Call(opts, &out, "v1DigitalMediaStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// V1DigitalMediaStore is a free data retrieval call binding the contract method 0xa0f01e08.
//
// Solidity: function v1DigitalMediaStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreSession) V1DigitalMediaStore() (common.Address, error) {
	return _DigitalMediaCore.Contract.V1DigitalMediaStore(&_DigitalMediaCore.CallOpts)
}

// V1DigitalMediaStore is a free data retrieval call binding the contract method 0xa0f01e08.
//
// Solidity: function v1DigitalMediaStore() view returns(address)
func (_DigitalMediaCore *DigitalMediaCoreCallerSession) V1DigitalMediaStore() (common.Address, error) {
	return _DigitalMediaCore.Contract.V1DigitalMediaStore(&_DigitalMediaCore.CallOpts)
}

// AddApprovedTokenCreator is a paid mutator transaction binding the contract method 0xfafa8a4b.
//
// Solidity: function addApprovedTokenCreator(address _creatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) AddApprovedTokenCreator(opts *bind.TransactOpts, _creatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "addApprovedTokenCreator", _creatorAddress)
}

// AddApprovedTokenCreator is a paid mutator transaction binding the contract method 0xfafa8a4b.
//
// Solidity: function addApprovedTokenCreator(address _creatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) AddApprovedTokenCreator(_creatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.AddApprovedTokenCreator(&_DigitalMediaCore.TransactOpts, _creatorAddress)
}

// AddApprovedTokenCreator is a paid mutator transaction binding the contract method 0xfafa8a4b.
//
// Solidity: function addApprovedTokenCreator(address _creatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) AddApprovedTokenCreator(_creatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.AddApprovedTokenCreator(&_DigitalMediaCore.TransactOpts, _creatorAddress)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Approve(&_DigitalMediaCore.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Approve(&_DigitalMediaCore.TransactOpts, _to, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Burn(&_DigitalMediaCore.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Burn(&_DigitalMediaCore.TransactOpts, tokenId)
}

// BurnDigitalMedia is a paid mutator transaction binding the contract method 0x23077f58.
//
// Solidity: function burnDigitalMedia(uint256 _digitalMediaId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) BurnDigitalMedia(opts *bind.TransactOpts, _digitalMediaId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "burnDigitalMedia", _digitalMediaId)
}

// BurnDigitalMedia is a paid mutator transaction binding the contract method 0x23077f58.
//
// Solidity: function burnDigitalMedia(uint256 _digitalMediaId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) BurnDigitalMedia(_digitalMediaId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.BurnDigitalMedia(&_DigitalMediaCore.TransactOpts, _digitalMediaId)
}

// BurnDigitalMedia is a paid mutator transaction binding the contract method 0x23077f58.
//
// Solidity: function burnDigitalMedia(uint256 _digitalMediaId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) BurnDigitalMedia(_digitalMediaId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.BurnDigitalMedia(&_DigitalMediaCore.TransactOpts, _digitalMediaId)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) BurnToken(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "burnToken", _tokenId)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) BurnToken(_tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.BurnToken(&_DigitalMediaCore.TransactOpts, _tokenId)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) BurnToken(_tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.BurnToken(&_DigitalMediaCore.TransactOpts, _tokenId)
}

// ChangeCreator is a paid mutator transaction binding the contract method 0x5cf09fee.
//
// Solidity: function changeCreator(address _creator, address _newCreator) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) ChangeCreator(opts *bind.TransactOpts, _creator common.Address, _newCreator common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "changeCreator", _creator, _newCreator)
}

// ChangeCreator is a paid mutator transaction binding the contract method 0x5cf09fee.
//
// Solidity: function changeCreator(address _creator, address _newCreator) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) ChangeCreator(_creator common.Address, _newCreator common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.ChangeCreator(&_DigitalMediaCore.TransactOpts, _creator, _newCreator)
}

// ChangeCreator is a paid mutator transaction binding the contract method 0x5cf09fee.
//
// Solidity: function changeCreator(address _creator, address _newCreator) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) ChangeCreator(_creator common.Address, _newCreator common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.ChangeCreator(&_DigitalMediaCore.TransactOpts, _creator, _newCreator)
}

// ChangeSingleCreator is a paid mutator transaction binding the contract method 0xa7df572c.
//
// Solidity: function changeSingleCreator(address _newCreatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) ChangeSingleCreator(opts *bind.TransactOpts, _newCreatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "changeSingleCreator", _newCreatorAddress)
}

// ChangeSingleCreator is a paid mutator transaction binding the contract method 0xa7df572c.
//
// Solidity: function changeSingleCreator(address _newCreatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) ChangeSingleCreator(_newCreatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.ChangeSingleCreator(&_DigitalMediaCore.TransactOpts, _newCreatorAddress)
}

// ChangeSingleCreator is a paid mutator transaction binding the contract method 0xa7df572c.
//
// Solidity: function changeSingleCreator(address _newCreatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) ChangeSingleCreator(_newCreatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.ChangeSingleCreator(&_DigitalMediaCore.TransactOpts, _newCreatorAddress)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x059dfe13.
//
// Solidity: function createCollection(string _metadataPath) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) CreateCollection(opts *bind.TransactOpts, _metadataPath string) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "createCollection", _metadataPath)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x059dfe13.
//
// Solidity: function createCollection(string _metadataPath) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) CreateCollection(_metadataPath string) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateCollection(&_DigitalMediaCore.TransactOpts, _metadataPath)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x059dfe13.
//
// Solidity: function createCollection(string _metadataPath) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) CreateCollection(_metadataPath string) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateCollection(&_DigitalMediaCore.TransactOpts, _metadataPath)
}

// CreateDigitalMedia is a paid mutator transaction binding the contract method 0x603417fb.
//
// Solidity: function createDigitalMedia(uint32 _totalSupply, uint256 _collectionId, string _metadataPath) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) CreateDigitalMedia(opts *bind.TransactOpts, _totalSupply uint32, _collectionId *big.Int, _metadataPath string) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "createDigitalMedia", _totalSupply, _collectionId, _metadataPath)
}

// CreateDigitalMedia is a paid mutator transaction binding the contract method 0x603417fb.
//
// Solidity: function createDigitalMedia(uint32 _totalSupply, uint256 _collectionId, string _metadataPath) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) CreateDigitalMedia(_totalSupply uint32, _collectionId *big.Int, _metadataPath string) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMedia(&_DigitalMediaCore.TransactOpts, _totalSupply, _collectionId, _metadataPath)
}

// CreateDigitalMedia is a paid mutator transaction binding the contract method 0x603417fb.
//
// Solidity: function createDigitalMedia(uint32 _totalSupply, uint256 _collectionId, string _metadataPath) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) CreateDigitalMedia(_totalSupply uint32, _collectionId *big.Int, _metadataPath string) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMedia(&_DigitalMediaCore.TransactOpts, _totalSupply, _collectionId, _metadataPath)
}

// CreateDigitalMediaAndReleases is a paid mutator transaction binding the contract method 0xeca211e3.
//
// Solidity: function createDigitalMediaAndReleases(uint32 _totalSupply, uint256 _collectionId, string _metadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) CreateDigitalMediaAndReleases(opts *bind.TransactOpts, _totalSupply uint32, _collectionId *big.Int, _metadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "createDigitalMediaAndReleases", _totalSupply, _collectionId, _metadataPath, _numReleases)
}

// CreateDigitalMediaAndReleases is a paid mutator transaction binding the contract method 0xeca211e3.
//
// Solidity: function createDigitalMediaAndReleases(uint32 _totalSupply, uint256 _collectionId, string _metadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) CreateDigitalMediaAndReleases(_totalSupply uint32, _collectionId *big.Int, _metadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMediaAndReleases(&_DigitalMediaCore.TransactOpts, _totalSupply, _collectionId, _metadataPath, _numReleases)
}

// CreateDigitalMediaAndReleases is a paid mutator transaction binding the contract method 0xeca211e3.
//
// Solidity: function createDigitalMediaAndReleases(uint32 _totalSupply, uint256 _collectionId, string _metadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) CreateDigitalMediaAndReleases(_totalSupply uint32, _collectionId *big.Int, _metadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMediaAndReleases(&_DigitalMediaCore.TransactOpts, _totalSupply, _collectionId, _metadataPath, _numReleases)
}

// CreateDigitalMediaAndReleasesInNewCollection is a paid mutator transaction binding the contract method 0x8b40e8aa.
//
// Solidity: function createDigitalMediaAndReleasesInNewCollection(uint32 _totalSupply, string _digitalMediaMetadataPath, string _collectionMetadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) CreateDigitalMediaAndReleasesInNewCollection(opts *bind.TransactOpts, _totalSupply uint32, _digitalMediaMetadataPath string, _collectionMetadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "createDigitalMediaAndReleasesInNewCollection", _totalSupply, _digitalMediaMetadataPath, _collectionMetadataPath, _numReleases)
}

// CreateDigitalMediaAndReleasesInNewCollection is a paid mutator transaction binding the contract method 0x8b40e8aa.
//
// Solidity: function createDigitalMediaAndReleasesInNewCollection(uint32 _totalSupply, string _digitalMediaMetadataPath, string _collectionMetadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) CreateDigitalMediaAndReleasesInNewCollection(_totalSupply uint32, _digitalMediaMetadataPath string, _collectionMetadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMediaAndReleasesInNewCollection(&_DigitalMediaCore.TransactOpts, _totalSupply, _digitalMediaMetadataPath, _collectionMetadataPath, _numReleases)
}

// CreateDigitalMediaAndReleasesInNewCollection is a paid mutator transaction binding the contract method 0x8b40e8aa.
//
// Solidity: function createDigitalMediaAndReleasesInNewCollection(uint32 _totalSupply, string _digitalMediaMetadataPath, string _collectionMetadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) CreateDigitalMediaAndReleasesInNewCollection(_totalSupply uint32, _digitalMediaMetadataPath string, _collectionMetadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMediaAndReleasesInNewCollection(&_DigitalMediaCore.TransactOpts, _totalSupply, _digitalMediaMetadataPath, _collectionMetadataPath, _numReleases)
}

// CreateDigitalMediaReleases is a paid mutator transaction binding the contract method 0x7fb08c9b.
//
// Solidity: function createDigitalMediaReleases(uint256 _digitalMediaId, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) CreateDigitalMediaReleases(opts *bind.TransactOpts, _digitalMediaId *big.Int, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "createDigitalMediaReleases", _digitalMediaId, _numReleases)
}

// CreateDigitalMediaReleases is a paid mutator transaction binding the contract method 0x7fb08c9b.
//
// Solidity: function createDigitalMediaReleases(uint256 _digitalMediaId, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) CreateDigitalMediaReleases(_digitalMediaId *big.Int, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMediaReleases(&_DigitalMediaCore.TransactOpts, _digitalMediaId, _numReleases)
}

// CreateDigitalMediaReleases is a paid mutator transaction binding the contract method 0x7fb08c9b.
//
// Solidity: function createDigitalMediaReleases(uint256 _digitalMediaId, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) CreateDigitalMediaReleases(_digitalMediaId *big.Int, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.CreateDigitalMediaReleases(&_DigitalMediaCore.TransactOpts, _digitalMediaId, _numReleases)
}

// DisableOboAddress is a paid mutator transaction binding the contract method 0x92b8bde1.
//
// Solidity: function disableOboAddress(address _oboAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) DisableOboAddress(opts *bind.TransactOpts, _oboAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "disableOboAddress", _oboAddress)
}

// DisableOboAddress is a paid mutator transaction binding the contract method 0x92b8bde1.
//
// Solidity: function disableOboAddress(address _oboAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) DisableOboAddress(_oboAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.DisableOboAddress(&_DigitalMediaCore.TransactOpts, _oboAddress)
}

// DisableOboAddress is a paid mutator transaction binding the contract method 0x92b8bde1.
//
// Solidity: function disableOboAddress(address _oboAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) DisableOboAddress(_oboAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.DisableOboAddress(&_DigitalMediaCore.TransactOpts, _oboAddress)
}

// OboCreateDigitalMediaAndReleases is a paid mutator transaction binding the contract method 0x6ee17a78.
//
// Solidity: function oboCreateDigitalMediaAndReleases(address _owner, uint32 _totalSupply, uint256 _collectionId, string _metadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) OboCreateDigitalMediaAndReleases(opts *bind.TransactOpts, _owner common.Address, _totalSupply uint32, _collectionId *big.Int, _metadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "oboCreateDigitalMediaAndReleases", _owner, _totalSupply, _collectionId, _metadataPath, _numReleases)
}

// OboCreateDigitalMediaAndReleases is a paid mutator transaction binding the contract method 0x6ee17a78.
//
// Solidity: function oboCreateDigitalMediaAndReleases(address _owner, uint32 _totalSupply, uint256 _collectionId, string _metadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) OboCreateDigitalMediaAndReleases(_owner common.Address, _totalSupply uint32, _collectionId *big.Int, _metadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.OboCreateDigitalMediaAndReleases(&_DigitalMediaCore.TransactOpts, _owner, _totalSupply, _collectionId, _metadataPath, _numReleases)
}

// OboCreateDigitalMediaAndReleases is a paid mutator transaction binding the contract method 0x6ee17a78.
//
// Solidity: function oboCreateDigitalMediaAndReleases(address _owner, uint32 _totalSupply, uint256 _collectionId, string _metadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) OboCreateDigitalMediaAndReleases(_owner common.Address, _totalSupply uint32, _collectionId *big.Int, _metadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.OboCreateDigitalMediaAndReleases(&_DigitalMediaCore.TransactOpts, _owner, _totalSupply, _collectionId, _metadataPath, _numReleases)
}

// OboCreateDigitalMediaAndReleasesInNewCollection is a paid mutator transaction binding the contract method 0x06a186d1.
//
// Solidity: function oboCreateDigitalMediaAndReleasesInNewCollection(address _owner, uint32 _totalSupply, string _digitalMediaMetadataPath, string _collectionMetadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) OboCreateDigitalMediaAndReleasesInNewCollection(opts *bind.TransactOpts, _owner common.Address, _totalSupply uint32, _digitalMediaMetadataPath string, _collectionMetadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "oboCreateDigitalMediaAndReleasesInNewCollection", _owner, _totalSupply, _digitalMediaMetadataPath, _collectionMetadataPath, _numReleases)
}

// OboCreateDigitalMediaAndReleasesInNewCollection is a paid mutator transaction binding the contract method 0x06a186d1.
//
// Solidity: function oboCreateDigitalMediaAndReleasesInNewCollection(address _owner, uint32 _totalSupply, string _digitalMediaMetadataPath, string _collectionMetadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) OboCreateDigitalMediaAndReleasesInNewCollection(_owner common.Address, _totalSupply uint32, _digitalMediaMetadataPath string, _collectionMetadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.OboCreateDigitalMediaAndReleasesInNewCollection(&_DigitalMediaCore.TransactOpts, _owner, _totalSupply, _digitalMediaMetadataPath, _collectionMetadataPath, _numReleases)
}

// OboCreateDigitalMediaAndReleasesInNewCollection is a paid mutator transaction binding the contract method 0x06a186d1.
//
// Solidity: function oboCreateDigitalMediaAndReleasesInNewCollection(address _owner, uint32 _totalSupply, string _digitalMediaMetadataPath, string _collectionMetadataPath, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) OboCreateDigitalMediaAndReleasesInNewCollection(_owner common.Address, _totalSupply uint32, _digitalMediaMetadataPath string, _collectionMetadataPath string, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.OboCreateDigitalMediaAndReleasesInNewCollection(&_DigitalMediaCore.TransactOpts, _owner, _totalSupply, _digitalMediaMetadataPath, _collectionMetadataPath, _numReleases)
}

// OboCreateDigitalMediaReleases is a paid mutator transaction binding the contract method 0xb36f0faa.
//
// Solidity: function oboCreateDigitalMediaReleases(address _owner, uint256 _digitalMediaId, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) OboCreateDigitalMediaReleases(opts *bind.TransactOpts, _owner common.Address, _digitalMediaId *big.Int, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "oboCreateDigitalMediaReleases", _owner, _digitalMediaId, _numReleases)
}

// OboCreateDigitalMediaReleases is a paid mutator transaction binding the contract method 0xb36f0faa.
//
// Solidity: function oboCreateDigitalMediaReleases(address _owner, uint256 _digitalMediaId, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) OboCreateDigitalMediaReleases(_owner common.Address, _digitalMediaId *big.Int, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.OboCreateDigitalMediaReleases(&_DigitalMediaCore.TransactOpts, _owner, _digitalMediaId, _numReleases)
}

// OboCreateDigitalMediaReleases is a paid mutator transaction binding the contract method 0xb36f0faa.
//
// Solidity: function oboCreateDigitalMediaReleases(address _owner, uint256 _digitalMediaId, uint32 _numReleases) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) OboCreateDigitalMediaReleases(_owner common.Address, _digitalMediaId *big.Int, _numReleases uint32) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.OboCreateDigitalMediaReleases(&_DigitalMediaCore.TransactOpts, _owner, _digitalMediaId, _numReleases)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) Pause() (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Pause(&_DigitalMediaCore.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) Pause() (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Pause(&_DigitalMediaCore.TransactOpts)
}

// RemoveApprovedTokenCreator is a paid mutator transaction binding the contract method 0x4751ae99.
//
// Solidity: function removeApprovedTokenCreator(address _creatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) RemoveApprovedTokenCreator(opts *bind.TransactOpts, _creatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "removeApprovedTokenCreator", _creatorAddress)
}

// RemoveApprovedTokenCreator is a paid mutator transaction binding the contract method 0x4751ae99.
//
// Solidity: function removeApprovedTokenCreator(address _creatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) RemoveApprovedTokenCreator(_creatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.RemoveApprovedTokenCreator(&_DigitalMediaCore.TransactOpts, _creatorAddress)
}

// RemoveApprovedTokenCreator is a paid mutator transaction binding the contract method 0x4751ae99.
//
// Solidity: function removeApprovedTokenCreator(address _creatorAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) RemoveApprovedTokenCreator(_creatorAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.RemoveApprovedTokenCreator(&_DigitalMediaCore.TransactOpts, _creatorAddress)
}

// ResetApproval is a paid mutator transaction binding the contract method 0xcf8dae05.
//
// Solidity: function resetApproval(uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) ResetApproval(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "resetApproval", _tokenId)
}

// ResetApproval is a paid mutator transaction binding the contract method 0xcf8dae05.
//
// Solidity: function resetApproval(uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) ResetApproval(_tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.ResetApproval(&_DigitalMediaCore.TransactOpts, _tokenId)
}

// ResetApproval is a paid mutator transaction binding the contract method 0xcf8dae05.
//
// Solidity: function resetApproval(uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) ResetApproval(_tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.ResetApproval(&_DigitalMediaCore.TransactOpts, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SafeTransferFrom(&_DigitalMediaCore.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SafeTransferFrom(&_DigitalMediaCore.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SafeTransferFrom0(&_DigitalMediaCore.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SafeTransferFrom0(&_DigitalMediaCore.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SetApprovalForAll(&_DigitalMediaCore.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _to, bool _approved) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SetApprovalForAll(&_DigitalMediaCore.TransactOpts, _to, _approved)
}

// SetOboApprovalForAll is a paid mutator transaction binding the contract method 0x774f99d0.
//
// Solidity: function setOboApprovalForAll(address _to, bool _approved) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) SetOboApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "setOboApprovalForAll", _to, _approved)
}

// SetOboApprovalForAll is a paid mutator transaction binding the contract method 0x774f99d0.
//
// Solidity: function setOboApprovalForAll(address _to, bool _approved) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) SetOboApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SetOboApprovalForAll(&_DigitalMediaCore.TransactOpts, _to, _approved)
}

// SetOboApprovalForAll is a paid mutator transaction binding the contract method 0x774f99d0.
//
// Solidity: function setOboApprovalForAll(address _to, bool _approved) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) SetOboApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SetOboApprovalForAll(&_DigitalMediaCore.TransactOpts, _to, _approved)
}

// SetV1DigitalMediaStoreAddress is a paid mutator transaction binding the contract method 0x91c60788.
//
// Solidity: function setV1DigitalMediaStoreAddress(address _dmsAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) SetV1DigitalMediaStoreAddress(opts *bind.TransactOpts, _dmsAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "setV1DigitalMediaStoreAddress", _dmsAddress)
}

// SetV1DigitalMediaStoreAddress is a paid mutator transaction binding the contract method 0x91c60788.
//
// Solidity: function setV1DigitalMediaStoreAddress(address _dmsAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) SetV1DigitalMediaStoreAddress(_dmsAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SetV1DigitalMediaStoreAddress(&_DigitalMediaCore.TransactOpts, _dmsAddress)
}

// SetV1DigitalMediaStoreAddress is a paid mutator transaction binding the contract method 0x91c60788.
//
// Solidity: function setV1DigitalMediaStoreAddress(address _dmsAddress) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) SetV1DigitalMediaStoreAddress(_dmsAddress common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.SetV1DigitalMediaStoreAddress(&_DigitalMediaCore.TransactOpts, _dmsAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.TransferFrom(&_DigitalMediaCore.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.TransferFrom(&_DigitalMediaCore.TransactOpts, _from, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.TransferOwnership(&_DigitalMediaCore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.TransferOwnership(&_DigitalMediaCore.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DigitalMediaCore.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DigitalMediaCore *DigitalMediaCoreSession) Unpause() (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Unpause(&_DigitalMediaCore.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_DigitalMediaCore *DigitalMediaCoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _DigitalMediaCore.Contract.Unpause(&_DigitalMediaCore.TransactOpts)
}

// DigitalMediaCoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DigitalMediaCore contract.
type DigitalMediaCoreApprovalIterator struct {
	Event *DigitalMediaCoreApproval // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreApproval)
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
		it.Event = new(DigitalMediaCoreApproval)
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
func (it *DigitalMediaCoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreApproval represents a Approval event raised by the DigitalMediaCore contract.
type DigitalMediaCoreApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*DigitalMediaCoreApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreApprovalIterator{contract: _DigitalMediaCore.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreApproval)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseApproval(log types.Log) (*DigitalMediaCoreApproval, error) {
	event := new(DigitalMediaCoreApproval)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DigitalMediaCore contract.
type DigitalMediaCoreApprovalForAllIterator struct {
	Event *DigitalMediaCoreApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreApprovalForAll)
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
		it.Event = new(DigitalMediaCoreApprovalForAll)
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
func (it *DigitalMediaCoreApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreApprovalForAll represents a ApprovalForAll event raised by the DigitalMediaCore contract.
type DigitalMediaCoreApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*DigitalMediaCoreApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreApprovalForAllIterator{contract: _DigitalMediaCore.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreApprovalForAll)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseApprovalForAll(log types.Log) (*DigitalMediaCoreApprovalForAll, error) {
	event := new(DigitalMediaCoreApprovalForAll)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreChangedCreatorIterator is returned from FilterChangedCreator and is used to iterate over the raw logs and unpacked data for ChangedCreator events raised by the DigitalMediaCore contract.
type DigitalMediaCoreChangedCreatorIterator struct {
	Event *DigitalMediaCoreChangedCreator // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreChangedCreatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreChangedCreator)
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
		it.Event = new(DigitalMediaCoreChangedCreator)
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
func (it *DigitalMediaCoreChangedCreatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreChangedCreatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreChangedCreator represents a ChangedCreator event raised by the DigitalMediaCore contract.
type DigitalMediaCoreChangedCreator struct {
	Creator    common.Address
	NewCreator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangedCreator is a free log retrieval operation binding the contract event 0xde6cfdf21fe76bcb45258138e27bcd332b76941b24d226b5da8dc5f9cd531c3e.
//
// Solidity: event ChangedCreator(address creator, address newCreator)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterChangedCreator(opts *bind.FilterOpts) (*DigitalMediaCoreChangedCreatorIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "ChangedCreator")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreChangedCreatorIterator{contract: _DigitalMediaCore.contract, event: "ChangedCreator", logs: logs, sub: sub}, nil
}

// WatchChangedCreator is a free log subscription operation binding the contract event 0xde6cfdf21fe76bcb45258138e27bcd332b76941b24d226b5da8dc5f9cd531c3e.
//
// Solidity: event ChangedCreator(address creator, address newCreator)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchChangedCreator(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreChangedCreator) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "ChangedCreator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreChangedCreator)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "ChangedCreator", log); err != nil {
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

// ParseChangedCreator is a log parse operation binding the contract event 0xde6cfdf21fe76bcb45258138e27bcd332b76941b24d226b5da8dc5f9cd531c3e.
//
// Solidity: event ChangedCreator(address creator, address newCreator)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseChangedCreator(log types.Log) (*DigitalMediaCoreChangedCreator, error) {
	event := new(DigitalMediaCoreChangedCreator)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "ChangedCreator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreDigitalMediaBurnEventIterator is returned from FilterDigitalMediaBurnEvent and is used to iterate over the raw logs and unpacked data for DigitalMediaBurnEvent events raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaBurnEventIterator struct {
	Event *DigitalMediaCoreDigitalMediaBurnEvent // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreDigitalMediaBurnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreDigitalMediaBurnEvent)
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
		it.Event = new(DigitalMediaCoreDigitalMediaBurnEvent)
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
func (it *DigitalMediaCoreDigitalMediaBurnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreDigitalMediaBurnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreDigitalMediaBurnEvent represents a DigitalMediaBurnEvent event raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaBurnEvent struct {
	Id                   *big.Int
	Caller               common.Address
	StoreContractAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDigitalMediaBurnEvent is a free log retrieval operation binding the contract event 0x327ecc068f1b41267f69376098f6a50da487e4a4d762d53c01197d6a2f294b3e.
//
// Solidity: event DigitalMediaBurnEvent(uint256 id, address caller, address storeContractAddress)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterDigitalMediaBurnEvent(opts *bind.FilterOpts) (*DigitalMediaCoreDigitalMediaBurnEventIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "DigitalMediaBurnEvent")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreDigitalMediaBurnEventIterator{contract: _DigitalMediaCore.contract, event: "DigitalMediaBurnEvent", logs: logs, sub: sub}, nil
}

// WatchDigitalMediaBurnEvent is a free log subscription operation binding the contract event 0x327ecc068f1b41267f69376098f6a50da487e4a4d762d53c01197d6a2f294b3e.
//
// Solidity: event DigitalMediaBurnEvent(uint256 id, address caller, address storeContractAddress)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchDigitalMediaBurnEvent(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreDigitalMediaBurnEvent) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "DigitalMediaBurnEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreDigitalMediaBurnEvent)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaBurnEvent", log); err != nil {
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

// ParseDigitalMediaBurnEvent is a log parse operation binding the contract event 0x327ecc068f1b41267f69376098f6a50da487e4a4d762d53c01197d6a2f294b3e.
//
// Solidity: event DigitalMediaBurnEvent(uint256 id, address caller, address storeContractAddress)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseDigitalMediaBurnEvent(log types.Log) (*DigitalMediaCoreDigitalMediaBurnEvent, error) {
	event := new(DigitalMediaCoreDigitalMediaBurnEvent)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaBurnEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreDigitalMediaCollectionCreateEventIterator is returned from FilterDigitalMediaCollectionCreateEvent and is used to iterate over the raw logs and unpacked data for DigitalMediaCollectionCreateEvent events raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaCollectionCreateEventIterator struct {
	Event *DigitalMediaCoreDigitalMediaCollectionCreateEvent // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreDigitalMediaCollectionCreateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreDigitalMediaCollectionCreateEvent)
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
		it.Event = new(DigitalMediaCoreDigitalMediaCollectionCreateEvent)
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
func (it *DigitalMediaCoreDigitalMediaCollectionCreateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreDigitalMediaCollectionCreateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreDigitalMediaCollectionCreateEvent represents a DigitalMediaCollectionCreateEvent event raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaCollectionCreateEvent struct {
	Id                   *big.Int
	StoreContractAddress common.Address
	Creator              common.Address
	MetadataPath         string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDigitalMediaCollectionCreateEvent is a free log retrieval operation binding the contract event 0x01e2312dcdafe7cd3f82579d8c121fdb930d46ef2eb231953a521ac62093e277.
//
// Solidity: event DigitalMediaCollectionCreateEvent(uint256 id, address storeContractAddress, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterDigitalMediaCollectionCreateEvent(opts *bind.FilterOpts) (*DigitalMediaCoreDigitalMediaCollectionCreateEventIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "DigitalMediaCollectionCreateEvent")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreDigitalMediaCollectionCreateEventIterator{contract: _DigitalMediaCore.contract, event: "DigitalMediaCollectionCreateEvent", logs: logs, sub: sub}, nil
}

// WatchDigitalMediaCollectionCreateEvent is a free log subscription operation binding the contract event 0x01e2312dcdafe7cd3f82579d8c121fdb930d46ef2eb231953a521ac62093e277.
//
// Solidity: event DigitalMediaCollectionCreateEvent(uint256 id, address storeContractAddress, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchDigitalMediaCollectionCreateEvent(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreDigitalMediaCollectionCreateEvent) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "DigitalMediaCollectionCreateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreDigitalMediaCollectionCreateEvent)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaCollectionCreateEvent", log); err != nil {
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

// ParseDigitalMediaCollectionCreateEvent is a log parse operation binding the contract event 0x01e2312dcdafe7cd3f82579d8c121fdb930d46ef2eb231953a521ac62093e277.
//
// Solidity: event DigitalMediaCollectionCreateEvent(uint256 id, address storeContractAddress, address creator, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseDigitalMediaCollectionCreateEvent(log types.Log) (*DigitalMediaCoreDigitalMediaCollectionCreateEvent, error) {
	event := new(DigitalMediaCoreDigitalMediaCollectionCreateEvent)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaCollectionCreateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreDigitalMediaCreateEventIterator is returned from FilterDigitalMediaCreateEvent and is used to iterate over the raw logs and unpacked data for DigitalMediaCreateEvent events raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaCreateEventIterator struct {
	Event *DigitalMediaCoreDigitalMediaCreateEvent // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreDigitalMediaCreateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreDigitalMediaCreateEvent)
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
		it.Event = new(DigitalMediaCoreDigitalMediaCreateEvent)
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
func (it *DigitalMediaCoreDigitalMediaCreateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreDigitalMediaCreateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreDigitalMediaCreateEvent represents a DigitalMediaCreateEvent event raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaCreateEvent struct {
	Id                   *big.Int
	StoreContractAddress common.Address
	Creator              common.Address
	TotalSupply          uint32
	PrintIndex           uint32
	CollectionId         *big.Int
	MetadataPath         string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDigitalMediaCreateEvent is a free log retrieval operation binding the contract event 0x794c5cd70604d9d8dc2cbca1f8be65f167e4147b6512541d41e8e410594098a0.
//
// Solidity: event DigitalMediaCreateEvent(uint256 id, address storeContractAddress, address creator, uint32 totalSupply, uint32 printIndex, uint256 collectionId, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterDigitalMediaCreateEvent(opts *bind.FilterOpts) (*DigitalMediaCoreDigitalMediaCreateEventIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "DigitalMediaCreateEvent")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreDigitalMediaCreateEventIterator{contract: _DigitalMediaCore.contract, event: "DigitalMediaCreateEvent", logs: logs, sub: sub}, nil
}

// WatchDigitalMediaCreateEvent is a free log subscription operation binding the contract event 0x794c5cd70604d9d8dc2cbca1f8be65f167e4147b6512541d41e8e410594098a0.
//
// Solidity: event DigitalMediaCreateEvent(uint256 id, address storeContractAddress, address creator, uint32 totalSupply, uint32 printIndex, uint256 collectionId, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchDigitalMediaCreateEvent(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreDigitalMediaCreateEvent) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "DigitalMediaCreateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreDigitalMediaCreateEvent)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaCreateEvent", log); err != nil {
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

// ParseDigitalMediaCreateEvent is a log parse operation binding the contract event 0x794c5cd70604d9d8dc2cbca1f8be65f167e4147b6512541d41e8e410594098a0.
//
// Solidity: event DigitalMediaCreateEvent(uint256 id, address storeContractAddress, address creator, uint32 totalSupply, uint32 printIndex, uint256 collectionId, string metadataPath)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseDigitalMediaCreateEvent(log types.Log) (*DigitalMediaCoreDigitalMediaCreateEvent, error) {
	event := new(DigitalMediaCoreDigitalMediaCreateEvent)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaCreateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreDigitalMediaReleaseBurnEventIterator is returned from FilterDigitalMediaReleaseBurnEvent and is used to iterate over the raw logs and unpacked data for DigitalMediaReleaseBurnEvent events raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaReleaseBurnEventIterator struct {
	Event *DigitalMediaCoreDigitalMediaReleaseBurnEvent // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreDigitalMediaReleaseBurnEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreDigitalMediaReleaseBurnEvent)
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
		it.Event = new(DigitalMediaCoreDigitalMediaReleaseBurnEvent)
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
func (it *DigitalMediaCoreDigitalMediaReleaseBurnEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreDigitalMediaReleaseBurnEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreDigitalMediaReleaseBurnEvent represents a DigitalMediaReleaseBurnEvent event raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaReleaseBurnEvent struct {
	TokenId *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDigitalMediaReleaseBurnEvent is a free log retrieval operation binding the contract event 0x1e8df141f42ed659a8fe7e7c5966cbdf2d240d0c45f4c30cbe02526c618075ef.
//
// Solidity: event DigitalMediaReleaseBurnEvent(uint256 tokenId, address owner)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterDigitalMediaReleaseBurnEvent(opts *bind.FilterOpts) (*DigitalMediaCoreDigitalMediaReleaseBurnEventIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "DigitalMediaReleaseBurnEvent")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreDigitalMediaReleaseBurnEventIterator{contract: _DigitalMediaCore.contract, event: "DigitalMediaReleaseBurnEvent", logs: logs, sub: sub}, nil
}

// WatchDigitalMediaReleaseBurnEvent is a free log subscription operation binding the contract event 0x1e8df141f42ed659a8fe7e7c5966cbdf2d240d0c45f4c30cbe02526c618075ef.
//
// Solidity: event DigitalMediaReleaseBurnEvent(uint256 tokenId, address owner)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchDigitalMediaReleaseBurnEvent(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreDigitalMediaReleaseBurnEvent) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "DigitalMediaReleaseBurnEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreDigitalMediaReleaseBurnEvent)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaReleaseBurnEvent", log); err != nil {
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

// ParseDigitalMediaReleaseBurnEvent is a log parse operation binding the contract event 0x1e8df141f42ed659a8fe7e7c5966cbdf2d240d0c45f4c30cbe02526c618075ef.
//
// Solidity: event DigitalMediaReleaseBurnEvent(uint256 tokenId, address owner)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseDigitalMediaReleaseBurnEvent(log types.Log) (*DigitalMediaCoreDigitalMediaReleaseBurnEvent, error) {
	event := new(DigitalMediaCoreDigitalMediaReleaseBurnEvent)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaReleaseBurnEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreDigitalMediaReleaseCreateEventIterator is returned from FilterDigitalMediaReleaseCreateEvent and is used to iterate over the raw logs and unpacked data for DigitalMediaReleaseCreateEvent events raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaReleaseCreateEventIterator struct {
	Event *DigitalMediaCoreDigitalMediaReleaseCreateEvent // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreDigitalMediaReleaseCreateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreDigitalMediaReleaseCreateEvent)
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
		it.Event = new(DigitalMediaCoreDigitalMediaReleaseCreateEvent)
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
func (it *DigitalMediaCoreDigitalMediaReleaseCreateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreDigitalMediaReleaseCreateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreDigitalMediaReleaseCreateEvent represents a DigitalMediaReleaseCreateEvent event raised by the DigitalMediaCore contract.
type DigitalMediaCoreDigitalMediaReleaseCreateEvent struct {
	Id             *big.Int
	Owner          common.Address
	PrintEdition   uint32
	TokenURI       string
	DigitalMediaId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDigitalMediaReleaseCreateEvent is a free log retrieval operation binding the contract event 0x775f53e4c75ce0c74e611f7f0bb660e4cd647e0522ef0f8aefd4ecef373c5df9.
//
// Solidity: event DigitalMediaReleaseCreateEvent(uint256 id, address owner, uint32 printEdition, string tokenURI, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterDigitalMediaReleaseCreateEvent(opts *bind.FilterOpts) (*DigitalMediaCoreDigitalMediaReleaseCreateEventIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "DigitalMediaReleaseCreateEvent")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreDigitalMediaReleaseCreateEventIterator{contract: _DigitalMediaCore.contract, event: "DigitalMediaReleaseCreateEvent", logs: logs, sub: sub}, nil
}

// WatchDigitalMediaReleaseCreateEvent is a free log subscription operation binding the contract event 0x775f53e4c75ce0c74e611f7f0bb660e4cd647e0522ef0f8aefd4ecef373c5df9.
//
// Solidity: event DigitalMediaReleaseCreateEvent(uint256 id, address owner, uint32 printEdition, string tokenURI, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchDigitalMediaReleaseCreateEvent(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreDigitalMediaReleaseCreateEvent) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "DigitalMediaReleaseCreateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreDigitalMediaReleaseCreateEvent)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaReleaseCreateEvent", log); err != nil {
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

// ParseDigitalMediaReleaseCreateEvent is a log parse operation binding the contract event 0x775f53e4c75ce0c74e611f7f0bb660e4cd647e0522ef0f8aefd4ecef373c5df9.
//
// Solidity: event DigitalMediaReleaseCreateEvent(uint256 id, address owner, uint32 printEdition, string tokenURI, uint256 digitalMediaId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseDigitalMediaReleaseCreateEvent(log types.Log) (*DigitalMediaCoreDigitalMediaReleaseCreateEvent, error) {
	event := new(DigitalMediaCoreDigitalMediaReleaseCreateEvent)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "DigitalMediaReleaseCreateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreOboApprovalForAllIterator is returned from FilterOboApprovalForAll and is used to iterate over the raw logs and unpacked data for OboApprovalForAll events raised by the DigitalMediaCore contract.
type DigitalMediaCoreOboApprovalForAllIterator struct {
	Event *DigitalMediaCoreOboApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreOboApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreOboApprovalForAll)
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
		it.Event = new(DigitalMediaCoreOboApprovalForAll)
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
func (it *DigitalMediaCoreOboApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreOboApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreOboApprovalForAll represents a OboApprovalForAll event raised by the DigitalMediaCore contract.
type DigitalMediaCoreOboApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOboApprovalForAll is a free log retrieval operation binding the contract event 0x017e8a478826a4348bfb695968246edfab885f8a76b03279cf4630ac073945c9.
//
// Solidity: event OboApprovalForAll(address _owner, address _operator, bool _approved)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterOboApprovalForAll(opts *bind.FilterOpts) (*DigitalMediaCoreOboApprovalForAllIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "OboApprovalForAll")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreOboApprovalForAllIterator{contract: _DigitalMediaCore.contract, event: "OboApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchOboApprovalForAll is a free log subscription operation binding the contract event 0x017e8a478826a4348bfb695968246edfab885f8a76b03279cf4630ac073945c9.
//
// Solidity: event OboApprovalForAll(address _owner, address _operator, bool _approved)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchOboApprovalForAll(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreOboApprovalForAll) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "OboApprovalForAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreOboApprovalForAll)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "OboApprovalForAll", log); err != nil {
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

// ParseOboApprovalForAll is a log parse operation binding the contract event 0x017e8a478826a4348bfb695968246edfab885f8a76b03279cf4630ac073945c9.
//
// Solidity: event OboApprovalForAll(address _owner, address _operator, bool _approved)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseOboApprovalForAll(log types.Log) (*DigitalMediaCoreOboApprovalForAll, error) {
	event := new(DigitalMediaCoreOboApprovalForAll)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "OboApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreOboDisabledForAllIterator is returned from FilterOboDisabledForAll and is used to iterate over the raw logs and unpacked data for OboDisabledForAll events raised by the DigitalMediaCore contract.
type DigitalMediaCoreOboDisabledForAllIterator struct {
	Event *DigitalMediaCoreOboDisabledForAll // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreOboDisabledForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreOboDisabledForAll)
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
		it.Event = new(DigitalMediaCoreOboDisabledForAll)
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
func (it *DigitalMediaCoreOboDisabledForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreOboDisabledForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreOboDisabledForAll represents a OboDisabledForAll event raised by the DigitalMediaCore contract.
type DigitalMediaCoreOboDisabledForAll struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOboDisabledForAll is a free log retrieval operation binding the contract event 0xfd0e0c743dbdd84ef4e7c513db9b7e085970164787288791343fda28575652dd.
//
// Solidity: event OboDisabledForAll(address _operator)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterOboDisabledForAll(opts *bind.FilterOpts) (*DigitalMediaCoreOboDisabledForAllIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "OboDisabledForAll")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreOboDisabledForAllIterator{contract: _DigitalMediaCore.contract, event: "OboDisabledForAll", logs: logs, sub: sub}, nil
}

// WatchOboDisabledForAll is a free log subscription operation binding the contract event 0xfd0e0c743dbdd84ef4e7c513db9b7e085970164787288791343fda28575652dd.
//
// Solidity: event OboDisabledForAll(address _operator)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchOboDisabledForAll(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreOboDisabledForAll) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "OboDisabledForAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreOboDisabledForAll)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "OboDisabledForAll", log); err != nil {
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

// ParseOboDisabledForAll is a log parse operation binding the contract event 0xfd0e0c743dbdd84ef4e7c513db9b7e085970164787288791343fda28575652dd.
//
// Solidity: event OboDisabledForAll(address _operator)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseOboDisabledForAll(log types.Log) (*DigitalMediaCoreOboDisabledForAll, error) {
	event := new(DigitalMediaCoreOboDisabledForAll)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "OboDisabledForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DigitalMediaCore contract.
type DigitalMediaCoreOwnershipTransferredIterator struct {
	Event *DigitalMediaCoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreOwnershipTransferred)
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
		it.Event = new(DigitalMediaCoreOwnershipTransferred)
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
func (it *DigitalMediaCoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreOwnershipTransferred represents a OwnershipTransferred event raised by the DigitalMediaCore contract.
type DigitalMediaCoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DigitalMediaCoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreOwnershipTransferredIterator{contract: _DigitalMediaCore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreOwnershipTransferred)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseOwnershipTransferred(log types.Log) (*DigitalMediaCoreOwnershipTransferred, error) {
	event := new(DigitalMediaCoreOwnershipTransferred)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCorePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the DigitalMediaCore contract.
type DigitalMediaCorePauseIterator struct {
	Event *DigitalMediaCorePause // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCorePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCorePause)
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
		it.Event = new(DigitalMediaCorePause)
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
func (it *DigitalMediaCorePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCorePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCorePause represents a Pause event raised by the DigitalMediaCore contract.
type DigitalMediaCorePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterPause(opts *bind.FilterOpts) (*DigitalMediaCorePauseIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCorePauseIterator{contract: _DigitalMediaCore.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *DigitalMediaCorePause) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCorePause)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParsePause(log types.Log) (*DigitalMediaCorePause, error) {
	event := new(DigitalMediaCorePause)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreSingleCreatorChangedIterator is returned from FilterSingleCreatorChanged and is used to iterate over the raw logs and unpacked data for SingleCreatorChanged events raised by the DigitalMediaCore contract.
type DigitalMediaCoreSingleCreatorChangedIterator struct {
	Event *DigitalMediaCoreSingleCreatorChanged // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreSingleCreatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreSingleCreatorChanged)
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
		it.Event = new(DigitalMediaCoreSingleCreatorChanged)
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
func (it *DigitalMediaCoreSingleCreatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreSingleCreatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreSingleCreatorChanged represents a SingleCreatorChanged event raised by the DigitalMediaCore contract.
type DigitalMediaCoreSingleCreatorChanged struct {
	PreviousCreatorAddress common.Address
	NewCreatorAddress      common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSingleCreatorChanged is a free log retrieval operation binding the contract event 0x384c948063df3740539b4b000658c1a22348e7f18c87f808085662e461e48e71.
//
// Solidity: event SingleCreatorChanged(address indexed previousCreatorAddress, address indexed newCreatorAddress)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterSingleCreatorChanged(opts *bind.FilterOpts, previousCreatorAddress []common.Address, newCreatorAddress []common.Address) (*DigitalMediaCoreSingleCreatorChangedIterator, error) {

	var previousCreatorAddressRule []interface{}
	for _, previousCreatorAddressItem := range previousCreatorAddress {
		previousCreatorAddressRule = append(previousCreatorAddressRule, previousCreatorAddressItem)
	}
	var newCreatorAddressRule []interface{}
	for _, newCreatorAddressItem := range newCreatorAddress {
		newCreatorAddressRule = append(newCreatorAddressRule, newCreatorAddressItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "SingleCreatorChanged", previousCreatorAddressRule, newCreatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreSingleCreatorChangedIterator{contract: _DigitalMediaCore.contract, event: "SingleCreatorChanged", logs: logs, sub: sub}, nil
}

// WatchSingleCreatorChanged is a free log subscription operation binding the contract event 0x384c948063df3740539b4b000658c1a22348e7f18c87f808085662e461e48e71.
//
// Solidity: event SingleCreatorChanged(address indexed previousCreatorAddress, address indexed newCreatorAddress)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchSingleCreatorChanged(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreSingleCreatorChanged, previousCreatorAddress []common.Address, newCreatorAddress []common.Address) (event.Subscription, error) {

	var previousCreatorAddressRule []interface{}
	for _, previousCreatorAddressItem := range previousCreatorAddress {
		previousCreatorAddressRule = append(previousCreatorAddressRule, previousCreatorAddressItem)
	}
	var newCreatorAddressRule []interface{}
	for _, newCreatorAddressItem := range newCreatorAddress {
		newCreatorAddressRule = append(newCreatorAddressRule, newCreatorAddressItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "SingleCreatorChanged", previousCreatorAddressRule, newCreatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreSingleCreatorChanged)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "SingleCreatorChanged", log); err != nil {
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

// ParseSingleCreatorChanged is a log parse operation binding the contract event 0x384c948063df3740539b4b000658c1a22348e7f18c87f808085662e461e48e71.
//
// Solidity: event SingleCreatorChanged(address indexed previousCreatorAddress, address indexed newCreatorAddress)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseSingleCreatorChanged(log types.Log) (*DigitalMediaCoreSingleCreatorChanged, error) {
	event := new(DigitalMediaCoreSingleCreatorChanged)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "SingleCreatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DigitalMediaCore contract.
type DigitalMediaCoreTransferIterator struct {
	Event *DigitalMediaCoreTransfer // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreTransfer)
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
		it.Event = new(DigitalMediaCoreTransfer)
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
func (it *DigitalMediaCoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreTransfer represents a Transfer event raised by the DigitalMediaCore contract.
type DigitalMediaCoreTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*DigitalMediaCoreTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreTransferIterator{contract: _DigitalMediaCore.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreTransfer)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseTransfer(log types.Log) (*DigitalMediaCoreTransfer, error) {
	event := new(DigitalMediaCoreTransfer)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the DigitalMediaCore contract.
type DigitalMediaCoreUnpauseIterator struct {
	Event *DigitalMediaCoreUnpause // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreUnpause)
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
		it.Event = new(DigitalMediaCoreUnpause)
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
func (it *DigitalMediaCoreUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreUnpause represents a Unpause event raised by the DigitalMediaCore contract.
type DigitalMediaCoreUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterUnpause(opts *bind.FilterOpts) (*DigitalMediaCoreUnpauseIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreUnpauseIterator{contract: _DigitalMediaCore.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreUnpause) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreUnpause)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseUnpause(log types.Log) (*DigitalMediaCoreUnpause, error) {
	event := new(DigitalMediaCoreUnpause)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator is returned from FilterUpdateDigitalMediaPrintIndexEvent and is used to iterate over the raw logs and unpacked data for UpdateDigitalMediaPrintIndexEvent events raised by the DigitalMediaCore contract.
type DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator struct {
	Event *DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent // Event containing the contract specifics and raw log

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
func (it *DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent)
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
		it.Event = new(DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent)
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
func (it *DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent represents a UpdateDigitalMediaPrintIndexEvent event raised by the DigitalMediaCore contract.
type DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent struct {
	DigitalMediaId *big.Int
	PrintEdition   uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateDigitalMediaPrintIndexEvent is a free log retrieval operation binding the contract event 0x12d99f5e49ef761c52953e4f9a109827fc3540292ba88c10d309fef470685259.
//
// Solidity: event UpdateDigitalMediaPrintIndexEvent(uint256 digitalMediaId, uint32 printEdition)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) FilterUpdateDigitalMediaPrintIndexEvent(opts *bind.FilterOpts) (*DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator, error) {

	logs, sub, err := _DigitalMediaCore.contract.FilterLogs(opts, "UpdateDigitalMediaPrintIndexEvent")
	if err != nil {
		return nil, err
	}
	return &DigitalMediaCoreUpdateDigitalMediaPrintIndexEventIterator{contract: _DigitalMediaCore.contract, event: "UpdateDigitalMediaPrintIndexEvent", logs: logs, sub: sub}, nil
}

// WatchUpdateDigitalMediaPrintIndexEvent is a free log subscription operation binding the contract event 0x12d99f5e49ef761c52953e4f9a109827fc3540292ba88c10d309fef470685259.
//
// Solidity: event UpdateDigitalMediaPrintIndexEvent(uint256 digitalMediaId, uint32 printEdition)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) WatchUpdateDigitalMediaPrintIndexEvent(opts *bind.WatchOpts, sink chan<- *DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent) (event.Subscription, error) {

	logs, sub, err := _DigitalMediaCore.contract.WatchLogs(opts, "UpdateDigitalMediaPrintIndexEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent)
				if err := _DigitalMediaCore.contract.UnpackLog(event, "UpdateDigitalMediaPrintIndexEvent", log); err != nil {
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

// ParseUpdateDigitalMediaPrintIndexEvent is a log parse operation binding the contract event 0x12d99f5e49ef761c52953e4f9a109827fc3540292ba88c10d309fef470685259.
//
// Solidity: event UpdateDigitalMediaPrintIndexEvent(uint256 digitalMediaId, uint32 printEdition)
func (_DigitalMediaCore *DigitalMediaCoreFilterer) ParseUpdateDigitalMediaPrintIndexEvent(log types.Log) (*DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent, error) {
	event := new(DigitalMediaCoreUpdateDigitalMediaPrintIndexEvent)
	if err := _DigitalMediaCore.contract.UnpackLog(event, "UpdateDigitalMediaPrintIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
