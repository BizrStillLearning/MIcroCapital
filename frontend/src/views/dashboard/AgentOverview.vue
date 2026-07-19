<script setup>
import { useRouter } from 'vue-router'
import {
  Landmark,
  Users,
  ShieldAlert,
  ArrowDownLeft,
  ArrowUpRight,
  ChevronRight,
  History
} from '@lucide/vue'

const router = useRouter()

const stats = [
  { label: "KAS FISIK DI TANGAN", value: "Rp 4.350.000", desc: "Total uang tunai masuk dikurangi penarikan warga" },
  { label: "TRANSAKSI HARI INI", value: "18 Transaksi", desc: "Aktivitas Cash-In & Cash-Out" },
  { label: "ANTRAN VERIFIKASI", value: "2 Warga", desc: "Menunggu pencocokan identitas fisik" },
]

const recentActivities = [
  { id: 1, type: 'cash-in', text: 'Memproses isi saldo Ibu Siti Rohmah', amount: '+ Rp 200.000', time: '10 menit yang lalu' },
  { id: 2, type: 'kyc', text: 'Memverifikasi identitas Bapak Supardi', amount: 'Terverifikasi', time: '1 jam yang lalu' },
  { id: 3, type: 'cash-out', text: 'Memproses tarik tunai Kang Asep', amount: '- Rp 500.000', time: '2 jam yang lalu' },
]
</script>

<template>
  <div class="space-y-6 max-w-6xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Pusat Kendali Agen
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Pantau likuiditas kas fisik dan kelola permintaan layanan warga di wilayah Anda.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div
          v-for="(stat, index) in stats"
          :key="stat.label"
          class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col justify-between relative overflow-hidden"
      >
        <div>
          <p class="text-xs font-bold text-muted-foreground tracking-wider mb-2" style="font-family: 'DM Mono', monospace;">
            {{ stat.label }}
          </p>
          <h2
              class="text-3xl font-bold mb-2"
              :class="index === 0 ? 'text-primary' : 'text-foreground'"
              style="font-family: 'Fraunces', serif;"
          >
            {{ stat.value }}
          </h2>
        </div>
        <p class="text-xs text-muted-foreground leading-relaxed mt-4 border-t border-border/50 pt-3">
          {{ stat.desc }}
        </p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">

      <div class="lg:col-span-2 space-y-6">
        <div class="bg-card border border-border rounded-3xl p-6">
          <h3 class="text-lg font-bold text-foreground mb-4" style="font-family: 'Fraunces', serif;">Operasi Lapangan</h3>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div
                @click="router.push('/dashboard/cash-management')"
                class="border border-border rounded-2xl p-5 hover:border-primary/50 cursor-pointer transition-colors bg-muted/20 flex items-start gap-4 group"
            >
              <div class="w-10 h-10 bg-primary/10 text-primary rounded-xl flex items-center justify-center flex-shrink-0">
                <Landmark :size="20" />
              </div>
              <div class="flex-1">
                <h4 class="font-bold text-sm text-foreground flex items-center justify-between">
                  Manajemen Tunai
                  <ChevronRight :size="16" class="text-muted-foreground group-hover:text-primary transition-colors" />
                </h4>
                <p class="text-xs text-muted-foreground mt-1 leading-relaxed">Proses setor tunai atau penarikan saldo digital warga menggunakan nomor telepon.</p>
              </div>
            </div>

            <div
                @click="router.push('/dashboard/verification')"
                class="border border-border rounded-2xl p-5 hover:border-accent/50 cursor-pointer transition-colors bg-muted/20 flex items-start gap-4 group"
            >
              <div class="w-10 h-10 bg-accent/10 text-accent rounded-xl flex items-center justify-center flex-shrink-0">
                <Users :size="20" />
              </div>
              <div class="flex-1">
                <h4 class="font-bold text-sm text-foreground flex items-center justify-between">
                  Verifikasi Warga
                  <ChevronRight :size="16" class="text-muted-foreground group-hover:text-accent transition-colors" />
                </h4>
                <p class="text-xs text-muted-foreground mt-1 leading-relaxed">Lakukan pencocokan berkas identitas resmi (NIK/KTP) bagi pendaftar baru.</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col">
        <div class="flex items-center gap-2 mb-4">
          <History :size="18" class="text-muted-foreground" />
          <h3 class="font-bold text-foreground" style="font-family: 'Fraunces', serif;">Log Aktivitas</h3>
        </div>

        <div class="flex-1 space-y-4 overflow-y-auto">
          <div
              v-for="activity in recentActivities"
              :key="activity.id"
              class="p-3 bg-muted/30 rounded-xl border border-border/50 text-xs space-y-2"
          >
            <div class="flex items-center justify-between font-medium">
              <span class="text-muted-foreground" style="font-family: 'DM Mono', monospace;">{{ activity.time }}</span>
              <span
                  class="font-bold"
                  :class="activity.type === 'cash-in' ? 'text-green-600' : activity.type === 'cash-out' ? 'text-red-600' : 'text-accent'"
              >
                {{ activity.amount }}
              </span>
            </div>
            <p class="text-foreground leading-relaxed font-medium">{{ activity.text }}</p>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

