<script setup>
import { ref, onMounted, computed } from 'vue'
import { Search, Filter, CheckCircle, XCircle, MoreVertical } from '@lucide/vue'
import { useAdminStore } from '../../stores/adminStore'
import Swal from 'sweetalert2'

const adminStore = useAdminStore()
const agents = ref([])
const searchQuery = ref('')

const loadAgents = async () => {
  agents.value = await adminStore.fetchAgents()
}

onMounted(() => {
  loadAgents()
})

const formatRupiah = (angka) => {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(angka || 0)
}

const filteredAgents = computed(() => {
  if (!searchQuery.value) return agents.value
  return agents.value.filter(a =>
      a.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const handleApprove = async (agent) => {
  const confirm = await Swal.fire({
    title: 'Setujui Agen?',
    text: `Beri akses operasi kepada ${agent.name}?`,
    icon: 'question',
    showCancelButton: true,
    confirmButtonColor: '#10b981',
    confirmButtonText: 'Ya, Setujui'
  })

  if (confirm.isConfirmed) {
    const success = await adminStore.approveAgent(agent.id)
    if (success) {
      Swal.fire('Berhasil!', 'Agen telah aktif.', 'success')
      loadAgents()
    } else {
      Swal.fire('Gagal', 'Terjadi kesalahan sistem.', 'error')
    }
  }
}
</script>

<template>
  <div class="space-y-6 max-w-7xl mx-auto">

    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-5">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
          Manajemen Agen Lokal
        </h1>
        <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
          Tinjau pendaftaran agen baru dan pantau likuiditas agen aktif di lapangan.
        </p>
      </div>
    </div>

    <div class="flex flex-col sm:flex-row gap-3">
      <div class="relative flex-1">
        <Search class="absolute left-4 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
        <input type="text" placeholder="Cari nama agen, toko, atau wilayah..." class="w-full bg-card border border-border rounded-xl py-3 pl-11 pr-4 text-sm focus:outline-none focus:border-primary transition-colors shadow-sm" />
      </div>
      <button class="flex items-center justify-center gap-2 bg-card border border-border text-foreground px-5 py-3 rounded-xl font-bold text-sm hover:bg-muted transition-colors shadow-sm">
        <Filter :size="18" /> Filter Status
      </button>
    </div>

    <div class="bg-card border border-border rounded-3xl overflow-hidden shadow-sm">
      <div class="overflow-x-auto">
        <table class="w-full text-left text-sm border-collapse">
          <thead>
          <tr class="bg-muted/30 text-muted-foreground" style="font-family: 'DM Mono', monospace;">
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">NAMA AGEN & TOKO</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">WILAYAH</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">SALDO KAS (DIGITAL)</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">STATUS</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider text-right">AKSI</th>
          </tr>
          </thead>
          <tbody class="divide-y divide-border">
          <tr v-for="agent in filteredAgents" :key="agent.id" class="hover:bg-muted/10 transition-colors">
            <td class="py-4 px-6">
              <p class="font-bold text-foreground">{{ agent.name }}</p>
              <p class="text-xs text-muted-foreground">{{ agent.phone }}</p>
            </td>
            <td class="py-4 px-6 text-muted-foreground text-xs">Pusat</td>
            <td class="py-4 px-6 font-bold" style="font-family: 'DM Mono', monospace;">{{ formatRupiah(agent.balance) }}</td>
            <td class="py-4 px-6">
        <span
            class="px-3 py-1 rounded-full text-xs font-bold"
            :class="agent.is_verified ? 'bg-green-100 text-green-700' : 'bg-amber-100 text-amber-700'"
        >
          {{ agent.is_verified ? 'Aktif' : 'Menunggu' }}
        </span>
            </td>
            <td class="py-4 px-6 text-right">
              <div v-if="!agent.is_verified" class="flex items-center justify-end gap-2">
                <button class="p-1.5 text-red-500 hover:bg-red-50 rounded-lg transition-colors" title="Tolak"><XCircle :size="20" /></button>
                <button @click="handleApprove(agent)" class="p-1.5 text-green-600 hover:bg-green-50 rounded-lg transition-colors" title="Setujui"><CheckCircle :size="20" /></button>
              </div>
              <button v-else class="p-1.5 text-muted-foreground hover:text-foreground rounded-lg transition-colors">
                <MoreVertical :size="20" />
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

