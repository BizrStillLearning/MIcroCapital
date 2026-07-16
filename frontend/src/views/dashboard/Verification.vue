<script setup>
import { ShieldAlert, Check, X, Search, FileText } from '@lucide/vue'

const pendingUsers = [
  { id: 1, name: "Ibu Siti Rohmah", phone: "+62 852 1122 3344", requestDate: "Hari ini, 08:15", idNumber: "3578123456780001" },
  { id: 2, name: "Bapak Yanto", phone: "+62 813 9988 7766", requestDate: "Kemarin, 16:30", idNumber: "3578987654320002" }
]
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">

    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-5">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
          Verifikasi Identitas Warga
        </h1>
        <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
          Pastikan wajah warga sesuai dengan dokumen identitas asli mereka.
        </p>
      </div>

      <div class="bg-amber-100 text-amber-700 px-4 py-2 rounded-xl text-sm font-bold border border-amber-200 flex items-center gap-2">
        <ShieldAlert :size="18" />
        {{ pendingUsers.length }} Menunggu Verifikasi
      </div>
    </div>

    <div class="relative max-w-md">
      <Search class="absolute left-4 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
      <input type="text" placeholder="Cari nama atau NIK..." class="w-full bg-card border border-border rounded-xl py-3 pl-11 pr-4 text-sm focus:outline-none focus:border-primary transition-colors shadow-sm" />
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

      <div v-for="user in pendingUsers" :key="user.id" class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col justify-between">

        <div class="flex items-start justify-between mb-6">
          <div>
            <h3 class="font-bold text-lg text-foreground mb-1">{{ user.name }}</h3>
            <p class="text-sm text-muted-foreground font-medium" style="font-family: 'DM Mono', monospace;">{{ user.phone }}</p>
          </div>
          <span class="text-xs text-muted-foreground">{{ user.requestDate }}</span>
        </div>

        <div class="bg-muted/30 p-4 rounded-xl border border-border/50 mb-6 flex items-center gap-3">
          <FileText :size="24" class="text-primary opacity-70" />
          <div>
            <p class="text-xs text-muted-foreground mb-0.5">Nomor Identitas (NIK)</p>
            <p class="text-sm font-bold font-mono tracking-widest text-foreground">{{ user.idNumber }}</p>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-3 mt-auto">
          <button class="flex items-center justify-center gap-2 bg-red-50 text-red-600 border border-red-200 py-2.5 rounded-xl text-sm font-bold hover:bg-red-100 transition-colors">
            <X :size="18" /> Tolak
          </button>
          <button class="flex items-center justify-center gap-2 bg-primary text-primary-foreground py-2.5 rounded-xl text-sm font-bold hover:opacity-90 shadow-md transition-all">
            <Check :size="18" /> Verifikasi
          </button>
        </div>

      </div>

    </div>
  </div>
</template>