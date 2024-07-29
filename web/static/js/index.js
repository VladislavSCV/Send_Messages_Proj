document.addEventListener('DOMContentLoaded', () => {
    const pingBtn = document.getElementById('pingBtn');
    const pingResult = document.getElementById('pingResult');
    const messageForm = document.getElementById('messageForm');
    const messageResult = document.getElementById('messageResult');

    pingBtn.addEventListener('click', () => {
        fetch('http://localhost:8000/ping')
            .then(response => response.json())
            .then(data => {
                pingResult.textContent = data.message;
            })
            .catch(error => {
                pingResult.textContent = 'Error pinging server';
                console.error('Error:', error);
            });
    });

    messageForm.addEventListener('submit', (event) => {
        event.preventDefault();
        const formData = new FormData(messageForm);
        const message = formData.get('message');

        fetch('http://localhost:8000/sendMessage', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `message=${encodeURIComponent(message)}`
        })
            .then(response => response.text())
            .then(data => {
                messageResult.textContent = data;
            })
            .catch(error => {
                messageResult.textContent = 'Error sending message';
                console.error('Error:', error);
            });
    });
});
