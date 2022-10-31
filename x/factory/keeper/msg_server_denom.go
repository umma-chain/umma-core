package keeper

import (
	"context"

	"github.com/umma-chain/core-umma/x/factory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(ctx, msg.Denom)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		SiteUrl:            msg.SiteUrl,
		LogoUrl:            msg.LogoUrl,
		MaxSupply:          msg.MaxSupply,
		Supply:             0,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	if !valFound.CanChangeMaxSupply && valFound.MaxSupply != msg.MaxSupply {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot change maxsupply")
	}
	if !valFound.CanChangeMaxSupply && msg.CanChangeMaxSupply {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot revert change maxsupply flag")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             valFound.Ticker,
		Precision:          valFound.Precision,
		SiteUrl:            msg.SiteUrl,
		LogoUrl:            msg.LogoUrl,
		MaxSupply:          msg.MaxSupply,
		Supply:             valFound.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}
