const { JSDOM } = require("jsdom");
const { performance } = require("perf_hooks");

const dom = new JSDOM(`<!DOCTYPE html><html><body><pre id="output"></pre></body></html>`);
const output = dom.window.document.getElementById("output");

const ITERATIONS = 1000;

// Benchmark innerHTML +=
let start = performance.now();
for (let i = 0; i < ITERATIONS; i++) {
    output.innerHTML += "Message " + i + "\n";
}
let end = performance.now();
const timeInnerHTML = end - start;

// Reset
output.innerHTML = "";

// Benchmark appendChild
start = performance.now();
for (let i = 0; i < ITERATIONS; i++) {
    output.appendChild(dom.window.document.createTextNode("Message " + i + "\n"));
}
end = performance.now();
const timeAppendChild = end - start;

console.log(`innerHTML += time: ${timeInnerHTML.toFixed(2)} ms`);
console.log(`appendChild time: ${timeAppendChild.toFixed(2)} ms`);
console.log(`Improvement: ${(timeInnerHTML / timeAppendChild).toFixed(2)}x faster`);
