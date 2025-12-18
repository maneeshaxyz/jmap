package main

type Accounts map[string]Account

type Account struct {
	Name                string              `json:"name"`
	IsPersonal          bool                `json:"isPersonal"`
	IsReadOnly          bool                `json:"isReadOnly"`
	AccountCapabilities AccountCapabilities `json:"accountCapabilities"`
}

type AccountCapabilities map[string]any
