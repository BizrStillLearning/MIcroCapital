<script setup>
import { onMounted, computed } from 'vue'
import { UserCheck, Clock, Search, ShieldCheck } from '@lucide/vue'
import { useAgentStore } from '../../stores/agentStore'

const agentStore = useAgentStore()
const unverifiedUsers = computed(() => agentStore.unverifiedUsers)

onMounted(() => {
  agentStore.fetchUnverifiedUsers()
})

const handleVerify = async (userId, name) => {
  const confirmVerify = confirm(`Pastikan wajah dan KTP ${name} sudah sesuai. Lanjutkan verifikasi?`)
  if (!confirmVerify) return

  const success = await agentStore.verifyUser(userId)
  if (success) {
    alert(`Verifikasi ${name} berhasil!`)
  } else {
    alert(agentStore.error)
  }
}
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">Verifikasi Identitas (KYC)</h1>
        <p class="text-sm text-muted-foreground mt-1">Cocokkan KTP fisik warga sebelum membuka akses finansial mereka.</p>
      </div>
      <div class="relative w-full sm:w-64">
        <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground" :size="16" />
        <input type="text" placeholder="Cari nama warga..." class="w-full bg-card border border-border rounded-lg py-2 pl-9 pr-4 text-sm focus:border-primary" />
      </div>
    </div>

    <div v-if="unverifiedUsers.length === 0" class="text-center py-20 bg-card border border-border rounded-3xl shadow-sm">
      <ShieldCheck class="mx-auto text-green-500 mb-4" :size="48" />
      <h3 class="text-lg font-bold text-foreground">Semua Warga Terverifikasi</h3>
      <p class="text-sm text-muted-foreground mt-1">Tidak ada antrean verifikasi saat ini.</p>
    </div>

    <div v-else class="bg-card border border-border rounded-2xl overflow-hidden shadow-sm">
      <table class="w-full text-left text-sm">
        <thead class="bg-muted/50 border-b border-border">
        <tr>
          <th class="p-4 font-bold text-muted-foreground">ID</th>
          <th class="p-4 font-bold text-muted-foreground">Nama Lengkap (Sesuai KTP)</th>
          <th class="p-4 font-bold text-muted-foreground">Nomor Telepon</th>
          <th class="p-4 font-bold text-muted-foreground">Status</th>
          <th class="p-4 font-bold text-muted-foreground text-right">Aksi</th>
        </tr>
        </thead>
        <tbody class="divide-y divide-border">
        <tr v-for="user in unverifiedUsers" :key="user.id" class="hover:bg-muted/20 transition-colors">
          <td class="p-4" style="font-family: 'DM Mono', monospace;">#{{ user.id }}</td>
          <td class="p-4 font-bold">{{ user.name }}</td>
          <td class="p-4">{{ user.phone }}</td>
          <td class="p-4">
              <span class="inline-flex items-center gap-1.5 bg-amber-50 text-amber-700 px-2.5 py-1 rounded-md text-xs font-bold">
                <Clock :size="14" /> Menunggu Fisik
              </span>
          </td>
          <td class="p-4 text-right">
            <button
                @click="handleVerify(user.id, user.name)"
                :disabled="agentStore.isLoading"
                class="bg-primary text-primary-foreground px-4 py-2 rounded-lg font-bold text-xs hover:opacity-90 transition-opacity flex items-center justify-end gap-2 ml-auto"
            >
              <UserCheck :size="16" /> Verifikasi Warga
            </button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>