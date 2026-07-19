import { defineStore } from 'pinia'
import apiClient from '../api/axios'
import { useAuthStore } from './authStore'

export const useLoanStore = defineStore('loan', {
    state: () => ({
        isLoading: false,
        error: null
    }),

    actions: {
        async applyLoan(payload) {
            this.isLoading = true
            this.error = null
            try {
                await apiClient.post('/loans', payload)
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal mengajukan pinjaman'
                return false
            } finally {
                this.isLoading = false
            }
        },

        async payInstallment(loanId) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post(`/loans/${loanId}/pay`)

                const authStore = useAuthStore()
                if (authStore.user) {
                    authStore.user.balance = response.data.remaining_balance
                    localStorage.setItem('user', JSON.stringify(authStore.user))
                }
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal membayar cicilan'
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})