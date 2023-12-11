import java.util.ArrayList;
import java.util.List;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

public class Wallet {
    private double balance;
    private List<Transaction> transactions;
    private Card idCard;
    private Card drivingLicense;

    public Wallet() {
        this.balance = 0;
        this.transactions = new ArrayList<>();
        this.idCard = new Card("ID Card", "present", getDefaultCardInfo("ID Card"));
        this.drivingLicense = new Card("Driving License", "present", getDefaultCardInfo("Driving License"));
    }

    private Map<String, String> getDefaultCardInfo(String cardType) {
        Map<String, String> defaultInfo = new HashMap<>();
        if (cardType.equals("ID Card")) {
            defaultInfo.put("Name", "John Doe");
            defaultInfo.put("Birthdate", "01/01/1990");
        } else if (cardType.equals("Driving License")) {
            defaultInfo.put("Name", "John Doe");
            defaultInfo.put("Birthdate", "01/01/1990");
            defaultInfo.put("License Number", "ABC123");
        }
        return defaultInfo;
    }

    public void addMoney(double amount) {
        this.balance += amount;
        this.transactions.add(new Transaction("Add Money", amount, this.balance));
    }

    public void withdrawMoney(double amount) {
        if (amount > this.balance) {
            System.out.println("Insufficient balance to withdraw this amount.");
            return;
        }

        this.balance -= amount;
        this.transactions.add(new Transaction("Withdraw Money", amount, this.balance));
    }

    public void showBalance() {
        System.out.printf("Current balance: %.2f MGA%n", this.balance);
    }

    public void displayHistory() {
        System.out.println("Transaction history:");
        for (int i = 0; i < transactions.size(); i++) {
            Transaction transaction = transactions.get(i);
            String formattedDate = Transaction.formatDate(transaction.getDate());
            System.out.printf("%d. %s: %.2f MGA (Balance: %.2f MGA) - Date: %s%n",
                    i + 1, transaction.getType(), transaction.getAmount(), transaction.getBalance(), formattedDate);
        }
    }

    public void displayCardsInfo() {
        idCard.displayInfo();
        drivingLicense.displayInfo();
    }

    public void withdrawCard(String type) {
        if (type.equals("ID Card")) {
            idCard.withdraw();
        } else if (type.equals("Driving License")) {
            drivingLicense.withdraw();
        } else {
            System.out.println("Unrecognized card type.");
        }
    }

    public void putBackCard(String type) {
        if (type.equals("ID Card")) {
            idCard.putBack();
        } else if (type.equals("Driving License")) {
            drivingLicense.putBack();
        } else {
            System.out.println("Unrecognized card type.");
        }
    }
}