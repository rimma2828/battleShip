package errors

type DatabaseError struct {
	baseError
}

func NewDatabaseError(code databaseErrorCode, err error) error {
	return CertificateError{newBaseError(int(code), err)}
}

type CertificateError struct {
	baseError
}

var (
	ErrDatabaseCounter = NewDatabaseError(DatabaseCounterError, nil)
)

func NewCertificateError(code certificateErrorCode, err error) error {
	return CertificateError{newBaseError(int(code), err)}
}

var (
	ErrActivationCodeNotValid                = NewCertificateError(CertificateActivationCodeNotValidError, nil)
	ErrCertAlreadyActivated                  = NewCertificateError(CertificateAlreadyActivatedError, nil)
	ErrCertStatusNotNew                      = NewCertificateError(CertificateStatusNotNewError, nil)
	ErrCertificateCustomerIDNotValid         = NewCertificateError(CertificateCustomerIDNotValidError, nil)
	ErrCertificateNotActive                  = NewCertificateError(CertificateNotActiveError, nil)
	ErrCertificateAmountTooLarge             = NewCertificateError(CertificateAmountTooLargeError, nil)
	ErrCertificateAmountNotInRange           = NewCertificateError(CertificateAmountNotInRangeError, nil)
	ErrCertificateNotPaid                    = NewCertificateError(CertificateNotPaidError, nil)
	ErrCertificateAlreadySpent               = NewCertificateError(CertificateAlreadySpentError, nil)
	ErrCertificateAlreadyExpired             = NewCertificateError(CertificateAlreadyExpiredError, nil)
	ErrCertificatePurchaseDateAlreadyExpired = NewCertificateError(CertificatePurchaseDateAlreadyExpiredError, nil)
	ErrCertificatePurchaseDateEmptyError     = NewCertificateError(CertificatePurchaseDateEmptyError, nil)
	ErrCertificateValidationError            = NewCertificateError(CertificateValidationError, nil)
	ErrCertificateAlreadyBlocked             = NewCertificateError(CertificateAlreadyBlockedError, nil)
)

type TransactionError struct {
	baseError
}

func NewTransactionError(code transactionErrorCode, err error) error {
	return TransactionError{newBaseError(int(code), err)}
}

var (
	ErrTransactionNotHold        = NewTransactionError(TransactionNotHoldError, nil)
	ErrTransactionAlreadyHold    = NewTransactionError(TransactionAlreadyHoldError, nil)
	ErrTransactionAmountTooLarge = NewTransactionError(TransactionAmountTooLargeError, nil)
	ErrTransactionAlreadyClosed  = NewTransactionError(TransactionAlreadyClosedError, nil)
)

type CertificateTransactionError struct {
	baseError
}

func NewCertificateTransactionError(code certificateTransactionErrorCode, err error) error {
	return CertificateTransactionError{newBaseError(int(code), err)}
}

type BobClientError struct {
	baseError
}

func NewBobClientError(code bobClientErrorCode, err error) error {
	return BobClientError{newBaseError(int(code), err)}
}

type SoapClientError struct {
	baseError
	ResponseBody []byte
}

func NewSoapClientError(code soapClientErrorCode, err error) error {
	return SoapClientError{newBaseError(int(code), err), nil}
}

func NewSoapClientErrorWithResponse(code soapClientErrorCode, err error, response []byte) error {
	return SoapClientError{newBaseError(int(code), err), response}
}

type SystemError struct {
	baseError
}

func NewSystemError(code systemErrorCode, err error) error {
	return SystemError{newBaseError(int(code), err)}
}

var (
	ErrSystemDBQueryType = NewSystemError(SystemDBQueryTypeError, nil)
)

type AxaptaError struct {
	baseError
}

func NewAxaptaError(code axaptaErrorCode, err error) error {
	return AxaptaError{newBaseError(int(code), err)}
}

type OrdersManagementError struct {
	baseError
}

func NewOrdersManagementError(code ordersManagementErrorCode, err error) error {
	return OrdersManagementError{newBaseError(int(code), err)}
}

var (
	ErrOrdersManagementNumberNotGenerated = NewOrdersManagementError(OrdersManagementNumberNotGeneratedError, nil)
)

type PaymentProcessingError struct {
	baseError
}

func NewPaymentProcessingError(code paymentProcessingErrorCode, err error) error {
	return PaymentProcessingError{newBaseError(int(code), err)}
}

type SchedulerError struct {
	baseError
}

func NewSchedulerError(code schedulerErrorCode, err error) error {
	return SchedulerError{newBaseError(int(code), err)}
}

var (
	ErrSchedulerAlreadyInWork = NewSchedulerError(SchedulerAlreadyInWorkError, nil)
)

type EventsError struct {
	baseError
}

func NewEventsError(code eventsErrorCode, err error) error {
	return EventsError{newBaseError(int(code), err)}
}

var (
	ErrEventsWasNotProduced = NewEventsError(EventsWasNotProducedError, nil)
)

type NotificationsError struct {
	baseError
}

func NewNotificationsError(code notificationErrorCode, err error) error {
	return NotificationsError{newBaseError(int(code), err)}
}

type MarkPaidScheduleError struct {
	baseError
}

func NewMarkPaidScheduleError(code paidScheduleErrorCode, err error) error {
	return MarkPaidScheduleError{newBaseError(int(code), err)}
}

type MarkExpiredScheduleError struct {
	baseError
}

func NewMarkExpiredScheduleError(code expiredScheduleErrorCode, err error) error {
	return MarkExpiredScheduleError{newBaseError(int(code), err)}
}

type StateMachineError struct {
	baseError
}

func NewStateMachineError(code stateMachineErrorCode, err error, currentState string, destinationState string) error {
	return StateMachineError{newStateMachineCustomError(int(code), err, currentState, destinationState)}
}
