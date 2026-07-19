<script setup>
import { ref, computed } from 'vue'
import { CreditCard, Plus, ArrowRight, AlertCircle, X } from '@lucide/vue'
import { useLoanStore } from '../../stores/loanStore'
import { useAuthStore } from '../../stores/authStore'

const loanStore = useLoanStore()
const authStore = useAuthStore()

const isVerified = computed(() => authStore.user?.is_verified)
const isModalOpen = ref(false)

const loanForm = ref({
  title: '',
  total_amount: '',
  monthly_installment: ''
})

const submitLoan = async () => {
  const payload = {
    title: loanForm.value.title,
    total_amount: Number(loanForm.value.total_amount),
    monthly_installment: Number(loanForm.value.monthly_installment)
  }

  const success = await loanStore.applyLoan(payload)
  if (success) {
    alert("Pinjaman berhasil diajukan dan menunggu persetujuan Admin!")
    isModalOpen.value = false
    loanForm.value = { title: '', total_amount: '', monthly_installment: '' }
  } else {
    alert(loanStore.error)
  }
}

const payLoan = async (loanId) => {
  const success = await loanStore.payInstallment(loanId)
  if (success) {
    alert("Cicilan berhasil dibayar!")
  } else {
    alert(loanStore.error)
  }
}

const formatRupiah = (angka) => new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka || 0)

const activeLoans = ref([
  { id: 1, title: 'Modal Pupuk Musim Tanam', total_amount: 2000000, paid_amount: 500000, monthly_installment: 500000, status: 'active', due_date: '2026-08-15' }
])
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">
    <div class="flex justify-between items-center">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">Pinjaman Mikro</h1>
        <p class="text-sm text-muted-foreground mt-1">Kelola kredit usaha Anda dengan bunga 0% berbasis komunitas.</p>
      </div>
      <button @click="isModalOpen = true" :disabled="!isVerified" class="bg-primary text-primary-foreground px-5 py-2.5 rounded-xl font-bold text-sm flex items-center gap-2 hover:opacity-90 disabled:opacity-50">
        <Plus :size="18" /> Ajukan Pinjaman
      </button>
    </div>

    <div v-if="!isVerified" class="bg-amber-50 text-amber-700 p-4 rounded-xl text-sm font-bold flex items-center gap-3">
      <AlertCircle :size="20" /> Anda harus melakukan verifikasi fisik di Agen Lokal untuk mengajukan pinjaman.
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
      <div v-for="loan in activeLoans" :key="loan.id" class="bg-card border border-border rounded-3xl p-6 shadow-sm">
        <div class="flex justify-between items-start mb-4">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 bg-primary/10 rounded-full flex items-center justify-center text-primary">
              <CreditCard :size="20" />
            </div>
            <div>
              <h3 class="font-bold text-foreground">{{ loan.title }}</h3>
              <p class="text-xs text-muted-foreground">Jatuh Tempo: {{ loan.due_date }}</p>
            </div>
          </div>
          <span class="bg-blue-100 text-blue-700 text-xs font-bold px-2 py-1 rounded-md">Aktif</span>
        </div>

        <div class="space-y-2 mb-6">
          <div class="flex justify-between text-sm">
            <span class="text-muted-foreground">Terbayar</span>
            <span class="font-bold">{{ formatRupiah(loan.paid_amount) }} / {{ formatRupiah(loan.total_amount) }}</span>
          </div>
          <div class="w-full h-2 bg-muted rounded-full overflow-hidden">
            <div class="h-full bg-primary" :style="`width: ${(loan.paid_amount / loan.total_amount) * 100}%`"></div>
          </div>
        </div>

        <button @click="payLoan(loan.id)" :disabled="loanStore.isLoading" class="w-full bg-accent text-accent-foreground py-3 rounded-xl text-sm font-bold flex justify-center items-center gap-2 hover:opacity-90">
          {{ loanStore.isLoading ? 'Memproses...' : `Bayar Cicilan (${formatRupiah(loan.monthly_installment)})` }}
        </button>
      </div>
    </div>

    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="isModalOpen = false"></div>
      <div class="bg-card rounded-3xl w-full max-w-md shadow-2xl relative z-10 p-6">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-xl font-bold" style="font-family: 'Fraunces', serif;">Formulir Pinjaman</h2>
          <button @click="isModalOpen = false"><X :size="20" class="text-muted-foreground" /></button>
        </div>
        <form @submit.prevent="submitLoan" class="space-y-4">
          <input v-model="loanForm.title" type="text" placeholder="Tujuan Pinjaman (cth: Beli Bibit)" class="w-full bg-muted/30 border rounded-xl py-3 px-4 text-sm" required />
          <input v-model="loanForm.total_amount" type="number" placeholder="Total Pinjaman (Rp)" class="w-full bg-muted/30 border rounded-xl py-3 px-4 text-sm" required />
          <input v-model="loanForm.monthly_installment" type="number" placeholder="Kesanggupan Cicilan per Bulan (Rp)" class="w-full bg-muted/30 border rounded-xl py-3 px-4 text-sm" required />
          <button type="submit" class="w-full bg-primary text-white py-3 rounded-xl font-bold mt-2">Ajukan Sekarang</button>
        </form>
      </div>
    </div>
  </div>
</template>