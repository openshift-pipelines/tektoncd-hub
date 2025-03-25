package proto

// Subset of possible errors.
const (
	ErrUnsupportedMethod                            Error = 1
	ErrUnsupportedParameter                         Error = 2
	ErrUnexpectedEndOfFile                          Error = 3
	ErrExpectedEndOfFile                            Error = 4
	ErrCannotParseText                              Error = 6
	ErrIncorrectNumberOfColumns                     Error = 7
	ErrThereIsNoColumn                              Error = 8
	ErrSizesOfColumnsDoesntMatch                    Error = 9
	ErrNotFoundColumnInBlock                        Error = 10
	ErrPositionOutOfBound                           Error = 11
	ErrParameterOutOfBound                          Error = 12
	ErrSizesOfColumnsInTupleDoesntMatch             Error = 13
	ErrDuplicateColumn                              Error = 15
	ErrNoSuchColumnInTable                          Error = 16
	ErrDelimiterInStringLiteralDoesntMatch          Error = 17
	ErrCannotInsertElementIntoConstantColumn        Error = 18
	ErrSizeOfFixedStringDoesntMatch                 Error = 19
	ErrNumberOfColumnsDoesntMatch                   Error = 20
	ErrCannotReadAllDataFromTabSeparatedInput       Error = 21
	ErrCannotParseAllValueFromTabSeparatedInput     Error = 22
	ErrCannotReadFromIstream                        Error = 23
	ErrCannotWriteToOstream                         Error = 24
	ErrCannotParseEscapeSequence                    Error = 25
	ErrCannotParseQuotedString                      Error = 26
	ErrCannotParseInputAssertionFailed              Error = 27
	ErrCannotPrintFloatOrDoubleNumber               Error = 28
	ErrCannotPrintInteger                           Error = 29
	ErrCannotReadSizeOfCompressedChunk              Error = 30
	ErrCannotReadCompressedChunk                    Error = 31
	ErrAttemptToReadAfterEOF                        Error = 32
	ErrCannotReadAllData                            Error = 33
	ErrTooManyArgumentsForFunction                  Error = 34
	ErrTooLessArgumentsForFunction                  Error = 35
	ErrBadArguments                                 Error = 36
	ErrUnknownElementInAst                          Error = 37
	ErrCannotParseDate                              Error = 38
	ErrTooLargeSizeCompressed                       Error = 39
	ErrChecksumDoesntMatch                          Error = 40
	ErrCannotParseDatetime                          Error = 41
	ErrNumberOfArgumentsDoesntMatch                 Error = 42
	ErrIllegalTypeOfArgument                        Error = 43
	ErrIllegalColumn                                Error = 44
	ErrIllegalNumberOfResultColumns                 Error = 45
	ErrUnknownFunction                              Error = 46
	ErrUnknownIdentifier                            Error = 47
	ErrNotImplemented                               Error = 48
	ErrLogicalError                                 Error = 49
	ErrUnknownType                                  Error = 50
	ErrEmptyListOfColumnsQueried                    Error = 51
	ErrColumnQueriedMoreThanOnce                    Error = 52
	ErrTypeMismatch                                 Error = 53
	ErrStorageDoesntAllowParameters                 Error = 54
	ErrStorageRequiresParameter                     Error = 55
	ErrUnknownStorage                               Error = 56
	ErrTableAlreadyExists                           Error = 57
	ErrTableMetadataAlreadyExists                   Error = 58
	ErrIllegalTypeOfColumnForFilter                 Error = 59
	ErrUnknownTable                                 Error = 60
	ErrOnlyFilterColumnInBlock                      Error = 61
	ErrSyntaxError                                  Error = 62
	ErrUnknownAggregateFunction                     Error = 63
	ErrCannotReadAggregateFunctionFromText          Error = 64
	ErrCannotWriteAggregateFunctionAsText           Error = 65
	ErrNotAColumn                                   Error = 66
	ErrIllegalKeyOfAggregation                      Error = 67
	ErrCannotGetSizeOfField                         Error = 68
	ErrArgumentOutOfBound                           Error = 69
	ErrCannotConvertType                            Error = 70
	ErrCannotWriteAfterEndOfBuffer                  Error = 71
	ErrCannotParseNumber                            Error = 72
	ErrUnknownFormat                                Error = 73
	ErrCannotReadFromFileDescriptor                 Error = 74
	ErrCannotWriteToFileDescriptor                  Error = 75
	ErrCannotOpenFile                               Error = 76
	ErrCannotCloseFile                              Error = 77
	ErrUnknownTypeOfQuery                           Error = 78
	ErrIncorrectFileName                            Error = 79
	ErrIncorrectQuery                               Error = 80
	ErrUnknownDatabase                              Error = 81
	ErrDatabaseAlreadyExists                        Error = 82
	ErrDirectoryDoesntExist                         Error = 83
	ErrDirectoryAlreadyExists                       Error = 84
	ErrFormatIsNotSuitableForInput                  Error = 85
	ErrReceivedErrorFromRemoteIoServer              Error = 86
	ErrCannotSeekThroughFile                        Error = 87
	ErrCannotTruncateFile                           Error = 88
	ErrUnknownCompressionMethod                     Error = 89
	ErrEmptyListOfColumnsPassed                     Error = 90
	ErrSizesOfMarksFilesAreInconsistent             Error = 91
	ErrEmptyDataPassed                              Error = 92
	ErrUnknownAggregatedDataVariant                 Error = 93
	ErrCannotMergeDifferentAggregatedDataVariants   Error = 94
	ErrCannotReadFromSocket                         Error = 95
	ErrCannotWriteToSocket                          Error = 96
	ErrCannotReadAllDataFromChunkedInput            Error = 97
	ErrCannotWriteToEmptyBlockOutputStream          Error = 98
	ErrUnknownPacketFromClient                      Error = 99
	ErrUnknownPacketFromServer                      Error = 100
	ErrUnexpectedPacketFromClient                   Error = 101
	ErrUnexpectedPacketFromServer                   Error = 102
	ErrReceivedDataForWrongQueryID                  Error = 103
	ErrTooSmallBufferSize                           Error = 104
	ErrCannotReadHistory                            Error = 105
	ErrCannotAppendHistory                          Error = 106
	ErrFileDoesntExist                              Error = 107
	ErrNoDataToInsert                               Error = 108
	ErrCannotBlockSignal                            Error = 109
	ErrCannotUnblockSignal                          Error = 110
	ErrCannotManipulateSigset                       Error = 111
	ErrCannotWaitForSignal                          Error = 112
	ErrThereIsNoSession                             Error = 113
	ErrCannotClockGettime                           Error = 114
	ErrUnknownSetting                               Error = 115
	ErrThereIsNoDefaultValue                        Error = 116
	ErrIncorrectData                                Error = 117
	ErrEngineRequired                               Error = 119
	ErrCannotInsertValueOfDifferentSizeIntoTuple    Error = 120
	ErrUnknownSetDataVariant                        Error = 121
	ErrIncompatibleColumns                          Error = 122
	ErrUnknownTypeOfAstNode                         Error = 123
	ErrIncorrectElementOfSet                        Error = 124
	ErrIncorrectResultOfScalarSubquery              Error = 125
	ErrCannotGetReturnType                          Error = 126
	ErrIllegalIndex                                 Error = 127
	ErrTooLargeArraySize                            Error = 128
	ErrFunctionIsSpecial                            Error = 129
	ErrCannotReadArrayFromText                      Error = 130
	ErrTooLargeStringSize                           Error = 131
	ErrCannotCreateTableFromMetadata                Error = 132
	ErrAggregateFunctionDoesntAllowParameters       Error = 133
	ErrParametersToAggregateFunctionsMustBeLiterals Error = 134
	ErrZeroArrayOrTupleIndex                        Error = 135
	ErrUnknownElementInConfig                       Error = 137
	ErrExcessiveElementInConfig                     Error = 138
	ErrNoElementsInConfig                           Error = 139
	ErrAllRequestedColumnsAreMissing                Error = 140
	ErrSamplingNotSupported                         Error = 141
	ErrNotFoundNode                                 Error = 142
	ErrFoundMoreThanOneNode                         Error = 143
	ErrFirstDateIsBiggerThanLastDate                Error = 144
	ErrUnknownOverflowMode                          Error = 145
	ErrQuerySectionDoesntMakeSense                  Error = 146
	ErrNotFoundFunctionElementForAggregate          Error = 147
	ErrNotFoundRelationElementForCondition          Error = 148
	ErrNotFoundRHSElementForCondition               Error = 149
	ErrNoAttributesListed                           Error = 150
	ErrIndexOfColumnInSortClauseIsOutOfRange        Error = 151
	ErrUnknownDirectionOfSorting                    Error = 152
	ErrIllegalDivision                              Error = 153
	ErrAggregateFunctionNotApplicable               Error = 154
	ErrUnknownRelation                              Error = 155
	ErrDictionariesWasNotLoaded                     Error = 156
	ErrIllegalOverflowMode                          Error = 157
	ErrTooManyRows                                  Error = 158
	ErrTimeoutExceeded                              Error = 159
	ErrTooSlow                                      Error = 160
	ErrTooManyColumns                               Error = 161
	ErrTooDeepSubqueries                            Error = 162
	ErrTooDeepPipeline                              Error = 163
	ErrReadonly                                     Error = 164
	ErrTooManyTemporaryColumns                      Error = 165
	ErrTooManyTemporaryNonConstColumns              Error = 166
	ErrTooDeepAst                                   Error = 167
	ErrTooBigAst                                    Error = 168
	ErrBadTypeOfField                               Error = 169
	ErrBadGet                                       Error = 170
	ErrBlocksHaveDifferentStructure                 Error = 171
	ErrCannotCreateDirectory                        Error = 172
	ErrCannotAllocateMemory                         Error = 173
	ErrCyclicAliases                                Error = 174
	ErrChunkNotFound                                Error = 176
	ErrDuplicateChunkName                           Error = 177
	ErrMultipleAliasesForExpression                 Error = 178
	ErrMultipleExpressionsForAlias                  Error = 179
	ErrThereIsNoProfile                             Error = 180
	ErrIllegalFinal                                 Error = 181
	ErrIllegalPrewhere                              Error = 182
	ErrUnexpectedExpression                         Error = 183
	ErrIllegalAggregation                           Error = 184
	ErrUnsupportedMyisamBlockType                   Error = 185
	ErrUnsupportedCollationLocale                   Error = 186
	ErrCollationComparisonFailed                    Error = 187
	ErrUnknownAction                                Error = 188
	ErrTableMustNotBeCreatedManually                Error = 189
	ErrSizesOfArraysDoesntMatch                     Error = 190
	ErrSetSizeLimitExceeded                         Error = 191
	ErrUnknownUser                                  Error = 192
	ErrWrongPassword                                Error = 193
	ErrRequiredPassword                             Error = 194
	ErrIPAddressNotAllowed                          Error = 195
	ErrUnknownAddressPatternType                    Error = 196
	ErrServerRevisionIsTooOld                       Error = 197
	ErrDNSError                                     Error = 198
	ErrUnknownQuota                                 Error = 199
	ErrQuotaDoesntAllowKeys                         Error = 200
	ErrQuotaExpired                                 Error = 201
	ErrTooManySimultaneousQueries                   Error = 202
	ErrNoFreeConnection                             Error = 203
	ErrCannotFsync                                  Error = 204
	ErrNestedTypeTooDeep                            Error = 205
	ErrAliasRequired                                Error = 206
	ErrAmbiguousIdentifier                          Error = 207
	ErrEmptyNestedTable                             Error = 208
	ErrSocketTimeout                                Error = 209
	ErrNetworkError                                 Error = 210
	ErrEmptyQuery                                   Error = 211
	ErrUnknownLoadBalancing                         Error = 212
	ErrUnknownTotalsMode                            Error = 213
	ErrCannotStatvfs                                Error = 214
	ErrNotAnAggregate                               Error = 215
	ErrQueryWithSameIDIsAlreadyRunning              Error = 216
	ErrClientHasConnectedToWrongPort                Error = 217
	ErrTableIsDropped                               Error = 218
	ErrDatabaseNotEmpty                             Error = 219
	ErrDuplicateInterserverIoEndpoint               Error = 220
	ErrNoSuchInterserverIoEndpoint                  Error = 221
	ErrAddingReplicaToNonEmptyTable                 Error = 222
	ErrUnexpectedAstStructure                       Error = 223
	ErrReplicaIsAlreadyActive                       Error = 224
	ErrNoZookeeper                                  Error = 225
	ErrNoFileInDataPart                             Error = 226
	ErrUnexpectedFileInDataPart                     Error = 227
	ErrBadSizeOfFileInDataPart                      Error = 228
	ErrQueryIsTooLarge                              Error = 229
	ErrNotFoundExpectedDataPart                     Error = 230
	ErrTooManyUnexpectedDataParts                   Error = 231
	ErrNoSuchDataPart                               Error = 232
	ErrBadDataPartName                              Error = 233
	ErrNoReplicaHasPart                             Error = 234
	ErrDuplicateDataPart                            Error = 235
	ErrAborted                                      Error = 236
	ErrNoReplicaNameGiven                           Error = 237
	ErrFormatVersionTooOld                          Error = 238
	ErrCannotMunmap                                 Error = 239
	ErrCannotMremap                                 Error = 240
	ErrMemoryLimitExceeded                          Error = 241
	ErrTableIsReadOnly                              Error = 242
	ErrNotEnoughSpace                               Error = 243
	ErrUnexpectedZookeeperError                     Error = 244
	ErrCorruptedData                                Error = 246
	ErrIncorrectMark                                Error = 247
	ErrInvalidPartitionValue                        Error = 248
	ErrNotEnoughBlockNumbers                        Error = 250
	ErrNoSuchReplica                                Error = 251
	ErrTooManyParts                                 Error = 252
	ErrReplicaIsAlreadyExist                        Error = 253
	ErrNoActiveReplicas                             Error = 254
	ErrTooManyRetriesToFetchParts                   Error = 255
	ErrPartitionAlreadyExists                       Error = 256
	ErrPartitionDoesntExist                         Error = 257
	ErrUnionAllResultStructuresMismatch             Error = 258
	ErrClientOutputFormatSpecified                  Error = 260
	ErrUnknownBlockInfoField                        Error = 261
	ErrBadCollation                                 Error = 262
	ErrCannotCompileCode                            Error = 263
	ErrIncompatibleTypeOfJoin                       Error = 264
	ErrNoAvailableReplica                           Error = 265
	ErrMismatchReplicasDataSources                  Error = 266
	ErrStorageDoesntSupportParallelReplicas         Error = 267
	ErrCPUIDError                                   Error = 268
	ErrInfiniteLoop                                 Error = 269
	ErrCannotCompress                               Error = 270
	ErrCannotDecompress                             Error = 271
	ErrAioSubmitError                               Error = 272
	ErrAioCompletionError                           Error = 273
	ErrAioReadError                                 Error = 274
	ErrAioWriteError                                Error = 275
	ErrIndexNotUsed                                 Error = 277
	ErrLeadershipLost                               Error = 278
	ErrAllConnectionTriesFailed                     Error = 279
	ErrNoAvailableData                              Error = 280
	ErrDictionaryIsEmpty                            Error = 281
	ErrIncorrectIndex                               Error = 282
	ErrUnknownDistributedProductMode                Error = 283
	ErrUnknownGlobalSubqueriesMethod                Error = 284
	ErrTooLessLiveReplicas                          Error = 285
	ErrUnsatisfiedQuorumForPreviousWrite            Error = 286
	ErrUnknownFormatVersion                         Error = 287
	ErrDistributedInJoinSubqueryDenied              Error = 288
	ErrReplicaIsNotInQuorum                         Error = 289
	ErrLimitExceeded                                Error = 290
	ErrDatabaseAccessDenied                         Error = 291
	ErrLeadershipChanged                            Error = 292
	ErrMongodbCannotAuthenticate                    Error = 293
	ErrInvalidBlockExtraInfo                        Error = 294
	ErrReceivedEmptyData                            Error = 295
	ErrNoRemoteShardFound                           Error = 296
	ErrShardHasNoConnections                        Error = 297
	ErrCannotPipe                                   Error = 298
	ErrCannotFork                                   Error = 299
	ErrCannotDlsym                                  Error = 300
	ErrCannotCreateChildProcess                     Error = 301
	ErrChildWasNotExitedNormally                    Error = 302
	ErrCannotSelect                                 Error = 303
	ErrCannotWaitpid                                Error = 304
	ErrTableWasNotDropped                           Error = 305
	ErrTooDeepRecursion                             Error = 306
	ErrTooManyBytes                                 Error = 307
	ErrUnexpectedNodeInZookeeper                    Error = 308
	ErrFunctionCannotHaveParameters                 Error = 309
	ErrInvalidShardWeight                           Error = 317
	ErrInvalidConfigParameter                       Error = 318
	ErrUnknownStatusOfInsert                        Error = 319
	ErrValueIsOutOfRangeOfDataType                  Error = 321
	ErrBarrierTimeout                               Error = 335
	ErrUnknownDatabaseEngine                        Error = 336
	ErrDdlGuardIsActive                             Error = 337
	ErrUnfinished                                   Error = 341
	ErrMetadataMismatch                             Error = 342
	ErrSupportIsDisabled                            Error = 344
	ErrTableDiffersTooMuch                          Error = 345
	ErrCannotConvertCharset                         Error = 346
	ErrCannotLoadConfig                             Error = 347
	ErrCannotInsertNullInOrdinaryColumn             Error = 349
	ErrIncompatibleSourceTables                     Error = 350
	ErrAmbiguousTableName                           Error = 351
	ErrAmbiguousColumnName                          Error = 352
	ErrIndexOfPositionalArgumentIsOutOfRange        Error = 353
	ErrZlibInflateFailed                            Error = 354
	ErrZlibDeflateFailed                            Error = 355
	ErrBadLambda                                    Error = 356
	ErrReservedIdentifierName                       Error = 357
	ErrIntoOutfileNotAllowed                        Error = 358
	ErrTableSizeExceedsMaxDropSizeLimit             Error = 359
	ErrCannotCreateCharsetConverter                 Error = 360
	ErrSeekPositionOutOfBound                       Error = 361
	ErrCurrentWriteBufferIsExhausted                Error = 362
	ErrCannotCreateIoBuffer                         Error = 363
	ErrReceivedErrorTooManyRequests                 Error = 364
	ErrOutputIsNotSorted                            Error = 365
	ErrSizesOfNestedColumnsAreInconsistent          Error = 366
	ErrTooManyFetches                               Error = 367
	ErrBadCast                                      Error = 368
	ErrAllReplicasAreStale                          Error = 369
	ErrDataTypeCannotBeUsedInTables                 Error = 370
	ErrInconsistentClusterDefinition                Error = 371
	ErrSessionNotFound                              Error = 372
	ErrSessionIsLocked                              Error = 373
	ErrInvalidSessionTimeout                        Error = 374
	ErrCannotDlopen                                 Error = 375
	ErrCannotParseUUID                              Error = 376
	ErrIllegalSyntaxForDataType                     Error = 377
	ErrDataTypeCannotHaveArguments                  Error = 378
	ErrUnknownStatusOfDistributedDdlTask            Error = 379
	ErrCannotKill                                   Error = 380
	ErrHTTPLengthRequired                           Error = 381
	ErrCannotLoadCatboostModel                      Error = 382
	ErrCannotApplyCatboostModel                     Error = 383
	ErrPartIsTemporarilyLocked                      Error = 384
	ErrMultipleStreamsRequired                      Error = 385
	ErrNoCommonType                                 Error = 386
	ErrExternalLoadableAlreadyExists                Error = 387
	ErrCannotAssignOptimize                         Error = 388
	ErrInsertWasDeduplicated                        Error = 389
	ErrCannotGetCreateTableQuery                    Error = 390
	ErrExternalLibraryError                         Error = 391
	ErrQueryIsProhibited                            Error = 392
	ErrThereIsNoQuery                               Error = 393
	ErrQueryWasCancelled                            Error = 394
	ErrFunctionThrowIfValueIsNonZero                Error = 395
	ErrTooManyRowsOrBytes                           Error = 396
	ErrQueryIsNotSupportedInMaterializedView        Error = 397
	ErrCannotParseDomainValueFromString             Error = 441
	ErrKeeperException                              Error = 999
	ErrPocoException                                Error = 1000
	ErrStdException                                 Error = 1001
	ErrUnknownException                             Error = 1002
	ErrConditionalTreeParentNotFound                Error = 2001
	ErrIllegalProjectionManipulator                 Error = 2002

	ErrAuthenticationFailed Error = 516
)
