// ==============================================================================
// IMPORTS
// ==============================================================================
// Import the CSS file so styles are applied to our HTML
import './style.css'

// ==============================================================================
// CONSTANTS
// ==============================================================================
// Constants are values that don't change during program execution
// We use UPPERCASE names for constants to show they're fixed values

// API_BASE_URL is the address of our Go backend server
// Reference: Think of this like a street address - it tells the browser
// where to find our backend API
const API_BASE_URL = 'http://localhost:8080'

// ==============================================================================
// TYPE DEFINITIONS
// ==============================================================================
// TypeScript allows us to define the "shape" of our data
// This helps catch errors before running the code

// TimerState describes what data we expect from the backend
// Reference: This is like a contract - it says "when we get timer data,
// it must have these fields with these types"
interface TimerState {
  isRunning: boolean     // true or false - is the timer counting?
  milliseconds: number   // a number - how many milliseconds have elapsed?
}

// ==============================================================================
// DOM ELEMENTS
// ==============================================================================
// The DOM (Document Object Model) is the tree structure of HTML elements
// We need to "grab" elements from the HTML so we can modify them

// Get the timer display element
// querySelector finds the first element matching the CSS selector '#timer-display'
// The <HTMLDivElement> part tells TypeScript what type of element this is
// The ! at the end says "I promise this element exists"
//
// Reference: document.querySelector is like finding a specific book in a library
// using its call number (#timer-display is like a call number)
const timerDisplay = document.querySelector<HTMLDivElement>('#timer-display')!

// Get the button elements
// These will be the buttons users click to control the timer
const startButton = document.querySelector<HTMLButtonElement>('#start-btn')!
const stopButton = document.querySelector<HTMLButtonElement>('#stop-btn')!
const resetButton = document.querySelector<HTMLButtonElement>('#reset-btn')!

// ==============================================================================
// GLOBAL STATE
// ==============================================================================
// Variables that track the current state of our application

// updateInterval stores the ID of the interval timer
// An interval is a repeated action (like checking the timer every 50ms)
// We need to store the ID so we can stop it later
// null means "no interval is currently running"
//
// Reference: Think of this like a subscription ID - you need it to cancel the subscription
let updateInterval: number | null = null

// ==============================================================================
// UTILITY FUNCTIONS
// ==============================================================================

/**
 * formatTime converts milliseconds to a readable time format
 * 
 * Format: MM:SS.mmm
 * - MM = minutes (2 digits)
 * - SS = seconds (2 digits)
 * - mmm = milliseconds (3 digits)
 * 
 * Example: 125000 milliseconds becomes "02:05.000" (2 minutes, 5 seconds)
 * 
 * Reference: We break down time like converting money:
 * - 1 second = 1000 milliseconds (like 1 dollar = 100 cents)
 * - 1 minute = 60 seconds
 * 
 * @param milliseconds - The total time in milliseconds
 * @returns A formatted string like "02:05.000"
 */
