from wallet import Wallet
from utils import format_date

def display_menu():
    print("\nMenu:")
    print("1. Add Money")
    print("2. Withdraw Money")
    print("3. Show Balance")
    print("4. Display History")
    print("5. Display Cards Info")
    print("6. Withdraw ID Card")
    print("7. Withdraw Driving License")
    print("8. Put Back ID Card")
    print("9. Put Back Driving License")
    print("10. Exit")

def start_menu(wallet):
    while True:
        display_menu()
        option = input("Select an option (1-10): ")

        if option == '1':
            amount = float(input("Enter the amount to add: "))
            wallet.add_money(amount)
        elif option == '2':
            amount = float(input("Enter the amount to withdraw: "))
            wallet.withdraw_money(amount)
        elif option == '3':
            wallet.show_balance()
        elif option == '4':
            wallet.display_history()
        elif option == "5":
            wallet.display_cards_info()
        elif option == "6":
            wallet.withdraw_card("ID Card")
        elif option == "7":
            wallet.withdraw_card("Driving License")
        elif option == "8":
            wallet.put_back_card("ID Card")
        elif option == "9":
            wallet.put_back_card("Driving License")
        elif option == "10":
            break
        else:
            print("Invalid option. Please enter a number from 1 to 10.")