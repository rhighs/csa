package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	errorIota             uint32 = 1110
	ErrSample                    = sdkerrors.Register(ModuleName, errorIota+1, "sample error")
	ErrGameNotFound              = sdkerrors.Register(ModuleName, errorIota+2, "game by id not found")
	ErrCreatorNotPlayer          = sdkerrors.Register(ModuleName, errorIota+3, "message creator is not a player")
	ErrNotPlayerTurn             = sdkerrors.Register(ModuleName, errorIota+4, "player tried to play out of turn")
	ErrWrongMove                 = sdkerrors.Register(ModuleName, errorIota+5, "wrong move")
	ErrBlackAlreadyPlayed        = sdkerrors.Register(ModuleName, errorIota+6, "black player has already played")
	ErrRedAlreadyPlayed          = sdkerrors.Register(ModuleName, errorIota+7, "red player has already played")
)
