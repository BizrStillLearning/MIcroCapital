<script setup>
import { onMounted, computed } from 'vue'
import { Users, ShieldCheck, Activity, BarChart4 } from '@lucide/vue'
import { useAdminStore } from '../../stores/adminStore'
import { useAuthStore } from '../../stores/authStore'

const adminStore = useAdminStore()
const authStore = useAuthStore()

const analytics = computed(() => adminStore.analytics)
const adminName = computed(() => authStore.user?.name || 'Administrator')

onMounted(() => {
  adminStore.fetchAnalytics()
})

const formatRupiah = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka || 0)
}
</script>

<template>
  <div class="space-y-6 max-w-6xl mx-auto">
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">Pusat Kendali Sistem</h1>
        <p class="text-sm text-muted-foreground mt-1">Selamat bertugas, {{ adminName }}. Berikut adalah metrik likuiditas platform saat ini.</p>
      </div>
      <button class="bg-primary text-primary-foreground px-4 py-2 rounded-xl text-sm font-bold flex items-center gap-2 hover:opacity-90">
        <BarChart4 :size="16" /> Unduh Laporan
      </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mt-6">

      <div class="bg-primary text-primary-foreground rounded-3xl p-6 shadow-md relative overflow-hidden lg:col-span-2 flex flex-col justify-center">
        <Activity class="absolute -right-5 -bottom-5 text-white/10" :size="120" />
        <p class="text-sm font-medium opacity-80 mb-2" style="font-family: 'DM Mono', monospace;">TOTAL LIKUIDITAS TERSALURKAN</p>
        <h2 class="text-4xl font-bold tracking-tight relative z-10" style="font-family: 'Fraunces', serif;">
          {{ formatRupiah(analytics.total_funded_amount) }}
        </h2>
      </div>

      <!-- Total Anggota Komunitas -->
      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col justify-center">
        <div class="w-10 h-10 bg-blue-50 text-blue-600 rounded-full flex items-center justify-center mb-4">
          <Users :size="20" />
        </div>
        <p class="text-xs text-muted-foreground font-bold mb-1" style="font-family: 'DM Mono', monospace;">ANGGOTA AKTIF</p>
        <h3 class="text-3xl font-bold text-foreground">{{ analytics.total_users }}</h3>
      </div>

      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col justify-center">
        <div class="w-10 h-10 bg-green-50 text-green-600 rounded-full flex items-center justify-center mb-4">
          <ShieldCheck :size="20" />
        </div>
        <p class="text-xs text-muted-foreground font-bold mb-1" style="font-family: 'DM Mono', monospace;">AGEN LOKAL</p>
        <h3 class="text-3xl font-bold text-foreground">{{ analytics.total_agents }}</h3>
      </div>
    </div>
  </div>
</template>