function displayMenu() {
  console.log("\nMenu:");
  console.log("1. Add Money");
  console.log("2. Withdraw Money");
  console.log("3. Show Balance");
  console.log("4. Display History");
  console.log("5. Display Cards Info");
  console.log("6. Withdraw ID Card");
  console.log("7. Withdraw Driving License");
  console.log("8. Put Back ID Card");
  console.log("9. Put Back Driving License");
  console.log("10. Exit");
}

function startMenu(wallet, rl) {
  displayMenu();
  rl.question("Select an option (1-10): ", function (option) {
    switch (option) {
      case '1':
        rl.question("Enter the amount to add: ", function (amount) {
          wallet.addMoney(parseFloat(amount));
          startMenu(wallet, rl);
        });
        break;

      case '2':
        rl.question("Enter the amount to withdraw: ", function (amount) {
          wallet.withdrawMoney(parseFloat(amount));
          startMenu(wallet, rl);
        });
        break;

      case '3':
        wallet.showBalance();
        startMenu(wallet, rl);
        break;

      case '4':
        wallet.displayHistory();
        startMenu(wallet, rl);
        break;

      case '5':
        wallet.displayCardsInfo();
        startMenu(wallet, rl);
        break;

      case '6':
        wallet.withdrawCard("ID Card");
        startMenu(wallet, rl);
        break;

      case '7':
        wallet.withdrawCard("Driving License");
        startMenu(wallet, rl);
        break;

      case '8':
        wallet.putBackCard("ID Card");
        startMenu(wallet, rl);
        break;

      case '9':
        wallet.putBackCard("Driving License");
        startMenu(wallet, rl);
        break;

      case '10':
        rl.close();
        break;

      default:
        console.log("Invalid option. Please enter a number from 1 to 10.");
        startMenu(wallet, rl);
        break;
    }
  });
}

module.exports = { displayMenu, startMenu };
