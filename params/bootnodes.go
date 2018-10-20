// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Ethereum Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstrem bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}

var CallistoBootnodes = []string{
	"enode://ab0d9fe81f23653c3217303d3a4bc035be908b05f8b12d6f155ab4598d7760cfea1e009e414771279c9ffc3499950283afaabe099a5329c5a6013f57d77de0a6@104.236.96.118:30303",
	"enode://9e2d9dc2639e02893aa17c80e6ba8e8803fd3166083b622a841fc852161112720281514436f7605c89041d5efa1738215185c4c4024ff812b0f500c403cc0ab1@206.189.47.198:30303",
	"enode://149ba679e8851c3e0d030e0dc0336984b97c83ef649e68ec113dcf266449364cc1ec8ee27950f71b00c2182ef504894fa7bff19f6741978ced67e9e4b6536d2a@206.189.45.31:30303",
	"enode://eeb3b1680f651b291a19454345721b5196a2a689dcd280e5f66dc3207636366d4b25e84d205303c2f8aa0a38467dad9e6f2536e195a4760df56aeac428ebea0b@199.247.18.157:30303",
	"enode://c21418c02f5ba480d64fbdf3dd7e1a276cbea441f9d55b8bb1c653fa3a05cc07e32a332f63df53f51d275dadc9b50925375a699dd39bdad991594326d6b8afab@199.247.3.3:30303",
	"enode://40aa8ff2c3d98ffc12004cb6a3636e7c9b79400153667163cfc24123f2ee3ee693ac45775183f5f6a7e315a4884899ac32ef0616e26cdd23a7b00f80d07cdeae@45.32.126.82:30303",
	"enode://3beb80913887d985a857076621baca66ea27b62ff159c5a41243d02a8614f537003c03ba1fe082b63a47e7f6f7ba1caf6bb14343560dbff6ba1e456e99e6119d@144.202.73.111:30303",
	"enode://b79a50393b16b76a6c94d7ddae80c44464c9e5ecc59fc2b7e83d0c248190de781e7e2aebeab8d466b3869677e6388e6fac8bd36f3925cbeebbe4cf0372a7eac1@207.148.31.238:30303",
	"enode://dcc091d6da928681e76993cffa047e434b0c0b1388c33330499cd006547de6641f8e5b885c2be12dee07f79d8bd86612cd1b16b4ab2536cbcdfceb278bf403f7@165.227.5.237:30303",
	"enode://019faf0a35f3bcf6ac47c54d9d73917378f82acc6ae7197e0cdc037325e934cf12c7587bcbf0358abb2ab5b139780c97ed7d15ec2bd7784a5570e2cc58479bed@104.156.230.179:30303",
	"enode://cd85a09fddddfed9ad6858ba1cffb7b6688e4197b85b722f79ccd8b97a7ee8d6d6ef173af91618c24032b5dbaff3265b63bb85dc768d58d830e9aa654cc51838@46.101.190.119:30303",
	"enode://d34e34dc0d57694a71611d9bb3df73c928ef77feb05f1e6e3a6243d26b75034c09a6eab37872f7ce830890d87ca180a8d359fa0a649c2d12a05c853c52bae264@159.65.237.29:30303",
}

var CallistoTestnetBootnodes = []string{
	"enode://33c47b8d28daa3773751ab02fcb486806b3260df02a58bbfb4d00af0fcec79ed64a96dd5cfe4e78b48af59fb4d549f659e4943cdd0c73ccb217b59164ca2549a@45.55.38.193:30303",
}