function formatTime(milliseconds: number): string {
  // Step 1: Convert milliseconds to total seconds
  // Math.floor rounds down to the nearest whole number
  // Example: 125000ms ÷ 1000 = 125 seconds
  const totalSeconds = Math.floor(milliseconds / 1000)
  
  // Step 2: Calculate minutes
  // Divide total seconds by 60 and round down
  // Example: 125 seconds ÷ 60 = 2.08... → 2 minutes
  const minutes = Math.floor(totalSeconds / 60)
  
  // Step 3: Calculate remaining seconds
  // The % operator gives us the remainder after division
  // Example: 125 % 60 = 5 (2 complete minutes, 5 seconds left over)
  //
  // Reference: % is the "modulo" operator
  // Think of it like: "125 seconds = 2 minutes with 5 seconds left over"
  const seconds = totalSeconds % 60
  
  // Step 4: Get the milliseconds part (the decimal portion)
  // Example: 125456ms → we want the 456 part
  // 125456 % 1000 = 456
  const ms = Math.floor(milliseconds % 1000)
  
  // Step 5: Format with leading zeros
  // padStart adds zeros to the left if the number is too short
  // Example: "5" becomes "05", "5" becomes "005" for 3 digits
  //
  // Reference: String() converts a number to text so we can use padStart
  // padStart(2, '0') means "make it 2 characters, add '0' on the left if needed"
  return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}.${String(ms).padStart(3, '0')}`
}

// ==============================================================================
// API FUNCTIONS
// ==============================================================================
// These functions communicate with our Go backend server

/**
 * fetchTimerState gets the current timer state from the backend
 * 
 * This function is "async" which means it does work that takes time
 * (like sending a request over the network)
 * 
 * Reference: async/await is like ordering at a restaurant:
 * - You place an order (send request)
 * - You wait for the food (await response)
 * - You receive the food (get data back)
 * Meanwhile, other things can happen (like talking to friends)
 * 
 * @returns A Promise that resolves to TimerState
 */
async function fetchTimerState(): Promise<TimerState> {
  // fetch() sends an HTTP request to the specified URL
  // By default, fetch() sends a GET request
  //
  // await means "pause here until we get a response"
  // The browser doesn't freeze though - other code can still run
  const response = await fetch(`${API_BASE_URL}/api/timer`)
  
  // Check if the request was successful
  // response.ok is true if status code is 200-299
  if (!response.ok) {
    // throw creates an error that stops execution
    // This will be caught by try/catch blocks that use this function
    throw new Error('Failed to fetch timer state')
  }
  
  // response.json() converts the JSON text response to a JavaScript object
  // await means "wait for the conversion to finish"
  //
  // Reference: JSON is text format like: {"isRunning": true, "milliseconds": 5000}
  // response.json() converts this text into a real object we can use
  return await response.json()
}

/**
 * sendCommand sends a command to the backend (start, stop, or reset)
 * 
 * @param endpoint - The API endpoint to call (e.g., 'start', 'stop', 'reset')
 */
async function sendCommand(endpoint: string): Promise<void> {
  // Send a POST request to the specified endpoint
  // POST is used when we want to change something on the server
  // (GET is for reading, POST is for writing/changing)
  //
  // Reference: Think of HTTP methods like actions:
  // - GET = "show me the data" (reading)
  // - POST = "do this action" (writing/changing)
  const response = await fetch(`${API_BASE_URL}/api/timer/${endpoint}`, {
    method: 'POST', // Specify this is a POST request
  })
  
  if (!response.ok) {
    throw new Error(`Failed to ${endpoint} timer`)
  }
}

// ==============================================================================
// UI UPDATE FUNCTIONS
// ==============================================================================

/**
 * updateDisplay refreshes the timer display and button states
 * 
 * This function:
 * 1. Gets current timer state from backend
 * 2. Updates the display with formatted time
 * 3. Enables/disables buttons based on timer state
 */
async function updateDisplay(): Promise<void> {
  try {
    // Get current state from backend
    const state = await fetchTimerState()
    
    // Update the timer display
    // .textContent changes the text inside the HTML element
    //
    // Reference: timerDisplay.textContent is like changing the text in a text box
    // If the HTML is <div id="timer-display">00:00.000</div>
    // This changes it to <div id="timer-display">02:05.123</div>
    timerDisplay.textContent = formatTime(state.milliseconds)
    
    // Update button states based on whether timer is running
    // When running: Start button is disabled, Stop button is enabled
    // When stopped: Start button is enabled, Stop button is disabled
    //
    // Reference: .disabled is a property that controls if a button can be clicked
    // disabled = true → button is grayed out and can't be clicked
    // disabled = false → button is active and can be clicked
    if (state.isRunning) {
      startButton.disabled = true  // Can't start if already running
      stopButton.disabled = false  // Can stop if running
    } else {
      startButton.disabled = false // Can start if not running
      stopButton.disabled = true   // Can't stop if not running
    }
  } catch (error) {
    // If something goes wrong, log the error to the console
    // console.error prints error messages (helpful for debugging)
    console.error('Error updating display:', error)
    timerDisplay.textContent = 'ERROR'
  }
}

// ==============================================================================
// EVENT HANDLERS
// ==============================================================================
// These functions are called when the user interacts with the UI

/**
 * handleStart runs when the user clicks the Start button
 */
async function handleStart(): Promise<void> {
  try {
    // Send start command to backend
    await sendCommand('start')
    
    // Start updating the display repeatedly
    // setInterval runs a function repeatedly at a specified interval
    // Here we update every 50 milliseconds (50ms = 0.05 seconds)
    //
    // Reference: setInterval is like setting an alarm that goes off every 50ms
    // Each time it "rings", it calls updateDisplay()
    // We store the ID so we can cancel it later with clearInterval()
    if (updateInterval === null) {
      updateInterval = window.setInterval(updateDisplay, 50)
    }
    
    // Update display immediately (don't wait for first interval)
    await updateDisplay()
  } catch (error) {
    console.error('Error starting timer:', error)
  }
}

/**
 * handleStop runs when the user clicks the Stop button
 */
async function handleStop(): Promise<void> {
  try {
    // Send stop command to backend
    await sendCommand('stop')
    
    // Stop the interval that's updating the display
    // clearInterval cancels the repeated updates we started with setInterval
    //
    // Reference: This is like canceling a subscription
    // We use the ID we saved earlier to tell it which interval to stop
    if (updateInterval !== null) {
      clearInterval(updateInterval)
      updateInterval = null // Set back to null to show no interval is running
    }
    
    // Update display one final time to show the stopped time
    await updateDisplay()
  } catch (error) {
    console.error('Error stopping timer:', error)
  }
}

/**
 * handleReset runs when the user clicks the Reset button
 */
async function handleReset(): Promise<void> {
  try {
    // Send reset command to backend
    await sendCommand('reset')
    
    // Stop the update interval if it's running
    if (updateInterval !== null) {
      clearInterval(updateInterval)
      updateInterval = null
    }
    
    // Update display to show 00:00.000
    await updateDisplay()
  } catch (error) {
    console.error('Error resetting timer:', error)
  }
}

// ==============================================================================
// EVENT LISTENERS
// ==============================================================================
// Wire up the buttons to their handler functions

// addEventListener connects a button click to a function
// Syntax: element.addEventListener('event-name', function-to-call)
//
// Reference: Think of this like programming a remote control button
// "When this button is pressed, do this action"

startButton.addEventListener('click', handleStart)
stopButton.addEventListener('click', handleStop)
resetButton.addEventListener('click', handleReset)

// ==============================================================================
// INITIALIZATION
// ==============================================================================
// Code that runs when the page first loads

// Get the initial state and display it
// This shows 00:00.000 when the page first loads
updateDisplay()

// Log a message to show the app is ready
// console.log prints messages to the browser's developer console
// (Open with F12 or right-click → Inspect → Console tab)
console.log('✅ Stopwatch app initialized')
