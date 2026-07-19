<script setup>
import { ref, onMounted, computed } from 'vue'
import {
  Search,
  Filter,
  Plus,
  Users,
  X,
  Info
} from '@lucide/vue'
import { useCampaignStore } from '../../stores/campaignStore'
import { useAuthStore } from '../../stores/authStore.js'
import { useTransactionStore } from '../../stores/transactionStore'
const transactionStore = useTransactionStore()

const isFundModalOpen = ref(false)
const selectedCampaign = ref(null)
const fundAmount = ref('')

const openFundModal = (campaign) => {
  selectedCampaign.value = campaign
  isFundModalOpen.value = true
}

const handleFunding = async () => {
  if (!fundAmount.value || fundAmount.value <= 0) return

  const success = await transactionStore.fundCampaign(selectedCampaign.value.id, Number(fundAmount.value))

  if (success) {
    isFundModalOpen.value = false
    fundAmount.value = ''
    campaignStore.fetchCampaigns()
    alert("Pendanaan berhasil! Saldo Anda telah dipotong.")
  } else {
    alert(transactionStore.error)
  }
}

const isModalOpen = ref(false)

const campaignStore = useCampaignStore()
const authStore = useAuthStore()

const campaigns = computed(() => campaignStore.campaigns)
const isVerified = computed(() => authStore.user?.is_verified)

onMounted(() => {
  campaignStore.fetchCampaigns()
})

const campaignForm = ref({
  title: '',
  target: '',
  description: '',
  duration: '30'
})

const submitCampaign = async () => {
  const payload = {
    title: campaignForm.value.title,
    target_amount: Number(campaignForm.value.target),
    description: campaignForm.value.description,
    duration_days: Number(campaignForm.value.duration)
  }

  const success = await campaignStore.createCampaign(payload)

  if (success) {
    isModalOpen.value = false
    campaignForm.value = { title: '', target: '', description: '', duration: '30' }
  }
}
</script>

