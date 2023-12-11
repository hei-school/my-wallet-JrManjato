from card import Card
from datetime import datetime
from utils import format_date

class Wallet:
    def __init__(self):
        self.balance = 0
        self.actions = []
        self.id_card = Card("ID Card", "present", {"Name": "John Doe", "Birthdate": "01/01/1990"})
        self.driving_license = Card("Driving License", "present", {"Name": "John Doe", "Birthdate": "01/01/1990", "License Number": "ABC123"})


    def add_money(self, amount):
        self.balance += amount
        self.actions.append({'type': 'Add Money', 'amount': amount, 'balance': self.balance, 'date': datetime.now()})

    def withdraw_money(self, amount):
        if amount > self.balance:
            print("Insufficient balance to withdraw this amount.")
            return

        self.balance -= amount
        self.actions.append({'type': 'Withdraw Money', 'amount': amount, 'balance': self.balance, 'date': datetime.now()})

    def show_balance(self):
        print(f"Current balance: {self.balance} MGA")

    def display_history(self):
        print("Action history:")
        for index, action in enumerate(self.actions, start=1):
            formatted_date = format_date(action['date'])
            print(f"{index}. {action['type']}: {action['amount']} MGA (Balance: {action['balance']} MGA) - Date: {formatted_date}")

    def display_cards_info(self):
        self.id_card.display_info()
        self.driving_license.display_info()

    def withdraw_card(self, card_type):
        if card_type == "ID Card":
            self.id_card.withdraw()
        elif card_type == "Driving License":
            self.driving_license.withdraw()
        else:
            print("Unrecognized card type.")

    def put_back_card(self, card_type):
        if card_type == "ID Card":
            self.id_card.put_back()
        elif card_type == "Driving License":
            self.driving_license.put_back()
        else:
            print("Unrecognized card type.")