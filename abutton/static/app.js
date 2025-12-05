const API_URL = 'http://localhost:8080/api/message'

const textArea = document.getElementById("displayText");

const displayBtn = document.getElementById("displayBtn");

async function displayMessage() {
  try {
    const response = await fetch(API_URL);
    const data = await response.json();
    textArea.textContent = data.message;
    
  } catch (error) {
    console.error("Error fetching message:", error);
  }
}

displayBtn.addEventListener("click", displayMessage);
