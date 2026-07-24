import { defineStore } from 'pinia'
import apiClient from '../api/axios'

export const useAdminStore = defineStore('admin', {
    state: () => ({
        analytics: {
            total_users: 0,
            total_agents: 0,
            total_funded_amount: 0
        },
        isLoading: false,
        error: null
    }),
    actions: {
        async fetchAnalytics() {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.get('/admin/analytics')
                this.analytics = response.data.data
            } catch (error) {
                console.error("Gagal mengambil data analitik:", error)
                this.error = "Gagal memuat statistik platform"
            } finally {
                this.isLoading = false
            }
        },

        async fetchAgents() {
            this.isLoading = true
            try {
                const response = await apiClient.get('/admin/agents')
                return response.data.data
            } catch (error) {
                console.error(error)
                return []
            } finally {
                this.isLoading = false
            }
        },

        async approveAgent(agentId) {
            try {
                await apiClient.post(`/admin/agents/${agentId}/approve`)
                return true
            } catch (error) {
                console.error(error)
                return false
            }
        },

        async fetchPendingLoans() {
            try {
                const response = await apiClient.get('/admin/loans/pending')
                return response.data.data
            } catch (error) {
                console.error("Gagal mengambil pinjaman:", error)
                return []
            }
        },

        async approveLoan(loanId) {
            try {
                const response = await apiClient.post(`/admin/loans/${loanId}/approve`)
                return { success: true, message: response.data.message }
            } catch (error) {
                return { success: false, error: error.response?.data?.error || "Gagal menyetujui pinjaman" }
            }
        }
    }
})