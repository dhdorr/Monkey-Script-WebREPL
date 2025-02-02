// document.getElementById('createButton').addEventListener('click', async () => {
//     try {
//         const response = await fetch('/create', {
//             method: 'GET',
//         });

//         if (response.ok) {
//             // The session_id is automatically stored in the cookie
//             console.log('Session ID stored in cookie');

//             // Handle the new HTML page (if needed)
//             const newPageHtml = await response.text();
//             document.body.innerHTML = newPageHtml;
//         } else {
//             console.error('Failed to fetch:', response.statusText);
//         }
//     } catch (error) {
//         console.error('Error:', error);
//     }
// });


// Function to read the session_id from cookies
// function getSessionIdFromCookie() {
//     const cookies = document.cookie.split(';');
//     for (const cookie of cookies) {
//         const [name, value] = cookie.trim().split('=');
//         if (name === 'session_id') {
//             return value;
//         }
//     }
//     return null;
// }

// // Example: Check for session_id on page load
// window.addEventListener('load', () => {
//     const sessionId = getSessionIdFromCookie();
//     if (sessionId) {
//         console.log('Session ID from cookie:', sessionId);
//     } else {
//         console.log('No session ID found in cookies');
//     }
// });