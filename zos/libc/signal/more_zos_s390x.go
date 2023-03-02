// Copy of linux:s390x implementation

package signal

const (
	//	#define SIG_DFL	((__sighandler_t)0)	/* default signal handling */
	SIG_DFL = 0
	//	#define SIG_IGN	((__sighandler_t)1)	/* ignore signal */
	SIG_IGN = 1
)
