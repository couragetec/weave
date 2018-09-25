package ticker

import (
	"regexp"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/x/nft"
)

var _ weave.Msg = (*IssueTokenMsg)(nil)

const (
	pathIssueTokenMsg = "nft/ticker/issue"
)

var (
	//todo: revisit pattern
	IsValidID = regexp.MustCompile(`^[A-Z0-9]{3,4}$`).MatchString
)

// Path returns the routing path for this message
func (*IssueTokenMsg) Path() string {
	return pathIssueTokenMsg
}

func (i *IssueTokenMsg) Validate() error {
	if i == nil {
		return errors.ErrInternal("must not be nil")
	}
	if err := weave.Address(i.Owner).Validate(); err != nil {
		return err
	}

	if !IsValidID(string(i.Id)) {
		return nft.ErrInvalidID()
	}
	if err := i.Details.Validate(); err != nil {
		return err
	}
	// TODO: impl proper approval validation
	return nil
}