<template>
  <div class="space-y-8 max-w-6xl mx-auto relative">

    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-5">
      <div>
        <h1 class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
          Urun Dana Komunitas
        </h1>
        <p class="text-sm text-muted-foreground mt-1" style="font-family: 'DM Sans', sans-serif;">
          Bantu wujudkan usaha lokal atau mulai kampanye Anda sendiri.
        </p>
      </div>

      <button
          @click="isModalOpen = true"
          :disabled="!isVerified"
          class="flex items-center justify-center gap-2 bg-primary text-primary-foreground px-5 py-3 rounded-xl font-bold text-sm hover:opacity-90 transition-opacity shadow-md shadow-primary/20 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <Plus :size="18" />
        Buat Kampanye
      </button>
    </div>

    <div class="flex flex-col sm:flex-row gap-3">
      <div class="relative flex-1">
        <Search class="absolute left-4 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
        <input type="text" placeholder="Cari nama usaha atau jenis kampanye..." class="w-full bg-card border border-border rounded-xl py-3 pl-11 pr-4 text-sm focus:outline-none focus:border-primary transition-colors shadow-sm" />
      </div>
      <button class="flex items-center justify-center gap-2 bg-card border border-border text-foreground px-5 py-3 rounded-xl font-bold text-sm hover:bg-muted transition-colors shadow-sm">
        <Filter :size="18" />
        Filter
      </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
          v-for="camp in campaigns"
          :key="camp.id"
          class="bg-card border border-border rounded-3xl overflow-hidden shadow-sm hover:shadow-md transition-shadow group cursor-pointer flex flex-col"
      >
        <div class="relative h-48 overflow-hidden bg-muted">
          <img src="https://images.unsplash.com/photo-1595841696677-6489ff3f8cd1?w=400&h=300&fit=crop" :alt="camp.title" class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105" />
          <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>

          <div class="absolute bottom-3 left-4 right-4 flex justify-between items-end">
            <span class="bg-primary/90 backdrop-blur-sm text-white text-xs font-bold px-3 py-1.5 rounded-lg">
              {{ camp.duration_days > 0 ? `${camp.duration_days} Hari Lagi` : 'Selesai' }}
            </span>
          </div>
        </div>

        <div class="p-5 flex-1 flex flex-col">
          <div class="mb-4 flex-1">
            <p class="text-xs text-muted-foreground mb-1 font-medium">{{ camp.raiser?.name || 'Anggota Komunitas' }}</p>
            <h3 class="text-lg font-bold text-foreground leading-tight" style="font-family: 'Fraunces', serif;">
              {{ camp.title }}
            </h3>
          </div>

          <div class="mb-4">
            <div class="flex justify-between items-center text-xs mb-2" style="font-family: 'DM Mono', monospace;">
              <span class="font-bold text-primary">{{ Math.round(((camp.current_amount || 0) / camp.target_amount) * 100) }}% Terkumpul</span>
            </div>
            <div class="w-full h-2 bg-muted rounded-full overflow-hidden">
              <div
                  class="h-full rounded-full transition-all duration-1000"
                  :class="(camp.current_amount || 0) >= camp.target_amount ? 'bg-accent' : 'bg-primary'"
                  :style="`width: ${Math.min(((camp.current_amount || 0) / camp.target_amount) * 100, 100)}%`"
              ></div>
            </div>
            <div class="flex justify-between items-center text-xs mt-2 text-muted-foreground" style="font-family: 'DM Mono', monospace;">
              <span>Rp {{ (camp.current_amount || 0).toLocaleString('id-ID') }}</span>
              <span>Target: Rp {{ (camp.target_amount || 0).toLocaleString('id-ID') }}</span>
            </div>
          </div>

          <div class="flex items-center justify-between pt-4 border-t border-border">
            <div class="flex items-center gap-1.5 text-xs text-muted-foreground font-medium">
              <Users :size="14" />
              <span>Mendukung</span>
            </div>
            <button
                @click="openFundModal(camp)"
                :disabled="(camp.current_amount || 0) >= camp.target_amount"
                class="text-sm font-bold transition-colors"
                :class="(camp.current_amount || 0) >= camp.target_amount ? 'text-muted-foreground cursor-not-allowed' : 'text-primary hover:text-primary/80'"
            >
              {{ (camp.current_amount || 0) >= camp.target_amount ? 'Terdanai Penuh' : 'Beri Dana' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <Transition
        enter-active-class="transition-opacity duration-300"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition-opacity duration-300"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
    >
      <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-5">
        <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="isModalOpen = false"></div>

        <div class="bg-card border border-border rounded-3xl w-full max-w-lg shadow-2xl relative z-10 flex flex-col max-h-[90vh]">
          <div class="flex items-center justify-between p-6 border-b border-border flex-shrink-0">
            <h2 class="text-xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">Buat Kampanye Baru</h2>
            <button @click="isModalOpen = false" class="p-2 -mr-2 text-muted-foreground hover:bg-muted rounded-xl transition-colors"><X :size="20" /></button>
          </div>

          <div class="p-6 overflow-y-auto">
            <form @submit.prevent="submitCampaign" class="space-y-5">
              <div class="space-y-1.5">
                <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Nama Usaha / Tujuan</label>
                <input v-model="campaignForm.title" type="text" placeholder="Contoh: Pembelian Mesin Jahit Baru" class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm focus:outline-none focus:border-primary transition-colors" required />
              </div>

              <div class="space-y-1.5">
                <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Target Dana (Rp)</label>
                <input v-model="campaignForm.target" type="number" placeholder="Contoh: 1500000" class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm focus:outline-none focus:border-primary transition-colors" required />
              </div>

              <div class="space-y-1.5">
                <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Durasi Kampanye</label>
                <select v-model="campaignForm.duration" class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm focus:outline-none focus:border-primary transition-colors cursor-pointer appearance-none">
                  <option value="15">15 Hari</option>
                  <option value="30">30 Hari</option>
                  <option value="60">60 Hari</option>
                </select>
              </div>

              <div class="space-y-1.5">
                <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Ceritakan Kebutuhan Anda</label>
                <textarea v-model="campaignForm.description" rows="4" placeholder="Jelaskan secara singkat untuk apa dana ini akan digunakan..." class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm focus:outline-none focus:border-primary transition-colors resize-none" required></textarea>
              </div>

              <div class="bg-accent/10 border border-accent/20 rounded-xl p-4 flex gap-3 text-sm text-accent">
                <Info :size="20" class="flex-shrink-0" />
                <p>Kampanye Anda akan ditinjau oleh sistem berbasis komunitas sebelum dipublikasikan.</p>
              </div>

              <button type="submit" class="w-full bg-primary text-primary-foreground font-semibold py-3.5 rounded-xl hover:opacity-90 shadow-md transition-all mt-4">Ajukan Kampanye</button>
            </form>
          </div>
        </div>
      </div>
    </Transition>
  </div>

  <Transition enter-active-class="transition-opacity duration-300" enter-from-class="opacity-0" enter-to-class="opacity-100" leave-active-class="transition-opacity duration-300" leave-from-class="opacity-100" leave-to-class="opacity-0">
    <div v-if="isFundModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="isFundModalOpen = false"></div>
      <div class="bg-card border border-border rounded-3xl w-full max-w-sm shadow-2xl relative z-10 p-6">
        <h2 class="text-xl font-bold mb-2" style="font-family: 'Fraunces', serif;">Dukung Kampanye</h2>
        <p class="text-xs text-muted-foreground mb-4">{{ selectedCampaign?.title }}</p>

        <div class="mb-4">
          <p class="text-xs font-bold text-muted-foreground mb-1">Saldo Anda: {{ formatRupiah(authStore.user?.balance) }}</p>
          <input v-model="fundAmount" type="number" placeholder="Nominal (contoh: 25000)" class="w-full bg-muted/30 border border-border rounded-xl py-3 px-4 text-sm" />
        </div>

        <button @click="handleFunding" :disabled="transactionStore.isLoading" class="w-full bg-primary text-white py-3 rounded-xl font-bold">
          {{ transactionStore.isLoading ? 'Memproses...' : 'Kirim Dana' }}
        </button>
      </div>
    </div>
  </Transition>
</template>