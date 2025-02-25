package status

// FTP reply codes as specified in RFC 959
const (
	FileOK = "150 File status okay; about to open data connection."

	OK              = "200 Command okay."
	Created         = "201 Created."
	CommandNotImpl  = "202 Command not implemented."
	System          = "211 System status, or system help reply."
	Directory       = "212 Directory status."
	File            = "213 File status."
	Help            = "214 Help message."
	Name            = "215 NAME system type."
	Ready           = "220 Service ready for new user."
	Closing         = "221 Service closing control connection."
	DataSessionOpen = "225 Data connection open; no transfer in progress."
	ClosingData     = "226 Closing data connection. Requested file action successful."
	PassiveMode     = "227 Entering Passive Mode (%s)."
	ExtendedPassive = "229 Entering Extended Passive Mode (|||%s|)."
	LoggedIn        = "230 User %s logged in, proceed."
	FileActionOK    = "250 Requested file action okay, completed."
	PathCreated     = "257 \"%s\" created."

	UserOK      = "331 User name okay, need password."
	NeedAccount = "332 Need account for login."
	PendingMore = "350 Requested file action pending more information."

	ServiceNotAvail    = "421 Service not available, closing control connection."
	DataSessionNotOpen = "425 Can't open data connection."
	SessionClosed      = "426 Sessionection closed; transfer aborted."
	TransientFile      = "450 Requested file action not taken. File unavailable."
	LocalError         = "451 Requested action aborted. Local error in processing."
	NoSpace            = "452 Requested action not taken. Insufficient storage space."

	BadCommand       = "500 Syntax error, command unrecognized."
	BadArguments     = "501 Syntax error in parameters or arguments."
	NotImplemented   = "502 Command not implemented."
	BadSequence      = "503 Bad sequence of commands."
	ParameterNotImpl = "504 Command not implemented for that parameter."
	NotLoggedIn      = "530 Not logged in."
	NeedAccountStore = "532 Need account for storing files."
	FileUnavailable  = "550 Requested action not taken. File unavailable."
	PageTypeUnknown  = "551 Requested action aborted. Page type unknown."
	ExceededStorage  = "552 Requested file action aborted. Exceeded storage allocation."
	BadFileName      = "553 Requested action not taken. File name not allowed."
)
