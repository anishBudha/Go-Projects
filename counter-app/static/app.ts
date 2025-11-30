 const API_URL = "http://localhost:8080/api/counter";

 const counterDisplay = document.getElementById("counter") as HTMLElement;
 const incrementBtn = document.getElementById("incrementBtn") as HTMLButtonElement;
 const decrementBtn = document.getElementById("decrementBtn") as HTMLButtonElement;
 const statusDisplay = document.getElementById("status") as HTMLElement;

 async function fetchCounter(): Promise<void> {
   try {
     statusDisplay.textContent = "loading...";
     const response = await fetch (API_URL);
     const data = await response.json();
     counterDisplay.textContent = data.count.toString();

     statusDisplay.textContent = "";
   } catch (error) {
     console.error("Error fetching counter:", error);
     statusDisplay.textContent = "Error loading counter";
   }
 }

 async function incrementCounter(): Promise<void> {
   try {
     incrementBtn.disabled = true;
     decrementBtn.disabled = true;

     const response = await fetch(`${API_URL}/increment`, {
       method: "POST", 
     });

     const data = await response.json();

     counterDisplay.textContent = data.count.toString();

     counterDisplay.style.transform = "scale(1.1)";
     setTimeout(() => {
       counterDisplay.style.transform = "scale(1)";
     }, 200);
   } catch (error) {
     console.error("Error incrementing counter:", error);
     statusDisplay.textContent = "Error updating counter increment";
   } finally {
     incrementBtn.disabled = false;
     decrementBtn.disabled = false;
   }
 }

 async function decrementCounter(): Promise<void> {
   try {
     incrementBtn.disabled = true;
     decrementBtn.disabled = true;

     const response = await fetch(`${API_URL}/decrement`, {
       method: "POST",
     });

     const data = await response.json();
 
     counterDisplay.textContent = data.count.toString();

     counterDisplay.style.transform = "scale(1.1)";
     setTimeout(() => {
       counterDisplay.style.transform = "scale(1)";
     }, 200);
   } catch (error) {
     console.error("Error decrementing counter:", error);
     statusDisplay.textContent = "Error updating counter decrement";
   } finally {
     incrementBtn.disabled = false;
     decrementBtn.disabled = false;
   }
 }

 incrementBtn.addEventListener("click", incrementCounter);
 decrementBtn.addEventListener("click", decrementCounter);

 fetchCounter();
