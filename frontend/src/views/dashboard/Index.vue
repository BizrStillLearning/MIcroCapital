<script setup>
import { ref, computed, onMounted } from 'vue'
import {
  ArrowUpRight,
  ArrowDownLeft,
  Wallet,
  Sprout,
  History,
  ChevronRight,
  Users
} from '@lucide/vue'
import { useAuthStore } from '../../stores/authStore'
import apiClient from '../../api/axios'
import Swal from 'sweetalert2'

const authStore = useAuthStore()
const user = computed(() => authStore.user || {})

const formatRupiah = (angka) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(angka || 0)
}

const recentTransactions = ref([])
const isProcessing = ref(false)

const isTopUpModalOpen = ref(false)
const topUpAmount = ref('')

const isFundModalOpen = ref(false)
const fundAmount = ref('')

const activeCampaign = ref({
  id: 1,
  title: "Perbaikan Atap Warung Bu Siti",
  raised: 850000,
  target: 1000000,
  progress: 85
})

const activeGroup = ref({
  id: 1,
  name: 'Arisan Petani Maju',
  monthly_fee: 50000
})

const fetchTransactions = async () => {
  try {
    const response = await apiClient.get('/transactions/history')
    recentTransactions.value = response.data.data
  } catch (error) {
    console.error('Gagal mengambil riwayat transaksi:', error)
  }
}

const fetchProfile = async () => {
  try {
    const response = await apiClient.get('/profile')

    const userData = response.data.user

    if (userData) {
      authStore.user.balance = userData.balance
      authStore.user.name = userData.name
      localStorage.setItem('user', JSON.stringify(authStore.user))
    }
  } catch (error) {
    console.error("Gagal mengambil profil terbaru:", error)
  }
}

const handleTopUp = async () => {
  if (!topUpAmount.value || topUpAmount.value <= 0) {
    Swal.fire({ icon: 'warning', title: 'Oops...', text: 'Masukkan nominal yang valid!' })
    return
  }

  isProcessing.value = true
  try {
    const response = await apiClient.post('/topup', {
      amount: Number(topUpAmount.value)
    })

    Swal.fire({ icon: 'success', title: 'Berhasil!', text: response.data.message || 'Saldo berhasil ditambahkan.', timer: 2000, showConfirmButton: false })

    isTopUpModalOpen.value = false
    topUpAmount.value = ''

    fetchTransactions()

    if (response.data.new_balance !== undefined) {
      authStore.user.balance = response.data.new_balance
    } else {
      authStore.user.balance += Number(topUpAmount.value)
    }

  } catch (error) {
    const errorMsg = error.response?.data?.error || "Gagal memproses Top Up"
    Swal.fire({ icon: 'error', title: 'Gagal', text: errorMsg })
  } finally {
    isProcessing.value = false
  }
}

const handleFund = async () => {
  if (!fundAmount.value || fundAmount.value <= 0) {
    Swal.fire({ icon: 'warning', title: 'Oops...', text: 'Masukkan nominal yang valid!' })
    return
  }

  if (user.value.balance < Number(fundAmount.value)) {
    Swal.fire({ icon: 'error', title: 'Saldo Tidak Cukup', text: 'Silakan isi saldo Anda terlebih dahulu.' })
    return
  }

  isProcessing.value = true
  try {
    const response = await apiClient.post('/fund', {
      campaign_id: activeCampaign.value.id,
      amount: Number(fundAmount.value)
    })

    Swal.fire({ icon: 'success', title: 'Terima Kasih!', text: 'Donasi Anda telah disalurkan.', timer: 2000, showConfirmButton: false })

    isFundModalOpen.value = false
    fundAmount.value = ''

    fetchTransactions()
    authStore.user.balance -= Number(fundAmount.value)

    activeCampaign.value.raised += Number(fundAmount.value)
    activeCampaign.value.progress = (activeCampaign.value.raised / activeCampaign.value.target) * 100

  } catch (error) {
    const errorMsg = error.response?.data?.error || "Gagal menyalurkan dana"
    Swal.fire({ icon: 'error', title: 'Transaksi Gagal', text: errorMsg })
  } finally {
    isProcessing.value = false
  }
}

