const { JSDOM } = require("jsdom");
const { performance } = require("perf_hooks");

const dom = new JSDOM(`<!DOCTYPE html><html><body><pre id="output"></pre></body></html>`);
const output = dom.window.document.getElementById("output");

const ITERATIONS = 1000;

// Benchmark textContent +=
let start = performance.now();
for (let i = 0; i < ITERATIONS; i++) {
    output.textContent += "Message " + i + "\n";
}
let end = performance.now();
const timeTextContent = end - start;

// Reset
output.textContent = "";

// Benchmark appendChild
start = performance.now();
for (let i = 0; i < ITERATIONS; i++) {
    output.appendChild(dom.window.document.createTextNode("Message " + i + "\n"));
}
end = performance.now();
const timeAppendChild = end - start;

console.log(`textContent += time: ${timeTextContent.toFixed(2)} ms`);
console.log(`appendChild time: ${timeAppendChild.toFixed(2)} ms`);
console.log(`Improvement: ${(timeTextContent / timeAppendChild).toFixed(2)}x faster`);
