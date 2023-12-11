package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Type    string
	Amount  float64
	Balance float64
	Date    time.Time
}

type Card struct {
	Type   string
	Status string
	Info   map[string]string
}

type Wallet struct {
	Balance      float64
	Transactions []Transaction
	IDCard       Card
	DrivingLicense Card
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0,
		Transactions: []Transaction{},
		IDCard: Card{
			Type:   "ID Card",
			Status: "present",
			Info: map[string]string{
				"Name":      "John Doe",
				"Birthdate": "01/01/1990",
			},
		},
		DrivingLicense: Card{
			Type:   "Driving License",
			Status: "present",
			Info: map[string]string{
				"Name":           "John Doe",
				"Birthdate":      "01/01/1990",
				"License Number": "ABC123",
			},
		},
	}
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, Transaction{Type: "Add Money", Amount: amount, Balance: w.Balance, Date: time.Now()})
}

func (w *Wallet) WithdrawMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Insufficient balance to withdraw this amount.")
		return
	}

	w.Balance -= amount
	w.Transactions = append(w.Transactions, Transaction{Type: "Withdraw Money", Amount: amount, Balance: w.Balance, Date: time.Now()})
}

func (w *Wallet) ShowBalance() {
	fmt.Printf("Current balance: %.2f MGA\n", w.Balance)
}

func (w *Wallet) DisplayHistory() {
	fmt.Println("Transaction history:")
	for i, transaction := range w.Transactions {
		formattedDate := transaction.Date.Format("02/01/2006 15:04:05")
		fmt.Printf("%d. %s: %.2f MGA (Balance: %.2f MGA) - Date: %s\n",
			i+1, transaction.Type, transaction.Amount, transaction.Balance, formattedDate)
	}
}

func (w *Wallet) DisplayCardsInfo() {
	fmt.Println("===================================")
	fmt.Println("Card information:")
	fmt.Printf("1. %s: Status - %s\n", w.IDCard.Type, w.IDCard.Status)
	w.displayCardInfo(w.IDCard.Info)
	fmt.Println("===================================")
	fmt.Printf("2. %s: Status - %s\n", w.DrivingLicense.Type, w.DrivingLicense.Status)
	w.displayCardInfo(w.DrivingLicense.Info)
	fmt.Println("===================================")

}

func (w *Wallet) displayCardInfo(info map[string]string) {
	fmt.Println("Card Details:")
	for key, value := range info {
		fmt.Printf("- %s: %s\n", key, value)
	}
}

func (w *Wallet) WithdrawCard(cardType string) {
	card := w.getCardByName(cardType)
	if card != nil {
		if card.Status == "present" {
			fmt.Printf("Withdrawing %s.\n", card.Type)
			card.Status = "withdrawn"
		} else {
			fmt.Printf("%s already withdrawn.\n", card.Type)
		}
	} else {
		fmt.Println("Invalid card type.")
	}
}

func (w *Wallet) PutBackCard(cardType string) {
	card := w.getCardByName(cardType)
	if card != nil {
		if card.Status == "withdrawn" {
			fmt.Printf("Putting back %s.\n", card.Type)
			card.Status = "present"
		} else {
			fmt.Printf("%s already present.\n", card.Type)
		}
	} else {
		fmt.Println("Invalid card type.")
	}
}

func (w *Wallet) getCardByName(cardType string) *Card {
	switch cardType {
	case "ID Card":
		return &w.IDCard
	case "Driving License":
		return &w.DrivingLicense
	default:
		return nil
	}
}

func DisplayMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Add Money")
	fmt.Println("2. Withdraw Money")
	fmt.Println("3. Show Balance")
	fmt.Println("4. Display History")
	fmt.Println("5. Display Cards Info")
	fmt.Println("6. Withdraw ID Card")
	fmt.Println("7. Withdraw Driving License")
	fmt.Println("8. Put Back ID Card")
	fmt.Println("9. Put Back Driving License")
	fmt.Println("10. Exit")
}

func StartMenu(wallet *Wallet) {
	reader := bufio.NewReader(os.Stdin)

	for {
		DisplayMenu()
		fmt.Print("Select an option (1-10): ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Enter the amount to add: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
			wallet.AddMoney(amount)
		case "2":
			fmt.Print("Enter the amount to withdraw: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
			wallet.WithdrawMoney(amount)
		case "3":
			wallet.ShowBalance()
		case "4":
			wallet.DisplayHistory()
		case "5":
			wallet.DisplayCardsInfo()
		case "6":
			wallet.WithdrawCard("ID Card")
		case "7":
			wallet.WithdrawCard("Driving License")
		case "8":
			wallet.PutBackCard("ID Card")
		case "9":
			wallet.PutBackCard("Driving License")
		case "10":
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please enter a number from 1 to 10.")
		}
	}
}

func main() {
	wallet := NewWallet()
	StartMenu(wallet)
}