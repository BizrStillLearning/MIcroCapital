<script setup>
import {
  Users,
  RotateCw,
  Check,
  Info,
  ArrowUpRight
} from '@lucide/vue'

const groupInfo = {
  name: "Arisan Pedagang Pasar Baru",
  membersCount: 10,
  monthlyFee: 50000,
  totalPool: 500000,
  currentCycle: 4,
  recipientThisMonth: "Bapak Supardi"
}

const membersList = [
  { id: 1, name: "Grace O. (Anda)", hasPaid: true, hasReceived: false },
  { id: 2, name: "Bapak Supardi", hasPaid: true, hasReceived: true, isCurrent: true },
  { id: 3, name: "Ibu Siti", hasPaid: false, hasReceived: false },
  { id: 4, name: "Kang Asep", hasPaid: true, hasReceived: true },
]
</script>

<template>
  <div class="space-y-6 max-w-5xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Tabungan Kelompok
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Kelola iuran rutin dan pantau pencairan dana kelompok secara transparan.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">

      <div class="bg-primary text-primary-foreground rounded-3xl p-6 relative overflow-hidden shadow-xl shadow-primary/20">
        <div class="absolute -right-8 -top-8 w-32 h-32 bg-white/10 rounded-full blur-2xl"></div>

        <div class="relative z-10">
          <div class="flex items-center gap-2 mb-6 opacity-90 text-sm font-semibold tracking-wider" style="font-family: 'DM Mono', monospace;">
            <Users :size="18" /> {{ groupInfo.name }}
          </div>

          <div class="mb-2">
            <p class="text-xs opacity-80 mb-1">TOTAL DANA TERKUMPUL BULAN INI</p>
            <p class="text-3xl font-bold" style="font-family: 'Fraunces', serif;">
              Rp {{ (groupInfo.totalPool).toLocaleString('id-ID') }}
            </p>
          </div>

          <div class="flex items-center justify-between mt-8 pt-4 border-t border-white/20">
            <div>
              <p class="text-xs opacity-80">Iuran Rutin</p>
              <p class="font-bold text-sm" style="font-family: 'DM Mono', monospace;">Rp {{ (groupInfo.monthlyFee).toLocaleString('id-ID') }} / bln</p>
            </div>
            <div class="text-right">
              <p class="text-xs opacity-80">Anggota</p>
              <p class="font-bold text-sm">{{ groupInfo.membersCount }} Orang</p>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-card border border-border rounded-3xl p-6 flex flex-col justify-between shadow-sm">
        <div>
          <div class="flex items-center gap-2 mb-3">
            <RotateCw :size="20" class="text-accent" />
            <h3 class="font-bold text-foreground text-lg" style="font-family: 'Fraunces', serif;">Siklus ke-{{ groupInfo.currentCycle }}</h3>
          </div>
          <p class="text-sm text-muted-foreground leading-relaxed">
            Dana bulan ini akan dicairkan kepada <strong class="text-foreground">{{ groupInfo.recipientThisMonth }}</strong> setelah semua anggota membayar iuran.
          </p>
        </div>

        <div class="mt-6">
          <button class="w-full flex items-center justify-center gap-2 bg-primary/10 text-primary px-5 py-3.5 rounded-xl font-bold text-sm hover:bg-primary/20 transition-colors border border-primary/20">
            Bayar Iuran Saya
            <ArrowUpRight :size="18" />
          </button>
        </div>
      </div>

    </div>

    <div class="bg-card border border-border rounded-3xl overflow-hidden shadow-sm">
      <div class="p-5 border-b border-border flex items-center justify-between">
        <h3 class="text-lg font-bold text-foreground" style="font-family: 'Fraunces', serif;">Status Pembayaran Anggota</h3>
        <span class="text-xs bg-muted text-muted-foreground px-3 py-1 rounded-full font-bold">Bulan Ini</span>
      </div>

      <div class="divide-y divide-border">
        <div
            v-for="member in membersList"
            :key="member.id"
            class="p-4 flex items-center justify-between"
            :class="{ 'bg-accent/5': member.isCurrent }"
        >
          <div class="flex items-center gap-4">
            <div
                class="w-10 h-10 rounded-full flex items-center justify-center text-sm font-bold shadow-sm border border-border"
                :class="member.isCurrent ? 'bg-accent text-white border-transparent' : 'bg-muted text-foreground'"
            >
              {{ member.name.charAt(0) }}
            </div>
            <div>
              <p class="text-sm font-bold text-foreground flex items-center gap-2">
                {{ member.name }}
                <span v-if="member.isCurrent" class="text-[10px] bg-accent text-white px-2 py-0.5 rounded text-center tracking-wider">PENERIMA</span>
              </p>
              <p class="text-xs text-muted-foreground mt-0.5">
                Status: {{ member.hasReceived ? 'Sudah Dapat Giliran' : 'Belum Dapat Giliran' }}
              </p>
            </div>
          </div>

          <div>
            <div v-if="member.hasPaid" class="flex items-center gap-1.5 text-xs font-bold text-green-600 bg-green-50 px-3 py-1.5 rounded-full border border-green-200">
              <Check :size="14" /> Lunas
            </div>
            <div v-else class="text-xs font-bold text-amber-600 bg-amber-50 px-3 py-1.5 rounded-full border border-amber-200">
              Belum Bayar
            </div>
          </div>
        </div>
      </div>

      <div class="bg-muted/30 p-4 flex gap-3 text-xs text-muted-foreground">
        <Info :size="16" class="flex-shrink-0 mt-0.5" />
        <p>Dana akan otomatis ditransfer ke dompet penerima segera setelah 100% anggota melunasi iuran pada siklus bulan ini.</p>
      </div>
    </div>

  </div>
</template>



