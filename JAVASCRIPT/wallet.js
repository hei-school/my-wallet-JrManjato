// wallet.js
const { formatDate } = require('./utils');
const Card = require('./card');

class Wallet {
  constructor() {
    this.balance = 0;
    this.actions = [];
    this.idCard = new Card("ID Card", "present", { name: 'John Doe', birthdate: '01/01/1990' });
    this.drivingLicense = new Card("Driving License", "present", { name: 'John Doe', birthdate: '01/01/1990', licence: 'A | B | D' });
  }

  addMoney(amount) {
    this.balance += amount;
    this.actions.push({ type: 'Add Money', amount, balance: this.balance, date: new Date() });
  }

  withdrawMoney(amount) {
    if (amount > this.balance) {
      console.log("Insufficient balance to withdraw this amount.");
      return;
    }

    this.balance -= amount;
    this.actions.push({ type: 'Withdraw Money', amount, balance: this.balance, date: new Date() });
  }

  showBalance() {
    console.log(`Current balance: ${this.balance} MGA`);
  }

  displayHistory() {
    console.log("Action history:");
    this.actions.forEach((action, index) => {
      const formattedDate = formatDate(action.date);
      console.log(`${index + 1}. ${action.type}: ${action.amount} MGA (Balance: ${action.balance} MGA) - Date: ${formattedDate}`);
    });
  }

  displayCardsInfo() {
    this.idCard.displayInfo();
    this.drivingLicense.displayInfo();
  }

  withdrawCard(type) {
    if (type === "ID Card") {
      this.idCard.withdraw();
    } else if (type === "Driving License") {
      this.drivingLicense.withdraw();
    } else {
      console.log("Unrecognized card type.");
    }
  }

  putBackCard(type) {
    if (type === "ID Card") {
      this.idCard.putBack();
    } else if (type === "Driving License") {
      this.drivingLicense.putBack();
    } else {
      console.log("Unrecognized card type.");
    }
  }
}

module.exports = Wallet;