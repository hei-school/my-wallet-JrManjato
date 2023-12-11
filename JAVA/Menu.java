import java.util.Scanner;

public class Menu {
    public static void displayMenu() {
        System.out.println("\nMenu:");
        System.out.println("1. Add Money");
        System.out.println("2. Withdraw Money");
        System.out.println("3. Show Balance");
        System.out.println("4. Display History");
        System.out.println("5. Display Cards Info");
        System.out.println("6. Withdraw ID Card");
        System.out.println("7. Withdraw Driving License");
        System.out.println("8. Put Back ID Card");
        System.out.println("9. Put Back Driving License");
        System.out.println("10. Exit");
    }

    public static void startMenu(Wallet wallet) {
        Scanner scanner = new Scanner(System.in);

        while (true) {
            displayMenu();
            System.out.print("Select an option (1-10): ");
            String option = scanner.nextLine();

            switch (option) {
                case "1":
                    System.out.print("Enter the amount to add: ");
                    double amountToAdd = Double.parseDouble(scanner.nextLine());
                    wallet.addMoney(amountToAdd);
                    break;
                case "2":
                    System.out.print("Enter the amount to withdraw: ");
                    double amountToWithdraw = Double.parseDouble(scanner.nextLine());
                    wallet.withdrawMoney(amountToWithdraw);
                    break;
                case "3":
                    wallet.showBalance();
                    break;
                case "4":
                    wallet.displayHistory();
                    break;
                case "5":
                    wallet.displayCardsInfo();
                    break;
                case "6":
                    wallet.withdrawCard("ID Card");
                    break;
                case "7":
                    wallet.withdrawCard("Driving License");
                    break;
                case "8":
                    wallet.putBackCard("ID Card");
                    break;
                case "9":
                    wallet.putBackCard("Driving License");
                    break;
                case "10":
                    scanner.close();
                    break;
                default:
                    System.out.println("Invalid option. Please enter a number from 1 to 10.");
            }
        }
    }
}
