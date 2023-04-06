package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidateAuctionMsg validates that the MsgAuctionBid can be included in the auction.
func (k Keeper) ValidateAuctionMsg(ctx sdk.Context, bidder sdk.AccAddress, bid, highestBid sdk.Coin, transactions []sdk.Tx) error {
	// Validate the bundle size.
	maxBundleSize, err := k.GetMaxBundleSize(ctx)
	if err != nil {
		return err
	}

	if uint32(len(transactions)) > maxBundleSize {
		return fmt.Errorf("bundle size (%d) exceeds max bundle size (%d)", len(transactions), maxBundleSize)
	}

	// Validate the bid amount.
	if err := k.ValidateAuctionBid(ctx, bidder, bid, highestBid); err != nil {
		return err
	}

	// Validate the bundle of transactions if front-running protection is enabled.
	protectionEnabled, err := k.FrontRunningProtectionEnabled(ctx)
	if err != nil {
		return err
	}

	if protectionEnabled {
		if err := k.ValidateAuctionBundle(bidder, transactions); err != nil {
			return err
		}
	}

	return nil
}

// ValidateAuctionBid validates that the bidder has sufficient funds to participate in the auction and that the bid amount
// is sufficiently high enough.
func (k Keeper) ValidateAuctionBid(ctx sdk.Context, bidder sdk.AccAddress, bid, highestBid sdk.Coin) error {
	if bid.IsNil() {
		return fmt.Errorf("bid amount cannot be nil")
	}

	// Get the bid floor.
	reserveFee, err := k.GetReserveFee(ctx)
	if err != nil {
		return err
	}

	// Ensure that the bid denomination matches the fee denominations.
	if bid.Denom != reserveFee.Denom {
		return fmt.Errorf("bid denom (%s) does not match the reserve fee denom (%s)", bid, reserveFee)
	}

	// Bid must be greater than the bid floor.
	if !bid.IsGTE(reserveFee) {
		return fmt.Errorf("bid amount (%s) is less than the reserve fee (%s)", bid, reserveFee)
	}

	if !highestBid.IsNil() {
		// Ensure the bid is greater than the highest bid + min bid increment.
		minBidIncrement, err := k.GetMinBidIncrement(ctx)
		if err != nil {
			return err
		}

		minBid := highestBid.Add(minBidIncrement)
		if !bid.IsGTE(minBid) {
			return fmt.Errorf("bid amount (%s) is less than the highest bid (%s) + min bid increment (%s)", bid, highestBid, minBidIncrement)
		}
	}

	// Get the pay-to-play fee.
	minBuyInFee, err := k.GetMinBuyInFee(ctx)
	if err != nil {
		return err
	}

	// Ensure the bidder has enough funds to cover all the inclusion fees.
	minBalance := bid.Add(minBuyInFee)
	balances := k.bankKeeper.GetAllBalances(ctx, bidder)
	if !balances.IsAllGTE(sdk.NewCoins(minBalance)) {
		return fmt.Errorf("insufficient funds to bid %s (reserve fee + bid) with balance %s", minBalance, balances)
	}

	return nil
}

// ValidateAuctionBundle validates the ordering of the referenced transactions. Bundles are valid if
//  1. all of the transactions are signed by the signer.
//  2. some subset of contiguous transactions starting from the first tx are signed by the same signer, and all other tranasctions
//     are signed by the bidder.
//
// example:
//  1. valid: [tx1, tx2, tx3] where tx1 is signed by the signer 1 and tx2 and tx3 are signed by the bidder.
//  2. valid: [tx1, tx2, tx3, tx4] where tx1 - tx4 are signed by the bidder.
//  3. invalid: [tx1, tx2, tx3] where tx1 and tx3 are signed by the bidder and tx2 is signed by some other signer. (possible sandwich attack)
//  4. invalid: [tx1, tx2, tx3] where tx1 is signed by the bidder, and tx2 - tx3 are signed by some other signer. (possible front-running attack)
func (k Keeper) ValidateAuctionBundle(bidder sdk.AccAddress, transactions []sdk.Tx) error {
	if len(transactions) <= 1 {
		return nil
	}

	// prevSigners is used to track whether the signers of the current transaction overlap.
	prevSigners, err := k.getTxSigners(transactions[0])
	if err != nil {
		return err
	}
	seenBidder := prevSigners[bidder.String()]

	// Check that all subsequent transactions are signed by either
	// 1. the same party as the first transaction
	// 2. the same party for some arbitrary number of txs and then are all remaining txs are signed by the bidder.
	for _, refTx := range transactions[1:] {
		txSigners, err := k.getTxSigners(refTx)
		if err != nil {
			return err
		}

		// Filter the signers to only those that signed the current transaction.
		filterSigners(prevSigners, txSigners)

		// If there are no overlapping signers from the previous tx and the bidder address has not been seen, then the bundle can still be valid
		// as long as all subsequent transactions are signed by the bidder.
		if len(prevSigners) == 0 {
			if seenBidder {
				return fmt.Errorf("bundle contains transactions signed by multiple parties; possible front-running or sandwich attack")
			}

			seenBidder = true
			prevSigners = map[string]bool{bidder.String(): true}
			filterSigners(prevSigners, txSigners)

			if len(prevSigners) == 0 {
				return fmt.Errorf("bundle contains transactions signed by multiple parties; possible front-running or sandwich attack")
			}
		}
	}

	return nil
}

// getTxSigners returns the signers of a transaction.
func (k Keeper) getTxSigners(tx sdk.Tx) (map[string]bool, error) {
	signers := make(map[string]bool, 0)
	for _, msg := range tx.GetMsgs() {
		for _, signer := range msg.GetSigners() {
			// TODO: check for multi-sig accounts
			// https://github.com/skip-mev/pob/issues/14
			signers[signer.String()] = true
		}
	}

	return signers, nil
}

// filterSigners removes any signers from the currentSigners map that are not in the txSigners map.
func filterSigners(currentSigners, txSigners map[string]bool) {
	for signer := range currentSigners {
		if _, ok := txSigners[signer]; !ok {
			delete(currentSigners, signer)
		}
	}
}