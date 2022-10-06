export default async function compiler(url = '', data = {}) {
  // console.log((typeof data))
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data)
    });
    return response; // parses JSON response into native JavaScript objects
  }