const handlePayKas = async () => {
  const result = await Swal.fire({
    title: 'Bayar Iuran Kas?',
    text: `Saldo Anda akan dipotong sebesar ${formatRupiah(activeGroup.value.monthly_fee)} untuk ${activeGroup.value.name}.`,
    icon: 'question',
    showCancelButton: true,
    confirmButtonColor: '#10b981',
    cancelButtonColor: '#ef4444',
    confirmButtonText: 'Ya, Bayar Sekarang',
    cancelButtonText: 'Batal'
  })

  if (result.isConfirmed) {
    isProcessing.value = true
    try {
      const response = await apiClient.post('/pay-kas')

      Swal.fire({ icon: 'success', title: 'Berhasil!', text: response.data.message, timer: 2000, showConfirmButton: false })

      authStore.user.balance = response.data.new_balance
      fetchTransactions()

    } catch (error) {
      const errorMsg = error.response?.data?.error || "Gagal membayar Kas"
      Swal.fire({ icon: 'error', title: 'Gagal', text: errorMsg })
    } finally {
      isProcessing.value = false
    }
  }
}

onMounted(() => {
  fetchTransactions()
  fetchProfile()
})
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">
    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Selamat datang, {{ user.name }}!
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Berikut adalah ringkasan aktivitas keuangan Anda hari ini.
      </p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-primary rounded-3xl p-6 sm:p-8 text-primary-foreground relative overflow-hidden shadow-xl shadow-primary/10">
          <div class="absolute -right-10 -top-10 w-40 h-40 bg-white/10 rounded-full blur-2xl"></div>
          <div class="absolute right-20 -bottom-10 w-32 h-32 bg-accent/20 rounded-full blur-xl"></div>

          <div class="relative z-10">
            <div class="flex items-center gap-2 mb-4 opacity-90">
              <Wallet :size="20" />
              <span class="text-sm font-semibold tracking-wider" style="font-family: 'DM Mono', monospace;">SALDO TERSEDIA</span>
            </div>
            <h2 class="text-4xl sm:text-5xl font-bold mb-8" style="font-family: 'Fraunces', serif;">
              {{ formatRupiah(user.balance) }}
            </h2>

            <div class="flex flex-wrap gap-3">
              <button @click="isTopUpModalOpen = true" class="flex items-center gap-2 bg-white text-primary px-5 py-2.5 rounded-xl font-bold text-sm hover:bg-white/90 transition-colors shadow-sm">
                <ArrowDownLeft :size="18" />
                Isi Saldo
              </button>
              <button @click="isFundModalOpen = true" class="flex items-center gap-2 bg-primary-foreground/15 text-white border border-white/20 px-5 py-2.5 rounded-xl font-bold text-sm hover:bg-primary-foreground/25 transition-colors">
                <ArrowUpRight :size="18" />
                Kirim Dana
              </button>
            </div>
          </div>
        </div>

        <div class="bg-card border border-border rounded-3xl p-6">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Aktivitas Terakhir</h3>
          </div>

          <div class="space-y-4">
            <div v-if="recentTransactions.length === 0" class="text-center text-sm text-muted-foreground py-4">
              Belum ada riwayat transaksi.
            </div>

            <div
                v-for="trx in recentTransactions"
                :key="trx.id"
                class="flex items-center justify-between p-3 hover:bg-muted/50 rounded-2xl transition-colors cursor-pointer"
            >
              <div class="flex items-center gap-4">
                <div
                    class="w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0"
                    :class="trx.type === 'topup' ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'"
                >
                  <ArrowDownLeft v-if="trx.type === 'topup'" :size="18" />
                  <ArrowUpRight v-else :size="18" />
                </div>
                <div>
                  <p class="text-sm font-bold text-foreground capitalize" style="font-family: 'DM Sans', sans-serif;">
                    {{ trx.type === 'fund_campaign' ? 'Kirim Dana' : (trx.type === 'pay_dues' ? 'Bayar Kas' : trx.type) }}
                  </p>
                  <p class="text-xs text-muted-foreground mt-0.5" style="font-family: 'DM Mono', monospace;">{{ new Date(trx.created_at).toLocaleDateString('id-ID') }}</p>
                </div>
              </div>
              <div
                  class="text-sm font-bold"
                  :class="trx.type === 'topup' ? 'text-green-600' : 'text-foreground'"
                  style="font-family: 'DM Mono', monospace;"
              >
                {{ trx.type === 'topup' ? '+' : '-' }} {{ formatRupiah(trx.amount) }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="space-y-6">
        <div class="bg-card border border-border rounded-3xl p-6 relative overflow-hidden group cursor-pointer hover:border-primary/50 transition-colors">
          <div class="flex items-center justify-between mb-4">
            <div class="flex items-center gap-2 text-accent">
              <Sprout :size="18" />
              <span class="text-xs font-bold tracking-wider" style="font-family: 'DM Mono', monospace;">KAMPANYE DIDUKUNG</span>
            </div>
          </div>

          <h4 class="text-base font-bold text-foreground mb-4 leading-snug" style="font-family: 'Fraunces', serif;">
            {{ activeCampaign.title }}
          </h4>

          <div class="w-full h-2 bg-muted rounded-full overflow-hidden mb-3">
            <div
                class="h-full bg-accent rounded-full transition-all duration-1000"
                :style="`width: ${Math.min(activeCampaign.progress, 100)}%`"
            ></div>
          </div>

          <div class="flex justify-between items-center text-xs" style="font-family: 'DM Mono', monospace;">
            <span class="font-bold text-foreground">{{ formatRupiah(activeCampaign.raised) }}</span>
            <span class="text-muted-foreground">dari {{ formatRupiah(activeCampaign.target) }}</span>
          </div>

          <button @click="isFundModalOpen = true" class="w-full mt-5 bg-primary/10 text-primary font-bold py-2 rounded-xl hover:bg-primary/20 transition-colors text-sm">
            Bantu Sekarang
          </button>
        </div>

        <div class="bg-muted/50 border border-border rounded-3xl p-6 text-center">
          <div class="w-12 h-12 bg-card rounded-full flex items-center justify-center mx-auto mb-3 shadow-sm border border-border">
            <Users :size="20" class="text-primary" />
          </div>
          <h4 class="text-sm font-bold text-foreground mb-1" style="font-family: 'Fraunces', serif;">Tagihan Iuran Rutin</h4>
          <p class="text-xs text-muted-foreground mb-4 leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
            Kas {{ activeGroup.name }} bulan ini.
          </p>
          <button @click="handlePayKas" :disabled="isProcessing" class="w-full bg-primary text-primary-foreground text-xs font-bold py-2.5 rounded-xl hover:opacity-90 transition-opacity disabled:opacity-50">
            {{ isProcessing ? 'Memproses...' : `Bayar ${formatRupiah(activeGroup.monthly_fee)}` }}
          </button>
        </div>

      </div>
    </div>
  </div>

  <Transition enter-active-class="transition-opacity duration-300" enter-from-class="opacity-0" enter-to-class="opacity-100" leave-active-class="transition-opacity duration-300" leave-from-class="opacity-100" leave-to-class="opacity-0">
    <div v-if="isTopUpModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="isTopUpModalOpen = false"></div>
      <div class="bg-card border border-border rounded-3xl w-full max-w-sm shadow-2xl relative z-10 p-6">
        <h2 class="text-xl font-bold mb-4" style="font-family: 'Fraunces', serif;">Isi Saldo Simulasi</h2>
        <input v-model.number="topUpAmount" type="number" placeholder="Nominal (contoh: 50000)" class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm mb-4 focus:outline-none focus:border-primary" />

        <button @click="handleTopUp" :disabled="isProcessing" class="w-full bg-primary text-white py-3 rounded-xl font-bold disabled:opacity-50">
          {{ isProcessing ? 'Memproses...' : 'Konfirmasi Top-Up' }}
        </button>
      </div>
    </div>
  </Transition>

  <Transition enter-active-class="transition-opacity duration-300" enter-from-class="opacity-0" enter-to-class="opacity-100" leave-active-class="transition-opacity duration-300" leave-from-class="opacity-100" leave-to-class="opacity-0">
    <div v-if="isFundModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="isFundModalOpen = false"></div>
      <div class="bg-card border border-border rounded-3xl w-full max-w-sm shadow-2xl relative z-10 p-6">
        <h2 class="text-xl font-bold mb-1" style="font-family: 'Fraunces', serif;">Kirim Dana</h2>
        <p class="text-xs text-muted-foreground mb-4">Untuk: {{ activeCampaign.title }}</p>

        <input v-model.number="fundAmount" type="number" placeholder="Nominal Donasi (contoh: 20000)" class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm mb-4 focus:outline-none focus:border-primary" />

        <button @click="handleFund" :disabled="isProcessing" class="w-full bg-primary text-white py-3 rounded-xl font-bold disabled:opacity-50">
          {{ isProcessing ? 'Memproses...' : 'Salurkan Dana' }}
        </button>
      </div>
    </div>
  </Transition>
</template>

