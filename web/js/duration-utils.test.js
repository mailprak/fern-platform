// Simple test file for formatDuration function
// Run this in browser console to verify functionality

function testFormatDuration() {
    console.log("Testing formatDuration function...");
    
    // Test cases based on acceptance criteria
    const testCases = [
        { input: 0, expected: "0ms" },
        { input: 500, expected: "500ms" },
        { input: 999, expected: "999ms" },
        { input: 1000, expected: "1.0s" },
        { input: 1500, expected: "1.5s" },
        { input: 12300, expected: "12.3s" },
        { input: 59999, expected: "59.9s" },
        { input: 60000, expected: "1m 0s" },
        { input: 65000, expected: "1m 5s" },
        { input: 125000, expected: "2m 5s" },
        { input: 3600000, expected: "1h 0m 0s" },
        { input: 3725000, expected: "1h 2m 5s" },
        { input: null, expected: "0ms" },
        { input: undefined, expected: "0ms" },
        { input: -100, expected: "0ms" }
    ];
    
    let allPassed = true;
    
    testCases.forEach((testCase, index) => {
        const result = formatDuration(testCase.input);
        const passed = result === testCase.expected;
        
        if (!passed) {
            console.error(`Test ${index + 1} FAILED:`, {
                input: testCase.input,
                expected: testCase.expected,
                actual: result
            });
            allPassed = false;
        } else {
            console.log(`Test ${index + 1} PASSED:`, {
                input: testCase.input,
                result: result
            });
        }
    });
    
    if (allPassed) {
        console.log("✅ All tests passed!");
    } else {
        console.log("❌ Some tests failed!");
    }
    
    return allPassed;
}

// Auto-run tests if this file is loaded
if (typeof formatDuration === 'function') {
    testFormatDuration();
} else {
    console.error("formatDuration function not found. Make sure duration-utils.js is loaded first.");
}
