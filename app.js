// Create Go object from wasm_exec.js
const go = new Go();
// Load WASM binary
WebAssembly.instantiateStreaming(fetch("main.wasm"), go?.importObject)
  ?.then((result) => {
    go?.run(result.instance);
    console?.log(addFunc(5, 5));
  })
  ?.catch(console?.error);
