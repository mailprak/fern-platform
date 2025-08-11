function formatDuration(milliseconds) {
  // Handle null, undefined, NaN, and negative values
  if (milliseconds === null || milliseconds === undefined || isNaN(milliseconds) || milliseconds < 0) return "0ms";
  
  // Ensure we're working with a number
  const ms = Number(milliseconds);
  if (isNaN(ms)) return "0ms";
  
  // Handle very small durations (less than 1ms)
  if (ms === 0) return "0ms";
  if (ms < 1) return "<1ms";
  
  if (ms < 1000) {
    return Math.round(ms) + "ms";
  }
  
  const seconds = ms / 1000;
  if (seconds < 60) {
    return seconds.toFixed(1) + "s";
  }
  
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = Math.round(seconds % 60);
  
  if (minutes < 60) {
    return minutes + "m " + remainingSeconds + "s";
  }
  
  const hours = Math.floor(minutes / 60);
  const remainingMinutes = minutes % 60;
  return hours + "h " + remainingMinutes + "m " + remainingSeconds + "s";
}

// Make sure it's available globally in multiple ways
window.formatDuration = formatDuration;
if (typeof globalThis !== 'undefined') {
  globalThis.formatDuration = formatDuration;
}
// Also make it available as a regular global
if (typeof global !== 'undefined') {
  global.formatDuration = formatDuration;
}

// Log that the function is loaded
console.log("formatDuration function loaded successfully");
console.log("Testing with 2501ms:", formatDuration(2501));