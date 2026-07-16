<script setup>
import {
  CreditCard,
  CalendarCheck,
  AlertCircle,
  CheckCircle2,
  ChevronRight,
  Plus
} from '@lucide/vue'

const activeLoan = {
  title: "Modal Gerobak Bakso",
  totalAmount: 1000000,
  paidAmount: 600000,
  remainingAmount: 400000,
  nextInstallment: 100000,
  dueDate: "20 Jul 2026",
  status: "Lancar"
}

const paymentHistory = [
  { id: 1, date: "20 Jun 2026", amount: 100000, status: "Lunas" },
  { id: 2, date: "20 Mei 2026", amount: 100000, status: "Lunas" },
  { id: 3, date: "20 Apr 2026", amount: 100000, status: "Lunas" },
]
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">

    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-5">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
          Pinjaman Mikro
        </h1>
        <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
          Pantau status pinjaman aktif dan jadwal cicilan Anda.
        </p>
      </div>

      <button class="flex items-center justify-center gap-2 bg-muted text-muted-foreground px-5 py-3 rounded-xl font-bold text-sm cursor-not-allowed">
        <Plus :size="18" />
        Ajukan Pinjaman Baru
      </button>
    </div>

    <div class="bg-card border border-border rounded-3xl p-6 sm:p-8 shadow-sm relative overflow-hidden">
      <div class="absolute top-6 right-6 bg-green-100 text-green-700 px-3 py-1 rounded-full text-xs font-bold flex items-center gap-1.5 border border-green-200">
        <CheckCircle2 :size="14" />
        Status: {{ activeLoan.status }}
      </div>

      <div class="mb-8 pr-32">
        <p class="text-xs font-bold text-muted-foreground tracking-wider mb-2" style="font-family: 'DM Mono', monospace;">PINJAMAN AKTIF</p>
        <h2 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">{{ activeLoan.title }}</h2>
      </div>

      <div class="mb-8">
        <div class="flex justify-between items-end mb-3">
          <div>
            <p class="text-3xl font-bold text-primary" style="font-family: 'Fraunces', serif;">
              Rp {{ (activeLoan.paidAmount).toLocaleString('id-ID') }}
            </p>
            <p class="text-xs text-muted-foreground font-medium mt-1">Telah dibayar dari Rp {{ (activeLoan.totalAmount).toLocaleString('id-ID') }}</p>
          </div>
          <div class="text-right">
            <span class="text-sm font-bold text-foreground" style="font-family: 'DM Mono', monospace;">
              {{ Math.round((activeLoan.paidAmount / activeLoan.totalAmount) * 100) }}%
            </span>
          </div>
        </div>

        <div class="w-full h-3 bg-muted rounded-full overflow-hidden">
          <div
              class="h-full bg-primary rounded-full transition-all duration-1000"
              :style="`width: ${(activeLoan.paidAmount / activeLoan.totalAmount) * 100}%`"
          ></div>
        </div>
      </div>

      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-5 bg-accent/10 border border-accent/20 rounded-2xl p-5">
        <div class="flex items-center gap-4">
          <div class="w-12 h-12 bg-accent/20 text-accent rounded-full flex items-center justify-center flex-shrink-0">
            <CalendarCheck :size="24" />
          </div>
          <div>
            <p class="text-xs font-bold text-accent mb-1" style="font-family: 'DM Mono', monospace;">TAGIHAN BERIKUTNYA</p>
            <p class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">
              Rp {{ (activeLoan.nextInstallment).toLocaleString('id-ID') }}
            </p>
            <p class="text-xs text-muted-foreground flex items-center gap-1 mt-0.5">
              <AlertCircle :size="12" /> Jatuh tempo: {{ activeLoan.dueDate }}
            </p>
          </div>
        </div>
        <button class="w-full sm:w-auto bg-primary text-primary-foreground font-bold px-6 py-3 rounded-xl hover:opacity-90 transition-opacity shadow-md shadow-primary/20">
          Bayar Cicilan
        </button>
      </div>
    </div>

    <div class="bg-card border border-border rounded-3xl overflow-hidden shadow-sm">
      <div class="p-6 border-b border-border">
        <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Riwayat Pembayaran</h3>
      </div>
      <div class="divide-y divide-border">
        <div
            v-for="payment in paymentHistory"
            :key="payment.id"
            class="p-5 flex items-center justify-between hover:bg-muted/30 transition-colors"
        >
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 bg-green-100 text-green-600 rounded-full flex items-center justify-center">
              <CheckCircle2 :size="20" />
            </div>
            <div>
              <p class="text-sm font-bold text-foreground">Cicilan Bulanan</p>
              <p class="text-xs text-muted-foreground mt-0.5" style="font-family: 'DM Mono', monospace;">{{ payment.date }}</p>
            </div>
          </div>
          <div class="text-right">
            <p class="text-sm font-bold text-foreground" style="font-family: 'DM Mono', monospace;">Rp {{ (payment.amount).toLocaleString('id-ID') }}</p>
            <p class="text-xs font-bold text-green-600 mt-0.5">{{ payment.status }}</p>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>