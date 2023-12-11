import java.util.HashMap;
import java.util.Map;
public class Card {
    private String type;
    private String status;
    private boolean present;
    private Map<String, String> info;

    public Card(String type, String status, Map<String, String> info) {
        this.type = type;
        this.status = status;
        this.present = !status.equals("lost");
        this.info = info;
    }

    public void displayInfo() {
        if (status.equals("lost")) {
            System.out.println("You have lost your " + type + ".");
        } else if (present) {
            System.out.println("===================================");
            System.out.println(type + ": Status - " + status);
            System.out.println("Card Information:");
            for (Map.Entry<String, String> entry : info.entrySet()) {
                System.out.println(entry.getKey() + ": " + entry.getValue());
            }
            System.out.println("===================================");
        } else {
            System.out.println("The " + type + " has been withdrawn. Please put it back.");
        }
    }

    public void withdraw() {
        if (present) {
            System.out.println("Withdrawing " + type + ".");
            present = false;
        } else {
            System.out.println(type + " already withdrawn.");
        }
    }

    public void putBack() {
        if (!present) {
            System.out.println("Putting back " + type + ".");
            present = true;
        } else {
            System.out.println(type + " already present.");
        }
    }
}