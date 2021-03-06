// Copyright 2019 Bruno Miguel Custodio
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Entity represents a party involved in a payment.
type Entity struct {
	// AccountNumber is the account number for the entity.
	AccountNumber string `bson:"account_number" json:"account_number"`
	// BankID is the bank ID for the entity.
	BankID string `bson:"bank_id" json:"bank_id"`
	// Name is the name of the entity.
	Name string `bson:"name" json:"name"`
}

// Validate validates the current Entity object.
func (e *Entity) Validate() error {
	if e.AccountNumber == "" {
		return errors.New("the entity's account number must not be empty")
	}
	if e.BankID == "" {
		return errors.New("the entity's bank id must not be empty")
	}
	if e.Name == "" {
		return errors.New("the entity's name must not be empty")
	}
	return nil
}

// Payment represents a payment to an entity (the beneficiary) made by another entity (the debtor).
type Payment struct {
	// ID is the ID of the payments.
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// UpdatedAt is the record's modification date.
	UpdatedAt time.Time `bson:"updated_at" json:"-"`
	// DeletedAt is the record's deletion date.
	DeletedAt *time.Time `bson:"deleted_at" json:"-"`

	// Beneficiary is the entity that received the payment.
	Beneficiary Entity `bson:"beneficiary" json:"beneficiary"`
	// Debtor is the entity that sent the payment.
	Debtor Entity `bson:"debtor" json:"debtor"`

	// Amount is the amount involved in the payment.
	// It is a required field.
	Amount float64 `bson:"amount" json:"amount"`
	// Currency is the currency in which the payment was made.
	// It is a required field.
	Currency string `bson:"currency" json:"currency"`
	// Date is the date at which the payment was processed.
	// It is a required field.
	Date time.Time `bson:"date" json:"date"`
	// Description is the description associated with the payment.
	// It is an optional field.
	Description string `bson:"description" json:"description"`
}

// Validate validates the current Payment object.
func (p *Payment) Validate() error {
	if err := p.Beneficiary.Validate(); err != nil {
		return fmt.Errorf("beneficiary: %v", err)
	}
	if err := p.Debtor.Validate(); err != nil {
		return fmt.Errorf("debtor: %v", err)
	}
	if p.Amount <= 0 {
		return errors.New("the amount must be positive")
	}
	if p.Currency == "" {
		return errors.New("the currency must not be empty")
	}
	if p.Date.IsZero() {
		return errors.New("the date must not be empty")
	}
	if p.Description == "" {
		return errors.New("the description must not be empty")
	}
	return nil
}
