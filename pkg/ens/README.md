curl "https://api.etherscan.io/api?module=contract&action=getsourcecode&address=0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e" | jq -r '.result[0].SourceCode' > ENSRegistryWithFallback.sol

curl "https://api.etherscan.io/api?module=contract&action=getsourcecode&address=0x4976fb03c32e5b8cfe2b6ccb31c09ba78ebaba41" | jq -r '.result[0].SourceCode' > contracts/PublicResolver.sol

abigen --solc ~/Downloads/solc-static-linux-5 --sol contracts/ENSRegistryWithFallback.sol --pkg ens --out ENSRegistryWithFallback.go

abigen --solc ~/Downloads/solc-static-linux-5 --sol contracts/PublicResolver.sol --pkg ens --out PublicResolver.go --exc contracts/PublicResolver.sol:ENS
