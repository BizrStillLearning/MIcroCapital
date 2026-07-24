import { defineStore } from 'pinia'
import apiClient from '../api/axios'
import { useAuthStore } from './authStore'

export const useLoanStore = defineStore('loan', {
    state: () => ({
        loans: [],
        isLoading: false,
        error: null
    }),
    actions: {
        async fetchLoans() {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.get('/loans')
                this.loans = response.data.data
            } catch (error) {
                console.error(error)
                this.error = "Gagal memuat daftar pinjaman"
            } finally {
                this.isLoading = false
            }
        },

        async applyLoan(payload) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/loans', payload)
                this.loans.unshift(response.data.data)
                return true
            } catch (error) {
                this.error = error.response?.data?.error || "Gagal mengajukan pinjaman"
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
                authStore.user.balance = response.data.remaining_balance

                await this.fetchLoans()
                return true
            } catch (error) {
                this.error = error.response?.data?.error || "Gagal membayar cicilan"
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})