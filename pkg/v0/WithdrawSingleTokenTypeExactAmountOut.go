// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package spl_token_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawSingleTokenTypeExactAmountOut is the `withdrawSingleTokenTypeExactAmountOut` instruction.
type WithdrawSingleTokenTypeExactAmountOut struct {
	DestinationTokenAmount *uint64
	MaximumPoolTokenAmount *uint64

	// [0] = [] swap
	//
	// [1] = [] authority
	//
	// [2] = [SIGNER] userTransferAuthority
	//
	// [3] = [WRITE] poolMint
	//
	// [4] = [WRITE] poolTokenSource
	//
	// [5] = [WRITE] swapTokenA
	//
	// [6] = [WRITE] swapTokenB
	//
	// [7] = [WRITE] destination
	//
	// [8] = [WRITE] feeAccount
	//
	// [9] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type WithdrawSingleTokenTypeExactAmountOutAccounts struct {
	Swap                  ag_solanago.PublicKey
	Authority             ag_solanago.PublicKey
	UserTransferAuthority ag_solanago.PublicKey
	PoolMint              ag_solanago.PublicKey
	PoolTokenSource       ag_solanago.PublicKey
	SwapTokenA            ag_solanago.PublicKey
	SwapTokenB            ag_solanago.PublicKey
	Destination           ag_solanago.PublicKey
	FeeAccount            ag_solanago.PublicKey
	TokenProgram          ag_solanago.PublicKey
}

// NewWithdrawSingleTokenTypeExactAmountOutInstructionBuilder creates a new `WithdrawSingleTokenTypeExactAmountOut` instruction builder.
func NewWithdrawSingleTokenTypeExactAmountOutInstructionBuilder() *WithdrawSingleTokenTypeExactAmountOut {
	nd := &WithdrawSingleTokenTypeExactAmountOut{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

func (inst *WithdrawSingleTokenTypeExactAmountOut) GetWithdrawSingleTokenTypeExactAmountOutAccounts() *WithdrawSingleTokenTypeExactAmountOutAccounts {
	res := &WithdrawSingleTokenTypeExactAmountOutAccounts{}
	if inst.AccountMetaSlice[0] != nil {
		res.Swap = inst.AccountMetaSlice[0].PublicKey
	}
	if inst.AccountMetaSlice[1] != nil {
		res.Authority = inst.AccountMetaSlice[1].PublicKey
	}
	if inst.AccountMetaSlice[2] != nil {
		res.UserTransferAuthority = inst.AccountMetaSlice[2].PublicKey
	}
	if inst.AccountMetaSlice[3] != nil {
		res.PoolMint = inst.AccountMetaSlice[3].PublicKey
	}
	if inst.AccountMetaSlice[4] != nil {
		res.PoolTokenSource = inst.AccountMetaSlice[4].PublicKey
	}
	if inst.AccountMetaSlice[5] != nil {
		res.SwapTokenA = inst.AccountMetaSlice[5].PublicKey
	}
	if inst.AccountMetaSlice[6] != nil {
		res.SwapTokenB = inst.AccountMetaSlice[6].PublicKey
	}
	if inst.AccountMetaSlice[7] != nil {
		res.Destination = inst.AccountMetaSlice[7].PublicKey
	}
	if inst.AccountMetaSlice[8] != nil {
		res.FeeAccount = inst.AccountMetaSlice[8].PublicKey
	}
	if inst.AccountMetaSlice[9] != nil {
		res.TokenProgram = inst.AccountMetaSlice[9].PublicKey
	}
	return res
}

// SetDestinationTokenAmount sets the "destinationTokenAmount" parameter.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetDestinationTokenAmount(destinationTokenAmount uint64) *WithdrawSingleTokenTypeExactAmountOut {
	inst.DestinationTokenAmount = &destinationTokenAmount
	return inst
}

// SetMaximumPoolTokenAmount sets the "maximumPoolTokenAmount" parameter.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetMaximumPoolTokenAmount(maximumPoolTokenAmount uint64) *WithdrawSingleTokenTypeExactAmountOut {
	inst.MaximumPoolTokenAmount = &maximumPoolTokenAmount
	return inst
}

// SetSwapAccount sets the "swap" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetSwapAccount(swap ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swap)
	return inst
}

