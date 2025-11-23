import './style.css'

interface ConvertResponse {
  from: string
  to: string
  amount: number
  converted: number
  date: string
  rate: number
}

class CurrencyConverter {
  private fromInput: HTMLInputElement
  private toInput: HTMLInputElement
  private amountInput: HTMLInputElement
  private convertButton: HTMLButtonElement
  private resultDiv: HTMLDivElement
  private errorDiv: HTMLDivElement

  constructor() {
    this.initializeElements()
    this.setupEventListeners()
  }

  private initializeElements(): void {
    const app = document.querySelector<HTMLDivElement>('#app')!
    
    app.innerHTML = `
      <div class="min-h-screen bg-white flex items-center justify-center p-4">
        <div class="w-full max-w-md border-2 border-black p-8">
          <h1 class="text-3xl font-bold text-black mb-8 text-center">Currency Converter</h1>
          
          <div class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-black mb-2">From Currency</label>
              <input 
                type="text" 
                id="from" 
                placeholder="USD" 
                class="w-full px-4 py-2 border-2 border-black focus:outline-none focus:ring-2 focus:ring-black text-black bg-white"
                maxlength="3"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-black mb-2">To Currency</label>
              <input 
                type="text" 
                id="to" 
                placeholder="EUR" 
                class="w-full px-4 py-2 border-2 border-black focus:outline-none focus:ring-2 focus:ring-black text-black bg-white"
                maxlength="3"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-black mb-2">Amount</label>
              <input 
                type="number" 
                id="amount" 
                placeholder="100" 
                step="0.01"
                min="0"
                class="w-full px-4 py-2 border-2 border-black focus:outline-none focus:ring-2 focus:ring-black text-black bg-white"
              />
            </div>
            
            <button 
              id="convert" 
              class="w-full py-3 bg-black text-white font-medium hover:bg-gray-800 transition-colors border-2 border-black"
            >
              Convert
            </button>
            
            <div id="error" class="hidden text-sm text-black bg-gray-100 border-2 border-black p-3"></div>
            
            <div id="result" class="hidden space-y-2 pt-4 border-t-2 border-black">
              <div class="text-sm text-black">
                <span class="font-medium">Result:</span>
                <span id="converted-amount" class="text-lg font-bold ml-2"></span>
              </div>
              <div class="text-xs text-black">
                <span>Rate: </span>
                <span id="rate"></span>
              </div>
              <div class="text-xs text-black">
                <span>Date: </span>
                <span id="date"></span>
              </div>
            </div>
          </div>
        </div>
      </div>
    `

    this.fromInput = document.getElementById('from') as HTMLInputElement
    this.toInput = document.getElementById('to') as HTMLInputElement
    this.amountInput = document.getElementById('amount') as HTMLInputElement
    this.convertButton = document.getElementById('convert') as HTMLButtonElement
    this.resultDiv = document.getElementById('result') as HTMLDivElement
    this.errorDiv = document.getElementById('error') as HTMLDivElement
  }

  private setupEventListeners(): void {
    this.convertButton.addEventListener('click', () => this.handleConvert())
    
    // Allow Enter key to trigger conversion
    [this.fromInput, this.toInput, this.amountInput].forEach(input => {
      input.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') {
          this.handleConvert()
        }
      })
    })

    // Convert currency codes to uppercase
    this.fromInput.addEventListener('input', (e) => {
      const target = e.target as HTMLInputElement
      target.value = target.value.toUpperCase()
    })

    this.toInput.addEventListener('input', (e) => {
      const target = e.target as HTMLInputElement
      target.value = target.value.toUpperCase()
    })
  }

  private async handleConvert(): Promise<void> {
    const from = this.fromInput.value.trim().toUpperCase()
    const to = this.toInput.value.trim().toUpperCase()
    const amount = parseFloat(this.amountInput.value)

    // Hide previous results/errors
    this.resultDiv.classList.add('hidden')
    this.errorDiv.classList.add('hidden')

    // Validation
    if (!from || !to || !amount || amount <= 0) {
      this.showError('Please fill in all fields with valid values')
      return
    }

    if (from.length !== 3 || to.length !== 3) {
      this.showError('Currency codes must be 3 characters (e.g., USD, EUR)')
      return
    }

    try {
      this.convertButton.disabled = true
      this.convertButton.textContent = 'Converting...'

      const response = await fetch(
        `/api/convert?from=${encodeURIComponent(from)}&to=${encodeURIComponent(to)}&amount=${encodeURIComponent(amount)}`
      )

      if (!response.ok) {
        const errorText = await response.text()
        throw new Error(errorText || 'Conversion failed')
      }

      const data: ConvertResponse = await response.json()
      this.showResult(data)
    } catch (error) {
      this.showError(error instanceof Error ? error.message : 'Failed to convert currency')
    } finally {
      this.convertButton.disabled = false
      this.convertButton.textContent = 'Convert'
    }
  }

  private showResult(data: ConvertResponse): void {
    const convertedAmountEl = document.getElementById('converted-amount')!
    const rateEl = document.getElementById('rate')!
    const dateEl = document.getElementById('date')!

    convertedAmountEl.textContent = `${data.converted.toFixed(2)} ${data.to}`
    rateEl.textContent = `1 ${data.from} = ${data.rate.toFixed(6)} ${data.to}`
    dateEl.textContent = data.date

    this.resultDiv.classList.remove('hidden')
  }

  private showError(message: string): void {
    this.errorDiv.textContent = message
    this.errorDiv.classList.remove('hidden')
  }
}

// Initialize the converter when DOM is ready
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', () => {
    new CurrencyConverter()
  })
} else {
  new CurrencyConverter()
}

