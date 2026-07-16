<script setup>
import { Search, Filter, CheckCircle, XCircle, MoreVertical } from '@lucide/vue'

const agents = [
  { id: 1, name: "Bapak Budi Santoso", store: "Toko Makmur", region: "Desa Sukamaju, Jatim", balance: "Rp 5.500.000", status: "Aktif" },
  { id: 2, name: "Ibu Siti Aminah", store: "Warung Barokah", region: "Desa Karanganyar, Jateng", balance: "Rp 2.100.000", status: "Aktif" },
  { id: 3, name: "Kang Asep", store: "Koperasi Tani", region: "Desa Cibaduyut, Jabar", balance: "Rp 0", status: "Menunggu" },
]
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
          <tr v-for="agent in agents" :key="agent.id" class="hover:bg-muted/10 transition-colors">
            <td class="py-4 px-6">
              <p class="font-bold text-foreground">{{ agent.name }}</p>
              <p class="text-xs text-muted-foreground">{{ agent.store }}</p>
            </td>
            <td class="py-4 px-6 text-muted-foreground text-xs">{{ agent.region }}</td>
            <td class="py-4 px-6 font-bold" style="font-family: 'DM Mono', monospace;">{{ agent.balance }}</td>
            <td class="py-4 px-6">
                <span
                    class="px-3 py-1 rounded-full text-xs font-bold"
                    :class="agent.status === 'Aktif' ? 'bg-green-100 text-green-700' : 'bg-amber-100 text-amber-700'"
                >
                  {{ agent.status }}
                </span>
            </td>
            <td class="py-4 px-6 text-right">
              <div v-if="agent.status === 'Menunggu'" class="flex items-center justify-end gap-2">
                <button class="p-1.5 text-red-500 hover:bg-red-50 rounded-lg transition-colors" title="Tolak"><XCircle :size="20" /></button>
                <button class="p-1.5 text-green-600 hover:bg-green-50 rounded-lg transition-colors" title="Setujui"><CheckCircle :size="20" /></button>
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