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
            try {
                const response = await apiClient.get('/campaigns')
                this.campaigns = response.data.data
            } catch (err) {
                this.error = 'Gagal memuat daftar kampanye'
            } finally {
                this.isLoading = false
            }
        },

        async createCampaign(campaignData) {
            this.isLoading = true
            try {
                await apiClient.post('/campaigns', campaignData)
                await this.fetchCampaigns() 
                return true
            } catch (err) {
                this.error = 'Gagal membuat kampanye'
                return false
            } finally {
                this.isLoading = false
            }
        }
    }
})