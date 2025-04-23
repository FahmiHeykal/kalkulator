const display = document.getElementById('display');

function appendNumber(num) {
  if (display.innerText === '0') {
    display.innerText = num;
  } else {
    display.innerText += num;
  }
}

function appendOperator(op) {
  const lastChar = display.innerText.slice(-1);
  if ('+-*/'.includes(lastChar)) return; // prevent duplicate ops
  display.innerText += op;
}

function clearDisplay() {
  display.innerText = '0';
}

function deleteLast() {
  const text = display.innerText;
  display.innerText = text.length > 1 ? text.slice(0, -1) : '0';
}

function calculate() {
    try {
      let expression = display.innerText
        .replace(/×/g, '*')
        .replace(/÷/g, '/');
  
      if (/[\+\-\*\/]$/.test(expression) || expression === '') {
        display.innerText = 'Error';
        return;
      }
  
      if (/[\+\-\*\/]{2,}/.test(expression)) {
        display.innerText = 'Error';
        return;
      }
  
      const result = eval(expression);
      display.innerText = result;
    } catch {
      display.innerText = 'Error';
    }
  
}
