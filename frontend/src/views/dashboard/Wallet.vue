<script setup>
import {
  Wallet,
  ArrowDownLeft,
  ArrowUpRight,
  Search,
  Download,
  Landmark
} from '@lucide/vue'

const balance = "Rp 1.250.000"

const transactions = [
  { id: 1, type: 'in', title: 'Isi Saldo via Agen (Pak Budi)', category: 'Top Up', date: '16 Jul 2026, 09:30', amount: '+ Rp 500.000' },
  { id: 2, type: 'out', title: 'Pendanaan: Kebun Sayur', category: 'Urun Dana', date: '15 Jul 2026, 14:15', amount: '- Rp 150.000' },
  { id: 3, type: 'out', title: 'Iuran Tabungan Kelompok', category: 'Tabungan', date: '12 Jul 2026, 10:00', amount: '- Rp 50.000' },
  { id: 4, type: 'in', title: 'Pelunasan Pinjaman (Siti)', category: 'Pengembalian', date: '10 Jul 2026, 16:45', amount: '+ Rp 200.000' },
  { id: 5, type: 'out', title: 'Tarik Tunai via Agen (Bu Tejo)', category: 'Penarikan', date: '05 Jul 2026, 08:20', amount: '- Rp 300.000' },
]
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Dompet & Kas
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Kelola saldo dan pantau seluruh riwayat transaksi Anda.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="md:col-span-2 bg-card border border-border rounded-3xl p-6 sm:p-8 flex flex-col sm:flex-row sm:items-center justify-between gap-6 shadow-sm">
        <div>
          <div class="flex items-center gap-2 mb-2 text-muted-foreground">
            <Wallet :size="18" />
            <span class="text-xs font-bold tracking-wider" style="font-family: 'DM Mono', monospace;">TOTAL SALDO</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
            {{ balance }}
          </h2>
        </div>

        <div class="flex gap-3">
          <button class="flex-1 sm:flex-none flex items-center justify-center gap-2 bg-primary/10 text-primary px-5 py-3 rounded-xl font-bold text-sm hover:bg-primary/20 transition-colors">
            <ArrowDownLeft :size="18" />
            Isi Saldo
          </button>
          <button class="flex-1 sm:flex-none flex items-center justify-center gap-2 bg-primary text-primary-foreground px-5 py-3 rounded-xl font-bold text-sm hover:opacity-90 transition-opacity shadow-md shadow-primary/20">
            <ArrowUpRight :size="18" />
            Tarik Tunai
          </button>
        </div>
      </div>

      <div class="bg-accent/10 border border-accent/20 rounded-3xl p-6 flex flex-col justify-center">
        <div class="w-10 h-10 bg-accent/20 rounded-full flex items-center justify-center text-accent mb-3">
          <Landmark :size="20" />
        </div>
        <h3 class="font-bold text-foreground text-sm mb-1" style="font-family: 'Fraunces', serif;">Butuh bantuan tunai?</h3>
        <p class="text-xs text-muted-foreground leading-relaxed mb-3">
          Kunjungi agen lokal terdekat untuk melakukan pengisian atau penarikan saldo secara langsung.
        </p>
        <button class="text-xs font-bold text-accent hover:underline text-left">
          Cari Agen Terdekat &rarr;
        </button>
      </div>
    </div>

    <div class="bg-card border border-border rounded-3xl overflow-hidden shadow-sm">
      <div class="p-6 border-b border-border flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Mutasi Kas</h3>

        <div class="flex items-center gap-3">
          <div class="relative flex-1 sm:w-64">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground" :size="16" />
            <input type="text" placeholder="Cari transaksi..." class="w-full bg-muted/50 border border-border rounded-lg py-2 pl-9 pr-4 text-sm focus:outline-none focus:border-primary transition-colors" />
          </div>
          <button class="p-2 border border-border rounded-lg text-muted-foreground hover:bg-muted hover:text-foreground transition-colors" title="Unduh Laporan">
            <Download :size="18" />
          </button>
        </div>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left text-sm border-collapse">
          <thead>
          <tr class="bg-muted/30 text-muted-foreground" style="font-family: 'DM Mono', monospace;">
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">TANGGAL</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">KETERANGAN</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider">KATEGORI</th>
            <th class="py-4 px-6 font-semibold text-xs tracking-wider text-right">JUMLAH</th>
          </tr>
          </thead>
          <tbody class="divide-y divide-border">
          <tr v-for="trx in transactions" :key="trx.id" class="hover:bg-muted/20 transition-colors">
            <td class="py-4 px-6 text-muted-foreground text-xs" style="font-family: 'DM Mono', monospace;">
              {{ trx.date }}
            </td>
            <td class="py-4 px-6">
              <div class="flex items-center gap-3">
                <div
                    class="w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0"
                    :class="trx.type === 'in' ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'"
                >
                  <ArrowDownLeft v-if="trx.type === 'in'" :size="14" />
                  <ArrowUpRight v-else :size="14" />
                </div>
                <span class="font-bold text-foreground">{{ trx.title }}</span>
              </div>
            </td>
            <td class="py-4 px-6">
                <span class="bg-muted px-2.5 py-1 rounded-md text-xs font-medium text-muted-foreground">
                  {{ trx.category }}
                </span>
            </td>
            <td
                class="py-4 px-6 text-right font-bold"
                :class="trx.type === 'in' ? 'text-green-600' : 'text-foreground'"
                style="font-family: 'DM Mono', monospace;"
            >
              {{ trx.amount }}
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

