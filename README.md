# 🍜 Food Ordering Queue Application

This repository contains an implementation of a food ordering queue system converted from 3 JavaScript repositories into **Golang** using Goroutines, Channels, and WaitGroup.

---

## 📂 Code Structure
This application handles the following data:

| Variable Name | Data Type | Description |
| :--- | :--- | :--- |
| `Name` | `string` | Name of the food customer. |
| `Wait` | `int` | Waiting time in seconds. |

---

## 🛠️ Code Explanation

### 1. Channels (`chan`)
Channels are used as a medium to send data between Goroutines.

* **Declaration**: `queueChan := make(chan Person)` creates a dedicated channel to send `Person` objects.  
* **Sending**: `queueChan <- p` puts order data into the queue.  
* **Receiving**: `for p := range queueChan` inside a Goroutine continuously waits for and retrieves data from the channel as long as it is not closed.

### 2. WaitGroup (`sync.WaitGroup`)
A WaitGroup acts as a counter to ensure the application does not stop before all asynchronous processes are completed.

* **`wg.Add(1)`**: Called whenever a new order is added to be processed (increments the counter).  
* **`wg.Done()`**: Called using `defer` inside the processing function to signal that one order has finished.  
* **`wg.Wait()`**: Keeps the main program running until the counter returns to zero (all orders are completed).

---

## 💻 How to Run
1. Open the terminal in VS Code.  
2. Make sure the Go code is inside the `main.go` file.  
3. Run the command:

    ```
    go run main.go
    ```

---

## 📝 Program Logic
The program processes the list of orders sequentially:

1. **Validation**: If a name is provided, the program prints the status "Waiting in queue...". If empty, it prints a warning.  
2. **Time Simulation**: Uses `time.Sleep` based on each person's `Wait` attribute.  
3. **Completion**: After all orders are processed and the channel is empty, the message "Finished" will appear as the closing output.

