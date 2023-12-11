class Card:
    def __init__(self, card_type, status, info=None):
        self.type = card_type
        self.status = status
        self.present = status != "lost"
        self.info = info if info else {}

    def display_info(self):
        if self.status == "lost":
            print(f"You have lost your {self.type}.")
        elif self.present:
            print("===================================")
            print(f"{self.type}: Status - {self.status}")
            print("Card Information:")
            for key, value in self.info.items():
                print(f"{key}: {value}")
            print("===================================")
        else:
            print(f"The {self.type} has been withdrawn. Please put it back.")

    def withdraw(self):
        if self.present:
            print(f"Withdrawing {self.type}.")
            self.present = False
        else:
            print(f"{self.type} already withdrawn.")

    def put_back(self):
        if not self.present:
            print(f"Putting back {self.type}.")
            self.present = True
        else:
            print(f"{self.type} already present.")