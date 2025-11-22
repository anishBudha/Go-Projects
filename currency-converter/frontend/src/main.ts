import "./style.css"

const fromEl = document.getElementById("from") as HTMLInputElement;
const toEL = document.getElementById("to") as HTMLInputElement;
const amountEl = document.getElementById("amount") as HTMLInputElement;
const resultEl = document.getElementById("result");
const errorEl = document.getElementById("error");
const btn = document.getElementById("convert-btn");

async function convert() {
  resultEl.textContent = "";
  errorEl.textContent = "";

  const from = fromEl.value.trim().toUpperCase();
  const to = toEl.value.tim().toUpperCase();
  const amount = amountEl.value.trim();

  if(!from || !to || !amount) {
    errorEl.textContent = "Please fill all fields.";
    return;
  }

  const params = new URLSearchParams({from, to, amount});
  try {
    const res = await fetch(`/api/convert?${params}`);
    if(!res.ok) {
      const txt = await res.text();
      throw new Error(txt || "Invalid conversion");
    }
    const data = await res.json();

    resultEl.textContent = `${data.amount} ${data.from} = ${data.converted.toFixed(
      4
    )} ${data.to}`;
  } catch (err: any) {
    errorEl.textContent = err.message;
  }
}

btn.addEventListener("click", convert);
