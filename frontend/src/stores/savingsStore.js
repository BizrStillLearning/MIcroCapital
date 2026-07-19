import { defineStore } from 'pinia'
import apiClient from '../api/axios'
import { useAuthStore } from './authStore'

export const useSavingsStore = defineStore('savings', {
    state: () => ({
        isLoading: false,
        error: null
    }),

    actions: {
        async payFee(groupId) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post(`/savings/${groupId}/pay`)

                const authStore = useAuthStore()
                if (authStore.user) {
                    authStore.user.balance -= response.data.amount_deducted || 0
                }
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal membayar iuran'
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})