class Card {
    constructor(type, status, info = {}) {
      this.type = type;
      this.status = status;
      this.present = status !== 'lost';
      this.info = info;
    }
  
    displayInfo() {
      if (this.status === 'lost') {
        console.log(`You have lost your ${this.type}.`);
      } else if (this.present) {
        console.log('===================================');
        console.log(`${this.type}: Status - ${this.status}`);
        console.log(`Card Information:`);
        Object.entries(this.info).forEach(([key, value]) => {
          console.log(`${key}: ${value}`);
        });
        console.log('===================================');
      } else {
        console.log(`The ${this.type} has been withdrawn. Please put it back.`);
      }
    }
  
    withdraw() {
      if (this.present) {
        console.log(`Withdrawing ${this.type}.`);
        this.present = false;
      } else {
        console.log(`${this.type} already withdrawn.`);
      }
    }
  
    putBack() {
      if (!this.present) {
        console.log(`Putting back ${this.type}.`);
        this.present = true;
      } else {
        console.log(`${this.type} already present.`);
      }
    }
  }
  
  module.exports = Card;  