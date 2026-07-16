<script setup>
import { ref } from 'vue'
import { Search, ArrowDownLeft, ArrowUpRight, UserCircle2, CheckCircle2, History } from '@lucide/vue'

const searchQuery = ref('')
const isUserFound = ref(false)

const searchUser = () => {
  if (searchQuery.value.length > 8) {
    isUserFound.value = true
  } else {
    isUserFound.value = false
  }
}

const foundUser = {
  name: "Bapak Supardi",
  phone: "+62 812 3456 7890",
  status: "Terverifikasi"
}

const agentTransactions = [
  { id: 1, type: 'in', client: 'Ibu Siti', phone: '...1234', amount: 'Rp 200.000', time: '14:30' },
  { id: 2, type: 'out', client: 'Kang Asep', phone: '...9876', amount: 'Rp 500.000', time: '11:15' },
]
</script>

<template>
  <div class="space-y-6 max-w-6xl mx-auto">

    <div>
      <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
        Manajemen Tunai
      </h1>
      <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
        Fasilitasi isi saldo (Cash-In) dan tarik tunai (Cash-Out) untuk warga komunitas.
      </p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">

      <div class="lg:col-span-2 space-y-6">

        <div class="bg-card border border-border rounded-3xl p-6 sm:p-8 shadow-sm">
          <label class="text-xs font-bold text-muted-foreground tracking-wider mb-3 block" style="font-family: 'DM Mono', monospace;">CARI WARGA (NO. TELEPON)</label>
          <div class="relative flex items-center mb-6">
            <Search class="absolute left-4 text-muted-foreground" :size="20" />
            <input
                v-model="searchQuery"
                @input="searchUser"
                type="tel"
                placeholder="Contoh: 081234567890"
                class="w-full bg-muted/30 border border-border rounded-xl py-4 pl-12 pr-4 text-lg font-medium focus:outline-none focus:border-primary transition-colors"
            />
          </div>

          <Transition
              enter-active-class="transition-opacity duration-300"
              enter-from-class="opacity-0"
              enter-to-class="opacity-100"
          >
            <div v-if="isUserFound" class="bg-muted/30 border border-border rounded-2xl p-5">
              <div class="flex items-center justify-between mb-5">
                <div class="flex items-center gap-4">
                  <div class="w-12 h-12 bg-primary/10 text-primary rounded-full flex items-center justify-center">
                    <UserCircle2 :size="24" />
                  </div>
                  <div>
                    <h3 class="font-bold text-foreground text-lg">{{ foundUser.name }}</h3>
                    <p class="text-sm text-muted-foreground" style="font-family: 'DM Mono', monospace;">{{ foundUser.phone }}</p>
                  </div>
                </div>
                <div class="bg-green-100 text-green-700 text-xs font-bold px-3 py-1.5 rounded-full border border-green-200 flex items-center gap-1">
                  <CheckCircle2 :size="14" /> {{ foundUser.status }}
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4 pt-5 border-t border-border">
                <button class="flex flex-col items-center justify-center gap-2 bg-primary text-primary-foreground p-4 rounded-xl hover:opacity-90 transition-opacity shadow-md shadow-primary/20">
                  <ArrowDownLeft :size="24" />
                  <span class="font-bold text-sm">Terima Tunai (Isi Saldo)</span>
                </button>
                <button class="flex flex-col items-center justify-center gap-2 bg-card border border-border text-foreground p-4 rounded-xl hover:bg-muted transition-colors shadow-sm">
                  <ArrowUpRight :size="24" />
                  <span class="font-bold text-sm">Berikan Tunai (Tarik)</span>
                </button>
              </div>
            </div>
          </Transition>

          <div v-if="!isUserFound && searchQuery.length > 0" class="text-center py-8 text-muted-foreground text-sm">
            Nomor telepon tidak ditemukan atau belum terdaftar di sistem.
          </div>
        </div>

      </div>

      <div class="bg-card border border-border rounded-3xl p-6 shadow-sm flex flex-col">
        <div class="flex items-center gap-2 mb-6">
          <History :size="20" class="text-muted-foreground" />
          <h3 class="font-bold text-foreground text-lg" style="font-family: 'Fraunces', serif;">Riwayat Hari Ini</h3>
        </div>

        <div class="flex-1 overflow-y-auto space-y-4">
          <div v-for="trx in agentTransactions" :key="trx.id" class="flex items-center justify-between p-3 bg-muted/20 rounded-xl border border-border/50">
            <div class="flex items-center gap-3">
              <div
                  class="w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0"
                  :class="trx.type === 'in' ? 'bg-primary/10 text-primary' : 'bg-accent/10 text-accent'"
              >
                <ArrowDownLeft v-if="trx.type === 'in'" :size="18" />
                <ArrowUpRight v-else :size="18" />
              </div>
              <div>
                <p class="text-sm font-bold text-foreground">{{ trx.client }}</p>
                <p class="text-xs text-muted-foreground">{{ trx.phone }} • {{ trx.time }}</p>
              </div>
            </div>
            <p class="text-sm font-bold" style="font-family: 'DM Mono', monospace;">{{ trx.amount }}</p>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>