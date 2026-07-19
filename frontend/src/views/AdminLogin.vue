<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ShieldAlert, Mail, Lock, ArrowRight } from '@lucide/vue'
import { useAuthStore } from '../stores/authStore'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const errorMessage = ref('')

const handleAdminLogin = async () => {
  errorMessage.value = ''

  if (!email.value || !password.value) {
    errorMessage.value = 'Email dan Password wajib diisi.'
    return
  }

  const success = await authStore.adminLogin(email.value, password.value)
  if (success) {
    router.push('/dashboard/admin')
  } else {
    errorMessage.value = authStore.error
  }
}
</script>

<template>
  <div class="min-h-screen bg-slate-900 flex items-center justify-center p-4" style="font-family: 'DM Sans', sans-serif;">
    <div class="bg-slate-800 border border-slate-700 rounded-3xl w-full max-w-md p-8 shadow-2xl relative overflow-hidden">

      <div class="absolute -top-10 -right-10 w-40 h-40 bg-blue-500/10 rounded-full blur-2xl"></div>

      <div class="relative z-10">
        <div class="flex justify-center mb-6">
          <div class="w-16 h-16 bg-blue-500/20 text-blue-400 rounded-2xl flex items-center justify-center border border-blue-500/30">
            <ShieldAlert :size="32" />
          </div>
        </div>

        <div class="text-center mb-8">
          <h1 class="text-2xl font-bold text-white mb-2" style="font-family: 'Fraunces', serif;">Akses Terotorisasi</h1>
          <p class="text-sm text-slate-400">Portal masuk khusus Super Administrator Umoja.</p>
        </div>

        <div v-if="errorMessage" class="bg-red-500/10 border border-red-500/50 text-red-400 p-3 rounded-xl mb-6 text-sm font-bold text-center">
          {{ errorMessage }}
        </div>

        <form @submit.prevent="handleAdminLogin" class="space-y-5">
          <div class="space-y-1.5">
            <label class="text-xs font-bold text-slate-400 tracking-wider">EMAIL ADMINISTRATOR</label>
            <div class="relative">
              <Mail class="absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-500" :size="18" />
              <input v-model="email" type="email" placeholder="admin@umoja.id" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl py-3 pl-11 pr-4 text-sm text-white focus:outline-none focus:border-blue-500 transition-colors" required />
            </div>
          </div>

          <div class="space-y-1.5">
            <label class="text-xs font-bold text-slate-400 tracking-wider">PASSWORD</label>
            <div class="relative">
              <Lock class="absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-500" :size="18" />
              <input v-model="password" type="password" placeholder="••••••••" class="w-full bg-slate-900/50 border border-slate-700 rounded-xl py-3 pl-11 pr-4 text-sm text-white focus:outline-none focus:border-blue-500 transition-colors" required />
            </div>
          </div>

          <button type="submit" :disabled="authStore.isLoading" class="w-full bg-blue-600 hover:bg-blue-500 text-white font-bold py-3.5 rounded-xl transition-colors flex items-center justify-center gap-2 mt-4 disabled:opacity-50">
            {{ authStore.isLoading ? 'Memverifikasi...' : 'Masuk Portal' }} <ArrowRight v-if="!authStore.isLoading" :size="18" />
          </button>
        </form>

        <div class="mt-8 text-center">
          <button @click="router.push('/signin')" class="text-slate-500 text-xs hover:text-white transition-colors">
            Kembali ke Portal Warga
          </button>
        </div>
      </div>
    </div>
  </div>
</template>