// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package spl_token_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawAllTokenTypes is the `withdrawAllTokenTypes` instruction.
type WithdrawAllTokenTypes struct {
	PoolTokenAmount     *uint64
	MinimumTokenAAmount *uint64
	MinimumTokenBAmount *uint64

	// [0] = [] swap
	//
	// [1] = [] authority
	//
	// [2] = [SIGNER] userTransferAuthority
	//
	// [3] = [WRITE] poolMint
	//
	// [4] = [WRITE] source
	//
	// [5] = [WRITE] swapTokenA
	//
	// [6] = [WRITE] swapTokenB
	//
	// [7] = [WRITE] destinationTokenA
	//
	// [8] = [WRITE] destinationTokenB
	//
	// [9] = [WRITE] feeAccount
	//
	// [10] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type WithdrawAllTokenTypesAccounts struct {
	Swap                  ag_solanago.PublicKey
	Authority             ag_solanago.PublicKey
	UserTransferAuthority ag_solanago.PublicKey
	PoolMint              ag_solanago.PublicKey
	Source                ag_solanago.PublicKey
	SwapTokenA            ag_solanago.PublicKey
	SwapTokenB            ag_solanago.PublicKey
	DestinationTokenA     ag_solanago.PublicKey
	DestinationTokenB     ag_solanago.PublicKey
	FeeAccount            ag_solanago.PublicKey
	TokenProgram          ag_solanago.PublicKey
}

// NewWithdrawAllTokenTypesInstructionBuilder creates a new `WithdrawAllTokenTypes` instruction builder.
func NewWithdrawAllTokenTypesInstructionBuilder() *WithdrawAllTokenTypes {
	nd := &WithdrawAllTokenTypes{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

func (inst *WithdrawAllTokenTypes) GetWithdrawAllTokenTypesAccounts() *WithdrawAllTokenTypesAccounts {
	res := &WithdrawAllTokenTypesAccounts{}
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
		res.Source = inst.AccountMetaSlice[4].PublicKey
	}
	if inst.AccountMetaSlice[5] != nil {
		res.SwapTokenA = inst.AccountMetaSlice[5].PublicKey
	}
	if inst.AccountMetaSlice[6] != nil {
		res.SwapTokenB = inst.AccountMetaSlice[6].PublicKey
	}
	if inst.AccountMetaSlice[7] != nil {
		res.DestinationTokenA = inst.AccountMetaSlice[7].PublicKey
	}
	if inst.AccountMetaSlice[8] != nil {
		res.DestinationTokenB = inst.AccountMetaSlice[8].PublicKey
	}
	if inst.AccountMetaSlice[9] != nil {
		res.FeeAccount = inst.AccountMetaSlice[9].PublicKey
	}
	if inst.AccountMetaSlice[10] != nil {
		res.TokenProgram = inst.AccountMetaSlice[10].PublicKey
	}
	return res
}

// SetPoolTokenAmount sets the "poolTokenAmount" parameter.
func (inst *WithdrawAllTokenTypes) SetPoolTokenAmount(poolTokenAmount uint64) *WithdrawAllTokenTypes {
	inst.PoolTokenAmount = &poolTokenAmount
	return inst
}

// SetMinimumTokenAAmount sets the "minimumTokenAAmount" parameter.
func (inst *WithdrawAllTokenTypes) SetMinimumTokenAAmount(minimumTokenAAmount uint64) *WithdrawAllTokenTypes {
	inst.MinimumTokenAAmount = &minimumTokenAAmount
	return inst
}

// SetMinimumTokenBAmount sets the "minimumTokenBAmount" parameter.
func (inst *WithdrawAllTokenTypes) SetMinimumTokenBAmount(minimumTokenBAmount uint64) *WithdrawAllTokenTypes {
	inst.MinimumTokenBAmount = &minimumTokenBAmount
	return inst
}

// SetSwapAccount sets the "swap" account.
func (inst *WithdrawAllTokenTypes) SetSwapAccount(swap ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(swap)
	return inst
}

// GetSwapAccount gets the "swap" account.
func (inst *WithdrawAllTokenTypes) GetSwapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawAllTokenTypes) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawAllTokenTypes) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserTransferAuthorityAccount sets the "userTransferAuthority" account.
func (inst *WithdrawAllTokenTypes) SetUserTransferAuthorityAccount(userTransferAuthority ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	return inst
}

