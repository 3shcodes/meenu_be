package models

import (
   time "time"
)
// BillSummary represents a summary of a bill.
type BillSummary struct {
    id          int
    billedOn    time.Time
    billStatus  string
    totalAmount float64
    amountPaid  float64
}

// Getters
func (bs *BillSummary) ID() int {
    return bs.id
}

func (bs *BillSummary) BilledOn() time.Time {
    return bs.billedOn
}

func (bs *BillSummary) BillStatus() string {
    return bs.billStatus
}

func (bs *BillSummary) TotalAmount() float64 {
    return bs.totalAmount
}

func (bs *BillSummary) AmountPaid() float64 {
    return bs.amountPaid
}

// Setters
func (bs *BillSummary) SetID(id int) {
    bs.id = id
}

func (bs *BillSummary) SetBilledOn(billedOn time.Time) {
    bs.billedOn = billedOn
}

func (bs *BillSummary) SetBillStatus(status string) {
    bs.billStatus = status
}

func (bs *BillSummary) SetTotalAmount(totalAmount float64) {
    bs.totalAmount = totalAmount
}

func (bs *BillSummary) SetAmountPaid(amountPaid float64) {
    bs.amountPaid = amountPaid
}

