<script setup>
import { ref, onMounted } from 'vue'
import { Users, CheckCircle, Search, XCircle } from '@lucide/vue'
import apiClient from '../../api/axios'
import Swal from 'sweetalert2'

const unverifiedUsers = ref([])
const isProcessing = ref(false)
const searchQuery = ref('')

const fetchUnverifiedUsers = async () => {
  try {
    const response = await apiClient.get('/agent/unverified-users')
    unverifiedUsers.value = response.data.data
  } catch (error) {
    console.error("Gagal mengambil data pendaftar:", error)
    Swal.fire('Error', 'Gagal memuat daftar warga', 'error')
  }
}

const verifyUser = async (user) => {
  const confirm = await Swal.fire({
    title: 'Verifikasi Warga?',
    text: `Apakah Anda sudah memeriksa fisik KTP/Identitas milik ${user.name}?`,
    icon: 'question',
    showCancelButton: true,
    confirmButtonColor: '#10b981',
    cancelButtonColor: '#64748b',
    confirmButtonText: 'Ya, Verifikasi'
  })

  if (confirm.isConfirmed) {
    isProcessing.value = true
    try {
      const response = await apiClient.post(`/agent/verify/${user.id}`)
      Swal.fire('Berhasil!', response.data.message, 'success')

      unverifiedUsers.value = unverifiedUsers.value.filter(u => u.id !== user.id)
    } catch (error) {
      Swal.fire('Gagal', error.response?.data?.error || 'Gagal memverifikasi', 'error')
    } finally {
      isProcessing.value = false
    }
  }
}

onMounted(() => {
  fetchUnverifiedUsers()
})
</script>

<template>
  <div class="space-y-6 max-w-4xl mx-auto">
    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Antrean Verifikasi (KYC)
      </h1>
      <p class="text-sm text-muted-foreground mt-1">
        Lakukan pencocokan identitas fisik (KTP) sebelum menyetujui akun warga.
      </p>
    </div>

    <div class="relative w-full max-w-md">
      <Search class="absolute left-4 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
      <input
          v-model="searchQuery"
          type="text"
          placeholder="Cari nama atau nomor telepon..."
          class="w-full bg-card border border-border rounded-xl py-3 pl-11 pr-4 text-sm focus:outline-none focus:border-primary transition-colors"
      />
    </div>

    <div class="bg-card border border-border rounded-3xl p-6 shadow-sm">
      <div v-if="unverifiedUsers.length === 0" class="text-center py-10 flex flex-col items-center">
        <CheckCircle :size="48" class="text-green-500 mb-4 opacity-50" />
        <h3 class="font-bold text-foreground">Semua Bersih!</h3>
        <p class="text-sm text-muted-foreground">Tidak ada warga yang menunggu verifikasi saat ini.</p>
      </div>

      <div v-else class="space-y-4">
        <div
            v-for="user in unverifiedUsers"
            :key="user.id"
            class="flex flex-col sm:flex-row justify-between items-start sm:items-center p-4 border border-border rounded-2xl bg-muted/20"
        >
          <div class="flex gap-4 items-center mb-4 sm:mb-0">
            <div class="w-12 h-12 bg-accent/10 text-accent rounded-full flex items-center justify-center">
              <Users :size="24" />
            </div>
            <div>
              <h4 class="font-bold text-foreground">{{ user.name }}</h4>
              <p class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">{{ user.phone }}</p>
              <p class="text-xs text-muted-foreground mt-1">Mendaftar: {{ new Date(user.created_at).toLocaleDateString('id-ID') }}</p>
            </div>
          </div>

          <div class="w-full sm:w-auto flex gap-2">
            <button
                @click="verifyUser(user)"
                :disabled="isProcessing"
                class="flex-1 sm:flex-none bg-primary text-primary-foreground px-5 py-2.5 rounded-xl text-sm font-bold flex items-center justify-center gap-2 hover:opacity-90 transition-opacity"
            >
              <CheckCircle :size="16" /> Setujui
            </button>
            <button class="bg-muted text-muted-foreground px-3 py-2.5 rounded-xl hover:bg-muted/80 transition-colors">
              <XCircle :size="18" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>