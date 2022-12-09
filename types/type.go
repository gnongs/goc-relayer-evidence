package types

import "time"

type ReturnTxsBlock struct {
	Txs []struct {
		Body struct {
			Messages []struct {
				Type     string `json:"@type"`
				ClientID string `json:"client_id,omitempty"`
				Header   struct {
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
						Commit struct {
							Height  string `json:"height"`
							Round   int    `json:"round"`
							BlockID struct {
								Hash          string `json:"hash"`
								PartSetHeader struct {
									Total int    `json:"total"`
									Hash  string `json:"hash"`
								} `json:"part_set_header"`
							} `json:"block_id"`
							Signatures []struct {
								BlockIDFlag      string    `json:"block_id_flag"`
								ValidatorAddress string    `json:"validator_address"`
								Timestamp        time.Time `json:"timestamp"`
								Signature        string    `json:"signature"`
							} `json:"signatures"`
						} `json:"commit"`
					} `json:"signed_header"`
					ValidatorSet struct {
						Validators []struct {
							Address string `json:"address"`
							PubKey  struct {
								Ed25519 string `json:"ed25519"`
							} `json:"pub_key"`
							VotingPower      string `json:"voting_power"`
							ProposerPriority string `json:"proposer_priority"`
						} `json:"validators"`
						Proposer struct {
							Address string `json:"address"`
							PubKey  struct {
								Ed25519 string `json:"ed25519"`
							} `json:"pub_key"`
							VotingPower      string `json:"voting_power"`
							ProposerPriority string `json:"proposer_priority"`
						} `json:"proposer"`
						TotalVotingPower string `json:"total_voting_power"`
					} `json:"validator_set"`
					TrustedHeight struct {
						RevisionNumber string `json:"revision_number"`
						RevisionHeight string `json:"revision_height"`
					} `json:"trusted_height"`
					TrustedValidators struct {
						Validators []struct {
							Address string `json:"address"`
							PubKey  struct {
								Ed25519 string `json:"ed25519"`
							} `json:"pub_key"`
							VotingPower      string `json:"voting_power"`
							ProposerPriority string `json:"proposer_priority"`
						} `json:"validators"`
						Proposer struct {
							Address string `json:"address"`
							PubKey  struct {
								Ed25519 string `json:"ed25519"`
							} `json:"pub_key"`
							VotingPower      string `json:"voting_power"`
							ProposerPriority string `json:"proposer_priority"`
						} `json:"proposer"`
						TotalVotingPower string `json:"total_voting_power"`
					} `json:"trusted_validators"`
				} `json:"header,omitempty"`
				Signer string `json:"signer"`
				Packet struct {
					Sequence           string `json:"sequence"`
					SourcePort         string `json:"source_port"`
					SourceChannel      string `json:"source_channel"`
					DestinationPort    string `json:"destination_port"`
					DestinationChannel string `json:"destination_channel"`
					Data               string `json:"data"`
					TimeoutHeight      struct {
						RevisionNumber string `json:"revision_number"`
						RevisionHeight string `json:"revision_height"`
					} `json:"timeout_height"`
					TimeoutTimestamp string `json:"timeout_timestamp"`
				} `json:"packet,omitempty"`
				ProofCommitment string `json:"proof_commitment,omitempty"`
				ProofHeight     struct {
					RevisionNumber string `json:"revision_number"`
					RevisionHeight string `json:"revision_height"`
				} `json:"proof_height,omitempty"`
			} `json:"messages"`
			Memo                        string        `json:"memo"`
			TimeoutHeight               string        `json:"timeout_height"`
			ExtensionOptions            []interface{} `json:"extension_options"`
			NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
		} `json:"body"`
		AuthInfo struct {
			SignerInfos []struct {
				PublicKey struct {
					Type string `json:"@type"`
					Key  string `json:"key"`
				} `json:"public_key"`
				ModeInfo struct {
					Single struct {
						Mode string `json:"mode"`
					} `json:"single"`
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
	BlockID struct {
		Hash          string `json:"hash"`
		PartSetHeader struct {
			Total int    `json:"total"`
			Hash  string `json:"hash"`
		} `json:"part_set_header"`
	} `json:"block_id"`
	Block struct {
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
		Data struct {
			Txs []string `json:"txs"`
		} `json:"data"`
		Evidence struct {
			Evidence []interface{} `json:"evidence"`
		} `json:"evidence"`
		LastCommit struct {
			Height  string `json:"height"`
			Round   int    `json:"round"`
			BlockID struct {
				Hash          string `json:"hash"`
				PartSetHeader struct {
					Total int    `json:"total"`
					Hash  string `json:"hash"`
				} `json:"part_set_header"`
			} `json:"block_id"`
			Signatures []struct {
				BlockIDFlag      string    `json:"block_id_flag"`
				ValidatorAddress string    `json:"validator_address"`
				Timestamp        time.Time `json:"timestamp"`
				Signature        string    `json:"signature"`
			} `json:"signatures"`
		} `json:"last_commit"`
	} `json:"block"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

type RetrunTxs struct {
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
	ChainName       string
	TotalHeights    []string
	VerifiedHeights []string
}
