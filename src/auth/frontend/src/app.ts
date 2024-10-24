import axios from 'axios';

document.getElementById('register-form')?.addEventListener('submit', async (event) => {
    event.preventDefault();

    const name = (document.getElementById('name') as HTMLInputElement).value;
    const email = (document.getElementById('email') as HTMLInputElement).value;
    const password = (document.getElementById('password') as HTMLInputElement).value;
    const phone = (document.getElementById('phone') as HTMLInputElement).value;

    try {
        const response = await axios.post('http://localhost:8080/register', {
            name, email, password, phone
        });
        alert('Registro exitoso');
    } catch (error) {
        alert('Error en el registro');
    }
});
