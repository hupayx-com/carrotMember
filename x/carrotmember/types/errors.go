package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/carrotmember module sentinel errors
var (
	ErrSample        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNotIntinalize = sdkerrors.Register(ModuleName, 1200, "Not intinalize error")
	ErrNotYetTime    = sdkerrors.Register(ModuleName, 1300, "Not yet reward time error")
)
