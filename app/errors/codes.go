package errors

type (
	systemErrorCode                 int
	databaseErrorCode               int
	certificateErrorCode            int
	certificateTransactionErrorCode int
	transactionErrorCode            int
	notificationErrorCode           int
	soapClientErrorCode             int
	bobClientErrorCode              int
	axaptaErrorCode                 int
	ordersManagementErrorCode       int
	paymentProcessingErrorCode      int
	schedulerErrorCode              int
	eventsErrorCode                 int
	paidScheduleErrorCode           int
	expiredScheduleErrorCode        int
	stateMachineErrorCode           int
)

const errorCodesRangeLength = 1000 // Length of range of error codes for error type

const undefinedError = -1

// New error codes MUST BE appended to end of const block!!!
const (
	SystemBasicError systemErrorCode = iota + 0*errorCodesRangeLength
	SystemConfigurationError
	SystemMoneyConversionError
	SystemTimeError
	SystemDBQueryTypeError
)

// New error codes MUST BE appended to end of const block!!!
const (
	DatabaseBasicError databaseErrorCode = iota + 1*errorCodesRangeLength
	DatabaseConnectionError
	DatabaseCounterError
)

// New error codes MUST BE appended to end of const block!!!
const (
	CertificateBasicError certificateErrorCode = iota + 2*errorCodesRangeLength
	CertificateValidationError
	CertificateNotFoundError
	CertificateActivationCodeNotValidError
	CertificateAlreadyActivatedError
	CertificateStatusNotNewError
	CertificateNotActiveError
	CertificatePaymentAlreadyFinishedError
	CertificateCustomerIDNotValidError
	CertificateAmountTooLargeError
	CertificateNotPaidError
	CertificateAlreadySpentError
	CertificateAlreadyExpiredError
	CertificateAmountNotInRangeError
	CertificatePartnerIdNotValidError
	CertificatePurchaseDateAlreadyExpiredError
	CertificatePurchaseDateEmptyError
	CertificateAlreadyBlockedError
)

// New error codes MUST BE appended to end of const block!!!
const (
	CertificateTransactionBasicError certificateTransactionErrorCode = iota + 3*errorCodesRangeLength
)

// New error codes MUST BE appended to end of const block!!!
const (
	TransactionBasicError transactionErrorCode = iota + 4*errorCodesRangeLength
	TransactionNotHoldError
	TransactionAlreadyHoldError
	TransactionAmountTooLargeError
	TransactionAlreadyClosedError
)

// New error codes MUST BE appended to end of const block!!!
const (
	NotificationBasicError notificationErrorCode = iota + 5*errorCodesRangeLength
	NotificationNotFoundError
)

// New error codes MUST BE appended to end of const block!!!
const (
	SoapClientBasicError soapClientErrorCode = iota + 6*errorCodesRangeLength
	SoapClientRequestXmlParsingError
	SoapClientResponseXmlParsingError
	SoapClientConnectionError
	SoapClientServiceSideError
)

// New error codes MUST BE appended to end of const block!!!
const (
	BobClientBasicError bobClientErrorCode = iota + 7*errorCodesRangeLength
	BobClientConnectionError
	BobClientServiceSideError
	BobClientHTTPError
)

// New error codes MUST BE appended to end of const block!!!
const (
	AxaptaBasicError axaptaErrorCode = iota + 8*errorCodesRangeLength
	AxaptaResponseParsingError
	AxaptaServiceSideError
)

// New error codes MUST BE appended to end of const block!!!
const (
	OrdersManagementBasicError ordersManagementErrorCode = iota + 9*errorCodesRangeLength
	OrdersManagementCommunicationError
	OrdersManagementNumberNotGeneratedError
	OrdersManagementResponseValidationError
)

// New error codes MUST BE appended to end of const block!!!
const (
	PaymentProcessingBasicError paymentProcessingErrorCode = iota + 10*errorCodesRangeLength
	PaymentProcessingCommunicationError
	PaymentProcessingServiceSideError
)

// New error codes MUST BE appended to end of const block!!!
const (
	SchedulerBasicError schedulerErrorCode = iota + 11*errorCodesRangeLength
	SchedulerCronSpecParseError
	SchedulerAlreadyInWorkError
)

// New error codes MUST BE appended to end of const block!!!
const (
	EventsBasicError eventsErrorCode = iota + 12*errorCodesRangeLength
	EventsUnmarshalError
	EventsMarshalError
	EventsWasNotProducedError
	EventsRPCClientError
	EventsValidationError
)

// New error codes MUST BE appended to end of const block!!!
const (
	PaidScheduleBasicError paidScheduleErrorCode = iota + 13*errorCodesRangeLength
	PaidScheduleNotFoundError
)

const (
	ExpiredScheduleBasicError expiredScheduleErrorCode = iota + 14*errorCodesRangeLength
	ExpiredScheduleNotFoundError
)

const (
	StateMachineBasicError stateMachineErrorCode = iota + 15*errorCodesRangeLength
	StateMachineTransitionNotAvailableError
)

