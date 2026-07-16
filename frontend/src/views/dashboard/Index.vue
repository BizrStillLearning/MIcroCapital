<script setup>
import {
  ArrowUpRight,
  ArrowDownLeft,
  Wallet,
  Sprout,
  History,
  ChevronRight
} from '@lucide/vue'

const balance = "Rp 1.250.000"
const activeCampaign = {
  title: "Perbaikan Atap Warung Bu Siti",
  raised: 850000,
  target: 1000000,
  progress: 85
}

const recentTransactions = [
  { id: 1, type: 'in', title: 'Top Up via Agen (Pak Budi)', date: 'Hari ini, 09:30', amount: '+ Rp 500.000' },
  { id: 2, type: 'out', title: 'Mendanai: Kebun Sayur', date: 'Kemarin, 14:15', amount: '- Rp 150.000' },
  { id: 3, type: 'out', title: 'Iuran Tabungan Kelompok', date: '12 Jul, 10:00', amount: '- Rp 50.000' },
]
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Selamat datang, Grace!
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
              {{ balance }}
            </h2>

            <div class="flex flex-wrap gap-3">
              <button class="flex items-center gap-2 bg-white text-primary px-5 py-2.5 rounded-xl font-bold text-sm hover:bg-white/90 transition-colors shadow-sm">
                <ArrowDownLeft :size="18" />
                Isi Saldo
              </button>
              <button class="flex items-center gap-2 bg-primary-foreground/15 text-white border border-white/20 px-5 py-2.5 rounded-xl font-bold text-sm hover:bg-primary-foreground/25 transition-colors">
                <ArrowUpRight :size="18" />
                Kirim
              </button>
            </div>
          </div>
        </div>

        <div class="bg-card border border-border rounded-3xl p-6">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Aktivitas Terakhir</h3>
            <button class="text-sm text-primary font-bold hover:underline" style="font-family: 'DM Mono', monospace;">Lihat Semua</button>
          </div>

          <div class="space-y-4">
            <div
                v-for="trx in recentTransactions"
                :key="trx.id"
                class="flex items-center justify-between p-3 hover:bg-muted/50 rounded-2xl transition-colors cursor-pointer"
            >
              <div class="flex items-center gap-4">
                <div
                    class="w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0"
                    :class="trx.type === 'in' ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'"
                >
                  <ArrowDownLeft v-if="trx.type === 'in'" :size="18" />
                  <ArrowUpRight v-else :size="18" />
                </div>
                <div>
                  <p class="text-sm font-bold text-foreground" style="font-family: 'DM Sans', sans-serif;">{{ trx.title }}</p>
                  <p class="text-xs text-muted-foreground mt-0.5" style="font-family: 'DM Mono', monospace;">{{ trx.date }}</p>
                </div>
              </div>
              <div
                  class="text-sm font-bold"
                  :class="trx.type === 'in' ? 'text-green-600' : 'text-foreground'"
                  style="font-family: 'DM Mono', monospace;"
              >
                {{ trx.amount }}
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
            <ChevronRight :size="18" class="text-muted-foreground group-hover:text-primary transition-colors" />
          </div>

          <h4 class="text-base font-bold text-foreground mb-4 leading-snug" style="font-family: 'Fraunces', serif;">
            {{ activeCampaign.title }}
          </h4>

          <div class="w-full h-2 bg-muted rounded-full overflow-hidden mb-3">
            <div
                class="h-full bg-accent rounded-full transition-all duration-1000"
                :style="`width: ${activeCampaign.progress}%`"
            ></div>
          </div>

          <div class="flex justify-between items-center text-xs" style="font-family: 'DM Mono', monospace;">
            <span class="font-bold text-foreground">Rp {{ (activeCampaign.raised).toLocaleString('id-ID') }}</span>
            <span class="text-muted-foreground">dari Rp {{ (activeCampaign.target).toLocaleString('id-ID') }}</span>
          </div>
        </div>

        <div class="bg-muted/50 border border-border rounded-3xl p-6 text-center">
          <div class="w-12 h-12 bg-card rounded-full flex items-center justify-center mx-auto mb-3 shadow-sm border border-border">
            <History :size="20" class="text-primary" />
          </div>
          <h4 class="text-sm font-bold text-foreground mb-2" style="font-family: 'Fraunces', serif;">Waktunya Iuran Rutin</h4>
          <p class="text-xs text-muted-foreground mb-4 leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
            Tabungan Kelompok "Maju Bersama" jatuh tempo dalam 3 hari.
          </p>
          <button class="w-full bg-primary text-primary-foreground text-xs font-bold py-2.5 rounded-xl hover:opacity-90 transition-opacity">
            Bayar Iuran Sekarang
          </button>
        </div>

      </div>
    </div>
  </div>
</template>