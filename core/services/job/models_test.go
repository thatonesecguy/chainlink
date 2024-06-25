package job

import (
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pelletier/go-toml/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/codec"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	pkgworkflows "github.com/smartcontractkit/chainlink-common/pkg/workflows"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestJob_RelayIdentifier(t *testing.T) {
	type fields struct {
		Relay       string
		ChainID     string
		RelayConfig JSONConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.RelayID
		wantErr bool
	}{
		{name: "err no chain id",
			fields:  fields{},
			want:    types.RelayID{},
			wantErr: true,
		},
		{
			name: "evm explicitly configured",
			fields: fields{
				Relay:   relay.NetworkEVM,
				ChainID: "1",
			},
			want: types.RelayID{Network: relay.NetworkEVM, ChainID: "1"},
		},
		{
			name: "evm implicitly configured",
			fields: fields{
				Relay:       relay.NetworkEVM,
				RelayConfig: map[string]any{"chainID": 1},
			},
			want: types.RelayID{Network: relay.NetworkEVM, ChainID: "1"},
		},
		{
			name: "evm implicitly configured with bad value",
			fields: fields{
				Relay:       relay.NetworkEVM,
				RelayConfig: map[string]any{"chainID": float32(1)},
			},
			want:    types.RelayID{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			j := &Job{
				Relay:       tt.fields.Relay,
				ChainID:     tt.fields.ChainID,
				RelayConfig: tt.fields.RelayConfig,
			}
			got, err := j.RelayID()
			if (err != nil) != tt.wantErr {
				t.Errorf("OCR2OracleSpec.RelayIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OCR2OracleSpec.RelayIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	//go:embed testdata/ocr2spec_compact.toml
	specCompact string
	//go:embed testdata/ocr2spec_pretty.toml
	specPretty string
	//go:embed testdata/job_compact.toml
	jobCompact string
	//go:embed testdata/job_pretty.toml
	jobPretty string
)

func TestOCR2OracleSpec(t *testing.T) {
	val := OCR2OracleSpec{
		Relay:                             relay.NetworkEVM,
		PluginType:                        types.Median,
		ContractID:                        "foo",
		OCRKeyBundleID:                    null.StringFrom("bar"),
		TransmitterID:                     null.StringFrom("baz"),
		ContractConfigConfirmations:       1,
		ContractConfigTrackerPollInterval: *models.NewInterval(time.Second),
		RelayConfig: map[string]interface{}{
			"chainID":   1337,
			"fromBlock": 42,
			"chainReader": evmtypes.ChainReaderConfig{
				Contracts: map[string]evmtypes.ChainContractReader{
					"median": {
						ContractABI: `[
	 {
	   "anonymous": false,
	   "inputs": [
	     {
	       "indexed": true,
	       "internalType": "address",
	       "name": "requester",
	       "type": "address"
	     },
	     {
	       "indexed": false,
	       "internalType": "bytes32",
	       "name": "configDigest",
	       "type": "bytes32"
	     },
	     {
	       "indexed": false,
	       "internalType": "uint32",
	       "name": "epoch",
	       "type": "uint32"
	     },
	     {
	       "indexed": false,
	       "internalType": "uint8",
	       "name": "round",
	       "type": "uint8"
	     }
	   ],
	   "name": "RoundRequested",
	   "type": "event"
	 },
	 {
	   "inputs": [],
	   "name": "latestTransmissionDetails",
	   "outputs": [
	     {
	       "internalType": "bytes32",
	       "name": "configDigest",
	       "type": "bytes32"
	     },
	     {
	       "internalType": "uint32",
	       "name": "epoch",
	       "type": "uint32"
	     },
	     {
	       "internalType": "uint8",
	       "name": "round",
	       "type": "uint8"
	     },
	     {
	       "internalType": "int192",
	       "name": "latestAnswer_",
	       "type": "int192"
	     },
	     {
	       "internalType": "uint64",
	       "name": "latestTimestamp_",
	       "type": "uint64"
	     }
	   ],
	   "stateMutability": "view",
	   "type": "function"
	 }
	]
	`,
						Configs: map[string]*evmtypes.ChainReaderDefinition{
							"LatestTransmissionDetails": {
								ChainSpecificName: "latestTransmissionDetails",
								OutputModifications: codec.ModifiersConfig{
									&codec.EpochToTimeModifierConfig{
										Fields: []string{"LatestTimestamp_"},
									},
									&codec.RenameModifierConfig{
										Fields: map[string]string{
											"LatestAnswer_":    "LatestAnswer",
											"LatestTimestamp_": "LatestTimestamp",
										},
									},
								},
							},
							"LatestRoundRequested": {
								ChainSpecificName: "RoundRequested",
								ReadType:          evmtypes.Event,
							},
						},
					},
				},
			},
			"codec": evmtypes.CodecConfig{
				Configs: map[string]evmtypes.ChainCodecConfig{
					"MedianReport": {
						TypeABI: `[
	 {
	   "Name": "Timestamp",
	   "Type": "uint32"
	 },
	 {
	   "Name": "Observers",
	   "Type": "bytes32"
	 },
	 {
	   "Name": "Observations",
	   "Type": "int192[]"
	 },
	 {
	   "Name": "JuelsPerFeeCoin",
	   "Type": "int192"
	 }
	]
	`,
					},
				},
			},
		},
	}

	t.Run("marshal", func(t *testing.T) {
		gotB, err := toml.Marshal(jb)
		require.NoError(t, err)
		t.Log("marshaled:", string(gotB))
		require.Equal(t, jobCompact, string(gotB))
	})

	t.Run("round-trip", func(t *testing.T) {
		var gotVal Job
		require.NoError(t, toml.Unmarshal([]byte(jobCompact), &gotVal))
		gotB, err := toml.Marshal(gotVal)
		require.NoError(t, err)
		require.Equal(t, jobCompact, string(gotB))
		t.Run("specPretty", func(t *testing.T) {
			var gotVal Job
			require.NoError(t, toml.Unmarshal([]byte(jobPretty), &gotVal))
			gotB, err := toml.Marshal(gotVal)
			require.NoError(t, err)
			t.Log("marshaled specCompact:", string(gotB))
			require.Equal(t, jobCompact, string(gotB))
		})
	})
}

func TestJobRelayConfig(t *testing.T) {
	jbUUID, err := uuid.Parse("ee50cdf7-5d72-4d90-8def-06bb42b932ef")
	require.NoError(t, err)
	jb := Job{
		ExternalJobID:     jbUUID,
		PipelineSpecID:    999,
		JobSpecErrors:     []SpecError{},
		Type:              "",
		SchemaVersion:     0,
		GasLimit:          clnull.Uint32{Uint32: 1234, Valid: true},
		ForwardingAllowed: false,
		Name:              null.StringFrom("some name"),
		MaxTaskDuration:   models.Interval(time.Second * 2),
		CreatedAt:         time.Unix(1718894576, 0), // 20.06.2024 14:42:56 GMT
		ChainID:           "323211",

		Relay: types.NetworkEVM,
		RelayConfig: map[string]interface{}{
			"chainID":   1337,
			"fromBlock": 42,
			"chainReader": evmtypes.ChainReaderConfig{
				Contracts: map[string]evmtypes.ChainContractReader{
					"median": {
						ContractABI: `[
	 {
	   "anonymous": false,
	   "inputs": [
	     {
	       "indexed": true,
	       "internalType": "address",
	       "name": "requester",
	       "type": "address"
	     },
	     {
	       "indexed": false,
	       "internalType": "bytes32",
	       "name": "configDigest",
	       "type": "bytes32"
	     },
	     {
	       "indexed": false,
	       "internalType": "uint32",
	       "name": "epoch",
	       "type": "uint32"
	     },
	     {
	       "indexed": false,
	       "internalType": "uint8",
	       "name": "round",
	       "type": "uint8"
	     }
	   ],
	   "name": "RoundRequested",
	   "type": "event"
	 },
	 {
	   "inputs": [],
	   "name": "latestTransmissionDetails",
	   "outputs": [
	     {
	       "internalType": "bytes32",
	       "name": "configDigest",
	       "type": "bytes32"
	     },
	     {
	       "internalType": "uint32",
	       "name": "epoch",
	       "type": "uint32"
	     },
	     {
	       "internalType": "uint8",
	       "name": "round",
	       "type": "uint8"
	     },
	     {
	       "internalType": "int192",
	       "name": "latestAnswer_",
	       "type": "int192"
	     },
	     {
	       "internalType": "uint64",
	       "name": "latestTimestamp_",
	       "type": "uint64"
	     }
	   ],
	   "stateMutability": "view",
	   "type": "function"
	 }
	]
	`,
						Configs: map[string]*evmtypes.ChainReaderDefinition{
							"LatestTransmissionDetails": {
								ChainSpecificName: "latestTransmissionDetails",
								OutputModifications: codec.ModifiersConfig{
									&codec.EpochToTimeModifierConfig{
										Fields: []string{"LatestTimestamp_"},
									},
									&codec.RenameModifierConfig{
										Fields: map[string]string{
											"LatestAnswer_":    "LatestAnswer",
											"LatestTimestamp_": "LatestTimestamp",
										},
									},
								},
							},
							"LatestRoundRequested": {
								ChainSpecificName: "RoundRequested",
								ReadType:          evmtypes.Event,
							},
						},
					},
				},
			},
			"codec": evmtypes.CodecConfig{
				Configs: map[string]evmtypes.ChainCodecConfig{
					"MedianReport": {
						TypeABI: `[
	 {
	   "Name": "Timestamp",
	   "Type": "uint32"
	 },
	 {
	   "Name": "Observers",
	   "Type": "bytes32"
	 },
	 {
	   "Name": "Observations",
	   "Type": "int192[]"
	 },
	 {
	   "Name": "JuelsPerFeeCoin",
	   "Type": "int192"
	 }
	]
	`,
					},
				},
			},
		},
	}

	t.Run("marshal", func(t *testing.T) {
		gotB, err := toml.Marshal(jb)
		require.NoError(t, err)
		t.Log("marshaled:", string(gotB))
		require.Equal(t, jobCompact, string(gotB))
	})

	t.Run("round-trip", func(t *testing.T) {
		var gotVal Job
		require.NoError(t, toml.Unmarshal([]byte(jobCompact), &gotVal))
		gotB, err := toml.Marshal(gotVal)
		require.NoError(t, err)
		require.Equal(t, jobCompact, string(gotB))
		t.Run("specPretty", func(t *testing.T) {
			var gotVal Job
			require.NoError(t, toml.Unmarshal([]byte(jobPretty), &gotVal))
			gotB, err := toml.Marshal(gotVal)
			require.NoError(t, err)
			t.Log("marshaled specCompact:", string(gotB))
			require.Equal(t, jobCompact, string(gotB))
		})
	})
}

func TestOCR2OracleSpec(t *testing.T) {
	val := OCR2OracleSpec{
		PluginType:                        types.Median,
		ContractID:                        "foo",
		OCRKeyBundleID:                    null.StringFrom("bar"),
		TransmitterID:                     null.StringFrom("baz"),
		ContractConfigConfirmations:       1,
		ContractConfigTrackerPollInterval: *models.NewInterval(time.Second),

		OnchainSigningStrategy: map[string]interface{}{
			"strategyName": "single-chain",
			"config": map[string]interface{}{
				"evm":       "",
				"publicKey": "0xdeadbeef",
			},
		},
		PluginConfig: map[string]interface{}{"juelsPerFeeCoinSource": `  // data source 1
  ds1          [type=bridge name="%s"];
  ds1_parse    [type=jsonparse path="data"];
  ds1_multiply [type=multiply times=2];

  // data source 2
  ds2          [type=http method=GET url="%s"];
  ds2_parse    [type=jsonparse path="data"];
  ds2_multiply [type=multiply times=2];

  ds1 -> ds1_parse -> ds1_multiply -> answer1;
  ds2 -> ds2_parse -> ds2_multiply -> answer1;

  answer1 [type=median index=0];
`,
		},
	}

	t.Run("marshal", func(t *testing.T) {
		gotB, err := toml.Marshal(val)
		require.NoError(t, err)
		t.Log("marshaled:", string(gotB))
		require.Equal(t, specCompact, string(gotB))
	})

	t.Run("round-trip", func(t *testing.T) {
		var gotVal OCR2OracleSpec
		require.NoError(t, toml.Unmarshal([]byte(specCompact), &gotVal))
		gotB, err := toml.Marshal(gotVal)
		require.NoError(t, err)
		require.Equal(t, specCompact, string(gotB))
		t.Run("specPretty", func(t *testing.T) {
			var gotVal OCR2OracleSpec
			require.NoError(t, toml.Unmarshal([]byte(specPretty), &gotVal))
			gotB, err := toml.Marshal(gotVal)
			require.NoError(t, err)
			t.Log("marshaled specCompact:", string(gotB))
			require.Equal(t, specCompact, string(gotB))
		})
	})
}

func TestWorkflowSpec_Validate(t *testing.T) {
	type fields struct {
		Workflow string
	}
	tests := []struct {
		name              string
		fields            fields
		wantWorkflowOwner string
		wantWorkflowName  string

		wantError bool
	}{
		{
			name: "valid",
			fields: fields{
				Workflow: pkgworkflows.WFYamlSpec(t, "workflow01", "0x0123456789012345678901234567890123456789"),
			},
			wantWorkflowOwner: "0123456789012345678901234567890123456789", // the workflow job spec strips the 0x prefix to limit to 40	characters
			wantWorkflowName:  "workflow01",
		},
		{
			name: "valid no name",
			fields: fields{
				Workflow: pkgworkflows.WFYamlSpec(t, "", "0x0123456789012345678901234567890123456789"),
			},
			wantWorkflowOwner: "0123456789012345678901234567890123456789", // the workflow job spec strips the 0x prefix to limit to 40	characters
			wantWorkflowName:  "",
		},
		{
			name: "valid no owner",
			fields: fields{
				Workflow: pkgworkflows.WFYamlSpec(t, "workflow01", ""),
			},
			wantWorkflowOwner: "",
			wantWorkflowName:  "workflow01",
		},
		{
			name: "invalid ",
			fields: fields{
				Workflow: "garbage",
			},
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WorkflowSpec{
				Workflow: tt.fields.Workflow,
			}
			err := w.Validate()
			require.Equal(t, tt.wantError, err != nil)
			if !tt.wantError {
				assert.NotEmpty(t, w.WorkflowID)
				assert.Equal(t, tt.wantWorkflowOwner, w.WorkflowOwner)
				assert.Equal(t, tt.wantWorkflowName, w.WorkflowName)
			}
		})
	}
}
