package lockdownd

import (
	"fmt"
)

// Error list
var (
	ErrInvalidArg                          = fmt.Errorf("invalid argument")
	ErrInvalidConf                         = fmt.Errorf("invalid configuration")
	ErrPlist                               = fmt.Errorf("plist error")
	ErrPairingFailed                       = fmt.Errorf("pairing failed")
	ErrSSL                                 = fmt.Errorf("ssl error")
	ErrDict                                = fmt.Errorf("dict error")
	ErrReceiveTimeout                      = fmt.Errorf("receive timeout")
	ErrMux                                 = fmt.Errorf("mux error")
	ErrNoRunningSession                    = fmt.Errorf("no running session")
	ErrInvalidReponse                      = fmt.Errorf("invalid reponse")
	ErrMissingKey                          = fmt.Errorf("missing key")
	ErrMissingValue                        = fmt.Errorf("missing value")
	ErrGetProhibited                       = fmt.Errorf("get prohibited")
	ErrSetProhibited                       = fmt.Errorf("set prohibited")
	ErrRemoveProhibited                    = fmt.Errorf("remove prohibited")
	ErrImmutableValue                      = fmt.Errorf("immutable value")
	ErrPasswordProtected                   = fmt.Errorf("password protected")
	ErrUserDeniedPairing                   = fmt.Errorf("user denied pairing")
	ErrPairingDialogResponsePending        = fmt.Errorf("pairing dialog response pending")
	ErrMissingHostID                       = fmt.Errorf("missing host id")
	ErrInvalidHostID                       = fmt.Errorf("invalid host id")
	ErrSessionActive                       = fmt.Errorf("session active")
	ErrSessionInactive                     = fmt.Errorf("session inactive")
	ErrMissingSessionID                    = fmt.Errorf("missing session id")
	ErrInvalidSessionID                    = fmt.Errorf("invalid session id")
	ErrMissingService                      = fmt.Errorf("Missing service")
	ErrInvalidService                      = fmt.Errorf("invalid service")
	ErrServiceLimit                        = fmt.Errorf("service limit")
	ErrMissingPairRecord                   = fmt.Errorf("missing pair record")
	ErrSavePairRecordFailed                = fmt.Errorf("save pair record failed")
	ErrInvalidPairRecord                   = fmt.Errorf("invalid pair record")
	ErrInvalidActivationRecord             = fmt.Errorf("invalid activation record")
	ErrMissingActivationRecord             = fmt.Errorf("missing activation record")
	ErrServiceProhibited                   = fmt.Errorf("service prohibited")
	ErrEscrowLocked                        = fmt.Errorf("escrow locked")
	ErrPairingProhibitedOverThisConnection = fmt.Errorf("pairing prohibited over this connection")
	ErrFMIPProtected                       = fmt.Errorf("fmip protected")
	ErrMCProtected                         = fmt.Errorf("mc protected")
	ErrMCChallengeRequired                 = fmt.Errorf("mc challenge required")
	ErrUnknown                             = fmt.Errorf("unknown error")
)

var errorMap = map[int]error{
	-1:   ErrInvalidArg,
	-2:   ErrInvalidConf,
	-3:   ErrPlist,
	-4:   ErrPairingFailed,
	-5:   ErrSSL,
	-6:   ErrDict,
	-7:   ErrReceiveTimeout,
	-8:   ErrMux,
	-9:   ErrNoRunningSession,
	-10:  ErrInvalidReponse,
	-11:  ErrMissingKey,
	-12:  ErrMissingValue,
	-13:  ErrGetProhibited,
	-14:  ErrSetProhibited,
	-15:  ErrRemoveProhibited,
	-16:  ErrImmutableValue,
	-17:  ErrPasswordProtected,
	-18:  ErrUserDeniedPairing,
	-19:  ErrPairingDialogResponsePending,
	-20:  ErrMissingHostID,
	-21:  ErrInvalidHostID,
	-22:  ErrSessionActive,
	-23:  ErrSessionInactive,
	-24:  ErrMissingSessionID,
	-25:  ErrInvalidSessionID,
	-26:  ErrMissingService,
	-27:  ErrInvalidService,
	-28:  ErrServiceLimit,
	-29:  ErrMissingPairRecord,
	-30:  ErrSavePairRecordFailed,
	-31:  ErrInvalidPairRecord,
	-32:  ErrInvalidActivationRecord,
	-33:  ErrMissingActivationRecord,
	-34:  ErrServiceProhibited,
	-35:  ErrEscrowLocked,
	-36:  ErrPairingProhibitedOverThisConnection,
	-37:  ErrFMIPProtected,
	-38:  ErrMCProtected,
	-39:  ErrMCChallengeRequired,
	-256: ErrUnknown,
}

func handleError(code int) error {
	e, ok := errorMap[code]
	if !ok {
		e = ErrUnknown
	}

	return e
}
