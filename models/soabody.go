package models

import "time"

type Soa struct {
	BalanceInquiryCASAResponse BalanceInquiryCASAResponse `json:"balanceInquiryCASAResponse"`
}
type SoaHeader struct {
	MessageVersion     string `json:"messageVersion"`
	MessageType        string `json:"messageType"`
	MessageSubType     string `json:"messageSubType"`
	MessageSender      string `json:"messageSender"`
	SenderDomain       string `json:"senderDomain"`
	MessageTimeStamp   string `json:"messageTimeStamp"`
	InitiatedTimeStamp string `json:"initiatedTimeStamp"`
	TrackingID         string `json:"trackingID"`
	ExceptionCode      string `json:"exceptionCode"`
}
type Property struct {
	PropertyKey   string `json:"propertyKey"`
	PropertyValue string `json:"propertyValue"`
}
type MessageHeader struct {
	Property []Property `json:"property"`
}
type AccountInfo struct {
	BranchNumber       string `json:"branchNumber"`
	ShortName          string `json:"shortName"`
	CifNumber          string `json:"cifNumber"`
	AlternateAddress   string `json:"alternateAddress"`
	AdditionalName     string `json:"additionalName"`
	AliasName          string `json:"aliasName"`
	Officer            string `json:"officer"`
	Status             string `json:"status"`
	AccountGroupNumber string `json:"accountGroupNumber"`
	CloseStatus        string `json:"closeStatus"`
	Class              string `json:"class"`
	CurrencyType       string `json:"currencyType"`
	CurrencyDecimals   string `json:"currencyDecimals"`
	GlGroupCode        string `json:"glGroupCode"`
}
type AccountAmountDetails struct {
	HoldAmount          string `json:"holdAmount"`
	TotalOCAmount       string `json:"totalOCAmount"`
	TotalOCPAmount      string `json:"totalOCPAmount"`
	TotalODOCAmount     string `json:"totalODOCAmount"`
	TotalMCAmount       string `json:"totalMCAmount"`
	TotalIBTAmount      string `json:"totalIBTAmount"`
	TotalLLCAmount      string `json:"totalLLCAmount"`
	TotalEXTRAAmount    string `json:"totalEXTRAAmount"`
	TotalSuspenseAmount string `json:"totalSuspenseAmount"`
	FloatDrawingAmount  string `json:"floatDrawingAmount"`
}
type BalanceInfo struct {
	LedgerBalance            string `json:"ledgerBalance"`
	YesterdayLedgerBalance   string `json:"yesterdayLedgerBalance"`
	PreviousStatementBalance string `json:"previousStatementBalance"`
	PrevYearEndLedgerBalance string `json:"prevYearEndLedgerBalance"`
	MinimumBalance           string `json:"minimumBalance"`
	MaximumBalance           string `json:"maximumBalance"`
	CollectedBalance         string `json:"collectedBalance"`
}
type InterestInfo struct {
	AccruedODInterest        string `json:"accruedODInterest"`
	AccruedInterest          string `json:"accruedInterest"`
	AlternateAccruedInterest string `json:"alternateAccruedInterest"`
	YtdInterestPaid          string `json:"ytdInterestPaid"`
	PreviousYTDInterestPaid  string `json:"previousYTDInterestPaid"`
	AmtOfLastIntPd           string `json:"amtOfLastIntPd"`
}
type UserCode struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}
type UserCodeInfo struct {
	UserCode []UserCode `json:"userCode"`
}
type DebitInfos struct {
	DebitName  string `json:"debitName"`
	DebitValue string `json:"debitValue"`
}
type DebitInfo struct {
	DebitInfo []DebitInfos `json:"debitInfo"`
}
type CreditInfos struct {
	CreditName  string `json:"creditName"`
	CreditValue string `json:"creditValue"`
}
type CreditInfo struct {
	CreditInfo []CreditInfos `json:"creditInfo"`
}
type AtmDebitAndCreditInfo struct {
	NoOfOnUsATMCredits  string `json:"noOfOnUsATMCredits"`
	AmtOfONusATMCredits string `json:"amtOfONusATMCredits"`
	NoOfOnUsATMDebits   string `json:"noOfOnUsATMDebits"`
	AmtOfOnUsATMDebits  string `json:"amtOfOnUsATMDebits"`
	NoOfFrgnATMCredits  string `json:"noOfFrgnATMCredits"`
	AmtOfFrgnATMCredits string `json:"amtOfFrgnATMCredits"`
	NoOfFrgnATMDebits   string `json:"noOfFrgnATMDebits"`
	AmtOfFrgnATMDebits  string `json:"amtOfFrgnATMDebits"`
}
type OverDraftCycleInfo struct {
	TimesODThisCycle  string `json:"timesODThisCycle"`
	TimesODThisQtr    string `json:"timesODThisQtr"`
	TimesOD2NdQtr     string `json:"timesOD2ndQtr"`
	TimesOD3RdQtr     string `json:"timesOD3rdQtr"`
	TimesOD4ThQtr     string `json:"timesOD4thQtr"`
	TimesODLastYear   string `json:"timesODLastYear"`
	TimesODLifeToDate string `json:"timesODLifeToDate"`
}
type AccountRateInfo struct {
	AcctTierRateNumber   string `json:"acctTierRateNumber"`
	InterestRate         string `json:"interestRate"`
	RateVariance         string `json:"rateVariance"`
	RateVarianceCode     string `json:"rateVarianceCode"`
	RateFloor            string `json:"rateFloor"`
	RateCeiling          string `json:"rateCeiling"`
	RateReviewDate       string `json:"rateReviewDate"`
	RateReviewTerm       string `json:"rateReviewTerm"`
	RateReviewTermCode   string `json:"rateReviewTermCode"`
	OdInterestRate       string `json:"odInterestRate"`
	OdRateVariance       string `json:"odRateVariance"`
	OdRateVarianceCode   string `json:"odRateVarianceCode"`
	OdRateFloor          string `json:"odRateFloor"`
	OdRateCeiling        string `json:"odRateCeiling"`
	OdRateReviewDate     string `json:"odRateReviewDate"`
	OdRateReviewTerm     string `json:"odRateReviewTerm"`
	OdRateReviewTermCode string `json:"odRateReviewTermCode"`
}
type AggregateBalanceInfo struct {
	AggregateDaysThisQtr                  string `json:"aggregateDaysThisQtr"`
	AggregateLedgerBalanceThisQtr         string `json:"aggregateLedgerBalanceThisQtr"`
	AggregateCollectedBalanceThisQtr      string `json:"aggregateCollectedBalanceThisQtr"`
	AverageLedgerBalancePriorFirstQtr     string `json:"averageLedgerBalancePriorFirstQtr"`
	AverageLedgerBalancePriorSecondQtr    string `json:"averageLedgerBalancePriorSecondQtr"`
	AverageLedgerBalancePriorThirdQtr     string `json:"averageLedgerBalancePriorThirdQtr"`
	AverageLedgerBalancePriorFourthQtr    string `json:"averageLedgerBalancePriorFourthQtr"`
	AverageCollectedBalancePriorFirstQtr  string `json:"averageCollectedBalancePriorFirstQtr"`
	AverageCollectedBalancePriorSecondQtr string `json:"averageCollectedBalancePriorSecondQtr"`
	AverageCollectedBalancePriorThirdQtr  string `json:"averageCollectedBalancePriorThirdQtr"`
	AverageCollectedBalancePriorFourthQtr string `json:"averageCollectedBalancePriorFourthQtr"`
}
type SpecialInstruction struct {
	SiCode  string `json:"siCode"`
	SiValue string `json:"siValue"`
}
type SpecialInstructionInfo struct {
	SpecialInstruction []SpecialInstruction `json:"specialInstruction"`
}
type NsfCycleInfo struct {
	NoNSFThisCycle  string `json:"noNSFThisCycle"`
	NoNSFThisQtr    string `json:"noNSFThisQtr"`
	NoNSF2NdQtr     string `json:"noNSF2ndQtr"`
	NoNSF3RdQtr     string `json:"noNSF3rdQtr"`
	NoNSthQtr       string `json:"noNSthQtr"`
	NoNSFLastYear   string `json:"noNSFLastYear"`
	NoNSFLifeToDate string `json:"noNSFLifeToDate"`
}
type AccountWHInfo struct {
	FederalWHCode        string `json:"federalWHCode"`
	StateWHCode          string `json:"stateWHCode"`
	State                string `json:"state"`
	DateWHCodeChanged    string `json:"dateWHCodeChanged"`
	AlternateAccountNo   string `json:"alternateAccountNo"`
	AlternateAccountType string `json:"alternateAccountType"`
	FedWHAlternateRate   string `json:"fedWHAlternateRate"`
	StateWHAlternateRate string `json:"stateWHAlternateRate"`
	FedWHToday           string `json:"fedWHToday"`
	FederalWHMTD         string `json:"federalWHMTD"`
	FederalWHQTD         string `json:"federalWHQTD"`
	FederalWHYTD         string `json:"federalWHYTD"`
	StateWHToday         string `json:"stateWHToday"`
	StateWHMTD           string `json:"stateWHMTD"`
	StateWHQTD           string `json:"stateWHQTD"`
	StateWHYTD           string `json:"stateWHYTD"`
	FederalWHLastYear    string `json:"federalWHLastYear"`
	StateWHLastYear      string `json:"stateWHLastYear"`
}
type AccountMTDInfo struct {
	MtdServiceCharge        string `json:"mtdServiceCharge"`
	MtdInterestPaid         string `json:"mtdInterestPaid"`
	MtdODInterestCharged    string `json:"mtdODInterestCharged"`
	MtdFeesPaid             string `json:"mtdFeesPaid"`
	MtdAggregateBalance     string `json:"mtdAggregateBalance"`
	MtdAggregateDays        string `json:"mtdAggregateDays"`
	MtdBeginningAccrual     string `json:"mtdBeginningAccrual"`
	MtdReturnedDebitCount   string `json:"mtdReturnedDebitCount"`
	MtdReturnedDebitAmount  string `json:"mtdReturnedDebitAmount"`
	MtdReturnedCreditCount  string `json:"mtdReturnedCreditCount"`
	MtdReturnedCreditAmount string `json:"mtdReturnedCreditAmount"`
}
type Payload struct {
	ResponseCode                  string                 `json:"responseCode"`
	ResponseMessage               string                 `json:"responseMessage"`
	ResponseTimestamp             time.Time              `json:"responseTimestamp"`
	AccountInfo                   AccountInfo            `json:"accountInfo"`
	StopCode                      string                 `json:"stopCode"`
	HoldCode                      string                 `json:"holdCode"`
	AlertCode                     string                 `json:"alertCode"`
	SpecialInstructionCode        string                 `json:"specialInstructionCode"`
	PrintInterestCheck            string                 `json:"printInterestCheck"`
	DateOpened                    string                 `json:"dateOpened"`
	DateOfStatus                  string                 `json:"dateOfStatus"`
	DateEntered                   string                 `json:"dateEntered"`
	DateLastOverdrawn             string                 `json:"dateLastOverdrawn"`
	PriorLastActiveDate           string                 `json:"priorLastActiveDate"`
	DateLastActive                string                 `json:"dateLastActive"`
	DateLastContact               string                 `json:"dateLastContact"`
	DateLastInterestPaid          string                 `json:"dateLastInterestPaid"`
	PreviousIntPayDate            string                 `json:"previousIntPayDate"`
	DateLastStatement             string                 `json:"dateLastStatement"`
	DateLastMaintenance           string                 `json:"dateLastMaintenance"`
	DateLastDeposit               string                 `json:"dateLastDeposit"`
	DateSCWaiveExpires            string                 `json:"dateSCWaiveExpires"`
	DateLastCARCodeMaint          string                 `json:"dateLastCARCodeMaint"`
	AmtOfLastDeposit              string                 `json:"amtOfLastDeposit"`
	AccountAmountDetails          AccountAmountDetails   `json:"accountAmountDetails"`
	BalanceInfo                   BalanceInfo            `json:"balanceInfo"`
	InterestInfo                  InterestInfo           `json:"interestInfo"`
	StatementCycle                string                 `json:"statementCycle"`
	StatementCode                 string                 `json:"statementCode"`
	InterestCycle                 string                 `json:"interestCycle"`
	ServiceChargeCycle            string                 `json:"serviceChargeCycle"`
	DepositTypeCode               string                 `json:"depositTypeCode"`
	ServiceChargeMode             string                 `json:"serviceChargeMode"`
	UserCodeInfo                  UserCodeInfo           `json:"userCodeInfo"`
	DebitInfo                     DebitInfo              `json:"debitInfo"`
	CreditInfo                    CreditInfo             `json:"creditInfo"`
	AtmDebitAndCreditInfo         AtmDebitAndCreditInfo  `json:"atmDebitAndCreditInfo"`
	AmountOfNSFItems              string                 `json:"amountOfNSFItems"`
	OverdraftLimitCode            string                 `json:"overdraftLimitCode"`
	SignatureVerification         string                 `json:"signatureVerification"`
	OverdraftProtection           string                 `json:"overdraftProtection"`
	ValidForCheckBook             string                 `json:"validForCheckBook"`
	AtmCard                       string                 `json:"atmCard"`
	IncludeOnCombinedStatement    string                 `json:"includeOnCombinedStatement"`
	HoldMailCode                  string                 `json:"holdMailCode"`
	AutomaticNSFCharge            string                 `json:"automaticNSFCharge"`
	HighVolumeAccount             string                 `json:"highVolumeAccount"`
	BadCheckIncident              string                 `json:"badCheckIncident"`
	OverDraftCycleInfo            OverDraftCycleInfo     `json:"overDraftCycleInfo"`
	RelatedAccountNo              string                 `json:"relatedAccountNo"`
	RelatedAccountType            string                 `json:"relatedAccountType"`
	NoCksDepositOnUs              string                 `json:"noCksDepositOnUs"`
	NoCksDepositForeign           string                 `json:"noCksDepositForeign"`
	ListPostCode                  string                 `json:"listPostCode"`
	AccountRateInfo               AccountRateInfo        `json:"accountRateInfo"`
	AutomaticalyMaintainNPLStatus string                 `json:"automaticalyMaintainNPLStatus"`
	NonperformingStatus           string                 `json:"nonperformingStatus"`
	FirstReachedExcessDate        string                 `json:"firstReachedExcessDate"`
	NoDebitsUpTo3                 string                 `json:"noDebitsUpTo3"`
	NoDebitsUpTo6                 string                 `json:"noDebitsUpTo6"`
	NoDebitsUpTo3Exceed           string                 `json:"noDebitsUpTo3Exceed"`
	NoDebitsUpTo6Exceed           string                 `json:"noDebitsUpTo6Exceed"`
	CashDepositThisCycle          string                 `json:"cashDepositThisCycle"`
	CashPaidThisCycle             string                 `json:"cashPaidThisCycle"`
	AggregateBalanceInfo          AggregateBalanceInfo   `json:"aggregateBalanceInfo"`
	SpecialInstructionInfo        SpecialInstructionInfo `json:"specialInstructionInfo"`
	SpecialMessage                string                 `json:"specialMessage"`
	ClubPlan                      string                 `json:"clubPlan"`
	NsfCycleInfo                  NsfCycleInfo           `json:"nsfCycleInfo"`
	TotalAccrualFloat             string                 `json:"totalAccrualFloat"`
	TotalAvailibilityFloat        string                 `json:"totalAvailibilityFloat"`
	TotalCashFloat                string                 `json:"totalCashFloat"`
	NoOfEnclosures                string                 `json:"noOfEnclosures"`
	EftFlag                       string                 `json:"eftFlag"`
	MinimumCharge                 string                 `json:"minimumCharge"`
	NoLargeBalanceFluctuationMTD  string                 `json:"noLargeBalanceFluctuationMTD"`
	NoLargeBalanceFluctuationYTD  string                 `json:"noLargeBalanceFluctuationYTD"`
	NoTimesOnKiteSuspectReport    string                 `json:"noTimesOnKiteSuspectReport"`
	AccountWHInfo                 AccountWHInfo          `json:"accountWHInfo"`
	HighestInterestRateEarned     string                 `json:"highestInterestRateEarned"`
	NsfItemsExist                 string                 `json:"nsfItemsExist"`
	AftCode                       string                 `json:"aftCode"`
	StatementPassbookCode         string                 `json:"statementPassbookCode"`
	TellerStatus                  string                 `json:"tellerStatus"`
	PrintChecksInOrderOnStatement string                 `json:"printChecksInOrderOnStatement"`
	AccountMTDInfo                AccountMTDInfo         `json:"accountMTDInfo"`
	ClassCARCode                  string                 `json:"classCARCode"`
	SelectionNumber               string                 `json:"selectionNumber"`
	OdInterestRateNo              string                 `json:"odInterestRateNo"`
	AccrualMethod                 string                 `json:"accrualMethod"`
	AccrueOnLedColMinBal          string                 `json:"accrueOnLedColMinBal"`
	YearBaseCode                  string                 `json:"yearBaseCode"`
	MinimumBalToMaintain          string                 `json:"minimumBalToMaintain"`
	MinimumAccrualBalance         string                 `json:"minimumAccrualBalance"`
	AverageLedgerBalance          string                 `json:"averageLedgerBalance"`
	AvailableBalance              string                 `json:"availableBalance"`
	OneDayFloatLC                 string                 `json:"oneDayFloatLC"`
	TwoDayFloatLC                 string                 `json:"twoDayFloatLC"`
	ThreeDayFloatLC               string                 `json:"threeDayFloatLC"`
	TotalDepositToday             string                 `json:"totalDepositToday"`
	TotalWithdrawalToday          string                 `json:"totalWithdrawalToday"`
	TotalOCDepositedToday         string                 `json:"totalOCDepositedToday"`
	TotalLCDepositedToday         string                 `json:"totalLCDepositedToday"`
	TotalAuthLimit                string                 `json:"totalAuthLimit"`
	TotalDrawLimit                string                 `json:"totalDrawLimit"`
	OthFloatOSOC                  string                 `json:"othFloatOSOC"`
	OthFloatOSFS                  string                 `json:"othFloatOSFS"`
	CreditAccrInt                 string                 `json:"creditAccrInt"`
	DebitAccrInt                  string                 `json:"debitAccrInt"`
	CommitmentFee                 string                 `json:"commitmentFee"`
	Total1DayFloat                string                 `json:"total1DayFloat"`
	OdLimit                       string                 `json:"odLimit"`
	WhRate                        string                 `json:"whRate"`
	IDNumber                      string                 `json:"idNumber"`
	CustomerTypeCode              string                 `json:"customerTypeCode"`
	SexCode                       string                 `json:"sexCode"`
	VipCustomerCode               string                 `json:"vipCustomerCode"`
	CountryOfHeritage             string                 `json:"countryOfHeritage"`
	InsiderCode                   string                 `json:"insiderCode"`
	ResidentCode                  string                 `json:"residentCode"`
	Country                       string                 `json:"country"`
	BusinessType                  string                 `json:"businessType"`
	BnmbumiNRCCCode               string                 `json:"bnmbumiNRCCCode"`
	TimesBadCheques               string                 `json:"timesBadCheques"`
	InterestRate0                 string                 `json:"interestRate0"`
	GroupedFields                 string                 `json:"groupedFields"`
}
type BalanceInquiryCASAResponse struct {
	SoaHeader     SoaHeader     `json:"soaHeader"`
	MessageHeader MessageHeader `json:"messageHeader"`
	Payload       Payload       `json:"payload"`
}