// GetUserTransferAuthorityAccount gets the "userTransferAuthority" account.
func (inst *WithdrawAllTokenTypes) GetUserTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolMintAccount sets the "poolMint" account.
func (inst *WithdrawAllTokenTypes) SetPoolMintAccount(poolMint ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

// GetPoolMintAccount gets the "poolMint" account.
func (inst *WithdrawAllTokenTypes) GetPoolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSourceAccount sets the "source" account.
func (inst *WithdrawAllTokenTypes) SetSourceAccount(source ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceAccount gets the "source" account.
func (inst *WithdrawAllTokenTypes) GetSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSwapTokenAAccount sets the "swapTokenA" account.
func (inst *WithdrawAllTokenTypes) SetSwapTokenAAccount(swapTokenA ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(swapTokenA).WRITE()
	return inst
}

// GetSwapTokenAAccount gets the "swapTokenA" account.
func (inst *WithdrawAllTokenTypes) GetSwapTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSwapTokenBAccount sets the "swapTokenB" account.
func (inst *WithdrawAllTokenTypes) SetSwapTokenBAccount(swapTokenB ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(swapTokenB).WRITE()
	return inst
}

// GetSwapTokenBAccount gets the "swapTokenB" account.
func (inst *WithdrawAllTokenTypes) GetSwapTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetDestinationTokenAAccount sets the "destinationTokenA" account.
func (inst *WithdrawAllTokenTypes) SetDestinationTokenAAccount(destinationTokenA ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(destinationTokenA).WRITE()
	return inst
}

// GetDestinationTokenAAccount gets the "destinationTokenA" account.
func (inst *WithdrawAllTokenTypes) GetDestinationTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDestinationTokenBAccount sets the "destinationTokenB" account.
func (inst *WithdrawAllTokenTypes) SetDestinationTokenBAccount(destinationTokenB ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(destinationTokenB).WRITE()
	return inst
}

// GetDestinationTokenBAccount gets the "destinationTokenB" account.
func (inst *WithdrawAllTokenTypes) GetDestinationTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetFeeAccountAccount sets the "feeAccount" account.
func (inst *WithdrawAllTokenTypes) SetFeeAccountAccount(feeAccount ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(feeAccount).WRITE()
	return inst
}

// GetFeeAccountAccount gets the "feeAccount" account.
func (inst *WithdrawAllTokenTypes) GetFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *WithdrawAllTokenTypes) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *WithdrawAllTokenTypes) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst WithdrawAllTokenTypes) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawAllTokenTypes,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawAllTokenTypes) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawAllTokenTypes) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PoolTokenAmount == nil {
			return errors.New("PoolTokenAmount parameter is not set")
		}
		if inst.MinimumTokenAAmount == nil {
			return errors.New("MinimumTokenAAmount parameter is not set")
		}
		if inst.MinimumTokenBAmount == nil {
			return errors.New("MinimumTokenBAmount parameter is not set")
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
			return errors.New("accounts.Source is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SwapTokenA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SwapTokenB is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.DestinationTokenA is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.DestinationTokenB is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.FeeAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawAllTokenTypes) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawAllTokenTypes")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    PoolTokenAmount", *inst.PoolTokenAmount))
						paramsBranch.Child(ag_format.Param("MinimumTokenAAmount", *inst.MinimumTokenAAmount))
						paramsBranch.Child(ag_format.Param("MinimumTokenBAmount", *inst.MinimumTokenBAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 swap", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("userTransferAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             poolMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               source", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           swapTokenA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           swapTokenB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("    destinationTokenA", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    destinationTokenB", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                  fee", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj WithdrawAllTokenTypes) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PoolTokenAmount` param:
	err = encoder.Encode(obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `MinimumTokenAAmount` param:
	err = encoder.Encode(obj.MinimumTokenAAmount)
	if err != nil {
		return err
	}
	// Serialize `MinimumTokenBAmount` param:
	err = encoder.Encode(obj.MinimumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawAllTokenTypes) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PoolTokenAmount`:
	err = decoder.Decode(&obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `MinimumTokenAAmount`:
	err = decoder.Decode(&obj.MinimumTokenAAmount)
	if err != nil {
		return err
	}
	// Deserialize `MinimumTokenBAmount`:
	err = decoder.Decode(&obj.MinimumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawAllTokenTypesInstruction declares a new WithdrawAllTokenTypes instruction with the provided parameters and accounts.
func NewWithdrawAllTokenTypesInstruction(
	// Parameters:
	poolTokenAmount uint64,
	minimumTokenAAmount uint64,
	minimumTokenBAmount uint64,
	// Accounts:
	swap ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	userTransferAuthority ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	source ag_solanago.PublicKey,
	swapTokenA ag_solanago.PublicKey,
	swapTokenB ag_solanago.PublicKey,
	destinationTokenA ag_solanago.PublicKey,
	destinationTokenB ag_solanago.PublicKey,
	feeAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	return NewWithdrawAllTokenTypesInstructionBuilder().
		SetPoolTokenAmount(poolTokenAmount).
		SetMinimumTokenAAmount(minimumTokenAAmount).
		SetMinimumTokenBAmount(minimumTokenBAmount).
		SetSwapAccount(swap).
		SetAuthorityAccount(authority).
		SetUserTransferAuthorityAccount(userTransferAuthority).
		SetPoolMintAccount(poolMint).
		SetSourceAccount(source).
		SetSwapTokenAAccount(swapTokenA).
		SetSwapTokenBAccount(swapTokenB).
		SetDestinationTokenAAccount(destinationTokenA).
		SetDestinationTokenBAccount(destinationTokenB).
		SetFeeAccountAccount(feeAccount).
		SetTokenProgramAccount(tokenProgram)
}