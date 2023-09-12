import http from 'k6/http';

export const options = {
    vus: 1,
    iterations: 1
};


export default function () {

    const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXJhZ3JhcGgiOiJ0ZXh0IGlzIGEgcGFyYWdyYXBoIGhlcmUiLCJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.Z9vMaZjaEuPNSpJ1a5pzjRPaV_Gm-oNqIIqMN5XOJZI';
    const url = 'http://localhost:8080/';

    const params = {
        headers : {
            'Authorization': `Bearer ${token}`,
        }
      };

    const response = http.get(url, params);
    if (response.status !== 200) {
        console.error(`Request failed with status code ${response.status}`);
    }
}
