DROP TABLE {{ .Schema }}.receipts
DROP TABLE {{ .Schema }}.tx_attempts
DROP TABLE {{ .Schema }}.upkeep_states
DROP TABLE {{ .Schema }}.txes
DROP TABLE {{ .Schema }}.logs
DROP TABLE {{ .Schema }}.log_poller_filters
DROP TABLE {{ .Schema }}.log_poller_blocks
DROP TABLE {{ .Schema }}.key_states;
DROP TABLE {{ .Schema }}.heads;
DROP TABLE {{ .Schema }}.forwarders;