// GetSwapAccount gets the "swap" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetSwapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserTransferAuthorityAccount sets the "userTransferAuthority" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetUserTransferAuthorityAccount(userTransferAuthority ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	return inst
}

// GetUserTransferAuthorityAccount gets the "userTransferAuthority" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetUserTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolMintAccount sets the "poolMint" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetPoolMintAccount(poolMint ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

// GetPoolMintAccount gets the "poolMint" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetPoolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPoolTokenSourceAccount sets the "poolTokenSource" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetPoolTokenSourceAccount(poolTokenSource ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(poolTokenSource).WRITE()
	return inst
}

// GetPoolTokenSourceAccount gets the "poolTokenSource" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetPoolTokenSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSwapTokenAAccount sets the "swapTokenA" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetSwapTokenAAccount(swapTokenA ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(swapTokenA).WRITE()
	return inst
}

// GetSwapTokenAAccount gets the "swapTokenA" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetSwapTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSwapTokenBAccount sets the "swapTokenB" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetSwapTokenBAccount(swapTokenB ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(swapTokenB).WRITE()
	return inst
}

// GetSwapTokenBAccount gets the "swapTokenB" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetSwapTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetDestinationAccount sets the "destination" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetDestinationAccount(destination ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetFeeAccountAccount sets the "feeAccount" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetFeeAccountAccount(feeAccount ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(feeAccount).WRITE()
	return inst
}

// GetFeeAccountAccount gets the "feeAccount" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *WithdrawSingleTokenTypeExactAmountOut) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst WithdrawSingleTokenTypeExactAmountOut) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawSingleTokenTypeExactAmountOut,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawSingleTokenTypeExactAmountOut) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawSingleTokenTypeExactAmountOut) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DestinationTokenAmount == nil {
			return errors.New("DestinationTokenAmount parameter is not set")
		}
		if inst.MaximumPoolTokenAmount == nil {
			return errors.New("MaximumPoolTokenAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Swap is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PoolTokenSource is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SwapTokenA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SwapTokenB is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Destination is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.FeeAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawSingleTokenTypeExactAmountOut) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawSingleTokenTypeExactAmountOut")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DestinationTokenAmount", *inst.DestinationTokenAmount))
						paramsBranch.Child(ag_format.Param("MaximumPoolTokenAmount", *inst.MaximumPoolTokenAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 swap", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("userTransferAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             poolMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      poolTokenSource", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           swapTokenA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           swapTokenB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          destination", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                  fee", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj WithdrawSingleTokenTypeExactAmountOut) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestinationTokenAmount` param:
	err = encoder.Encode(obj.DestinationTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `MaximumPoolTokenAmount` param:
	err = encoder.Encode(obj.MaximumPoolTokenAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawSingleTokenTypeExactAmountOut) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestinationTokenAmount`:
	err = decoder.Decode(&obj.DestinationTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `MaximumPoolTokenAmount`:
	err = decoder.Decode(&obj.MaximumPoolTokenAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawSingleTokenTypeExactAmountOutInstruction declares a new WithdrawSingleTokenTypeExactAmountOut instruction with the provided parameters and accounts.
func NewWithdrawSingleTokenTypeExactAmountOutInstruction(
	// Parameters:
	destinationTokenAmount uint64,
	maximumPoolTokenAmount uint64,
	// Accounts:
	swap ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	userTransferAuthority ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	poolTokenSource ag_solanago.PublicKey,
	swapTokenA ag_solanago.PublicKey,
	swapTokenB ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	feeAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *WithdrawSingleTokenTypeExactAmountOut {
	return NewWithdrawSingleTokenTypeExactAmountOutInstructionBuilder().
		SetDestinationTokenAmount(destinationTokenAmount).
		SetMaximumPoolTokenAmount(maximumPoolTokenAmount).
		SetSwapAccount(swap).
		SetAuthorityAccount(authority).
		SetUserTransferAuthorityAccount(userTransferAuthority).
		SetPoolMintAccount(poolMint).
		SetPoolTokenSourceAccount(poolTokenSource).
		SetSwapTokenAAccount(swapTokenA).
		SetSwapTokenBAccount(swapTokenB).
		SetDestinationAccount(destination).
		SetFeeAccountAccount(feeAccount).
		SetTokenProgramAccount(tokenProgram)
}