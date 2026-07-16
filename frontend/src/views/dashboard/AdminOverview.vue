<script setup>
import {
  BarChart3,
  TrendingUp,
  AlertTriangle,
  Activity,
  ArrowUpRight
} from '@lucide/vue'

const platformStats = [
  { label: "VOLUME TRANSAKSI (30 HARI)", value: "Rp 1.25M", trend: "+12.5%", isPositive: true },
  { label: "PENDAPATAN PLATFORM (FEE)", value: "Rp 12.500.000", trend: "+8.2%", isPositive: true },
  { label: "KREDIT MACET (NPL)", value: "1.2%", trend: "-0.4%", isPositive: true }, // NPL turun berarti bagus
]

const systemLogs = [
  { id: 1, action: "Pencairan Urun Dana 'Desa Makmur' otomatis dieksekusi", time: "10 menit lalu", type: "info" },
  { id: 2, action: "Lonjakan pendaftaran (50+ user) terdeteksi di Region Jatim", time: "1 jam lalu", type: "warning" },
  { id: 3, action: "Agen 'Budi Toko' disetujui oleh Super Admin", time: "3 jam lalu", type: "success" },
]
</script>

<template>
  <div class="space-y-6 max-w-7xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Dasbor Super Admin
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Pantau kesehatan finansial platform dan aktivitas sistem secara *real-time*.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div
          v-for="stat in platformStats"
          :key="stat.label"
          class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col justify-between"
      >
        <p class="text-xs font-bold text-muted-foreground tracking-wider mb-2" style="font-family: 'DM Mono', monospace;">
          {{ stat.label }}
        </p>
        <div class="flex items-end justify-between">
          <h2 class="text-3xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
            {{ stat.value }}
          </h2>
          <span
              class="text-xs font-bold px-2 py-1 rounded-md"
              :class="stat.isPositive ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'"
          >
            {{ stat.trend }}
          </span>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">

      <div class="lg:col-span-2 bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Pertumbuhan Arus Kas</h3>
          <div class="flex gap-2">
            <span class="flex items-center gap-1.5 text-xs font-bold text-muted-foreground"><span class="w-3 h-3 rounded-sm bg-primary"></span> Cash In</span>
            <span class="flex items-center gap-1.5 text-xs font-bold text-muted-foreground"><span class="w-3 h-3 rounded-sm bg-accent"></span> Cash Out</span>
          </div>
        </div>

        <div class="flex-1 flex items-end justify-between gap-2 mt-4 pt-10 border-b border-border/50 pb-2 relative h-48">
          <div v-for="i in 7" :key="i" class="w-full flex flex-col justify-end gap-1 group">
            <div class="w-full bg-accent rounded-t-sm transition-all duration-500 hover:opacity-80" :style="`height: ${Math.floor(Math.random() * 40 + 20)}%`"></div>
            <div class="w-full bg-primary rounded-t-sm transition-all duration-500 hover:opacity-80" :style="`height: ${Math.floor(Math.random() * 60 + 40)}%`"></div>
          </div>
        </div>
        <div class="flex justify-between mt-2 text-xs text-muted-foreground font-medium" style="font-family: 'DM Mono', monospace;">
          <span>Sen</span><span>Sel</span><span>Rab</span><span>Kam</span><span>Jum</span><span>Sab</span><span>Min</span>
        </div>
      </div>

      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col">
        <div class="flex items-center gap-2 mb-6">
          <Activity :size="20" class="text-muted-foreground" />
          <h3 class="font-bold text-foreground text-lg" style="font-family: 'Fraunces', serif;">Log Sistem</h3>
        </div>

        <div class="flex-1 space-y-4">
          <div v-for="log in systemLogs" :key="log.id" class="flex gap-3">
            <div class="mt-1">
              <div v-if="log.type === 'info'" class="w-2 h-2 rounded-full bg-blue-500"></div>
              <div v-else-if="log.type === 'warning'" class="w-2 h-2 rounded-full bg-amber-500"></div>
              <div v-else class="w-2 h-2 rounded-full bg-green-500"></div>
            </div>
            <div>
              <p class="text-sm font-medium text-foreground leading-snug">{{ log.action }}</p>
              <p class="text-xs text-muted-foreground mt-1">{{ log.time }}</p>
            </div>
          </div>
        </div>
        <button class="mt-4 text-xs font-bold text-primary hover:underline text-left" style="font-family: 'DM Mono', monospace;">Lihat Semua Log &rarr;</button>
      </div>

    </div>
  </div>
</template>