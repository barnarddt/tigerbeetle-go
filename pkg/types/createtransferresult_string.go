// Code generated by "stringer -type=CreateTransferResult -trimprefix=Transfer"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TransferLinkedEventFailed-1]
	_ = x[TransferReservedFlag-2]
	_ = x[TransferReservedField-3]
	_ = x[TransferIdMustNotBeZero-4]
	_ = x[TransferIdMustNotBeMaxInt-5]
	_ = x[TransferDebitAccountIdMustNotBeZero-6]
	_ = x[TransferDebitAccountIdMustNotBeMaxInt-7]
	_ = x[TransferCreditAccountIdMustNotBeZero-8]
	_ = x[TransferCreditAccountIdMustNotBeMaxInt-9]
	_ = x[TransferAccountsMustBeDifferent-10]
	_ = x[TransferPendingIdMustBeZero-11]
	_ = x[TransferPendingTransferMustTimeout-12]
	_ = x[TransferLedgerMustNotBeZero-13]
	_ = x[TransferCodeMustNotBeZero-14]
	_ = x[TransferAmountMustNotBeZero-15]
	_ = x[TransferDebitAccountNotFound-16]
	_ = x[TransferCreditAccountNotFound-17]
	_ = x[TransferAccountsMustHaveTheSameLedger-18]
	_ = x[TransferTransferMustHaveTheSameLedgerAsAccounts-19]
	_ = x[TransferExistsWithDifferentFlags-20]
	_ = x[TransferExistsWithDifferentDebitAccountId-21]
	_ = x[TransferExistsWithDifferentCreditAccountId-22]
	_ = x[TransferExistsWithDifferentUserData-23]
	_ = x[TransferExistsWithDifferentPendingId-24]
	_ = x[TransferExistsWithDifferentTimeout-25]
	_ = x[TransferExistsWithDifferentCode-26]
	_ = x[TransferExistsWithDifferentAmount-27]
	_ = x[TransferExists-28]
	_ = x[TransferOverflowsDebitsPending-29]
	_ = x[TransferOverflowsCreditsPending-30]
	_ = x[TransferOverflowsDebitsPosted-31]
	_ = x[TransferOverflowsCreditsPosted-32]
	_ = x[TransferOverflowsDebits-33]
	_ = x[TransferOverflowsCredits-34]
	_ = x[TransferExceedsCredits-35]
	_ = x[TransferExceedsDebits-36]
	_ = x[TransferCannotPostAndVoidPendingTransfer-37]
	_ = x[TransferPendingTransferCannotPostOrVoidAnother-38]
	_ = x[TransferTimeoutReservedForPendingTransfer-39]
	_ = x[TransferPendingIdMustNotBeZero-40]
	_ = x[TransferPendingIdMustNotBeMaxInt-41]
	_ = x[TransferPendingIdMustBeDifferent-42]
	_ = x[TransferPendingTransferNotFound-43]
	_ = x[TransferPendingTransferNotPending-44]
	_ = x[TransferPendingTransferHasDifferentDebitAccountId-45]
	_ = x[TransferPendingTransferHasDifferentCreditAccountId-46]
	_ = x[TransferPendingTransferHasDifferentLedger-47]
	_ = x[TransferPendingTransferHasDifferentCode-48]
	_ = x[TransferExceedsPendingTransferAmount-49]
	_ = x[TransferPendingTransferHasDifferentAmount-50]
	_ = x[TransferPendingTransferAlreadyPosted-51]
	_ = x[TransferPendingTransferAlreadyVoided-52]
	_ = x[TransferPendingTransferExpired-53]
}

const _CreateTransferResult_name = "LinkedEventFailedReservedFlagReservedFieldIdMustNotBeZeroIdMustNotBeMaxIntDebitAccountIdMustNotBeZeroDebitAccountIdMustNotBeMaxIntCreditAccountIdMustNotBeZeroCreditAccountIdMustNotBeMaxIntAccountsMustBeDifferentPendingIdMustBeZeroPendingTransferMustTimeoutLedgerMustNotBeZeroCodeMustNotBeZeroAmountMustNotBeZeroDebitAccountNotFoundCreditAccountNotFoundAccountsMustHaveTheSameLedgerTransferMustHaveTheSameLedgerAsAccountsExistsWithDifferentFlagsExistsWithDifferentDebitAccountIdExistsWithDifferentCreditAccountIdExistsWithDifferentUserDataExistsWithDifferentPendingIdExistsWithDifferentTimeoutExistsWithDifferentCodeExistsWithDifferentAmountExistsOverflowsDebitsPendingOverflowsCreditsPendingOverflowsDebitsPostedOverflowsCreditsPostedOverflowsDebitsOverflowsCreditsExceedsCreditsExceedsDebitsCannotPostAndVoidPendingTransferPendingTransferCannotPostOrVoidAnotherTimeoutReservedForPendingTransferPendingIdMustNotBeZeroPendingIdMustNotBeMaxIntPendingIdMustBeDifferentPendingTransferNotFoundPendingTransferNotPendingPendingTransferHasDifferentDebitAccountIdPendingTransferHasDifferentCreditAccountIdPendingTransferHasDifferentLedgerPendingTransferHasDifferentCodeExceedsPendingTransferAmountPendingTransferHasDifferentAmountPendingTransferAlreadyPostedPendingTransferAlreadyVoidedPendingTransferExpired"

var _CreateTransferResult_index = [...]uint16{0, 17, 29, 42, 57, 74, 101, 130, 158, 188, 211, 230, 256, 275, 292, 311, 331, 352, 381, 420, 444, 477, 511, 538, 566, 592, 615, 640, 646, 668, 691, 712, 734, 749, 765, 779, 792, 824, 862, 895, 917, 941, 965, 988, 1013, 1054, 1096, 1129, 1160, 1188, 1221, 1249, 1277, 1299}

func (i CreateTransferResult) String() string {
	i -= 1
	if i >= CreateTransferResult(len(_CreateTransferResult_index)-1) {
		return "CreateTransferResult(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CreateTransferResult_name[_CreateTransferResult_index[i]:_CreateTransferResult_index[i+1]]
}
