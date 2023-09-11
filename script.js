import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    vus: 100,
    iterations: 100
};


export default function () {
    const url = 'http://localhost:8080';
    const payload = JSON.stringify({ "paragraph": "Sample paragraph" }); // Replace with your JSON data

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    // Send a POST request with the JSON payload
    const response = http.post(url, payload, params);

    // Check the response status code, add more checks if needed
    if (response.status !== 200) {
        console.error(`Request failed with status code ${response.status}`);
    }

    // Sleep for a short duration (e.g., 1 second) between requests to simulate some delay
    sleep(1);
}
