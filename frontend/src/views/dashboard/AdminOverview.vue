<script setup>
import { ref, onMounted, computed } from 'vue'
import { Users, ShieldCheck, Activity, BarChart4, CheckCircle, XCircle } from '@lucide/vue'
import { useAdminStore } from '../../stores/adminStore'
import { useAuthStore } from '../../stores/authStore'
import Swal from 'sweetalert2'

const adminStore = useAdminStore()
const authStore = useAuthStore()

const analytics = computed(() => adminStore.analytics)
const adminName = computed(() => authStore.user?.name || 'Administrator')

const pendingLoans = ref([])

const loadData = async () => {
  adminStore.fetchAnalytics()
  pendingLoans.value = await adminStore.fetchPendingLoans()
}

onMounted(() => {
  loadData()
})

const formatRupiah = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka || 0)
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('id-ID', { year: 'numeric', month: 'short', day: 'numeric' })
}

const handleApproveLoan = async (loan) => {
  const confirm = await Swal.fire({
    title: 'Setujui Pinjaman?',
    text: `Cairkan dana sebesar ${formatRupiah(loan.total_amount)} ke dompet ${loan.borrower?.name || 'warga ini'}?`,
    icon: 'question',
    showCancelButton: true,
    confirmButtonColor: '#10b981',
    cancelButtonColor: '#ef4444',
    confirmButtonText: 'Ya, Cairkan'
  })

  if (confirm.isConfirmed) {
    const result = await adminStore.approveLoan(loan.id)
    if (result.success) {
      Swal.fire('Berhasil!', result.message, 'success')
      loadData()
    } else {
      Swal.fire('Gagal', result.error, 'error')
    }
  }
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

    <div class="mt-8">
      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Antrean Persetujuan Pinjaman</h3>
          <span class="bg-amber-100 text-amber-700 text-xs font-bold px-3 py-1 rounded-full">
            {{ pendingLoans ? pendingLoans.length : 0 }} Menunggu
          </span>
        </div>

        <div v-if="!pendingLoans || pendingLoans.length === 0" class="flex-1 flex flex-col items-center justify-center py-10 text-muted-foreground">
          <CheckCircle :size="40" class="text-green-500 mb-3 opacity-50" />
          <p class="text-sm font-medium">Semua pinjaman telah diproses atau antrean kosong.</p>
        </div>

        <div v-else class="space-y-4">
          <div v-for="loan in pendingLoans" :key="loan.id" class="p-4 bg-muted/20 border border-border rounded-2xl flex flex-col sm:flex-row justify-between sm:items-center gap-4">
            <div>
              <div class="flex items-center gap-2 mb-1">
                <h4 class="font-bold text-foreground">{{ loan.borrower?.name || 'Warga' }}</h4>
                <span class="text-[10px] bg-primary/10 text-primary px-2 py-0.5 rounded-full font-bold">Warga</span>
              </div>
              <p class="text-sm text-foreground font-medium">{{ loan.title }}</p>
              <p class="text-xs text-muted-foreground mt-1" style="font-family: 'DM Mono', monospace;">
                {{ formatDate(loan.created_at) }} • Cicilan: {{ formatRupiah(loan.monthly_installment) }}/bln
              </p>
            </div>

            <div class="flex flex-col sm:items-end gap-3">
              <span class="font-bold text-lg text-foreground" style="font-family: 'DM Mono', monospace;">{{ formatRupiah(loan.total_amount) }}</span>
              <div class="flex gap-2 w-full sm:w-auto">
                <button @click="handleApproveLoan(loan)" class="flex-1 sm:flex-none bg-primary text-primary-foreground px-4 py-2 rounded-xl text-xs font-bold flex items-center justify-center gap-1 hover:opacity-90 transition-opacity">
                  <CheckCircle :size="14" /> Cairkan
                </button>
                <button class="bg-red-50 text-red-600 px-3 py-2 rounded-xl hover:bg-red-100 transition-colors">
                  <XCircle :size="16" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>