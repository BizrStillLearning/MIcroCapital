import axios from 'axios'

const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api',
    headers: {
        'Content-Type': 'application/json',
    },
})

apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token')
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

apiClient.interceptors.response.use(
    (response) => {
        return response
    },
    (error) => {
        if (error.response && error.response.status === 401) {
            console.warn('Sesi kedaluwarsa. Mengeluarkan pengguna otomatis...')

            localStorage.removeItem('token')
            localStorage.removeItem('user')

            alert('Sesi Anda telah berakhir demi keamanan. Silakan masuk kembali.')

            window.location.href = '/signin'
        }

        return Promise.reject(error)
    }
)

export default apiClient