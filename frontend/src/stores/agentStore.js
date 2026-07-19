import { defineStore } from 'pinia'
import apiClient from '../api/axios'

export const useAgentStore = defineStore('agent', {
    state: () => ({
        unverifiedUsers: [],
        isLoading: false,
        error: null
    }),

    actions: {
        async fetchUnverifiedUsers() {
            this.isLoading = true
            try {
                const response = await apiClient.get('/agent/unverified-users')
                this.unverifiedUsers = response.data.data
            } catch (err) {
                this.error = 'Gagal memuat daftar warga'
            } finally {
                this.isLoading = false
            }
        },

        async verifyUser(userId) {
            this.isLoading = true
            try {
                await apiClient.post(`/agent/verify/${userId}`)
                this.unverifiedUsers = this.unverifiedUsers.filter(u => u.id !== userId)
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Gagal memverifikasi warga'
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})