var errorMessagePrefixes = map[int]string{
	int(undefinedError): "undefined error",

	int(SystemBasicError):           "system error",
	int(SystemConfigurationError):   "system configuration error",
	int(SystemMoneyConversionError): "system money conversion error",
	int(SystemTimeError):            "system time error",
	int(SystemDBQueryTypeError):     "db query is not string",

	int(DatabaseBasicError):      "database error",
	int(DatabaseConnectionError): "database connection error",
	int(DatabaseCounterError):    "database 'counter' function return 0 rows",

	int(CertificateBasicError):                      "certificate error",
	int(CertificateValidationError):                 "certificate validation error",
	int(CertificateNotFoundError):                   "certificate is not found",                   // to var
	int(CertificateActivationCodeNotValidError):     "certificate's activation code is not found", // to var
	int(CertificateAlreadyActivatedError):           "certificate already activated",              // to var
	int(CertificateStatusNotNewError):               "certificate status is not 'new'",            // to var
	int(CertificateNotActiveError):                  "certificate not active",                     // to var
	int(CertificatePaymentAlreadyFinishedError):     "certificate's payment status is not 'new'",
	int(CertificateCustomerIDNotValidError):         "certificate' customer is not valid",
	int(CertificateAmountTooLargeError):             "certificate amount too large",
	int(CertificateNotPaidError):                    "certificate is not paid",
	int(CertificateAlreadySpentError):               "certificate already spent",
	int(CertificateAlreadyExpiredError):             "certificate already expired",
	int(CertificateAmountNotInRangeError):           "certificate amount not in limits range",
	int(CertificatePartnerIdNotValidError):          "certificate partner id is not valid",
	int(CertificatePurchaseDateAlreadyExpiredError): "certificate purchase date was already expired",
	int(CertificatePurchaseDateEmptyError):          "certificate purchase date is empty",
	int(CertificateAlreadyBlockedError):             "certificate already blocked",

	int(CertificateTransactionBasicError): "certificate transaction error",

	int(TransactionBasicError):          "transaction error",
	int(TransactionNotHoldError):        "transaction was not held",
	int(TransactionAlreadyHoldError):    "transaction already was held",
	int(TransactionAmountTooLargeError): "transaction amount is too large",
	int(TransactionAlreadyClosedError):  "transaction is already closed",

	int(NotificationBasicError):    "notification error",
	int(NotificationNotFoundError): "notification not found",

	int(SoapClientBasicError):              "soap client error",
	int(SoapClientRequestXmlParsingError):  "soap client request xml parsing error",
	int(SoapClientResponseXmlParsingError): "soap client response xml parsing error",
	int(SoapClientConnectionError):         "soap client connection error",
	int(SoapClientServiceSideError):        "soap client service error",

	int(BobClientBasicError):       "bob error",
	int(BobClientConnectionError):  "bob connection error",
	int(BobClientServiceSideError): "bob service error",
	int(BobClientHTTPError):        "bob service HTTP error",

	int(AxaptaBasicError):           "axapta error",
	int(AxaptaResponseParsingError): "axapta response parsing error",
	int(AxaptaServiceSideError):     "axapta service error",

	int(OrdersManagementBasicError):              "orders-management error",
	int(OrdersManagementCommunicationError):      "orders-management communication error",
	int(OrdersManagementNumberNotGeneratedError): "orders-management orderNr not generated",
	int(OrdersManagementResponseValidationError): "orders-management error response validation error",

	int(PaymentProcessingBasicError):         "payment-processing error",
	int(PaymentProcessingCommunicationError): "payment-processing communication error",
	int(PaymentProcessingServiceSideError):   "payment-processing service error",

	int(SchedulerBasicError):         "scheduler error",
	int(SchedulerCronSpecParseError): "scheduler cron spec parse error",
	int(SchedulerAlreadyInWorkError): "scheduler already is in work",

	int(EventsBasicError):          "events error",
	int(EventsUnmarshalError):      "events unmarshall error",
	int(EventsMarshalError):        "events marshall error",
	int(EventsWasNotProducedError): "events message was not produced",
	int(EventsRPCClientError):      "events RPC client error",
	int(EventsValidationError):     "events validation error",

	int(PaidScheduleBasicError):    "paid schedule error",
	int(PaidScheduleNotFoundError): "paid schedule not found",

	int(ExpiredScheduleBasicError):    "expired schedule error",
	int(ExpiredScheduleNotFoundError): "expired schedule not found",

	int(StateMachineBasicError):                  "state machine error",
	int(StateMachineTransitionNotAvailableError): "state machine transition is not available error (from %s to %s)",
}

func getErrorMessagePrefix(code int) string {
	prefix, ok := errorMessagePrefixes[code]
	if ok {
		return prefix
	}

	prefix, ok = errorMessagePrefixes[code-(code%errorCodesRangeLength)] // Try to get basic error
	if ok {
		return prefix
	}

	return errorMessagePrefixes[undefinedError]
}
