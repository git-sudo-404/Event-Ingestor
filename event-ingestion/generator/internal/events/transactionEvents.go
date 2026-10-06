/*
 * MIT License
 *
 * Copyright (c) 2026 git-sudo-404
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * Of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * Copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * Copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package events

type TransactionType string

const (
	TransactionOrderCreated      TransactionType = "order_created"
	TransactionPaymentCompleted  TransactionType = "payment_completed"
	TransactionPaymentFailed     TransactionType = "payment_failed"
	TransactionRefund            TransactionType = "refund"
	TransactionShipmentCreated   TransactionType = "shipment_created"
	TransactionShipmentDelivered TransactionType = "shipment_delivered"
)

type TransactionStatus string

const (
	TransactionPending   TransactionStatus = "PENDING"
	TransactionCompleted TransactionStatus = "COMPLETED"
	TransactionFailed    TransactionStatus = "FAILED"
	TransactionRefunded  TransactionStatus = "REFUNDED"
)

type TransactionEvent struct {
	EventMetaData

	TransactionID string            `json:"transaction_id"`
	Type          TransactionType   `json:"type"`
	Amount        *float64          `json:"amount,omitempty"`
	Currency      *string           `json:"currency,omitempty"`
	Status        TransactionStatus `json:"status"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

func (te *TransactionEvent) SetTransactionID(TransactionID string) *TransactionEvent {
	te.TransactionID = TransactionID
	return te
}

func (te *TransactionEvent) SetType(Type TransactionType) *TransactionEvent {
	te.Type = Type
	return te
}

func (te *TransactionEvent) SetAmount(Amount float64) *TransactionEvent {
	te.Amount = &Amount
	return te
}

func (te *TransactionEvent) SetCurrency(Currency string) *TransactionEvent {
	te.Currency = &Currency
	return te
}

func (te *TransactionEvent) SetStatus(Status TransactionStatus) *TransactionEvent {
	te.Status = Status
	return te
}

func (te *TransactionEvent) SetMetadata(Metadata map[string]string) *TransactionEvent {
	te.Metadata = Metadata
	return te
}
