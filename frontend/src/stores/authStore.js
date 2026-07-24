import { defineStore } from 'pinia'
import apiClient from '../api/axios'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')) || null,
        token: localStorage.getItem('token') || null,
        isLoading: false,
        error: null,
    }),

    getters: {
        isAuthenticated: (state) => !!state.token,
        userRole: (state) => state.user?.role || 'member',
    },

    actions: {
        async register(name, phone, pin) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/register', { name, phone, pin })
                return { success: true, message: response.data.message }
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal melakukan pendaftaran'
                return { success: false }
            } finally {
                this.isLoading = false
            }
        },

        async login(phone, pin) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/login', { phone, pin })

                const receivedToken = response.data.token
                const receivedUser = response.data.user

                if (!receivedToken) {
                    throw new Error("Token tidak diterima dari server.")
                }

                this.token = receivedToken
                this.user = receivedUser

                localStorage.setItem('token', this.token)
                localStorage.setItem('user', JSON.stringify(this.user))

                return true
            } catch (err) {
                this.error = err.response?.data?.error || err.message || 'Gagal masuk'
                return false
            } finally {
                this.isLoading = false
            }
        },

        logout() {
            this.user = null
            this.token = null
            localStorage.removeItem('token')
            localStorage.removeItem('user')
        },

        async updatePin(currentPin, newPin) {
            this.isLoading = true
            this.error = null
            try {
                await apiClient.put('/profile/pin', { current_pin: currentPin, new_pin: newPin })
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal memperbarui PIN'
                return false
            } finally {
                this.isLoading = false
            }
        },

        async adminLogin(email, password) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/admin/login', {
                    email: email,
                    password: password
                })

                this.token = response.data.token
                this.user = response.data.user

                localStorage.setItem('token', this.token)
                localStorage.setItem('user', JSON.stringify(this.user))

                return true
            } catch (error) {
                this.error = error.response?.data?.error || "Gagal terhubung ke server."
                return false
            } finally {
                this.isLoading = false
            }
        },
    }
})