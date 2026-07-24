<script setup>
import { ref, onMounted } from 'vue'
import { Wallet, History, Users } from '@lucide/vue'
import { useAuthStore } from '../../stores/authStore'
import { useSavingsStore } from '../../stores/savingsStore'
import apiClient from '../../api/axios'
import Swal from 'sweetalert2'

const authStore = useAuthStore()
const savingsStore = useSavingsStore()

const formatRupiah = (angka) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(angka || 0)
}

const activeGroup = ref({
  id: 1,
  name: 'Arisan Petani Maju',
  monthly_fee: 50000,
  total_pool: 1500000
})

const fetchProfile = async () => {
  try {
    const response = await apiClient.get('/profile')
    if (response.data.user) {
      authStore.user.balance = response.data.user.balance
    }
  } catch (error) {
    console.error("Gagal sinkronisasi saldo:", error)
  }
}

const paySavings = async (groupId) => {
  const result = await Swal.fire({
    title: 'Bayar Iuran Kas?',
    text: `Saldo Anda akan dipotong sebesar ${formatRupiah(activeGroup.value.monthly_fee)}.`,
    icon: 'question',
    showCancelButton: true,
    confirmButtonColor: '#10b981',
    cancelButtonColor: '#ef4444',
    confirmButtonText: 'Ya, Bayar'
  })

  if (result.isConfirmed) {
    const success = await savingsStore.payFee(groupId, activeGroup.value.monthly_fee)

    if (success) {
      Swal.fire({
        icon: 'success',
        title: 'Berhasil!',
        text: 'Iuran tabungan kelompok berhasil dibayarkan.',
        timer: 2000,
        showConfirmButton: false
      })

      authStore.user.balance -= activeGroup.value.monthly_fee
      activeGroup.value.total_pool += activeGroup.value.monthly_fee

    } else {
      Swal.fire({
        icon: 'error',
        title: 'Pembayaran Gagal',
        text: savingsStore.error || 'Terjadi kesalahan sistem'
      })
    }
  }
}

onMounted(() => {
  fetchProfile()
})
</script>

<template>
  <div class="space-y-6 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">Dompet Kas & Tabungan</h1>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

      <div class="bg-primary text-primary-foreground rounded-3xl p-6 shadow-md relative overflow-hidden">
        <Wallet class="absolute -right-5 -bottom-5 text-white/10" :size="120" />
        <p class="text-sm font-medium opacity-80 mb-1" style="font-family: 'DM Mono', monospace;">SALDO AKTIF</p>
        <h2 class="text-4xl font-bold tracking-tight mb-8" style="font-family: 'Fraunces', serif;">
          {{ formatRupiah(authStore.user?.balance) }}
        </h2>
        <div class="flex gap-3 relative z-10">
          <button class="bg-white/20 px-4 py-2 rounded-lg text-sm font-bold hover:bg-white/30 transition-colors">
            Riwayat Transaksi
          </button>
        </div>
      </div>

      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col justify-between hover:border-primary/50 transition-colors">
        <div>
          <div class="flex items-center gap-2 text-accent mb-4">
            <Users :size="20" />
            <h3 class="font-bold text-foreground">{{ activeGroup.name }}</h3>
          </div>
          <p class="text-sm text-muted-foreground mb-1">Dana Terkumpul Komunitas:</p>
          <p class="text-2xl font-bold text-foreground" style="font-family: 'DM Mono', monospace;">
            {{ formatRupiah(activeGroup.total_pool) }}
          </p>
        </div>

        <div class="mt-6 pt-4 border-t border-border flex justify-between items-center">
          <div>
            <p class="text-xs text-muted-foreground">Iuran Bulan Ini</p>
            <p class="font-bold text-sm text-foreground">{{ formatRupiah(activeGroup.monthly_fee) }}</p>
          </div>
          <button
              @click="paySavings(activeGroup.id)"
              :disabled="savingsStore.isLoading"
              class="bg-accent text-accent-foreground px-5 py-2.5 rounded-xl text-sm font-bold hover:opacity-90 disabled:opacity-50 transition-all shadow-sm"
          >
            {{ savingsStore.isLoading ? 'Memproses...' : 'Bayar Iuran' }}
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

