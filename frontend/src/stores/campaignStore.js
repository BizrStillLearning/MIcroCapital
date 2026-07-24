import { defineStore } from 'pinia'
import apiClient from '../api/axios'

export const useCampaignStore = defineStore('campaign', {
    state: () => ({
        campaigns: [],
        isLoading: false,
        error: null
    }),
    actions: {
        async fetchCampaigns() {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.get('/campaigns')
                this.campaigns = response.data.data
            } catch (error) {
                this.error = 'Gagal mengambil data kampanye'
                console.error(error)
            } finally {
                this.isLoading = false
            }
        },

        async createCampaign(payload) {
            this.isLoading = true
            this.error = null
            try {
                const response = await apiClient.post('/campaigns', payload)

                this.campaigns.unshift(response.data.data)

                return true
            } catch (error) {
                this.error = error.response?.data?.error || 'Gagal membuat kampanye'
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})