import { defineStore } from 'pinia'
import apiClient from '../api/axios'

export const useAdminStore = defineStore('admin', {
    state: () => ({
        analytics: {
            total_users: 0,
            total_agents: 0,
            total_campaigns: 0,
            total_funded_amount: 0
        },
        isLoading: false,
        error: null
    }),

    actions: {
        async fetchAnalytics() {
            this.isLoading = true
            try {
                const response = await apiClient.get('/admin/analytics')
                this.analytics = response.data.data
            } catch (err) {
                this.error = 'Gagal memuat analitik platform'
            } finally {
                this.isLoading = false
            }
        }
    }
})