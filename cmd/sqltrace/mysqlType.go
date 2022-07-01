package main

type INNODB_TRX struct {
	TrxId                   string `orm:"trx_id" json:"trx_id"`
	TrxState                string `orm:"trx_state" json:"trx_state"`
	TrxStarted              string `orm:"trx_started" json:"trx_started"`
	TrxRequestedLockId      string `orm:"trx_requested_lock_id" json:"trx_requested_lock_id"`
	TrxWaitStarted          string `orm:"trx_wait_started" json:"trx_wait_started"`
	TrxWeight               int    `orm:"trx_weight" json:"trx_weight"`
	TrxMysqlThreadId        int    `orm:"trx_mysql_thread_id" json:"trx_mysql_thread_id"`
	TrxQuery                string `orm:"trx_query" json:"trx_query"`
	TrxOperationState       string `orm:"trx_operation_state" json:"trx_operation_state"`
	TrxTablesInUse          int    `orm:"trx_tables_in_use" json:"trx_tables_in_use"`
	TrxTablesLocked         int    `orm:"trx_tables_locked" json:"trx_tables_locked"`
	TrxLockStructs          int    `orm:"trx_lock_structs" json:"trx_lock_structs"`
	TrxLockMemoryBytes      int    `orm:"trx_lock_memory_bytes" json:"trx_lock_memory_bytes"`
	TrxRowsLocked           int    `orm:"trx_rows_locked" json:"trx_rows_locked"`
	TrxRowsModified         int    `orm:"trx_rows_modified" json:"trx_rows_modified"`
	TrxConcurrencyTickets   int    `orm:"trx_concurrency_tickets" json:"trx_concurrency_tickets"`
	TrxIsolationLevel       string `orm:"trx_isolation_level" json:"trx_isolation_level"`
	TrxUniqueChecks         int    `orm:"trx_unique_checks" json:"trx_unique_checks"`
	TrxForeignKeyChecks     int    `orm:"trx_foreign_key_checks" json:"trx_foreign_key_checks"`
	TrxLastForeignKeyError  string `orm:"trx_last_foreign_key_error" json:"trx_last_foreign_key_error"`
	TrxAdaptiveHashLatched  int    `orm:"trx_adaptive_hash_latched" json:"trx_adaptive_hash_latched"`
	TrxAdaptiveHashTimeout  int    `orm:"trx_adaptive_hash_timeout" json:"trx_adaptive_hash_timeout"`
	TrxIsReadOnly           int    `orm:"trx_is_read_only" json:"trx_is_read_only"`
	TrxAutocommitNonLocking int    `orm:"trx_autocommit_non_locking" json:"trx_autocommit_non_locking"`
}

type INNODB_LOCKS struct {
	LockId    string `orm:"lock_id" json:"lock_id"`
	LockTrxId string `orm:"lock_trx_id" json:"lock_trx_id"`
	LockMode  string `orm:"lock_mode" json:"lock_mode"`
	LockType  string `orm:"lock_type" json:"lock_type"`
	LockTable string `orm:"lock_table" json:"lock_table"`
	LockIndex string `orm:"lock_index" json:"lock_index"`
	LockSpace int    `orm:"lock_space" json:"lock_space"`
	LockPage  int    `orm:"lock_page" json:"lock_page"`
	LockRec   int    `orm:"lock_rec" json:"lock_rec"`
	LockData  string `orm:"lock_data" json:"lock_data"`
}

type PROCESSLIST struct {
	ID      int    `orm:"ID" json:"ID"`
	USER    string `orm:"USER" json:"USER"`
	HOST    string `orm:"HOST" json:"HOST"`
	DB      string `orm:"DB" json:"DB"`
	COMMAND string `orm:"COMMAND" json:"COMMAND"`
	TIME    int    `orm:"TIME" json:"TIME"`
	STATE   string `orm:"STATE" json:"STATE"`
	INFO    string `orm:"INFO" json:"INFO"`
}

type INNODB_LOCK_WAITS struct {
	RequestingTrxId string `orm:"requesting_trx_id" json:"requesting_trx_id"`
	RequestedLockId string `orm:"requested_lock_id" json:"requested_lock_id"`
	BlockingTrxId   string `orm:"blocking_trx_id" json:"blocking_trx_id"`
	BlockingLockId  string `orm:"blocking_lock_id" json:"blocking_lock_id"`
}
