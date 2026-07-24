import { defineStore } from 'pinia'
import apiClient from '../api/axios'
import { useAuthStore } from './authStore'

export const useTransactionStore = defineStore('transaction', {
    state: () => ({
        isLoading: false,
        error: null
    }),

    actions: {
        async topUp(phone, amount) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/topup', { phone, amount })

                const authStore = useAuthStore()
                if (authStore.user) {
                    authStore.user.balance = response.data.new_balance
                    localStorage.setItem('user', JSON.stringify(authStore.user))
                }
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal melakukan isi saldo'
                return false
            } finally {
                this.isLoading = false
            }
        },

        async fundCampaign(campaignId, amount) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/fund', {
                    campaign_id: campaignId,
                    amount: amount
                })

                const authStore = useAuthStore()
                authStore.user.balance -= amount

                return true
            } catch (error) {
                this.error = error.response?.data?.error || 'Saldo tidak cukup atau transaksi gagal'
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})