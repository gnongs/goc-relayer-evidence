package types

import "time"

type ReturnData struct {
	Txs []struct {
		Body struct {
			Messages []struct {
				TypeURL string `json:"@type"`
				Header  struct {
					Type         string `json:"@type"`
					SignedHeader struct {
						Header struct {
							Version struct {
								Block string `json:"block"`
								App   string `json:"app"`
							} `json:"version"`
							ChainID     string    `json:"chain_id"`
							Height      string    `json:"height"`
							Time        time.Time `json:"time"`
							LastBlockID struct {
								Hash          string `json:"hash"`
								PartSetHeader struct {
									Total int    `json:"total"`
									Hash  string `json:"hash"`
								} `json:"part_set_header"`
							} `json:"last_block_id"`
							LastCommitHash     string `json:"last_commit_hash"`
							DataHash           string `json:"data_hash"`
							ValidatorsHash     string `json:"validators_hash"`
							NextValidatorsHash string `json:"next_validators_hash"`
							ConsensusHash      string `json:"consensus_hash"`
							AppHash            string `json:"app_hash"`
							LastResultsHash    string `json:"last_results_hash"`
							EvidenceHash       string `json:"evidence_hash"`
							ProposerAddress    string `json:"proposer_address"`
						} `json:"header"`
					} `json:"signed_header"`
				} `json:"header"`
				Value string `json:"value"`
			} `json:"messages"`
			Memo             string `json:"memo"`
			TimeoutHeight    string `json:"timeout_height"`
			ExtensionOptions []struct {
				TypeURL string `json:"type_url"`
				Value   string `json:"value"`
			} `json:"extension_options"`
			NonCriticalExtensionOptions []struct {
				TypeURL string `json:"type_url"`
				Value   string `json:"value"`
			} `json:"non_critical_extension_options"`
		} `json:"body"`
		AuthInfo struct {
			SignerInfos []struct {
				PublicKey struct {
					TypeURL string `json:"type_url"`
					Value   string `json:"value"`
				} `json:"public_key"`
				ModeInfo struct {
					Single struct {
						Mode string `json:"mode"`
					} `json:"single"`
					Multi struct {
						Bitarray struct {
							ExtraBitsStored int    `json:"extra_bits_stored"`
							Elems           string `json:"elems"`
						} `json:"bitarray"`
						ModeInfos []interface{} `json:"mode_infos"`
					} `json:"multi"`
				} `json:"mode_info"`
				Sequence string `json:"sequence"`
			} `json:"signer_infos"`
			Fee struct {
				Amount []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"amount"`
				GasLimit string `json:"gas_limit"`
				Payer    string `json:"payer"`
				Granter  string `json:"granter"`
			} `json:"fee"`
		} `json:"auth_info"`
		Signatures []string `json:"signatures"`
	} `json:"txs"`
	TxResponses []struct {
		Height    string `json:"height"`
		Txhash    string `json:"txhash"`
		Codespace string `json:"codespace"`
		Code      int    `json:"code"`
		Data      string `json:"data"`
		RawLog    string `json:"raw_log"`
		Logs      []struct {
			MsgIndex int    `json:"msg_index"`
			Log      string `json:"log"`
			Events   []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"logs"`
		Info      string `json:"info"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Tx        struct {
			TypeURL string `json:"type_url"`
			Body    struct {
				Messages []struct {
					TypeURL string `json:"@type"`
					Header  struct {
						Type         string `json:"@type"`
						SignedHeader struct {
							Header struct {
								Version struct {
									Block string `json:"block"`
									App   string `json:"app"`
								} `json:"version"`
								ChainID     string    `json:"chain_id"`
								Height      string    `json:"height"`
								Time        time.Time `json:"time"`
								LastBlockID struct {
									Hash          string `json:"hash"`
									PartSetHeader struct {
										Total int    `json:"total"`
										Hash  string `json:"hash"`
									} `json:"part_set_header"`
								} `json:"last_block_id"`
								LastCommitHash     string `json:"last_commit_hash"`
								DataHash           string `json:"data_hash"`
								ValidatorsHash     string `json:"validators_hash"`
								NextValidatorsHash string `json:"next_validators_hash"`
								ConsensusHash      string `json:"consensus_hash"`
								AppHash            string `json:"app_hash"`
								LastResultsHash    string `json:"last_results_hash"`
								EvidenceHash       string `json:"evidence_hash"`
								ProposerAddress    string `json:"proposer_address"`
							} `json:"header"`
						} `json:"signed_header"`
					} `json:"header"`
				} `json:"messages"`
				Memo             string `json:"memo"`
				TimeoutHeight    string `json:"timeout_height"`
				ExtensionOptions []struct {
					TypeURL string `json:"type_url"`
					Value   string `json:"value"`
				} `json:"extension_options"`
				NonCriticalExtensionOptions []struct {
					TypeURL string `json:"type_url"`
					Value   string `json:"value"`
				} `json:"non_critical_extension_options"`
			} `json:"body"`
			Value string `json:"value"`
		} `json:"tx"`
		Timestamp string `json:"timestamp"`
	} `json:"tx_responses"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}

type Evidence struct {
	ChainName string
	Heights   []string
}
